package api

import (
	"net/http"

	"github.com/gorilla/mux"
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
	Router *mux.Router
}

func New() *Api {
	return &Api{
		Router: mux.NewRouter(),
	}
}

func (a *Api) ConfigWith(file string) {
	a.Config = config.New(file)
	a.Logger = logger.New(a.Config)

	a.Route("/urlinfo/1", &Version1{})
}

func (a *Api) Route(path string, sr Subrouter) {
	r := a.Router.PathPrefix(path).Subrouter()
	sr.ConfigWith(r, a.Config)
	sr.Handle()
}

func (a *Api) Serve() {
	http.ListenAndServe(
		// a.Config.Listen.Address,
		":9000",
		a.Router,
	)
}
