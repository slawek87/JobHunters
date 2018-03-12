package user

import (
	"github.com/astaxie/beego"
)

type UserView struct {
	beego.Controller
	UserController UserController
}

func (view *UserView) Login() {
	userSession := view.GetSession("User")
	results := make(map[string]interface{})

	if userSession != nil {
		view.UserController.User = *userSession.(*User)
		results["results"] = &view.UserController.User
		view.Data["json"] = results
		view.ServeJSON()
	} else {
		view.Ctx.Input.Bind(&view.UserController.Authenticate.Code, "code")
		view.Ctx.Input.Bind(&view.UserController.Authenticate.State, "state")

		err := view.UserController.Auth()

		if err != nil {
			view.CustomAbort(300, err.Error())
		} else {
			view.SetSession("User", &view.UserController.User)
			results["results"] = &view.UserController.User
			view.Data["json"] = results
			view.ServeJSON()
		}
	}
}

func (view *UserView) UpdateUser() {
	userSession := view.GetSession("User")
	results := make(map[string]interface{})

	view.UserController.User = *userSession.(*User)
	view.UserController.SetUserID(userSession.(*User).UserID)
	view.ParseForm(&view.UserController.User)

	err := view.UserController.Update()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = view.UserController.User
		view.Data["json"] = results
		view.ServeJSON()
	}
}

func (view *UserView) UpdateUserCompany() {
	userSession := view.GetSession("User")
	results := make(map[string]interface{})

	view.UserController.User = *userSession.(*User)
	view.UserController.SetUserID(userSession.(*User).UserID)
	view.ParseForm(&view.UserController.User.Company)

	err := view.UserController.Update()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = view.UserController.User
		view.Data["json"] = results
		view.ServeJSON()
	}
}
