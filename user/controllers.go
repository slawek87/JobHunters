package user

import (
	"github.com/slawek87/JobHunters/conf"
	"gopkg.in/mgo.v2/bson"
)

const MongoDBIndex = "User"

type UserController struct {}

func (uc *UserController) Create(u *User) error {
	session, db := conf.MongoDB()
	defer session.Close()

	c := db.C(MongoDBIndex)
	return c.Insert(u)
}

// b argument is a bson object with data to delete from db.
func (uc *UserController) Delete(b *bson.M) error {
	session, db := conf.MongoDB()
	defer session.Close()

	c := db.C(MongoDBIndex)
	return c.Remove(b)
}

// b argument is a bson object with data to update in User model.
func (uc *UserController) Update(u *User, b *bson.M) error {
	session, db := conf.MongoDB()
	defer session.Close()

	c := db.C(MongoDBIndex)
	return c.Update(b, u)
}

