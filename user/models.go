package user

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	UserID            bson.ObjectId `json:"user_id" bson:"user_id" form:"-" valid:"Required"`
	LinkedInID        string        `json:"-" bson:"linkedin_id" form:"-" valid:"Required"`
	FirstName         string        `json:"first_name" form:"first_name" bson:"first_name" valid:"Required"`
	LastName          string        `json:"last_name" bson:"last_name" form:"last_name" valid:"Required"`
	Avatar            []byte        `json:"avatar" bson:"avatar" form:"avatar"`
	Location          string        `json:"location" bson:"location" form:"location"`
	Headline          string        `json:"headline" bson:"headline" form:"headline"`
	Email             string        `json:"email" bson:"email" form:"email"`
	LinkedIn          string        `json:"linked_in,omitempty" bson:"linked_in" form:"linked_in"`
	Authenticate      Authenticate  `json:"authorization" bson:"-"`
	IsBusinessPartner bool          `json:"is_business_partner" form:"-" bson:"is_business_partner" default:"false" valid:"Required"`
	Company           Company       `json:"company" bson:"company" form:"-"`
	CreatedAt         time.Time     `json:"created_at, @timestamp" bson:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at, @timestamp" bson:"updated_at"`
}

type Authenticate struct {
	Code               string    `json:"code" form:"code"`
	State              string    `json:"state" form:"state"`
	AccessToken        string    `json:"access_token" form:"-"`
	ExpiresIn          int       `json:"expires_in" form:"-"`
	ExpirationDateTime time.Time `json:"expiration_datetime" form:"-"`
}

type Company struct {
	CompanyID         bson.ObjectId `json:"company_id" bson:"company_id" form:"-"`
	Name              string        `json:"name" bson:"name" form:"name" valid:"Required"`
	Logo              []byte        `json:"logo" bson:"logo" form:"logo"`
	IsBusinessPartner bool          `json:"is_business_partner" form:"-" bson:"is_business_partner" default:"false" valid:"Required"`
	Webpage           string        `json:"webpage" bson:"webpage" form:"webpage"`
	City              string        `json:"city" bson:"city" form:"city" valid:"Required"`
	Country           string        `json:"country" bson:"country" form:"country" valid:"Required"`
	Email             string        `json:"email,omitempty" bson:"email" form:"email" valid:"Required"`
	Description       string        `json:"description" bson:"description" form:"description"`
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

func (model *Authenticate) IsExpired() bool {
	return model.ExpirationDateTime.After(time.Now())
}
