package main

import "fmt"

func main() {
	count := 4
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
	fmt.Println("----------------")
	for count < 8 {
		fmt.Println(count)
		count++
	}
	fmt.Println("----------------")
	for count := 10; ; count++ {
		if count == 20 {
			continue
		}
		if count == 25 {
			break
		}
		fmt.Println(count)
	}
	fmt.Println("----------------")
	//for range
	name := "Air Asia Academy"
	fmt.Println(len(name))
	for i, v := range name {
		fmt.Printf("Hello... %v %c \n", i, rune(v))
	}
	fmt.Println("----------------")
	scores := [4]int{10, 20, 30, 40}
	for whereAmI, whoAmI := range scores {
		fmt.Println(whereAmI, whoAmI)
	}
	fmt.Println("----------------")
	scores1 := [4]int{10, 20, 30, 40}
	for index := range scores1 {
		fmt.Println(index)
	}
	fmt.Println("----------------")
	scores2 := [4]int{10, 20, 30, 40}
	for _, value := range scores2 {
		fmt.Println(value)
	}

}
