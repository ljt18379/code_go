package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//msg := []byte("Hello world. 你好，世界！")
	msg := []byte("nDQW6XZ7b_u2Sy9slofYLlG03sOEoug3I0aAPQ0exs4")
	// 标准编码
	encoded := base64.StdEncoding.EncodeToString(msg)
	fmt.Println(encoded)
	// SGVsbG8gd29ybGQuIOS9oOWlve+8jOS4lueVjO+8gQ==

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))
	// Hello world. 你好，世界！

	// 常规编码，末尾不补 =
	encoded = base64.RawStdEncoding.EncodeToString(msg)
	fmt.Println(encoded)
	// SGVsbG8gd29ybGQuIOS9oOWlve+8jOS4lueVjO+8gQ

	decoded, _ = base64.RawStdEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))
	// Hello world. 你好，世界！

	// URL safe 编码
	encoded = base64.URLEncoding.EncodeToString(msg)
	fmt.Println(encoded)
	// SGVsbG8gd29ybGQuIOS9oOWlve-8jOS4lueVjO-8gQ==

	decoded, _ = base64.URLEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))
	// Hello world. 你好，世界！

	// URL safe 编码，末尾不补 =
	encoded = base64.RawURLEncoding.EncodeToString(msg)
	fmt.Println(encoded)
	// SGVsbG8gd29ybGQuIOS9oOWlve-8jOS4lueVjO-8gQ

	decoded, _ = base64.RawURLEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))
	// Hello world. 你好，世界！
}
