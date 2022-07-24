package main

import "fmt"

func main() {
	// Declaring variable a and assigning it the string value "string"
	// This branch takes 'a' and prints its string value. It changes the string value for 'a' on lines 7, 9, and 11 Then prints the changed value on each subsequent line.
	var a = "string"
	fmt.Println(a)
	a = "new string value"
	fmt.Println(a)
	a = "1"
	fmt.Println(a)
	// This statement declares two variables 'b' and 'c' at once as an integer, and assigns them the int values of 1 and 2 respectively. Then prints the value
	var b, c int = 1, 2
	fmt.Println(b, c)
	// This statement declares variable 'd' and asigns it the bool type and value of true. Then it prints that boolean value.
	var d = true
	fmt.Println(d)
	// This statement declares variable 'e' and asigns it the int type along with the default int value of 0. Then prints the value for 'e'
	var e int
	fmt.Println(e)
	// This statement declares variable 'ef' and asigns it the type uint. 'ef' is assigned the value of 11 in the subsequent line. Then it prints the value for 'ef'
	var ef uint
	ef = 100
	fmt.Println(ef)
	// This statement declares variable 'f' and assigns it the string type and value of "apple" using only the ":=" operand. Then it prints that value.
	f := "apple"
	fmt.Println(f)

}
