package delivery

import (
	"context"
	"fmt"
	"time"

	"github.com/fatkulnurk/gostarter/internal/example/domain"
	"github.com/hibiken/asynq"
)

type ScheduleDelivery struct {
	usecase domain.IUsecase
}

func NewScheduleDelivery(usecase domain.IUsecase) *ScheduleDelivery {
	return &ScheduleDelivery{
		usecase: usecase,
	}
}

func (s ScheduleDelivery) HandleTaskScheduleExample(ctx context.Context, task *asynq.Task) error {
	fmt.Printf("HandleTaskScheduleExample")
	fmt.Println("Current time:", time.Now().Format(time.RFC3339))
	return nil
}
