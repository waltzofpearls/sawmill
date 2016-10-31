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
	Config *config.Config
	appLog *os.File
	srvLog *os.File
	zap.Logger
}

func New(c *config.Config) (*Logger, error) {
	l := &Logger{Config: c}

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

	switch l.Config.Application.LogLevel {
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
	if l.srvLog != nil {
		l.srvLog.Close()
	}
	if l.appLog != nil {
		l.appLog.Close()
	}
}

func (l *Logger) ServerLogWriter() (io.Writer, error) {
	var (
		w   io.Writer
		err error
	)

	logFile := l.Config.Server.LogFile
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
		l.srvLog = w.(*os.File)
	}
	return w, nil
}

var ErrConvertZapWriteSyncer = errors.New("Cannot convert application log writer to zap.WriteSyncer.")

func (l *Logger) ApplicationLogWriter() (zap.WriteSyncer, error) {
	var (
		w   io.Writer
		err error
	)

	logFile := l.Config.Application.LogFile
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
		l.appLog = w.(*os.File)
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
	return f, nil
}

type NullWriteSyncer struct {
	io.Writer
}

func (ws *NullWriteSyncer) Sync() error {
	return nil
}
