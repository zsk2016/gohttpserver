package dbcontrollers

import (
	"fmt"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/lunny/log"
)

var orm *xorm.Engine

func GetOrm() *xorm.Engine {
	if orm == nil {
		InitXorm()
	}
	return orm
}

func PanicIf(err error, format string, v ...interface{}) {
	if err != nil {
		log.Errorf(format+" [ %v ]", v, err)
		panic(err)
	}
}

type UserInfo struct {
	Id         string    `json:"Id" xorm:"varchar(36)"`
	UserId     string    `json:"UserId" xorm:"varchar(36)"`
	UserName   string    `json:"UserName" xorm:"varchar(255)"`
	UserPhone  string    `json:"UserPhone" xorm:"varchar(20)"`
	UserEmail  string    `json:"UserEmail" xorm:"varchar(255)"`
	UserType   string    `json:"UserType" xorm:"int"`
	CreateTime time.Time `json:"BDate" xorm:"datetime"`
	UpdateTime time.Time `json:"BDate" xorm:"datetime"`
}

func (p *UserInfo) GetUserInfoSQL(querystring string, args ...interface{}) []*UserInfo {
	userInfos := make([]*UserInfo, 0)
	err := GetOrm().Where(querystring, args...).Find(&userInfos)
	PanicIf(err, "fail to GetUserInfoSQL")
	return userInfos
}

func InitXorm() {
	var err error
	PrefixMapper := "AE_"
	dbc := "user=%s password=%s port=%s dbname=%s host=%s sslmode=disable"
	connString := fmt.Sprintf(dbc, "postgres", "1982911zsk", "5432", "database_userinfo", "localhost")
	dbdriver := "postgres"
	fmt.Println(connString)
	orm, err = xorm.NewEngine(dbdriver, connString)
	orm.TZLocation = time.Local
	orm.ShowSQL(true)
	log.Info(connString)
	tbMapper := core.NewPrefixMapper(core.SameMapper{}, PrefixMapper)
	orm.SetTableMapper(tbMapper)
	orm.SetColumnMapper(core.SameMapper{})
	orm.SetMaxIdleConns(100)
	orm.SetMaxOpenConns(780)
	fmt.Println(err)
}
