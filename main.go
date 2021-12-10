package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

//路由设置,用于注册路由
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}

func main() {

	cfg, err := tool.ParseConfig("./config/app.json")

	if err != nil {
		panic(err.Error())
	}

	app := gin.Default()

	registerRouter(app)

	err = app.Run(cfg.AppHost + ":" + cfg.AppPort)
	if err != nil {
		return
	}
}
