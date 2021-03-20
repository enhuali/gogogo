package main

import "fmt"

/**
goto 可跳转到对应的标签下。通常标签以大写命名
 */

func main()  {
	Here:
		fmt.Println("你来了，朋友")
	for i:=0;i<10;i++ {
		goto Here
	}
}