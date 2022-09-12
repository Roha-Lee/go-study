package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("roha")
	fmt.Println(totalLength, upperName)

}

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
