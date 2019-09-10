package cmd

import (
  "log"
  "fmt"
  "net"

  proto "github.com/nirajgeorgian/job/src/api"
  "github.com/nirajgeorgian/job/src/app"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
)

var Port int
var SecretKey string
var DatabaseURI string

var proxyCount string

func init() {
  serveCmd.Flags().IntVarP(&Port, "port", "p", 3000, "port configuration for this application")
  serveCmd.Flags().StringVarP(&SecretKey, "secretkey", "k", "", "secret key for application (required)")
  serveCmd.Flags().StringVarP(&DatabaseURI, "databaseuri", "d", "", "database URI for postgresql (required)")
  serveCmd.Flags().StringVarP(&proxyCount, "proxycount", "x", "", "no. of proxy count")

  serveCmd.MarkFlagRequired("secretkey")
  serveCmd.MarkFlagRequired("databaseuri")

  viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
  viper.BindPFlag("secretkey", serveCmd.Flags().Lookup("secretkey"))
  viper.BindPFlag("databaseuri", serveCmd.Flags().Lookup("databaseuri"))
  viper.BindPFlag("proxycount", serveCmd.Flags().Lookup("proxycount"))
}

var serveCmd = &cobra.Command{
  Use: "serve",
  Short: "serves the gRPC server",
  Long: `start the gRPC server on provided port along with the provided database URI`,
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
	RootCmd.AddCommand(serveCmd)
}
