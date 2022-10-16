package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

func main() {
	//输入字符串测试开始.
	input := "abcdefghijklmnopqrstuvwxyz"
	//MD5算法.
	hash := md5.New()
	_, err := hash.Write([]byte(input))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result := hash.Sum(nil)
	//或者result := hash.Sum([]byte(""))
	fmt.Printf("%x\n", result)

	//SHA1算法.
	hash = sha1.New()
	_, err = hash.Write([]byte(input))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result = hash.Sum(nil)
	//或者result = hash.Sum([]byte(""))
	fmt.Printf("%x\n", result)

	//SHA256算法.
	hash = sha256.New()
	_, err = hash.Write([]byte(input))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result = hash.Sum(nil)
	//或者result = hash.Sum([]byte(""))
	fmt.Printf("%x\n", result)

	//SHA512算法.
	hash = sha512.New()
	_, err = hash.Write([]byte(input))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result = hash.Sum(nil)
	//或者result = hash.Sum([]byte(""))
	fmt.Printf("%x\n\n", result)

	//输入字符串测试结束.

	//输入文件测试开始.

	input = "d:/Python.rar"
	filedata, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	//MD5算法.
	hash = md5.New()
	_, err = io.Copy(hash, filedata)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result = hash.Sum(nil)
	//或者result = hash.Sum([]byte(""))
	fmt.Printf("%x %s\n", result, input)

	//SHA1算法.
	hash = sha1.New()
	_, err = io.Copy(hash, filedata)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result = hash.Sum(nil)
	//或者result = hash.Sum([]byte(""))
	fmt.Printf("%x %s\n", result, input)

	//SHA256算法.
	hash = sha256.New()
	_, err = io.Copy(hash, filedata)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result = hash.Sum(nil)
	//或者result = hash.Sum([]byte(""))
	fmt.Printf("%x %s\n", result, input)

	//SHA512算法.
	hash = sha512.New()
	_, err = io.Copy(hash, filedata)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result = hash.Sum(nil)
	//或者result = hash.Sum([]byte(""))
	fmt.Printf("%x %s\n", result, input)

	//输入文件测试结束.

	//程序正常退出.
	os.Exit(0)
}
