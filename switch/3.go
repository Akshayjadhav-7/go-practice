//Write a go program to check whether an alphabet is vowel or consonant.
package main

import "fmt"

func main() {

	alphabet := 'm'

	switch alphabet {
	case 'a':
		fmt.Print("Given character is vowel")
	case 'e':
		fmt.Print("Given character is vowel")
	case 'i':
		fmt.Print("Given character is vowel")
	case 'o':
		fmt.Print("Given character is vowel")
	case 'u':
		fmt.Print("Given character is vowel")
	default:
		fmt.Print("Given character is consonant")
	}
}
