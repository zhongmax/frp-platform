package core

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"net/http"

	v1 "frp-platform/api/v1"
	"frp-platform/global"
)

func RunServer() {

	app := fiber.New()
	app.Use(recover.New())

	routerGroup := v1.RouterGroupApp

	publicGroup := app.Group(global.CONFIG.System.RouterPrefix)

	privateGroup := app.Group(global.CONFIG.System.RouterPrefix)

	{
		publicGroup.Get("/health", func(c fiber.Ctx) error {
			return c.Status(http.StatusOK).JSON("ok")
		})
	}
	{
		routerGroup.InitBaseRouter(publicGroup) // 注册登录相关, 不做鉴权
		routerGroup.InitInitRouter(publicGroup) // 自动初始化相关
	}

	{
		routerGroup.InitUserRouter(privateGroup) // 用户相关
	}
	global.LOG.Error(app.Listen(fmt.Sprintf(":%d", global.CONFIG.System.ServerPort)).Error())
}
