package db

import (
	log "github.com/Sirupsen/logrus"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"github.com/satori/go.uuid"

  "github.com/nirajgeorgian/job/src/model"
)

// CreateJob :- create's an job on job dashboard
func (db *Database) CreateJob(ctx context.Context, in *model.Job) (*model.Job, error) {
	log.Println("database: CreateJob")

	tx := db.Begin()
	jobORM, err := in.ToORM(ctx)
	if err != nil {
		return nil, errors.New("error converting input to ORM")
	}
	u1 := uuid.NewV4()
	jobORM.JobId = u1.String()

	if err = tx.Create(&jobORM).Error; err != nil {
		return nil, errors.New("error creating job")
	}

	job, err := jobORM.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("error converting back to PB")
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error commiting job create coommit")
	}

	return &job, nil
}
