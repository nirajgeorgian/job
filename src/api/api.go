package api

import (
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/nirajgeorgian/job/src/app"
	db "github.com/nirajgeorgian/job/src/db"
)

type JobServer struct {
	db *db.Database
}
type API struct {
	App       *app.App
	Config    *Config
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

func ListenGRPC(api *API, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	RegisterJobServiceServer(grpcServer, &api.JobServer)
	reflection.Register(grpcServer)

	fmt.Printf("starting to listen on tcp: %q\n", lis.Addr().String())

	return grpcServer.Serve(lis)
}

func (s *JobServer) CreateJob(ctx context.Context, in *CreateJobReq) (*CreateJobRes, error) {
	fmt.Println("creating job")

	err := s.db.CreateJob(ctx, in.Job)
	if err != nil {
		return nil, err
	}
	job := in.Job

	return &CreateJobRes{Job: job}, nil
}
