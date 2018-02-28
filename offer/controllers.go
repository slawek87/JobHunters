package offer

import (
	"github.com/slawek87/JobHunters/conf"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const MongoDBIndex = "Offer"

type OfferController struct {
	Offer	Offer
}

func (oc *OfferController) Create() error {
	session, db := conf.MongoDB()
	defer session.Close()

	oc.Offer.CreatedAt = time.Now()
	oc.Offer.UpdatedAt = time.Now()
	oc.Offer.ExpirationTime = time.Now().AddDate(0, 0, EXPIRATION_TIME_DAYS)

	c := db.C(MongoDBIndex)
	return c.Insert(oc.Offer)
}

// b argument is a bson object with data to delete from db.
func (oc *OfferController) Delete(b *bson.M) error {
	session, db := conf.MongoDB()
	defer session.Close()

	c := db.C(MongoDBIndex)
	return c.Remove(b)
}

// b argument is a bson object with data to update in User model.
func (oc *OfferController) Update(b bson.M) error {
	session, db := conf.MongoDB()
	defer session.Close()

	c := db.C(MongoDBIndex)
	return c.Update(b, oc.Offer)
}

// list all records
func (oc *OfferController) All(b bson.M) ([]Offer, error) {
	var offers []Offer

	session, db := conf.MongoDB()
	defer session.Close()

	c := db.C(MongoDBIndex).Find(b).All(&offers)

	return offers, c
}