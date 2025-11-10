// Functions

// A function can take zero or more arguments.
// In this example, add takes two parameters of type int.
// Notice that the type comes after the variable name.

package functions

import "fmt"

func add(x int, y int) int {
	return x + y
}

func PrintAdd() {
	fmt.Println(add(2, 5))
}