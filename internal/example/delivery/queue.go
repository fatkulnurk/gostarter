package delivery

import (
	"context"
	"fmt"
	"magicauth/internal/example/domain"

	"github.com/hibiken/asynq"
)

type QueueDelivery struct {
	usecase domain.IUsecase
}

func NewDeliveryQueue(usecase domain.IUsecase) *QueueDelivery {
	return &QueueDelivery{usecase: usecase}
}

func (d QueueDelivery) HandleTaskExample(ctx context.Context, task *asynq.Task) error {
	fmt.Printf("HandleTaskExample")
	return nil
}
