package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"size:120";not null`
	Email    string `gorm:"type:varchar(100);unique_index;not null"`
	Birthday *time.Time
	Address  string `gorm:"index:addr"` // create index with name `addr` for address
}
