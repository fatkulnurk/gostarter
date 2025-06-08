package delivery

import (
	"context"
	"fmt"
	"github.com/fatkulnurk/gostarter/internal/helloworld/domain"

	"github.com/hibiken/asynq"
)

type TaskDelivery struct {
	service domain.Service
}

func NewDeliveryQueue(service domain.Service) *TaskDelivery {
	return &TaskDelivery{service: service}
}

func (t TaskDelivery) HandleExample(ctx context.Context, task *asynq.Task) error {
	fmt.Printf("HandleExample")
	return nil
}
