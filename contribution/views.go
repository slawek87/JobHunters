package contribution

import (
	"github.com/astaxie/beego"
	"github.com/slawek87/JobHunters/user"
)

type ContributionView struct {
	beego.Controller
	ContributionController ContributionController
}

func (view *ContributionView) Post() {
	userSession := view.GetSession("User")
	results := make(map[string]interface{})
	view.ParseForm(&view.ContributionController.Contribution)
	view.ContributionController.SetOfferID(view.Ctx.Input.Param(":offerID"))
	view.ContributionController.SetUserID(userSession.(*user.User).UserID)

	err := view.ContributionController.Create()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = view.ContributionController.GetContribution()
		view.Data["json"] = results
		view.ServeJSON()
	}
}

func (view *ContributionView) Delete() {
	userSession := view.GetSession("User")
	results := make(map[string]interface{})

	view.ContributionController.SetContributionID(view.Ctx.Input.Param(":contributionID"))
	view.ContributionController.SetOfferID(view.Ctx.Input.Param(":offerID"))
	view.ContributionController.SetUserID(userSession.(*user.User).UserID)

	err := view.ContributionController.Delete()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = "Contribution has been deleted."
		view.Data["json"] = results
		view.ServeJSON()
	}
}
