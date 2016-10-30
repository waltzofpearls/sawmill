package api

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/waltzofpearls/sawmill/app/config"
	"github.com/waltzofpearls/sawmill/app/logger"
)

func TestCreateApi(t *testing.T) {
	a := New()

	assert.NotNil(t, a.Router)
	assert.Nil(t, a.Config)
	assert.Nil(t, a.Logger)
	assert.Nil(t, a.Writer)
}

var yamlConfigBadLogger = `
application:
  log_file: /i/am/a/troubl/maker.log
`

var yamlConfigBadWriter = `
server:
  log_file: /i/am/a/trouble/maker.log

application:
  log_file: "null"
`

var yamlConfigAllGood = `
server:
  log_file: "null"

application:
  log_file: "null"
`

func TestConfigWith(t *testing.T) {
	var (
		filePath string
		err      error
	)

	r := mux.NewRouter()
	a := &Api{Router: r}

	err = a.ConfigWith("/i/am/not/here.yml")
	assert.Error(t, err)
	assert.IsType(t, (*os.PathError)(nil), err)

	filePath = helpeCreateTestFile(t, "bad_logger", yamlConfigBadLogger)
	err = a.ConfigWith(filePath)
	assert.Error(t, err)
	assert.IsType(t, (*os.PathError)(nil), err)

	filePath = helpeCreateTestFile(t, "bad_writer", yamlConfigBadWriter)
	err = a.ConfigWith(filePath)
	assert.Error(t, err)
	assert.IsType(t, (*os.PathError)(nil), err)

	filePath = helpeCreateTestFile(t, "all_good", yamlConfigAllGood)
	err = a.ConfigWith(filePath)
	assert.NoError(t, err)
}

func helpeCreateTestFile(t *testing.T, prefix, fixture string) string {
	tf, err := ioutil.TempFile("", prefix)
	require.Nil(t, err)

	_, err = tf.WriteString(fixture)
	require.Nil(t, err)

	err = tf.Close()
	require.Nil(t, err)

	return tf.Name()
}

type FakeBilboBagginsSubroute struct{ Subroute }

func (sr *FakeBilboBagginsSubroute) Handle() {}

func TestRoute(t *testing.T) {
	a := New()
	a.Config = &config.Config{}
	a.Logger = &logger.Logger{}

	sr := &FakeBilboBagginsSubroute{}
	a.Route("/bilbobaggins", sr)

	assert.NotNil(t, sr.Router)
	assert.NotNil(t, sr.Config)
	assert.NotNil(t, sr.Logger)
}
