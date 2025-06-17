package example

import (
	"fmt"
	"github.com/fatkulnurk/gostarter/internal/helloworld/delivery"
	"github.com/fatkulnurk/gostarter/internal/helloworld/domain"
	"github.com/fatkulnurk/gostarter/internal/helloworld/repository"
	"github.com/fatkulnurk/gostarter/internal/helloworld/service"
	"github.com/fatkulnurk/gostarter/pkg/module"
	"github.com/fatkulnurk/gostarter/shared/infrastructure"
	"github.com/hibiken/asynq"
)

type Module struct {
	Adapter  *infrastructure.Adapter
	Delivery *infrastructure.Delivery
	Service  *domain.Service
}

func New(adapter *infrastructure.Adapter, delivery *infrastructure.Delivery) module.IModule {
	repo := repository.NewRepository(adapter.DB.Sql)
	svc := service.NewService(repo)

	return &Module{
		Adapter:  adapter,
		Delivery: delivery,
		Service:  &svc,
	}
}

func (m *Module) GetInfo() *module.Module {
	return &module.Module{
		Name:   "HelloWorld",
		Prefix: "hello-world",
	}
}

func (m *Module) RegisterHTTP() {
	if m.Delivery.HTTP == nil {
		panic("router is nil")
	}

	deliveryHttp := delivery.NewDeliveryHttp(m.Service)

	// app
	app := m.Delivery.HTTP.Group(fmt.Sprintf("/%s", m.GetInfo().Prefix))
	app.Get("", deliveryHttp.HandleHelloWorld)

	// api
	api := m.Delivery.HTTP.Group(fmt.Sprintf("/api/v1/%s", m.GetInfo().Prefix))
	api.Get("", deliveryHttp.HandleExampleApi)
}

func (m *Module) RegisterTask() {
	if m.Delivery.Task == nil {
		panic("task is nil")
	}

	deliveryTask := delivery.NewDeliveryQueue(m.Service)
	m.Delivery.Task.HandleFunc(m.GetInfo().Prefix+":example", deliveryTask.HandleExample)

	deliverySchedule := delivery.NewScheduleDelivery(m.Service)
	m.Delivery.Task.HandleFunc(m.GetInfo().Prefix+":schedule::example", deliverySchedule.HandleTaskScheduleExample)
}

func (m *Module) RegisterSchedule() {
	if m.Delivery.Schedule == nil {
		panic("schedule is nil")
	}

	// register schedule
	entryID, err := m.Delivery.Schedule.Register("*/1 * * * *", asynq.NewTask(m.GetInfo().Prefix+":schedule::example", nil))
	if err != nil {
		panic(err)
	}
	fmt.Println("Registered schedule:", entryID)
}
