package demo

import (
	"AB_Train/utils"
	"fmt"
	"strconv"
	"time"
)

// 基于Go语言实现好用的HTTP接口请求requests
// https://www.cnblogs.com/superhin/p/16332720.html
func GetBaasChains() string {
	time := strconv.FormatInt(time.Now().Unix(), 10)
	//passwd := strings.ToUpper(utils.GetSHA256HashCode(time + utils.GetConfIns().BaasSalt))
	r := utils.Request{
		Method: "GET",
		Url:    "http://localhost:8080",
		Params: map[string]string{
			//"passwd": passwd,
			"time": time},
		//Headers: map[string]string{"Cookie": "abc", "Token": "123"},
	}
	resp := r.Send()
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("原因: %s\n", resp.Reason)
	fmt.Printf("响应时间: %f秒\n", resp.Elapsed)
	//fmt.Printf("响应文本: %s\n", resp.Text)

	return resp.Text
	//JsonParseToStruct(resp.Text)
}
