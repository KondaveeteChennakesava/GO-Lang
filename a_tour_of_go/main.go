package main

import (
	"a_tour_of_go/exports"
	"a_tour_of_go/functions"
	"a_tour_of_go/imports"
	"a_tour_of_go/packages"
	typeconversions "a_tour_of_go/type_conversions"
	"a_tour_of_go/variables"
)

// main is the single entrypoint that calls the example functions from the subpackages.
func main() {
	packages.PrintRandom()
	imports.PrintSqrt()
	exports.PrintPi()
	functions.PrintAdd()
	functions.PrintNamedReturns()
	variables.PrintVariabels()
	variables.PrintVariablesExtended()
	variables.PrintVariablesDeclaration()
	typeconversions.PrintTypeConversions()
}
