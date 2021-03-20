package main

import "fmt"

/**
变量分为：
变量：
a.布尔型			bool
1.长度：			1字节
2.取值范围：		true,false
3.注意：			不能使用数字代表true or false

b.整型			int(正负数)/uint(正数)
1.取决于运行平台（32/64位）
2.int8/uint8别名为byte	int16/uint16	int32别名为rune/uint32	int64/uint64
3.取值范围	-2^64/2 ~ 2^64/2-1 / 0~2^64-1

c.浮点型			float

d.复数型			complex

e.字符串型		string

f.字节型			byte

常量：
不可变的值，定义格式：	const Pi = 3.1415926
支持iota枚举（逢iota由0计算，依次加1）

*/

/**
全局变量声明方式[]为可选
a.未赋值的变量必须声明变量类型,且赋值需在函数体内进行，否则报错syntax error: non-declaration statement outside function body
var variableName variableType
variableName = value

b.直接赋值时可省略变量类型
var variableName [variableType] = value

c.零值
bool	false
string	""
int		0

*/

// 全局下定义变量类型
// 布尔
var bl0 bool
var bl1 bool = true
var bl2 = true
var bl3 bool

// 字符串
var s0 string
var s1 string = "string1"
var s2 = "string2"
var s3 string

//数字型
var n0 int
var n1 int = 1
var n2 = 2
var n3 int

// 常量
const Pi = 3.1415926
const a1 = iota

func main() {
	bl3 = true
	s3 = "string3"
	n3 = 3
	const (
		b1 = iota
		c1
		d1
		e1, f1, g1 = iota, iota, iota
	)
	const h1 = iota

	fmt.Println(bl0, bl1, bl2, bl3)
	fmt.Println(s0, s1, s2, s3)
	fmt.Println(n0, n1, n2, n3)
	fmt.Println(a1, b1, c1, d1, e1, f1, g1, h1)

	/**
	  局部变量声明方式除了可以像全局变量那样定义还可以使用:=代替
	  func main{
	  	s1 := 1
	  	s2 int := 2
	  }
	*/
	a2 := 1
	b2 := false
	c2 := "掌握了吗？"
	fmt.Println(a2, b2, c2)

	// 多重赋值
	a3, b3, c3 := 1, true, "nihao"
	fmt.Println(a3, b3, c3)
	// 匿名变量
	d3, _, f3 := 1, false, "nihao" // go中定义了变量未使用则会报错 	declared and not used
	fmt.Println(d3, f3)
	// 变量间赋值
	g3 := "nihao"
	h3 := g3
	fmt.Println(g3, h3)
	// 变量相加，只有同等类型才可以进行此操作否则报错	k + l (mismatched types string and int)
	i3 := 1
	i3 += 10
	var j3 int
	j3 += 5
	k3 := "nihao"
	l3 := "世界"
	m3 := k3 + l3
	fmt.Println(i3, j3, m3)
}
