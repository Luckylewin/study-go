package main

func main() {
	// Go 语言函数经常使用两个返回值来表示执行是否成功，
	// 返回某个值以及true表示成功，返回零或者nil或者false 表示失败
	// 当不使用 true false 的时候，可以使用一个 error 类型的变量作为第二个值的返回值
	// 成功执行的话，error的值为nil
	// 否则包含相应的错误信息
	// 这样一来就需要用一个 if 语句来测试执行结果
	// 由于符号原因，这样的形式又称之为 comma,ok 模式

	// 在4.7节中 string_conversion.go 中
	// 函数 strconv.Atoi 作用是将一个字符串转换为一个整数，之前我们忽略了相关错误检查
	// anInt, _ = strconv.Atoi(originStr)
	// 如果 originStr 不能转换为整数，anInt 的值将会标为 0 而 _ 无视了错误，程序会继续运行
	// 这样的做法是非常不好的，程序应该在最接近的位置检查所有相关的错误
	// 至少要提示用户有错误发生，并对函数进行返回，或者中断程序.

	// 我们将会改进新的版本 见该文件夹中的string_conversion2.go

	// 习惯用法 [打印错误，返回err]
	value, err := pack1.Function1(param1)
	if err != nil {
		fmt.Printf("An error occured in pack1.Function1 with paramter %v", param1)
		return err
	}

	// 习惯用法 [打印错误，终止程序运行] 使用 os 包的 Exit 函数
	if err != nil {
		fmt.Printf("Program stopping with error %v",err)
		os.Exit(1)
	}

	// 当没有错误发生时，代码继续运行就是唯一要做的事情，所以if 语句不需要 else 分支
	// 示例: 我们尝试通过 os.Open 方法打开一个名为 name 的只读文件:
	f, err := os.Open(name)
	if err != nil {
		return err
	}

	doSomething(f)
	doSomething()

	// 习惯用法 可以将错误的获取位置放在 if 语句的初始化部分
	if err := file.Chmod(0664); err != nil {
		fmt.Println(err)
		return err
	}

	// 习惯用法 可以将 comma-ok pattern 的获取位置放在 if 的初始化部分，然后进行判断
	if value, ok := readData(); ok {
		// do something
	} 



}

// [注意事项1]

// 如果像下面一样，没有为多返回值函数准备足够多的变量来存放结果
func mySqrt(f float64) (v float64, ok bool) {
		if f < 0 { return } // error case
		return math.mySqrt(f), true
}
func main() {
	// 得到编译错误 multiple-value mySqrt() in single-value context	
	t := mySqrt(25.0) 
	fmt.Println(t) 

	// 正确的写法是
	t, ok := mySqrt(25.0)
}

// [注意事项2]
// 当你把字符串转换为整数时，且确定转换一定能够成功时
// 可以将 Atoi 函数进行一层忽略错误的封装
func atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return n
}

// 实际上 fmt 包里最简单的打印函数也有两个返回值
count, err := fmt.Println(x) // number of bytes printed,nil or 0 ,error