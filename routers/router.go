package routers

import (
	"fmt"
	controllers "github.com/54xiake/go-test/controllers"
	"github.com/astaxie/beego"
)

var (
	IsProd bool
)

func init() {
	fmt.Println("init router=============")
	beego.Router("/user/create", &controllers.UserController{}, "Get:Create")
	beego.Router("/jpush/send", &controllers.JpushController{}, "Get:Send")
	beego.Router("/user/create", &controllers.UserController{}, "Get:Create")
}
