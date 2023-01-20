package sceners

import "github.com/Defacto2/sceners/pkg/rename"

// Cleaner fixes the malformed string.
// This includes the removal of duplicate spaces, the stripping of incompatiable characters.
// The removal of excess whitespace and if found "The " prefix.
func Cleaner(s string) string {
	return rename.Cleaner(s)
}

// CleanURL deobfuscates the url and returns a human-readable and formatted group name.
func CleanURL(url string) string {
	return rename.DeObfuscateURL(url)
}
