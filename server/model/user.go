package model

import (
	"frp-platform/global"
)

type User struct {
	global.MODEL
	UUID        string `json:"uuid" gorm:"column:uuid;comment:用户uuid"`
	Username    string `json:"username" gorm:"column:username;comment:用户名"`
	Password    string `json:"-" gorm:"column:password;comment:用户密码"`
	SideMode    string `json:"sideMode" gorm:"column:sideMode;default:dark;comment:用户主题"`
	HeaderImg   string `json:"headerImg" gorm:"column:headerImg;comment:用户头像"`
	BaseColor   string `json:"baseColor" gorm:"column:baseColor;comment:基础颜色"`
	AuthorityId uint   `json:"authorityId" gorm:"column:authorityId;comment:角色id"`
	Phone       string `json:"phone" gorm:"column:phone;comment:手机号"`
	Email       string `json:"email" gorm:"column:email;comment:邮箱"`
	Enable      int    `json:"enable" gorm:"column:enable;comment:用户是否被冻结"` // 1 正常 2 冻结
}

func (User) TableName() string {
	return "users"
}
