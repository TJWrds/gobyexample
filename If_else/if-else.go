// declares program as an executable
package main

// imports default Golang library from main
import "fmt"

// Program entry point
func main() {
	// Prints "7 is even" if the remainder of 7 divided by 2 is equal to zero. Otherwise prints "7 is odd"
	if 7%2 == 0 {
		fmt.Println("7 is even")

	} else {
		fmt.Println("7 is odd")
	}
	// prints "8 is divisible by 4" if 8 divided by 4 equals 0.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")

	}
	// Declares num as a variable with int type and value of 9. prints "is negative" if the value for num is less than 0, "has 1 digit" if it's less than 10, and "has multiple digits" if it has any other undefined value.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		// else if is a compromise between if and else statements. an else if follows the preceding if, and is followed by any additional else if statements until it reaches the final else statement. Each if block can have only 1 if and 1 else statement.
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
