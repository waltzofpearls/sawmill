package api

import (
	"fmt"

	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/logger"
)

type ServiceProvider interface {
	ConfigWith(string)
	Serve()
}

type Api struct {
	Config *config.Config
	Logger *logger.Logger
}

func New() *Api {
	return &Api{}
}

func (a *Api) ConfigWith(file string) {
	a.Config = config.New(file)
	a.Logger = logger.New(a.Config)

}

func (a *Api) Serve() {
	fmt.Println("Hello world!")
	select {}
}
