package user

import (
	"github.com/slawek87/JobHunters/linkedin"
	"gopkg.in/mgo.v2/bson"
	"github.com/slawek87/JobHunters/conf"
	"time"
	"github.com/astaxie/beego/validation"
	"errors"
	"encoding/json"
)

const MongoDBIndex = "User"

type UserController struct {
	User         User
	Authenticate Authenticate
}

func _getProfileData(profileData interface{}) string {
	if profileData != nil {
		return profileData.(string)
	}
	return ""
}

func (controller *UserController) SetUserID(userID bson.ObjectId) {
	controller.User.UserID = userID
}

func (controller *UserController) Auth() error {
	auth := linkedin.Authorization{
		AccessTokenEndpoint: linkedin.ACCESS_TOKEN_ENDPOINT,
		Method:              linkedin.GET,
		GrantType:           linkedin.GRANT_TYPE,
		ResponseType:        linkedin.RESPONSE_TYPE,
		Scope:               linkedin.SCOPE,
		Code:                controller.Authenticate.Code,
		RedirectURI:         linkedin.REDIRECT_URI,
		ClientID:            linkedin.CLIENT_ID,
		ClientSecret:        linkedin.CLIENT_SECRET,
		State:               linkedin.STATE,
	}

	err := auth.ValidateState(controller.Authenticate.State)

	if err != nil {
		return err
	}

	controller.Authenticate.AccessToken, controller.Authenticate.ExpiresIn, controller.Authenticate.ExpirationDateTime = auth.GetToken()

	if controller.Authenticate.AccessToken == "" {
		return errors.New("You are not logged.")
	}

	err = controller.Login()

	return err
}

func (controller *UserController) Login() error {
	profileData := linkedin.RetrieveProfileData(controller.Authenticate.AccessToken, "")

	controller.User.LinkedInID = _getProfileData(profileData["id"])
	controller.User.FirstName = _getProfileData(profileData["firstName"])
	controller.User.LastName = _getProfileData(profileData["lastName"])
	controller.User.Headline = _getProfileData(profileData["headline"])
	controller.User.LinkedInID = _getProfileData(profileData["id"])
	controller.User.LinkedIn = _getProfileData(profileData["linkedIn"])
	controller.User.Location = _getProfileData(profileData["location"])
	controller.User.Email = _getProfileData(profileData["email"])

	err := controller.GetUser()

	if err != nil {
		err = controller.CreateUser()
	}

	controller.User.Authenticate = controller.Authenticate

	return err
}

func (controller *UserController) GetUser() error {
	session, db := conf.MongoDB()
	defer session.Close()

	return db.C(MongoDBIndex).Find(bson.M{"linked_in_id": controller.User.LinkedInID}).One(&controller.User)
}

func (controller *UserController) CreateUser() error {
	session, db := conf.MongoDB()
	defer session.Close()

	controller.User.UserID = bson.NewObjectId()
	controller.User.Company.CompanyID = bson.NewObjectId()
	controller.User.CreatedAt = time.Now()
	controller.User.UpdatedAt = time.Now()

	valid := validation.Validation{}
	isValid, _ := valid.Valid(controller.User)

	if !isValid {
		errorMsg := make(map[string]string)
		for _, err := range valid.Errors {
			errorMsg[err.Field] = err.Message
		}
		results, _ := json.Marshal(errorMsg)
		return errors.New(string(results))
	}

	collection := db.C(MongoDBIndex)
	return collection.Insert(controller.User)
}

func (controller *UserController) Update() error {
	controller.User.UpdatedAt = time.Now()

	if controller.User.Company.Name != "" {
		controller.User.Company.UpdatedAt = time.Now()

		valid := validation.Validation{}
		isValid, _ := valid.Valid(controller.User.Company)

		if !isValid {
			errorMsg := make(map[string]string)
			for _, err := range valid.Errors {
				errorMsg[err.Field] = err.Message
			}
			results, _ := json.Marshal(errorMsg)
			return errors.New(string(results))
		}
	}

	session, db := conf.MongoDB()
	defer session.Close()

	valid := validation.Validation{}
	isValid, _ := valid.Valid(controller.User)

	if !isValid {
		errorMsg := make(map[string]string)
		for _, err := range valid.Errors {
			errorMsg[err.Field] = err.Message
		}
		results, _ := json.Marshal(errorMsg)
		return errors.New(string(results))
	}

	collection := db.C(MongoDBIndex)
	return collection.Update(bson.M{"user_id": &controller.User.UserID}, &controller.User)
}
