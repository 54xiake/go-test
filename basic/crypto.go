package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

func main() {
	// 生成 RSA 密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("生成 RSA 密钥对失败: %v", err)
	}
	// 将私钥转换为 ASN.1 PKCS#1 DER 编码
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)
	// 将 DER 编码的私钥转换为 PEM 格式
	privPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privDER,
	})
	// 将公钥提取为 *rsa.PublicKey 类型
	publicKey := &privateKey.PublicKey
	// 将公钥转换为 ASN.1 PKIX DER 编码
	pubDER, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatalf("公钥编码失败: %v", err)
	}
	// 将 DER 编码的公钥转换为 PEM 格式
	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubDER,
	})
	// 打印私钥和公钥
	fmt.Println("私钥:")
	fmt.Println(string(privPEM))
	fmt.Println("公钥:")
	fmt.Println(string(pubPEM))
}
