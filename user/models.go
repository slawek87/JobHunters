package user

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	UserID        bson.ObjectId `json:"user_id" bson:"user_id" valid:"Required"`
	LinkedInID    string        `json:"-,omitempty" bson:"linked_in_id" valid:"Required"`
	FirstName     string        `json:"first_name" bson:"first_name" valid:"Required"`
	LastName      string        `json:"last_name" bson:"last_name" valid:"Required"`
	Avatar        string        `json:"avatar,omitempty" bson:"avatar"`
	Location      string        `json:"location,omitempty" bson:"location"`
	Headline      string        `json:"headline,omitempty" bson:"headline"`
	Email         string        `json:"email,omitempty" bson:"email"`
	LinkedIn      string        `json:"linked_in,omitempty" bson:"linked_in"`
	Authorization Authorization `json:"authorization" bson:"-"`
	CreatedAt     time.Time     `json:"created_at, @timestamp" bson:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at, @timestamp" bson:"updated_at"`
}

type Authorization struct {
	Code        string `json:"code" form:"code"`
	State       string `json:"state" form:"state"`
	AccessToken string `json:"access_token" form:"-"`
	ExpiresIn   int    `json:"expires_in" form:"-"`
}


func (model *User) IsActive() bool {
	if model.FirstName != "" && model.LastName != "" && model.Email != "" {
		return true
	}
	return false
}