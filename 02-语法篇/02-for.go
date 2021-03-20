package main

import "fmt"

/**
for ex1;ex2;ex3 {
	cmd1		//...同if一样，前/判断/后
}		// 没有定义expression1、expression3时，则与while功能相仿

break				// 跳出此循环，执行其他命令			可配合标签使用
continue		// 跳过此次循环，继续执行此循环

for k,v:=range map {											// for range配合可以遍历列表和字典的index与value
    fmt.Println("map's key:",k)
    fmt.Println("map's val:",v)
}

 */

func main()  {
	for a:=0;a<10;a++{
		fmt.Println(a)
		//break
		continue
	}
	a:=3
	for a==4{
		fmt.Println("没错我是4")
	}
		fmt.Println("我不等于4")

	// 计算1到100的和
	var sum = 0
	for i:=0;i<=100;i++ {
		sum+=i
	}
	fmt.Println(sum)
}