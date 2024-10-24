package main

import (
	"os"
	piscine "piscine/piscine1"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// checkInputs checks that inputs are permissible
func checkInputs(input []string) {
	inputLen := len(input)
	if inputLen != 2 {

		println("Incorrect number of inputs. go run main.go input_file output_file")
		os.Exit((1))
	}
	if _, err := os.Stat(input[0]); err != nil {
		check(err)
	}
}

// readInput reads the input
func readInput(inputName string) string {
	input, err := os.ReadFile(inputName)
	if err != nil {
		println("readfile error", err)
	}
	return (string(input))
}

// writeOutput writes the output
func writeOutput(outputName, result string) {
	err := os.WriteFile(outputName, []byte(result), 0o644)
	check(err)
}

func processIn(in []string) string {
	inLen := len(in)
	if inLen == 0 {
		return ""
	}
	for i := 0; i < inLen; i++ {
		if i > 0 && len(in[i]) > 1 && !piscine.IsConsonant(rune(in[i][0])) && (in[i-1] == "a" || in[i-1] == "A") {
			in[i-1] += "n"
		}
		if in[i] == "(hex)" {
			in[i-1], in[i] = piscine.HexToDec(in[i-1]), ""
		} else if in[i] == "(bin)" {
			in[i-1], in[i] = piscine.BinToDec(in[i-1]), ""
		} else if in[i] == "(up)" {
			in[i-1], in[i] = piscine.ToUpper(in[i-1]), ""
		} else if in[i] == "(low)" {
			in[i-1], in[i] = piscine.ToLower(in[i-1]), ""
		} else if in[i] == "(cap)" {
			in[i-1], in[i] = piscine.ToCapitalize(in[i-1]), ""
		} else if (in[i] == "(cap" || in[i] == "(up" || in[i] == "(low") && i+2 < inLen {
			if in[i+1] == "," && len(in[i+3]) > 1 {
				countInput := in[i+3][:len(in[i+3])-1]
				if piscine.IsNumber(countInput) {
					count, _ := strconv.Atoi(piscine.AtoiBase(countInput, "0123456789"))
					for j := 1; j <= count; j++ {
						if i-j < 0 {
							break
						}
						switch {
						case in[i] == "(cap":
							in[i-j] = piscine.ToCapitalize(in[i-j])
						case in[i] == "(up":
							in[i-j] = piscine.ToUpper(in[i-j])
						case in[i] == "(low":
							in[i-j] = piscine.ToLower(in[i-j])
						}
					}
				}
				in[i], in[i+1], in[i+3] = "", "", ""
			}
		}
	}

	s := in[0]
	iHaveFoundQuot := false
	for i := 1; i < inLen; i++ {
		if in[i] != "" {
			if piscine.IsPunctuationString(in[i]) {
				s = s + in[i]
			} else if piscine.IsQuotationString(in[i]) {
				if !iHaveFoundQuot {
					s = s + " " + in[i]
					if i+1 < inLen && in[i+1] != "" {
						s = s + in[i+1]
					} else {
						s = s + in[i+2]
					}
					i = i + 2
				} else {
					s = s + in[i]
				}
				iHaveFoundQuot = !iHaveFoundQuot
			} else {
				s = s + " " + in[i]
			}
		}
	}
	return s
}

func main() {

	checkInputs(os.Args[1:])
	inputStr := readInput(os.Args[1])
	inputStrArr := piscine.SplitWhiteSpaces(inputStr)
	outputStr := processIn(inputStrArr)
	writeOutput(os.Args[2], outputStr)
}
