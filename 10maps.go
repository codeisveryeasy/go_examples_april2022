package main

import "fmt"

func main() {
	//key:value pairs is map
	var studentScores = map[string]int{"okie": 22, "moi": 18, "lia": 24, "iot": 40}
	fmt.Println(studentScores)
	fmt.Println(studentScores["lia"])
	fmt.Println(studentScores["pia"])
	studentScores["pia"] = 26
	studentScores["yio"] = 29
	fmt.Println(studentScores)

	//declare map using make
	studentsScore := make(map[string]int)
	studentsScore["omg"] = 26
	fmt.Println(studentsScore)
	studentsScore["jio"] = 28
	studentsScore["dio"] = 30
	fmt.Println(studentsScore)
	delete(studentsScore, "dio")
	fmt.Println(studentsScore)

	students := make([]map[string]int, 4, 10)

	tempMap1 := make(map[string]int)
	tempMap1["obb"] = 5
	tempMap2 := make(map[string]int)
	tempMap2["hio"] = 15
	tempMap3 := make(map[string]int)
	tempMap3["obb"] = 25

	students[0] = tempMap1
	students[1] = tempMap2
	students[2] = tempMap3
	fmt.Println(students)
	fmt.Println("New length of slice:", len(students))
	fmt.Println("New capacity of slice:", cap(students))

	tempMap4 := make(map[string]int)
	tempMap4["obbo"] = 26
	students[3] = tempMap4
	fmt.Println(students)

	tempMap5 := make(map[string]int)
	tempMap5["ovo"] = 31
	//out of bounds
	//students[4] = tempMap5
	students = append(students, tempMap5)
	fmt.Println(students)
	fmt.Println("New length of slice:", len(students))
	fmt.Println("New capacity of slice:", cap(students))

}
