package main

import (
	"fmt"
	"strconv"
)

func(){
	
}

// 截取字符串
// 16进制转10进制
func test() {
	str := "100f4ceb0002642f7b89"
	str2 := str[len(str)-8 : len(str)] //截取后八位
	fmt.Println(str2)
	n, _ := strconv.ParseUint(str2, 16, 32) //16进制转10进制
	fmt.Println(n)
}
