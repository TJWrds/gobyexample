//Declares program as an executable
package main

// imports "fmt" and "math" Golang libraries from main
import (
	"fmt"
	"math"
)

// declares a new variable 's' as a string with the immutable value of "Constant"
const s string = "Constant"

// entry point for program
func main() {
	// prints 's'
	fmt.Println(s)
	// declares immutable variable 'n', assigning it the int value 500000000.
	const n = 500000000
	// declares immutable variable 'd', assigning it the value 3e20. Scientific notation
	const d = 3e20 / n
	// These statements print the values of the variables declared above. They are passed through an immutable value type and then a math function which is called.
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(int64(n))

	fmt.Println(math.Sin(n))
	fmt.Println(math.Sin(d))
}
