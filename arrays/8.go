//Matrix addition and subtraction (*USE: functions,switch cases)
package main

import "fmt"

func main() {

	var choice int

	for {

		fmt.Println("====MENU===")
		fmt.Println("1.Addition of matrix")
		fmt.Println("2.Subtraction of matrix")
		fmt.Println("Enter your choice")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			matrixAddition()

		case 2:
			matrixSubtraction()
		}

	}

}

func matrixAddition() {

	var matrix1 [2][3]int
	var matrix2 [2][3]int
	var resultMatrix [2][3]int

	fmt.Println("Enter the numbers in matrix 1")
	for i, j := range matrix1 {
		for l := range j {
			fmt.Scanln(&matrix1[i][l])
		}
	}

	fmt.Println("Enter the numbers in matrix 2")
	for k, m := range matrix2 {
		for p := range m {
			fmt.Scanln(&matrix2[k][p])
		}
	}

	for i, row := range resultMatrix {
		for j := range row {
			resultMatrix[i][j] = matrix1[i][j] + matrix2[i][j]
			fmt.Println("Addotion of matrix is:", resultMatrix[i][j])
		}
	}
}

func matrixSubtraction() {
	var matrix1 [2][3]int
	var matrix2 [2][3]int
	var resultMatrix2 [2][3]int

	fmt.Println("Enter the matrix1 elements")
	for i, j := range matrix1 {
		for k := range j {
			fmt.Scanln(&matrix1[i][k])
		}
	}

	fmt.Println("Enter the matrix2 elements")
	for l, m := range matrix2 {
		for p := range m {
			fmt.Scanln(&matrix2[l][p])
		}
	}

	for x, row := range resultMatrix2 {
		for y := range row {
			resultMatrix2[x][y] = matrix1[x][y] - matrix2[x][y]
		}
	}
	fmt.Println("Subtraction of two matix is", resultMatrix2)
}
