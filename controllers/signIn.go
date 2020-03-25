package controllers

import (
	"encoding/json"
	"fmt"
	"gohttpserver/tools"
	"strconv"
	"time"

	"gohttpserver/models"

	"strings"

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

		whereStr := `"UserName" = ?`
		userInfos := userInfo.GetUserInfoSQL(whereStr, signInInfo.UserName)
		if len(userInfos) > 0 {
			ctx.JSON(200, &ContextResult{
				Ok:   false,
				Data: "SignIn err has same name",
			})
			return
		}

		UId, flag := userInfo.GetMaxUserId()
		var UIdMaxInt int
		if flag == true {
			string_slice := strings.Split(UId, "-")
			strMax := string_slice[1]
			UIdMaxInt, err = strconv.Atoi(strMax)
		}
		strUId := fmt.Sprintf("U-%010d", UIdMaxInt+1)

		userInfo.Id = uuid.String()
		userInfo.UserId = strUId
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
