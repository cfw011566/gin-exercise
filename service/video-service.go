package service

import (
	"example/golang-gin-poc/entity"
	"example/golang-gin-poc/repository"
	"fmt"
)

type VideoService interface {
	Save(entity.Video) error
	Update(entity.Video) error
	Delete(entity.Video) error
	FindAll() []entity.Video
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video entity.Video) error {
	service.videos = append(service.videos, video)
	return nil
}

func (service *videoService) Update(video entity.Video) error {
	for i, v := range service.videos {
		if v.ID == video.ID {
			service.videos[i] = video
			return nil
		}
	}
	return fmt.Errorf("video id %d not exist", video.ID)
}

func (service *videoService) Delete(video entity.Video) error {
	s := service.videos
	for i, v := range s {
		if v.ID == video.ID {
			s[i] = s[len(s)-1]
			s = s[:len(s)-1]
			service.videos = s
			return nil
		}
	}
	return fmt.Errorf("video id %d not exist", video.ID)
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}
