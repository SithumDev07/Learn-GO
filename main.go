package main

import (
	"fmt"
	"strings"
)

// import (
// 	"fmt"

// 	"github.com/SithumDev07/LearnGO/something"
// )

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiply(a int, b int) int {
	return a * b
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// fmt.Println(multiply(45, 10))
	// something.SayHello()
	// totalLength, _ := lenAndUpper("Sithum")
	// fmt.Println(totalLength)

	repeatMe("Sithum", "Golang", "Learning", "Hustling")
}
