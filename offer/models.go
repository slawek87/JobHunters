package offer

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

const EXPIRATION_TIME_DAYS = 14

type Offer struct {
	OfferID        bson.ObjectId               `json:"offer_id" bson:"offer_id" form:"-" valid:"Required"`
	Name           string                      `json:"name" form:"name" valid:"Required"`
	Description    string                      `json:"description" form:"description" valid:"Required"`
	UserID         string                      `json:"user_id"  bson:"user_id" form:"-" valid:"Required"`
	City           string                      `json:"city" form:"city" bson:"city" valid:"Required"`
	Country        string                      `json:"country" form:"country" bson:"country" valid:"Required"`
	Commission     float64                     `json:"commission" form:"commission" bson:"commission" valid:"Required"`
	Currency       string                      `json:"currency" form:"currency" bson:"currency" valid:"Required"`
	Remote         bool                        `json:"remote" form:"remote" bson:"remote" valid:"Required"`
	ExpirationTime time.Time                   `json:"expiration_time, @timestamp" bson:"expiration_time" form:"-"`
	CreatedAt      time.Time                   `json:"created_at" bson:"created_at" form:"-"`
	UpdatedAt      time.Time                   `json:"updated_at" bson:"updated_at" form:"-"`
}
