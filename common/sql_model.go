package common

import "time"

type SQLModel struct {
	Id int `json:"id" gorm:"column:id;"`
	//FakeId   *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"create_at,omitempty" gorm:"created_at"`
	UpdatedAt *time.Time `json:"update_at,omitempty" gorm:"updated_at"`
}
