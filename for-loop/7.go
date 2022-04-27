//Print squares of first N numbers
package main

import "fmt"

func main() {

	n := 20

	for i := 1; i <= n; i++ {
		fmt.Println("Square of", i, "is:", square(i))
	}
}

func square(n int) int {
	return n * n
}
