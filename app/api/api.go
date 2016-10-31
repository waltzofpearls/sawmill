package api

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/database"
	"github.com/waltzofpearls/sawmill/app/logger"
)

type ServiceProvider interface {
	ConfigWith(string) error
	Serve()
	Shutdown()
}

type Api struct {
	Config   *config.Config
	Database *database.Database
	Logger   *logger.Logger
	Router   *mux.Router
	Writer   io.Writer
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
	if a.Writer, err = a.Logger.ServerLogWriter(); err != nil {
		return err
	}
	if a.Database, err = database.New(a.Config, a.Logger); err != nil {
		return err
	}

	a.Route("/urlinfo/1", &Version1{})

	return nil
}

func (a *Api) Route(path string, sr Subrouter) {
	r := a.Router.PathPrefix(path).Subrouter()
	sr.ConfigWith(r, a.Database, a.Config, a.Logger)
	sr.Handle()
}

func (a *Api) Serve() {
	a.Logger.Info("Start serving http request...")

	http.ListenAndServe(
		a.Config.Server.Listen,
		AttachMiddleware(
			a.Router,
			Catch404Handler(),
			LoggingHandler(a.Writer),
		),
	)
}

func (a *Api) Shutdown() {
	a.Logger.Close()
	a.Database.Close()
}
