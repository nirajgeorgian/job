package main

import (
  "net"
  "fmt"
  "time"
  "os"

  proto "github.com/nirajgeorgian/job/src/proto"
  context "golang.org/x/net/context"
  "github.com/sirupsen/logrus"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
)

type server struct{}
var log *logrus.Logger
const (
	listenPort  = "3030"
)

func init() {
	log = logrus.New()
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
}

func main() {
  port := listenPort
  lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  s := grpc.NewServer()
  proto.RegisterJobServiceServer(s, &server{})
  reflection.Register(s)
  log.Infof("starting to listen on tcp: %q", lis.Addr().String())

  if err := s.Serve(lis); err != nil {
    log.Fatalf("Failed to serve: %v", err)
  }
}

func (s *server) CreateJob(ctx context.Context, in *proto.CreateJobRequest) (*proto.CreateJobResponse, error) {
  job := in.Job

  return &proto.CreateJobResponse{Job: job}, nil
}
