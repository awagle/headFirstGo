package main

import "fmt"

func main() {
	number := 2
	doubleByReference(&number)
	fmt.Println("Doubled number from outside the function is ", number)
}

// Now Passing By reference
func doubleByReference(numberPointer *int) {
	//Read star as value At
	*numberPointer = (*numberPointer * 2)
	fmt.Println("Doubled numberPointer from inside the function is ", *numberPointer)
}
