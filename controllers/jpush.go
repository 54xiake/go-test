package test

import (
	"fmt"
	"github.com/54xiake/go-test/models"
)

type JpushController struct {
	BaseController
}

func (c *JpushController) Send() {
	fmt.Println("jpush start......")
	ids := []string{"100d855909f53538ebf"}
	push := models.Push(ids, "测试内容", "测试title", "aaa", "2f081f43b71d224aa205f6.jpg")
	fmt.Println(push)
	fmt.Println("jpush end......")

	c.SuccessJson(nil)
}
