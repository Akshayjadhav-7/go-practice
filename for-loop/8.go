//Print cubes of first N numbers
package main

import "fmt"

func main() {

	var n int = 20
	cubes(n)
}
func cubes(n int) int {
	cubes := 0
	for i := 1; i <= n; i++ {
		cubes = i * i * i
		fmt.Printf("% d", cubes)
	}
	return cubes
}
