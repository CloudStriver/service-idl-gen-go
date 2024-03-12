// Code generated by Kitex v0.8.0. DO NOT EDIT.

package system

import (
	"context"
	core_api "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/core_api"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetNotifications(ctx context.Context, Req *core_api.GetNotificationsReq, callOptions ...callopt.Option) (r *core_api.GetNotificationsResp, err error)
	GetNotificationCount(ctx context.Context, Req *core_api.GetNotificationCountReq, callOptions ...callopt.Option) (r *core_api.GetNotificationCountResp, err error)
	ReadNotifications(ctx context.Context, Req *core_api.ReadNotificationsReq, callOptions ...callopt.Option) (r *core_api.ReadNotificationsResp, err error)
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
	return &kSystemClient{
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

type kSystemClient struct {
	*kClient
}

func (p *kSystemClient) GetNotifications(ctx context.Context, Req *core_api.GetNotificationsReq, callOptions ...callopt.Option) (r *core_api.GetNotificationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetNotifications(ctx, Req)
}

func (p *kSystemClient) GetNotificationCount(ctx context.Context, Req *core_api.GetNotificationCountReq, callOptions ...callopt.Option) (r *core_api.GetNotificationCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetNotificationCount(ctx, Req)
}

func (p *kSystemClient) ReadNotifications(ctx context.Context, Req *core_api.ReadNotificationsReq, callOptions ...callopt.Option) (r *core_api.ReadNotificationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReadNotifications(ctx, Req)
}
