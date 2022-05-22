// 2. Write a Go program that prints the count of occurence of array elements.
// Example: {1, 2, 3, 2, 3, 4, 5, 4, 6, 6, 7, 7, 7}
// 1 = 1, 2 = 2, 3 = 2, 4 = 2, 5 = 1, 6 = 2, 7 = 3

package main

import "fmt"

func main() {

	a := []int{4, 3, 6, 8, 3, 2, 4, 4, 9}
	fmt.Println("Given array is", a)

	var result = make(map[int]int)

	for _, num := range a {

		result[num] = result[num] + 1

	}
	fmt.Println("Final result", result)
}
