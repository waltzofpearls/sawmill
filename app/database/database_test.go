package database

import (
	"errors"
	"testing"

	riak "github.com/basho/riak-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/logger"
)

type FakeRiak struct{ *Riak }

func (r *FakeRiak) AddNode(nodeAddress string) error { return nil }
func (r *FakeRiak) FormCluster() error               { return nil }
func (r *FakeRiak) Connect() error                   { return nil }
func (r *FakeRiak) Cluster() *riak.Cluster           { return &riak.Cluster{} }

func TestCreateDatabase(t *testing.T) {
	c := &config.Config{}
	c.Database.Nodes = []string{}
	c.Application.LogFile = "null"
	l, _ := logger.New(c)
	d, err := New(&FakeRiak{}, c, l)

	assert.NoError(t, err)
	assert.NotNil(t, d)
	assert.IsType(t, (*riak.Cluster)(nil), d.Cluster)
}

var (
	ErrRiakBadNode       = errors.New("Bad node")
	ErrRiakBadCluster    = errors.New("Bad cluster")
	ErrRiakBadConnection = errors.New("Bad connection")
)

type FakeRiakBadNode struct{ *FakeRiak }

func (r *FakeRiakBadNode) AddNode(nodeAddress string) error { return ErrRiakBadNode }

type FakeRiakBadCluster struct{ *FakeRiak }

func (r *FakeRiakBadCluster) FormCluster() error { return ErrRiakBadCluster }

type FakeRiakBadConnection struct{ *FakeRiak }

func (r *FakeRiakBadConnection) Connect() error { return ErrRiakBadConnection }

func TestConnect(t *testing.T) {
	var err error

	c := &config.Config{}
	c.Database.Nodes = []string{"a_riak_node"}
	c.Application.LogFile = "null"
	l, _ := logger.New(c)
	d := &Database{Config: c, Logger: l}

	d.Riak = &FakeRiakBadNode{}
	err = d.connect()
	assert.Error(t, err)
	assert.EqualError(t, err, ErrRiakBadNode.Error())

	d.Riak = &FakeRiakBadCluster{}
	err = d.connect()
	assert.Error(t, err)
	assert.EqualError(t, err, ErrRiakBadCluster.Error())

	d.Riak = &FakeRiakBadConnection{}
	err = d.connect()
	assert.Error(t, err)
	assert.EqualError(t, err, ErrRiakBadConnection.Error())

	d.Riak = &FakeRiak{}
	err = d.connect()
	assert.NoError(t, err)
}
