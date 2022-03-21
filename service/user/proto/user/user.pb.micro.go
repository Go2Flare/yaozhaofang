// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for User service

type UserService interface {
	MicroGetUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdateUserName(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*UpdateResp, error)
	UploadAvatar(ctx context.Context, in *UploadReq, opts ...client.CallOption) (*UploadResp, error)
	AuthUpdate(ctx context.Context, in *AuthReq, opts ...client.CallOption) (*AuthResp, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) MicroGetUser(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.MicroGetUser", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateUserName(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*UpdateResp, error) {
	req := c.c.NewRequest(c.name, "User.UpdateUserName", in)
	out := new(UpdateResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UploadAvatar(ctx context.Context, in *UploadReq, opts ...client.CallOption) (*UploadResp, error) {
	req := c.c.NewRequest(c.name, "User.UploadAvatar", in)
	out := new(UploadResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) AuthUpdate(ctx context.Context, in *AuthReq, opts ...client.CallOption) (*AuthResp, error) {
	req := c.c.NewRequest(c.name, "User.AuthUpdate", in)
	out := new(AuthResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	MicroGetUser(context.Context, *Request, *Response) error
	UpdateUserName(context.Context, *UpdateReq, *UpdateResp) error
	UploadAvatar(context.Context, *UploadReq, *UploadResp) error
	AuthUpdate(context.Context, *AuthReq, *AuthResp) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		MicroGetUser(ctx context.Context, in *Request, out *Response) error
		UpdateUserName(ctx context.Context, in *UpdateReq, out *UpdateResp) error
		UploadAvatar(ctx context.Context, in *UploadReq, out *UploadResp) error
		AuthUpdate(ctx context.Context, in *AuthReq, out *AuthResp) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) MicroGetUser(ctx context.Context, in *Request, out *Response) error {
	return h.UserHandler.MicroGetUser(ctx, in, out)
}

func (h *userHandler) UpdateUserName(ctx context.Context, in *UpdateReq, out *UpdateResp) error {
	return h.UserHandler.UpdateUserName(ctx, in, out)
}

func (h *userHandler) UploadAvatar(ctx context.Context, in *UploadReq, out *UploadResp) error {
	return h.UserHandler.UploadAvatar(ctx, in, out)
}

func (h *userHandler) AuthUpdate(ctx context.Context, in *AuthReq, out *AuthResp) error {
	return h.UserHandler.AuthUpdate(ctx, in, out)
}
