package candidate

import (
	"github.com/rs/xid"
	"time"
	"github.com/astaxie/beego/validation"
	"github.com/slawek87/JobHunters/conf"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"gopkg.in/mgo.v2/bson"
)

const MongoDBIndex = "Candidate"
const MongoDBFS = "Resume"

type CandidateController struct {
	Candidate Candidate
}

func (controller *CandidateController) SetCandidate(candidate Candidate) {
	controller.Candidate = candidate
}

func (controller *CandidateController) SetCandidateID(candidateID string) {
	controller.Candidate.CandidateID = candidateID
}

func (controller *CandidateController) SetOfferID(OfferID string) {
	controller.Candidate.OfferID = OfferID
}

func (controller *CandidateController) SetRecruiterID(RecruiterID string) {
	controller.Candidate.RecruiterID = RecruiterID
}

func (controller *CandidateController) SetResumeID(ResumeID string) {
	controller.Candidate.ResumeID = ResumeID
}

func (controller *CandidateController) GetCandidate() Candidate {
	return controller.Candidate
}

func (controller *CandidateController) GetCandidateFullName() string {
	return controller.Candidate.FirstName + " " + controller.Candidate.LastName
}

func (controller *CandidateController) Create() error {
	getUniqueID := xid.New()

	session, db := conf.MongoDB()
	defer session.Close()

	controller.Candidate.CandidateID = getUniqueID.String()
	controller.Candidate.CreatedAt = time.Now()
	controller.Candidate.UpdatedAt = time.Now()

	if controller.Candidate.Resume != nil {
		resumeName := strings.Replace(controller.GetCandidateFullName(), " ", "_", -1) + ".pdf"
		controller.SetResumeID(controller.Candidate.CandidateID)

		file, _ := db.GridFS(MongoDBFS).Create(resumeName)

		io.Copy(file, controller.Candidate.Resume)

		file.SetId(controller.Candidate.ResumeID)
		file.Close()
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

	c := db.C(MongoDBIndex)
	return c.Insert(controller.Candidate)
}

func (controller *CandidateController) Update() error {
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

	c := db.C(MongoDBIndex)
	err := c.Update(bson.M{"candidate_id": &controller.Candidate.CandidateID}, &controller.Candidate)

	if controller.Candidate.Resume != nil {
		controller.DeleteResume()

		resumeName := strings.Replace(controller.GetCandidateFullName(), " ", "_", -1) + ".pdf"
		controller.SetResumeID(controller.Candidate.CandidateID)

		file, _ := db.GridFS(MongoDBFS).Create(resumeName)

		io.Copy(file, controller.Candidate.Resume)

		file.SetId(controller.Candidate.ResumeID)
		file.Close()
	}

	return err
}

func (controller *CandidateController) DeleteResume() error {
	session, db := conf.MongoDB()
	defer session.Close()

	file := db.GridFS(MongoDBFS)
	return file.RemoveId(controller.Candidate.CandidateID)
}

func (controller *CandidateController) Delete() error {
	session, db := conf.MongoDB()
	defer session.Close()

	controller.DeleteResume()

	c := db.C(MongoDBIndex)
	return c.Remove(bson.M{
		"candidate_id": controller.Candidate.CandidateID,
		"offer_id": controller.Candidate.OfferID,
		"recruiter_id":  controller.Candidate.RecruiterID})
}
