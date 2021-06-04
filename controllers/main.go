package test

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}
