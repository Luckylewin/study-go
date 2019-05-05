package main

func main() {
	println(f())
}

// 请学习以下示例并思考（return_defer.go）：函数 f 返回时，变量 ret 的值是什么？
// 变量 ret 的值为 2，因为 ret++ 是在执行 return 1 语句后发生的。
func f() (ret int) {
	defer func() {
		ret++
		
	}()
	return 1
}