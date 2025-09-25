// Option à faire : gestion de la virgule dans l'inpout, les calculs et l'output.
// Option à faire : gérer les premiers nombres négatifs.

package main

import (
	"calculator/gestionerror"
	engine "calculator/resolveengine"
	"fmt"
)

func main() {

	fmt.Println("Hello, please, use this template to handle the calculator : \"1234-34/92(84+8)\"")

	gestionerror.ImportError()
	calcul := gestionerror.InputError()
	engine.FirstStep(calcul)
}
