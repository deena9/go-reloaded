package piscine

func IsConsonant(rn rune) bool {
	if !(rn == 'a' || rn == 'A' || rn == 'e' || rn == 'E' || rn == 'i' || rn == 'I' || rn == 'o' || rn == 'O' || rn == 'u' || rn == 'U') {
		return true
	}
	return false
}
