//Print table of a number N.
package main

import "fmt"

func main() {

	var num int = 20
	var table int = 0

	for i := 1; i <= 10; i++ {
		table = i * num
		fmt.Printf("%d ", table)
	}
}
