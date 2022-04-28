//1. Find maximum between two numbers.

package main

import "fmt"

func main() {
	num1 := 100
	num2 := 9

	if num1 > num2 {
		fmt.Printf("Max num between", num1, "and", num2, "is:", num1)
	} else {
		fmt.Printf("Max num between", num1, "and", num2, "is:", num2)
	}
}
