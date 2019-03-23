package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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
