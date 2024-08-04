package common

import "time"

type SQLModel struct {
	ID int `json:"id" gorm:"column:id;"`
	// FakeId *UID `json:"id" gorm:"-"`
	Status   int        `json:"status" gorm:"column:status;default:1;"`
	CreateAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime"`
	UpdateAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;autoCreateTime"`
}
