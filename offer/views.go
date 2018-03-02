package offer

import (
	"github.com/astaxie/beego"
)

type OfferView struct {
	beego.Controller
	OfferController    OfferController
	//Session session.Store
}

func (view *OfferView) Post() {
	results := make(map[string]interface{})
	//getSession := v.StartSession()
	view.ParseForm(&view.OfferController.Offer)
	view.OfferController.SetUserID("Xyz123") //getSession.Get("ID").(string)

	err := view.OfferController.Create()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = view.OfferController.GetOffer()
		view.Data["json"] = results
		view.ServeJSON()
	}
}

func (view *OfferView) Get() {
	results := make(map[string]interface{})

	view.OfferController.SetOfferID(view.Ctx.Input.Param(":offerID"))

	offer, err := view.OfferController.Get()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = offer
		view.Data["json"] = results
		view.ServeJSON()
	}
}

func (view *OfferView) Delete() {
	results := make(map[string]interface{})

	view.OfferController.SetOfferID(view.Ctx.Input.Param(":offerID"))
	view.OfferController.SetUserID("Xyz123") //getSession.Get("ID").(string)

	err := view.OfferController.Delete()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = "Offer has been deleted."
		view.Data["json"] = results
		view.ServeJSON()
	}
}


func (view *OfferView) Put() {
	results := make(map[string]interface{})
	offerID := view.Ctx.Input.Param(":offerID")

	view.ParseForm(&view.OfferController.Offer)
	view.OfferController.SetOfferID(offerID)
	view.OfferController.SetUserID("Xyz123") //getSession.Get("ID").(string)

	err := view.OfferController.Update()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = view.OfferController.GetOffer()
		view.Data["json"] = results
		view.ServeJSON()
	}
}

func (view *OfferView) List() {
	results := make(map[string]interface{})
	offers, err := view.OfferController.All(nil)

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = offers
		view.Data["json"] = results
		view.ServeJSON()
	}
}