package delivery

import (
	"context"
	"magicauth/internal/example/domain"

	"github.com/hibiken/asynq"
)

type QueueDelivery struct {
	usecase domain.IUsecase
}

func NewDeliveryQueue(usecase domain.IUsecase) *QueueDelivery {
	return &QueueDelivery{usecase: usecase}
}

func (d *QueueDelivery) SendMagicLink(ctx context.Context, task *asynq.Task) error {
	return nil
}
