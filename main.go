package main

import (
	"fmt"
	"strings"
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

func main() {
	// fmt.Println(multiply(45, 10))
	// something.SayHello()
	// totalLength, _ := lenAndUpper("Sithum")
	// fmt.Println(totalLength)

	// repeatMe("Sithum", "Golang", "Learning", "Hustling")

	// length, upper := lenAndUpper("Sithum")
	// fmt.Println(length, upper)

	result := superAdd(45, 885, 125, 485, 68, 748, 45)
	fmt.Println(result)
}