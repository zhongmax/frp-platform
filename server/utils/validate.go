package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v3"

	"frp-platform/global"
)

func BindAndValidate(obj any, c fiber.Ctx) error {
	err := c.Bind().JSON(&obj)
	if err != nil {
		return fmt.Errorf("bind json err: %s", err)
	}
	err = global.Validate(obj)
	if err != nil {
		return fmt.Errorf("validate struct err: %s", err)
	}
	return nil
}
