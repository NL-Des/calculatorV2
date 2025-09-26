package rules

// If the first caracter of the argument is an '-'.
func MinusFirst(caracter rune, i int, sample1 []rune, firstMinusCoord int) ([]rune, int) {
	if caracter == '-' && i == 0 {
		firstMinusCoord = i
		sample1 = append(sample1, '-') // integrate the '-' to be able to take care of it.
	}
	return sample1, firstMinusCoord
}
