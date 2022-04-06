package main

import (
	"fmt"
	"reflect"
)

func main() {
	//declare array
	var scores = [4]int{2, 4, 8, 13}
	var emptyArray = [8]int{}
	var emptyFriends = [4]string{}
	var countOfFriends = [...]string{"obb", "kia", "mia", "jon"}
	var countOfFriends1 = []string{"obb", "kia", "mia", "jon"}

	var specificPosition = [9]int{0: 8, 5: 2, 8: 17}
	//index 11 is out of bounds
	//var specificPosition1 = [9]int{0: 8, 5: 2, 11: 17}

	fmt.Println(scores, " ", len(scores))
	fmt.Println(emptyArray)
	fmt.Println(emptyFriends)
	fmt.Println(countOfFriends, " ", len(countOfFriends), " ", reflect.TypeOf(countOfFriends1).Kind())
	fmt.Println(countOfFriends1, " ", len(countOfFriends1), " ", reflect.TypeOf(countOfFriends1).Kind())
	fmt.Println(specificPosition, " ", len(specificPosition))
	//fmt.Println(specificPosition1, " ", len(specificPosition1))

}
