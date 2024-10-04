package main

import "fmt"

// 参考：https://www.jianshu.com/p/136f959628d5
func main() {

}

// 指针及地址的使用方法
// 1、当使用 & 操作符对普通变量进行取地址操作时，可以得到变量的指针。
// 此时可以对指针使用 * 操作符，可以得到变量值（此操作也叫指针取值），
// 如以下代码：
func pointerAndAddress() {
	// 定义一个字符串类型的变量
	var myAddr = "tree road 1025, 100"

	// 对字符串取地址， ptr类型为*string
	ptr := &myAddr

	// 打印ptr的类型
	fmt.Printf("ptr的类型是 : %T\n", ptr)

	// 打印ptr的指针地址
	fmt.Printf("ptr的地址是 : %p\n", ptr)

	// 对指针进行取值操作
	value := *ptr

	// 打印取值后的类型
	fmt.Printf("value的类型是 : %T\n", value)

	// 指针取值后就是指向变量的值
	fmt.Printf("value的值是 : %s\n", value)
}

// /**
// * 定义一个交换函数，参数为a、b，类型都为 *int 指针类型
// */
func swap(a, b *int) {

	// 取a指针的值，并把值赋给临时变量t，t此时是int类型
	t := *a

	// 取b指针的值，赋给a指针指向的变量，此时 *a的意思不是取a指针的值，而是 a指向的变量
	*a = *b

	// 将t的值赋给指针b指向的变量
	*b = t
}

// 定义一个人的结构体
type Person struct {
	Age     int    // 年龄
	Sex     int    // 性别
	Name    string // 名字
	Address string // 住址
}

// 新建一个人
func NewPerson(age int, sex int, name string, address string) *Person {

	var person Person

	person = Person{
		Age:     age,
		Sex:     sex,
		Name:    name,
		Address: address,
	}

	// 返回地址即可（因为指针指向的就是地址）
	return &person
}
