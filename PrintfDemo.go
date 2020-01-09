package main

import "fmt"

func main() {
	fmt.Printf("%s world \n", "Hello")

	fmt.Printf("%d is an integer and %s is a String", 29, "Aditya")

	fmt.Printf("\n %T will give you the type of 25", 25)

	fmt.Printf("\n %v will automatically take the right type for 25\n", 25)

	//Let's do a table
	fmt.Println("--------------------------------------------")
	fmt.Printf("%12s | %12s | %s\n", "Name", "Role", "Years")
	fmt.Println("--------------------------------------------")
	fmt.Printf("%12s | %12s | %d\n", "Aditya", "BW  Lead", 13)
	fmt.Printf("%12s | %12s | %d\n", "Vijay", "Flogo Lead", 12)
}
