package main

import "fmt"

func main() {
	// 严格来说 字符只是整数的特殊用例，byte 类型是 unit8 的别名
	var ch0 byte = 'A'      // 字符A
	var ch1 byte = 65		// ASCII A 对应的值
	var ch2 byte = '\x41'	// 十六进制

	fmt.Printf("%c - %c - %c\n", ch0, ch1, ch2) // character

	// [Unicode] Go支持 Unicode，字符称为 Unicode 代码点 或者 runes
	// 一般用格式 U+hhhh 来表示，其中 h 表示一个十六进制数
	// 在书写Unicode 字符时，需要在16进制数之前加上前缀 \u 或者 \U
	// Unicode 至少占用2个字节，使用 int16 或者 int类型来表示
	// \u 紧跟着长度为4的十六进制数
	// \U 紧跟着长度为8的十六进制数

	var uni1 int = '\u0041'
	var uni2 int = '\u03B2'
	var uni3 int = '\U00101234'

	// 整数
	fmt.Printf("%d - %d - %d\n", uni1, uni2, uni3)
	// 字符
	fmt.Printf("%c - %c - %c \n", uni1, uni2, uni3)
	// UTF-8 字节
	fmt.Printf("%X- %X - %X\n", uni1, uni2, uni3)
	// UTF-8 代码点
	fmt.Printf("%U - %U - %U\n", uni1, uni2, uni3)

	// [格式化说明]
	// %c 用于表示字符; 
	// 当和字符配合使用时，%v 或 %d 会输出用于表示该字符的整数
	// %U 输出 U+hhhh 的字符串

	// [包 unicode] 包含一些针对测试字符非常有用的函数
	// 判断是否为字母:unicode.IsLetter(ch)
	// 判断是否为数字:unicode.IsDigit(ch)
	// 判断是否为空白字符: unicode.isSpace(ch)
}