package piscine

func ToLower(s string) string {
	if s == "" {
		return s
	}
	sRune := []rune(s)
	for i := 0; i < len(s); i++ {
		if IsUpper(string(sRune[i])) {
			sRune[i] = sRune[i] + 32
		}
	}
	return string(sRune)
}
