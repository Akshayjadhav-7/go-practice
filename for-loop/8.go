//Print cubes of first N numbers
package main

import "fmt"

func main() {

	var n int = 20
	var cubes int = 0

	for i := 0; i <= n; i++ {
		cubes = i * i * i
		fmt.Printf("%d ", cubes)
	}
}
