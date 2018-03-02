package contribution

import "github.com/astaxie/beego"

type ContributionView struct {
	beego.Controller
	ContributionController    ContributionController
	//Session session.Store
}

func (view *ContributionView) Post() {
	results := make(map[string]interface{})
	//getSession := v.StartSession()
	view.ParseForm(&view.ContributionController.Contribution)
	view.ContributionController.SetOfferID(view.Ctx.Input.Param(":offerID"))
	view.ContributionController.SetUserID("Xyz123") //getSession.Get("ID").(string)

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
	results := make(map[string]interface{})

	view.ContributionController.SetContributionID(view.Ctx.Input.Param(":contributionID"))
	view.ContributionController.SetOfferID(view.Ctx.Input.Param(":offerID"))
	view.ContributionController.SetUserID("Xyz123") //getSession.Get("ID").(string)

	err := view.ContributionController.Delete()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = "Contribution has been deleted."
		view.Data["json"] = results
		view.ServeJSON()
	}
}
