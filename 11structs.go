package main

import "fmt"

type school struct {
	name            string
	countOfStudents int
	city            string
}

func main() {
	//empty instance
	var school1 = school{}
	fmt.Println(school1)

	school1.name = "AAA"
	school1.city = "KL"
	school1.countOfStudents = 88888
	fmt.Println(school1)

	//intialize new instance with values
	var s2 = school{name: "RBA", countOfStudents: 7171717171, city: "MY"}
	fmt.Println(s2)

	var newS = NewSchool("OBB", 44, "Chennai")
	fmt.Println(newS)
	fmt.Println(*newS)

}

func NewSchool(name string, count int, location string) *school {
	var s = school{
		name:            name,
		countOfStudents: count,
		city:            location,
	}
	return &s
}
