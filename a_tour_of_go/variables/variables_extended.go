// Variables with initializers

// A var declaration can include initializers, one per variable.
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

package variables

import "fmt"

var i, j = 1, 2

func PrintVariablesExtended() {
	var correct, name = false, "Sasi"
	fmt.Println(i, j, correct, name)
}