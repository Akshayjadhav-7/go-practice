// WAP to check entered number is even or odd using switch case
package main

import "fmt"

func main() {

	for {

		fmt.Print("\n====Even/Odd=====\n")
		fmt.Printf("1. Check whether the number is even or odd\n")
		fmt.Printf("2. Exit\n")
		fmt.Printf("Enter the choice\n")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Printf("\nEnter the number\n")
			var num int
			fmt.Scanln(&num)
			evenOdd(num)
		case 2:
			fmt.Print("\nExit\n")
			return
		default:
			fmt.Println("\nInvalid choice!!\n")
		}
	}

}
func evenOdd(a int) {
	if a%2 == 0 {
		fmt.Printf("\nEntered number is even\n")
	} else {
		fmt.Printf("\nEntered numbers is odd\n")
	}

}
