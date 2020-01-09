package main
//math/rand is the import path
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//If we don't add the below two lines it really doesn't generate a new random number on every run.
	seconds := time.Now().Unix();
	rand.Seed(seconds)


	randomNumber := rand.Intn(100)//rand here is the package name
	fmt.Println(randomNumber)
}
