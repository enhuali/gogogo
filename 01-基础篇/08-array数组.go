package main

import "fmt"


/**
数组定义格式:var arrayName [n]arraryType ，n为多少时就一定不能大于n个value，否则报错invalid array index 2 (out of bounds for 2-element array)
不确定有多少个key时，可以用...表示
*/
var a =[...]int{1,2,3}
var b [2]int

func main()  {
	b[0] = 0
	b[1]=1
	c:=[3]int{1,2,3}
	//b[2]=2
	fmt.Println(a,b,c)
}