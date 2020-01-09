package main

import "fmt"

func main() {
	number := 100

	//Define the pointer
	var numberPointer *int
	//Assign the address
	numberPointer = &number
	//Print the address of the variable
	fmt.Println("Address of the variable is ", numberPointer)

	//Print the value at the address
	fmt.Println("Value of the variable is ", *numberPointer)

	//Short declaration
	shortNumberPointerVariable := &number

	//Print the value at the address
	fmt.Println("Value of the variable is (using Short declaration)", *shortNumberPointerVariable)

}
