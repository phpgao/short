package main

import (
	"gorm.io/gorm"
	"time"
)

type CreateUrl struct {
	Url    string `form:"url"`
	Path   string `form:"path"`
	Expire uint64 `form:"expire"`
}

type Url struct {
	ID        uint   `gorm:"primaryKey"`
	Url       string `gorm:"size:100;uniqueIndex"`
	Path      string `gorm:"size:100;uniqueIndex"`
	Expire    uint64
	CreatedAt int64
	UpdatedAt int64
}

func (u *Url) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now().Unix()
	return nil
}

func (u *Url) BeforeCreate(tx *gorm.DB) error {
	if u.UpdatedAt == 0 {
		u.UpdatedAt = time.Now().Unix()
	}
	u.CreatedAt = time.Now().Unix()
	return nil
}
