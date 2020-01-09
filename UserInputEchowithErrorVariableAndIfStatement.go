package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	//the error is stored in err variable and logged
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(input)
}
