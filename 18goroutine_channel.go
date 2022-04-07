package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//using channel in place of wait groups
	channel := make(chan int)

	fmt.Println("Start of main thread")
	go printNumbers(1, 5, channel)
	go printNumbers(2, 5, channel)
	go printNumbers(3, 5, channel)
	go printNumbers(4, 5, channel)
	go printNumbers(5, 5, channel)
	go printNumbers(6, 5, channel)
	go printNumbers(7, 5, channel)
	go printNumbers(8, 5, channel)
	<-channel
	<-channel
	<-channel
	<-channel
	<-channel
	<-channel
	<-channel
	<-channel

	fmt.Println("End of main thread")
}

func printNumbers(id int, count int, channel chan int) {
	fmt.Println("In printNumbers with id", id)
	rand.Seed(time.Now().Unix())
	for i := 0; i < count; i++ {
		var random int
		random = rand.Intn(10)
		fmt.Printf("%v: %v \n", id, i)
		sleep := time.Millisecond * time.Duration(random*100)
		fmt.Printf("Waiting for %v ms \n", sleep)
		time.Sleep(sleep)
	}

	channel <- 1
}
