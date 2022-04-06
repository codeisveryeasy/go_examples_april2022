package main

import "fmt"

func main() {
	counter := 40
	if counter > 40 {
		fmt.Println("greater than 40")
	} else {
		fmt.Println("greater that or equal to 40")
	}

	if check := true; check {
		fmt.Println("Can do")
	} else {
		fmt.Println("Cannot do")
	}

	if total := 80; total > 70 && counter > 40 {
		fmt.Println("Can do ", total)
	} else {
		fmt.Println("Cannot do", total)
	}
}
