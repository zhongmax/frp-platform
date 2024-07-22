package model

import "frp-platform/global"

type Api struct {
	global.MODEL
	Path        string `json:"path" gorm:"comment:api路径"`
	Description string `json:"description" gorm:"column:description;comment:描述"`
	ApiGroup    string `json:"apiGroup" gorm:"column:apiGroup;comment:api组"`
	Method      string `json:"method" gorm:"column:method;default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (Api) TableName() string {
	return "apis"
}
