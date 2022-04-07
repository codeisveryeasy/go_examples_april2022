package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start of main thread")
	go printNumbers(1, 5)
	go printNumbers(2, 5)
	go printNumbers(3, 5)
	go printNumbers(4, 5)
	//fmt.Println("Sleep main thread for 10 seconds")
	//time.Sleep(time.Second * 2)
	fmt.Println("End of main thread")
}

func printNumbers(id int, count int) {
	fmt.Println("In printNumbers with id", id)
	for i := 0; i < count; i++ {
		fmt.Printf("%v: %v \n", id, i)
		time.Sleep(time.Millisecond * 500)

	}
}
