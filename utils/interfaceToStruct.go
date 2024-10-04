package main

import (
	"fmt"
)

// 定义一个结构体
type MyStruct struct {
	Field1 string
	Field2 int
}

func main() {
	// 创建一个空接口类型的变量，并赋值为MyStruct类型
	var data interface{} = MyStruct{"Hello", 123}

	// 将interface{}类型的变量转换为MyStruct类型
	if myData, ok := data.(MyStruct); ok {
		fmt.Printf("Field1: %s, Field2: %d\n", myData.Field1, myData.Field2)
	} else {
		fmt.Println("data is not of type MyStruct")
	}
}
