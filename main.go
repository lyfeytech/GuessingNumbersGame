package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	secretNumber := getRandomNumber()
	//fmt.Println(secretNumber)

	for matching := false; !matching; {

		guess := getUserInput()
		fmt.Println(secretNumber, guess)

		matching = isMatching(secretNumber, guess)
		//fmt.Println(matching)
	}
}

func isMatching(secretNumber, guess int) bool {
	if guess > secretNumber {
		fmt.Println("Your guess is too big")
		return false
	} else if guess < secretNumber {
		fmt.Println("Your guess is too small")
		return false
	} else {
		fmt.Println("YOU GOT IT!")
		return true
	}
}

func getUserInput() int {
	fmt.Print("Please input your guess: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Failed to parse your input", err)
	} else {
		fmt.Println("You guessed:", input)
	}
	return input

}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}
