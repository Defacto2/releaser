// Package name provides functionality for handling the URL path of a releaser.
// It contains a type Path that represents the URL path of a releaser and methods
// to validate and retrieve the well-known styled name of the releaser.
// It also contains a map of releasers and their well-known styled names.
package name

import (
	"errors"
	"regexp"
	"strings"
)

var ErrInvalidPath = errors.New("the path contains invalid characters")

// A Path is the partial URL path of the releaser.
type Path string

// String returns the well-known styled name of the releaser if it exists in the
// names, lowercase or uppercase lists. Otherwise it returns an empty string.
//
// Example:
//
//	name.Path("acid-productions").String() = "ACiD Productions"
//	name.Path("razor-1911").String() = "" // unlisted
func (path Path) String() string {
	p := Path(strings.ToLower(string(path)))
	list := Special()
	if _, ok := list[p]; ok {
		return list[p]
	}
	return ""
}

// Valid returns true if the URL path uses valid characters.
// Valid URL paths are all lowercase and contain only alphanumeric characters, dashes, underscores,
// ampersands and asterisks.
//
// Example:
//
//	name.Path("acid-productions").Valid() = true
//	name.Path("acid-productions!").Valid() = false
func (path Path) Valid() bool {
	re := regexp.MustCompile(`^[a-z0-9\&\-_\*]+$`)
	return re.MatchString(string(path))
}

// A List is a map of releasers and their well-known styled names.
type List map[Path]string

/*
The following list of styled names is used to test the Path type and its methods.

Stylized names should avoid using special characters that may get encoded in the URL
or converted due to their special uses within the name.
*/
var names = List{
	"2000-ad":                               "2000AD",
	"79th-trac":                             "79th TRAC",
	"acid-productions":                      "ACiD Productions",
	"biased":                                "bIASED",
	"binpda":                                "BiNPDA",
	"coop":                                  "TDT / TRSi",
	"core":                                  "CoRE",
	"copycats-inc":                          "CopyCats Inc",
	"coreutil":                              "The Utility Division of CORE",
	"crackpl":                               "CrackPL",
	"cybermail":                             "CyberMail",
	"dbcdemo":                               "DBCDemo",
	"dreadloc":                              "DREADLoC",
	"dumptruck":                             "dumpTruck",
	"defacto2net":                           "Defacto2 website",
	"drm-ftp":                               "dRM FTP",
	"dst-ftp":                               "dst FTP",
	"dvniso":                                "DVNiSO",
	"dvtiso":                                "DVTiSO",
	"esp-pirates":                           "ESP Pirates",
	"extreme-net":                           "ExtremeNET",
	"excretion-anarchy":                     "eXCReTION Anarchy",
	"hashx":                                 "Hash X",
	"htbzine":                               "HTBZine",
	"linezer0":                              "LineZer0",
	"ice-weekly-newsletter":                 "iCE Weekly Newsletter",
	"icon":                                  "iCON",
	"imars":                                 "iMARS",
	"jrp":                                   "Japanese Release Project",
	"oneup":                                 "OneUp",
	"orion":                                 "ORiON",
	"mmi":                                   "MMi",
	"mp2k":                                  "MP2K",
	"nc_17":                                 "NC-17",
	"nicjr":                                 "NicJr",
	"noclass":                               "NoClass",
	"nofx-bbs":                              "NoFX BBS",
	"nukethis":                              "NukeThis",
	"numbers":                               "The Numbers",
	"nrp":                                   "NoRePack",
	"paradox":                               "Paradox",
	"phoenixbbs":                            "Phoenix BBS",
	"pjs-tower-bbs":                         "PJs Tower BBS",
	"playme":                                "PlayMe",
	"pocketheaven":                          "PocketHeaven",
	"psico":                                 "PSiCO",
	"ptl-club":                              "PTL Club",
	"pouet":                                 "Pouët",
	"risciso":                               "RISCiSO",
	"sda-review":                            "SDA Review",
	"seek-n-destroy":                        "Seek n Destroy",
	"sma-posse":                             "SMA Posse",
	"shitonlygerman":                        "ShitOnlyGerman",
	"software-pirates-inc":                  "Software Pirates Inc",
	"sprint":                                "$print",
	"surprise-productions":                  "Surprise! Productions",
	"syndicate":                             "$yndicate",
	"razordox":                              "RazorDOX",
	"rhvid":                                 "RHViD",
	"rzsoft-ftp":                            "RZSoft FTP",
	"tdu_jam":                               "TDU Jam!",
	"team-xtx":                              "Team XTX",
	"thg-fx":                                "THG-FX",
	"tft-team":                              "TFT Team",
	"tpinc":                                 "TPiNC",
	"trsi":                                  "TRSi",
	"tristar-ampersand-red-sector-inc":      "Tristar & Red Sector Inc",
	"the-dvdr-releasing-standards":          "The DVDR Releasing Standards",
	"the-firm":                              "The FiRM",
	"tsg-ftp":                               "tSG FTP",
	"tport":                                 "tPORt",
	"underpl":                               "UnderPL",
	"unreal-magazine":                       "UnReal Magazine",
	"united-software-association*fairlight": "United Software Association + Fairlight PC Division",
	"vdr-lake-ftp":                          "VDR Lake FTP",
	"xdb":                                   "X-db",
	"xquizit-ftp":                           "XquiziT FTP",
}

