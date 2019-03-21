package main

import (
	"bufio"
	"fmt"
	"os"
)

//Print the content of a file
func PrintArt(artName string) {
	artFile, _ := os.Open(artName + ".art")
	scanner := bufio.NewScanner(artFile)

	for scanner.Scan() {
		//fmt.Printf(scanner.Text())
		fmt.Println(scanner.Text())
	}
}
