package main

// 常量只能是布尔型、数字型和字符串型 必须是编译时就能确定的值
const Pi = 3.14

const flag = true // 隐式类型
const flag2 bool = false //显式类型

// 并行赋值
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6

// 用于枚举
const (
	unknown = 0
	male = 1
	female = 2
)

// 与 iota 结合
const (
	Jan = iota
	Feb
	Mar
)

// 某个类型作为枚举量的类型
type color int

const (
	Red color = iota
	Green
	Blue
	Orange
)

func main() {
	println(Pi)
	println(Orange)
}