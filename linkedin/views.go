package linkedin

import (
	"github.com/astaxie/beego"
)

type LinkedinView struct {
	beego.Controller
}

func (view *LinkedinView) GetAuthorizationURL() {
	results := make(map[string]interface{})

	auth := Authorization {
		AuthorizationEndpoint: AUTHORIZATION_ENDPOINT,
		AccessTokenEndpoint:   ACCESS_TOKEN_ENDPOINT,
		Method:                GET,
		GrantType:             GRANT_TYPE,
		ResponseType:          RESPONSE_TYPE,
		Scope:                 SCOPE,
		RedirectURI:           REDIRECT_URI,
		ClientID:              CLIENT_ID,
		ClientSecret:          CLIENT_SECRET,
		State:                 STATE,
	}

	results["results"] = auth.GetAuthorizationURL()
	view.Data["json"] = results
	view.ServeJSON()
}