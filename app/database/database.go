package database

import (
	"fmt"

	riak "github.com/basho/riak-go-client"
	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/logger"
)

type Adapter interface {
	AddNode(string) error
	FormCluster() error
	Connect() error
	Cluster() *riak.Cluster
	Close() error
}

type Database struct {
	Cluster RiakCluster
	Config  *config.Config
	Logger  *logger.Logger
	Riak    Adapter
}

func New(r Adapter, c *config.Config, l *logger.Logger) (*Database, error) {
	d := &Database{Config: c, Logger: l, Riak: r}

	if err := d.connect(); err != nil {
		return nil, err
	}

	d.Cluster = d.Riak.Cluster()
	return d, nil
}

func (d *Database) connect() error {
	db := d.Config.Database
	d.Logger.Info(fmt.Sprintf("Connecting to riak cluster with [%d] nodes...", len(db.Nodes)))

	for _, n := range db.Nodes {
		if err := d.Riak.AddNode(n); err != nil {
			return err
		}
	}
	if err := d.Riak.FormCluster(); err != nil {
		return err
	}
	if err := d.Riak.Connect(); err != nil {
		return err
	}
	return nil
}

func (d *Database) Close() error {
	return d.Riak.Close()
}
