// Option à faire : gérer les parenthèses.
// Option à faire : gérer les espaces ?
package main

import (
	"calculator/gestionerror"
	engine "calculator/resolveengine"
)

func main() {

	gestionerror.ImportError()
	calcul := gestionerror.InputError()
	engine.Solution(calcul)
}
