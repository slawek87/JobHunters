package user

import (
	"github.com/astaxie/beego/context"
	"os/user"
)

// need to be modify - set correct redirection url.
var OnlyAuthorizedUserCanModify = func(ctx *context.Context) {
	authorizedMethods := []string{"PUT", "DELETE", "POST"}

	for _, method := range authorizedMethods {
		if ctx.Request.Method == method {
			_, ok := ctx.Input.Session("User").(user.User)
			if !ok {
				ctx.Redirect(302, "/user/login")
				break
			}
		}
	}
}

// need to be modify - set correct redirection url.
var OnlyAuthorizedUser = func(ctx *context.Context) {
	authorizedMethods := []string{"GET", "PUT", "DELETE", "POST"}

	for _, method := range authorizedMethods {
		if ctx.Request.Method == method {
			_, ok := ctx.Input.Session("User").(user.User)
			if !ok {
				ctx.Redirect(302, "/user/login")
				break
			}
		}
	}
}

