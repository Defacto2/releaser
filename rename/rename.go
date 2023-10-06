// Package rename provides functions for cleaning and formatting strings of known words and group names.
package rename

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const space = " "

// Amp formats the special ampersand (&) character in the string
// to be usable with a URL path in use by the group.
//
// Example:
//
//	Amp("hello&&world") = "hello & world"
func Amp(s string) string {
	if !strings.Contains(s, "&") {
		return s
	}
	x := s
	trimDupes := regexp.MustCompile(`\&+`)
	x = trimDupes.ReplaceAllString(x, "&")

	trimPrefix := regexp.MustCompile(`^\&+`)
	x = trimPrefix.ReplaceAllString(x, "")

	trimSuffix := regexp.MustCompile(`\&+$`)
	x = trimSuffix.ReplaceAllString(x, "")

	addWhitespace := regexp.MustCompile(`(\S)\&(\S)`) // \S matches any character that's not whitespace
	x = addWhitespace.ReplaceAllString(x, "$1 & $2")
	return x
}

// Case returns the exact syntax of the known named group.
//
// Example:
//
//	Case("beer") = "BEER"
//	Case("lkcc") = "LKCC"
//	Case("noclass") = "NoClass"
func Case(name string) string {
	s := strings.ToLower(name)
	switch s {
	// all uppercase full groups
	case "2nd2none bbs",
		"3wa bbs",
		"acb bbs",
		"anz ftp",
		"beer",
		"bcp bbs",
		"ckc bbs",
		"cnx ftp",
		"core",
		"crsiso",
		"cwl bbs",
		"dv8 bbs",
		"es bbs",
		"fic bbs",
		"lkcc",
		"lms bbs",
		"ls bbs",
		"lsdiso",
		"lpc bbs",
		"lta bbs",
		"mor ftp",
		"msv ftp",
		"new dtl",
		"nsdap",
		"nos ftp",
		"og bbs",
		"okc bbs",
		"pmc bbs",
		"pp bbs",
		"ppps bbs",
		"pox ftp",
		"psi bbs",
		"qed bbs",
		"scf ftp",
		"scsi ftp",
		"swat",
		"tiw bbs",
		"tbb ftp",
		"tcsm bbs",
		"tfz 2 bbs",
		"tog ftp",
		"top ftp",
		"tph-qqt",
		"tph-qqt ftp",
		"trt 2001 bbs",
		"tsi bbs",
		"tsc bbs",
		"uct bbs",
		"u4ea ftp",
		"zoo ftp":
		return strings.ToUpper(s)
	case "scenet":
		// all lowercase full groups
		return strings.ToLower(s)
	}
	return mixCase(s)
}

// mixCase returns the exact mix-case syntax of the named group.
func mixCase(name string) string {
	s := strings.ToLower(name)
	switch s {
	case "79th trac":
		return "79th TRAC"
	case "biased":
		return "bIASED"
	case "cybermail":
		return "CyberMail"
	case "dreadloc":
		return "DREADLoC"
	case "drm ftp":
		return "dRM FTP"
	case "dst ftp":
		return "dst FTP"
	case "dvtiso":
		return "DVTiSO"
	case "excretion anarchy":
		return "eXCReTION Anarchy"
	case "htbzine":
		return "HTBZine"
	case "ice weekly newsletter":
		return "iCE Weekly Newsletter"
	case "mci escapes":
		return "mci escapes"
	case "noclass":
		return "NoClass"
	case "nofx bbs":
		return "NoFX BBS"
	case "pjs tower":
		return "PJs Tower BBS"
	case "pocketheaven":
		return "PocketHeaven"
	case "ptl club":
		return "PTL Club"
	case "rhvid":
		return "RHViD"
	case "rzsoft ftp":
		return "RZSoft FTP"
	case "trsi":
		return "TRSi"
	case "tsg ftp":
		return "tSG FTP"
	case "unreal magazine":
		return "UnReal Magazine"
	case "vdr lake ftp":
		return "VDR Lake FTP"
	case "xquizit ftp":
		return "XquiziT FTP"
	}
	// rename groups (demozoo vs defacto2 formatting etc.)
	switch s {
	case "2000 ad":
		return "2000AD"
	case "hashx":
		return "Hash X"
	case "phoenixbbs":
		return "Phoenix BBS"
	}
	return ""
}

// Connect formats common connecting word as the w string based on its position in a words slice.
func Connect(w string, position, last int) string {
	const first = 0
	if position == first || position == last {
		return ""
	}
	switch strings.ToLower(w) {
	case "a", "as", "and", "at", "by", "el", "of", "for", "from", "in", "is", "or", "tha",
		"the", "to", "with":
		return strings.ToLower(w)
	}
	return ""
}

// Fix formats the w string based on its position in the words slice.
// The position is the index of the word in the words slice.
// The last is the index of the last word in the words slice.
func Fix(w string, position, last int) string {
	if fix := Connect(w, position, last); fix != "" {
		return fix
	}
	if fix := Word(w); fix != "" {
		return fix
	}
	title := cases.Title(language.English, cases.NoLower)
	if fix := PreSuffix(w, title); fix != "" {
		return fix
	}
	if fix := Sequence(w, position); fix != "" {
		return fix
	}
	return title.String(w)
}

// FixHyphen applies [rename.Fix] to hyphenated words.
func FixHyphen(w string) string {
	const hyphen = "-"
	if !strings.Contains(w, hyphen) {
		return ""
	}
	compounds := strings.Split(w, hyphen)
	last := len(compounds) - 1
	for i, word := range compounds {
		compounds[i] = Fix(word, i, last)
	}
	return strings.Join(compounds, hyphen)
}

