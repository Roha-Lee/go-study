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
	superAdd(1, 3, 2, 4, 5, 6, 1)
	fmt.Println(canIDrink(31), canIDrink(17))

	names := []string{"Roha", "Gru"}
	names = append(names, "Lee")
	fmt.Println(names)
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

func superAdd(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return total
}

func canIDrink(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge > 19:
		return true
	case koreanAge <= 19:
		return false
	}
	return false
}
