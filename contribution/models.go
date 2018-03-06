package contribution

import (
"time"
"gopkg.in/mgo.v2/bson"
)

type Contribution struct {
	ContributionID bson.ObjectId `json:"contribution_id" bson:"contribution_id" valid:"Required"`
	OfferID        bson.ObjectId `json:"offer_id" bson:"offer_id" form:"-" valid:"Required"`
	UserID         string        `json:"user_id"  bson:"user_id" form:"-" valid:"Required"`
	FirstName      string        `json:"first_name" form:"first_name" bson:"first_name" valid:"Required"`
	LastName       string        `json:"last_name" form:"last_name" bson:"last_name" valid:"Required"`
	Avatar         string        `json:"avatar,omitempty" form:"avatar" bson:"avatar"`
	LinkedIn       string        `json:"linked_in,omitempty" form:"linked_in" bson:"linked_in" valid:"Required"`
	Description    string        `json:"description" form:"description" bson:"description"  valid:"Required"`
	ExpirationTime time.Time     `json:"expiration_time, @timestamp" bson:"expiration_time" form:"-"`
	CreatedAt      time.Time     `json:"created_at" bson:"created_at" form:"-"`
	UpdatedAt      time.Time     `json:"updated_at" bson:"updated_at" form:"-"`
}
