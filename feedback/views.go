package feedback

import "github.com/astaxie/beego"

type FeedbackView struct {
	beego.Controller
	MainController MainController
	//Session session.Store
}

func (view *FeedbackView) SendFeedback() {
	results := make(map[string]interface{})

	view.ParseForm(&view.MainController.FeedbackController.Feedback)
	view.MainController.FeedbackController.SetFeedbackID(view.Ctx.Input.Param(":FeedbackID"))
	view.MainController.FeedbackController.SetCandidateID(view.Ctx.Input.Param(":candidateID"))
	view.MainController.FeedbackController.SetOfferID(view.Ctx.Input.Param(":offerID"))

	if view.MainController.FeedbackController.Feedback.FeedbackID == "" {
		err := view.MainController.CreateFeedback()
		if err != nil {
			view.CustomAbort(300, err.Error())
		}
	}

	view.ParseForm(&view.MainController.MessageController.Message)
	view.MainController.MessageController.SetSenderID("xyz")
	view.MainController.MessageController.SetSenderFullName("Sławek Ka.")

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
	results := make(map[string]interface{})

	view.ParseForm(&view.MainController.FeedbackController.Feedback)
	view.MainController.FeedbackController.SetFeedbackID(view.Ctx.Input.Param(":FeedbackID"))
	view.MainController.FeedbackController.SetCandidateID(view.Ctx.Input.Param(":candidateID"))
	view.MainController.FeedbackController.SetOfferID(view.Ctx.Input.Param(":offerID"))

	err := view.MainController.Get()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = view.MainController.FeedbackController.Feedback
		view.Data["json"] = results
		view.ServeJSON()
	}
}