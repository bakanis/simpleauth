package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var Cfg = beego.AppConfig

func init() {

	var FilterAdmin = func(ctx *context.Context) {
		adminEmail := Cfg.String("admin_Email")

		if sessionMap, ok := ctx.Input.Session("admin").(map[string]interface{}); !ok {
			beego.Debug("Session: ", sessionMap)

			if sessionMap["adminEmail"] != adminEmail {
				ctx.Redirect(302, "/login")
			}
		}
	}

	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterAdmin)
	beego.InsertFilter("/admin", beego.BeforeRouter, FilterAdmin)
}
