package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode/utf8"
)

//PrintArt prints the content of a file
func PrintArt(artName string) {
	artFile, _ := os.Open(artName + ".art")
	scanner := bufio.NewScanner(artFile)

	for scanner.Scan() {
		//fmt.Printf(scanner.Text())
		fmt.Println(scanner.Text())
	}
}

//PrintDelim prints a line of the same character to break up sections
func PrintDelim(delimChar string, length int) {
	var delim bytes.Buffer

	for i := 0; i < length; i++ {
		delim.WriteString(delimChar)
	}

	fmt.Println(delim.String())
}

//ClearScreen clears the console
func ClearScreen() {
	currentOs := runtime.GOOS
	if strings.Contains(strings.ToLower(currentOs), "windows") {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		//handle Mac and Linux
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

//PrintArtWithText repalces tokens with 33 chars of text per line
func PrintArtWithText(artName string, line1 string, line2 string) {
	artFile, _ := os.Open(artName + ".art")
	scanner := bufio.NewScanner(artFile)
	token := "#################################"
	line := line1

	for scanner.Scan() {
		text := scanner.Text()
		var message string
		if strings.Contains(text, token) {
			textLength := utf8.RuneCountInString(line)
			if textLength < 33 {
				pad := math.Floor(float64((33 - textLength) / 2))
				message = strings.Repeat(" ", int(pad)) + line
				message += strings.Repeat(" ", int(pad))

				messageLength := utf8.RuneCountInString(message)
				if messageLength < 33 {
					message = message + strings.Repeat(" ", 33-messageLength)
				}
			}

			message = strings.Replace(text, token, message, 1)
			fmt.Println(message)

			line = line2
		} else {
			fmt.Println(text)
		}
	}
}
