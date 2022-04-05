package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string
	name = "Air Asia Academy"
	fmt.Println("What if variable is declared and not used! " + name)
	fmt.Println(reflect.TypeOf(name))

	//auto type inference
	score := 88
	fmt.Println(reflect.TypeOf(score))

	location := "KL"
	fmt.Println(reflect.TypeOf(location))
	//creating the message by combining variables
	fmt.Println(name + "headquarters are in" + location)
	fmt.Println(name, "headquarters are in", location)
	//using format specifiers to print the message
	fmt.Printf("%v headquarters are in %v \n", name, location)
	fmt.Printf("I am not in new line \n")
	fmt.Printf("%50v headquarters are in %v \n", name, location)
	fmt.Printf("%-50v headquarters are in %v \n", name, location)
	//using format specifier to get the datatype
	fmt.Printf("Datatype of %v is %T \n", name, name)
	fmt.Printf("%T \n", location)
	fmt.Printf("%T \n", score)
	life := 88.888
	fmt.Printf("%T \n", life)
	check := true
	fmt.Printf("%T \n", check)

	//constants
	const PI = 3.14
	fmt.Printf("%T \n", PI)
	fmt.Println("PI:", PI)
	//single line comment
	/*
		constant values
		cannot be
		changed
	*/
	//PI = 22 / 7

}
