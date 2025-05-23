package delivery

import (
	"context"
	"fmt"
	"github.com/fatkulnurk/gostarter/internal/helloworld/domain"

	"github.com/hibiken/asynq"
)

type TaskDelivery struct {
	usecase domain.IUsecase
}

func NewDeliveryQueue(usecase domain.IUsecase) *TaskDelivery {
	return &TaskDelivery{usecase: usecase}
}

func (t TaskDelivery) HandleExample(ctx context.Context, task *asynq.Task) error {
	fmt.Printf("HandleExample")
	return nil
}
