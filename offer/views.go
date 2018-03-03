package offer

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
	"fmt"
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

	var description string
	var currency string
	var commissionGte int
	var commissionLte int
	var lastDays int
	var commissionQuery bson.M

	view.Ctx.Input.Bind(&description, "description")
	view.Ctx.Input.Bind(&currency, "currency")
	view.Ctx.Input.Bind(&commissionGte, "commission_gte")
	view.Ctx.Input.Bind(&commissionLte, "commission_lte")
	view.Ctx.Input.Bind(&commissionLte, "commission_lte")
	view.Ctx.Input.Bind(&lastDays, "last_days")

	if commissionLte > 0 && commissionGte > 0 {
		commissionQuery = bson.M{"$lte": commissionLte, "$gte": commissionGte}
	} else if commissionGte == 0 && commissionLte > 0 {
		commissionQuery = bson.M{"$lte": commissionLte}
	} else if commissionLte == 0 && commissionGte > 0 {
		commissionQuery = bson.M{"$gte": commissionGte}
	} else {
		commissionQuery = bson.M{"$gte": 0}
	}

	if currency == "" {
		currency = "EUR"
	} else {
		currency = strings.ToUpper(currency)
	}

	createdAt := time.Now().Add(-30 * 24 * time.Hour)

	switch lastDays {
	case 7:
		createdAt = time.Now().Add(-7 * 24 * time.Hour)
	case 14:
		createdAt = time.Now().Add(-14 * 24 * time.Hour)
	case 30:
		createdAt = time.Now().Add(-30 * 24 * time.Hour)
	}

	offers, err := view.OfferController.Find(
		bson.M{
			"description": bson.RegEx{Pattern: description, Options: "i"},
			"currency": currency,
			"commission": commissionQuery,
			"created_at": bson.M{"$gte": createdAt}})

	if err != nil {
		view.CustomAbort(300, err.Error())
	} else {
		results["results"] = offers
		view.Data["json"] = results
		view.ServeJSON()
	}
}