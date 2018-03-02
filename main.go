package main

import (
	"github.com/astaxie/beego"
	"github.com/slawek87/JobHunters/offer"
	"github.com/slawek87/JobHunters/contribution"
)

func main() {
	beego.Router("/offer", &offer.OfferView{}, "post:Post")
	beego.Router("/offer/list", &offer.OfferView{}, "get:List")
	beego.Router("/offer/:offerID:string", &offer.OfferView{}, "get:Get")
	beego.Router("/offer/:offerID:string", &offer.OfferView{}, "put:Put")
	beego.Router("/offer/:offerID:string", &offer.OfferView{}, "delete:Delete")
	beego.Router("/offer/:offerID:string/contribution", &contribution.ContributionView{}, "post:Post")
	beego.Router("/offer/:offerID:string/contribution/:contributionID:string", &contribution.ContributionView{}, "delete:Delete")
	beego.Run("localhost:8000")
}
