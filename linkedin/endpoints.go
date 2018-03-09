package linkedin

var profileDataFullParams = `
	:(id,first-name,last-name,headline,picture-url,industry,summary,specialties,
	positions:(id,title,summary,start-date,end-date,is-current,company:(id,name,type,size,industry,ticker)),
	educations:(id,school-name,field-of-study,start-date,end-date,degree,activities,notes),
	associations,interests,num-recommenders,date-of-birth,
	publications:(id,title,publisher:(name),authors:(id,name),date,url,summary),
	patents:(id,title,summary,number,status:(id,name),
	office:(name),inventors:(id,name),date,url),
	languages:(id,language:(name),proficiency:(level,name)),
	skills:(id,skill:(name)),certifications:(id,name,authority:(name),
	number,start-date,end-date),courses:(id,name,number),
	recommendations-received:(id,recommendation-type,recommendation-text,recommender),
	honors-awards,three-current-positions,three-past-positions,volunteer)
`

// In default function returns whole data what Linkedin returns for profile endpoint.
// To change this, you need to modify `params` argument.
func RetrieveProfileData(token string, params string) map[string]interface{} {
	endpoint := "https://api.linkedin.com/v1/people/~"

	if params == "" {
		params = profileDataFullParams
	}

	linkedInClient := LinkedInClient{
		Token: token,
		Endpoint: endpoint,
		Params: params,
		Method: GET,
	}

	return linkedInClient.Retrieve()
}

// Universal function to retrieve data from Linkedin API.
func RetrieveData(token string, endpoint string, params string, method string) map[string]interface{} {
	linkedInClient := LinkedInClient{
		Token: token,
		Endpoint: endpoint,
		Params: params,
		Method: method,
	}

	return linkedInClient.Retrieve()
}