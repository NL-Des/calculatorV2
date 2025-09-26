package answer

import (
	"fmt"
	"os"
	"strconv"
)

// Realise the calcul.
func Calculs(sample1 []rune, symbol rune, sample2 []rune, openBracket bool, endBracket bool) float64 {

	var result float64
	var firstPart float64
	var secondPart float64
	var err error

	// Security check.
	if len(sample1) == 0 {
		fmt.Println("sample1 empty in func Calculs")
		os.Exit(0)
	}
	if len(sample2) == 0 {
		fmt.Println("sample2 empty in func Calculs")
		os.Exit(0)
	}

	// Conversion of format.
	firstPart, err = strconv.ParseFloat(string(sample1), 64)
	if err != nil {
		fmt.Println("firstPart error in func Calculs")
		os.Exit(0)
	}

	secondPart, err = strconv.ParseFloat(string(sample2), 64)
	if err != nil {
		fmt.Println("secondPart error in func Calculs")
		os.Exit(0)
	}

	// Handle the symbol and realise the calcul.
	switch symbol {
	case '+':
		result = firstPart + secondPart
	case '-':
		result = firstPart - secondPart
	case '*':
		result = firstPart * secondPart
	case '/':
		if secondPart == 0 {
			fmt.Println("division by zero impossible, in func Calculs")
			os.Exit(0)
		}
		result = firstPart / secondPart
	}

	return result
}
