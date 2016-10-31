package database

import (
	"fmt"

	riak "github.com/basho/riak-go-client"
	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/logger"
)

type Database struct {
	Cluster *riak.Cluster
	Config  *config.Config
	Logger  *logger.Logger
}

func New(c *config.Config, l *logger.Logger) (*Database, error) {
	var err error
	d := &Database{Config: c, Logger: l}
	if d.Cluster, err = d.connect(); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Database) connect() (*riak.Cluster, error) {
	var nodes []*riak.Node

	db := d.Config.Database
	d.Logger.Info(fmt.Sprintf("Connecting to riak cluster with [%d] nodes...", len(db.Nodes)))

	for _, n := range db.Nodes {
		rn, err := riak.NewNode(&riak.NodeOptions{
			RemoteAddress: n,
		})
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, rn)
	}
	c, err := riak.NewCluster(&riak.ClusterOptions{
		Nodes: nodes,
	})
	if err != nil {
		return nil, err
	}
	if err := c.Start(); err != nil {
		return nil, err
	}
	return c, nil
}

func (d *Database) Close() error {
	if d.Cluster != nil {
		if err := d.Cluster.Stop(); err != nil {
			return err
		}
	}
	return nil
}
