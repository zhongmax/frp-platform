package v1

import "github.com/gofiber/fiber/v3"

type InitRouter struct {
}

func (s *InitRouter) InitInitRouter(router fiber.Router) {
	initRouter := router.Group("initialize")
	dbApi := new(initDBApi)
	{
		initRouter.Post("initDB", dbApi.InitDB)
		initRouter.Post("checkDB", dbApi.CheckDB)
	}
}

type initDBApi struct {
}

func (s *initDBApi) InitDB(c fiber.Ctx) error {
	return nil
}

func (s *initDBApi) CheckDB(c fiber.Ctx) error {
	return nil
}
