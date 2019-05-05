package main

func main() {
	// 当我们不希望给函数起名字的适合，可以使用匿名函数
    // 例如 func(x, y int) int { return x + y}
    // 但是需要赋值给某个变量 然后通过变量进行调用

    // 当然也可以直接调用 func (x, y int) int { return x + y } (3, 4)
    // defer 语句和匿名函数

	/*
	关键字 defer （详见第 6.4 节）经常配合匿名函数使用，它可以用于改变函数的命名返回值。
	匿名函数还可以配合 go 关键字来作为 goroutine 使用（详见第 14 章和第 16.9 节）。
	匿名函数同样被称之为闭包（函数式语言的术语）：它们被允许调用定义在其它环境下的变量。
	闭包可使得某个函数捕捉到一些外部状态
	例如：函数被创建时的状态。
	另一种表示方式为：一个闭包继承了函数所声明时的作用域。
	这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁，详见第 6.9 节中的示例。
	闭包经常被用作包装函数：它们会预先定义好 1 个或多个参数以用于包装，详见下一节中的示例。另一个不错的应用就是使用闭包来完成更加简洁的错误检查（详见第 16.10.2 节）。*/
}