// Write a C program to check whether a number is even or odd.
package main

import "fmt"

func main() {
	num := 7

	switch (num%2 == 0) || (num%2 != 0) {
	case num%2 == 0:
		fmt.Print("Given number is Even")
	case num%2 != 0:
		fmt.Print("Given number is Odd")
	}
}
