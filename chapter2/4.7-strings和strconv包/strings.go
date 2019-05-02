package main

import "strings"

func main() {
	// [包含关系]
	// strings.Contains(s, substr string) bool
	str1 := "Hello world golang"
	println(strings.Contains(str1, "world"))

	// [字符出现的索引]
	// 第一次出现索引
	// strings.Index(s,str string) int -1 表示字符串s不包含str
	// 最后一次出现的索引
	// strings.LastIndex(s, str string) int
	// 查询非ASCII编码字符在父字符串中的索引
	// strings.IndexRune(s string, r rune) int
	// strings.LastIndex()

	// [字符串替换]
	// 将前n个old替换成new,n=-1时是替换所有
	// strings.Replace(str, old, new, n)

	// [统计字符串出现次数]
	// strings.Count(s,string) int

	// [重复字符串] 重复 count 次字符串 s 并返回一个新的字符串
	// strings.Repeat(s, count int) string 

	// [大小写转换]
	// strings.ToLower(s) string
	// strings.ToUpper(s) string

	// [修剪字符串]
	// strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号

	// [剔除指定的字符]
	// strings.Trim(s,"cut") 剔除字符串开头和结尾的cut字符串
	// strings.TrimLeft(s,"cut") 左剔除
	// strings.TrimRight(s,"cut") 右剔除

	// [拼接slice到字符串]
	// Join 用于将字符串类型为 string  的 slice 使用分割符号来拼接成一个字符串
	// strings.Join(sl []string, sep string) string

	// [从字符串中读取内容]
	// 函数 strings.NewReader(str) 用于生成一个 Reader 并读取字符串中的内容,然后返回该Reader的指针

	// Read() 从[]byte 中读取内容
	// ReadByte() 和 ReadRune() 从字符串中读取下一个Byte

	// [字符串与其它类型的转换]
	// 任何类型T转换成字符串总是OK的
	// strconv.Itoa(i int) string 返回数字 i 所表示的字符串类型的十进制
	// strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string
	// fmt 表示格式 其值可以是 'b','e','f','g'
	// prec 表示精度
	// bitSize 使用32表示 float32, 使用64则表示 float64

	// 字符串转换为数字类型，Go 提供了以下函数
	// strcov.Atoi(s string) (i int, err error) 将字符串转换成int型
	// strocv.ParaseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型号
	// 多返回值特性, 这些函数会返回两个值，第一个是转换后的结果，第二个是可能出现的错误
	// var, err = strconv.Atoi(s)
}

