package gestionerror

import (
	"fmt"
	"os"
)

func ImportError() {
	//Input Error.
	switch {
	case len(os.Args) != 2:
		fmt.Println("Please, use this template : 1234-34/92(84+8)")
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
	bracketOpen := false
	withoutSymbol := true

	// Check first caracter and last caracter.
	for i := 0; i < len(calcul); i++ {
		if calcul[0] == '+' || calcul[0] == '/' || calcul[0] == '*' || calcul[0] == ')' {
			fmt.Println("Please do not put in first position an ')', an '+', an '/' or an '*' ")
			exitError()
		}
		if calcul[len(calcul)-1] == '-' || calcul[len(calcul)-1] == '+' || calcul[len(calcul)-1] == '/' || calcul[len(calcul)-1] == '*' || calcul[len(calcul)-1] == '(' {
			fmt.Println("Please do not put in last position an '(', an '+', an '/' or an '*' ")
			exitError()
		}
	}

	// Check for symbol together.
	for i := 1; i < len(calcul); i++ {
		if (calcul[i-1] == '-' || calcul[i-1] == '+' || calcul[i-1] == '/' || calcul[i-1] == '*') && (calcul[i] == '-' || calcul[i] == '+' || calcul[i] == '/' || calcul[i] == '*') {
			fmt.Println("Please do not put wrong symbols together")
			exitError()
		}
	}

	// Check for symbol wrongly together with an open or an end bracket.
	for i := 1; i < len(calcul); i++ {
		if calcul[i] == '(' && (calcul[i-1] == '(') && (calcul[i+1] == ')' || calcul[i+1] == '-' || calcul[i+1] == '+' || calcul[i+1] == '/' || calcul[i+1] == '*' || calcul[i+1] == '-') {
			fmt.Println("Please do not put symbols wronlgy together")
			exitError()
		}
		if calcul[i] == ')' && (calcul[i-1] == '(') && (calcul[i+1] == ')') {
			fmt.Println("Please do not put symbols wronlgy together")
			exitError()
		}
	}

	for _, caracter := range inpout {

		// Caracters forgotten.
		if (caracter < 0 && caracter > 9) && caracter != '+' && caracter != '-' && caracter != '/' && caracter != '*' && caracter != '(' && caracter != ')' {
			fmt.Println("Please use only numbers (0123456789) and specials characters (+ - / () )")
			exitError()
		}
		// Handle equal.
		if caracter == '=' {
			fmt.Println("We search the result, why do you use this calculator if you already have it ^^? '='")
			exitError()
		}

		// Handle brackets.
		if caracter == '(' {
			bracketOpen = true
		}
		if caracter == ')' && bracketOpen == true {
			bracketOpen = false
		} else if caracter == ')' && bracketOpen == false {
			fmt.Println("You have a bad gestion of brackets.")
		}

		// If it's whitout symbols.
		if caracter == '+' || caracter == '-' || caracter == '/' || caracter == '*' || caracter == '/' || caracter == '(' || caracter == ')' {
			withoutSymbol = false
		}
	}

	// If without symbols, return the problem.
	if withoutSymbol == true {
		fmt.Println(inpout)
		exitError()
	}
	return calcul
}

// Kill the program
func exitError() {
	os.Exit(0)
}
