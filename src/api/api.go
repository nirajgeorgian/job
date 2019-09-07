package api

import (
  context "golang.org/x/net/context"

  "github.com/nirajgeorgian/job/src/app"
)

type JobServer struct{}
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
  s := JobServer{}

  api.JobServer = s

	return api, nil
}

func (s *JobServer) CreateJob(ctx context.Context, in *CreateJobRequest) (*CreateJobResponse, error) {
  job := in.Job

  return &CreateJobResponse{Job: job}, nil
}
