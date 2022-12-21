package models

import (
	"time"
)

type StorageUserRepository struct {
	CreatedAt          time.Time `json:"created_at" xorm:"created DATETIME"`
	DeletedAt          time.Time `json:"deleted_at" xorm:"deleted DATETIME"`
	Ext                string    `json:"ext" xorm:"comment('文件或文件夹类型') VARCHAR(255)"`
	Id                 int       `json:"id" xorm:"not null pk autoincr INT"`
	Identity           string    `json:"identity" xorm:"VARCHAR(36)"`
	Name               string    `json:"name" xorm:"VARCHAR(255)"`
	ParentId           int       `json:"parent_id" xorm:"INT"`
	RepositoryIdentity string    `json:"repository_identity" xorm:"VARCHAR(36)"`
	UpdatedAt          time.Time `json:"updated_at" xorm:"updated DATETIME"`
	UserIdentity       string    `json:"user_identity" xorm:"VARCHAR(36)"`
}

func (r *StorageUserRepository) TableName() string {
	return "storage_user_repository"
}
