package rules

func BracketGestion(caracter rune, symbolOpenBracket rune, symbolEndBracket rune) (rune, rune) {
	if caracter == '(' {
		symbolOpenBracket = '('
	}
	if caracter == ')' {
		symbolEndBracket = ')'
	}

	return symbolOpenBracket, symbolEndBracket
}
