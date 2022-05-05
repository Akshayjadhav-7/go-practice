//1. Write a go program to print day of week name.
package main

import "fmt"

func main() {
	var day int
	print("Enter the day")
	fmt.Scanln(&day)

	switch day {
	case 1:
		fmt.Print("The day is Monday")
	case 2:
		fmt.Print("The day is tuesday")
	case 3:
		fmt.Print("The day is Wednesday")
	case 4:
		fmt.Print("The day is Thursday")
	case 5:
		fmt.Print("The day is Friday")
	case 6:
		fmt.Print("The day is Saturday")
	case 7:
		fmt.Print("The day is Sunday")
	default:
		fmt.Print("Invalid")
	}
}
