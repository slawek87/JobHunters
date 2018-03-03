package candidate

import "time"

type Candidate struct {
	CandidateID     string    `json:"contribution_id" bson:"contribution_id" valid:"Required"`
	OfferID         string    `json:"offer_id" bson:"offer_id" form:"-" valid:"Required"`
	RecruiterID     string    `json:"recruiter_id"  bson:"recruiter_id" form:"-" valid:"Required"`
	FirstName       string    `json:"first_name" form:"first_name" bson:"first_name" valid:"Required"`
	LastName        string    `json:"last_name" form:"last_name" bson:"last_name" valid:"Required"`
	LinkedIn        string    `json:"linked_in,omitempty" form:"linked_in" bson:"linked_in" valid:"Required"`
	Salary          int       `json:"salary" bson:"salary" form:"salary" valid:"Required"`
	CooperationForm string    `json:"cooperation_form" form:"cooperation_form" bson:"cooperation_form" valid:"Required"`
	Description     string    `json:"description" form:"description" bson:"description" valid:"Required"`
	Resume          []byte    `json:"resume" form:"resume" bson:"resume"`
	Accepted        bool      `json:"accepted" form:"accepted" bson:"accepted" default:"false" valid:"Required"`
	Interview       bool      `json:"interview" form:"interview" bson:"interview" default:"false" valid:"Required"`
	Successful      bool      `json:"successful" form:"successful" bson:"successful" default:"false" valid:"Required"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at" form:"-"`
	UpdatedAt       time.Time `json:"updated_at" bson:"updated_at" form:"-"`
}
