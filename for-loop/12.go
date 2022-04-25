//12.

//  * * * * *

//	* * * * *

//	* * * * *
package main

import "fmt"

func main() {

	for i := 0; i <= 5; i++ {
		for j := 0; j <= 4; j++ {
			if i%2 != 0 {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}
}
