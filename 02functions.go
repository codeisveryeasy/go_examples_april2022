package main

import "fmt"

func main() {
	welcome()

	total := addition(5, 4)
	fmt.Println("Total: ", total)

	n1, n2, n3 := multipleValues("Mars")
	fmt.Printf("%v supports life (%v) is %v from sun \n", n1, n2, n3)

	planet := "Mars"
	n4, n5 := multipleValuesWithoutVariables(planet)
	fmt.Printf("%v supports life (%v) is %v from sun \n", planet, n4, n5)

	planet1 := "Mercury"
	n6, n7 := multipleValuesWithoutVariables(planet1)
	fmt.Printf("%v supports life (%v) is %v from sun \n", planet1, n6, n7)
}

func addition(i1 int, i2 int) int {
	return i1 + i2
}

func welcome() {
	fmt.Println("Welcome to Go function!")
}

func multipleValues(message string) (string, bool, int) {
	supportLife := true
	distanceFromrSun := 5646542165465
	return message, supportLife, distanceFromrSun
}

func multipleValuesWithoutVariables(message string) (supportLife bool, distanceFromSun int) {
	planet := message
	if planet == "Mars" {
		supportLife = true
	}
	if planet == "Mercury" {
		supportLife = false
	}

	distanceFromSun = 46521321
	return
}
