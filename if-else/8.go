// Input any alphabet and check whether it is vowel or consonant.
package main

import "fmt"

func main() {

	alphabet := 'a'

	checkVowelConsonant(alphabet)
}

func checkVowelConsonant(alphabet rune) {

	if alphabet == 'a' || alphabet == 'e' || alphabet == 'i' || alphabet == 'o' || alphabet == 'u' {
		fmt.Print("Given alphabet is vowel")
	} else {
		fmt.Print("Given alphabet is consonant")
	}
}