// Format returns a copy of s with custom formatting.
// Certain words and known acronyms will be upper cased, lower cased or title cased.
// Known named groups will be returned in their special casing.
// Trailing dots will be removed.
//
// Example:
//
//	Format("hello world.") = "Hello World"
//	Format("the 12am group.") = "The 12AM Group"
func Format(s string) string {
	const (
		acronym   = 3
		separator = ", "
	)
	if len(s) <= acronym {
		return strings.ToUpper(s)
	}
	groups := strings.Split(s, separator)
	for j, group := range groups {
		g := strings.ToLower(strings.TrimSpace(group))
		g = Amp(g)
		if fix := Case(g); fix != "" {
			groups[j] = fix
			continue
		}

		words := strings.Split(g, space)
		last := len(words) - 1
		for i, word := range words {
			word = TrimDot(word)
			if fix := FixHyphen(word); fix != "" {
				words[i] = fix
				continue
			}
			words[i] = Fix(word, i, last)
		}
		groups[j] = strings.Join(words, space)
	}
	return strings.Join(groups, separator)
}

// PreSuffix formats the w string if a known prefix or suffix is found.
// The title caser needs to be a language-specific title casing.
//
// Example:
//
//	PreSuffix("12am", cases.Title(language.English, cases.NoLower)) = "12AM"
func PreSuffix(s string, title cases.Caser) string {
	w := strings.ToLower(s)
	switch {
	case strings.HasSuffix(w, "ad"):
		x := strings.TrimSuffix(w, "ad")
		if val, err := strconv.Atoi(x); err == nil {
			return fmt.Sprintf("%dAD", val)
		}
	case strings.HasSuffix(w, "bc"):
		x := strings.TrimSuffix(w, "bc")
		if val, err := strconv.Atoi(x); err == nil {
			return fmt.Sprintf("%dBC", val)
		}
	case strings.HasSuffix(w, "am"):
		x := strings.TrimSuffix(w, "am")
		if val, err := strconv.Atoi(x); err == nil {
			return fmt.Sprintf("%dAM", val)
		}
	case strings.HasSuffix(w, "pm"):
		x := strings.TrimSuffix(w, "pm")
		if val, err := strconv.Atoi(x); err == nil {
			return fmt.Sprintf("%dPM", val)
		}
	case strings.HasSuffix(w, "dox"):
		val := strings.TrimSuffix(w, "dox")
		return fmt.Sprintf("%sDox", title.String(val))
	case strings.HasSuffix(w, "fxp"):
		val := strings.TrimSuffix(w, "fxp")
		return fmt.Sprintf("%sFXP", title.String(val))
	case strings.HasSuffix(w, "iso"):
		val := strings.TrimSuffix(w, "iso")
		return fmt.Sprintf("%sISO", title.String(val))
	case strings.HasSuffix(w, "nfo"):
		val := strings.TrimSuffix(w, "nfo")
		return fmt.Sprintf("%sNFO", title.String(val))
	case strings.HasPrefix(w, "pc-"):
		val := strings.TrimPrefix(w, "pc-")
		return fmt.Sprintf("PC-%s", title.String(val))
	case strings.HasPrefix(w, "lsd"):
		val := strings.TrimPrefix(w, "lsd")
		return fmt.Sprintf("LSD%s", title.String(val))
	}
	return ""
}

// Sequence formats the w string if it is the first word in the words slice.
func Sequence(w string, i int) string {
	if i != 0 {
		return ""
	}
	switch w { //nolint:gocritic
	case "inc":
		// note: Format() applies UPPER to all 3 letter or smaller words
		return strings.ToUpper(w)
	}
	return ""
}

// TrimDot removes a trailing dot from s.
//
// Example:
//
//	TrimDot("hello.") = "hello"
//	TrimDot("hello..") = "hello."
func TrimDot(s string) string {
	const short = 2
	if len(s) < short {
		return s
	}
	if l := s[len(s)-1:]; l == "." {
		return s[:len(s)-1]
	}
	return s
}

// TrimThe drops "The" prefix whenever the named site ends with "BBS" or "FTP".
// This is to avoid the same site name being both "The X BBS" and "X BBS".
//
// Example:
//
//	TrimThe("The X BBS") = "X BBS"
//	TrimThe("The X") = "The X"
func TrimThe(name string) string {
	const short = 2
	a := strings.Split(name, space)
	if len(a) < short {
		return name
	}
	l := strings.ToUpper(a[len(a)-1])
	if strings.EqualFold(a[0], "the") && (l == "BBS" || l == "FTP") {
		return strings.Join(a[1:], space) // drop "the" prefix
	}
	return name
}

// Word applies upper casing to known acronyms, initalisms and abbreviations.
// And lower casing to ordinal numbers 1st through to 13th.
// Otherwise it returns an empty string.
//
// Example:
//
//	Word("1sT") = "1st"
//	Word("iso") = "ISO"
func Word(s string) string {
	x := strings.ToLower(s)
	switch x {
	case "1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th", "9th",
		"10th", "11th", "12th", "13th":
		return strings.ToLower(s)
	case "3d", "abc", "acdc", "ad", "am", "amf", "ansi", "asm", "au", "bbc", "bbs", "bc",
		"cd", "cgi", "diz", "dox", "eu", "faq", "fbi", "fm", "ftp", "fr", "fx", "fxp",
		"gbc", "gif", "hq", "id", "ii", "iii", "iso", "kgb", "mp3", "pc", "pcb", "pcp",
		"pda", "pm", "psx", "pwa", "rom", "rpm", "ssd", "st", "tnt", "tsr", "ufo", "uk",
		"us", "usa", "uss", "ussr", "vcd", "whq", "xxx":
		return strings.ToUpper(s)
	case "7of9":
		return strings.ToLower(s)
	default:
		return ""
	}
}
