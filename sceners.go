// Package sceners provides functions for cleaning and formatting the scene names and titles.
package sceners

import (
	"regexp"
	"strings"

	"github.com/Defacto2/sceners/rename"
	"github.com/Defacto2/sceners/str"
)

// Clean fixes the malformed string.
// This includes the removal of duplicate spaces and the stripping of incompatible characters.
// The removal of excess whitespace and if found "The " prefix from BBS and FTP named sites.
//
// Example:
//
//	Clean("  Defacto2  demo  group.") = "Defacto2 Demo Group"
//	Clean("the x bbs") = "X BBS"
//	Clean("The X Ftp") = "X FTP"
func Clean(s string) string {
	x := str.TrimSP(s)
	x = str.StripChars(x)
	x = str.StripStart(x)
	x = strings.TrimSpace(x)
	x = rename.TrimThe(x)
	return rename.Format(x)
}

// Humanize deobfuscates the URL path and returns the formatted, human-readable group name.
// The path is expected to be in the format of a URL path without the scheme or domain.
//
// Example:
//
//	Humanize("defacto2") = "Defacto2"
//	Humanize("razor-1911-demo") = "Razor 1911 Demo"
//	Humanize("razor-1911-demo-ampersand-skillion") = "Razor 1911 Demo & Skillion"
//	Humanize("north-american-pirate_phreak-association") = "North American Pirate-Phreak Association"
//	Humanize("razor-1911-demo*trsi") = "Razor 1911 Demo, TRSi"
func Humanize(path string) string {
	s := strings.TrimSpace(strings.ToLower(path))
	re := regexp.MustCompile(`-ampersand-`)
	s = re.ReplaceAllString(s, " & ")
	re = regexp.MustCompile(`-`)
	s = re.ReplaceAllString(s, " ")
	re = regexp.MustCompile(`_`)
	s = re.ReplaceAllString(s, "-")
	re = regexp.MustCompile(`\*`)
	s = re.ReplaceAllString(s, ", ")
	return Clean(s)
}
