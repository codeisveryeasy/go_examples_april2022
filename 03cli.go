package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"time"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("Executable: ", os.Args)

	fmt.Println("1st CLA:", os.Args[1])

	fmt.Println("2nd CLA:", os.Args[2])

	fmt.Println("Count of CLA: ", len(os.Args))
	fmt.Println(runtime.Version())
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())
	fmt.Println(os.Getenv("PATH"))

	checkEnvironmentVariable := "TEMP"
	value, present := os.LookupEnv(checkEnvironmentVariable)
	if present {
		fmt.Printf("%v:%v", checkEnvironmentVariable, value)
	} else {
		fmt.Printf("%v is not present", checkEnvironmentVariable)
	}

	//date time library
	fmt.Println()
	current := time.Now()
	fmt.Println(current)
	fmt.Println(current.Hour())
	fmt.Println(current.Minute())
	fmt.Println(current.Second())
	fmt.Println(current.Month())
	fmt.Println(current.Weekday())

	fmt.Println(reflect.TypeOf(current))
}
