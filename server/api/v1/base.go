package v1

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"

	"frp-platform/global"
	"frp-platform/model"
	"frp-platform/model/request"
	"frp-platform/model/response"
	"frp-platform/service"
	"frp-platform/utils"
)

var store = base64Captcha.DefaultMemStore

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(router fiber.Router) {
	baseRouter := router.Group("base")
	api := new(baseApi)
	{
		baseRouter.Post("register", api.Register)
		baseRouter.Post("login", api.Login)
		baseRouter.Post("captcha", api.Captcha)
	}
}

type baseApi struct {
}

func (b *baseApi) Register(c fiber.Ctx) error {
	var req request.Register
	err := utils.BindAndValidate(req, c)
	if err != nil {
		global.LOG.Error("bind and validate error", zap.Error(err))
		return response.FailWithMessage("注册失败", c)
	}

	openCaptcha := global.CONFIG.Captcha.OpenCaptcha
	openCaptchaTimeout := global.CONFIG.Captcha.OpenCaptchaTimeOut
	key := c.IP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeout))
	}

	var oc = openCaptcha == 0 || openCaptcha < interfaceToInt(v)
	if oc || req.CaptchaId == "" || req.Captcha == "" || !store.Verify(req.CaptchaId, req.Captcha, true) {
		return response.FailWithMessage("验证码错误", c)
	}

	userService := service.ServiceGroupApp.UserService
	if err = userService.Register(req); err != nil {
		global.LOG.Error("register error", zap.String("username", req.Username), zap.Error(err))
		return response.FailWithMessage("注册失败", c)
	} else {
		return response.OkWithMessage("注册成功", c)
	}
}

func (b *baseApi) Login(c fiber.Ctx) error {
	var req request.Login
	if err := utils.BindAndValidate(&req, c); err != nil {
		global.LOG.Error("bind and validate error", zap.String("username", req.Username), zap.Error(err))
		return response.FailWithMessage("登录失败", c)
	}

	openCaptcha := global.CONFIG.Captcha.OpenCaptcha
	openCaptchaTimeout := global.CONFIG.Captcha.OpenCaptchaTimeOut
	key := c.IP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeout))
	}

	var oc = openCaptcha == 0 || openCaptcha < interfaceToInt(v)
	if oc || req.CaptchaId == "" || req.Captcha == "" || !store.Verify(req.CaptchaId, req.Captcha, true) {
		return response.FailWithMessage("验证码错误", c)
	}

	userService := service.ServiceGroupApp.UserService
	u := model.User{Username: req.Username, Password: req.Password}
	user, token, expires, err := userService.Login(u, c)
	if err != nil {
		global.LOG.Error("登录失败", zap.String("username", req.Username), zap.Error(err))
		return response.FailWithMessage("登录失败", c)
	} else {
		return response.OkWithDetailed(response.LoginResp{
			User:      user,
			Token:     token,
			ExpiresAt: expires,
		}, "登录成功", c)
	}
}

func (b *baseApi) Captcha(c fiber.Ctx) error {
	openCaptcha := global.CONFIG.Captcha.OpenCaptcha
	openCaptchaTimeout := global.CONFIG.Captcha.OpenCaptchaTimeOut
	key := c.IP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeout))
	}
	var oc = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.LOG.Error("验证码获取失败", zap.Error(err))
		return response.FailWithMessage("验证码获取失败", c)
	}
	return response.OkWithDetailed(response.CaptchaResp{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

func interfaceToInt(v any) int {
	switch v.(type) {
	case int:
		return v.(int)
	default:
		return 0
	}
}
