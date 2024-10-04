package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int64
	//num = 1234567890
	str := strconv.FormatInt(num, 10)
	fmt.Println(str) // 输出: "1234567890"

	var s *string
	s1 := ""
	s = &s1
	s2 := ""
	s = &s2
	fmt.Println(s)
}
