//Print cubes of first N numbers
package main

import "fmt"

func main() {

	var n int = 20
	var cubes int = 0

	cube(n, cubes)
}

func cube(n, cubes int) {

	for i := 1; i <= n; i++ {

		cubes = i * i * i
		fmt.Printf("% d", cubes)
	}
}
