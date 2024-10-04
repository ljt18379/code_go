package main

import (
	"encoding/json"
	"fmt"
)

//Go的json序列化：Marshal与Unmarshal
type Stu struct {
	Name  string `json:"name"`
	Age   int    //来源公众号：「码农编程进阶笔记」
	sex   string //小写的字段不可导出，所以无法序列化
	Class *Class `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}
func marshal(stu Stu){
	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err.Error())
	}
	// jsonStu是[]byte类型，转化为string类型便于查看
	fmt.Println(string(jsonStu))
}

func unmarshal(stu Stu){


	//json字符串沿用上面的结果，但对key进行了大小的修改，并添加了sex数据
	data:="{\"name\":\"张三\",\"Age\":18,\"high\":true,\"sex\":\"男\",\"CLASS\":{\"naME\":\"1班\",\"GradE\":3}}"
	str:=[]byte(data)

	//1.Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构。
	//第二个参数必须是指针，否则无法接收解析的数据，如stu仍为空对象StuRead{}
	//2.可以直接stu:=new(StuRead),此时的stu自身就是指针      //来源公众号：「码农编程进阶笔记」    stu:=StuRead{}
	err:=json.Unmarshal(str,&stu)

	//解析失败会报错，如json字符串格式不对，缺"号，缺}等。
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(stu)
}


func main(){
	stu := Stu{
		Name: "张三",
		Age:  18,
		sex:  "男",
	}
	//指针变量
	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class=cla
	marshal(stu)
	unmarshal(stu)
}