package main

import "fmt"

/*
当出现多个if-else时可考虑用其代替
	!思考：if和switch的区别？
			if 可以嵌套 可以判断区间  执行效率比较低
			switch 执行效率高  不能嵌套和区间判断
	特性：
		1.switch 选择项可以是一个整型变量
		2.swich中的值不能是浮点型数据 浮点型数据是一个约等于的数据




 */

func main()  {
	// 判断今天是星期几
	var a int
	INPUT:
		fmt.Println("请输入1～7数字")
		fmt.Scan(&a)
	switch a {
	case 1:
		fmt.Println("今天是周一")
	case 2:
		fmt.Println("今天是周二")
	case 3:
		fmt.Println("今天是周三")
	case 4:
		fmt.Println("今天是周四")
	case 5:
		fmt.Println("今天是周五")
	case 6:
		fmt.Println("今天是周六")
	case 7:
		fmt.Println("今天是周日")
	default:
		goto INPUT
	}

	// 判断成绩是否合格
	var results int
	fmt.Println("考了多少分")
	fmt.Scan(&results)
	switch results >=60{
	case true:
		fmt.Println("及格")
	case false:
		fmt.Println("不及格")

	i := 5

	switch i {											// 需要注意的是i和下面case N中的N类型一定是一样的
	case 1:													// 当i未声明时代表true   ex: switch {}
		fmt.Println("i 等于 1")				// 由于默认每个case的最后都带break所以匹配到结果既跳出判断
		fallthrough
	case 2,3,4:											// 如果需要匹配成功后继续向下执行，则通过fallthrough控制
		fmt.Println("i 为 2,3,4 中的一个")
		fallthrough
	case 5:
		fmt.Println("i 等于 5")
		fallthrough
	default:
		fmt.Println("i 比 5 大")
		}
	}
}