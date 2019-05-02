package main

import (
	"fmt"
	"time"
)

// time 包
// 为我们提供了一个数据类型 time.Time (作为值使用)
// 以及 显示和测量时间和日期的功能函数

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t) 
	fmt.Printf("Y-m-d: %04d-%02d-%02d\n",t.Year(), t.Month(), t.Day())

	t = time.Now().UTC()
	fmt.Printf("time.Now().UTC(): %s\n", t)
	fmt.Printf("time.Now()      : %s\n", time.Now())
	// 计算时间
	week = 60 * 60 * 24 * 1e9
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)
	// 格式化输出
	fmt.Printf("time.RFC822 : %s\n",t.Format(time.RFC822))
	fmt.Printf("time.ANSIC : %s\n",t.Format(time.RFC822))
	fmt.Printf("format:02 Jan 2010 15:00: %s\n",t.Format("format:02 Jan 2010 15:00"))

	s := t.Format("20100816")
	fmt.Println(t, "=>", s)
}