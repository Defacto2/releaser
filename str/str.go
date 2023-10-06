package str

import "regexp"

// StripChars removes all the incompatible characters that cannot be used for groups and author names.
func StripChars(s string) string {
	const validChars = `[^A-Za-zÀ-ÖØ-öø-ÿ0-9\-,& ]`
	r := regexp.MustCompile(validChars)
	return r.ReplaceAllString(s, "")
}

// StripStart removes the non-alphanumeric characters from the start of the string.
func StripStart(s string) string {
	const latinChars = `[A-Za-z0-9À-ÖØ-öø-ÿ]`
	r := regexp.MustCompile(latinChars)
	f := r.FindStringIndex(s)
	if f == nil {
		return ""
	}
	if f[0] != 0 {
		return s[f[0]:]
	}
	return s
}

// TrimSP removes duplicate spaces from the string.
func TrimSP(s string) string {
	const spaces = `\s+`
	r := regexp.MustCompile(spaces)
	return r.ReplaceAllString(s, " ")
}
