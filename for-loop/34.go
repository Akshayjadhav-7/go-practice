//Sum of digits of a given number.

package main

import "fmt"

func main() {
	n := 342
	fmt.Print("Sum of digita of a given number is: ", getSum(n))
}

func getSum(n int) int {
	sum := 0

	for ; n > 0; n = n / 10 {
		sum = sum + n%10
	}
	return sum
}
