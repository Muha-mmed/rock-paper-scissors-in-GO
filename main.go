// ROCK PAPER SCISSORS GAME

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	playerScore := 0
	computerScore := 0
	option := []string{"rock", "paper", "scissors"}
	var playerChoice string

	for {
		computerChoice := option[rand.Intn(len(option))]

		fmt.Print("choose an option, rock, paper and scissors: ")
		fmt.Scan(&playerChoice)
		fmt.Printf("Computer Choice:  %v \n", computerChoice)
		fmt.Printf("player Choice:  %v  \n", playerChoice)
		if playerChoice == computerChoice {
			fmt.Println("no winner(tie)")
			fmt.Printf("game score Player %v, computer %v \n", playerScore, computerScore)
		} else if playerChoice == "rock" && computerChoice == "paper" || playerChoice == "paper" && computerChoice == "rock" || playerChoice == "scissor" && computerChoice == "paper" {
			fmt.Println("player wins")
			playerScore++
			fmt.Printf("game score Player %v, computer %v \n", playerScore, computerScore)
		} else {
			fmt.Println("computer Wins")
			computerScore++
			fmt.Printf("game score Player %v, computer %v \n", playerScore, computerScore)
		}
	}
}
