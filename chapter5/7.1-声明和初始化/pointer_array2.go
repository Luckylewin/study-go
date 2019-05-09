package main

import "fmt"

func main() {
	for i:=0; i<3; i++ {
		// 可以取任意数组常量的地址来作为指向新实例的指针
		fp(&[3]int{i, i * i, i * i * i})
	}
}

func fp(a *[3]int) { fmt.Println(a)}