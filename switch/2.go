//Write a C program print total number of days in a month.
package main

import "fmt"

func main() {
	var month int
	fmt.Print("Enter the month number")
	fmt.Scanln(&month)

	switch month {
	case 1:
		fmt.Print("No. of days in month of January are 31")
	case 2:
		fmt.Print("No. of days in month of Feb are 28/29")
	case 3:
		fmt.Print("No. of days in month of March are 31")
	case 4:
		fmt.Print("No. of days in month of April are 30")
	case 5:
		fmt.Print("No. of days in month of May are 31")
	case 6:
		fmt.Print("No. of days in month of June are 30")
	case 7:
		fmt.Print("No. of days in month of July are 31")
	case 8:
		fmt.Print("No. of days in month of August are 31")
	case 9:
		fmt.Print("No. of days in month of September are 30")
	case 10:
		fmt.Print("No. of days in month of October are 31")
	case 11:
		fmt.Print("No. of days in month of November are 30")
	case 12:
		fmt.Print("No. of days in month of December are 31")
	}
}
