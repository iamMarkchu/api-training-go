// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package action

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type ActionHTTPServer interface {
	CreateAction(context.Context, *CreateActionRequest) (*CreateActionReply, error)
	DeleteAction(context.Context, *DeleteActionRequest) (*DeleteActionReply, error)
	GetAction(context.Context, *GetActionRequest) (*GetActionReply, error)
	ListAction(context.Context, *ListActionRequest) (*ListActionReply, error)
	MGetAction(context.Context, *MGetActionRequest) (*MGetActionReply, error)
	UpdateAction(context.Context, *UpdateActionRequest) (*UpdateActionReply, error)
}

func RegisterActionHTTPServer(s *http.Server, srv ActionHTTPServer) {
	r := s.Route("/")
	r.POST("/action/create", _Action_CreateAction0_HTTP_Handler(srv))
	r.POST("/action/update", _Action_UpdateAction0_HTTP_Handler(srv))
	r.POST("/action/delete", _Action_DeleteAction0_HTTP_Handler(srv))
	r.GET("/action/get/{id}", _Action_GetAction0_HTTP_Handler(srv))
	r.GET("/action/list", _Action_ListAction0_HTTP_Handler(srv))
	r.POST("/action/mget", _Action_MGetAction0_HTTP_Handler(srv))
}

func _Action_CreateAction0_HTTP_Handler(srv ActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.action.Action/CreateAction")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAction(ctx, req.(*CreateActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateActionReply)
		return ctx.Result(200, reply)
	}
}

func _Action_UpdateAction0_HTTP_Handler(srv ActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.action.Action/UpdateAction")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAction(ctx, req.(*UpdateActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateActionReply)
		return ctx.Result(200, reply)
	}
}

func _Action_DeleteAction0_HTTP_Handler(srv ActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.action.Action/DeleteAction")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAction(ctx, req.(*DeleteActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteActionReply)
		return ctx.Result(200, reply)
	}
}

func _Action_GetAction0_HTTP_Handler(srv ActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetActionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.action.Action/GetAction")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAction(ctx, req.(*GetActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetActionReply)
		return ctx.Result(200, reply)
	}
}

func _Action_ListAction0_HTTP_Handler(srv ActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListActionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.action.Action/ListAction")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListAction(ctx, req.(*ListActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListActionReply)
		return ctx.Result(200, reply)
	}
}

func _Action_MGetAction0_HTTP_Handler(srv ActionHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MGetActionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.action.Action/MGetAction")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MGetAction(ctx, req.(*MGetActionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MGetActionReply)
		return ctx.Result(200, reply)
	}
}

type ActionHTTPClient interface {
	CreateAction(ctx context.Context, req *CreateActionRequest, opts ...http.CallOption) (rsp *CreateActionReply, err error)
	DeleteAction(ctx context.Context, req *DeleteActionRequest, opts ...http.CallOption) (rsp *DeleteActionReply, err error)
	GetAction(ctx context.Context, req *GetActionRequest, opts ...http.CallOption) (rsp *GetActionReply, err error)
	ListAction(ctx context.Context, req *ListActionRequest, opts ...http.CallOption) (rsp *ListActionReply, err error)
	MGetAction(ctx context.Context, req *MGetActionRequest, opts ...http.CallOption) (rsp *MGetActionReply, err error)
	UpdateAction(ctx context.Context, req *UpdateActionRequest, opts ...http.CallOption) (rsp *UpdateActionReply, err error)
}

type ActionHTTPClientImpl struct {
	cc *http.Client
}

func NewActionHTTPClient(client *http.Client) ActionHTTPClient {
	return &ActionHTTPClientImpl{client}
}

func (c *ActionHTTPClientImpl) CreateAction(ctx context.Context, in *CreateActionRequest, opts ...http.CallOption) (*CreateActionReply, error) {
	var out CreateActionReply
	pattern := "/action/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.action.Action/CreateAction"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ActionHTTPClientImpl) DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...http.CallOption) (*DeleteActionReply, error) {
	var out DeleteActionReply
	pattern := "/action/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.action.Action/DeleteAction"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ActionHTTPClientImpl) GetAction(ctx context.Context, in *GetActionRequest, opts ...http.CallOption) (*GetActionReply, error) {
	var out GetActionReply
	pattern := "/action/get/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.action.Action/GetAction"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ActionHTTPClientImpl) ListAction(ctx context.Context, in *ListActionRequest, opts ...http.CallOption) (*ListActionReply, error) {
	var out ListActionReply
	pattern := "/action/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.action.Action/ListAction"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ActionHTTPClientImpl) MGetAction(ctx context.Context, in *MGetActionRequest, opts ...http.CallOption) (*MGetActionReply, error) {
	var out MGetActionReply
	pattern := "/action/mget"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.action.Action/MGetAction"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ActionHTTPClientImpl) UpdateAction(ctx context.Context, in *UpdateActionRequest, opts ...http.CallOption) (*UpdateActionReply, error) {
	var out UpdateActionReply
	pattern := "/action/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.action.Action/UpdateAction"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
