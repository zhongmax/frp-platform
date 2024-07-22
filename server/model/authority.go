package model

import "frp-platform/global"

type Authority struct {
	global.MODEL
	AuthorityName string `json:"authorityName" gorm:"column:authority_name;comment:角色名称"`
	ParentId      *uint  `json:"parentId" gorm:"column:parent_id;comment:父角色id"`
	DefaultRouter string `json:"defaultRouter" gorm:"column:default_router;default:dashboard;comment:默认菜单"`
}

func (Authority) TableName() string {
	return "authorities"
}
