package example

import (
	"fmt"
	"magicauth/internal/example/delivery"
	"magicauth/internal/example/domain"
	"magicauth/internal/example/repository"
	"magicauth/internal/example/usecase"
	"magicauth/pkg"
)

type Module struct {
	Adapter  *pkg.Adapter
	Delivery *pkg.Delivery
	Usecase  *domain.IUsecase
}

func New(adapter *pkg.Adapter, delivery *pkg.Delivery) pkg.IModule {
	repo := repository.NewRepository(adapter.DB)
	svc := usecase.NewUseCase(repo)

	return &Module{
		Adapter:  adapter,
		Delivery: delivery,
		Usecase:  &svc,
	}
}

func (m *Module) RegisterHTTP() {
	if m.Delivery.HTTP == nil {
		panic("router is nil")
	}

	deliveryHttp := delivery.NewDeliveryHttp(m.Usecase)

	// app
	app := m.Delivery.HTTP.Group(fmt.Sprintf("/%s", m.GetInfo().Prefix))
	app.Get("", deliveryHttp.HandleHelloWorld)

	// api
	api := m.Delivery.HTTP.Group(fmt.Sprintf("/api/v1/%s", m.GetInfo().Prefix))
	api.Get("", deliveryHttp.HandleExampleApi)
}

func (m *Module) RegisterQueue() {
	if m.Delivery.Worker == nil {
		panic("queue is nil")
	}

	deliveryQueue := delivery.NewDeliveryQueue(m.Usecase)
	m.Delivery.Worker.HandleFunc(m.GetInfo().Prefix+":send", deliveryQueue.SendMagicLink)
}

func (m *Module) GetInfo() *pkg.Module {
	return &pkg.Module{
		Name:   "Module",
		Prefix: "example",
	}
}
