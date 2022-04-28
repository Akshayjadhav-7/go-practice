//6. Check whether a year is leap year or not.
package main

import "fmt"

func main() {
	year := 1997
	if leapYear(year) {
		fmt.Print("Year is leap year")
		return
	}
	fmt.Print("Year is not leap year")
}

func leapYear(year int) bool {

	if year%400 == 0 {
		return true
	} else if year%4 == 0 {
		return true
	}
	return false
}
