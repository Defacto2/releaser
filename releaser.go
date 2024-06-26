// Package releaser provides functions for cleaning and formatting the scene names and titles.
package releaser

import (
	"strings"

	"github.com/Defacto2/releaser/fix"
	"github.com/Defacto2/releaser/name"
)

// Cell formats the string to be used as a cell in a database table.
// This includes the removal of duplicate spaces and the stripping of incompatible characters.
// The removal of excess whitespace and if found "The " prefix from BBS and FTP named sites.
//
// Example:
//
//	Cell("  Defacto2  demo  group.") = "DEFACTO2 DEMO GROUP"
//	Cell("the x bbs") = "X BBS"
//	Cell("defacto2.net") = "DEFACTO2NET"
func Cell(s string) string {
	x := fix.TrimSP(s)
	x = fix.StripChars(x)
	x = fix.StripStart(x)
	x = strings.TrimSpace(x)
	x = fix.TrimThe(x)
	return fix.Cell(x)
}

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
	x := fix.TrimSP(s)
	x = fix.StripChars(x)
	x = fix.StripStart(x)
	x = strings.TrimSpace(x)
	x = fix.TrimThe(x)
	return fix.Format(x)
}

// Humanize deobfuscates the URL path and returns the formatted, human-readable group name.
// The path is expected to be in the format of a URL path without the scheme or domain.
// If the URL path contains invalid characters then an empty string is returned.
//
// Example:
//
//	Humanize("defacto2") = "Defacto2"
//	Humanize("razor-1911-demo") = "Razor 1911 Demo"
//	Humanize("razor-1911-demo-ampersand-skillion") = "Razor 1911 Demo & Skillion"
//	Humanize("north-american-pirate_phreak-association") = "North American Pirate-Phreak Association"
//	Humanize("razor-1911-demo*trsi") = "Razor 1911 Demo, TRSi"
//	Humanize("razor-1911-demo#trsi") = "" // invalid # character
func Humanize(path string) string {
	p := name.Path(strings.ToLower(path))
	if special := p.String(); special != "" {
		return special
	}
	s, err := name.Humanize(p)
	if err != nil {
		return ""
	}
	return Clean(s)
}

// Humanize deobfuscates the URL path and applies [releaser.Humanize].
// In addition, the humanized name is formatted to be used as a link description.
// If the URL path contains invalid characters then an empty string is returned.
func Link(path string) string {
	s := Humanize(path)
	return strings.ReplaceAll(s, ", ", " + ")
}

// Obfuscate formats the string to be used as a URL path.
//
// Example:
//
//	Obfuscate("ACiD Productions") = "acid-productions"
//	Obfuscate("Razor 1911 Demo & Skillion") = "razor-1911-demo-ampersand-skillion"
//	Obfuscate("TDU-Jam!") = "tdu_jam"
func Obfuscate(s string) string {
	c := Clean(s)
	x := name.Obfuscate(c)
	return string(x)
}
