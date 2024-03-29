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
	SetCommentAttrs(ctx context.Context, Req *comment.SetCommentAttrsReq, callOptions ...callopt.Option) (r *comment.SetCommentAttrsResp, err error)
	GetCommentSubject(ctx context.Context, Req *comment.GetCommentSubjectReq, callOptions ...callopt.Option) (r *comment.GetCommentSubjectResp, err error)
	CreateCommentSubject(ctx context.Context, Req *comment.CreateCommentSubjectReq, callOptions ...callopt.Option) (r *comment.CreateCommentSubjectResp, err error)
	UpdateCommentSubject(ctx context.Context, Req *comment.UpdateCommentSubjectReq, callOptions ...callopt.Option) (r *comment.UpdateCommentSubjectResp, err error)
	DeleteCommentSubject(ctx context.Context, Req *comment.DeleteCommentSubjectReq, callOptions ...callopt.Option) (r *comment.DeleteCommentSubjectResp, err error)
	CreateLabel(ctx context.Context, Req *comment.CreateLabelReq, callOptions ...callopt.Option) (r *comment.CreateLabelResp, err error)
	DeleteLabel(ctx context.Context, Req *comment.DeleteLabelReq, callOptions ...callopt.Option) (r *comment.DeleteLabelResp, err error)
	GetLabel(ctx context.Context, Req *comment.GetLabelReq, callOptions ...callopt.Option) (r *comment.GetLabelResp, err error)
	GetLabelsInBatch(ctx context.Context, Req *comment.GetLabelsInBatchReq, callOptions ...callopt.Option) (r *comment.GetLabelsInBatchResp, err error)
	UpdateLabel(ctx context.Context, Req *comment.UpdateLabelReq, callOptions ...callopt.Option) (r *comment.UpdateLabelResp, err error)
	GetLabels(ctx context.Context, Req *comment.GetLabelsReq, callOptions ...callopt.Option) (r *comment.GetLabelsResp, err error)
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

func (p *kCommentServiceClient) CreateLabel(ctx context.Context, Req *comment.CreateLabelReq, callOptions ...callopt.Option) (r *comment.CreateLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateLabel(ctx, Req)
}

func (p *kCommentServiceClient) DeleteLabel(ctx context.Context, Req *comment.DeleteLabelReq, callOptions ...callopt.Option) (r *comment.DeleteLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteLabel(ctx, Req)
}

func (p *kCommentServiceClient) GetLabel(ctx context.Context, Req *comment.GetLabelReq, callOptions ...callopt.Option) (r *comment.GetLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLabel(ctx, Req)
}

func (p *kCommentServiceClient) GetLabelsInBatch(ctx context.Context, Req *comment.GetLabelsInBatchReq, callOptions ...callopt.Option) (r *comment.GetLabelsInBatchResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLabelsInBatch(ctx, Req)
}

func (p *kCommentServiceClient) UpdateLabel(ctx context.Context, Req *comment.UpdateLabelReq, callOptions ...callopt.Option) (r *comment.UpdateLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateLabel(ctx, Req)
}

func (p *kCommentServiceClient) GetLabels(ctx context.Context, Req *comment.GetLabelsReq, callOptions ...callopt.Option) (r *comment.GetLabelsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLabels(ctx, Req)
}
