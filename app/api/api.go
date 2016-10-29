package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/logger"
)

type ServiceProvider interface {
	ConfigWith(string) error
	Serve() error
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

	if a.Config, err = config.New(filePath); err != nil {
		return err
	}
	if a.Logger, err = logger.New(a.Config); err != nil {
		return err
	}
	return nil
}

func (a *Api) Serve() error {
	a.Route("/urlinfo/1", &Version1{})

	a.Logger.Info("Blah Blah Blah")
	a.Logger.Debug("Blah Blah Blah")

	w, err := a.Logger.ServerLogWriter()
	if err != nil {
		return err
	}

	http.ListenAndServe(
		a.Config.Server.Listen,
		AttachMiddleware(
			a.Router,
			Catch404Handler(),
			LoggingHandler(w),
		),
	)

	return nil
}

func (a *Api) Route(path string, sr Subrouter) {
	r := a.Router.PathPrefix(path).Subrouter()
	sr.ConfigWith(r, a.Config, a.Logger)
	sr.Handle()
}
