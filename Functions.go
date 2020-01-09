package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	length, breadth := 5.0, 3.0
	area, err := areaRectangle(length, breadth)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Area of rectangle with length ", length, "and breadth ", breadth, "is =", area)
	}

	number := 2
	double(number)
	fmt.Println("Doubled number from outside the function is ", number)
}

//Example function with Multiple return types
func areaRectangle(length float64, breadth float64) (area float64, err error) {
	if length == 0 {
		err = errors.New("Length cannot be 0")
		return 0, err
	} else if breadth == 0 {
		err = errors.New("Breadth cannot be 0")
		return 0, err
	} else {
		return length * breadth, nil
	}

}

//Example function which is Passing by Value (no pointers)
func double(number int) {
	number = number * 2 // This will not change the original value since it's passing by value and not reference
	fmt.Println("Doubled number from inside the function is ", number)
}
