//Factorial of a number N. (5! = 1 * 2 * 3 * 4 * 5 = 120)
package main

import "fmt"

func main() {

	var n int = 5
	var factorial int = 1

	facto(n, factorial)

}

func facto(n, factorial int) {

	for i := n; i >= 1; i-- {
		factorial = factorial * i
	}
	fmt.Printf("%d ", factorial)

}
