package offer

import (
	"github.com/astaxie/beego"
	"github.com/rs/xid"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type OfferView struct {
	beego.Controller
	//Session session.Store
}

func (v *OfferView) Post() {
	//getSession := v.StartSession()
	getUniqueID := xid.New()
	offer := Offer{}

	v.ParseForm(&offer)

	offer.OfferID = getUniqueID.String()
	offer.CreatedAt = time.Now()
	offer.UpdatedAt = time.Now()
	offer.ExpirationTime = time.Now().AddDate(0, 0, EXPIRATION_TIME_DAYS)
	offer.UserID = "Xyz123"//getSession.Get("ID").(string)

	controller := OfferController{Offer: offer}

	controller.Create()

	v.Data["json"] = controller.Offer
	v.ServeJSON()
}

func (v *OfferView) Put() {
	var offer Offer

	offerID := v.Ctx.Input.Param(":offerID")
	v.ParseForm(&offer)

	offer.OfferID = offerID

	controller := OfferController{Offer: offer}

	controller.Update()

}

func (v *OfferView) All() {
	controller := OfferController{}

	offers, _ := controller.All(bson.M{"user_id": "Xyz123"})

	v.Data["json"] = offers
	v.ServeJSON()
}