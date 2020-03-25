package controllers

import (
	"encoding/json"
	"fmt"

	// "gohttpserver/dbcontrollers"
	"gohttpserver/models"

	"gopkg.in/macaron.v1"
)

type LogInInfo struct {
	UserName string
	Passwd   string
}

func LogInFun(ctx *macaron.Context) {

	jsonStr := ctx.Query("LogIn")
	logInInfo := &LogInInfo{}
	err := json.Unmarshal([]byte(jsonStr), logInInfo)
	fmt.Println("------login  sadsad -----", err, logInInfo.UserName, logInInfo.Passwd)
	var ret bool = false
	if err == nil {
		userInfo := &models.UserInfo{}
		userInfos := make([]*models.UserInfo, 0)
		userInfos = userInfo.GetUserInfoSQL("")
		for _, v := range userInfos {
			if v.UserName == logInInfo.UserName && v.UserPasswd == logInInfo.Passwd {
				fmt.Println("--------has -----", true)
				ret = true
			}
		}
	}

	if err == nil {
		ret = true
	}

	fmt.Println("--------ret-----", ret)

	ctx.JSON(200, &ContextResult{
		Ok:   ret,
		Data: "LogIn",
	})
}
