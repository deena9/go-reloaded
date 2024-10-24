package piscine

//this only checks for numerical characters
func IsNumber(s string) bool {
	for _, number := range s {
		if number < '0' || number > '9' {
			return false
		}
	}
	return true
}
