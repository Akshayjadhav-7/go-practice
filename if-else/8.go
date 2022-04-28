// Input any alphabet and check whether it is vowel or consonant.
package main

import "fmt"

func main() {

	alphabet := 'a'

	if checkVowelConsonant(alphabet) {
		fmt.Printf("Given character is vowel")
		return
	}
	fmt.Printf("Given character is consonant")
}

func checkVowelConsonant(alphabet rune) bool {

	if alphabet == 'a' || alphabet == 'e' || alphabet == 'i' || alphabet == 'o' || alphabet == 'u' {
		return true
	}
	return false
}
