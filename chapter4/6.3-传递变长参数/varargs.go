package main

import "fmt"

func main() {
	// 写一个函数，该函数接受一个变长参数并对每个元素进行换行打印。
	names := []string{"lychee","apple","pear","watermalon"}
	F1(names...)
}

// 一个接受变长参数的函数可以将这个参数作为其它函数的参数进行传递
// 变长参数可以作为对应类型的 slice 进行二次传递。
func F1(s ...string) {
	F2(s...)
	F3(s)
}

func F2(s ...string) {
	for index, value := range s {
		fmt.Printf("%d-%s\n",index,value)
	}
}

func F3(s []string) {
	for index, value := range s {
		fmt.Printf("%d-%s\n",index,value)
	}
}