package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func (m *Admin) DeleteAllAdmins() {

	var admins []Admin

	if _, err := m.Query().Limit(-1).All(&admins); err != nil {
		beego.Error("Error quering admin users: ", err)
	} else {
		for _, admin := range admins {
			if err := admin.Delete(); err != nil {
				beego.Error("Error deleting user: ", err)
			}
		}
	}
}

func (m *Admin) InsertNewAdmin() {

	if err := m.Insert(); err != nil {
		beego.Error("Error inserting admin user: ", err)
	} else {
		beego.Debug("OK. Admin user created successfully.")
	}

}

//// Orm

type Admin struct {
	Id       int64  `orm:"column(id);pk;auto"`
	Email    string `orm:"column(email);size(255)"`
	Password string `orm:"column(password)"`
}

func init() {
	orm.RegisterModel(new(Admin))
}

func (m *Admin) TableName() string {
	return "admin"
}

func (m *Admin) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Admin) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Admin) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Admin) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m).RelatedSel()
}
