package gestionerror

import (
	"fmt"
	"os"
	"unicode"
)

// Handle the errors's arguments.
func ImportError() {
	switch {
	case len(os.Args) != 2:
		fmt.Println("Please, use this template as exemple : 23,43-43")
		exitError()
	case len(os.Args[1]) == 0:
		fmt.Println("Please, enter something")
		exitError()
	case len(os.Args[1]) > 99:
		fmt.Println("Please, limit your usage for less than 99 caracters")
		exitError()
	}
}

func InputError() []rune {

	// Host the inpout.
	inpout := os.Args[1]

	calcul := []rune(inpout)
	//bracketOpen := false
	withoutSymbol := true

	for i := 0; i < len(calcul); i++ {

		// Check first caracter is not an symbol.
		if calcul[0] == '+' || calcul[0] == '/' || calcul[0] == '*' || calcul[0] == ')' {
			fmt.Println("Please do not put in first position an ')', an '+', an '/' or an '*' ")
			exitError()
		}

		// Check last caracter is not an symbol.
		if calcul[len(calcul)-1] == '-' || calcul[len(calcul)-1] == '+' || calcul[len(calcul)-1] == '/' || calcul[len(calcul)-1] == '*' || calcul[len(calcul)-1] == '(' {
			fmt.Println("Please do not put in last position an '(', an '+', an '/' or an '*' ")
			exitError()
		}

		// Check two symbols are not together.
		if i > 0 && (calcul[i-1] == '-' || calcul[i-1] == '+' || calcul[i-1] == '/' || calcul[i-1] == '*') && (calcul[i] == '-' || calcul[i] == '+' || calcul[i] == '/' || calcul[i] == '*') {
			fmt.Println("Please do not put wrong symbols together")
			exitError()
		}

		// Caracters forgotten.
		if !unicode.IsDigit(calcul[i]) &&
			calcul[i] != '+' &&
			calcul[i] != '-' &&
			calcul[i] != '/' &&
			calcul[i] != '*' &&
			calcul[i] != '.' &&
			calcul[i] != ',' /* && calcul[i] != '(' && calcul[i] != ')' */ {
			fmt.Println("Please use only numbers (0123456789) and allowed specials characters (+ - / * )")
			exitError()
		}

		// Handle equal.
		if calcul[i] == '=' {
			fmt.Println("We search the result, why do you use this calculator if you already have it ^^? '='")
			exitError()
		}

		// If it's whitout symbols.
		if calcul[i] == '+' || calcul[i] == '-' || calcul[i] == '/' || calcul[i] == '*' || calcul[i] == '/' || calcul[i] == '(' || calcul[i] == ')' {
			withoutSymbol = false
		}
	}

	// Check for symbol wrongly together with an open or an end bracket.
	/* 	for i := 1; i < len(calcul); i++ {
		if calcul[i] == '(' && (calcul[i-1] == '(') && (calcul[i+1] == ')' || calcul[i+1] == '-' || calcul[i+1] == '+' || calcul[i+1] == '/' || calcul[i+1] == '*' || calcul[i+1] == '-') {
			fmt.Println("Please do not put symbols wronlgy together")
			exitError()
		}
		if calcul[i] == ')' && (calcul[i-1] == '(') && (calcul[i+1] == ')') {
			fmt.Println("Please do not put symbols wronlgy together")
			exitError()
		}
	} */

	/* 	for _, caracter := range inpout { */

	// Handle brackets.
	/* 		if caracter == '(' {
	   			bracketOpen = true
	   		}
	   		if caracter == ')' && bracketOpen == true {
	   			bracketOpen = false
	   		} else if caracter == ')' && bracketOpen == false {
	   			fmt.Println("You have a bad gestion of brackets.")
	   		} */

	/* 	} */

	// If without symbols, return the problem.
	if withoutSymbol {
		fmt.Printf("Please, don't forget symbols. This is your inpout : %s\n", inpout)
		exitError()
	}
	return calcul
}

// Kill the program
func exitError() {
	os.Exit(0)
}
