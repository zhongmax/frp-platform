package model

import "frp-platform/global"

type Menu struct {
	global.MODEL
	ParentId  uint   `json:"parentId" gorm:"column:parent_id;comment:父菜单id"`
	Path      string `json:"path" gorm:"column:path;comment:路由path"`
	Name      string `json:"name" gorm:"column:name;comment:路由name"`
	Hidden    bool   `json:"hidden" gorm:"column:hidden;comment:是否隐藏"`
	Component string `json:"component" gorm:"column:component;comment:对应前端文件路径"`
	Sort      int    `json:"sort" gorm:"column:sort;comment:排序标记"`
}

type MenuMeta struct {
	ActiveName  string `json:"activeName" gorm:"column:active_name;comment:高亮菜单"`
	KeepAlive   bool   `json:"keepAlive" gorm:"column:keep_alive;comment:是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"column:default_menu;comment:是否为基础路由"`
	Title       string `json:"title" gorm:"column:title;comment:菜单名称"`
	Icon        string `json:"icon" gorm:"column:icon;comment:菜单图标"`
	CloseTab    bool   `json:"closeTab" gorm:"column:close_tab;comment:自动关闭tab"`
}

type MenuParameter struct {
	global.MODEL
	SystemBaseId uint
	Type         string `json:"type" gorm:"column:type;comment:"`
}
