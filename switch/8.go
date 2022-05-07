//Write a go program to create Simple Calculator.(addition|subtraction|multiplication|division|modulas|square|cube)
package main

import "fmt"

func main() {

	for {

		fmt.Printf("\n\n=====WELCOME TO CALCULATOR===== \n\n")
		fmt.Printf("1.Addition \n\n")
		fmt.Printf("2.Subtraction \n\n")
		fmt.Printf("3.Multiplication \n\n")
		fmt.Printf("4.Division \n\n")
		fmt.Printf("5.Modulas \n\n")
		fmt.Printf("6.Square \n\n")
		fmt.Printf("7.Cube \n\n")
		fmt.Printf("8.Exit \n\n")
		fmt.Printf("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Printf("Enter the first number: \n")
			var a int
			fmt.Scanln(&a)
			fmt.Printf("Enter the second number: \n")
			var b int
			fmt.Scanln(&b)
			fmt.Print("The addition of ", a, " and ", b, " is: ", addition(a, b))

		case 2:
			fmt.Printf("Enter the first number: \n")
			var a int
			fmt.Scanln(&a)
			fmt.Printf("Enter the second number: \n")
			var b int
			fmt.Scanln(&b)
			fmt.Print("The subtraction of ", a, " and ", b, " is: ", subtraction(a, b))

		case 3:
			fmt.Printf("Enter the first number: \n")
			var a int
			fmt.Scanln(&a)
			fmt.Printf("Enter the second number: \n")
			var b int
			fmt.Scanln(&b)
			fmt.Print("The multiplication of ", a, " and ", b, " is: ", multiplication(a, b))

		case 4:
			fmt.Printf("Enter the first number: \n")
			var a int
			fmt.Scanln(&a)
			fmt.Printf("Enter the second number: \n")
			var b int
			fmt.Scanln(&b)
			fmt.Print("The Division of ", a, " and ", b, " is: ", division(a, b))

		case 5:
			fmt.Printf("Enter the first number: \n")
			var a int
			fmt.Scanln(&a)
			fmt.Printf("Enter the second number: \n")
			var b int
			fmt.Scanln(&b)
			fmt.Print("The Modulas of ", a, " and ", b, " is: ", modulas(a, b))

		case 6:
			fmt.Printf("Enter the number: \n")
			var a int
			fmt.Scanln(&a)
			fmt.Print("The Square of the ", a, " is: ", square(a))

		case 7:
			fmt.Printf("Enter the number: \n")
			var a int
			fmt.Scanln(&a)
			fmt.Print("The Cube of the ", a, " is:", cube(a))

		case 8:
			fmt.Print("Exit")
			return

		default:
			fmt.Print("Invalid choice!!")
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
