package main

var a = "G"

func main() {
	n() // G
	m() // O 这是因为 a = "O" 作用域为全局
	n() // O
}

func n() {
	print(a)
}

func m() {
	a := "O"
	print(a)
}