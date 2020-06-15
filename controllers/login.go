package controllers

import (
	"encoding/json"
	"fmt"

	"gohttpserver/dbcontrollers"
	"gohttpserver/models"

	"gopkg.in/macaron.v1"
)

type LogInInfo struct {
	UserName string
	Passwd   string
	CpuId    string
}

func changeCpuId(cpuId, userName string) {
	fmt.Println("---changeCpuId---")
	sqlUpdate := `UPDATE "AK_UserInfo" SET "CpuId" = ? WHERE "AK_UserInfo"."UserName" = ?`
	res, err := dbcontrollers.GetOrm().Exec(sqlUpdate, cpuId, userName)
	fmt.Println(res, err)
}

func LogInFun(ctx *macaron.Context) {
	//ctx.Resp.Header().Set("Access-Control-Allow-Origin", "*")
	// ctx.Resp.Header().Set("Access-Control-Allow-Headers", "content-type,token")
	// ctx.Resp.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	jsonStr := ctx.Query("LogIn")
	logInInfo := &LogInInfo{}
	err := json.Unmarshal([]byte(jsonStr), logInInfo)
	fmt.Println("------login  sadsad -----", err, logInInfo.UserName, logInInfo.Passwd)
	var ret bool = false
	userInfo := &models.UserInfo{}
	if err == nil {
		userInfos := make([]*models.UserInfo, 0)
		userInfos = userInfo.GetUserInfoSQL("")
		for _, v := range userInfos {
			if v.UserName == logInInfo.UserName && v.UserPasswd == logInInfo.Passwd {
				fmt.Println("--------has -----", true)
				ret = true
				userInfo.UserId = v.UserId
				userInfo.UserName = v.UserName
				userInfo.CpuId = v.CpuId
			}
		}
		changeCpuId(logInInfo.CpuId, logInInfo.UserName)
	}

	fmt.Println("--------ret-----", userInfo)

	ctx.JSON(200, &ContextResult{
		Ok:    ret,
		Data:  "LogIn",
		Value: &userInfo,
	})
}
