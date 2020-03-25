package models

import (
	"fmt"
	"gohttpserver/dbcontrollers"
	"time"
)

type Order struct {
	Id           string    `xorm:"not null VARCHAR(36)"`
	OrderId      string    `xorm:"VARCHAR(255)"`
	UserId       string    `xorm:"VARCHAR(36)"`
	RealizeId    string    `xorm:"VARCHAR(36)"`
	CreateTime   time.Time `xorm:"DATETIME"`
	ValidityTime string    `xorm:"VARCHAR(36)"`
}

func (p *Order) InsertOrderSQL() bool {
	rows, err := dbcontrollers.GetOrm().Insert(p)
	dbcontrollers.PanicIf(err, "fail to InsertOrderSQL")
	if err != nil {
		return false
	}
	if rows == 0 {
		return false
	}
	return true
}

func (p *Order) GetMaxOrderId() (string, bool) {
	gsql := `SELECT MAX("OrderId") as "OrderId" from "AK_Order"`
	gres, gerr := dbcontrollers.GetOrm().Query(gsql)
	dbcontrollers.PanicIf(gerr, "fail to GetMaxOrderId")
	var retStr string
	if len(gres) > 0 {
		retStr = string(gres[0]["OrderId"])
		fmt.Println(retStr, "-------------", len(gres))
	} else {
		retStr = ""
	}
	if gerr != nil {
		return retStr, false
	}
	return retStr, true
}
