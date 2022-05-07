//Write a go program to check whether a number is positive, negative or zero.
package main

import "fmt"

func main() {
	num := 5

	switch {
	case num > 0:
		fmt.Print("Given number is Positive")
	case num < 0:
		fmt.Print("Given number is Negative")
	case num == 0:
		fmt.Print("Given numbe is Zero")
	}
}
