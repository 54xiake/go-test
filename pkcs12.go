package main

import (
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"os"
)

func main() {
	//content := "abc"
	//// read .p12 file
	//buf, err := ioutil.ReadFile("files/2000804843@38.pfx")
	//if err != nil {
	//	//
	//}
	//// extract key and cert
	//pk, cert, err := pkcs12.Decode(buf, "1")
	//fmt.Println(pk)
	//fmt.Println(cert)
	//privateKey := pk.(*rsa.PrivateKey)
	//// create hash
	//h := crypto.SHA256.New()
	//_, err = h.Write([]byte(content))
	//if err != nil {
	//	//
	//}
	//hashed := h.Sum(nil)
	//// how to pass intermediate cert??
	//sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256,
	//	hashed)
	//fmt.Println(sign)
	//sig := base64.RawURLEncoding.EncodeToString(sign)
	//fmt.Println(sig)

	getPrivateKey("files/2000804843@38.pfx", "1")

}

func getPrivateKey(privateKeyName, privatePassword string) (*rsa.PrivateKey, error) {
	f, err := os.Open(privateKeyName)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	// 因为pfx证书公钥和密钥是成对的，所以要先转成pem.Block
	blocks, err := pkcs12.ToPEM(bytes, privatePassword)
	if err != nil {
		return nil, err
	}
	if len(blocks) != 2 {
		return nil, errors.New("解密错误")
	}
	// 拿到第一个block，用x509解析出私钥（当然公钥也是可以的）
	privateKey, err := x509.ParsePKCS1PrivateKey(blocks[0].Bytes)
	if err != nil {
		return nil, err
	}
	fmt.Println(privateKey)
	return privateKey, nil
}
