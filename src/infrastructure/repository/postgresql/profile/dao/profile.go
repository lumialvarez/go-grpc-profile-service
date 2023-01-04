package dao

import (
	"gorm.io/datatypes"
)

type Profile struct {
	Id         int64  `gorm:"primaryKey"`
	Language   string `gorm:"uniqueIndex"`
	Data       datatypes.JSON
	LastUpdate datatypes.Date
	Status     bool
}
