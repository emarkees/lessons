package main

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
	target := rand.Intn(100) + 1
	fmt.Println("I have choosen a random number between 1 to 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses <= 10; guesses++ {
		fmt.Println("You have", 10-guesses, "left")
		fmt.Print("Make a guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Trim Spaces
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess == target {
			success = true
			fmt.Println("Good JOB! Your guess RIGHT")
			break
		} else if guess > target {
			fmt.Println("Opps! your guess was HIGH")
		} else {
			fmt.Println("Ops! you guess LOW")
		}
	}

	if !success {
		fmt.Println("Sorry! You didn't guess the number right. It was :", target)
	}

}
