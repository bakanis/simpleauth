package adminauth

import (
	"github.com/bakanis/simpleauth/controllers"

	_ "github.com/bakanis/simpleauth/filters"

	"github.com/astaxie/beego"
)

func init() {
	beego.Debug("Initializing auth module")
	beego.SessionOn = true
	beego.EnableXSRF = true
	controllers.InitializeModule()
}
