package response

import (
	"github.com/gofiber/fiber/v3"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data any, msg string, c fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c fiber.Ctx) error {
	return Result(SUCCESS, fiber.Map{}, "操作成功", c)
}

func OkWithMessage(message string, c fiber.Ctx) error {
	return Result(SUCCESS, fiber.Map{}, message, c)
}

func OkWithDetailed(data any, message string, c fiber.Ctx) error {
	return Result(SUCCESS, data, message, c)
}

func Fail(c fiber.Ctx) error {
	return Result(ERROR, fiber.Map{}, "操作失败", c)
}

func FailWithMessage(message string, c fiber.Ctx) error {
	return Result(ERROR, fiber.Map{}, message, c)
}

func FailWithDetailed(data any, message string, c fiber.Ctx) error {
	return Result(ERROR, data, message, c)
}
