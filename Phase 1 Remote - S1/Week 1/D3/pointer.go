package main

import "fmt"

func main() {
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson = ", firstPerson)
	fmt.Println("secondPerson = ", *secondPerson)

	*secondPerson = "David"
	fmt.Println("firstPerson = ", firstPerson)
	fmt.Println("secondPerson = ", *secondPerson)
}
