//4. Check whether a number is divisible by 5 and 11 or not.
package main

import "fmt"

func main() {
	num1 := 2

	if num1%5 == 0 {
		fmt.Print("Given number is divisible by 5: ", num1)
	} else if num1%11 == 0 {
		fmt.Print("Given num is divisible 11: ", num1)
	} else {
		fmt.Print("Given num is not divisible by 5 or 11")
	}
}
