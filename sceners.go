package sceners

import "github.com/Defacto2/sceners/rename"

// Cleaner fixes the malformed string.
// This includes the removal of duplicate spaces, the stripping of incompatible characters.
// The removal of excess whitespace and if found "The " prefix.
func Cleaner(s string) string {
	return rename.Cleaner(s)
}

// Clean deobfuscates the URL path and returns the formatted, human-readable group name.
// The path is expected to be in the format of a URL path without the scheme or domain.
//
// Example:
//
//	Clean("razor-1911") = "Razor 1911"
//	Clean("razor-1911-ampersand-trsi") = "Razor 1911 & Trsi"
//	Clean("north-american-pirate_phreak-association") = "North American Pirate-Phreak Association"
func Clean(path string) string {
	return rename.DeObfuscate(path)
}
