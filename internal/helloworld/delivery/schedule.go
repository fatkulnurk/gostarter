package delivery

import (
	"context"
	"fmt"
	"time"

	"github.com/fatkulnurk/gostarter/internal/example/domain"
	"github.com/hibiken/asynq"
)

type ScheduleDelivery struct {
	service domain.Service
}

func NewScheduleDelivery(service domain.Service) *ScheduleDelivery {
	return &ScheduleDelivery{
		service: service,
	}
}

func (s ScheduleDelivery) HandleTaskScheduleExample(ctx context.Context, task *asynq.Task) error {
	fmt.Printf("HandleTaskScheduleExample")
	fmt.Println("Current time:", time.Now().Format(time.RFC3339))
	return nil
}
