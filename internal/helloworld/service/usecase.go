package service

import "github.com/fatkulnurk/gostarter/internal/helloworld/domain"

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) domain.Service {
	return &Service{repo: repo}
}
