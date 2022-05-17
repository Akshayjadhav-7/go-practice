//  Delete an element from array
package main

import "fmt"

func main() {

	a := []int{4, 7, 3, 9, 1}

	fmt.Println("Given array is:", a)
	deleteElement(a)
}

func deleteElement(a []int) {

	s1 := a[:1]
	s2 := a[2:]
	s3 := append(s1, s2...)
	fmt.Println("sliced array is:", s1)
	fmt.Println("sliced array is:", s2)
	fmt.Println("sliced array is:", s3)
}
