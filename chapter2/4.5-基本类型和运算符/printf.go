package main

import "fmt"

func main() {
	var a float32 = 3.14159
	var b int = 1680325

	// %n.mg 表示数字n并精确到小数点后m位
	fmt.Printf("%2.2g\n",a)
	// %g 格式化浮点数
	fmt.Printf("%g\n", a)
	// %f 输出浮点数
	fmt.Printf("%f\n",a)
	// %e 科学计数法
	fmt.Printf("%e\n",a)
	// %d 格式化整数
	fmt.Printf("%d\n",b)
	// %0d 输出定长的整数 0是必须
	fmt.Printf("%0d\n",b)
}