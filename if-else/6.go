//6. Check whether a year is leap year or not.
package main

import "fmt"

func main() {
	year := 1997
	leapYear(year)
}

func leapYear(year int) {

	if year%400 == 0 {
		fmt.Print("Given year is Leap year")
	} else if year%4 == 0 {
		fmt.Print("Given year is Leap year")
	} else {
		fmt.Print("Given year is not Leap year")
	}
}
