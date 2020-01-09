package main

import (
	"fmt"
	"time"
)


func main() {
	var now = time.Now()
	fmt.Println("Year is ",now.Year(), "Month is ", now.Month())

}

