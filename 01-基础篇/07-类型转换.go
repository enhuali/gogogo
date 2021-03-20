package main

import (
	"fmt"
	"reflect"
)

func main()  {
	a:=1
	var b string
	fmt.Println(reflect.TypeOf(a))			// 打印数据类型
	fmt.Println(reflect.TypeOf(b))
	c := string(a)							// 类型转换
	fmt.Println(reflect.TypeOf(c))


	//小案例 A
	//编程实现107653秒是几天几小时几分钟几秒?
	time := 107653
	e:= time/60/60/24
	fmt.Println(e)
	fmt.Println("天：", time/60/60/24%365)
	fmt.Println("时：", time/60/60%24)
	fmt.Println("分：", time/60%60)
	fmt.Println("秒：", time%60)

}