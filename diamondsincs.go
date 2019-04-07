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
var options = []string{"a", "b", "c", "d", "e", "f", "g"}

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
	time.Sleep(1750 * time.Millisecond)
	PrintDelim("=", 80)
	fmt.Println("As the store manager it's your job to make sure that our customers are happy. \nIt is important to make sure that you stock the right merch and properly engage the consumers.")
	time.Sleep(1750 * time.Millisecond)
	PrintDelim("=", 80)
	fmt.Println("What would you like to do first? I recommend that you check the books.")

	//Start game loop
	for !gameOver {
		//@todo Add chance element here, adjust costs
		answer := showOptions()

		if !validateAnswer(answer, options) {
			fmt.Println("Invalid answer, try again.")
			continue
		}

		if answer == "a" {
			for true {
				answerA := showBuyMerch()

				//buyAnswers := append(options[0:3], "e")
				buyAnswers := options[0:4]

				if validateAnswer(answerA, buyAnswers) {
					for true {
						purchaseResult := askQty(answerA)
						if purchaseResult {
							break
						} else {
							ClearScreen()
							fmt.Println("Error submiting order, please try again.")
						}
					}

					break
				} else {
					ClearScreen()
					fmt.Println("Invalid answer, try again.")
				}
			}
		}

		if answer == "d" {
			//This action shouldn't advancce round
			showBooks()
			continue
		}

		if answer == "e" {
			for true {
				setPriceAnswer := askSetPrices()

				if setPriceAnswer {
					continue
				}
			}
		}

		if answer == "f" {
			playRound()

			round++
			if round == 3 {
				gameOver = true
			}

			break
		}

		if answer == "g" {
			fmt.Println("Sorry to see you go!")
			ClearScreen()
			os.Exit(0)
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
	fmt.Println("B) Run a Flash Sale") //show set prices with fee paid to run campaign
	fmt.Println("C) Run an Ad campaign")
	fmt.Println("D) Check the Books")
	fmt.Println("E) Set prices")
	fmt.Println("F) Play")
	fmt.Println("G) Quit")
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

func askQty(option string) bool {
	reader := bufio.NewReader(os.Stdin)
	options := map[string]string{
		"A": "Necklaces | Cost: $" + strconv.Itoa(necklaceCost),
		"B": "Buy Rings: | Cost: $" + strconv.Itoa(ringCost),
		"C": "Buy Bracelets: | Cost: $" + strconv.Itoa(braceletCost),
		"D": "Buy Watches: | Cost: $" + strconv.Itoa(watchCost),
		"E": "Buy Earrings : | Cost: $" + strconv.Itoa(earringCost),
	}

	PrintDelim("*", 80)
	fmt.Println("Money: $" + strconv.Itoa(cash))
	PrintDelim("-", 80)
	fmt.Println("How many " + options[option] + " do you want to buy?")
	PrintDelim("*", 80)

	answer, err := reader.ReadString('\n')
	if err != nil {
		return false
	}

	answer = strings.Trim(strings.ToLower(answer), "\n")

	qty, qtyErr := strconv.ParseInt(answer, 10, 64)

	if qtyErr != nil {
		return false
	}

	itemCost := 0
	switch option {
	case "a":
		itemCost = necklaceCost
	case "b":
		itemCost = ringCost
	case "c":
		itemCost = braceletCost
	case "d":
		itemCost = watchCost
	case "e":
		itemCost = earringCost
	}

	total := int(qty) * itemCost
	//fmt.Println("item cost: " + strconv.Itoa(itemCost))
	//fmt.Println("total paid: " + strconv.Itoa(total) + " | " + "QTY: " + strconv.Itoa(int(qty)))

	if total <= cash {
		cash = cash - total

		switch option {
		case "a":
			necklaceQty += int(qty)
		case "b":
			ringQty += int(qty)
		case "c":
			braceletQty += int(qty)
		case "d":
			watchQty += int(qty)
		case "e":
			earringQty += int(qty)
		}

	} else {
		return false
	}

	return true
}

func askSetPrices() bool {
	reader := bufio.NewReader(os.Stdin)

	PrintDelim("*", 80)
	fmt.Println("A) Set necklaces price | Cost: $" + strconv.Itoa(necklaceCost) + " | Current price: $" + strconv.Itoa(necklaceRetail))
	fmt.Println("B) Set rings price | Cost: $" + strconv.Itoa(ringCost) + " | Current price: $" + strconv.Itoa(ringRetail))
	fmt.Println("C) Set bracelets price: | Cost: $" + strconv.Itoa(braceletCost) + " | Current price: $" + strconv.Itoa(braceletRetail))
	fmt.Println("D) Set watches price: | Cost: $" + strconv.Itoa(watchCost) + " | Current price: $" + strconv.Itoa(watchRetail))
	fmt.Println("E) Set Earrings price : | Cost: $" + strconv.Itoa(earringCost) + " | Current price: $" + strconv.Itoa(earringRetail))
	PrintDelim("*", 80)

	answer, err := reader.ReadString('\n')

	if err != nil {
		return false
	}

	if validateAnswer(answer, options[0:4]) {
		return false
	}

	price, priceErr := reader.ReadString('\n')

	if priceErr != nil {
		return false
	}

	priceVal, priceValErr := strconv.ParseInt(price, 10, 64)

	if priceValErr != nil {
		return false
	}

	switch answer {
	case "a":
		necklaceRetail = int(priceVal)
	case "b":
		ringRetail = int(priceVal)
	case "c":
		braceletRetail = int(priceVal)
	case "d":
		watchRetail = int(priceVal)
	case "e":
		earringRetail = int(priceVal)
	}

	return true
}

func playRound() {
	//roll dice
	//pick chance
	//calc results

}
