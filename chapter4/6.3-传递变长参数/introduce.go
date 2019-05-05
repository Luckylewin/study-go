// 如果函数的最后一个参数时采用 ...type 的形式,
// 那么这个函数就可以处理一个变长的参数
// 长度可以为0，这样的函数称之为变参函数
func myFunc(a, b, arg ...int) {}

// 这个函数接受一个类似某个类型的slice 的参数
// 该参数可以使用 for 进行迭代

func Greeting(prefix string, who ..string)
Greeting("Hi", "Joe", "Anna", "Elieen")

