package models

import (
	"time"
)

type AkPrice struct {
	Id            string    `xorm:"not null pk VARCHAR(36)"`
	Priceid       string    `xorm:"VARCHAR(36)"`
	Pricevalue    string    `xorm:"VARCHAR(36)"`
	Pricediscount string    `xorm:"VARCHAR(36)"`
	Createtime    time.Time `xorm:"DATETIME"`
}
