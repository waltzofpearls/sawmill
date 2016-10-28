package app

import (
	"os"

	"github.com/urfave/cli"
	"github.com/waltzofpearls/sawmill/app/api"
)

type App struct {
	Api api.ServiceProvider
	Cmd *cli.App
}

func New() *App {
	return &App{
		Api: api.New(),
		Cmd: cli.NewApp(),
	}
}

func (a *App) Run() {
	a.Cmd.Name = "sawmill"
	a.Cmd.Version = "1.0.0"
	a.Cmd.Usage = "Look up possible malware infected URLs"
	a.Cmd.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "config.yml",
			Usage: "Load configuration from `FILE`",
		},
	}
	a.Cmd.Action = func(c *cli.Context) error {
		file := c.String("config")
		a.Api.ConfigWith(file)
		a.Api.Serve()
		return nil
	}
	a.Cmd.Run(os.Args)
}
