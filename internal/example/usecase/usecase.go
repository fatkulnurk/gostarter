package usecase

import "magicauth/internal/example/domain"

type UseCase struct {
	repo domain.IRepository
}

func NewUseCase(repo domain.IRepository) domain.IUsecase {
	return &UseCase{repo: repo}
}
