package gormate

import (
	"gorm.io/gorm"
)

// Model for gorm
type Model struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	CreatedAt TimeFormat     `json:"createdAt" gorm:"autoCreateTime;type:time"`
	UpdatedAt TimeFormat     `json:"updatedAt" gorm:"autoUpdateTime;type:time"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// PK returns primaryKey
func (m *Model) PK() uint64 {
	return m.ID
}
