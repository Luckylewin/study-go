package main

import "fmt"

func main() {
	str := "Golang is a beautiful language"
	fmt.Printf("The length of this str is %d\n", len(str))
	
	for pos, char := range str {
		fmt.Printf("Character on postion %d is:%c \n", pos, char)
	}

	fmt.Println()
}