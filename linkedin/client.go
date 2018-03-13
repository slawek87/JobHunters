package linkedin

import (
	"gopkg.in/resty.v0"
	"time"
)

type LinkedInClient struct {
	Token         string
	Endpoint      string
	Params        string
	Method        string
}

// Generic method to retrieve data from Linkedin API.
func (linkedInClient *LinkedInClient) Retrieve() map[string]interface{} {
	var result map[string]interface{}
	var response *resty.Response
	var err error

	url := PrepareURL(linkedInClient.Endpoint, linkedInClient.Params)
	headers := map[string]string{
		"Accept": "application/json",
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + linkedInClient.Token}

	resty.SetTimeout(time.Duration(1 * time.Minute))
	resty.SetHeaders(headers)

	request := resty.R().
		SetResult(&result)

    if linkedInClient.Method == POST {
		response, err = request.Post(url)
	} else {
		response, err = request.Get(url)
	}

	if ValidateResponse(response, err) == true {
		return result
	}

	return nil
}
