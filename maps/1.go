// 1. Write a Go program which takes country as input from user and prints the country's capital as output.
// Example
// input = India  output = New Delhi
// input = Bangladesh output = Dhaka
// input = Pakistan output = Islamabad
// input = Srilanka output = Colombo

package main

import (
	"fmt"
)

func main() {

	map1 := map[string]string{
		"India":      "New delhi",
		"Bangladesh": "Dhaka",
		"Pakistan":   "Islamabad",
		"Srilanka":   "Colombo",
	}

	for {
		var country string

		// a, _ := strconv.Atoi(country)
		fmt.Println("Enter the country name from the below list to find out the capital city:")
		fmt.Println("1.India")
		fmt.Println("2.Bangladesh")
		fmt.Println("3.Pakistan")
		fmt.Println("4.Srilanka")
		fmt.Println("5.Exit")
		fmt.Scanln(&country)

		switch country {
		case "1":
			capital := map1[country]
			fmt.Println("The capital is:", capital)

		case "2":
			capital := map1[country]
			fmt.Println("The capital is:", capital)

		case "3":
			capital := map1[country]
			fmt.Println("The capital is:", capital)

		case "4":
			capital := map1[country]
			fmt.Println("The capital is:", capital)

		case "5":
			return
		}
	}
}
