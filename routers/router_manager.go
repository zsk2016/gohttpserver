package routers

import (
	"gohttpserver/controllers"
	"net/http"

	"github.com/go-macaron/pongo2"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func Regist(m *macaron.Macaron) {
	RegistBeforeRouter(m)
	RegistGlobalRouter(m)
	RegistGroupRouter(m)
	RegistRouter(m)
	RegistAfterRouter(m)
}

//注册全局的路由处理函数
func RegistGlobalRouter(m *macaron.Macaron) {
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(session.Sessioner())
	m.Use(macaron.Static("static"))
	m.Use(macaron.Static("views"))
	m.Use(pongo2.Pongoer(pongo2.Options{
		Directory:  "views",
		IndentJSON: macaron.Env != macaron.PROD,
		IndentXML:  macaron.Env != macaron.PROD,
	}))
}

func RegistGroupRouter(m *macaron.Macaron) {
	//	m.Group("/dataserver", func() { //数据服务
	//		m.Any("/forgetpasswordbyphone", controllers.ForgetPasswordByPhone) //通过手机号找回密码
	//		m.Any("/forgetpasswordbyemail", controllers.ForgetPasswordByEmail) //通过邮箱找回密码
	//	})
}

func RegistRouter(m *macaron.Macaron) {
	m.Any("/", controllers.Default)
	m.Any("/imageShow", controllers.ImageShow)
	m.Any("/getData", controllers.GetData)
	m.Any("/GetUserCpuId", controllers.GetUserCpuId)
	m.Any("/SignIn", controllers.SignInFun)
	m.Any("/LogIn", controllers.LogInFun)
	m.Any("/PlaceOrder", controllers.PlaceOrderFun)
	m.Any("/GetNowOrderId", controllers.GetNowOrderId)
	m.Any("/GetRealize", controllers.GetRealize)
	m.Any("/GetOrderByUser", controllers.GetOrderByUser)
	m.Any("/GetQRCodeUrl", controllers.GetQRCodeUrl)
	//m.Any("/GetTimeRemain", controllers.GetTimeRemain)
	//	m.Any("/regist", controllers.Regist)
	//	m.Any("/register", controllers.Register)
	//	m.Any("/useragreement", controllers.UserAgreement)
	//	m.Any("/promisebook", controllers.PromiseBook)
	//	m.Any("/forgetpassword", controllers.ForgetPassword)
}

func RegistBeforeRouter(m *macaron.Macaron) {
	m.Before(func(rw http.ResponseWriter, req *http.Request) bool {
		return false
	})
}

func RegistAfterRouter(m *macaron.Macaron) {
	m.Use(func(ctx *macaron.Context) {
		ctx.Next()
	})
}
