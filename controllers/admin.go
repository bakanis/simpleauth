package controllers

import (
	"github.com/bakanis/simpleauth/models"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"

	"html/template"
	"time"
)

var Cfg = beego.AppConfig

type AdminController struct {
	beego.Controller
	i18n.Locale
}

func InitializeModule() {
	adminEmail := Cfg.String("admin_Email")
	adminPassword := Cfg.String("admin_password")

	if adminEmail == "" {
		beego.Error("No admin email found in application config file")
	} else {
		beego.Debug("Admin email found in application config file")
	}

	if adminPassword == "" {
		beego.Error("No admin password found in application config file")
	} else {
		beego.Debug("Admin password found in application config file")
	}

	p := []byte(adminPassword)

	h, err := bcrypt.GenerateFromPassword(p, 10)
	if err != nil {
		beego.Error("Error: ", err)
		return
	}

	adminUser := models.Admin{Email: adminEmail}
	if err := adminUser.Read("email"); err != nil && err != orm.ErrNoRows {
		beego.Error("Error: ", err)
	} else if err == orm.ErrNoRows {
		beego.Debug("No admin user found. Creating...")
		adminUser.DeleteAllAdmins()
		adminUser.Password = string(h)
		adminUser.InsertNewAdmin()
	}
}

func (c *AdminController) LoginDo() {

	flash := beego.NewFlash()
	email := c.GetString("username")
	password := c.GetString("password")

	valid := validation.Validation{}
	valid.Email(email, "email")
	valid.Required(password, "password")
	if valid.HasErrors() {
		beego.Debug(valid.ErrorsMap)
		errormap := make(map[string]string)
		for _, err := range valid.Errors {
			errormap[err.Key] = err.Message
		}
		c.Data["Errors"] = errormap
		flash.Error("Validation errors")
		flash.Store(&c.Controller)
		beego.Error("Validation errors")
		c.Redirect("/login", 302)
		return
	}
	beego.Debug("Authorization is", email, ":", password)

	admin := models.Admin{Email: email}
	if err := admin.Read("email"); err != nil {
		beego.Error("Error: ", err)
		flash.Error("Wrong username or password")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	}

	h := []byte(admin.Password)
	p := []byte(password)

	if err := bcrypt.CompareHashAndPassword(h, p); err != nil {
		beego.Warn("Error: ", err)
		flash.Error("Wrong username or password")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
		return
	} else {
		//******** Create session and go back to previous page
		sessionMap := make(map[string]interface{})
		sessionMap["adminId"] = admin.Id
		sessionMap["adminEmail"] = email
		sessionMap["timestamp"] = time.Now()
		c.SetSession("admin", sessionMap)
		c.Redirect("/admin/home", 302)
		return
	}
	flash.Error("Wrong username or password")
	flash.Store(&c.Controller)
	c.Redirect("/login", 302)
}

func (c *AdminController) LoginShow() {
	beego.XSRFKEY = Cfg.String("xsrf_key")
	beego.XSRFExpire, _ = Cfg.Int("xsrf_expire")
	c.Data["xsrfdata"] = template.HTML(c.XsrfFormHtml())
	c.TplNames = "simpleauth/login.tpl"
}

func (c *AdminController) LogoutDo() {
	c.DelSession("admin")
	c.Redirect("/", 302)
}
