package main

import (
	"encoding/json"
	"io/ioutil"

	"go.uber.org/dig"
)

const (
	configFile = "config.json"
)

func main() {
	c := dig.New()

	err := c.Provide(func() (*Config, error) {
		out := &Config{}
		bytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bytes, out)
		if err != nil {
			return nil, err
		}

		return out, nil
	})
	if err != nil {
		panic(err)
	}

	err = c.Provide(func(cfg *Config) *Logger {
		return &Logger{level: cfg.Level}
	})
	if err != nil {
		panic(err)
	}

	err = c.Provide(func(logger *Logger) *Server {
		return &Server{logger: logger}
	})
	if err != nil {
		panic(err)
	}

	err = c.Invoke(func(server *Server) {
		server.Listen()
	})
	if err != nil {
		panic(err)
	}
}

type Config struct {
	Level string
}

type Logger struct {
	level string
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	// not implemented
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	// not implemented
}

func (l *Logger) Error(msg string, args ...interface{}) {
	// not implemented
}

type Server struct {
	logger *Logger
}

func (s *Server) Listen() {
	// not implemented
}
