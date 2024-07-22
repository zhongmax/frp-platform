package v1

import "github.com/gofiber/fiber/v3"

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(router fiber.Router) {
	userRouter := router.Group("user")
	api := new(userApi)
	{
		userRouter.Post("changePassword", api.ChangePassword)
	}
	{

	}
}

type userApi struct {
}

func (s *userApi) ChangePassword(c fiber.Ctx) error {
	return nil
}
