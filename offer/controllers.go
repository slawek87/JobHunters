package offer

import (
	"github.com/slawek87/JobHunters/conf"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/rs/xid"
	"github.com/astaxie/beego/validation"
	"encoding/json"
	"errors"
	"github.com/slawek87/JobHunters/contribution"
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
	var result struct{
		Offer
		Contributions    []contribution.Contribution `json:"contributions"`
	}

	session, db := conf.MongoDB()
	defer session.Close()

	err := db.C(MongoDBIndex).Find(bson.M{"offer_id": controller.Offer.OfferID}).One(&offer)

	if err != nil {
		return nil, err
	}

	c := contribution.ContributionController{}
	contributions, err := c.All(bson.M{"offer_id": controller.Offer.OfferID})

	result.Offer = offer
    result.Contributions = contributions


    return result, err
}

func (controller *OfferController) Delete() error {
	session, db := conf.MongoDB()
	defer session.Close()

	c := db.C(MongoDBIndex)
	return c.Remove(bson.M{
		"offer_id": controller.Offer.OfferID,
		"user_id": controller.Offer.UserID})
}

func (controller *OfferController) Update() error {
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