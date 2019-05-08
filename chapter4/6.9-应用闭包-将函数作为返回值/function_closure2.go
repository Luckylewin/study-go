package main

func main() {
	// 在闭包中使用到的变量可以在函数内部声明 也可以在函数外部声明
	// 这样闭包函数就能够被应用到整个集合的元素上，并修改它们的值。然后这些变量就可以用于表示或计算全局或平均值。
	var g int
	func(i int) {
		sum := 0
		for j := 0; j <= i; j++ {
			sum += j
		}
		g = sum
	}(100)
	println(g)
}