package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/logger"
)

type ServiceProvider interface {
	ConfigWith(string) error
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

func (a *Api) ConfigWith(filePath string) error {
	var err error
	a.Config, err = config.New(filePath)
	if err != nil {
		return err
	}
	a.Logger = logger.New(a.Config)
	return nil
}

func (a *Api) Serve() {
	a.Route("/urlinfo/1", &Version1{})
	sr := &Subroute{}
	a.Router.NotFoundHandler = http.HandlerFunc(sr.JsonNotFoundHandler)

	http.ListenAndServe(
		a.Config.Listen.Address,
		AttachMiddleware(
			a.Router,
			Catch404Handler(),
			LoggingHandler(os.Stdout),
		),
	)
}

func (a *Api) Route(path string, sr Subrouter) {
	r := a.Router.PathPrefix(path).Subrouter()
	sr.ConfigWith(r, a.Config, a.Logger)
	sr.Handle()
}
