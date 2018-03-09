package feedback

import (
"time"
"gopkg.in/mgo.v2/bson"
)

type Feedback struct {
	FeedbackID     bson.ObjectId `json:"feedback_id" bson:"feedback_id" valid:"Required"`
	OfferID        bson.ObjectId `json:"offer_id" bson:"offer_id" form:"-" valid:"Required"`
	CandidateID    bson.ObjectId `json:"candidate_id" bson:"candidate_id" valid:"Required"`
	Messages       []Message     `json:"messages" form:"-" bson:"messages"`
	CreatedAt      time.Time     `json:"created_at" bson:"created_at" form:"-"`
	UpdatedAt      time.Time     `json:"updated_at" bson:"updated_at" form:"-"`
}

type Message struct {
	MessageID        bson.ObjectId `json:"message_id" form:"-" bson:"message_id" valid:"Required"`
	SenderID         string `json:"receiver_id" form:"-" bson:"receiver_id" valid:"Required"`
	SenderFullName   string        `json:"sender_fullname" form:"-" bson:"sender_fullname" valid:"Required"`
	Content          string        `json:"content" form:"content" bson:"content" valid:"Required"`
	IsRead           bool          `json:"is_read" form:"-" bson:"is_read" default:"false"`
	CreatedAt        time.Time     `json:"created_at" bson:"created_at" form:"-"`
	UpdatedAt        time.Time     `json:"updated_at" bson:"updated_at" form:"-"`
}
