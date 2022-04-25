//Print squares of first N numbers
package main

import "fmt"

func main() {

	n := 20
	squares(n)
}
func squares(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result = i * i
		fmt.Printf("%d ", result)
	}
	return result
}
