package feedback

import (
	"github.com/astaxie/beego"
	"github.com/slawek87/JobHunters/user"
)

type FeedbackView struct {
	beego.Controller
	MainController MainController
}

func (view *FeedbackView) SendFeedback() {
	userSession := view.GetSession("User")
	results := make(map[string]interface{})

	view.ParseForm(&view.MainController.FeedbackController.Feedback)
	view.MainController.FeedbackController.SetFeedbackID(view.Ctx.Input.Param(":feedbackID"))
	view.MainController.FeedbackController.SetCandidateID(view.Ctx.Input.Param(":candidateID"))
	view.MainController.FeedbackController.SetOfferID(view.Ctx.Input.Param(":offerID"))

	if view.MainController.FeedbackController.Feedback.FeedbackID == "" {
		err := view.MainController.CreateFeedback()
		if err != nil {
			view.CustomAbort(300, err.Error())
		}
	}

	view.ParseForm(&view.MainController.MessageController.Message)
	view.MainController.MessageController.SetSenderID(userSession.(*user.User).UserID)

	fullName := userSession.(*user.User).FirstName + " " + userSession.(*user.User).LastName
	view.MainController.MessageController.SetSenderFullName(fullName)

	err := view.MainController.CreateMessage()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = view.MainController.MessageController.Message
		view.Data["json"] = results
		view.ServeJSON()
	}
}

func (view *FeedbackView) ReceiveFeedback() {
	userSession := view.GetSession("User")
	results := make(map[string]interface{})

	view.ParseForm(&view.MainController.FeedbackController.Feedback)
	view.MainController.FeedbackController.SetFeedbackID(view.Ctx.Input.Param(":feedbackID"))
	view.MainController.FeedbackController.SetCandidateID(view.Ctx.Input.Param(":candidateID"))
	view.MainController.FeedbackController.SetOfferID(view.Ctx.Input.Param(":offerID"))

	err := view.MainController.Get()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		view.MainController.SetRead(userSession.(*user.User).UserID)
		results["results"] = view.MainController.FeedbackController.Feedback
		view.Data["json"] = results
		view.ServeJSON()
	}
}