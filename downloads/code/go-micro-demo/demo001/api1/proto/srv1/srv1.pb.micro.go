// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv1.proto

package srv1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"

	context "context"

	client "github.com/micro/go-micro/v2/client"

	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Srv1 service

type Srv1Service interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
}

type srv1Service struct {
	c    client.Client
	name string
}

func NewSrv1Service(name string, c client.Client) Srv1Service {
	return &srv1Service{
		c:    c,
		name: name,
	}
}

func (c *srv1Service) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.name, "Srv1.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Srv1 service

type Srv1Handler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error
}

func RegisterSrv1Handler(s server.Server, hdlr Srv1Handler, opts ...server.HandlerOption) error {
	type srv1 interface {
		Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error
	}
	type Srv1 struct {
		srv1
	}
	h := &srv1Handler{hdlr}
	return s.Handle(s.NewHandler(&Srv1{h}, opts...))
}

type srv1Handler struct {
	Srv1Handler
}

func (h *srv1Handler) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.Srv1Handler.Hello(ctx, in, out)
}
