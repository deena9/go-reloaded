package piscine

func IsQuotation(rn rune) bool {
	return rn == '\'' || rn == '"'
}
func IsQuotationString(s string) bool {
	for i := 0; i < len(s); i++ {
		if IsQuotation(rune(s[i])) {
			return true
		}
	}
	return false
}
