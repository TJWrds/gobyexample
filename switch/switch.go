// declares the program as an executable
package main

// imports default Golang libraries from main
import (
	"fmt"
	"time"
)

// func whatElseAmI(any interface{}) {
// 	switch t := any.(type) {
// 	case bool:
// 		fmt.Println("I'm a bool")
// 	case int:
// 		fmt.Println("I'm an int")
// 	default:
// 		fmt.Printf("Don't know type: %T\n", t)
// 	}
// }
func printsPenis() {
	fmt.Println("PENIS")
}

// Entry point for program. All data, operations, and arguments are supposed to pass through here
func main() {
	// declares 'i' as a variable, and assigns it the value of 5.
	i := 5
	fmt.Print("Write ", i, " as ")
	// The switch statement takes 'i' as an input
	// if i is 1 through 3, i is printed as a word.
	// if i is none of these(ex: a negative or greater than 3), print "Four or above"
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Four or above")
	}
	//  The switch statement takes the current time and day of the week(from the time golang library) as an input
	// if the current day of the week is Saturday or Sunday time.Weekday is printed as "it's the weekend"
	// if it is neither, it is printed as "It's a weekday" instead.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	// if the current hour is less than 12, it is printed as "It's before noon"
	// for any other cases, it is printed as "It's after noon" by default.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")

	}
	// This line declares a new variable called whatAmI which is a function that takes an empty interface called any as an input
	whatAmI := func(any interface{}) {
		switch t := any.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type: %T\n", t)
		}
	}
	whatElseAmI := func(any interface{}) {}
	aFunctionThatTakesAString := func(any string) {}
	anFunctionThatTakesAFunctionThatTakesNothing := func(emptyFunc func()) {
		emptyFunc()
	}

	// whatAmI(true)
	// whatAmI(1)
	// whatAmI("hey")
	// whatAmI(whatAmI)
	whatAmI(whatElseAmI)
	whatAmI(aFunctionThatTakesAString)
	whatAmI(anFunctionThatTakesAFunctionThatTakesNothing)
	anFunctionThatTakesAFunctionThatTakesNothing(main)
	// fmtPrintln := fmt.Println
	// fmtPrintln("I'm reusing an external function through a local variable")
}
