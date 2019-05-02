package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main () {
	for i := 0; i < 5; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}

	fmt.Println("\n")

	for i := 0; i < 6 ; i++ {
		b := rand.Intn(8)
		fmt.Printf("%d / ", b)
	}

	fmt.Println("\n")

	timens := int64(time.Now().Second())
	rand.Seed(timens)
	for i := 0;i < 5; i++ {
		fmt.Printf("%2.2f /", 100 * rand.Float32())
	}

	fmt.Println("\n")

	// [rand.Float32] [rand.Float64] 返回介于[0.0, 1.0)之间的伪随机数
	// [rand.Intn] 返回介于 [0,n) 之间的伪随机数
	// [rand.Seed] 提供伪随机数的生成种子，一般情况下都会使用当前时间的纳秒级数字
}