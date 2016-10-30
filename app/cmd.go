package app

import "github.com/urfave/cli"

type ActionFunc func(*cli.Context) error

type Commander interface {
	SetName(string)
	SetVersion(string)
	SetUsage(string)
	SetFlags([]cli.Flag)
	SetAction(ActionFunc)
	Run([]string) error
}

type Cmd struct {
	*cli.App
}

func (c *Cmd) SetName(name string) {
	c.Name = name
}

func (c *Cmd) SetVersion(version string) {
	c.Version = version
}

func (c *Cmd) SetUsage(usage string) {
	c.Usage = usage
}

func (c *Cmd) SetFlags(flags []cli.Flag) {
	c.Flags = flags
}

func (c *Cmd) SetAction(action ActionFunc) {
	c.Action = action
}
