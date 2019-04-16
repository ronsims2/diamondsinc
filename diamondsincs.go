package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
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
var circumstances []Circumstance
var showIntructions = false

var adBought = false
var campaignRan = false

func main() {
	//savekey := "Um9uYWxkIGlzIHRoZSBjb29sZXN0Lgo" // Add = to end to get full string, decode for fun message!
	ClearScreen()

	//Generate circumstances
	circumstances = GenerateCircumstances(10)

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
		answer := showOptions(showIntructions)
		showIntructions = true
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

		if answer == "b" {
			//start campaign
			if cash >= flashsaleCost && !campaignRan {
				cash -= flashsaleCost
				campaignRan = true
				PrintDelim("*", 80)
				fmt.Println("Flash sale scheduled!")
			} else {
				PrintDelim("*", 80)
				fmt.Println("You cannot run a flash sale right now.")
			}
		}

		if answer == "c" {
			//buy ads
			if cash >= adCost && !adBought {
				cash -= adCost
				adBought = true
				PrintDelim("*", 80)
				fmt.Println("Ads are running everywhere now!")
			} else {
				PrintDelim("*", 80)
				fmt.Println("You cannot purchase an ad right now.")
			}
		}

		if answer == "d" {
			//This action shouldn't advancce round
			showBooks()
			continue
		}

		if answer == "e" {
			for true {
				ClearScreen()
				setPriceAnswer := askSetPrices()

				if setPriceAnswer {
					break
				}
			}
		}

		if answer == "f" {
			playRound()
			if !checkStats() {
				gameOver = true
				ClearScreen()
				PrintArtWithText("manager", "Things aren't working out...", "I gotta' let you go...")
				PrintArt("fin")
				PrintDelim("*", 80)
				fmt.Println("I guess I didn't survie the 2019 Quarte Quell...may the odds ever be in YOUR favor!")
				PrintDelim("*", 80)
				break
			}
			fmt.Println("Round: " + strconv.Itoa(round+1))
			round++
			if round == len(circumstances) {
				gameOver = true
				ClearScreen()
				PrintArt("fin")
				PrintDelim("*", 80)
				fmt.Println("I guess I didn't survie the 2019 Quarte Quell...may the odds ever be in YOUR favor!")
				PrintDelim("*", 80)
				break

			}
			continue
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

func showOptions(showIntruction bool) string {
	reader := bufio.NewReader(os.Stdin)

	if showIntruction {
		PrintDelim("=", 80)
		fmt.Println("What would you like to do next?")
	}

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
	ClearScreen()
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

	time.Sleep(300 * time.Millisecond)
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
	answer = strings.Trim(strings.ToLower(answer), "\n")

	if err != nil {
		return false
	}

	if !validateAnswer(answer, options[0:4]) {
		fmt.Println("Incorrect answer, please try again.")
		time.Sleep(2000 * time.Millisecond)
		return false
	}

	PrintDelim("*", 80)
	fmt.Println("Enter a price")
	PrintDelim("*", 80)
	price, priceErr := reader.ReadString('\n')
	price = strings.Trim(strings.ToLower(price), "\n")

	if priceErr != nil {
		fmt.Println("Invaild price, please try again.")
		time.Sleep(2000 * time.Millisecond)
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

func rollDice(count int) int {
	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(count)

	return roll
}

func playRound() {
	result := rollDice(len(circumstances))
	buy(result)
	time.Sleep(5000 * time.Millisecond)
	adBought = false
	campaignRan = false
	ClearScreen()
	time.Sleep(300 * time.Millisecond)

}

func buy(diceRoll int) {
	circumstance := circumstances[diceRoll]
	boost := 0.0

	if adBought {
		boost = 0.25
	}

	if campaignRan {
		boost = 0.33
	}

	influence := circumstance.influence + boost
	shoppersMoney := int64(math.RoundToEven(float64(circumstance.goal) * influence))
	lowestPrice := checkMinPrice()
	dailySales := 0

	for shoppersMoney > lowestPrice {
		if int64(necklaceCost) <= shoppersMoney && necklaceQty > 0 {
			necklaceQty--
			shoppersMoney -= int64(necklaceCost)
			dailySales += necklaceCost
		}

		if int64(ringCost) <= shoppersMoney && ringQty > 0 {
			ringQty--
			shoppersMoney -= int64(ringCost)
			dailySales += ringCost
		}

		if int64(braceletCost) <= shoppersMoney && braceletQty > 0 {
			braceletQty--
			shoppersMoney -= int64(braceletCost)
			dailySales += braceletCost
		}

		if int64(watchCost) <= shoppersMoney && watchQty > 0 {
			watchQty--
			shoppersMoney -= int64(watchCost)
			dailySales += watchCost
		}

		if int64(earringCost) <= shoppersMoney && earringQty > 0 {
			earringQty--
			shoppersMoney -= int64(earringCost)
			dailySales += earringCost
		}
	}
	cash += dailySales
	desc := "Sales could be better."

	if shoppersMoney > int64(circumstance.goal) {
		desc = "We made our numbers."

		if adBought || campaignRan {
			desc += " Good Job!"
		}
	}

	ClearScreen()
	PrintArtWithText("manager", desc, "We made $"+strconv.Itoa(dailySales)+" today.")
}

func checkMinPrice() int64 {
	lowest := necklaceCost

	if ringCost < lowest {
		lowest = ringCost
	}

	if braceletCost < lowest {
		lowest = braceletCost
	}

	if watchCost < lowest {
		lowest = watchCost
	}

	if earringCost < lowest {
		lowest = earringCost
	}

	fmt.Println("Low func :" + string(lowest))
	return int64(lowest)
}

func checkStats() bool {
	count := 0
	if necklaceQty == 0 {
		count++
	}

	if ringQty == 0 {
		count++
	}
	if braceletQty == 0 {
		count++
	}
	if watchQty == 0 {
		count++
	}

	if earringQty == 0 {
		count++
	}

	if count == 5 && int64(cash) < checkMinPrice() {
		return false
	} else {
		return true
	}
}
