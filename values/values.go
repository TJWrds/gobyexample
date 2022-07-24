// Declares the program as an executable
package main

// imports external library from main
import "fmt"

//entry point for program.
func main() {
	// This block just displays the different ways different types of values can be displayed in Golang.
	fmt.Println("go"+"lang", "go", "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7/3 =", 7/3.14)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}
