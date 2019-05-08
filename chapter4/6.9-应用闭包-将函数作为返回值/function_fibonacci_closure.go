package main

func main() {
	fib := fibonacci()
	for i:=0; i<=10; i++ {
		println(fib())
	}
}

func fibonacci() func() int {
	var prevTwo int = 0
	var brevOne int = 1
	return func() int {
		tmp := prevTwo
		prevTwo = brevOne
		brevOne = tmp + prevTwo

		return tmp
	}
}