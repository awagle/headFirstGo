package main

import (
	"fmt"
	"reflect"
)

func main() {
	var length float64=3.1
	var width int = 2

	fmt.Println("Type of length is ",reflect.TypeOf(length))
	fmt.Println("Type of width is ",reflect.TypeOf(width))

	fmt.Println("Conversion of width to float", reflect.TypeOf(float64(width)))
}
