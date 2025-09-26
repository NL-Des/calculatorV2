package rules

func BracketDetection(i int, caracter rune, openBracket bool, endBracket bool, openBracketID int, endBracketID int) (bool, bool, int, int) {
	if caracter == '(' {
		openBracketID = i
		openBracket = true
	}
	if caracter == ')' {
		endBracketID = i
		endBracket = true

	}

	return openBracket, endBracket, openBracketID, endBracketID
}
