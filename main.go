package main

import (
	"github.com/astaxie/beego"
	"github.com/slawek87/JobHunters/offer"
	"github.com/slawek87/JobHunters/contribution"
	"github.com/slawek87/JobHunters/candidate"
	"github.com/slawek87/JobHunters/feedback"
	"fmt"
	"github.com/slawek87/JobHunters/linkedin"
	"github.com/slawek87/JobHunters/user"
)

func main() {
	offer.MigrateDB()

	auth := linkedin.Authorization {
		AuthorizationEndpoint: linkedin.AUTHORIZATION_ENDPOINT,
		AccessTokenEndpoint: linkedin.ACCESS_TOKEN_ENDPOINT,
		Method: linkedin.GET,
		GrantType: linkedin.GRANT_TYPE,
		ResponseType: linkedin.RESPONSE_TYPE,
		Scope: linkedin.SCOPE,
		RedirectURI: linkedin.REDIRECT_URI,
		ClientID: linkedin.CLIENT_ID,
		ClientSecret: linkedin.CLIENT_SECRET,
		State: linkedin.STATE,
	}

	// Generates Authorization Url. You have to visit it to authorized.
	fmt.Println(auth.GetAuthorizationURL())

	beego.Router("/user/login", &user.UserView{}, "get:Login")
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
	beego.Router("/offer/:offerID:string/candidate/:candidateID:string", &candidate.CandidateView{}, "get:Get")
	beego.Router("/offer/:offerID:string/candidate/:candidateID:string", &candidate.CandidateView{}, "put:Put")
	beego.Router("/offer/:offerID:string/candidate/:candidateID:string", &candidate.CandidateView{}, "delete:Delete")
	beego.Run("localhost:8000")
}
