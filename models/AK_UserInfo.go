package models

import (
	"fmt"
	"gohttpserver/dbcontrollers"
	"time"
)

type UserInfo struct {
	Id          string    `xorm:"not null VARCHAR(36)"`
	UserId      string    `xorm:"VARCHAR(36)"`
	UserName    string    `xorm:"VARCHAR(255)"`
	UserPasswd  string    `xorm:"VARCHAR(36)"`
	UserPhone   string    `xorm:"VARCHAR(36)"`
	UserEmail   string    `xorm:"VARCHAR(255)"`
	UserCompany string    `xorm:"VARCHAR(255)"`
	UserType    string    `xorm:"VARCHAR(36)"`
	CreateTime  time.Time `xorm:"DATETIME"`
	UpdateTime  time.Time `xorm:"DATETIME"`
}

func (p *UserInfo) GetUserInfoSQL(querystring string, args ...interface{}) []*UserInfo {
	userInfos := make([]*UserInfo, 0)
	err := dbcontrollers.GetOrm().Where(querystring, args...).Find(&userInfos)
	dbcontrollers.PanicIf(err, "fail to GetUserInfoSQL")
	return userInfos
}

func (p *UserInfo) InsertUserInfoSQL() bool {
	rows, err := dbcontrollers.GetOrm().Insert(p)
	dbcontrollers.PanicIf(err, "fail to InsertUserInfoSQL")
	if err != nil {
		return false
	}
	if rows == 0 {
		return false
	}
	return true
}

func (p *UserInfo) GetMaxUserId() (string, bool) {
	gsql := `SELECT MAX("UserId") as "UserId" from "AK_UserInfo"`
	gres, gerr := dbcontrollers.GetOrm().Query(gsql)
	dbcontrollers.PanicIf(gerr, "fail to GetMaxUserId")
	var retStr string
	if len(gres) > 0 {
		retStr = string(gres[0]["UserId"])
		fmt.Println(retStr, "-------------", len(gres))
	} else {
		retStr = ""
	}
	if gerr != nil {
		return retStr, false
	}
	return retStr, true
}
