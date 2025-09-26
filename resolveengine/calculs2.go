package engine

import (
	"errors"
	"fmt"
	"strconv"
)

func secondStep(sample1 []rune, sample2 []rune, symbol rune, firstMinus bool, openBracket rune, endBracket rune, end bool, result int) (int, error) {

	// Handle Error
	// Conversion from Runes to Integer of sample1.
	firstPart, err := strconv.Atoi(string(sample1))
	if err != nil {
		return 0, errors.New("error in calculs2, on firstPart")
	} else if len(sample1) == 0 {
		return 0, errors.New("error in calculs2, nothing in sample1")
	}
	if firstMinus { // If first caracter is an minus.
		firstPart = -firstPart
	}

	// Conversion from Runes to Integer of sample2.
	secondPart, err := strconv.Atoi(string(sample2))
	if err != nil {
		return 0, errors.New("error in calculs2, on secondPart")
	} else if len(sample2) == 0 {
		return 0, errors.New("error in calculs2, nothing in sample2")
	}

	fmt.Println(symbol)

	switch symbol {
	case '+':
		return result + firstPart + secondPart, nil
	case '-':
		return result + (firstPart - secondPart), nil
	case '*':
		return result + (firstPart * secondPart), nil
	case '/':
		if secondPart == 0 {
			return 0, errors.New("division by zero impossible")
		}
		return result + (firstPart / secondPart), nil
	default:
		return 0, errors.New("operator not handle")
	}

	return result, nil
}
