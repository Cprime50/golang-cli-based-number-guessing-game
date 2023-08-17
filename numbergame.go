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
	now := time.Now()
	fmt.Println("Date:", now.Year(), now.Month(), now.Day())
	fmt.Println("Time:", now.Hour(), ":", now.Minute(), ":", now.Second())

	seedSecs := time.Now().Unix()
	fmt.Println("THE GAME STARTS IN 5 SECS. GETT READDYY!!")

	//Timer animation
	timeAnimation := []string{
		"5.........", "4.......", "3....", "2..", "1.\n",
	}

	for _, timeCount := range timeAnimation {
		fmt.Printf("\r%s", timeCount)
		time.Sleep(time.Second)
	}

	//Random Number Guessing game

	//while loop to keep our window open
	for true {
		rand.NewSource(seedSecs)
		randNum := rand.Intn(50) + 1

		//Game logic loop
		for true {
			fmt.Print("Guess a number between 0 and 50\n")
			yourHints := []string{
				"Theres nothing here",
				"What comes first in races?",
				"A pair of ducks. ",
				"A magic number, the Trinity. ",
				"A lucky four-leaf clover. ",
				"A high five for this number.",
				"A roll of the dice, a single dot.",
				"A lucky number in many cultures.",
				"Two snowmen holding hands. ",
				"A cat's lives plus a tail. ",
				"A decade, all digits in a row. ",
				"The two ones stand side by side. ",
				"A dozen donuts in a box.  ",
				"Unlucky for some, but not for all. ",
				"Valentine's Day, a heart in the air.",
				"A driver's permit age. ",
				"Sweet sixteen, a milestone year. ",
				"Seventeen magazine, a teen's delight. ",
				"You can vote at this age.",
				"Prime and unique, almost 20.",
				"Double decades, a score.",
				"A blackjack hand, aces and faces.",
				"A pair of ducks, twice.",
				"A prime number, not divisible.",
				"Twenty-four hours in a day.",
				"A quarter-century milestone.",
				"Halfway to fifty.",
				"A cube's sides, three by three.",
				"The age when many start to thrive.",
				"A prime, a unique mark.",
				"A trio of decades, counting on.",
				"Halloween, a witch's magic.",
				"Like chess pieces on a board.",
				"Double digit, repeating.",
				"A TV show's address, the Simpson's.",
				"An odd number in the middle.",
				"Squared, six times six.",
				"A prime, a secret code.",
				"Two threes side by side.",
				"A prime, not divisible.",
				"Over the hill, halfway to eighty.",
				"A prime, a mystery.",
				"The answer to life, the universe, and everything.",
				"A prime, one after forty-two.",
				"Double digits, repeating fours.",
				"A multiple of nine, half of ninety.",
				"Just shy of fifty.",
				"A prime, one before forty-eight.",
				"Even and divisible, a box shape.",
				"Seven squared, a perfect square. ",
				"A half-century celebration. ",
			}
			for index, hints := range yourHints {
				if randNum == index {
					fmt.Printf("Hint: %s \n", hints)
				}

			}

			reader := bufio.NewReader(os.Stdin)
			yourGuess, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			yourGuess = strings.TrimSpace(yourGuess)
			intGuess, err := strconv.Atoi(yourGuess)
			if err != nil {
				fmt.Println("Seems you've entered a wrong value\n Please enter a postive integer Number")
				continue
			}
			if intGuess > randNum {
				fmt.Println("Nice try but your guess is a bit too high, guess a Lower Number")
				(time.Sleep(time.Second * 2))
			} else if (intGuess > randNum) && (intGuess < randNum+3) || (intGuess < randNum) && (intGuess > randNum-3) {
				fmt.Println("You are very very close, try again")
				time.Sleep(time.Second * 2)
			} else if intGuess < randNum {
				fmt.Println("Nice try but your guess is a bit too low, guess a Higher Number")
				time.Sleep(time.Second * 2)
			} else {
				youWin := []string{
					"You", "Guessed", "it", "Right", "CONGRATULATIONSSSS!!!", "YOU", "WON!!!\n",
				}
				for _, winText := range youWin {
					fmt.Print(winText + " ")
					(time.Sleep(time.Millisecond * 100))
				}
				break

			}

		}
		//End of Logic Loop

		fmt.Println("Would you like to play again?")
		time.Sleep(time.Second)
		fmt.Println("Press y if yes, press any other button if you want to exit the game")

		readNext := bufio.NewReader(os.Stdin)
		nextOption, err := readNext.ReadString('\n')
		if err != nil {
			fmt.Println("It seems you may have entered an Unknown command\n Please enter y to play again or n to exit the game")
		}
		nextOption = strings.TrimSpace(nextOption)

		if nextOption == "y" {
			fmt.Println("That's nice. Seems you're having a good time")
			time.Sleep(time.Second * 2)
			continue
		} else if nextOption != "y" {
			fmt.Println("Hope you had lot's of fun playing, come again soon")
			exitAnimate := []string{
				"Exiting.", "Exiting..", "Exiting...", "Exiting....", "Exiting.", "Exiting..", "Exiting...", "Exiting....",
				"Exiting.", "Exiting..", "Exiting...", "Exiting....", "Exiting.", "Exiting..", "Exiting...", "Exiting....",
				"Exiting.", "Exiting..", "Exiting...", "Exiting....", "Exiting.", "Exiting..", "Exiting...", "Exiting....",
			}
			for _, animate := range exitAnimate {
				fmt.Printf("\r%s", animate)
				time.Sleep(time.Millisecond * 100)
			}
			time.Sleep(time.Second * 4)
			break
		}
	}

}
