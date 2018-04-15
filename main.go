package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/slawek87/JobHunters/offer"
	"github.com/slawek87/JobHunters/contribution"
	"github.com/slawek87/JobHunters/candidate"
	"github.com/slawek87/JobHunters/feedback"
	"github.com/slawek87/JobHunters/user"
	"github.com/slawek87/JobHunters/conf"
	"github.com/slawek87/JobHunters/linkedin"
)

func main() {
	conf.SessionInit()
	offer.MigrateDB()

	beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowOrigins: []string{"http://*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Only Authorized User can modify those sites.
	beego.InsertFilter("/api/v1/user", beego.BeforeRouter, user.OnlyAuthorizedUserCanModify)
	beego.InsertFilter("/api/v1/user/company", beego.BeforeRouter, user.OnlyAuthorizedAndActiveUserCanModify)
	beego.InsertFilter("/api/v1/offer/:offerID:string", beego.BeforeRouter, user.OnlyAuthorizedAndActiveUserCanModify)
	beego.InsertFilter("/api/v1/offer/:offerID:string/contribution", beego.BeforeRouter, user.OnlyAuthorizedAndActiveUserCanModify)

	// Only Authorized User can visits those sites.
	beego.InsertFilter("/api/v1/offer/:offerID:string/candidate", beego.BeforeRouter, user.OnlyAuthorizedUser)
	beego.InsertFilter("/api/v1/offer/:offerID:string/list", beego.BeforeRouter, user.OnlyAuthorizedUser)
	beego.InsertFilter("/api/v1/offer/:offerID:string/candidate/resume/:resumeID:string.pdf", beego.BeforeRouter, user.OnlyAuthorizedUser)
	beego.InsertFilter("/api/v1/offer/:offerID:string/candidate/:candidateID:string/feedback", beego.BeforeRouter, user.OnlyAuthorizedUser)
	beego.InsertFilter("/api/v1/offer/:offerID:string/candidate/:candidateID:string/feedback/:feedbackID:string", beego.BeforeRouter, user.OnlyAuthorizedUser)
	beego.InsertFilter("/api/v1/offer/:offerID:string/candidate/:candidateID:string", beego.BeforeRouter, user.OnlyAuthorizedUser)

	// Routers
	beego.Router("/api/v1/user/authorization/url", &linkedin.LinkedinView{}, "get:GetAuthorizationURL")
	beego.Router("/api/v1/user/login", &user.UserView{}, "get:Login")
	beego.Router("/api/v1/user", &user.UserView{}, "put:UpdateUser")
	beego.Router("/api/v1/user/company", &user.UserView{}, "put:UpdateUserCompany")
	beego.Router("/api/v1/offer", &offer.OfferView{}, "post:Post")
	beego.Router("/api/v1/offer/list", &offer.OfferView{}, "get:List")
	beego.Router("/api/v1/offer/:offerID:string", &offer.OfferView{}, "get:Get")
	beego.Router("/api/v1/offer/:offerID:string", &offer.OfferView{}, "put:Put")
	beego.Router("/api/v1/offer/:offerID:string", &offer.OfferView{}, "delete:Delete")
	beego.Router("/api/v1/offer/:offerID:string/contribution", &contribution.ContributionView{}, "post:Post")
	beego.Router("/api/v1/offer/:offerID:string/contribution/:contributionID:string", &contribution.ContributionView{}, "delete:Delete")
	beego.Router("/api/v1/offer/:offerID:string/candidate", &candidate.CandidateView{}, "post:Post")
	beego.Router("/api/v1/offer/:offerID:string/candidate/list", &candidate.CandidateView{}, "get:List")
	beego.Router("/api/v1/offer/:offerID:string/candidate/resume/:resumeID:string.pdf", &candidate.CandidateView{}, "get:DownloadResume")
	beego.Router("/api/v1/offer/:offerID:string/candidate/:candidateID:string/feedback", &feedback.FeedbackView{}, "post:SendFeedback")
	beego.Router("/api/v1/offer/:offerID:string/candidate/:candidateID:string/feedback/:feedbackID:string", &feedback.FeedbackView{}, "post:SendFeedback")
	beego.Router("/api/v1/offer/:offerID:string/candidate/:candidateID:string/feedback/:feedbackID:string", &feedback.FeedbackView{}, "get:ReceiveFeedback")
	beego.Router("/api/v1/offer/:offerID:string/candidate/:candidateID:string", &candidate.CandidateView{}, "get:Get")
	beego.Router("/api/v1/offer/:offerID:string/candidate/:candidateID:string", &candidate.CandidateView{}, "put:Put")
	beego.Router("/api/v1/offer/:offerID:string/candidate/:candidateID:string", &candidate.CandidateView{}, "delete:Delete")
	beego.Run("localhost:9000")
}
