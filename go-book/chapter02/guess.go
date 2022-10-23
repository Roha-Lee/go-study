// guess프로그램은 플레이어가 난수를 맞히는 게임입니다.
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
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number from 1 to 100")
	fmt.Println("Can you guess it?")
	reader := bufio.NewReader(os.Stdin)
	const MAX_GUESS int = 10
	success := false
	for guesses := 0; guesses < MAX_GUESS; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Guess a number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}

		if guess < answer {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > answer {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Correct!")
			break
		}
	}
	if !success {
		fmt.Println("SOrry, you didn't guess my number. It was:", answer)
	}

}
