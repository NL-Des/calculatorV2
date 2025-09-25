package rules

func MinusFirst(caracter rune, i int, sample1 []rune, firstMinusCoord int) ([]rune, int) {
	if caracter == '-' && i == 0 {
		firstMinusCoord = i
		sample1 = append(sample1, '-')
	}
	return sample1, firstMinusCoord
}
