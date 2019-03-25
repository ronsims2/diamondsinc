package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

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

	playRound()

}

func playRound() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("")
	fmt.Println("A) Buy merch")
	fmt.Println("B) Run a Flash sale")
	fmt.Println("C) Run an Ad campaign")
	fmt.Println("E) Check the Books")
	fmt.Println("")

	answer, _ := reader.ReadString('\n')

	return answer

}