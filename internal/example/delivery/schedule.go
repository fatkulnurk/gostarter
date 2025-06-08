package delivery

import (
	"context"
	"fmt"
	"time"

	"github.com/fatkulnurk/gostarter/internal/example/domain"
	"github.com/hibiken/asynq"
)

type ScheduleDelivery struct {
	usecase domain.Service
}

func NewScheduleDelivery(usecase domain.Service) *ScheduleDelivery {
	return &ScheduleDelivery{
		usecase: usecase,
	}
}

func (s ScheduleDelivery) HandleTaskScheduleExample(ctx context.Context, task *asynq.Task) error {
	fmt.Printf("HandleTaskScheduleExample")
	fmt.Println("Current time:", time.Now().Format(time.RFC3339))
	return nil
}
