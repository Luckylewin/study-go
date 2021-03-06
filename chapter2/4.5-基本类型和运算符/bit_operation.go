package main

// 位左移常见实现存储单位的用例
type ByteSize float64
const (
	_ = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// 在通讯中使用位左移表示标识的用例
type BitFlag int
const (
	Active BitFlag = 1 << iota 
	Send 
	Receive
)

// 位运算
func main() {
	// [二元运算符]

	// [ 按位与 &] 全真才真
	// 1 & 1 -> 1
	// 1 & 0 -> 0
	// 0 & 1 -> 0
	// 0 & 0 -> 0

	// [ 按位或 |] 一真全真
	// 1 & 1 -> 1
	// 1 & 0 -> 1
	// 0 & 1 -> 1
	// 0 & 0 -> 0

	// [ 按位亦或 ^] 不同才真
	// 1 & 1 -> 0
	// 1 & 0 -> 1
	// 0 & 1 -> 1
	// 0 & 0 -> 0

	// [一元运算符]
	// [按位补足 ^]
	// 该运算符与异或运算符一同使用，即 m^x，对于无符号 x 使用“全部位设置为 1”，对于有符号 x 时使用 m=-1。例如：
	// ^10 = -01 ^ 10 = -11

	// [位左移 <<] bitP << n
	// bitP 的位向左移动 n 位，右侧空白部分使用 0 填充；如果 n 等于 2，则结果是 2 的相应倍数，即 2 的 n 次方
	// 1 << 10 // 1KB
	// 1 << 20 // 1MB
	// 1 << 30 // 1GB

	// [位右移 >>] bitP >> n
	// bitP 的位向又移动 n 位，左侧空白部分使用 0 填充；如果 n 等于 2，则结果是当前值除以 2 的 n 次方

	flag := Active | Send
	println(flag)

	// [逻辑运算符]
	// == != > >= < <= 之所以被称为逻辑运算符，是因为它们的结果总为布尔值

	// [算术运算符]
	// 对于整数运算而言，结果依旧为整数，例如 9 / 4 -> 2
	// 取余运算符只能作用于整数 9 % 4 -> 2
}