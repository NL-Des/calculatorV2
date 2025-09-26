package engine

import (
	answer "calculator/resolveengine/calculanswer"
	"calculator/resolveengine/rules"
	"fmt"
)

var result float64 // float64 than int, to be able to handle numbers after float.

// Logic of the system.
// Ex : result = sample (symbol) sample2

func Solution(calcul []rune) {

	length := len(calcul)

	var sample1 []rune
	var sample2 []rune
	var firstMinusCoord int
	symbol := ' '
	openBracket := false
	endBracket := false
	// var openBracketID int
	// var endBracketID int

	for i := 0; i < length; i++ {
		caracter := calcul[i]

		// Tentative to handle the question of the brackets.
		// Detection brackets.
		// openBracket, endBracket, openBracketID, endBracketID = rules.BracketDetection(i, caracter, openBracket, endBracket, openBracketID, endBracketID)

		// Bracket handler.
		//rules.HandleBrackets(i, caracter, openBracket, endBracket, sample1, sample2, symbol, firstMinusCoord)

		// If the first caracter(i==0) is a minus.
		sample1, firstMinusCoord = rules.MinusFirst(caracter, i, sample1, firstMinusCoord)

		// If we meet a number between 0 and 9 and without symbol.
		sample1, sample2 = rules.NumberParts(sample1, sample2, caracter, symbol)

		// If we meet a symbol.
		symbol = rules.SymbolGestion(caracter, symbol, i, firstMinusCoord)

		// if we meet the end of the calcul.
		if i == length-1 /* !(calcul[i] >= '0' && calcul[i] <= '9') && symbol != ' ' */ {
			result = answer.Calculs(sample1, symbol, sample2, openBracket, endBracket)
		}
	}

	fmt.Printf("Your result is : %g\n", result)
}
