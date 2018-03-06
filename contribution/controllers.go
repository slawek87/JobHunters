package contribution

import (
	"time"
	"github.com/astaxie/beego/validation"
	"github.com/slawek87/JobHunters/conf"
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

const MongoDBIndex = "Offer"

type ContributionController struct {
	Contribution Contribution
}

func (controller *ContributionController) SetContribution(contribution Contribution) {
	controller.Contribution = contribution
}

func (controller *ContributionController) SetContributionID(ContributionID string) {
	controller.Contribution.ContributionID = bson.ObjectIdHex(ContributionID)
}

func (controller *ContributionController) SetOfferID(OfferID string) {
	controller.Contribution.OfferID = bson.ObjectIdHex(OfferID)
}

func (controller *ContributionController) SetUserID(UserID string) {
	controller.Contribution.UserID = UserID
}

func (controller *ContributionController) GetContribution() Contribution {
	return controller.Contribution
}

func (controller *ContributionController) Create() error {
	session, db := conf.MongoDB()
	defer session.Close()

	controller.Contribution.ContributionID = bson.NewObjectId()
	controller.Contribution.CreatedAt = time.Now()
	controller.Contribution.UpdatedAt = time.Now()

	valid := validation.Validation{}

	isValid, _ := valid.Valid(controller.Contribution)

	if !isValid {
		errorMsg := make(map[string]string)
		for _, err := range valid.Errors {
			errorMsg[err.Field] = err.Message
		}
		results, _ := json.Marshal(errorMsg)
		return errors.New(string(results))
	}

	collection := db.C(MongoDBIndex)

	return collection.Update(
		bson.M{"offer_id": controller.Contribution.OfferID},
		bson.M{"$push": bson.M{"contributions": controller.Contribution}})
}

func (controller *ContributionController) Delete() error {
	session, db := conf.MongoDB()
	defer session.Close()

	collection := db.C(MongoDBIndex)

	return collection.Update(bson.M{
		"offer_id":        controller.Contribution.OfferID,
		"user_id":         controller.Contribution.UserID},
		bson.M{"$pull": bson.M{"contributions": bson.M{"contribution_id": controller.Contribution.ContributionID}}})
}

func (controller *ContributionController) All(query bson.M) ([]Contribution, error) {
	var contributions []Contribution

	session, db := conf.MongoDB()
	defer session.Close()

	err := db.C(MongoDBIndex).Find(query).All(&contributions)
	return contributions, err
}
