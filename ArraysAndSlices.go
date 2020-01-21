package main

import "fmt"

func main() {
	//Declaring an array
	var myArray [5]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4

	//Declaring a Slice

	//var mySlice []int
	//mySlice = make([]int,5)

	//shorter declaration of
	mySlice := make([]int, 5)
	mySlice[0] = 5
	mySlice[1] = 6
	mySlice[2] = 7
	mySlice[3] = 8

	//Printing contents of an Array using a for loop and range
	for _, myIntegers := range myArray {
		fmt.Println(myIntegers)
	}

	printSlice(mySlice)

	//add elements to the slice
	mySlice[4] = 9
	mySlice = append(mySlice, 10)
	printSlice(mySlice)

}

func printSlice(slice []int) {
	//Printing contents of a Slice using a for loop and range
	for _, myIntegers := range slice {
		fmt.Println(myIntegers)
	}
}
