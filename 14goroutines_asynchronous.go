package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(8)
	fmt.Println("Start of main thread")
	go printNumbers(1, 5, &wg)
	go printNumbers(2, 5, &wg)
	go printNumbers(3, 5, &wg)
	go printNumbers(4, 5, &wg)
	go printNumbers(5, 5, &wg)
	go printNumbers(6, 5, &wg)
	go printNumbers(7, 5, &wg)
	go printNumbers(8, 5, &wg)
	//fmt.Println("Sleep main thread for 10 seconds")
	//time.Sleep(time.Second * 2)
	fmt.Println("Waiting to wg.Done for each go routine....")
	wg.Wait()
	fmt.Println("End of main thread")
}

func printNumbers(id int, count int, wg *sync.WaitGroup) {
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
	//if id < 8 {
	wg.Done()
	//}

}
