// 13.
//  *   *   *
// 	*   *   *
// 	*   *   *
// 	*   *   *
// 	*   *   *
package main

import "fmt"

func main() {

	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			if j%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

}
