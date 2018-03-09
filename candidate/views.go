package candidate

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type CandidateView struct {
	beego.Controller
	CandidateController CandidateController
	//Session session.Store
}

func (view *CandidateView) Post() {
	results := make(map[string]interface{})
	//getSession := v.StartSession()
	view.ParseForm(&view.CandidateController.Candidate)
	view.CandidateController.SetOfferID(view.Ctx.Input.Param(":offerID"))
	view.CandidateController.SetRecruiterID("Xyz123") //getSession.Get("ID").(string)
	view.CandidateController.Candidate.Resume, _, _= view.GetFile("resume")

	err := view.CandidateController.Create()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		if view.CandidateController.Candidate.Resume != nil {
			view.SaveToFile("resume", view.CandidateController.GetResumePath())
		}
		results["results"] = view.CandidateController.GetCandidate()
		view.Data["json"] = results
		view.ServeJSON()
	}
}

func (view *CandidateView) Put() {
	results := make(map[string]interface{})

	view.ParseForm(&view.CandidateController.Candidate)
	view.CandidateController.SetCandidateID(view.Ctx.Input.Param(":candidateID"))
	view.CandidateController.SetOfferID(view.Ctx.Input.Param(":offerID"))
	view.CandidateController.SetRecruiterID("Xyz123") //getSession.Get("ID").(string)
	view.CandidateController.Candidate.Resume, _, _= view.GetFile("resume")

	err := view.CandidateController.Update()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		if view.CandidateController.Candidate.Resume != nil {
			view.SaveToFile("resume", view.CandidateController.GetResumePath())
		}
		results["results"] = view.CandidateController.Candidate
		view.Data["json"] = results
		view.ServeJSON()
	}
}

func (view *CandidateView) Get() {
	results := make(map[string]interface{})

	view.CandidateController.SetCandidateID(view.Ctx.Input.Param(":candidateID"))
	view.CandidateController.SetOfferID(view.Ctx.Input.Param(":offerID"))

	err := view.CandidateController.Get()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = view.CandidateController.Candidate
		view.Data["json"] = results
		view.ServeJSON()
	}
}

func (view *CandidateView) List() {
	results := make(map[string]interface{})
	view.CandidateController.SetOfferID(view.Ctx.Input.Param(":offerID"))

	candidates, err := view.CandidateController.Find(
		bson.M{"offer_id": view.CandidateController.Candidate.OfferID})

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = candidates
		view.Data["json"] = results
		view.ServeJSON()
	}
}

func (view *CandidateView) Delete() {
	results := make(map[string]interface{})

	view.CandidateController.SetCandidateID(view.Ctx.Input.Param(":candidateID"))
	view.CandidateController.SetOfferID(view.Ctx.Input.Param(":offerID"))
	view.CandidateController.SetRecruiterID("Xyz123") //getSession.Get("ID").(string)

	err := view.CandidateController.Delete()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = "Candidate has been deleted."
		view.Data["json"] = results
		view.ServeJSON()
	}
}

func (view *CandidateView) DownloadResume() {
	view.CandidateController.Candidate.ResumeID = view.Ctx.Input.Param(":resumeID") + ".pdf"
	view.Ctx.Output.Download(view.CandidateController.DownloadResume(), view.CandidateController.Candidate.ResumeID)
}