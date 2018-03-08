package main

import (
	"github.com/astaxie/beego"
	"github.com/slawek87/JobHunters/offer"
	"github.com/slawek87/JobHunters/contribution"
	"github.com/slawek87/JobHunters/candidate"
	"github.com/slawek87/JobHunters/feedback"
)

func main() {
	offer.MigrateDB()

	beego.Router("/offer", &offer.OfferView{}, "post:Post")
	beego.Router("/offer/list", &offer.OfferView{}, "get:List")
	beego.Router("/offer/:offerID:string", &offer.OfferView{}, "get:Get")
	beego.Router("/offer/:offerID:string", &offer.OfferView{}, "put:Put")
	beego.Router("/offer/:offerID:string", &offer.OfferView{}, "delete:Delete")
	beego.Router("/offer/:offerID:string/contribution", &contribution.ContributionView{}, "post:Post")
	beego.Router("/offer/:offerID:string/contribution/:contributionID:string", &contribution.ContributionView{}, "delete:Delete")
	beego.Router("/offer/:offerID:string/candidate", &candidate.CandidateView{}, "post:Post")
	beego.Router("/offer/:offerID:string/candidate/list", &candidate.CandidateView{}, "get:List")
	beego.Router("/offer/:offerID:string/candidate/resume/:resumeID:string.pdf", &candidate.CandidateView{}, "get:DownloadResume")
	beego.Router("/offer/:offerID:string/candidate/:candidateID:string/feedback", &feedback.FeedbackView{}, "post:SendFeedback")
	beego.Router("/offer/:offerID:string/candidate/:candidateID:string/feedback/:feedbackID:string", &feedback.FeedbackView{}, "post:SendFeedback")
	beego.Router("/offer/:offerID:string/candidate/:candidateID:string/feedback/:feedbackID:string", &feedback.FeedbackView{}, "get:ReceiveFeedback")
	beego.Router("/offer/:offerID:string/candidate/:candidateID:string", &candidate.CandidateView{}, "put:Put")
	beego.Router("/offer/:offerID:string/candidate/:candidateID:string", &candidate.CandidateView{}, "delete:Delete")
	beego.Run("localhost:8000")
}
