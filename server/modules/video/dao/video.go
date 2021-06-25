package dao

import (
	"tiger-go/tiger"
)

//var Video = &db.DB{
//	DB:         tiger.Db().Table("video").Session(&gorm.Session{}),
//	PrimaryKey: "video_id",
//}

var Video = tiger.Db().Table("video")
