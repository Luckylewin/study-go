package main

import (
	"math"
	"fmt"
	"errors"
)

func main() {
	// 编写一个名字为 MySqrt 的函数，
	// 计算一个 float64 类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。
	// 编写两个版本，一个是非命名返回值，一个是命名返回值。
	fmt.Printf("First example with -1:")
	ret1, err1 := MySqrt1(-1)
	if err1 != nil {
		fmt.Println("Error!Return values are:",ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are:",ret1)
	}

	fmt.Print("Second example with 5: ")

	if ret2, err2 := MySqrt2(25.00); err2 != nil {
		fmt.Println("Error!Return values are:",ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are:",ret2)
	}

}

func MySqrt1(f float64) (float64, error) {
	if f < 0 {
		return float64(math.NaN()), errors.New("can't not afford a negative number in sqrt operation")
	}

	return math.Sqrt(f), nil
}

func MySqrt2(f float64) (res float64, err error) {
	if f < 0 {
		res = float64(math.NaN())
		err = errors.New("can't not afford a negative number in sqrt operation")
	} else {
		res = math.Sqrt(f)
	}

	return
}

