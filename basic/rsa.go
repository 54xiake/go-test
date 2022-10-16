package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
)

func main() {
	str := "ApplicationID=77ea0eef-ba30-4bb6-8411-dc5c96c5caf5&RandomNumber=136063&SDKType=api&"
	base64Sig, _ := RSASign([]byte(str), "./files/x.pfx")
	fmt.Println("签名后信息", base64Sig)

	//err := RSAVerify([]byte(str1), base64Sig,"./files/publickey.cer")
	err := RSAVerify([]byte(str), base64Sig, "./files/publickey.cer")
	//err := RSAVerify([]byte(str), base64Sig,"./files/1_pub.cer")
	if err == nil {
		fmt.Println("验证签名ok!")
	} else {
		fmt.Println("验证失败！")
	}
}

func RSASign(data []byte, filename string) (string, error) {
	//1.选择hash算法，对需要签名的数据进行hash运算
	myhash := crypto.SHA256
	//myhash := crypto.MD5

	hashInstance := myhash.New()
	hashInstance.Write(data)
	hashed := hashInstance.Sum(nil)
	//2.读取私钥文件，解析出私钥对象
	privateKey, err := ReadParsePrivateKey(filename)
	if err != nil {
		return "", err
	}
	//3.RSA数字签名
	bytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, myhash, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

//公钥验证数据签名是否正确
func RSAVerify(data []byte, base64Sig, filename string) error {
	//- 对base64编码的签名内容进行解码，返回签名字节
	bytes, err := base64.StdEncoding.DecodeString((base64Sig))
	if err != nil {
		return err
	}
	//- 选择hash算法，对需要签名的数据进行hash运算
	myhash := crypto.SHA256
	//myhash := crypto.MD5
	hashInstance := myhash.New()
	hashInstance.Write(data)
	hashed := hashInstance.Sum(nil)
	//- 读取公钥文件，解析出公钥对象
	publicKey, err := ReadParsePublicKey(filename)
	fmt.Println(publicKey)

	if err != nil {
		return err
	}
	//- RSA验证数字签名
	return rsa.VerifyPKCS1v15(publicKey, myhash, hashed, bytes)
}

//读取公钥文件，解析出公钥对象
func ReadParsePublicKey(filename string) (*rsa.PublicKey, error) {
	//--1.读取公钥文件，获取公钥字节
	publicKeyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	//--2.解码公钥字节，生成加密对象
	block, _ := pem.Decode(publicKeyBytes)
	//fmt.Println(block)
	if block == nil {
		return nil, errors.New("公钥信息错误")
	}
	//--3.解析DER编码的公钥,生成公钥接口
	//publicKeyInterface,err :=x509.ParsePKIXPublicKey(block.Bytes)
	//fmt.Println(publicKeyInterface)
	publicKeyInterface, err := x509.ParseCertificate(block.Bytes)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//--4.公钥接口转型成公钥对象
	publicKey := publicKeyInterface.PublicKey.(*rsa.PublicKey)
	//publicKey := publicKeyInterface.(*rsa.PublicKey)
	//CN=051@kaifazhe@Z798648334@76,OU=Local RA+OU=Organizational-1,O=CFCA TEST CA,C=CN
	fmt.Println(publicKeyInterface.Subject)
	return publicKey, nil
}

//读取私钥文件，解析出私钥对象
func ReadParsePrivateKey(filename string) (*rsa.PrivateKey, error) {
	//--1.读取私钥文件，获取私钥字节
	privateKeyBytes, _ := ioutil.ReadFile(filename)
	//fmt.Println(privateKeyBytes)
	//--2.对私钥文件进行编码,生成加密对象
	//block,err1 := pem.Decode(privateKeyBytes)
	// 因为pfx证书公钥和密钥是成对的，所以要先转成pem.Block
	block, err1 := pkcs12.ToPEM(privateKeyBytes, "1")
	if block == nil {
		fmt.Println(err1)
		return nil, errors.New("私钥信息错误")
	}
	if len(block) != 2 {
		return nil, errors.New("解密错误")
	}
	//3.解析DER编码的私钥,生成私钥对象
	privateKey, err := x509.ParsePKCS1PrivateKey(block[0].Bytes)
	publicKeyInterface, err := x509.ParseCertificate(block[1].Bytes)
	fmt.Println(publicKeyInterface.Subject)
	rsaPublicKey := publicKeyInterface.PublicKey.(*rsa.PublicKey)
	fmt.Println(rsaPublicKey.N)
	fmt.Println(rsaPublicKey.E)
	//fmt.Println(x509.MarshalPKIXPublicKey(publicKeyInterface.PublicKey.(*rsa.PublicKey)))

	derPkix, err := x509.MarshalPKIXPublicKey(rsaPublicKey)
	if err != nil {
		return nil, errors.New("解密错误")
	}
	block1 := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	bufferPublic := new(bytes.Buffer)
	err = pem.Encode(bufferPublic, block1)
	if err != nil {
		return nil, errors.New("解密错误")
	}
	publicKeyStr := bufferPublic.String()
	fmt.Println(publicKeyStr)

	if err != nil {
		return nil, err
	}
	return privateKey, err
}
