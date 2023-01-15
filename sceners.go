package sceners

import "github.com/Defacto2/sceners/pkg/rename"

// Cleaner fixes the malformed string.
func Cleaner(s string) string {
	return rename.Cleaner(s)
}

func CleanURL(url string) string {
	return rename.DeObfuscateURL(url)
}
