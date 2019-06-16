package models

import (
	"time"
)

type AkOrder struct {
	Id           string    `xorm:"not null VARCHAR(36)"`
	Orderid      string    `xorm:"VARCHAR(255)"`
	Userid       string    `xorm:"VARCHAR(36)"`
	Realizeid    string    `xorm:"VARCHAR(36)"`
	Createtime   time.Time `xorm:"DATETIME"`
	Validitytime string    `xorm:"VARCHAR(36)"`
}
