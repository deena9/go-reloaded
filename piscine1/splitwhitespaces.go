package piscine

func SplitWhiteSpaces(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	result := []string{}
	temporary := ""

	for i := range s {
		if isWhiteSpace(rune(s[i])) {
			result = append(result, temporary)
			temporary = ""
		} else if IsPunctuation(rune(s[i])) || IsQuotation(rune(s[i])) {
			result = append(result, temporary, (string(s[i])))
			temporary = ""
			continue
		} else {
		temporary += string(s[i])
		}
	}
	result = append(result, temporary)

	return result
}

func isWhiteSpace(char rune) bool {
	return char == ' ' || char == '\n' || char == '\t'
}
