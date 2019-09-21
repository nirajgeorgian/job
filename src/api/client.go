package api

import (
	"context"

	"google.golang.org/grpc"

	"github.com/nirajgeorgian/job/src/model"
)

// Client :- account client structure
type Client struct {
	conn    *grpc.ClientConn
	service JobServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := NewJobServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) CreateJob(ctx context.Context, job model.Job) (*model.Job, error) {
	r, err := c.service.CreateJob(
		ctx,
		&CreateJobReq{Job: &job},
	)
	if err != nil {
		return nil, err
	}
	return r.Job, nil
}
