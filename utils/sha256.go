package main

import (
	"crypto/sha256"
	"encoding/hex"
)

//SHA256生成哈希值

func GetSHA256HashCode(stringMessage string) string {
	message := []byte(stringMessage)
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New() //sha-256加密
	//hash := sha512.New() //SHA-512加密
	hash.Write(message)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode

}
func GetSha256Byte(bodyByte []byte) []byte {
	sha := sha256.New()
	hash := sha.Sum(bodyByte)
	return hash
}
