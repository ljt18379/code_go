package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func main() {
	base64Str2publicKey()
}

// pem publicKey -> base64 string
func publicKey2base64Str() {
	// 生成RSA公钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// 编码公钥为PEM格式
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey),
		},
	)

	// 将PEM格式的公钥转换为Base64字符串
	pubBase64 := base64.StdEncoding.EncodeToString(pubPEM)

	fmt.Println(pubBase64)
}

// base64 string -> pem publicKey
func base64Str2publicKey() {
	// 假设base64Str是一个Base64编码的公钥字符串
	base64Str := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0="

	// 将Base64字符串解码成byte slice
	decodedBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		panic(err)
	}

	// 创建PEM块
	pemBlock := &pem.Block{
		Type:  "PUBLIC KEY", // PEM类型，根据实际情况可能是"RSA PUBLIC KEY"等
		Bytes: decodedBytes,
	}

	// 将PEM块转换成PEM格式的字符串
	pemStr := string(pem.EncodeToMemory(pemBlock))
	fmt.Println(pemStr)
}
