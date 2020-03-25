package models

import (
	"gohttpserver/dbcontrollers"
)

type RealizeAndPrice struct {
	Id    string `xorm:"VARCHAR(36)"`
	Name  string `xorm:"VARCHAR(36)"`
	Value string `xorm:"VARCHAR(36)"`
}

func (p *RealizeAndPrice) GetRealizeAndPriceSQL(querystring string, args ...interface{}) []*RealizeAndPrice {
	raps := make([]*RealizeAndPrice, 0)
	//err := dbcontrollers.GetOrm().Where(querystring, args...).Find(&raps)
	// dbcontrollers.PanicIf(err, "fail to GetRealizeAndPriceSQL")

	sqlll := `SELECT *from  "AK_RealizeAndPrice"`
	dataMap, _ := dbcontrollers.GetOrm().QueryString(sqlll)
	for _, v := range dataMap {
		rap := &RealizeAndPrice{}
		rap.Id = v["id"]
		rap.Name = v["name"]
		rap.Value = v["value"]
		raps = append(raps, rap)
	}
	return raps
}
