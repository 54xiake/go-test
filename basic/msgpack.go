package main

import (
	"bytes"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"strings"
	"unicode"
)

func main() {
	buf := &bytes.Buffer{}
	buf.Write([]byte{164, 97, 98, 99, 100})
	buf.Write([]byte{164, 97, 98, 99, 100})
	dec := msgpack.NewDecoder(buf)
	for {
		out, err := dec.DecodeBytes()
		if err != nil {
			break
		}
		fmt.Printf("%v %#v\n", err, string(out))
	}

	fmt.Println(strings.ToUpper("hello world"))
	fmt.Println(strings.ToLower("HELLO WORLD"))
	fmt.Println(strFirstToUpper("hello_world"))

	if IsStartUpper("Hello") {
		fmt.Println("首字母大写")
	} else {
		fmt.Println("首字母小写")
	}
	fmt.Println(Capitalize("hello world"))

	//b := []byte{232, 191, 153, 230, 152, 175, 228, 184, 128, 228, 184, 170, 229, 173, 151, 231, 172, 166, 228, 184, 178}
	b := []byte{144}
	s := string(b)
	fmt.Println(s)

}

/** * 字符串首字母转化为大写 ios_bbbbbbbb -> iosBbbbbbbbb */
func strFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		if y != 0 {
			for i := 0; i < len(vv); i++ {
				if i == 0 {
					vv[i] -= 32
					upperStr += string(vv[i]) // + string(vv[i+1])
				} else {
					upperStr += string(vv[i])
				}
			}
		}
	}
	return temp[0] + upperStr
}

func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func IsStartUpper(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}
