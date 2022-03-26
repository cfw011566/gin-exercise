package service

import (
	"example/golang-gin-poc/entity"
	"example/golang-gin-poc/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(entity.Video) entity.Video
	Delete(entity.Video) entity.Video
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

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) Update(video entity.Video) entity.Video {
	service.videoRepository.Update(video)
	return video
}

func (service *videoService) Delete(video entity.Video) entity.Video {
	service.videoRepository.Delete(video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}
