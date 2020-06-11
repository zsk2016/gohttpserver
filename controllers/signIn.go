package controllers

import (
	"encoding/json"
	"fmt"
	"gohttpserver/models"
	"gohttpserver/tools"
	"regexp"
	"strconv"
	"time"

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
	UserType  string
	PcNum     string
	CpuId     string
}

func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func SignInFun(ctx *macaron.Context) {
	//ctx.Resp.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("------sign in-----")
	jsonStr := ctx.Query("SignIn")
	signInInfo := &SignInInfo{}
	err := json.Unmarshal([]byte(jsonStr), signInInfo)
	if err == nil {
		uuid, _ := uuid.NewV4()
		fmt.Println(signInInfo.UserName, signInInfo.EmailAddr)
		userInfo := models.UserInfo{}

		if !VerifyEmailFormat(signInInfo.EmailAddr) {
			ctx.JSON(200, &ContextResult{
				Ok:   false,
				Data: "注册邮箱格式不对!!!",
			})
			return
		}

		whereStr := `"UserName" = ?`
		userInfos := userInfo.GetUserInfoSQL(whereStr, signInInfo.UserName)
		if len(userInfos) > 0 {
			ctx.JSON(200, &ContextResult{
				Ok:   false,
				Data: "该注册的用户名已经存在!!!",
			})
			return
		}

		whereStr1 := `"UserEmail" = ?`
		userInfos1 := userInfo.GetUserInfoSQL(whereStr1, signInInfo.EmailAddr)
		if len(userInfos1) > 0 {
			ctx.JSON(200, &ContextResult{
				Ok:   false,
				Data: "该注册的邮箱已经存在!!!",
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
		userInfo.UserType = signInInfo.UserType
		userInfo.PcNum = signInInfo.PcNum
		userInfo.CpuId = signInInfo.CpuId
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
