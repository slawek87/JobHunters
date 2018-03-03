package candidate

const MongoDBIndex = "Candidate"

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

func (controller *CandidateController) GetCandidate() Candidate {
	return controller.Candidate
}
