package main

import (
	"fmt"
	webserver "learn/sesi-5"
	"net/http"
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

func main() {

	// fmt.Println("Waiting for goroutines to finish...")
	// wg.Wait()
	// fmt.Println("Done!")

	// myData1()
	// myData2()

	// fizzBuzz(15)
	// countWord("selamat pagi")
	// minichallenge3.MiniChallenge3()

	// 2 Goroutines -- Mini Challenge 4
	// minichallenge4.DisorderArr()
	// minichallenge4.ArrangedArr()
	// minichallenge4.TestDum()

	// ======= Web Server -- Mini Challenge 5 =======
	http.HandleFunc("/", webserver.GetCustomers)
	http.HandleFunc("/create", webserver.CreateCust)
	http.HandleFunc("/login", webserver.LoginCust)
	http.HandleFunc("/profile", webserver.ProfilePage)
	http.HandleFunc("/login-failed", webserver.HandleLoginFail)

	fmt.Println("Listened on port 8090")
	http.ListenAndServe(":8090", nil)
	// ======= Web Server -- Mini Challenge 5 =======

	// wg.Add(1)
	// go egKedua(&wg)
	// wg.Wait()
}
