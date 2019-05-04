package main

// for switch select 语句都可以配合标签形式的标识符使用
// 即某一行第一个以冒号:结尾的单词(gofmt 会将后续代码自动移到下一行)
// 标签建议全部使用大写字母

func main() {
	i := 0
	HERE:
		print(i)
		i++
		if i == 5 {
			return 
		}
		goto HERE
}

// 使用逆向的goto 很快会导致意大利面条式的代码
// 一般情况下不建议使用 goto
// 仅仅在特殊的适合的场景中使用

// 注意 标签与goto语句之间不能出现定义新变量的语句
// 否则编译将会失败 
