package offer

import (
	"time"
)

const EXPIRATION_TIME_DAYS = 14


type Offer struct {
	OfferID        string    `json:"offer_id" bson:"offer_id" form:"-" valid:"Required"`
	Name           string    `json:"name" form:"name" valid:"Required"`
	Description    string    `json:"description" form:"description" valid:"Required"`
	UserID         string    `json:"user_id"  bson:"user_id" form:"-" valid:"Required"`
	City           string    `json:"city" form:"city" valid:"Required"`
	Country        string    `json:"country" form:"country" valid:"Required"`
	Commission     float64   `json:"commission" form:"commission" valid:"Required"`
	Currency       string    `json:"currency" form:"currency" valid:"Required"`
	Remote         bool      `json:"remote" form:"remote" valid:"Required"`
	ExpirationTime time.Time `json:"expiration_time, @timestamp" bson:"expiration_time" form:"-"`
	CreatedAt 	   time.Time `json:"offer_created_at" bson:"offer_created_at" form:"-"`
	UpdatedAt      time.Time `json:"offer_updated_at" bson:"offer_updated_at" form:"-"`
}
