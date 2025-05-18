package delivery

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"magicauth/internal/example/domain"
)

type ScheduleDelivery struct {
	usecase domain.IUsecase
}

func NewScheduleDelivery(usecase domain.IUsecase) *ScheduleDelivery {
	return &ScheduleDelivery{
		usecase: usecase,
	}
}

func (d QueueDelivery) HandleTaskScheduleExample(ctx context.Context, task *asynq.Task) error {
	fmt.Printf("HandleTaskScheduleExample")
	return nil
}
