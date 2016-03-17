package routers

import (
	"OxO/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &cat.SFO{})
    beego.Router("/s/check", &cat.URLChecker{})
    beego.Router("/s/gen", &cat.URLShortener{})
    beego.Router("/:shorten("+beego.AppConfig.String("minlength")+",})", &cat.URLWizard{})
}
