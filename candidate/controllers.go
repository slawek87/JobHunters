package candidate

import (
	"time"
	"github.com/astaxie/beego/validation"
	"github.com/slawek87/JobHunters/conf"
	"encoding/json"
	"errors"
	"os"
	"gopkg.in/mgo.v2/bson"
)

const MongoDBIndex = "Candidate"
const StoragePath = "static/resumes/"

type CandidateController struct {
	Candidate Candidate
}

func (controller *CandidateController) SetCandidate(candidate Candidate) {
	controller.Candidate = candidate
}

func (controller *CandidateController) SetCandidateID(candidateID string) {
	controller.Candidate.CandidateID = bson.ObjectIdHex(candidateID)
}

func (controller *CandidateController) SetOfferID(offerID string) {
	controller.Candidate.OfferID = bson.ObjectIdHex(offerID)
}

func (controller *CandidateController) SetRecruiterID(recruiterID string) {
	controller.Candidate.RecruiterID = recruiterID
}

func (controller *CandidateController) GetCandidate() Candidate {
	return controller.Candidate
}

func (controller *CandidateController) GetCandidateFullName() string {
	return controller.Candidate.FirstName + " " + controller.Candidate.LastName
}

func (controller *CandidateController) GetResumePath() string {
	return StoragePath + controller.Candidate.ResumeID
}


func (controller *CandidateController) Create() error {
	session, db := conf.MongoDB()
	defer session.Close()

	controller.Candidate.CandidateID = bson.NewObjectId()
	controller.Candidate.CreatedAt = time.Now()
	controller.Candidate.UpdatedAt = time.Now()

	if controller.Candidate.Resume != nil {
		controller.Candidate.ResumeID =
			string(controller.Candidate.CandidateID.Hex()) + "_" +
				controller.Candidate.FirstName + "_" + controller.Candidate.LastName + ".pdf"
	}

	valid := validation.Validation{}
	isValid, _ := valid.Valid(controller.Candidate)

	if !isValid {
		errorMsg := make(map[string]string)
		for _, err := range valid.Errors {
			errorMsg[err.Field] = err.Message
		}
		results, _ := json.Marshal(errorMsg)
		return errors.New(string(results))
	}

	collection := db.C(MongoDBIndex)
	return collection.Insert(controller.Candidate)
}

func (controller *CandidateController) Update() error {
	controller.Candidate.UpdatedAt = time.Now()

	session, db := conf.MongoDB()
	defer session.Close()

	valid := validation.Validation{}
	isValid, _ := valid.Valid(controller.Candidate)

	if !isValid {
		errorMsg := make(map[string]string)
		for _, err := range valid.Errors {
			errorMsg[err.Field] = err.Message
		}
		results, _ := json.Marshal(errorMsg)
		return errors.New(string(results))
	}

	if controller.Candidate.Resume != nil {
		controller.Candidate.ResumeID =
			string(controller.Candidate.CandidateID.Hex()) + "_" +
				controller.Candidate.FirstName + "_" + controller.Candidate.LastName + ".pdf"
	}

	collection := db.C(MongoDBIndex)
	return collection.Update(bson.M{"candidate_id": &controller.Candidate.CandidateID}, &controller.Candidate)
}

func (controller *CandidateController) DeleteResume() error {
	session, db := conf.MongoDB()
	defer session.Close()

	collection := db.C(MongoDBIndex)
	err := collection.Find(bson.M{
		"candidate_id": controller.Candidate.CandidateID,
		"offer_id": controller.Candidate.OfferID,
		"recruiter_id":  controller.Candidate.RecruiterID}).One(&controller.Candidate)

	if err != nil {
		return err
	}

	return os.Remove(controller.GetResumePath())
}

func (controller *CandidateController) Delete() error {
	session, db := conf.MongoDB()
	defer session.Close()

	controller.DeleteResume()

	collection := db.C(MongoDBIndex)
	return collection.Remove(bson.M{
		"candidate_id": controller.Candidate.CandidateID,
		"offer_id": controller.Candidate.OfferID})
}

func (controller *CandidateController) Find(query bson.M) ([]Candidate, error) {
	var candidates []Candidate

	session, db := conf.MongoDB()
	defer session.Close()

	collection := db.C(MongoDBIndex).Find(query).All(&candidates)
	return candidates, collection
}

func (controller *CandidateController) Get() error {
	session, db := conf.MongoDB()
	defer session.Close()

	return db.C(MongoDBIndex).Find(bson.M{"candidate_id": controller.Candidate.CandidateID}).One(&controller.Candidate)
}

func (controller *CandidateController) DownloadResume() (string) {
	return controller.GetResumePath()
}