package main

import "os"

// 变量声明的一般格式 var identifier type

// 基本形式
var a int
var b bool
var c string

// 因式分解声明变量 多用于声明全局变量
var (
	d int
	e bool
	f string
)

// [默认值]
// 变量声明后 初始值默认为 int:0,bool:false,string:空字符串,指针:nil

// [命名规范]
// 变量的命名规则遵循骆驼命名法，举例子 startDate
// 如果全局变量希望被外部包使用，则首字母也要大写 如MakeSignature

func main() {

// [赋值前提]
// 类型必须相同
	var a int
	var b int = 1
	b = a

// [编译器类型推断]
	var c = 1
	var d = false
	var e = "golang"

// 声明包级别的全局变量
	var (
		HOME = os.Getenv("HOME")
		USER = os.Getenv("USER")
	)	

// 函数内部变量声明
	f := 1
	g := "hello world"

// [注意]
// 相同的代码块中不可以再对 f,g 进行 f := x 或 g := x 变量声明

// 同一类型多个变量声明
	var h, i, j int

// 并行赋值
	k, l, m := 1, 2, 3

	println(a,b,c,d,e,HOME,USER,f,g,h,i,j,k,l,m)
}