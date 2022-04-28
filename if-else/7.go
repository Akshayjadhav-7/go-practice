// 7. Check whether a character is alphabet or not.
package main

import "fmt"

func main() {

	alphabet := '0'
	checkAlphabet(alphabet)

}

func checkAlphabet(alphabet rune) {

	if (alphabet >= 'a' && alphabet <= 'z') || (alphabet >= 'A' && alphabet <= 'Z') {
		fmt.Print("Given character is alphabet")
	} else {
		fmt.Print("Given character is not alphabet")
	}

}
