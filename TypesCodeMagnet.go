package main

import (
		"fmt"
		)

func main() {

	//Normal Variable declarations with value assignments
	var originalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", originalCount , "apples.")
	fmt.Println("Some jerk ate", eatenCount , "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")


	//Short Variable Declarations
	message := "Hello World"

	fmt.Println(message)
}
