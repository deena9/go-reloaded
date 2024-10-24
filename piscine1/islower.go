package piscine

func IsLower(s string) bool {
	for _, lowcase := range s {
		if lowcase < 'a' || lowcase > 'z' {
			return false
		}
	}
	return true
}
