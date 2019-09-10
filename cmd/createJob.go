package cmd

import (
	"context"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	proto "github.com/nirajgeorgian/job/src/api"
	model "github.com/nirajgeorgian/job/src/model"
)

var JobServiceURI string

func init() {
  createJob.Flags().StringVarP(&JobServiceURI, "jobserviceuri", "u", "", "job service uri (required)")
  createJob.MarkFlagRequired("jobserviceuri")
  viper.BindPFlag("jobserviceuri", createJob.Flags().Lookup("jobserviceuri"))
}

var createJob = &cobra.Command{
  Use: "createJob",
  Short: "create a job with gRPC server on:3000",
	RunE: func(cmd *cobra.Command, args []string) error {
		address     := viper.GetString("jobserviceuri")

		// Set up a connection to the server.
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := proto.NewJobServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		job := model.Job{
			JobName: "dodo duck",
		}
		r, err := c.CreateJob(ctx, &proto.CreateJobRequest{Job: &job})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Job.JobName)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(createJob)
}
