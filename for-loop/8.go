//Print cubes of first N numbers
package main

import "fmt"

func main() {

	var n int = 20

	for i := 1; i <= n; i++ {
		fmt.Println("Cube of", i, "is:", cube(i))
	}
}

func cube(n int) int {
	return n * n * n
}
