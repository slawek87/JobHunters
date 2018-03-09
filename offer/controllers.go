package offer

import (
	"github.com/slawek87/JobHunters/conf"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/astaxie/beego/validation"
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2"
	"github.com/slawek87/JobHunters/contribution"
)

const MongoDBIndex = "Offer"

type OfferController struct {
	Offer Offer
}

func MigrateDB() {
	session, db := conf.MongoDB()
	defer session.Close()

	offerDB := db.C(MongoDBIndex)
	offerIndex := mgo.Index{
		Key: []string{"$text:description"},
	}

	err := offerDB.EnsureIndex(offerIndex)
	if err != nil {
		panic(err)
	}
}

func (controller *OfferController) SetOffer(offer Offer) {
	controller.Offer = offer
}

func (controller *OfferController) SetOfferID(offerID string) {
	controller.Offer.OfferID = bson.ObjectIdHex(offerID)
}

func (controller *OfferController) SetUserID(userID string) {
	controller.Offer.UserID = userID
}

func (controller *OfferController) GetOffer() Offer {
	return controller.Offer
}

func (controller *OfferController) Create() error {
	session, db := conf.MongoDB()
	defer session.Close()

	controller.Offer.OfferID = bson.NewObjectId()
	controller.Offer.CreatedAt = time.Now()
	controller.Offer.UpdatedAt = time.Now()
	controller.Offer.ExpirationTime = time.Now().AddDate(0, 0, EXPIRATION_TIME_DAYS)

	valid := validation.Validation{}
	isValid, _ := valid.Valid(controller.Offer)

	if !isValid {
		errorMsg := make(map[string]string)
		for _, err := range valid.Errors {
			errorMsg[err.Field] = err.Message
		}
		results, _ := json.Marshal(errorMsg)
		return errors.New(string(results))
	}

	c := db.C(MongoDBIndex)
	return c.Insert(controller.Offer)
}

func (controller *OfferController) Get() (interface{}, error) {
	var offer Offer
	var result struct {
		Offer
		Contributions []contribution.Contribution `json:"contributions"`
	}

	session, db := conf.MongoDB()
	defer session.Close()

	err := db.C(MongoDBIndex).Find(bson.M{"offer_id": controller.Offer.OfferID}).One(&offer)
	if err != nil {
		return nil, err
	}

	collection := contribution.ContributionController{}
	contributions, err := collection.Find(bson.M{"offer_id": controller.Offer.OfferID})

	result.Offer = offer
	result.Contributions = contributions

	return result, err
}

func (controller *OfferController) Delete() error {
	session, db := conf.MongoDB()
	defer session.Close()

	collection := db.C(MongoDBIndex)
	return collection.Remove(bson.M{
		"offer_id": controller.Offer.OfferID,
		"user_id":  controller.Offer.UserID})
}

func (controller *OfferController) Update() error {
	controller.Offer.UpdatedAt = time.Now()

	session, db := conf.MongoDB()
	defer session.Close()

	valid := validation.Validation{}
	isValid, _ := valid.Valid(controller.Offer)

	if !isValid {
		errorMsg := make(map[string]string)
		for _, err := range valid.Errors {
			errorMsg[err.Field] = err.Message
		}
		results, _ := json.Marshal(errorMsg)
		return errors.New(string(results))
	}

	collection := db.C(MongoDBIndex)
	return collection.Update(bson.M{"offer_id": &controller.Offer.OfferID}, &controller.Offer)
}

// list all records
func (controller *OfferController) Find(query bson.M) ([]Offer, error) {
	var offers []Offer

	session, db := conf.MongoDB()
	defer session.Close()

	collection := db.C(MongoDBIndex).Find(query).All(&offers)
	return offers, collection
}
