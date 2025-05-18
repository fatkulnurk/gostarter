package usecase

import "github.com/fatkulnurk/gostarter/internal/helloworld/domain"

type UseCase struct {
	repo domain.IRepository
}

func NewUseCase(repo domain.IRepository) domain.IUsecase {
	return &UseCase{repo: repo}
}
