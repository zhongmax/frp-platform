package global

import (
	"gorm.io/gorm"
	"time"
)

type MODEL struct {
	ID        uint `json:"ID" gorm:"primarykey"`
	CreateAt  time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type MODEL_NO_DELETE struct {
	ID        uint `json:"ID" gorm:"primarykey"`
	CreateAt  time.Time
	UpdatedAt time.Time
}
