// 33. Sum of all integer between X and Y.
package main

import "fmt"

func main() {
	num1 := 2
	num2 := 5
	fmt.Printf("The sum of all integer is %d ", Sum(num1, num2))
}

func Sum(num1, num2 int) int {
	sum := 0
	for i := num1; i <= num2; i++ {
		sum += i
	}
	return sum
}
