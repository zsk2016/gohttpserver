package models

import (
	"time"
)

type AkRealize struct {
	Id            string    `xorm:"not null VARCHAR(36)"`
	Realizeid     string    `xorm:"VARCHAR(36)"`
	Priceid       string    `xorm:"VARCHAR(36)"`
	Realizetype   string    `xorm:"VARCHAR(36)"`
	Realizename   string    `xorm:"VARCHAR(255)"`
	Realizedetail string    `xorm:"VARCHAR(255)"`
	Createtime    time.Time `xorm:"DATETIME"`
}
