package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start of main thread")
	var wg sync.WaitGroup
	wg.Add(4)

	common := make(chan int)

	go printNumbers1(1, common, &wg)
	go printNumbers2(2, common, &wg)
	go printNumbers3(3, common, &wg)
	go printNumbers4(4, common, &wg)

	//fmt.Println("Press any key")
	//var pressAnyKey string
	//fmt.Scanln(&pressAnyKey)
	wg.Wait()
	fmt.Println("End of main thread")
}

func printNumbers1(n int, common chan int, wg *sync.WaitGroup) {
	fmt.Println("Printing ", n)
	fmt.Println("Writing to common....")
	time.Sleep(time.Second * 10)
	common <- 10
	fmt.Println("Written to common ", n)
	wg.Done()
}

func printNumbers2(n int, common chan int, wg *sync.WaitGroup) {
	fmt.Println("Printing ", n)
	fmt.Println("Wait till common is written to....")
	received := <-common
	fmt.Println("Common is read for n=", n, received)
	wg.Done()
}

func printNumbers3(n int, common chan int, wg *sync.WaitGroup) {
	fmt.Println("Printing ", n)
	common <- 100
	fmt.Println("Written to common ", n)
	wg.Done()
}

func printNumbers4(n int, common chan int, wg *sync.WaitGroup) {
	fmt.Println("Printing ", n)
	received := <-common
	fmt.Println("Common is read for n=", n, received)
	wg.Done()
}
