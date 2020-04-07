package controllers

import (
	"fmt"
	// "gohttpserver/dbcontrollers"
	"encoding/json"
	"gohttpserver/dbcontrollers"
	"gohttpserver/models"

	"gopkg.in/macaron.v1"
)

type ContextResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Value   interface{} `json:"value"`
	Ok      bool
	Err     interface{}
}

type UserNameInfo struct {
	UserName string
}

func GetData(ctx *macaron.Context) {
	fmt.Println("-------getuserinfo------")
	userInfos := make([]*models.UserInfo, 0)
	userInfo := &models.UserInfo{}
	userInfos = userInfo.GetUserInfoSQL("")
	fmt.Println("-------------", userInfos[0])
	ctx.JSON(200, &ContextResult{
		Ok:   true,
		Data: &userInfos,
	})
}

func GetUserCpuId(ctx *macaron.Context) {
	fmt.Println("-------GetUserCpuId------")
	jsonStr := ctx.Query("GetCpuId")
	uni := &UserNameInfo{}
	err := json.Unmarshal([]byte(jsonStr), uni)
	ret := false
	retValue := ""
	if err == nil {
		sqlll := `SELECT "AK_UserInfo"."CpuId" from "AK_UserInfo" WHERE "AK_UserInfo"."UserName" = ?`
		dataMap, _ := dbcontrollers.GetOrm().QueryString(sqlll, uni.UserName)
		if len(dataMap) == 1 {
			ret = true
			retValue = dataMap[0]["CpuId"]
		}
	}
	ctx.JSON(200, &ContextResult{
		Ok:    ret,
		Data:  "GetCpuId",
		Value: retValue,
	})
}

func GetUserIdByUserName(un string) string {
	fmt.Println("-------getuserid------")
	userInfos := make([]*models.UserInfo, 0)
	userInfo := &models.UserInfo{}
	whereStr := `"UserName" = ?`
	userInfos = userInfo.GetUserInfoSQL(whereStr, un)
	var reuserid = "0"
	if len(userInfos) >= 1 {
		reuserid = userInfos[0].UserId
		fmt.Println(reuserid)
	}
	return reuserid
}

func GetRealize(ctx *macaron.Context) {
	raps := make([]*models.RealizeAndPrice, 0)
	rap := &models.RealizeAndPrice{}
	raps = rap.GetRealizeAndPriceSQL("")
	fmt.Println(len(raps), raps[0])
	ctx.JSON(200, &ContextResult{
		Ok:    true,
		Data:  "GetRealize",
		Value: &raps,
	})
}
