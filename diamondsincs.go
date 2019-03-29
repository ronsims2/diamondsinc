package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//The money to buy stuff
var cash = 5000
var necklaceQty = 0
var ringQty = 0
var braceletQty = 0
var watchQty = 0
var earringQty = 0

//pricelist
var necklaceCost = 250
var ringCost = 1000
var braceletCost = 500
var watchCost = 100
var earringCost = 200

var adCost = 2500
var flashsaleCost = 1500

var necklaceRetail = 0
var ringRetail = 0
var braceletRetail = 0
var watchRetail = 0
var earringRetail = 0

var gameOver = false
var round = 0
var options = []string{"a", "b", "c", "d"}

func main() {
	//savekey := "Um9uYWxkIGlzIHRoZSBjb29sZXN0Lgo" // Add = to end to get full string, decode for fun message!
	ClearScreen()
	PrintArt("logo")
	fmt.Println("")
	fmt.Println("")
	//PrintArt("diamond")
	PrintArtWithText("manager", "Welcome to Diamonds Inc!", "Let me show you around.")
	fmt.Println("I'm Leroy Shaw the owner of this fine jewelry establishment.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("I'm sorry What is your name again?")
	playerName, _ := reader.ReadString('\n')
	playerName = strings.TrimSpace(playerName)
	ClearScreen()
	fmt.Println("Well then, welcome " + playerName + " I am sure you will do fine here.")
	time.Sleep(3 * time.Second)
	PrintDelim("=", 80)
	fmt.Println("As the store manager it's your job to make sure that our customers are happy. \nIt is important to make sure that you stock the right merch and properly engage the consumers.")
	time.Sleep(3 * time.Second)
	PrintDelim("=", 80)
	fmt.Println("What would you like to do first? I recommend that you check the books.")

	//Start game loop
	for !gameOver {
		//@todo Add chance element here, adjust costs
		answer := showOptions()

		if !validateAnswer(answer, options) {
			fmt.Println("Invalid answer, try again.")
			break
		}

		if answer == "a" {
			for true {
				answerA := showBuyMerch()

				buyAnswers := append(options[0:3], "e")

				if validateAnswer(answerA, buyAnswers) {

					break
				} else {
					fmt.Println("Invalid answer, try again.")
				}
			}
		}

		if answer == "d" {
			//This action shouldn't advancce round
			showBooks()
			continue
		}
	}

}

func validateAnswer(answer string, answers []string) bool {
	result := false
	for _, v := range answers {
		if answer == v {
			result = true
			break
		}
	}

	return result
}

func showOptions() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("")
	fmt.Println("A) Buy Merch")
	fmt.Println("B) Run a Flash Sale")
	fmt.Println("C) Run an Ad campaign")
	fmt.Println("D) Check the Books")
	fmt.Println("")

	answer, _ := reader.ReadString('\n')
	answer = strings.Trim(strings.ToLower(answer), "\n")

	return answer
}

func showBooks() {
	PrintDelim("*", 80)
	fmt.Println("Money: $" + strconv.Itoa(cash))
	PrintDelim("-", 80)
	fmt.Println("Necklaces QTY: " + strconv.Itoa(necklaceQty) + " | Cost: $" + strconv.Itoa(necklaceCost) + " | Retail: $" + strconv.Itoa(necklaceRetail))
	PrintDelim("-", 80)
	fmt.Println("Rings QTY: " + strconv.Itoa(ringQty) + " | Cost: $" + strconv.Itoa(ringCost) + " | Retail: $" + strconv.Itoa(ringRetail))
	PrintDelim("-", 80)
	fmt.Println("Bracelets QTY: " + strconv.Itoa(braceletQty) + " | Cost: $" + strconv.Itoa(braceletCost) + " | Retail: $" + strconv.Itoa(braceletRetail))
	PrintDelim("-", 80)
	fmt.Println("Watches QTY: " + strconv.Itoa(watchQty) + " | Cost: $" + strconv.Itoa(watchCost) + " | Retail: $" + strconv.Itoa(watchRetail))
	PrintDelim("-", 80)
	fmt.Println("Earring QTY: " + strconv.Itoa(earringQty) + " | Cost: $" + strconv.Itoa(earringCost) + " | Retail: $" + strconv.Itoa(earringRetail))
	PrintDelim("*", 80)
}

func showBuyMerch() string {
	reader := bufio.NewReader(os.Stdin)

	PrintDelim("*", 80)
	fmt.Println("Money: $" + strconv.Itoa(cash))
	PrintDelim("-", 80)
	fmt.Println("A) Buy Necklaces | Cost: $" + strconv.Itoa(necklaceCost))
	fmt.Println("B) Buy Rings: | Cost: $" + strconv.Itoa(ringCost))
	fmt.Println("C) Buy Bracelets: | Cost: $" + strconv.Itoa(braceletCost))
	fmt.Println("D) Buy Watches: | Cost: $" + strconv.Itoa(watchCost))
	fmt.Println("E) Buy Earring : | Cost: $" + strconv.Itoa(earringCost))
	PrintDelim("*", 80)

	answer, _ := reader.ReadString('\n')
	answer = strings.Trim(strings.ToLower(answer), "\n")

	return answer
}
