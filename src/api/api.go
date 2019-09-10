package api

import (
  "fmt"
  context "golang.org/x/net/context"

  "github.com/nirajgeorgian/job/src/app"
  db "github.com/nirajgeorgian/job/src/db"
)

type JobServer struct{
  db *db.Database
}
type API struct {
  App *app.App
  Config *Config
  JobServer JobServer
}

func New(a *app.App) (api *API, err error) {
	api = &API{App: a}
	api.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}
  s := JobServer{db: a.Database}

  api.JobServer = s

	return api, nil
}

func (s *JobServer) CreateJob(ctx context.Context, in *CreateJobRequest) (*CreateJobResponse, error) {
  fmt.Println("creating job")

  err := s.db.CreateJob(ctx, in.Job)
  if err != nil {
    return nil, err
  }
  job := in.Job

  return &CreateJobResponse{Job: job}, nil
}
