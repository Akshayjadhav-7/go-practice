//Average of first N numbers
package main

import "fmt"

func main() {

	var n int = 20
	var sum int = 0
	var avg int = 0

	for i := 1; i <= n; i++ {
		sum += i
	}
	avg = sum / n
	fmt.Printf("%d ", avg)
}
