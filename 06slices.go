package main

import (
	"fmt"
	"reflect"
)

func main() {
	//declare slice
	var dynamicArray = []int{1, 2, 3, 4}
	var scores = [4]int{1, 2, 3, 4}

	fmt.Println(reflect.TypeOf(dynamicArray).Kind())
	fmt.Println(reflect.TypeOf(scores).Kind())

	//use make to create slice
	//make(type, length, capacity)
	var mySlice1 = make([]int, 2, 5)
	fmt.Println("Length of slice:", len(mySlice1))
	fmt.Println("Capacity of slice:", cap(mySlice1))
	fmt.Println(mySlice1)
	mySlice1[0] = 9
	mySlice1[1] = 8
	//mySlice1[2] = 7
	mySlice1 = append(mySlice1, 7)
	fmt.Println(mySlice1)
	fmt.Println("Length of slice:", len(mySlice1))
	fmt.Println("Capacity of slice:", cap(mySlice1))
	mySlice1 = append(mySlice1, 1)
	mySlice1 = append(mySlice1, 2)
	fmt.Println("Capacity reaching to limit")
	fmt.Println("Length of slice:", len(mySlice1))
	fmt.Println("Capacity of slice:", cap(mySlice1))
	mySlice1 = append(mySlice1, 3)
	mySlice1 = append(mySlice1, 4)
	fmt.Println("Capacity excedding the limit")
	fmt.Println("Length of slice:", len(mySlice1))
	fmt.Println("Capacity of slice:", cap(mySlice1))
	fmt.Println(mySlice1)
	// fmt.Println("New length of slice:", len(mySlice1))
	// fmt.Println("New capacity of slice:", cap(mySlice1))
	mySlice1 = append(mySlice1, 11)
	mySlice1 = append(mySlice1, 12)
	mySlice1 = append(mySlice1, 13)
	mySlice1 = append(mySlice1, 14)
	fmt.Println(mySlice1)
	fmt.Println("New length of slice:", len(mySlice1))
	fmt.Println("New capacity of slice:", cap(mySlice1))
	mySlice1 = append(mySlice1, 21)
	mySlice1 = append(mySlice1, 22)
	mySlice1 = append(mySlice1, 23)
	mySlice1 = append(mySlice1, 24)
	mySlice1 = append(mySlice1, 25)
	mySlice1 = append(mySlice1, 26)
	mySlice1 = append(mySlice1, 27)
	mySlice1 = append(mySlice1, 28)
	mySlice1 = append(mySlice1, 29)
	mySlice1 = append(mySlice1, 30)
	fmt.Println(mySlice1)
	fmt.Println("New length of slice:", len(mySlice1))
	fmt.Println("New capacity of slice:", cap(mySlice1))

	mySlice2 := make([]int, 5, 20)
	fmt.Println(mySlice2)
	mySlice2[0] = 8
	mySlice2[2] = 9
	fmt.Println(mySlice2)
	mySlice2 = append(mySlice2, 1)
	fmt.Println(mySlice2)

}
