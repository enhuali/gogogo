package main

import "fmt"


//func main()  {
//	var s3,s4 string
//	fmt.Scan(&s3,&s4)
//	fmt.Println(&s3,&s4)
//	fmt.Println(s3,s4)
//
//}

func main()  {
	// 输入值并修改内存地址
	var s3,s4 string
	fmt.Scan(&s3,&s4)
	fmt.Println(&s3,&s4)
	fmt.Println(s3,s4)

	// 根据输入的值计算圆的周长及面积
	var r float64
	const Pi  = 3.1415926
	fmt.Scan(&r)											// 接收的值记录为内存地址
	fmt.Printf("面积为%.2f\n",Pi*r*r)          		// %.2f	保留两位小数
	fmt.Printf("周长为%.2f\n",2*Pi*r)

	// 输入账号密码
	var name string
	var pwd string
	fmt.Println("请您输入用户名")
	fmt.Scanf("%s",&name)
	fmt.Println("请您输入密码")
	fmt.Scanf("%s",&pwd)
	fmt.Printf("账号为：%s\n密码为：%v\n",name,pwd)

	var a int
	var b string
	fmt.Scanf("%3d", &a)
	_, _ = fmt.Scanf("%s", &b)
	fmt.Println(a, b)


}

