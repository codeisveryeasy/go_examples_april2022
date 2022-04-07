package main

import "fmt"

type student struct {
	studentName string
	schoolName  school
}

type school struct {
	name            string
	countOfStudents int
	city            string
}

func main() {
	school1 := school{name: "Air Asia Academy", countOfStudents: 44, city: "KL"}
	student1 := student{
		studentName: "Omi San",
		schoolName:  school1,
	}
	fmt.Println("Deference student1")
	fmt.Println(student1.schoolName)

	school2 := new(school)
	student2 := new(student)

	student2.schoolName = *school2
	fmt.Println(student2.schoolName)
	fmt.Println(*student2)

	school2.name = "RBA"
	school2.countOfStudents = 53
	school2.city = "PN"
	student2.studentName = "Whomi"
	//student2.schoolName = *school2
	fmt.Println(student2)
	//fmt.Println(*&student2.studentName)
	//fmt.Println(student2.studentName)

}
