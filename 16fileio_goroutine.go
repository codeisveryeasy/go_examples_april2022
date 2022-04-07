package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {

	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(8)

	go writeToFile(8000000, "filename10.txt", &wg)
	go writeToFile(8000000, "filename11.txt", &wg)
	go writeToFile(8000000, "filename12.txt", &wg)
	go writeToFile(8000000, "filename13.txt", &wg)
	go writeToFile(8000000, "filename14.txt", &wg)
	go writeToFile(8000000, "filename15.txt", &wg)
	go writeToFile(8000000, "filename16.txt", &wg)
	go writeToFile(8000000, "filename17.txt", &wg)
	wg.Wait()
	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	fmt.Printf("Time taken: %v \n", timeTaken)

}

func writeToFile(countOfRows int, filename string, wg *sync.WaitGroup) {

	filePointer, error := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if error != nil {
		panic(error)
	}
	for i := 0; i < countOfRows; i++ {
		content := fmt.Sprintf("Line number: %v \n", i)
		filePointer.WriteString(content)
	}
	filePointer.Close()

	//fmt.Printf("Written %v lines to file %v\n", countOfRows, filename)
	wg.Done()

}
