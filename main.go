package main

import (
	"github.com/astaxie/beego"
	"github.com/slawek87/JobHunters/offer"
)

func main() {
	beego.Router("/offer", &offer.OfferView{}, "post:Post")
	beego.Router("/offer/list", &offer.OfferView{}, "get:List")
	beego.Router("/offer/:offerID:string", &offer.OfferView{}, "put:Put")
	beego.Run("localhost:8000")
}
