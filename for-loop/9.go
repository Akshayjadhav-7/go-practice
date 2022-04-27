//Print sum of square of first N numbers
package main

import "fmt"

func main() {

	var n int = 20
	var sum int = 0

	for i := 1; i <= n; i++ {
		sum += square(i)
	}
	fmt.Printf("%d ", sum)
}

func square(n int) int {
	return n * n
}
