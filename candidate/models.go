package candidate

import (
	"time"
	"mime/multipart"
	"gopkg.in/mgo.v2/bson"
)

type Candidate struct {
	CandidateID     bson.ObjectId  `json:"candidate_id" bson:"candidate_id" valid:"Required"`
	OfferID         bson.ObjectId  `json:"offer_id" bson:"offer_id" form:"-" valid:"Required"`
	RecruiterID     string         `json:"recruiter_id"  bson:"recruiter_id" form:"-" valid:"Required"`
	FirstName       string         `json:"first_name" form:"first_name" bson:"first_name" valid:"Required"`
	LastName        string         `json:"last_name" form:"last_name" bson:"last_name" valid:"Required"`
	LinkedIn        string         `json:"linked_in,omitempty" form:"linked_in" bson:"linked_in" valid:"Required"`
	Salary          int            `json:"salary" bson:"salary" form:"salary" valid:"Required"`
	Currency        string         `json:"currency" form:"currency" bson:"currency" valid:"Required"`
	CooperationForm string         `json:"cooperation_form" form:"cooperation_form" bson:"cooperation_form" valid:"Required"`
	Description     string         `json:"description" form:"description" bson:"description" valid:"Required"`
	Resume          multipart.File `json:"resume" form:"resume" bson:"-"`
	ResumeID        string         `json:"resume_id" bson:"resume_id"`
	Accepted        bool           `json:"accepted" form:"accepted" bson:"accepted" default:"false"`
	Interview       bool           `json:"interview" form:"interview" bson:"interview" default:"false"`
	Successful      bool           `json:"successful" form:"successful" bson:"successful" default:"false"`
	CreatedAt       time.Time      `json:"created_at" bson:"created_at" form:"-"`
	UpdatedAt       time.Time      `json:"updated_at" bson:"updated_at" form:"-"`
}
