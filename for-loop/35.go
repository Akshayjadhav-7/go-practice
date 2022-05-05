//Sum of numbers between X and Y which are divisible by N.
package main

import "fmt"

func main() {

	num1, num2 := 3, 6
	fmt.Print("Sum of numbers between", num1, "and", num2, "which are divisible by N", findSum(num1, num2))

}

func findSum(num1, num2 int) int {

	sum := 0
	n := 2

	for i := num1; i <= num2; i++ {
		for j := num1; j <= num2; j++ {
			sum += j % n
		}

	}
	return sum

}
