package logger

import "github.com/waltzofpearls/sawmill/app/config"

type Logger struct{}

func New(c *config.Config) *Logger {
	return &Logger{}
}
