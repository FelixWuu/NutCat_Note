package variable

import "fmt"

var a int     // 自动初始化为 0
var b = false // 自动推断为 bool 类型
var c, d int  // 可以同时定义多个相同类型的变量

const c1 = 123
const c2, c3 = "Hello", "world"

// 也可以显式的指定常量的类型
const c4 string = "boom"

// 以组的方式定义
var (
	e, f int
	g, h = 777, "4396"
)

var x = 100

// Case1 使用简短模式应该注意作用域
func Case1() {
	fmt.Println(&x, x)
	res1 := x + 1
	fmt.Printf("res1: %v\n", res1)

	x := 1000
	fmt.Println(&x, x)
	res2 := x + 1
	fmt.Printf("res2: %v\n", res2)
}

// Case2 注意多变量赋值的先后顺序
func Case2() {
	x, y := 1, 2
	x, y = y+100, x+200

	fmt.Println(x, y)
}

const (
	c5 int = 999
	c6
	c7 = "OK"
	c8
)

// Case3 如果不指定类型和初始化的值，常量会与上一个非空常量相同。
func Case3() {
	fmt.Printf("c6 type: %T, value: %v\n", c6, c6)
	fmt.Printf("c8 type: %T, value: %v\n", c8, c8)
}

const (
	e1 = iota // 0
	e2        // 1
	e3        // 2
)

// 常见枚举
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

const (
	t1 = iota // 0
	t2        // 1
	t3 = 100  // 100
	t4        // 100, 因为未指定类型与初始化值，它与上一个非空常量相同
	t5 = iota // 4， 每新增一行，iota都会自增1
	t6        // 5
)

func Case4() {
	var a byte = 0x11
	var b uint8 = a
	var c uint8 = a + b
	test(c)
}

func test(n byte) {
	println(n)
}
