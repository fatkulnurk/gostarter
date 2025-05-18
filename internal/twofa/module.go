package twofa

import (
	"magicauth/pkg"
)

type TwoFA struct {
	Adapter  *pkg.Adapter
	Delivery *pkg.Delivery
}

func New(adapter *pkg.Adapter, delivery *pkg.Delivery) pkg.IModule {
	return &TwoFA{
		Adapter:  adapter,
		Delivery: delivery,
	}
}

func (t TwoFA) RegisterHTTP() {
	if t.Delivery.HTTP == nil {
		panic("router is nil")
	}
}

func (t TwoFA) RegisterQueue() {
	if t.Delivery.Worker == nil {
		panic("queue is nil")
	}
}

func (t TwoFA) GetInfo() *pkg.Module {
	return &pkg.Module{
		Name:   "2FA",
		Prefix: "2fa",
	}
}
