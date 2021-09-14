package main

import (
	"fmt"
	"net/http"
	"strings"
	// "github.com/SithumDev07/LearnGO/accounts"
)

// import (
// 	"fmt"

// 	"github.com/SithumDev07/LearnGO/something"
// )

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I' Done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func multiply(a int, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	switch sriLankanAge := age + 2; sriLankanAge {
	case 10:
		return false
	case 18:
		return true
	case 50:
		return false
	}
	return false
}

type person struct {
	name    string
	age     int
	favFood []string
}


type result struct {
	url string
	status string
}

func main() {
	
	results := make(map[string]string)
	c := make(chan result)

	urls:= []string{
		"https://www.google.com/",
		"https://www.airbnb.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i:=0; i < len(urls); i++{
		resultss := <-c
		results[resultss.url] = resultss.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}


 }


func hitURL(url string, c chan<- result) {
	res, err := http.Get(url)
	status := "OK"
	if err != nil || res.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- result{
		url: url,
		status: status,
	}
}
