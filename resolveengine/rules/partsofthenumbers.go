package rules

// Handle if the numbers mus to go in the first or the second part of the calcul.
func NumberParts(sample1, sample2 []rune, caracter rune, symbol rune) ([]rune, []rune) {
	// If it's a number or an '.' or an ','
	if (caracter >= '0' && caracter <= '9') || (caracter == '.') || (caracter == ',') {

		if caracter == ',' { // GO can't handle the ','
			caracter = '.' // So we convert it in an ','
		}

		if symbol == ' ' { // We use the symbol as an separator to distinguish the first to the second part of the calcul.
			sample1 = append(sample1, caracter)
		} else {
			sample2 = append(sample2, caracter)
		}

	}
	return sample1, sample2
}
