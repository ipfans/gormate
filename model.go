package gormate

import (
	"gorm.io/gorm"
)

type Model struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt TimeFormat     `json:"createdAt" gorm:"type:time"`
	UpdatedAt TimeFormat     `json:"updatedAt" gorm:"type:time"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
