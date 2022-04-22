//Factorial of a number N. (5! = 1 * 2 * 3 * 4 * 5 = 120)
package main

import "fmt"

func main() {

	var n int = 10
	var factorial int = 1

	for i := n; i >= 1; i-- {
		factorial = factorial * i
	}
	fmt.Printf("%d ", factorial)
}
