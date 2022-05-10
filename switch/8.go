//Write a go program to create Simple Calculator.(addition|subtraction|multiplication|division|modulas|square|cube)
package main

import "fmt"

func main() {

	for {
		fmt.Println("=====WELCOME TO CALCULATOR=====")
		fmt.Println("1.Addition ")
		fmt.Println("2.Subtraction")
		fmt.Println("3.Multiplication")
		fmt.Println("4.Division")
		fmt.Println("5.Modulas")
		fmt.Println("6.Square")
		fmt.Println("7.Cube")
		fmt.Println("8.Exit")
		fmt.Println("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var a, b int
			fmt.Println("Enter the first number:")
			fmt.Scanln(&a)
			fmt.Println("Enter the second number:")
			fmt.Scanln(&b)
			fmt.Println("The addition of ", a, " and ", b, " is: ", addition(a, b))

		case 2:
			fmt.Printf("Enter the first number:")
			var a, b int
			fmt.Scanln(&a)
			fmt.Println("Enter the second number:")
			fmt.Scanln(&b)
			fmt.Println("The subtraction of ", a, " and ", b, " is: ", subtraction(a, b))

		case 3:
			fmt.Println("Enter the first number:")
			var a, b int
			fmt.Scanln(&a)
			fmt.Println("Enter the second number:")
			fmt.Scanln(&b)
			fmt.Println("The multiplication of ", a, " and ", b, " is: ", multiplication(a, b))

		case 4:
			fmt.Println("Enter the first number:")
			var a, b int
			fmt.Scanln(&a)
			fmt.Println("Enter the second number:")
			fmt.Scanln(&b)
			fmt.Println("The Division of ", a, " and ", b, " is: ", division(a, b))

		case 5:
			fmt.Println("Enter the first number:")
			var a, b int
			fmt.Scanln(&a)
			fmt.Println("Enter the second number:")
			fmt.Scanln(&b)
			fmt.Println("The Modulas of ", a, " and ", b, " is: ", modulas(a, b))

		case 6:
			fmt.Println("Enter the number:")
			var a int
			fmt.Scanln(&a)
			fmt.Println("The Square of the ", a, " is: ", square(a))

		case 7:
			fmt.Println("Enter the number:")
			var a int
			fmt.Scanln(&a)
			fmt.Println("The Cube of the ", a, " is:", cube(a))

		case 8:
			fmt.Println("Exit")
			return

		default:
			fmt.Println("Invalid choice!!")
		}
	}
}

func addition(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) float64 {
	return float64(a / b)
}

func modulas(a, b int) float64 {
	return float64(a % b)
}

func square(a int) int {
	return a * a
}

func cube(a int) int {
	return a * a * a
}
