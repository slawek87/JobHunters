package linkedin

import (
	"gopkg.in/resty.v0"
	"net/url"
	"time"
	"errors"
)

type Authorization struct {
	AuthorizationEndpoint   string
	AccessTokenEndpoint     string
	Method       			string
	GrantType    			string
	ResponseType 			string
	Scope        			string
	Code         			string
	RedirectURI  			string
	ClientID     			string
	ClientSecret 			string
	State        			string
}

// Function validates state value. This function helps secure app from GET injections.
func (authorization *Authorization) ValidateState(state string) error {
	if state != authorization.State {
		return errors.New("State value is not correct.")
	}

	return nil
}

// Method returns authorization URL.
func (authorization *Authorization) GetAuthorizationURL() string {
	authorizationUrl, _ := url.Parse(authorization.AuthorizationEndpoint)

	query := authorizationUrl.Query()
	query.Set("response_type", authorization.ResponseType)
	query.Set("client_id", authorization.ClientID)
	query.Set("redirect_uri", authorization.RedirectURI)
	query.Set("state", authorization.State)
	query.Set("scope", authorization.Scope)

	authorizationUrl.RawQuery = query.Encode()

	return authorizationUrl.String()
}

// Method returns Linkedin Token and its expiration time in seconds.
func (authorization *Authorization) GetToken() (string, int) {
	return authorization.ExchangeAuthorizationCode()
}

// Method calls Linkedin API to exchange authorization code to Token.
func (authorization *Authorization) ExchangeAuthorizationCode() (string, int) {
	var result map[string]interface{}

	accessToken := ""
	expiresIn := 0

	formData := map[string]string {
			"grant_type": authorization.GrantType,
			"code": authorization.Code,
			"redirect_uri": authorization.RedirectURI,
			"client_id": authorization.ClientID,
			"client_secret": authorization.ClientSecret,
	}

	resty.SetTimeout(time.Duration(1 * time.Minute))
	resty.SetHeaders(map[string]string{"Content-Type":  "application/x-www-form-urlencoded"})

	if ValidateResponse(resty.R().
		SetFormData(formData).
		SetResult(&result).
		Post(authorization.AccessTokenEndpoint)) == true {
		accessToken = result["access_token"].(string)
		expiresIn = int(result["expires_in"].(float64))
	}

	return accessToken, expiresIn
}