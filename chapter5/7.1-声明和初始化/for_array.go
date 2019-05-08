package main

import "fmt"

func main() {
	var arr1 [8]int

	for i:=0; i<len(arr1); i++ {
		arr1[i] = i * i
	}

	for index,value := range arr1 {
		fmt.Printf("%d-%d\n", index, value);
	}
}