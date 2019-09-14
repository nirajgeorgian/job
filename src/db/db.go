package db

import (
  "fmt"

  "github.com/pkg/errors"
  "github.com/jinzhu/gorm"

  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
  *gorm.DB
}

func New(config *Config) (*Database, error) {
  db, err := gorm.Open("postgres", config.DatabaseURI)

  if err != nil {
    return nil, errors.Wrap(err, "unable to connect to database")
  }

  if err = db.DB().Ping(); err != nil {
    return nil, errors.Wrap(err, "database ping error")
  } else {
    fmt.Println("Connected to database")
  }

  return &Database{db}, nil
}
