// Functions

// A function can take zero or more arguments.
// In this example, add takes two parameters of type int.
// Notice that the type comes after the variable name.

package functions

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add_continued(x, y int) int {
	return x + y
}

func multiple_results(a, b string) (string, string) {
	return b, a
}

func PrintAdd() {
	fmt.Println("Addition of 2 and 5 ",add(2, 5))
	fmt.Println("Addition continued ", add_continued(2, 5))
	a, b := multiple_results("Hello", "World")
	fmt.Println("Multiple retun values ", a, b)
}