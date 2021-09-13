package main

import (
	"fmt"
	"strings"

	// "github.com/SithumDev07/LearnGO/accounts"
	"github.com/SithumDev07/LearnGO/mydict"
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
	dictionary := mydict.Dictionary{"first": "First Word"}

	word := "hello"
	defintion := "Greeting"

	error := dictionary.Add(word, defintion)

	if error != nil {
		fmt.Println(error)
	} 
	
	def, error := dictionary.Search(word)
	fmt.Println(def)
	error2 := dictionary.Add(word, defintion)
	
	if error2 != nil {
		fmt.Println(error2)
	} 
}
