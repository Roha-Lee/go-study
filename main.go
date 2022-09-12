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
	repeatMe("Hello", "My", "Name", "is", "Roha")
	totalLength, upperName = nakedReturn("gru")
	fmt.Println(totalLength, upperName)

}

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func nakedReturn(name string) (length int, uppercase string) {
	defer fmt.Println("I am done.")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
