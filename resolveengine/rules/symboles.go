package rules

// Handle symbols ('+' '-' '*' '/')
// Keep smybols in memory.
func SymbolGestion(caracter rune, symbol rune, i int, firstMinusCoord int) rune {

	if firstMinusCoord == i { // If positive, end the function. To avoid the repetition of the minus.

	} else {
		if !(caracter >= '0' && caracter <= '9') && (caracter != '.') && (symbol == ' ') {
			if caracter == '+' {
				symbol = caracter
			}
			if caracter == '-' {
				symbol = caracter
			}
			if caracter == '*' {
				symbol = caracter
			}
			if caracter == '/' {
				symbol = caracter
			}
		}
	}
	return symbol
}
