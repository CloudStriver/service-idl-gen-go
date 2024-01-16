// Code generated by Kitex v0.8.0. DO NOT EDIT.

package relationservice

import (
	"context"
	relation "github.com/CloudStriver/service-idl-gen-go/kitex_gen/platform/relation"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateRelation(ctx context.Context, Req *relation.CreateRelationReq, callOptions ...callopt.Option) (r *relation.CreateRelationResp, err error)
	GetRelations(ctx context.Context, Req *relation.GetRelationsReq, callOptions ...callopt.Option) (r *relation.GetRelationsResp, err error)
	GetRelationCount(ctx context.Context, Req *relation.GetRelationCountReq, callOptions ...callopt.Option) (r *relation.GetRelationCountResp, err error)
	DeleteRelation(ctx context.Context, Req *relation.DeleteRelationReq, callOptions ...callopt.Option) (r *relation.DeleteRelationResp, err error)
	GetRelation(ctx context.Context, Req *relation.GetRelationReq, callOptions ...callopt.Option) (r *relation.GetRelationResp, err error)
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
	return &kRelationServiceClient{
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

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) CreateRelation(ctx context.Context, Req *relation.CreateRelationReq, callOptions ...callopt.Option) (r *relation.CreateRelationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateRelation(ctx, Req)
}

func (p *kRelationServiceClient) GetRelations(ctx context.Context, Req *relation.GetRelationsReq, callOptions ...callopt.Option) (r *relation.GetRelationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRelations(ctx, Req)
}

func (p *kRelationServiceClient) GetRelationCount(ctx context.Context, Req *relation.GetRelationCountReq, callOptions ...callopt.Option) (r *relation.GetRelationCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRelationCount(ctx, Req)
}

func (p *kRelationServiceClient) DeleteRelation(ctx context.Context, Req *relation.DeleteRelationReq, callOptions ...callopt.Option) (r *relation.DeleteRelationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteRelation(ctx, Req)
}

func (p *kRelationServiceClient) GetRelation(ctx context.Context, Req *relation.GetRelationReq, callOptions ...callopt.Option) (r *relation.GetRelationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRelation(ctx, Req)
}
