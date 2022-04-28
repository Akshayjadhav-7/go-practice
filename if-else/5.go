//5. Check whether a number is even or odd.
package main

import "fmt"

func main() {
	num := 5

	if num%2 == 0 {
		fmt.Print("Given number is even ", num)
	} else if num%2 != 0 {
		fmt.Print("Given number is odd ", num)
	}
}
