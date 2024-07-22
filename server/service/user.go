package service

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"

	"frp-platform/global"
	"frp-platform/model"
	"frp-platform/model/request"
	"frp-platform/utils"
)

type UserService struct {
}

func (u *UserService) Register(req request.Register) error {
	err := global.DB.Where("username = ?", req.Username).First(&model.User{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("%s用户已注册", req.Username)
	}
	user := model.User{
		UUID:        uuid.New().String(),
		Username:    req.Username,
		Password:    utils.BcryptHash(req.Password),
		SideMode:    "dark",
		HeaderImg:   "",
		BaseColor:   "#fff",
		AuthorityId: 0, // 分配一个普通用户的权限
		Phone:       req.Phone,
		Email:       req.Email,
		Enable:      1,
	}
	err = global.DB.Create(&user).Error
	if err != nil {
		return fmt.Errorf("创建用户失败, err: %s", err)
	}
	return nil
}

func (u *UserService) Login(info model.User, c fiber.Ctx) (user model.User, token string, expire int64, err error) {
	err = global.DB.Where("username = ?", info.Username).First(&user).Error
	if err != nil {
		return user, token, expire, fmt.Errorf("query user err: %s", err)
	}
	if ok := utils.BcryptCheck(info.Password, user.Password); !ok {
		return user, token, expire, fmt.Errorf("密码错误")
	}

	// 签发token
	j := utils.NewJWT()
	claims := j.CreateClaims(request.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err = j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败", zap.Error(err))
		return user, token, expire, fmt.Errorf("获取token失败, err: %s", err)
	}
	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	return user, token, claims.RegisteredClaims.ExpiresAt.Unix() * 1000, nil
}