var lowercase = []string{
	"mci-escapes",
	"scenet",
}

var uppercase = []string{
	"2nd2none-bbs",
	"3wa-bbs",
	"acb-bbs",
	"anz-ftp",
	"beer",
	"bcp-bbs",
	"cusa",
	"ckc-bbs",
	"cnx-ftp",
	"core",
	"crsiso",
	"cwl-bbs",
	"dv8-bbs",
	"es-bbs",
	"fake",
	"fate",
	"fic-bbs",
	"hasp",
	"lkcc",
	"lms-bbs",
	"ls-bbs",
	"lsdiso",
	"lpc-bbs",
	"lta-bbs",
	"lube",
	"mor-ftp",
	"msv-ftp",
	"new-dtl",
	"nsdap",
	"nohk",
	"nos-ftp",
	"og-bbs",
	"okc-bbs",
	"pe*trsi*tdt",
	"petra",
	"pmc-bbs",
	"pp-bbs",
	"ppps-bbs",
	"pox-ftp",
	"ps5b",
	"psi-bbs",
	"qed-bbs",
	"scum",
	"swag",
	"scf-ftp",
	"scsi-ftp",
	"shot",
	"swat",
	"tiw-bbs",
	"tbb-ftp",
	"tcsm-bbs",
	"tfz-2-bbs",
	"tog-ftp",
	"top-ftp",
	"tph-qqt",
	"tph-qqt-ftp",
	"trt-2001-bbs",
	"tsi-bbs",
	"tsc-bbs",
	"uct-bbs",
	"u4ea-ftp",
	"x_ess",
	"zoo-ftp",
}

// Special returns the list of styled names that use special mix or all lower or upper casing.
func Special() List {
	l := make(List, len(names)+len(lowercase)+len(uppercase))
	for k, v := range Names() {
		l[k] = v
	}
	for k, v := range Lower() {
		l[k] = v
	}
	for k, v := range Upper() {
		l[k] = v
	}
	return l
}

// Names returns the list of well-known styled names.
func Names() List {
	return names
}

// Lower returns the list of styled names that use all lowercasing.
func Lower() List {
	l := make(List, len(lowercase))
	for _, s := range lowercase {
		p := Path(s)
		x, _ := Humanize(p)
		l[p] = strings.ToLower(x)
	}
	return l
}

// Upper returns the list of styled names that use all uppercasing.
func Upper() List {
	l := make(List, len(uppercase))
	for _, s := range uppercase {
		p := Path(s)
		x, _ := Humanize(p)
		l[p] = strings.ToUpper(x)
	}
	return l
}

// Humanize deobfuscates the URL path and returns the formatted, human-readable group name.
// If the URL path contains invalid characters then an error is returned.
func Humanize(path Path) (string, error) {
	if !path.Valid() {
		return "", ErrInvalidPath
	}

	s := strings.ToLower(string(path))

	re := regexp.MustCompile(`-ampersand-`)
	s = re.ReplaceAllString(s, " & ")

	re = regexp.MustCompile(`-`)
	s = re.ReplaceAllString(s, " ")

	re = regexp.MustCompile(`_`)
	s = re.ReplaceAllString(s, "-")

	re = regexp.MustCompile(`\*`)
	s = re.ReplaceAllString(s, ", ")
	return s, nil
}

// Obfuscate formats the named string to be used as a URL path.
//
// Example:
//
//	Obfuscate("ACiD Productions") = "acid-productions"
//	Obfuscate("Razor 1911 Demo & Skillion") = "razor-1911-demo-ampersand-skillion"
//	Obfuscate("TDU-Jam!") = "tdu_jam"
func Obfuscate(name string) Path {
	s := strings.TrimSpace(strings.ToLower(name))

	re := regexp.MustCompile(`[^a-z0-9\&\-\,\ ]`)
	s = re.ReplaceAllString(s, "")

	// the order of these expressions is critical

	re = regexp.MustCompile(`-`)
	s = re.ReplaceAllString(s, "_")

	re = regexp.MustCompile(` \& `)
	s = re.ReplaceAllString(s, "-ampersand-")

	re = regexp.MustCompile(`\, `)
	s = re.ReplaceAllString(s, "*")

	re = regexp.MustCompile(` `)
	s = re.ReplaceAllString(s, "-")

	return Path(s)
}
