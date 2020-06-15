package controllers

import (
	"fmt"
	"strings"

	// "sort"
	// "bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strconv"

	// "gohttpserver/dbcontrollers"
	"encoding/json"
	"gohttpserver/dbcontrollers"
	"gohttpserver/models"
	"time"

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
	UserId   string
	UserName string
}

type RetOrderInfo struct {
	UserId     string
	UserName   string
	RemainTime string
	IfValid    bool
	MustUpdate bool
	Info       string
}

type BufPayPar struct {
	RealizeId      string
	UserName       string
	ReturnLinkaddr string
}

type BufPayRequestPar struct {
	RealizeId  string
	Pname      string
	BType      string
	Price      string
	Ptime      string
	OrderId    string
	UserName   string
	ReturnLink string
}

type BufPayToBufPay struct {
	Name       string `json:"name"`
	Pay_type   string `json:"pay_type"`
	Price      string `json:"price"`
	Order_id   string `json:"order_id"`
	Order_uid  string `json:"order_uid"`
	Notify_url string `json:"notify_url"`
	Return_url string `json:"return_url"`
	Sign       string `json:"sign"`
}

type QrccodeInfoToClinet struct {
	OrderId string
	QrcType string
	QrcCode string
	QrPrice string
	Price   string
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
		fmt.Println(len(dataMap))
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

// func calTimeLeft(dm []map[string]string) {
// 	var keys []int
// 	timeTemplate1 := "2006-01-02T15:04:05Z"
// 	for _, v := range dm {
// 		t1 := v["CreateTime"]
// 		stamp, _ := time.ParseInLocation(timeTemplate1, t1, time.Local)
// 		vInt := stamp.Unix()
// 		fmt.Println("============", vInt)
// 		strInt64 := strconv.FormatInt(vInt, 10)
// 		id16, _ := strconv.Atoi(strInt64)
// 		keys = append(keys, id16)
// 	}

// 	sort.Ints(keys)

// 	for _, value := range keys {
// 		fmt.Println("---dmvalue---", dm[value]["CreateTime"])
// 	}

// }

func GetOrderByUser(ctx *macaron.Context) {
	//ctx.Resp.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("-------GetOrderByUser------")
	jsonStr := ctx.Query("GetOrderByUser")
	uni := &UserNameInfo{}
	err := json.Unmarshal([]byte(jsonStr), uni)
	ret := false
	roi := &RetOrderInfo{}
	roi.UserId = uni.UserId
	roi.UserName = uni.UserName
	roi.IfValid = false
	roi.RemainTime = "0"
	roi.MustUpdate = false
	roi.Info = "版本不是最新无法登入!!!"
	if err == nil {
		sql := `SELECT *from "AK_Order" WHERE "UserId" = ?`
		dataMap, _ := dbcontrollers.GetOrm().QueryString(sql, uni.UserId)
		timeTemplate1 := "2006-01-02T15:04:05Z"
		rt := 0
		for _, value := range dataMap {
			t1 := value["CreateTime"]
			vts := value["ValidityTime"]
			vti, _ := strconv.Atoi(vts)
			stamp, _ := time.ParseInLocation(timeTemplate1, t1, time.Local)
			tt := time.Now()
			residueTime := tt.Unix() - stamp.Unix()
			fmt.Println("---residue---", residueTime)

			rt = vti - int(residueTime/3600)
			break
		}
		//calTimeLeft(dataMap)
		roi.RemainTime = strconv.Itoa(rt)
	}
	ret = true
	ctx.JSON(200, &ContextResult{
		Ok:    ret,
		Data:  "GetOrderByUser",
		Value: &roi,
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

func httpPos(url, jsonStr string) *QrccodeInfoToClinet {
	fmt.Println("---httpPos---", jsonStr)
	// payInfoDataBuf, _ := json.Marshal(jsonStr)
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonStr) /*bytes.NewBuffer(payInfoDataBuf)*/)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {

		// handle error

	}
	body, _ := ioutil.ReadAll(resp.Body)
	var dat map[string]interface{}
	json.Unmarshal([]byte(body), &dat)
	qi := &QrccodeInfoToClinet{}
	if dat["status"] == "ok" {
		qi.QrcType = dat["pay_type"].(string)
		qi.QrPrice = dat["qr_price"].(string)
		qi.Price = dat["price"].(string)
		qi.QrcCode = dat["qr_img"].(string)
	}
	return qi
}

func GetQRCodeUrl(ctx *macaron.Context) {
	t := time.Now()
	orderId := t.Format("20060102150405")
	fmt.Println(orderId)
	var ret bool = false

	jsonStr := ctx.Query("GetQRCodeUrl")
	bpp := &BufPayPar{}
	err := json.Unmarshal([]byte(jsonStr), bpp)
	if err != nil {
		ret = false
		ctx.JSON(200, &ContextResult{
			Ok:    ret,
			Data:  "GetRealize",
			Value: "",
		})
		return
	}

	brp := &BufPayRequestPar{}
	brp.RealizeId = bpp.RealizeId
	brp.OrderId = orderId
	brp.BType = "wechat"
	brp.UserName = bpp.UserName
	brp.ReturnLink = bpp.ReturnLinkaddr

	sql := `SELECT *from "AK_RealizeAndPrice" WHERE "id" = ?`
	dataMap, err := dbcontrollers.GetOrm().QueryString(sql, bpp.RealizeId)

	for _, value := range dataMap {
		brp.Pname = value["name"]
		brp.Price = value["value"]
		brp.Ptime = value["ptime"]
		break
	}
	brp.ReturnLink += "PlaceOrder?Data=" + brp.UserName + "|" + brp.OrderId + "|" + brp.RealizeId + "|" + brp.Ptime
	str := brp.Pname + brp.BType + brp.Price + brp.OrderId + brp.UserName + brp.ReturnLink + "http://www.kaidany.com" + "1f93e4388f7b4e45959f6accaa1cff28"

	h := md5.New()
	h.Write([]byte(str))
	str = hex.EncodeToString(h.Sum(nil))
	pay_data := &BufPayToBufPay{}
	pay_data.Name = brp.Pname
	pay_data.Pay_type = brp.BType
	pay_data.Price = brp.Price
	pay_data.Order_id = brp.OrderId
	pay_data.Order_uid = brp.UserName
	pay_data.Notify_url = brp.ReturnLink
	pay_data.Return_url = "http://www.kaidany.com"
	pay_data.Sign = str

	bufStr := "name=" + brp.Pname + "&" + "pay_type=" + brp.BType + "&" + "price=" + brp.Price + "&" + "order_id=" + brp.OrderId + "&" + "order_uid=" + brp.UserName + "&" + "notify_url=" + brp.ReturnLink + "&" + "return_url=" + "http://www.kaidany.com" + "&" + "sign=" + str

	// b, err := json.Marshal(pay_data)
	// if err == nil {
	// 	fmt.Println("json:", string(b))
	// }
	qi := httpPos("https://bufpay.com/api/pay/96967?format=json", bufStr)
	qi.OrderId = orderId
	ctx.JSON(200, &ContextResult{
		Ok:    true,
		Data:  "GetRealize",
		Value: &qi,
	})
}
