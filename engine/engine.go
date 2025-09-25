package engine

import (
	answer "calculator/resolvemachine/calculanswer"
	"calculator/resolvemachine/rules"
	"fmt"
)

var result int
var err error

func FirstStep(calcul []rune) {

	length := len(calcul)

	var sample1 []rune
	var sample2 []rune
	var firstMinusCoord int
	symbol := ' '
	var symbolOpenBracket rune
	var symbolEndBracket rune
	//end := false
	//var save []rune
	//var savefirstMinus bool

	// Conditions, d'avant calculs.
	// result = sample (symbol) sample2
	for i := 0; i < length; i++ {
		caracter := calcul[i]

		// If the first caracter(i==0) is a minus.
		sample1, firstMinusCoord = rules.MinusFirst(caracter, i, sample1, firstMinusCoord)

		// If we meet a number between 0 and 9 and without symbol.
		// We stock it.
		sample1 = rules.FirstPart(sample1, caracter, symbol)

		// If between 0 and 9, and already with a symbol in memory.
		sample2 = rules.SecondPart(sample2, caracter, symbol)

		// If we meet a symbol and we haven't a symbol in memory.
		symbol = rules.SymbolGestion(caracter, symbol, i, firstMinusCoord)

		// Handle brackets.
		symbolOpenBracket, symbolEndBracket = rules.BracketGestion(caracter, symbolOpenBracket, symbolEndBracket)

		// if a second symbol is detected, we launch the calcul
		if i == length-1 /* !(calcul[i] >= '0' && calcul[i] <= '9') && symbol != ' ' */ {
			result = answer.Calculs(sample1, symbol, sample2, symbolOpenBracket, symbolEndBracket)
		}
	}

	fmt.Println(result)
}

/*
		// If end of the table of rune of calcul.
		if i+1 == len(calcul) {
			end = true
		}

		// When we meet the end of the calcul.
		switch end {
		case (calcul[i] >= '0' && calcul[i] <= '9'): // If we meet an number.
			if !savefirstMinus {
				firstMinus = true
			}
			sample2 = append(sample2, calcul[i])
			result, err = secondStep(sample1, sample2, symbol, firstMinus, symbolOpenBracket, symbolEndBracket, end, result)
		case calcul[i] == ')': // If we meet an end bracket.
			if !savefirstMinus {
				firstMinus = true
			}
			symbolEndBracket = ')'
			result, err = secondStep(sample1, sample2, symbol, firstMinus, symbolOpenBracket, symbolEndBracket, end, result)
		}

		// If second symbol is detected.
		switch !(calcul[i] >= '0' && calcul[i] <= '9') {
		case i < length-1 && (calcul[i+1] >= '0' && calcul[i+1] <= '9'): // If we are not at the end and i+1 is a number.
			if !savefirstMinus {
				firstMinus = true
			}
			result, err = secondStep(sample1, sample2, symbol, firstMinus, symbolOpenBracket, symbolEndBracket, end, result)
			sample1 = nil
			sample2 = nil
			symbol = ' '
			firstMinus = false
		case calcul[i] == '(': // If we are in face of an open bracket.
			if sample1 != nil && sample2 != nil { // If we have data in sample1 and sample2, realise the calcul.
				if !savefirstMinus {
					firstMinus = true
				}
				symbolOpenBracket = ' '
				result, err = secondStep(sample1, sample2, symbol, firstMinus, symbolOpenBracket, symbolEndBracket, end, result)
				symbolOpenBracket = '('
				sample1 = nil
				sample2 = nil
				symbol = ' '
				firstMinus = false
			} else if sample1 != nil && sample2 == nil { // If we have only sample1, we save it.
				save = sample1
				savefirstMinus = firstMinus
				symbolOpenBracket = '('
				sample1 = nil
				sample2 = nil
				symbol = ' '
			} else if sample1 == nil && sample2 == nil && i == 0 && i+1 < length && calcul[i+1] == '-' { // If we haven't data but an minus in the begining.
				firstMinus = true
			}
		case calcul[i] == ')': // If we are in face of an closed bracket.
			switch save == nil {
			case sample1 != nil && sample2 == nil:
				sample2 = sample1
				sample1 = []rune(string(result))
				symbol = '*'
				result, err = secondStep(sample1, sample2, symbol, firstMinus, symbolOpenBracket, symbolEndBracket, end, result)
			case sample1 != nil && sample2 != nil:
				result, err = secondStep(sample1, sample2, symbol, firstMinus, symbolOpenBracket, symbolEndBracket, end, result)
			}
			switch save != nil {
			case sample1 != nil && sample2 == nil:
				sample2 = sample1
				sample1 = save
				symbol = '*'
				result, err = secondStep(sample1, sample2, symbol, firstMinus, symbolOpenBracket, symbolEndBracket, end, result)
			case sample1 != nil && sample2 != nil:
				result, err = secondStep(sample1, sample2, symbol, firstMinus, symbolOpenBracket, symbolEndBracket, end, result)
				temporary, _ := strconv.Atoi(string(save))
				result = temporary * result
			}
		}
	}
	fmt.Println(result)
} */
