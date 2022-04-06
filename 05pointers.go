package main

import "fmt"

func main() {

	name := "Air Asia Academy"
	//pass by value
	printMessage(name)

	//pass by reference
	fmt.Println("Name:", name)
	fmt.Println("Name address: ", &name)
	printMessageRef(&name)
	printMessage(name)
}

func printMessage(message string) {
	fmt.Println(message)
}

func printMessageRef(message *string) {
	fmt.Println(message)
	fmt.Println(*message)
	*message = "Air Asia Academy, KL"
}
