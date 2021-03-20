package main

import "fmt"


func main()  {
	/**
	  Print		无换行，不可读取变量，需使用\n进行换行
	  Println		有换行，不可读取变量
	  Printf		无换行，可以读取变量，需使用\n进行换行
	*/
	a1 := 1
	fmt.Println("这个数字为：%d",a1)
	fmt.Printf("这个数字为：%d",a1)
	fmt.Printf("这个数字为：%d\n",a1)
	fmt.Print("这个数字为：%d",a1)
	fmt.Print("这个数字为：%d\n",a1)
	fmt.Println(`用这个符号
你怎么输入的
   它怎么输出`)

	/**
	  占位符：
	  string		%s
	  int			%d
	  bool			%t
	  字符			%c
	  内存地址		%p
	  通用			%v		除字符及内存地址外
	*/
	s2:="世界"
	i2:=2019
	var b2 bool
	c2:='a'
	m2:=100
	fmt.Printf("s2:%s,i2:%d,b2:%t,c2:%c,m2:%p\n",s2,i2,b2,c2,&m2)
	fmt.Printf("s2:%v,i2:%v,b2:%v,c2:%v,m2:%v\n",s2,i2,b2,c2)
}
