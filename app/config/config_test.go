package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var badYamlFixture string = `
foo: bar
    - baz: qux
`

var goodYamlFixture string = `
server:
  listen: 0.0.0.0:8000
  log_file: stdout

application:
  log_file: stdout
  log_level: debug

database:
  nodes:
    - 127.0.0.1:8087
`

func TestCreateConfig(t *testing.T) {
	var (
		filePath string
		err      error
	)

	// Negative case: yaml file doesn't exist
	_, err = New("i_am_no_here.yml")

	assert.Error(t, err)

	// Negative case: bad yaml config format
	filePath = helpeCreateTestFile(t, "bad_yaml", badYamlFixture)
	defer os.Remove(filePath)
	_, err = New(filePath)

	assert.Error(t, err)

	// Positive case: success
	filePath = helpeCreateTestFile(t, "good_yaml", goodYamlFixture)
	defer os.Remove(filePath)
	expected := &Config{
		Server: struct {
			Listen  string
			LogFile string `yaml:"log_file"`
		}{
			Listen:  "0.0.0.0:8000",
			LogFile: "stdout",
		},
		Application: struct {
			LogFile  string `yaml:"log_file"`
			LogLevel string `yaml:"log_level"`
		}{
			LogFile:  "stdout",
			LogLevel: "debug",
		},
		Database: struct {
			Nodes []string
		}{
			Nodes: []string{"127.0.0.1:8087"},
		},
	}
	actual, err := New(filePath)

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
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
