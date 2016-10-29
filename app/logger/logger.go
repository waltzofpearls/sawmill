package logger

import (
	"errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/uber-go/zap"
	"github.com/waltzofpearls/sawmill/app/config"
)

type Logger struct {
	config *config.Config
	file   *os.File
	zap.Logger
}

func New(c *config.Config) (*Logger, error) {
	l := &Logger{config: c}

	w, err := l.ApplicationLogWriter()
	if err != nil {
		return nil, err
	}
	l.Logger = zap.New(
		zap.NewTextEncoder(),
		zap.Output(w),
		l.logLevel(),
	)
	return l, nil
}

func (l *Logger) logLevel() zap.Level {
	var lvl zap.Level

	switch l.config.Application.LogLevel {
	case "info":
		lvl = zap.InfoLevel
	case "warn":
		lvl = zap.WarnLevel
	case "error":
		lvl = zap.ErrorLevel
	case "panic":
		lvl = zap.PanicLevel
	case "fatal":
		lvl = zap.FatalLevel
	case "debug":
		fallthrough
	default:
		lvl = zap.DebugLevel
	}
	return lvl
}

func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}

func (l *Logger) ServerLogWriter() (io.Writer, error) {
	var (
		w   io.Writer
		err error
	)

	logFile := l.config.Server.LogFile
	switch logFile {
	case "null":
		w = ioutil.Discard
	case "stdout":
		w = os.Stdout
	default:
		w, err = l.logFileWriter(logFile)
		if err != nil {
			return nil, err
		}
	}
	return w, nil
}

var ErrConvertZapWriteSyncer = errors.New("Cannot convert application log writer to zap.WriteSyncer.")

func (l *Logger) ApplicationLogWriter() (zap.WriteSyncer, error) {
	var (
		w   io.Writer
		err error
	)

	logFile := l.config.Application.LogFile
	switch logFile {
	case "null":
		w = &NullWriteSyncer{ioutil.Discard}
	case "stdout":
		w = os.Stdout
	default:
		w, err = l.logFileWriter(logFile)
		if err != nil {
			return nil, err
		}
	}

	ws, ok := w.(zap.WriteSyncer)
	if !ok {
		return nil, ErrConvertZapWriteSyncer
	}
	return ws, nil
}

func (l *Logger) logFileWriter(logFile string) (*os.File, error) {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}
	l.file = f
	return f, nil
}

type NullWriteSyncer struct {
	io.Writer
}

func (ws *NullWriteSyncer) Sync() error {
	return nil
}
