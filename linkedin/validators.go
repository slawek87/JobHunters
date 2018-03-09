package linkedin

import (
	"gopkg.in/resty.v0"
	"log"
)

// Function validates response. If response is not valid returns `false` and call logs with error details
func ValidateResponse(response *resty.Response, err error) bool {
	var validation bool

	if response.RawResponse.StatusCode != 200 {
		log.Println(response.RawResponse.Status + " " + response.Request.URL)
		validation = false
	} else if err != nil {
		log.Println(err)
		validation = false
	} else {
		validation = true
	}

	return validation
}
