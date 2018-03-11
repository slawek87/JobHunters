package user

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	UserID            bson.ObjectId `json:"user_id" bson:"user_id" valid:"Required"`
	LinkedInID        string        `json:"-,omitempty" bson:"linked_in_id" valid:"Required"`
	FirstName         string        `json:"first_name" bson:"first_name" valid:"Required"`
	LastName          string        `json:"last_name" bson:"last_name" valid:"Required"`
	Avatar            string        `json:"avatar,omitempty" bson:"avatar"`
	Location          string        `json:"location,omitempty" bson:"location"`
	Headline          string        `json:"headline,omitempty" bson:"headline"`
	Email             string        `json:"email,omitempty" bson:"email"`
	LinkedIn          string        `json:"linked_in,omitempty" bson:"linked_in"`
	Authenticate      Authenticate  `json:"authorization" bson:"-"`
	IsBusinessPartner bool          `json:"is_business_partner" form:"-" bson:"is_business_partner" default:"false" valid:"Required"`
	Company           Company		`json:"company" bson:"company" form:"-"`
	CreatedAt         time.Time     `json:"created_at, @timestamp" bson:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at, @timestamp" bson:"updated_at"`
}

type Authenticate struct {
	Code        string `json:"code" form:"code"`
	State       string `json:"state" form:"state"`
	AccessToken string `json:"access_token" form:"-"`
	ExpiresIn   int    `json:"expires_in" form:"-"`
}

type Company struct {
	CompanyID         bson.ObjectId `json:"company_id" bson:"company_id" valid:"Required"`
	Name              string        `json:"name" bson:"name" valid:"Required"`
	Logo              string        `json:"logo" bson:"logo"`
	IsBusinessPartner bool          `json:"is_business_partner" form:"-" bson:"is_business_partner" default:"false" valid:"Required"`
	Webpage           string        `json:"webpage" bson:"webpage"`
	City              string        `json:"city" bson:"city" valid:"Required"`
	Country           string        `json:"country" bson:"country" valid:"Required"`
	Email             string        `json:"email,omitempty" bson:"email" valid:"Required"`
	Description       string        `json:"description" bson:"description"`
	CreatedAt         time.Time     `json:"created_at, @timestamp" bson:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at, @timestamp" bson:"updated_at"`
}

func (model *User) IsActive() bool {
	if model.FirstName != "" && model.LastName != "" && model.Email != "" {
		return true
	}
	return false
}

func (model *Company) IsActive() bool {
	if model.Name != "" && model.City != "" && model.Country != "" && model.Email != "" {
		return true
	}
	return false
}
