package test

import (
	"github.com/54xiake/gotest/utils"
	"github.com/astaxie/beego"
)

// 所有Controller继承BaseController,实现页面布局
type BaseController struct {
	beego.Controller
}

func (base *BaseController) Prepare() {
	// 设置Layouts
	base.Layout = "layouts/layout.html"
	base.LayoutSections = make(map[string]string)
	base.LayoutSections["Navbar"] = "layouts/navbar.tpl"
	base.LayoutSections["Breadcrumb"] = "layouts/breadcrumb.tpl"
	//base.LayoutSections["Scripts"] = "layouts/scripts.tpl"
	base.LayoutSections["Footer"] = "layouts/footer.tpl"
	base.LayoutSections["Sidebar"] = "layouts/sidebar.tpl"
	base.Data["BaseTitle"] = beego.AppConfig.String("appname")
	//base.Data["username"] = base.GetSession("username")

	// controller & action
	base.Data["controller"], base.Data["action"] = base.GetControllerAndAction()
	base.Data["uri"] = base.Ctx.Request.URL.Path
	base.Data["IsAdmin"] = 0
	//if isAdmin := base.GetSession("isadmin"); isAdmin != nil {
	//	base.Data["IsAdmin"] = isAdmin.(int)
	//}

}

func (base *BaseController) SetPaginator(per int, nums int64) *utils.Paginator {
	p := utils.NewPaginator(base.Ctx.Request, per, nums)
	base.Data["paginator"] = p
	return p
}

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (base *BaseController) SuccessJson(data interface{}) {

	res := JsonResponse{
		200, "success", data,
	}
	base.Data["json"] = res
	base.ServeJSON() //对json进行序列化输出
	base.StopRun()
}

func (base *BaseController) ErrorJson(code int, msg string) {

	res := JsonResponse{
		code, msg, nil,
	}

	base.Data["json"] = res
	base.ServeJSON() //对json进行序列化输出
	base.StopRun()
}
