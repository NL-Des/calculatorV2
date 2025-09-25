package rules

func SecondPart(sample2 []rune, caracter rune, symbol rune) []rune {
	if (caracter >= '0' && caracter <= '9') || (caracter == '.') || (caracter == ',') && symbol != ' ' {
		if caracter == ',' {
			caracter = '.'
		}
		sample2 = append(sample2, caracter)
	}
	return sample2
}
