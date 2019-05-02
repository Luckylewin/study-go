package main

type TZ int
func main() {
	var a, b, TZ = 3, 4
	c := a + b
	println(c)
}

// 类型别名得到的新类型并非与原类型完全相同，新类型不会拥有原类型所附带的方法