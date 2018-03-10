package user

import (
	"github.com/astaxie/beego"
)

type UserView struct {
	beego.Controller
	UserController UserController
}

func (view *UserView) Login() {
	results := make(map[string]interface{})

	view.Ctx.Input.Bind(&view.UserController.Authorization.Code, "code")
	view.Ctx.Input.Bind(&view.UserController.Authorization.State, "state")

	err := view.UserController.Authorize()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		view.SetSession("User", &view.UserController.User)
		results["results"] = &view.UserController.User
		view.Data["json"] = results
		view.ServeJSON()
	}
}
