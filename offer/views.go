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
	//getSession := v.StartSession()
	view.ParseForm(&view.OfferController.Offer)
	view.OfferController.SetUserID("Xyz123") //getSession.Get("ID").(string)

	err := view.OfferController.Create()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		view.Data["json"] = view.OfferController.GetOffer()
		view.ServeJSON()
	}
}

func (view *OfferView) Put() {
	offerID := view.Ctx.Input.Param(":offerID")

	view.ParseForm(&view.OfferController.Offer)
	view.OfferController.SetOfferID(offerID)
	view.OfferController.SetUserID("Xyz123") //getSession.Get("ID").(string)

	err := view.OfferController.Update()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		view.Data["json"] = view.OfferController.GetOffer()
		view.ServeJSON()
	}
}

func (view *OfferView) List() {
	offers, err := view.OfferController.All(nil)

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		view.Data["json"] = offers
		view.ServeJSON()
	}
}