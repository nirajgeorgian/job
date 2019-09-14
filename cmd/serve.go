package cmd

import (
  "log"

  proto "github.com/nirajgeorgian/job/src/api"
  "github.com/nirajgeorgian/job/src/app"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var Port int
var SecretKey string
var DatabaseURI string

var proxyCount string

func init() {
  serveCmd.Flags().IntVarP(&Port, "port", "p", 3000, "port configuration for this application")
  serveCmd.Flags().StringVarP(&SecretKey, "secretkey", "k", "dododuckN9", "secret key for application (required)")
  serveCmd.Flags().StringVarP(&DatabaseURI, "databaseuri", "d", "", "database URI for postgresql (required)")
  serveCmd.Flags().StringVarP(&proxyCount, "proxycount", "x", "", "no. of proxy count")

  // serveCmd.MarkFlagRequired("secretkey")
  // serveCmd.MarkFlagRequired("databaseuri")

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

    if err := proto.ListenGRPC(api, port); err != nil {
      log.Fatalf("Failed to serve: %v\n", err)
    }

		return nil
  },
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
