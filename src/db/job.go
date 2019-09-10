package db

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"

  "github.com/nirajgeorgian/job/src/model"
)

func (db *Database) CreateJob(ctx context.Context, in *model.Job) error {
	tx := db.Begin()
	jobORM, err := in.ToORM(ctx)
	if err != nil {
		return errors.New("error converting input to ORM")
	}
	if err = tx.Create(&jobORM).Error; err != nil {
		return errors.New("error creating job")
	}

	_, err = jobORM.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		return errors.New("error converting back to PB")
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.New("error commiting job create coommit")
	}

	return nil
}
