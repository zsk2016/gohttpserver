package controllers

import (
	"encoding/json"
	"fmt"
	"gohttpserver/tools"
	"time"

	"gohttpserver/models"

	"github.com/satori/go.uuid"
	"gopkg.in/macaron.v1"
)

type SignInInfo struct {
	UserName  string
	EmailAddr string
	Passwd    string
	Company   string
	Phone     string
}

func SignInFun(ctx *macaron.Context) {
	fmt.Println("------sign in-----")
	jsonStr := ctx.Query("SignIn")
	signInInfo := &SignInInfo{}
	err := json.Unmarshal([]byte(jsonStr), signInInfo)
	if err == nil {
		uuid, _ := uuid.NewV4()
		fmt.Println(signInInfo.UserName, signInInfo.EmailAddr)
		userInfo := models.UserInfo{}

		UId, flag := userInfo.GetMaxUserId()
		if flag == true {
			fmt.Println(UId)
		}

		userInfo.Id = uuid.String()
		userInfo.UserName = signInInfo.UserName
		userInfo.UserEmail = signInInfo.EmailAddr
		userInfo.UserPasswd = signInInfo.Passwd
		userInfo.UserCompany = signInInfo.Company
		userInfo.UserPhone = signInInfo.Phone
		userInfo.UserType = "0"
		userInfo.CreateTime = time.Now()
		userInfo.UpdateTime = time.Now()

		ret := userInfo.InsertUserInfoSQL()

		if ret == true {
			ei := &tools.EmailInfo{}
			ei.FromAddr = "2388662767@qq.com"
			ei.SendAddr = userInfo.UserEmail

			ei.SendInfo = "欢迎注册 树pdf, 账号注册成功! \t\n账号：" + userInfo.UserName +
				"密码：" + userInfo.UserPasswd
			isok := tools.SendEmail(ei)
			fmt.Println(isok)
		}

		ctx.JSON(200, &ContextResult{
			Ok:   ret,
			Data: "SignIn",
		})
	} else {
		fmt.Println(jsonStr)
		ctx.JSON(200, &ContextResult{
			Ok:   false,
			Data: "SignIn",
		})
	}
}
