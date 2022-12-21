package models

import (
	"time"
)

type StorageShare struct {
	ClickNum               int       `json:"click_num" xorm:"default 0 comment('点击次数') INT"`
	CreatedAt              time.Time `json:"created_at" xorm:"created DATETIME"`
	DeletedAt              time.Time `json:"deleted_at" xorm:"deleted DATETIME"`
	ExpiredTime            int       `json:"expired_time" xorm:"comment('失效时间，单位秒, 【0-永不失效】') INT"`
	Id                     int       `json:"id" xorm:"not null pk autoincr INT"`
	Identity               string    `json:"identity" xorm:"VARCHAR(36)"`
	RepositoryIdentity     string    `json:"repository_identity" xorm:"comment('公共池中的唯一标识') VARCHAR(36)"`
	UpdatedAt              time.Time `json:"updated_at" xorm:"updated DATETIME"`
	UserIdentity           string    `json:"user_identity" xorm:"VARCHAR(36)"`
	UserRepositoryIdentity string    `json:"user_repository_identity" xorm:"comment('用户池子中的唯一标识') VARCHAR(36)"`
}

func (s *StorageShare) TableName() string {
	return "storage_share"
}
