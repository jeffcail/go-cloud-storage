package models

import (
	"time"
)

type StorageUser struct {
	CreatedAt time.Time `json:"created_at" xorm:"created DATETIME"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted DATETIME"`
	Email     string    `json:"email" xorm:"VARCHAR(100)"`
	Id        int       `json:"id" xorm:"not null pk autoincr INT"`
	Identity  string    `json:"identity" xorm:"VARCHAR(36)"`
	Name      string    `json:"name" xorm:"VARCHAR(60)"`
	Password  string    `json:"password" xorm:"VARCHAR(100)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated DATETIME"`
}

func (u *StorageUser) TableName() string {
	return "storage_user"
}
