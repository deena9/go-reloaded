package piscine

import "strconv"

func AtoiBase(s string, base string) string {
	baseLen := len(base)
	sLen := len(s)
	match := false
	result := 0
	var final string
	for i := 0; i < sLen; i++ {
		match = false
		for j := 0; j < baseLen; j++ {
			if s[i] == base[j] {
				match = true
				result = result*baseLen + j
			}
		}
		if !match {
			return "Invalid Input"
		}
	}
	final = strconv.Itoa(result)
	return final
}
func BinToDec(in string) string {
	return AtoiBase(in, "01")
}
func HexToDec(in string) string {
	return AtoiBase(ToLower(in), "0123456789abcdef")
}

