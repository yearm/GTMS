package index

import (
	"GTMS/library/controller"
	"GTMS/library/helper"
)

type MainController struct {
	controller.BaseController
}

func (this *MainController) Index() {
	this.SuccessWithData(helper.JSON{
		"hello": "world",
	})
}
