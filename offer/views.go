package offer

import (
	"github.com/astaxie/beego"
	"github.com/rs/xid"
	"time"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

type OfferView struct {
	beego.Controller
	OfferController    OfferController
	//Session session.Store
}

func (view *OfferView) Post() {
	//getSession := v.StartSession()
	getUniqueID := xid.New()
	offer := Offer{}

	view.ParseForm(&offer)

	offer.OfferID = getUniqueID.String()
	offer.CreatedAt = time.Now()
	offer.UpdatedAt = time.Now()
	offer.ExpirationTime = time.Now().AddDate(0, 0, EXPIRATION_TIME_DAYS)
	offer.UserID = "Xyz123"//getSession.Get("ID").(string)

	view.OfferController.SetOffer(offer)

	err := view.OfferController.Create()

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		view.Data["json"] = view.OfferController.GetOffer()
		view.ServeJSON()
	}
}

func (view *OfferView) Put() {
	var offer Offer

	offerID := view.Ctx.Input.Param(":offerID")
	view.ParseForm(&offer)

	offer.OfferID = offerID

	view.OfferController.SetOffer(offer)

	query, _ := json.Marshal(offer)
	err := view.OfferController.Update(query)

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		view.Data["json"] = view.OfferController.GetOffer()
		view.ServeJSON()
	}
}

func (view *OfferView) List() {
	offers, err := view.OfferController.All(bson.M{"user_id": "Xyz123"})

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		view.Data["json"] = offers
		view.ServeJSON()
	}
}