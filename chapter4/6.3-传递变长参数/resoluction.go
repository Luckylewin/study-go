package main

// 如果变长参数的类型并不都是相同的呢
// 使用五个参数来进行传递并不是明智的选择
// 有两种方案可以解决这个问题

// [1.使用结构]
type Options struct {
	par1 int,
	par2 string
}

func main() {
	// 调用时使用正常的参数 a,b 以及没有初始化的Options 结构
	F1(1,2,Options {})
	// 如果需要对选项进行初始化,则可以使用 
	F1(1,2,Options {par1:10,par2:"Helloworld"})
}

func F1(a int,b int, option Options) {

}

// [方法二:使用空接口]
// 如果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}
// 这样就可以接受任何类型的参数
// 该方案不仅可用于长度未知的参数，还可以用于任何不确定类型的参数
// 一般而言我们会使用一个 for-range 循环 以及 switch 结构对每个参数的类型进行判断

func typecheck(...,...,values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
			case int: ...
			case float: ...
			case string: ...
			case bool: ...
			default: ...
		}
	}
}