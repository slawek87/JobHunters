package user

import (
	"github.com/astaxie/beego/context"
)

// need to be modify - set correct redirection url && update url.
var OnlyAuthorizedAndActiveUserCanModify = func(ctx *context.Context) {
	authorizedMethods := []string{"GET", "PUT", "DELETE", "POST"}
	user, ok := ctx.Input.Session("User").(*User)

	for _, method := range authorizedMethods {
		if ctx.Request.Method == method {
			if !ok {
				ctx.Redirect(302, "/user/login")
				break
			}
			if  !user.IsActive() {
				ctx.Redirect(302, "/user/update")
				break
			}
			if  user.Company.CompanyID != "" && !user.Company.IsActive() {
				ctx.Redirect(302, "/user/company/update")
				break
			}
		}
	}
}

// need to be modify - set correct redirection url && update url.
var OnlyAuthorizedUserCanModify = func(ctx *context.Context) {
	authorizedMethods := []string{"GET", "PUT", "DELETE", "POST"}
	_, ok := ctx.Input.Session("User").(*User)

	for _, method := range authorizedMethods {
		if ctx.Request.Method == method {
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
	user, ok := ctx.Input.Session("User").(*User)

	for _, method := range authorizedMethods {
		if ctx.Request.Method == method {
			if !ok {
				ctx.Redirect(302, "/user/login")
				break
			}
			if  !user.IsActive() {
				ctx.Redirect(302, "/user/update")
				break
			}
			if  user.Company.CompanyID != "" && !user.Company.IsActive() {
				ctx.Redirect(302, "/user/company/update")
				break
			}
		}
	}
}

