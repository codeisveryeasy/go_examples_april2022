package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	writeToFile(8000000, "filename1.txt")
	writeToFile(8000000, "filename2.txt")
	writeToFile(8000000, "filename3.txt")
	writeToFile(8000000, "filename4.txt")
	writeToFile(8000000, "filename5.txt")
	writeToFile(8000000, "filename6.txt")
	writeToFile(8000000, "filename7.txt")
	writeToFile(8000000, "filename8.txt")

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	fmt.Printf("Time taken: %v \n", timeTaken)
}

func writeToFile(countOfRows int, filename string) {

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

}
