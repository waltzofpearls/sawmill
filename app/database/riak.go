package database

import riak "github.com/basho/riak-go-client"

type RiakCluster interface {
	Execute(riak.Command) error
	Start() error
	Stop() error
}

type Riak struct {
	Nodes   []*riak.Node
	cluster *riak.Cluster
}

func (r *Riak) AddNode(nodeAddress string) error {
	n, err := riak.NewNode(&riak.NodeOptions{
		RemoteAddress: nodeAddress,
	})
	if err == nil {
		r.Nodes = append(r.Nodes, n)
	}
	return err
}

func (r *Riak) FormCluster() error {
	c, err := riak.NewCluster(&riak.ClusterOptions{
		Nodes: r.Nodes,
	})
	if err == nil {
		r.cluster = c
	}
	return err
}

func (r *Riak) Connect() error {
	return r.cluster.Start()
}

func (r *Riak) Cluster() *riak.Cluster {
	return r.cluster
}

func (r *Riak) Close() error {
	if r.cluster != nil {
		if err := r.cluster.Stop(); err != nil {
			return err
		}
	}
	return nil
}
