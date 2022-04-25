//Print squares of first N numbers
package main

import "fmt"

func main() {

	var n int = 20

	var sqr int = 0

	square(n, sqr)

}

func square(n, sqr int) {

	for i := 1; i <= n; i++ {
		sqr = i * i
		fmt.Printf("%d ", sqr)
	}
}
