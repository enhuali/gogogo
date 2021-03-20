package main

import "fmt"

// 通常传入的是一个变量的copy值，对其进行操作不影响其内存所处位置；如需要修改变量内存所在地址，需通过"*变量名称"变为指针类型

func main()  {
	/**
	通常传入的是一个变量的copy值，对其进行操作不影响其内存所处位置；如需要修改变量内存所在地址，需通过"*变量名称"变为指针类型
	内存地址		&
	指针类型		*	空值为nil	一个指针指向了一个内存地址（类似变量和常量的关系），通过修改指针来改变内存地址
	 */
	var s1 string
	fmt.Println(&s1)
	s1="nihao"
	fmt.Println(&s1)


	// 指针声明格式 var variableName *variableType
	var s2 string = "nihao"
	var s3 *string
	fmt.Println(s3)
	s3 = &s2
	fmt.Println(s3)			// 取出内存地址
	fmt.Println(*s3)		// 取出值

}