package main
 
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)
 
func WriteToFile(path string, content string) {
	// 将PEM格式的私钥写入文件
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
 
	_, err = file.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}
 
func main() {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
 
	// 将私钥序列化为PEM格式
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	WriteToFile("private_key.pem", string(privateKeyPem))
 
	//生成公钥
	publicKey := privateKey.PublicKey
	// 将私钥序列化为PEM格式
	publicKeyBytes := x509.MarshalPKCS1PublicKey(&publicKey)
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	WriteToFile("public_key.pem", string(publicKeyPem))
}