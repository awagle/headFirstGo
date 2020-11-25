package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Starting HTTP Server")
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
