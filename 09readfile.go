package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filedata, error := ioutil.ReadFile("readme.txt")
	if error != nil {
		//fmt.Println(error)
		panic(error)
		//return
	}
	fmt.Println("File read is success")
	fmt.Println(filedata)
	fmt.Println(string(filedata))

	fmt.Printf(string(filedata))

	//write to file
	writeError := ioutil.WriteFile("writeme.txt", filedata, 0644)
	if writeError != nil {
		panic(writeError)
	}
}
