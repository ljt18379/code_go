package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `data:"test" test:"haha"`
	Age  string `json:"test"`
}

// Go语言结构体标签与反射(19) https://www.modb.pro/db/166045
// Binding 组件 - Go语言中文网 - Golang中文社区  https://studygolang.com/resources/4798
func Reflect() {
	user := User{"土味", "测试"}
	userValue := reflect.ValueOf(user) //获取变量的值

	userType := reflect.TypeOf(user) //获取类型
	fmt.Println(userValue, userType)

	//反射的用法，获取结构体标签的值
	nameData := userType.Field(0)           //user第一个字段
	nameTagData := nameData.Tag.Get("data") //获取结构体第一个字段的Tag（标签）里面键为data的值
	fmt.Println("name_data:", nameTagData)
}
