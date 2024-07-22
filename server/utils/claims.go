package utils

import (
	"github.com/gofiber/fiber/v3"
	"net"

	"frp-platform/global"
	"frp-platform/model/request"
)

func ClearToken(c fiber.Ctx) {
	c.ClearCookie("x-token")
	// 增加cookie x-token 向来源的web添加
	// host, _, err := net.SplitHostPort(c.Host())
	// if err != nil {
	// 	host = c.Host()
	// }
	//
	// if net.ParseIP(host) != nil {
	// 	c.Cookie(&fiber.Cookie{
	// 		Name:     "x-token",
	// 		Value:    "",
	// 		Path:     "/",
	// 		MaxAge: -1,
	// 		Domain: "",
	// 	})
	// } else {
	// 	c.Cookie(&fiber.Cookie{
	// 		Name:     "x-token",
	// 		Value:    "",
	// 		Path:     "/",
	// 		MaxAge: -1,
	// 		Domain:   host,
	// 	})
	// }
}

func SetToken(c fiber.Ctx, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Host())
	if err != nil {
		host = c.Host()
	}

	if net.ParseIP(host) != nil {
		c.Cookie(&fiber.Cookie{
			Name:   "x-token",
			Value:  token,
			Path:   "/",
			MaxAge: maxAge,
			Domain: "",
		})
	} else {
		c.Cookie(&fiber.Cookie{
			Name:   "x-token",
			Value:  token,
			Path:   "/",
			MaxAge: maxAge,
			Domain: host,
		})
	}
}

func GetToken(c fiber.Ctx) string {
	token := c.Cookies("x-token")
	if token == "" {
		token = c.Get("x-token")
	}
	return token
}

func GetClaims(c fiber.Ctx) (*request.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.LOG.Error("jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

func GetUserID(c fiber.Ctx) uint {
	data := c.Locals("claims")
	if data == nil {
		claims, err := GetClaims(c)
		if err != nil {
			return 0
		}
		return claims.BaseClaims.ID
	} else {
		claims := data.(*request.CustomClaims)
		return claims.BaseClaims.ID
	}
}

func GetUserUUID(c fiber.Ctx) string {
	data := c.Locals("claims")
	if data == nil {
		claims, err := GetClaims(c)
		if err != nil {
			return ""
		}
		return claims.BaseClaims.UUID
	} else {
		claims := data.(*request.CustomClaims)
		return claims.BaseClaims.UUID
	}
}

func GetAuthorityId(c fiber.Ctx) uint {
	data := c.Locals("claims")
	if data == nil {
		claims, err := GetClaims(c)
		if err != nil {
			return 0
		}
		return claims.AuthorityId
	} else {
		claims := data.(*request.CustomClaims)
		return claims.AuthorityId
	}
}

func GetClaimInfo(c fiber.Ctx) *request.CustomClaims {
	data := c.Locals("claims")
	if data == nil {
		claims, err := GetClaims(c)
		if err != nil {
			return nil
		}
		return claims
	} else {
		claims := data.(*request.CustomClaims)
		return claims
	}
}
