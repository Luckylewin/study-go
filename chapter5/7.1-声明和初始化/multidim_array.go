package main

import "fmt"

const (
   WIDTH = 1920
   HEIGHT = 1080
)

type pixel int
var screen [WIDTH][HEIGHT]pixel

func main() {
	// [ 多维数组 ]
	// 数组通常是一维的，但是可以组装成多维数组，例如: [3][5]int, [2][2][2]float64
	// 内部数组总是长度相同的。Go 语言的多维数组是矩形式的（唯一的例外是切片的数组，参见第 7.2.5 节）
	for y:=0; y<HEIGHT;y++ {
		for x :=0; x<WIDTH; x++ {
			screen[x][y] = 0
		}
	}
}	