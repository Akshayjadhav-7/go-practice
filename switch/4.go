// Write a C program to find maximum between two numbers.
package main

import "fmt"

func main() {
	num1, num2 := 9, 6

	switch (num1 > num2) || (num2 > num1) {
	case num1 > num2:
		fmt.Print("num1 is greater than num2")
	case num2 > num1:
		fmt.Print("num2 is greater than num1")
	}
}
