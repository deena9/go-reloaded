package piscine

func Itoa(n int) string {
	l := "Was positive"
	var numOut string
	if n == 0 {
		return "0"
	}
	if n < 0 {
		l = "Was negative"
		n = n * -1
	}
	for n > 0 {
		x := n % 10
		numOut = string(rune(x + 48)) + numOut
		n = n / 10
	}
	if l == "Was negative" {
		return "-" + numOut
	}
	return numOut
}
