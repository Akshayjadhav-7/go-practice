//3. Check whether a number is negative, positive or zero.
package main

import "fmt"

func main() {
	num1 := 98

	if num1 > 0 {
		fmt.Print("Given num is Positive", num1)
	} else if num1 < 0 {
		fmt.Print("Given num is negative", num1)
	} else if num1 == 0 {
		fmt.Print("GIven number is Zero", num1)
	}
}
