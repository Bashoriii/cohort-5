package main

import (
	"fmt"
	minichallenge3 "learn/mini-challenge-3"
	"strings"
)

func fizzBuzz(total int) {
	for i := 1; i <= total; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}

// Count a word -- Mini Challenge 2
func countWord(randomWords string) {
	myVar := randomWords
	varSplit := strings.Split(myVar, "")

	for _, arr := range varSplit {
		fmt.Println(arr)
	}

	myMap := map[string]int{}
	for _, val := range varSplit {
		myMap[val]++
	}

	fmt.Println(myMap)
}

// 2 GO Routines -- Mini Challenge 4
// func

func main() {

	// fizzBuzz(15)
	// countWord("selamat pagi")
	minichallenge3.MiniChallenge3()
}
