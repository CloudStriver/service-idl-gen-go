// Code generated by Kitex v0.9.0. DO NOT EDIT.

package relation

import (
	"context"
	core_api "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/core_api"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateRelation(ctx context.Context, Req *core_api.CreateRelationReq, callOptions ...callopt.Option) (r *core_api.CreateRelationResp, err error)
	GetFromRelations(ctx context.Context, Req *core_api.GetFromRelationsReq, callOptions ...callopt.Option) (r *core_api.GetFromRelationsResp, err error)
	GetToRelations(ctx context.Context, Req *core_api.GetToRelationsReq, callOptions ...callopt.Option) (r *core_api.GetToRelationsResp, err error)
	GetRelation(ctx context.Context, Req *core_api.GetRelationReq, callOptions ...callopt.Option) (r *core_api.GetRelationResp, err error)
	DeleteRelation(ctx context.Context, Req *core_api.DeleteRelationReq, callOptions ...callopt.Option) (r *core_api.DeleteRelationResp, err error)
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
	return &kRelationClient{
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

type kRelationClient struct {
	*kClient
}

func (p *kRelationClient) CreateRelation(ctx context.Context, Req *core_api.CreateRelationReq, callOptions ...callopt.Option) (r *core_api.CreateRelationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateRelation(ctx, Req)
}

func (p *kRelationClient) GetFromRelations(ctx context.Context, Req *core_api.GetFromRelationsReq, callOptions ...callopt.Option) (r *core_api.GetFromRelationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFromRelations(ctx, Req)
}

func (p *kRelationClient) GetToRelations(ctx context.Context, Req *core_api.GetToRelationsReq, callOptions ...callopt.Option) (r *core_api.GetToRelationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetToRelations(ctx, Req)
}

func (p *kRelationClient) GetRelation(ctx context.Context, Req *core_api.GetRelationReq, callOptions ...callopt.Option) (r *core_api.GetRelationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRelation(ctx, Req)
}

func (p *kRelationClient) DeleteRelation(ctx context.Context, Req *core_api.DeleteRelationReq, callOptions ...callopt.Option) (r *core_api.DeleteRelationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteRelation(ctx, Req)
}
