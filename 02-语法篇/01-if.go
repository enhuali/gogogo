package main

import "fmt"

/**
if ex1;ex2;ex3 {	// 最多可定义三个参数用分号隔开，第一个判断前执行；第二个是条件；第三个是判断后执行
	cmd1
} else if ex1;ex2;ex3 {			//	注意else 一定要跟在}之后，否则报错syntax error: unexpected else, expecting }
	cmd2
	} else {
		cmd3
	}
判断遵循就近原则，从上至下执行，匹配到即退出判断
 */

func main()  {
	var a,b int
	fmt.Scan(&a,&b)
	if a>b{
		fmt.Printf("a的值为%d大于b的值%d\n",a,b)
	} else if b > a{
		fmt.Printf("a的值为%d小于b的值%d\n",a,b)
		} else {
		fmt.Printf("a的值为%d等于b的值%d\n",a,b)
	}
}