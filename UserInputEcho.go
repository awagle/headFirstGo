package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	//Use Go's blank identifier for storing error and discarding it
	input, _ := reader.ReadString('\n')

	fmt.Println(input)
}
