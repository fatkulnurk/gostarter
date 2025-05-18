package usecase

import "magicauth/internal/magiclink/domain"

type UseCase struct {
	repo domain.IRepository
}

func NewUseCase(repo domain.IRepository) domain.IUsecase {
	return &UseCase{repo: repo}
}
