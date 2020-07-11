package gormate

import (
	"gorm.io/gorm"
)


type Model struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt TimeFormat     `json:"createdAt"`
	UpdatedAt TimeFormat     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
