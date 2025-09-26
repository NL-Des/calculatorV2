package rules

func FirstPart(sample1 []rune, caracter rune, symbol rune) []rune {
	if ((caracter >= '0' && caracter <= '9') || (caracter == '.') || (caracter == ',')) && symbol == ' ' {
		if caracter == ',' {
			caracter = '.'
		}
		sample1 = append(sample1, caracter)
	}
	return sample1
}
