package main

import (
	"fmt"
	//"strconv"
	//"strings"

	//"gohttpserver/tools"
	"net/http"

	"gopkg.in/macaron.v1"

	"gohttpserver/config"
	//"gohttpserver/dbcontrollers"
	"gohttpserver/routers"
)

func newInstance() *macaron.Macaron {
	m := macaron.New()
	return m
}

func main() {
	// var ss string = "U-00000001"
	// s := strings.Split(ss, "-")
	// fmt.Println(s[1])
	// int1164, _ := strconv.ParseInt(s[1], 10, 64)

	// fmt.Println(int1164)
	m := newInstance()

	// userInfos := make([]*dbcontrollers.UserInfo, 0)
	// userInfo := &dbcontrollers.UserInfo{}
	// userInfos = userInfo.GetUserInfoSQL("")
	// fmt.Println("-------------", userInfos[0])

	// ei := &tools.EmailInfo{}
	// ei.FromAddr = "277260074@qq.com"
	// ei.SendAddr = "shitou-1982911@163.com"
	// ei.SendInfo = "测试邮箱发送\n谢谢！"
	// isok := tools.SendEmail(ei)
	// fmt.Println(isok)

	routers.Regist(m)
	listenAddr := fmt.Sprintf("0.0.0.0:%d", conf.Httpport)
	fmt.Println(listenAddr)
	fmt.Println("start")
	if err := http.ListenAndServe(listenAddr, m); err != nil {
		fmt.Println("Err:", err)
	}
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type KeyBasic struct {
// 	UserName  string `json:"username"`
// 	Algorithm string `json:"algo"`
// 	Size      uint   `json:"size"`
// }
// type KeyUsage struct {
// 	UserName  string `json:"username"`
// 	Algorithm string `json:"algo"`
// 	Size      uint   `json:"size"`
// 	Usage     string `json:"usage,omitempty"`
// }

// func main() {
// 	key := KeyUsage{
// 		UserName:  "hello",
// 		Algorithm: "rsa",
// 		Size:      2048,
// 		Usage:     "",
// 	}
// 	keyBytes, err := json.Marshal(key)
// 	if err != nil {
// 		fmt.Println("json unmarshal err")
// 	}
// 	fmt.Println(string(keyBytes))

// 	keyUse := KeyBasic{}
// 	err = json.Unmarshal(keyBytes, &keyUse)
// 	if err != nil {
// 		fmt.Printf("err:%s\n", err.Error())
// 	}
// 	fmt.Printf("basic:%v\n", keyUse)
// }
