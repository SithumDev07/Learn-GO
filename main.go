package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
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

func main() {
	
	// var results = make(map[string]string)

	// urls:= []string{
	// 	"https://www.google.com/",
	// 	"https://www.airbnb.com/",
	// 	"https://www.amazon.com/",
	// 	"https://www.reddit.com/",
	// 	"https://soundcloud.com/",
	// 	"https://www.facebook.com/",
	// 	"https://www.instagram.com/",
	// }

	// for _, url := range urls {
	// 	result := "OK"
	// 	err := hitURL(url)
	// 	if err != nil {
	// 		result = "FAILED"
	// 	}
	// 	results[url] = result
	// }

	// for url, result := range results {
	// 	fmt.Println(url, " ", result)
	// }

	go Count("Sithum")
	Count("Dashantha")

}

func Count(person string) {
	for i:=0; i<10; i++ {
		fmt.Println(person, " is looking good", i)
		time.Sleep(time.Second)
	}
}

var requestFail = errors.New("Request Failed")

func hitURL(url string) error{
	fmt.Println("Checking", url)
	response, err := http.Get(url)
	if err != nil || response.StatusCode >= 400{
		fmt.Println(err, response.StatusCode)
		return requestFail
	}
	return nil
}
