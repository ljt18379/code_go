package main

import (
	"fmt"
)

// 截取字符串指定位置的子串
func Substr(s string, start, length int) string {
	// 将字符串转换为rune切片处理Unicode字符
	r := []rune(s)
	if start < 0 || start >= len(r) {
		return ""
	}
	if length < 0 {
		return ""
	}
	if start+length > len(r) {
		length = len(r) - start
	}
	return string(r[start : start+length])
}

func main() {
	str := "{\"Key\":\"did:cmid:001:CM001:bfa970da-6592-4380-b830-a9b3181d78a5\",\"Record\":{\"id\":\"did:cmid:001:CM001:bfa970da-6592-4380-b830-a9b3181d78a5\",\"createTime\":\"2024-05-20T16:02:20Z\",\"updateTime\":\"2024-06-12T16:02:20Z\",\"activated\":\"\",\"version\":\"1\",\"count\":\"\",\"docHash\":\"b97c144c0a27a0ee88e63c2af5bdd8ff\",\"url\":\"\"}}"
	pos := 7
	n := 5
	result := Substr(str, pos, n)
	fmt.Println(result) // 输出: World
}
