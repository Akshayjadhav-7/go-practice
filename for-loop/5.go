//Sum of first N numbers
package main

import "fmt"

func main() {

	var n int = 20
	var sum int = 0

	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("%d ", sum)
}
