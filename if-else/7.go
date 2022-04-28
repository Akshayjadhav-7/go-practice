// 7. Check whether a character is alphabet or not.
package main

import "fmt"

func main() {

	alphabet := 'a'
	if checkAlphabet(alphabet) {
		fmt.Print("Given character is alphabet")
		return
	}
	fmt.Print("Give character is not alphabet")

}

func checkAlphabet(alphabet rune) bool {

	if (alphabet >= 'a' && alphabet <= 'z') || (alphabet >= 'A' && alphabet <= 'Z') {
		return true
	}
	return false
}
