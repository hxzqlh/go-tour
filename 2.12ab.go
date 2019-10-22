package main

import (
	"fmt"
)

var numChan = make(chan bool, 1)
var strChan = make(chan bool, 1)
var finish = make(chan bool, 1)
var englishLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func printNumber() {
	for i := 1; i < 27; i += 2 {
		<-numChan
		fmt.Printf("%d%d", i, i+1)
		strChan <- true
	}
}

func printString() {
	size := len(englishLetters)
	for i := 0; i < size; i += 2 {
		<-strChan
		fmt.Printf("%c%c", englishLetters[i], englishLetters[i+1])
		numChan <- true
	}
	finish <- true
}

func main() {
	numChan <- true
	go printNumber()
	go printString()
	<-finish
	fmt.Println()
}
