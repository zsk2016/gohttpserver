package controllers

import (
	"fmt"
	"gohttpserver/dbcontrollers"

	"gopkg.in/macaron.v1"
)

type ContextResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Ok      bool
	Err     interface{}
}

func GetData(ctx *macaron.Context) {
	fmt.Println("-------123------")
	userInfos := make([]*dbcontrollers.UserInfo, 0)
	userInfo := &dbcontrollers.UserInfo{}
	userInfos = userInfo.GetUserInfoSQL("")
	fmt.Println("-------------", userInfos[0])
	ctx.JSON(200, &ContextResult{
		Ok:   true,
		Data: &userInfos,
	})
}
