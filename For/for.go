//Declares program as an executable
package main

// imports default Golang library
import "fmt"

// entry point for program
func main() {
	// declares 'i' and gives it the int type and value of 1. The For loop prints the value of i so long as that value remains less than or equal to 3.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
		// adds 1 to the value of 'i' each time it passes through the loop before terminating
	}

	fmt.Println("now j")
	// declares 'i; and gives it the int type and value of 7. The for loop will print j so long as it remains less than or equal to 9.
	for j := 7; j <= 9; j++ {
		// ++ operand automatically adds to the value of 'j' in increments of 1.
		fmt.Println(j)
	}
	// This loop prints a string and breaks after the initial pass.
	for {
		fmt.Println("loop")
		break
	}
	// Declares 'n' and gives it the int type and value of 0. The loop continues so long as 'n' is less than 5.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	// This loop prints the odd values for 'n' and moves to the next iteration of the loop if the remainder value of n%2 is equal to 0
}
