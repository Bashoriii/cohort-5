package main

import (
	"fmt"
	"strings"
)

func main() {

	// 	FizzBuzz -- Mini Challenge 1
	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("FizzBuzz")
	// 	} else if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	} else if i%3 == 0 {
	// 		fmt.Println("Fizz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	// Count a word in sentence -- Mini Challenge 2
	myVar := "chained together"
	varSplit := strings.Split(myVar, "")

	for _, arr := range varSplit {
		fmt.Println(arr)
	}

	myMap := map[string]int{}
	for _, val := range varSplit {
		myMap[val]++
	}

	fmt.Println(myMap)

	// Define an array
	// myArr := [3]int{49, 51, 12}

	// for i, arr := range myArr {
	// 	fmt.Println(i, arr)
	// }

	// // Define a slice
	// mySlice := []string{"Lala", "Lulu", "Lele", "Lili", "Lolo"}
	// for i, arr := range mySlice {
	// 	fmt.Println(i, arr)
	// }

	// var num int = 4
	// var num1 *int = &num

	// fmt.Println(num)
	// fmt.Println(num1)

	// mySlice := []int{10, 12, 49, 35, 27, 16}
	// fmt.Println("Original Slice:", mySlice)

	// dipilihArr := mySlice[3:6]
	// fmt.Println("After Selection:", dipilihArr)

	// for _, arr := range mySlice {
	// 	fmt.Println(arr)
	// }
}
