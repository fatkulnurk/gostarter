package example

import (
	"fmt"
	"magicauth/internal/example/delivery"
	"magicauth/internal/example/domain"
	"magicauth/internal/example/repository"
	"magicauth/internal/example/usecase"
	"magicauth/pkg"
)

type MagicLink struct {
	Adapter  *pkg.Adapter
	Delivery *pkg.Delivery
	Usecase  *domain.IUsecase
}

func New(adapter *pkg.Adapter, delivery *pkg.Delivery) pkg.IModule {
	repo := repository.NewRepository(adapter.DB)
	svc := usecase.NewUseCase(repo)

	return &MagicLink{
		Adapter:  adapter,
		Delivery: delivery,
		Usecase:  &svc,
	}
}

func (m *MagicLink) RegisterHTTP() {
	if m.Delivery.HTTP == nil {
		panic("router is nil")
	}

	deliveryHttp := delivery.NewDeliveryHttp(m.Usecase)
	api := m.Delivery.HTTP.Group(fmt.Sprintf("/api/v1/%s", m.GetInfo().Prefix))
	api.Post("/create", deliveryHttp.HandleCreateMagicLink)
	api.Get("/verify", deliveryHttp.HandleVerifyMagicLink)
}

func (m *MagicLink) RegisterQueue() {
	if m.Delivery.Worker == nil {
		panic("queue is nil")
	}

	deliveryQueue := delivery.NewDeliveryQueue(m.Usecase)
	m.Delivery.Worker.HandleFunc(m.GetInfo().Prefix+":send", deliveryQueue.SendMagicLink)
}

func (m *MagicLink) GetInfo() *pkg.Module {
	return &pkg.Module{
		Name:   "Example",
		Prefix: "example",
	}
}
