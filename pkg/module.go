package pkg

type IModule interface {
	RegisterHTTP()
	RegisterQueue()
	GetInfo() *Module
}

type Module struct {
	Name   string
	Prefix string
}
