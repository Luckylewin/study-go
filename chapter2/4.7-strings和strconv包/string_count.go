package main

import (
	"strings"
	"fmt"
)

func main() {
	var str string = "Hello,how is it goging,hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}