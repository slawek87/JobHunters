package user

import "time"

type User struct {
	UserID        string    `json:"user_id" bson:"user_id"`
	FirstName     string    `json:"first_name" bson:"first_name"`
	LastName      string    `json:"last_name" bson:"last_name"`
	Avatar        string    `json:"avatar,omitempty" bson:"avatar"`
	Location      string    `json:"location,omitempty" bson:"location"`
	Position      string    `json:"position,omitempty" bson:"position"`
	Headline      string    `json:"headline,omitempty" bson:"headline"`
	Email         string    `json:"email,omitempty" bson:"email"`
	LinkedIn      string    `json:"linked_in,omitempty" bson:"linked_in"`
	CreatedAt 	  time.Time `json:"user_created_at, @timestamp" bson:"user_created_at"`
	UpdatedAt 	  time.Time `json:"user_updated_at, @timestamp" bson:"user_updated_at"`
}

func (u *User) GetID() string {
    return u.UserID
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetAvatar() string {
	return u.Avatar
}

func (u *User) GetEmail() string {
	return u.Email
}