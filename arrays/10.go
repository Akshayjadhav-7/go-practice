//10. Sum of left & right diagonals of a matrices.
package main

import "fmt"

func main() {

	var matrix [3][3]int
	sum1 := 0
	sum2 := 0
	fmt.Println("Enter the matrix elements")
	for i, j := range matrix {
		for k := range j {
			fmt.Scanln(&matrix[i][k])
			if i == k {
				sum1 += matrix[i][k]
			}
			if i+k == 2 {
				sum2 += matrix[i][k]
			}
		}
	}
	fmt.Println("Sum of first diagonal is", sum1)
	fmt.Println("Sum of first diagonal is", sum2)
	fmt.Println("Entered matrix is\n", matrix)
}
