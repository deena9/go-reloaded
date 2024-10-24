package piscine

func IsPunctuation(rn rune) bool {
	if rn == '.' || rn == ',' || rn == '!' || rn == '?' || rn == ':' || rn == ';' {
		return true
	}
	return false
}
func IsPunctuationString(s string) bool {
	for i := 0; i < len(s); i++ {
		if !IsPunctuation(rune(s[i])) {
			return false
		}
	}
	return true
}
