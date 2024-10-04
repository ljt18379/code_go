package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func tt() {
	//给打印出来的字符串带上双引号,使用转义字符“\”引用字符串
	fmt.Println("\"GeeksforGeeks is a computer science portal\"")
}

func Type() {
	var v = 0
	fmt.Printf("variable v2=%v is of type %T \n", v, v)
	fmt.Printf("v type:%T\n ", v)

}

func TrimFirstCharPrefixSpace(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	firstCharIndex := -1
	for i, r := range runes {
		if firstCharIndex >= 0 && !strings.HasPrefix(string(rune(r)), " ") {
			return string(runes[:i]) + string(runes[firstCharIndex:])
		}
		if firstCharIndex < 0 && !strings.HasPrefix(string(rune(r)), " ") {
			firstCharIndex = i
		}
	}
	return s
}
func TestTrimSpace(t *testing.T) {
	str := "       abcdefg  "
	s := strings.TrimSpace(str)
	fmt.Println("处理之后的字符串:", s)
	// 处理之后的字符串: abcdefg
}

// compare string
func compareStrings(str1 string, str2 string) bool {
	return strings.Compare(str1, str2) == 0
}

// string-int
func strToInt() {
	hisUpdateTime, err := strconv.Atoi("did.UpdateTime")

	updateTime, err := strconv.ParseInt("did.UpdateTime", 10, 64)
}

// int to string
func intToStr() {
	var updateTime int64
	updateTime = 1717228940
	updateTimeStr := strconv.Itoa(int(updateTime))
	fmt.Println(updateTimeStr)
}

func int64ToStr() {
	//num := int64(1234567890)
	str := strconv.FormatInt(0, 10)
	fmt.Println(str) // 输出: "1234567890"
}

// *[]string 与[]*string
func arrString() {
	var vcList []*string
	for _, v := range vcList {
		fmt.Println("v:", v)
		vv := *v
		fmt.Println("vv:", vv)
	}

	//var vcList2 *[]string

}

func strSplit() {
	ss := "[dagah]"
	ss = strings.TrimPrefix(ss, "[")
	ss = strings.TrimSuffix(ss, "]")
}
func main() {
	str1 := "2024-06-12T16:02:20Z"
	str2 := "2024-06-12T16:02:20Z"
	fmt.Println(compareStrings(str1, str2))
}
