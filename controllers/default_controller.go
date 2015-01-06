package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.TplNames = "index.tpl"
}
