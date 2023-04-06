// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: crea/ms-users/v1/health.proto

package users

import (
	fmt "fmt"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for HealthService service

func NewHealthServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "HealthService.HealthCheck",
			Path:    []string{"/v1/health"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
	}
}

// Client API for HealthService service

type HealthService interface {
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*emptypb.Empty, error)
}

type healthService struct {
	c    client.Client
	name string
}

func NewHealthService(name string, c client.Client) HealthService {
	return &healthService{
		c:    c,
		name: name,
	}
}

func (c *healthService) HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "HealthService.HealthCheck", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HealthService service

type HealthServiceHandler interface {
	HealthCheck(context.Context, *emptypb.Empty, *emptypb.Empty) error
}

func RegisterHealthServiceHandler(s server.Server, hdlr HealthServiceHandler, opts ...server.HandlerOption) error {
	type healthService interface {
		HealthCheck(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error
	}
	type HealthService struct {
		healthService
	}
	h := &healthServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "HealthService.HealthCheck",
		Path:    []string{"/v1/health"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&HealthService{h}, opts...))
}

type healthServiceHandler struct {
	HealthServiceHandler
}

func (h *healthServiceHandler) HealthCheck(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error {
	return h.HealthServiceHandler.HealthCheck(ctx, in, out)
}