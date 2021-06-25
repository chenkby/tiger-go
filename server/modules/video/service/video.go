package service

import "tiger-go/tiger"

var Video = new(VideoService)

type VideoService struct {
}

func (s *VideoService) FindOne(dest interface{}, condition interface{}) error {
	err := tiger.Db().Table("video").Find(dest, condition).Error
	return err
}
