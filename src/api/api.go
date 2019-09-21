package api

import (
	"fmt"
	"net"
	"time"

	log "github.com/Sirupsen/logrus"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc/reflection"

	app "github.com/nirajgeorgian/job/src/app"
	db "github.com/nirajgeorgian/job/src/db"
)

type JobServer struct {
	db *db.Database
}

// API :- base API strcture
type API struct {
	App       *app.App
	Config    *Config
	JobServer JobServer
}

// New :- new api instance
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

	recoveryFunc := func(p interface{}) (err error) {
		log.Errorf("recovered from panic: %v", p)
		return fmt.Errorf("recovered from panic: %v", p)
	}
	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryFunc),
	}
	grpcServer := grpc.NewServer(
		grpc.ConnectionTimeout(time.Minute*30),
		grpc.MaxRecvMsgSize(1024*1024*128),
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(recoveryOpts...),
		),
	)

	RegisterJobServiceServer(grpcServer, &api.JobServer)
	reflection.Register(grpcServer)

	log.Printf("starting HTTP/2 gRPC API server: %q\n", lis.Addr().String())

	return grpcServer.Serve(lis)
}

// CreateJob :- CreateJob rpc network call
func (s *JobServer) CreateJob(ctx context.Context, in *CreateJobReq) (*CreateJobRes, error) {
	log.Println("server: CreateJob")

	job, err := s.db.CreateJob(ctx, in.Job)
	if err != nil {
		return nil, err
	}

	return &CreateJobRes{Job: job}, nil
}
