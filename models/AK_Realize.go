package models

import (
	"gohttpserver/dbcontrollers"
	"time"
)

type Realize struct {
	Id            string    `xorm:"not null VARCHAR(36)"`
	RealizeId     string    `xorm:"VARCHAR(36)"`
	PriceId       string    `xorm:"VARCHAR(36)"`
	RealizeType   string    `xorm:"VARCHAR(36)"`
	RealizeName   string    `xorm:"VARCHAR(255)"`
	RealizeDetail string    `xorm:"VARCHAR(255)"`
	CreateTime    time.Time `xorm:"DATETIME"`
}

func (p *Realize) GetRealizeSQL(querystring string, args ...interface{}) []*Realize {
	Realizes := make([]*Realize, 0)
	err := dbcontrollers.GetOrm().Where(querystring, args...).Find(&Realizes)
	dbcontrollers.PanicIf(err, "fail to GetRealizesSQL")
	return Realizes
}
