// 17.  *
// 	   * *
// 	  * * *
// 	 * * * *
// 	* * * * *
package main

import "fmt"

func main() {
	n := 5
	for i := 1; i <= n; i++ {
		for space := 1; space <= n-i; space++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
