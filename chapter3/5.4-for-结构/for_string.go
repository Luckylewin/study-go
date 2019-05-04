package main

import "fmt"

func main() {
	str := "Golang is a beautiful language!"
	fmt.Printf("The length of str is:%d", len(str))

	for ix:=0; ix<len(str); ix++ {
		fmt.Printf("Character on postion %d is:%c\n", ix, str[ix])
	}

	str2 := "日本语"
	fmt.Printf("The length of str is:%d", len(str2))
	for ix:=0; ix<len(str2); ix++ {
		fmt.Printf("Character on postion %d is:%c\n", ix, str2[ix])
	}
}