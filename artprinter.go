package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
