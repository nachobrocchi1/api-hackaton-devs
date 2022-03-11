package service

import (
	"api-hackaton-devs/entity"
	"api-hackaton-devs/mapper"
	"api-hackaton-devs/repository"
	"context"
)

type Service interface {
	GetHackatonWithBestDevs(ctx context.Context, req *entity.Request) (*entity.Response, error)
}

type service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetHackatonWithBestDevs(ctx context.Context, req *entity.Request) (*entity.Response, error) {
	hackatons, err := s.repo.GetHackatonWithBestDevs()
	if err != nil {
		return nil, err
	}
	return mapper.MapModelResponseToEntityResponse(hackatons), nil
}
