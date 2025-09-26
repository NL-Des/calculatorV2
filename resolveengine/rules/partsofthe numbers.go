package rules

func NumberParts(sample1, sample2 []rune, caracter rune, symbol rune) ([]rune, []rune) {
	if (caracter >= '0' && caracter <= '9') || (caracter == '.') || (caracter == ',') {
		if caracter == ',' {
			caracter = '.'
		}
		if symbol == ' ' {
			sample1 = append(sample1, caracter)
		} else {
			sample2 = append(sample2, caracter)
		}

	}
	return sample1, sample2
}
