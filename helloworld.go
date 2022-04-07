package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello Go!!!!")
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		random := rand.Intn(10)
		fmt.Println(random)
	}

}
