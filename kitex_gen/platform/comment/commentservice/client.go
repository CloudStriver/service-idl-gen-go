// Code generated by Kitex v0.8.0. DO NOT EDIT.

package commentservice

import (
	"context"
	comment "github.com/CloudStriver/service-idl-gen-go/kitex_gen/platform/comment"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetComment(ctx context.Context, Req *comment.GetCommentReq, callOptions ...callopt.Option) (r *comment.GetCommentResp, err error)
	GetCommentList(ctx context.Context, Req *comment.GetCommentListReq, callOptions ...callopt.Option) (r *comment.GetCommentListResp, err error)
	CreateComment(ctx context.Context, Req *comment.CreateCommentReq, callOptions ...callopt.Option) (r *comment.CreateCommentResp, err error)
	UpdateComment(ctx context.Context, Req *comment.UpdateCommentReq, callOptions ...callopt.Option) (r *comment.UpdateCommentResp, err error)
	DeleteComment(ctx context.Context, Req *comment.DeleteCommentReq, callOptions ...callopt.Option) (r *comment.DeleteCommentResp, err error)
	DeleteCommentWithUserId(ctx context.Context, Req *comment.DeleteCommentWithUserIdReq, callOptions ...callopt.Option) (r *comment.DeleteCommentWithUserIdResp, err error)
	SetCommentState(ctx context.Context, Req *comment.SetCommentStateReq, callOptions ...callopt.Option) (r *comment.SetCommentStateResp, err error)
	SetCommentAttrs(ctx context.Context, Req *comment.SetCommentAttrsReq, callOptions ...callopt.Option) (r *comment.SetCommentAttrsResp, err error)
	GetCommentSubject(ctx context.Context, Req *comment.GetCommentSubjectReq, callOptions ...callopt.Option) (r *comment.GetCommentSubjectResp, err error)
	CreateCommentSubject(ctx context.Context, Req *comment.CreateCommentSubjectReq, callOptions ...callopt.Option) (r *comment.CreateCommentSubjectResp, err error)
	UpdateCommentSubject(ctx context.Context, Req *comment.UpdateCommentSubjectReq, callOptions ...callopt.Option) (r *comment.UpdateCommentSubjectResp, err error)
	DeleteCommentSubject(ctx context.Context, Req *comment.DeleteCommentSubjectReq, callOptions ...callopt.Option) (r *comment.DeleteCommentSubjectResp, err error)
	SetCommentSubjectState(ctx context.Context, Req *comment.SetCommentSubjectStateReq, callOptions ...callopt.Option) (r *comment.SetCommentSubjectStateResp, err error)
	SetCommentSubjectAttrs(ctx context.Context, Req *comment.SetCommentSubjectAttrsReq, callOptions ...callopt.Option) (r *comment.SetCommentSubjectAttrsResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCommentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCommentServiceClient struct {
	*kClient
}

func (p *kCommentServiceClient) GetComment(ctx context.Context, Req *comment.GetCommentReq, callOptions ...callopt.Option) (r *comment.GetCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetComment(ctx, Req)
}

func (p *kCommentServiceClient) GetCommentList(ctx context.Context, Req *comment.GetCommentListReq, callOptions ...callopt.Option) (r *comment.GetCommentListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCommentList(ctx, Req)
}

func (p *kCommentServiceClient) CreateComment(ctx context.Context, Req *comment.CreateCommentReq, callOptions ...callopt.Option) (r *comment.CreateCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateComment(ctx, Req)
}

func (p *kCommentServiceClient) UpdateComment(ctx context.Context, Req *comment.UpdateCommentReq, callOptions ...callopt.Option) (r *comment.UpdateCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateComment(ctx, Req)
}

func (p *kCommentServiceClient) DeleteComment(ctx context.Context, Req *comment.DeleteCommentReq, callOptions ...callopt.Option) (r *comment.DeleteCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, Req)
}

func (p *kCommentServiceClient) DeleteCommentWithUserId(ctx context.Context, Req *comment.DeleteCommentWithUserIdReq, callOptions ...callopt.Option) (r *comment.DeleteCommentWithUserIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCommentWithUserId(ctx, Req)
}

func (p *kCommentServiceClient) SetCommentState(ctx context.Context, Req *comment.SetCommentStateReq, callOptions ...callopt.Option) (r *comment.SetCommentStateResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetCommentState(ctx, Req)
}

func (p *kCommentServiceClient) SetCommentAttrs(ctx context.Context, Req *comment.SetCommentAttrsReq, callOptions ...callopt.Option) (r *comment.SetCommentAttrsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetCommentAttrs(ctx, Req)
}

func (p *kCommentServiceClient) GetCommentSubject(ctx context.Context, Req *comment.GetCommentSubjectReq, callOptions ...callopt.Option) (r *comment.GetCommentSubjectResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCommentSubject(ctx, Req)
}

func (p *kCommentServiceClient) CreateCommentSubject(ctx context.Context, Req *comment.CreateCommentSubjectReq, callOptions ...callopt.Option) (r *comment.CreateCommentSubjectResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateCommentSubject(ctx, Req)
}

func (p *kCommentServiceClient) UpdateCommentSubject(ctx context.Context, Req *comment.UpdateCommentSubjectReq, callOptions ...callopt.Option) (r *comment.UpdateCommentSubjectResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateCommentSubject(ctx, Req)
}

func (p *kCommentServiceClient) DeleteCommentSubject(ctx context.Context, Req *comment.DeleteCommentSubjectReq, callOptions ...callopt.Option) (r *comment.DeleteCommentSubjectResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCommentSubject(ctx, Req)
}

func (p *kCommentServiceClient) SetCommentSubjectState(ctx context.Context, Req *comment.SetCommentSubjectStateReq, callOptions ...callopt.Option) (r *comment.SetCommentSubjectStateResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetCommentSubjectState(ctx, Req)
}

func (p *kCommentServiceClient) SetCommentSubjectAttrs(ctx context.Context, Req *comment.SetCommentSubjectAttrsReq, callOptions ...callopt.Option) (r *comment.SetCommentSubjectAttrsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetCommentSubjectAttrs(ctx, Req)
}