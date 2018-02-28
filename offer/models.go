package offer

import "time"

const EXPIRATION_TIME_DAYS = 14


type Offer struct {
	OfferID        string    `json:"offer_id" bson:"offer_id" form:"-"`
	Name           string    `json:"name" form:"name"`
	Description    string    `json:"description" form:"description"`
	UserID         string    `json:"user_id"  bson:"user_id" form:"-"`
	City           string    `json:"city" form:"city"`
	Country        string    `json:"country" form:"country"`
	Commission     float64   `json:"commission" form:"commission"`
	Currency       string    `json:"currency" form:"currency"`
	Remote         bool      `json:"remote" form:"remote"`
	ExpirationTime time.Time `json:"expiration_time, @timestamp" bson:"expiration_time" form:"-"`
	CreatedAt 	   time.Time `json:"offer_created_at" bson:"offer_created_at" form:"-"`
	UpdatedAt      time.Time `json:"offer_updated_at" bson:"offer_updated_at" form:"-"`
}
