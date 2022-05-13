//Seperate odd and even numbers of an array into separate arrays.
package main

import "fmt"

func main() {
	a := []int{34, 56, 77, 23, 98, 51, 87, 22}
	evenOdd(a)
}

func evenOdd(a []int) {
	even := []int{}
	odd := []int{}
	fmt.Println("Given array is:", a)

	for i, _ := range a {
		if a[i]%2 == 0 {
			even = append(even, a[i])
		} else {
			odd = append(odd, a[i])
		}
	}
	fmt.Println("Even numbers array:", even)
	fmt.Println("Odd numbers array:", odd)
}
