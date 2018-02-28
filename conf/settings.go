package conf

import (
	"gopkg.in/mgo.v2"
)

// Returns MongoDB session with selected DB.
func MongoDB() (*mgo.Session, *mgo.Database) {
	session, _ := mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)
	return session, session.DB("JobHunters")
}