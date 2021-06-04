package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"regexp"
)

func Md5Str(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Struct2Map(obj interface{}) map[string]interface{} {
	var data = make(map[string]interface{})
	d, _ := json.Marshal(obj)
	_ = json.Unmarshal(d, &data)

	return data
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
