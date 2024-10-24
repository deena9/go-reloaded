package piscine

func ToUpper(s string) string { 
	if s == "" {
		return s
	}
	upper := []rune(s)
	for i, word := range upper {
		if word >= 'a' && word <= 'z' {
			upper[i] = word - ('a' - 'A')
		}
	}
	return string(upper)
}
