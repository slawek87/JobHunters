package contribution

import (
	"time"
	"github.com/astaxie/beego/validation"
	"github.com/slawek87/JobHunters/conf"
	"encoding/json"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

const MongoDBIndex = "Contribution"

type ContributionController struct {
	Contribution Contribution
}

func (controller *ContributionController) SetContribution(contribution Contribution) {
	controller.Contribution = contribution
}

func (controller *ContributionController) SetContributionID(contributionID string) {
	controller.Contribution.ContributionID = bson.ObjectIdHex(contributionID)
}

func (controller *ContributionController) SetOfferID(offerID string) {
	controller.Contribution.OfferID = bson.ObjectIdHex(offerID)
}

func (controller *ContributionController) SetUserID(userID string) {
	controller.Contribution.UserID = userID
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
	return collection.Insert(controller.Contribution)
}

func (controller *ContributionController) Delete() error {
	session, db := conf.MongoDB()
	defer session.Close()

	collection := db.C(MongoDBIndex)
	return collection.Remove(bson.M{
		"contribution_id": controller.Contribution.ContributionID,
		"offer_id": controller.Contribution.OfferID,
		"user_id": controller.Contribution.UserID})
}

func (controller *ContributionController) Find(query bson.M) ([]Contribution, error) {
	var contributions []Contribution

	session, db := conf.MongoDB()
	defer session.Close()

	err := db.C(MongoDBIndex).Find(query).All(&contributions)
	return contributions, err
}
