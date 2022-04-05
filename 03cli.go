package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("Executable: ", os.Args)

	fmt.Println("1st CLA:", os.Args[1])

	fmt.Println("2nd CLA:", os.Args[2])

	fmt.Println("Count of CLA: ", len(os.Args))

}
