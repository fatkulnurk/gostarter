package delivery

import (
	"context"
	"fmt"
	"github.com/fatkulnurk/gostarter/internal/example/domain"

	"github.com/hibiken/asynq"
)

type TaskDelivery struct {
	usecase domain.Service
}

func NewDeliveryQueue(usecase domain.Service) *TaskDelivery {
	return &TaskDelivery{usecase: usecase}
}

func (t TaskDelivery) HandleExample(ctx context.Context, task *asynq.Task) error {
	fmt.Printf("HandleExample")
	return nil
}
