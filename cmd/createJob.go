package cmd

import (
	"context"
	"log"
	"time"
	"fmt"

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
		fmt.Println(address)

		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("could not connect to service: %v", err)
		}
		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		job := model.Job{
			JobId: "1",
			JobName: "dodo duck",
		}

		r, err := proto.NewJobServiceClient(conn).CreateJob(ctx, &proto.CreateJobReq{Job: &job})
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
