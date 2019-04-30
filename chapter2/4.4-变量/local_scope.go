package main

var a = "G"

func main() {
	n() // G
	m() // O 这是因为 a := "O" 作用域为局部
	n() // G 
}

func n() {
	print(a)
}

func m() {
	a := "O"
	print(a)
}