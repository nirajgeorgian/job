package app

import (
  "github.com/pkg/errors"

  "github.com/nirajgeorgian/job/src/model"

  "github.com/nirajgeorgian/job/src/db"
)

type App struct {
  Config *Config
  Database *db.Database
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

  // migrate database on startup
  if err := app.Database.AutoMigrate(
		&model.JobORM{},
		&model.SallaryORM{},
		&model.AttachmentORM{},
		).Error; err != nil {
		return nil, errors.Wrap(err, "unable to automatically migrate migrations table")
	}

  return app, err
}

func (a *App) Close() error {
	return a.Database.Close()
}
