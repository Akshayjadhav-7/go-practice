//Print odd numbers between N…1
package main

import "fmt"

func main() {

	var n int = 20

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			fmt.Printf("%d ", i)
		}
	}
}
