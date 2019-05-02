package main

func main() {
	// 字符串是 UTF-8 字符的一个序列
	// 当字符串ASCII码时则占用一个字节，其它字符根据需要占用2-4字节

	// [解释字符串]
	// \n : 换行符
	// \r : 回车符
	// \t : tab键盘
	// \u|\U: Unicode 字符
	// \\ 反斜杠自身表示

	// [非解释字符串]
	// 使用反引号括起来 
	// `This is a raw string \n` 其中 `\n` 会被原样输出

	// [纯字节的字符串内容](只对ASCII码的字符串有效) 
	// 可通过标准索引法来获取，在中括号[]内写入索引，索引从0开始计数
	// 字符串的第一个字节 str[0]
	// 字符串的第i个字节  str[i - 1]
	// 最后一个字节       str[len(str) - 1]

	// [字符串拼接 +]
	// 两个字符串 s1 和 s2 可以通过 s := s1 + s2 拼接在一起

	str := "hello world " + "by golang "
	str += "-2019-05-02"

	println(str)

	// 在循环中使用 + 拼接字符串并不是最搞笑的做法，更好的办法是使用函数 strings.Join()
	// 最好的办法是使用字节缓冲[bytes.Buffer] 拼接更加给力
}