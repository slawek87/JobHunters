package offer

import (
	"github.com/slawek87/JobHunters/conf"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/rs/xid"
)

const MongoDBIndex = "Offer"

type OfferController struct {
	Offer	Offer
}

func (controller *OfferController) SetOffer(offer Offer) {
	controller.Offer = offer
}

func (controller *OfferController) SetOfferID(OfferID string) {
	controller.Offer.OfferID = OfferID
}

func (controller *OfferController) SetUserID(UserID string) {
	controller.Offer.UserID = UserID
}

func (controller *OfferController) GetOffer() Offer {
	return controller.Offer
}

func (controller *OfferController) Create() error {
	getUniqueID := xid.New()

	session, db := conf.MongoDB()
	defer session.Close()

	controller.Offer.OfferID = getUniqueID.String()
	controller.Offer.CreatedAt = time.Now()
	controller.Offer.UpdatedAt = time.Now()
	controller.Offer.ExpirationTime = time.Now().AddDate(0, 0, EXPIRATION_TIME_DAYS)

	c := db.C(MongoDBIndex)
	return c.Insert(controller.Offer)
}

// b argument is a bson object with data to delete from db.
func (controller *OfferController) Delete(b *bson.M) error {
	session, db := conf.MongoDB()
	defer session.Close()

	c := db.C(MongoDBIndex)
	return c.Remove(b)
}

// b argument is a bson object with data to update in User model.
func (controller *OfferController) Update() error {
	session, db := conf.MongoDB()
	defer session.Close()

	c := db.C(MongoDBIndex)
	err := c.Update(bson.M{"offer_id": &controller.Offer.OfferID}, &controller.Offer)

	return err
}

// list all records
func (controller *OfferController) All(query bson.M) ([]Offer, error) {
	var offers []Offer

	session, db := conf.MongoDB()
	defer session.Close()

	c := db.C(MongoDBIndex).Find(query).All(&offers)
	return offers, c
}