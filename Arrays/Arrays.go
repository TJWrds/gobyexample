package main

import (
	"fmt"
)

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [10][10]int
	// i = 0, j = 0: twoD[0][0] = 0 + 0 -> 0
	// i = 0, j = 1: twoD[0][1] = 0 + 1 -> 1
	// i = 0, j = 2: twoD[0][2] = 0 + 2 -> 2
	// i[0] = [3]int{0, 1, 2}
	// i = 1, j = 0: twoD[1][0] = 1 + 0 -> 1
	// i = 1, j = 1: twoD[1][1] = 1 + 1 -> 2
	// i = 1, j = 2: twoD[1][2] = 1 + 2 -> 3
	// i[1] = [3]int{1, 2, 3}
	fmt.Println(len(twoD))
	fmt.Println(len(twoD[0]))
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[0]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
