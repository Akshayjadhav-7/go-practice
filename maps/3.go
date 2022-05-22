//3. Create a list of students. Print student details by taking roll number as input.

package main

import "fmt"

type Student struct {
	name    string
	class   string
	age     float64
	roll_no int
}

func main() {

	st1 := Student{
		roll_no: 101,
		name:    "Akshay",
		class:   "12th",
		age:     18,
	}
	st2 := Student{
		roll_no: 102,
		name:    "Ajay",
		class:   "12th",
		age:     17.5,
	}
	st3 := Student{
		roll_no: 103,
		name:    "VJ",
		class:   "12th",
		age:     18.2,
	}

	students := map[int]Student{
		101: st1,
		102: st2,
		103: st3,
	}

	var number int
	fmt.Println("Enter the roll number from below numbers to get the info of student")
	fmt.Println("101")
	fmt.Println("102")
	fmt.Println("103")
	fmt.Scanln(&number)

	student_info := students[number]

	fmt.Println(student_info)

}
