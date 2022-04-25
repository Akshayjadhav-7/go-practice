// 11.
//	       *
//     * *
//     * * *
//     * * * *
//     * * * * *

package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {

		for j := 1; j <= i; j++ {

			fmt.Printf("*")
		}
		fmt.Printf("\n")

	}

	for x := 1; x <= 5; x++ {

		fmt.Printf("\n")

		for y := 1; y <= x; y++ {

			if x+y <= 2 {

				fmt.Printf(" ")
			}

			fmt.Printf("*")
		}

	}
}
