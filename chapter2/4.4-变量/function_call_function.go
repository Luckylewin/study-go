package main

var a string

func main() {
	a = "G"
	print(a) // 全局变量
	f1()
 	// G O G
}

func f1() {
	a := "O"
	print(a) // 局部变量
	f2()
}

func f2() {
	print(a) // 访问全局变量
}