package rules

func SymbolGestion(caracter rune, symbol rune, i int, firstMinusCoord int) rune {

	if firstMinusCoord == i {

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
