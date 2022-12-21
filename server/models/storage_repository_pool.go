package models

import (
	"time"
)

type StorageRepositoryPool struct {
	CreatedAt time.Time `json:"created_at" xorm:"created DATETIME"`
	DeletedAt time.Time `json:"deleted_at" xorm:"deleted DATETIME"`
	Ext       string    `json:"ext" xorm:"comment('文件扩展名') VARCHAR(30)"`
	Hash      string    `json:"hash" xorm:"comment('文件的唯一标识') VARCHAR(32)"`
	Id        int       `json:"id" xorm:"not null pk autoincr INT"`
	Identity  string    `json:"identity" xorm:"VARCHAR(36)"`
	Name      string    `json:"name" xorm:"VARCHAR(255)"`
	Path      string    `json:"path" xorm:"comment('文件路径') VARCHAR(255)"`
	Size      int64     `json:"size" xorm:"comment('文件大小') INT"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated DATETIME"`
}

func (r *StorageRepositoryPool) TableName() string {
	return "storage_repository_pool"
}
