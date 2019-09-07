package app

import (
  "os"
  "time"

  "github.com/sirupsen/logrus"
  "github.com/nirajgeorgian/job/src/db"
)

type App struct {
  Config *Config
  Database *db.Database
}

func (a *App) NewContext() *Context {
  log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout

  return &Context{
		Logger:   log,
		Database: a.Database,
	}
}

func New() (app *App, err error) {
  app = &App{}
  app.Config, err = InitConfig()
  if err != nil {
		return nil, err
	}

  dbConfig, err := db.InitConfig()
	if err != nil {
		return nil, err
	}

  app.Database, err = db.New(dbConfig)
  if err != nil {
		return nil, err
	}

  return app, err
}

func (a *App) Close() error {
	return a.Database.Close()
}
