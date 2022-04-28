//2. Find maximum between three numbers.
package main

import "fmt"

func main() {
	num1 := 3
	num2 := 6
	num3 := 9

	if num1 > num2 && num1 > num3 {
		fmt.Print("Max num is:", num1)
	} else if num2 > num1 && num2 > num3 {
		fmt.Print("Max num is:", num2)
	} else if num3 > num1 && num3 > num2 {
		fmt.Print("Max num is:", num3)
	}
}
