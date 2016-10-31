package app

import (
	"os"

	"github.com/urfave/cli"
	"github.com/waltzofpearls/sawmill/app/api"
)

type App struct {
	Api api.ServiceProvider
	Cmd Commander
}

func New() *App {
	return &App{
		Api: api.New(),
		Cmd: &Cmd{App: cli.NewApp()},
	}
}

func (a *App) Run() error {
	a.Cmd.SetName("sawmill")
	a.Cmd.SetVersion("1.0.0")
	a.Cmd.SetUsage("Look up possible malware infected URLs")
	a.Cmd.SetFlags([]cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "config.yml",
			Usage: "Load configuration from `FILE`",
		},
	})
	a.Cmd.SetAction(func(c *cli.Context) error {
		defer func() {
			a.Api.Shutdown()
		}()
		file := c.String("config")
		if err := a.Api.ConfigWith(file); err != nil {
			return err
		}
		a.Api.Serve()
		return nil
	})
	return a.Cmd.Run(os.Args)
}
