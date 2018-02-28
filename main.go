package main

import (
	"github.com/astaxie/beego"
	"github.com/slawek87/JobHunters/offer"
)

func main() {
	beego.Router("/offer/", &offer.OfferView{}, "post:Post")
	beego.Router("/offer/all/", &offer.OfferView{}, "get:All")
	beego.Run("localhost:8000")
}
