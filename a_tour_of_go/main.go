package main

import (
	"a_tour_of_go/imports"
	"a_tour_of_go/packages"
	"a_tour_of_go/exports"
)

// main is the single entrypoint that calls the example functions from the subpackages.
func main() {
	packages.PrintRandom()
	imports.PrintSqrt()
	exports.PrintPi()
}
