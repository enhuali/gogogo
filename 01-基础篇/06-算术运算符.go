package main

import "fmt"


// 需要注意的是只有同种类型的变量才可相加否则报错a2 / b2 (mismatched types int and float64)
func main()  {
	// 算数运算符
	// 加减法
	var a1,b1 int=1,10
	c1:=a1+b1
	d1:=a1-b1
	fmt.Println(c1,d1)

	// 乘除法
	a2:=1
	//b2:=1.2		需要注意的是只有同种类型的变量才可相加否则报错a2 / b2 (mismatched types int and float64)
	b2:=10
	c2:=a2/b2
	d2:=a2*b2
	fmt.Println(c2,d2)

	// 取余 取模
	a3:=1
	b3:=10
	c3:=a3%b3
	fmt.Println(c3)

	// 自增减
	i:=1
	i++
	fmt.Println(i)
	i--
	i--
	fmt.Println(i)



	/**
	逻辑运算符：
	非[!] 非真为假，非假为真
	或[||] 同假为假，其余为真
	与[&&] 同真为真，其余为假
	 */
}