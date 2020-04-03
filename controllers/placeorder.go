package controllers

import (
	"fmt"
	"gohttpserver/models"

	// "strconv"
	"strings"
	"time"

	"github.com/satori/go.uuid"
	"gopkg.in/macaron.v1"
)

type OrderInfo struct {
	UserName     string
	OrderId      string
	RealizeId    string
	ValidityTime string
}

func GetNowOrderId(ctx *macaron.Context) {
	t := time.Now()
	orderId := t.Format("20060102150405")
	fmt.Println(orderId)
	ctx.JSON(200, &ContextResult{
		Ok:    true,
		Data:  "OrderId",
		Value: orderId,
	})
}

func PlaceOrderFun(ctx *macaron.Context) {
	fmt.Println("------PlaceOrder-----")
	orderInfo := &OrderInfo{}
	// orderInfo.UserName = ctx.Query("UserName")
	// orderInfo.OrderId = ctx.Query("OrderId")
	// orderInfo.RealizeId = ctx.Query("RealizeId")
	// orderInfo.ValidityTime = ctx.Query("ValidityTime")
	strData := ctx.Query("Data")
	kv := strings.Split(strData, "|")
	if len(kv) != 4 {
		ctx.JSON(200, &ContextResult{
			Ok:   false,
			Data: "PlaceOrder",
		})
		return
	}
	orderInfo.UserName = kv[0]
	orderInfo.OrderId = kv[1]
	orderInfo.RealizeId = kv[2]
	orderInfo.ValidityTime = kv[3]
	uid := GetUserIdByUserName(orderInfo.UserName)

	if uid == "0" {
		ctx.JSON(200, &ContextResult{
			Ok:   false,
			Data: "PlaceOrder",
		})
		return
	}

	order := models.Order{}
	orderId, _ := uuid.NewV4()
	order.Id = orderId.String()
	// UId, flag := order.GetMaxOrderId()
	// var UIdMaxInt int
	// if flag == true {
	// 	string_slice := strings.Split(UId, "-")
	// 	strMax := string_slice[1]
	// 	UIdMaxInt, _ = strconv.Atoi(strMax)
	// }
	// strUId := fmt.Sprintf("O-%010d", UIdMaxInt+1)
	order.OrderId = orderInfo.OrderId
	order.UserId = uid
	order.RealizeId = orderInfo.RealizeId
	order.CreateTime = time.Now()
	order.ValidityTime = orderInfo.ValidityTime
	ret := false
	ret = order.InsertOrderSQL()

	ctx.JSON(200, &ContextResult{
		Ok:   ret,
		Data: "PlaceOrder",
	})

}
