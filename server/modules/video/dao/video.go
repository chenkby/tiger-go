package dao

import (
	"gorm.io/gorm"
	"tiger-go/tiger"
	"tiger-go/tiger/db"
)

var Video = &db.DB{
	DB:         tiger.Db().Table("video").Session(&gorm.Session{}),
	PrimaryKey: "video_id",
}
