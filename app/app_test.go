package app

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli"
	"github.com/waltzofpearls/sawmill/app/api"
)

type FakeApi struct{}

func (a *FakeApi) ConfigWith(filePath string) error { return nil }

func (a *FakeApi) Serve()    {}
func (a *FakeApi) Shutdown() {}

type FakeCmd struct {
	*cli.App
	mock.Mock
}

func (c *FakeCmd) SetName(name string)         {}
func (c *FakeCmd) SetVersion(version string)   {}
func (c *FakeCmd) SetUsage(usage string)       {}
func (c *FakeCmd) SetFlags(flags []cli.Flag)   {}
func (c *FakeCmd) SetAction(action ActionFunc) {}
func (c *FakeCmd) Run(arguments []string) error {
	args := c.Called(arguments)
	return args.Error(0)
}

func TestCreateApp(t *testing.T) {
	a := New()

	assert.NotNil(t, a)
	assert.IsType(t, (*App)(nil), a)
	assert.IsType(t, (*api.Api)(nil), a.Api)
	assert.IsType(t, (*Cmd)(nil), a.Cmd)
}

func TestAppRun(t *testing.T) {
	var err error

	apiMock := &FakeApi{}
	cmdMock := &FakeCmd{App: cli.NewApp()}
	cmdMock.On("Run", []string{"good"}).Return(nil)
	cmdMock.On("Run", []string{"bad"}).Return(errors.New("error"))

	app := &App{
		Api: apiMock,
		Cmd: cmdMock,
	}

	os.Args = []string{"good"}
	err = app.Run()
	assert.NoError(t, err)

	os.Args = []string{"bad"}
	err = app.Run()
	assert.Error(t, err)
}
