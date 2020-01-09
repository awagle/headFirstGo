package main

//math/rand is the import path
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randomNumber := rand.Intn(100) //rand here is the package name

	for i := 0; i <= 10; i++ {
		//Prompt the user for a number
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input) //Required to trim the Enter key pressed at the end
		guess, err := strconv.Atoi(input)

		if guess == randomNumber {
			fmt.Println("Congratulations you correctly guessed the number")
			break
		} else {
			if guess > randomNumber {
				fmt.Println("Your guess was too high")
			} else {
				fmt.Println("Your guess was too low")
			}
		}
	}

}
