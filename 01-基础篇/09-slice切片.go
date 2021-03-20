package main

import "fmt"

/**
切片定义方式：var sliceName []sliceType 与数组相仿，[]中不用定义个数,切片操作遵循顾头不顾尾原则
`len` 获取`slice`的长度
`cap` 获取`slice`的最大容量
`append` 向`slice`里面追加一个或者多个元素，然后返回一个和`slice`一样类型的`slice`
`copy` 函数`copy`从源`slice`的`src`中复制元素到目标`dst`，并且返回复制的元素的个数

*/

var s []int
var a1 = [...]int{1,2,3,4,5,6,7,8,9,0}

func main()  {
	s = a1[:]
	fmt.Println(s)
	s = a1[0:len(a1)]
	fmt.Println(s)
	s = a1[:5]
	fmt.Println(s)
	s = a1[3:]
	fmt.Println(s)
	fmt.Println(cap(a1),cap(s),len(s))

	// len cap
	var s2 []int
	s2 = s[1:4]
	fmt.Println(len(s2),cap(s2))	// 切片得来的slice其len为实际key值数量，cap为赋值slice的len值

	// append必须是在已定义的slice下进行
	//var v []int
	s= append(s,1,2,3)
	fmt.Println(s)

	// copy
	var t  = []string{"a","b","c","d"}
	s3:=t[1:4]
	fmt.Println(s3)
	n:=copy(s3,t[1:4])
	fmt.Println(n)
}