package piscine

func IsUpper(s string) bool {
	for _, capOne := range s {
		if capOne < 'A' || capOne > 'Z' {
			return false
		}
	}
	return true
}
