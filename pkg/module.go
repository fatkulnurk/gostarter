package pkg

type IModule interface {
	GetInfo() *Module
	RegisterHTTP()
	RegisterTask()
	RegisterSchedule()
}

type Module struct {
	Name   string
	Prefix string
}
