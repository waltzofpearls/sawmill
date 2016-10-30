package logger

import (
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uber-go/zap"
	"github.com/waltzofpearls/sawmill/app/config"
)

func TestCreateLogger(t *testing.T) {
	var (
		l   *Logger
		err error
	)

	c := &config.Config{}

	c.Application.LogFile = "/i/am/not/here.log"
	l, err = New(c)

	assert.Nil(t, l)
	assert.Error(t, err)

	c.Application.LogFile = "null"
	c.Application.LogLevel = "warn"
	l, err = New(c)

	assert.NoError(t, err)
	assert.Implements(t, (*zap.Logger)(nil), l.Logger)
	assert.Equal(t, zap.WarnLevel, l.Logger.Level())
}

func TestLogLevel(t *testing.T) {
	c := &config.Config{}
	l := &Logger{Config: c}

	c.Application.LogLevel = "info"
	assert.Equal(t, zap.InfoLevel, l.logLevel())

	c.Application.LogLevel = "warn"
	assert.Equal(t, zap.WarnLevel, l.logLevel())

	c.Application.LogLevel = "error"
	assert.Equal(t, zap.ErrorLevel, l.logLevel())

	c.Application.LogLevel = "panic"
	assert.Equal(t, zap.PanicLevel, l.logLevel())

	c.Application.LogLevel = "fatal"
	assert.Equal(t, zap.FatalLevel, l.logLevel())

	c.Application.LogLevel = "debug"
	assert.Equal(t, zap.DebugLevel, l.logLevel())

	c.Application.LogLevel = "anything"
	assert.Equal(t, zap.DebugLevel, l.logLevel())
}

func TestCreateServerLogWriter(t *testing.T) {
	var (
		w   io.Writer
		err error
	)

	c := &config.Config{}
	l := &Logger{Config: c}

	// Negative test
	c.Server.LogFile = "/i/am/not/here.log"
	w, err = l.ServerLogWriter()

	assert.Error(t, err)
	assert.Nil(t, w)

	// Positive tests
	// 1. null logger
	c.Server.LogFile = "null"
	w, err = l.ServerLogWriter()

	assert.NoError(t, err)
	assert.Equal(t, ioutil.Discard, w)

	// 2. stdout logger
	c.Server.LogFile = "stdout"
	w, err = l.ServerLogWriter()

	assert.NoError(t, err)
	assert.Equal(t, os.Stdout, w)

	// 3. file logger
	tf, err := ioutil.TempFile("", "")
	require.NoError(t, err)

	c.Server.LogFile = tf.Name()
	w, err = l.ServerLogWriter()

	assert.NoError(t, err)
	assert.IsType(t, (*os.File)(nil), w)
}

func TestCreateApplicationLogWriter(t *testing.T) {
	var (
		w   zap.WriteSyncer
		err error
	)

	c := &config.Config{}
	l := &Logger{Config: c}

	// Negative test
	c.Application.LogFile = "/i/am/not/here.log"
	w, err = l.ApplicationLogWriter()

	assert.Error(t, err)
	assert.Nil(t, w)

	// Positive tests
	// 1. null logger
	c.Application.LogFile = "null"
	w, err = l.ApplicationLogWriter()

	assert.NoError(t, err)
	assert.IsType(t, (*NullWriteSyncer)(nil), w)

	// 2. stdout logger
	c.Application.LogFile = "stdout"
	w, err = l.ApplicationLogWriter()

	assert.NoError(t, err)
	assert.Equal(t, os.Stdout, w)

	// 3. file logger
	tf, err := ioutil.TempFile("", "")
	require.NoError(t, err)

	c.Application.LogFile = tf.Name()
	w, err = l.ApplicationLogWriter()

	assert.NoError(t, err)
	assert.IsType(t, (*os.File)(nil), w)
}
