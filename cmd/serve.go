package cmd

import (
  "log"
  "fmt"
  "net"

  proto "github.com/nirajgeorgian/job/src/api"
  "github.com/nirajgeorgian/job/src/app"

  "github.com/spf13/cobra"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
)

var serveCmd = &cobra.Command{
  Use: "serve",
  Short: "serves the gRPC server",
  RunE: func(cmd *cobra.Command, args []string) error {
    a, err := app.New()
    if err != nil {
			return err
		}
		defer a.Close()

    api, err := proto.New(a)
    if err != nil {
			return err
		}

		// serveRpc(api)
    port := api.Config.Port
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
      log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterJobServiceServer(grpcServer, &api.JobServer)
    reflection.Register(grpcServer)

    fmt.Printf("starting to listen on tcp: %q\n", lis.Addr().String())
    if err := grpcServer.Serve(lis); err != nil {
      log.Fatalf("Failed to serve: %v\n", err)
    }

		return nil
  },
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
