package delivery

import (
	"context"
	"fmt"
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
	return nil
}
