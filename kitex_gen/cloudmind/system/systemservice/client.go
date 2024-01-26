// Code generated by Kitex v0.8.0. DO NOT EDIT.

package systemservice

import (
	"context"
	system "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/system"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetNotifications(ctx context.Context, Req *system.GetNotificationsReq, callOptions ...callopt.Option) (r *system.GetNotificationsResp, err error)
	CleanNotification(ctx context.Context, Req *system.CleanNotificationReq, callOptions ...callopt.Option) (r *system.CleanNotificationResp, err error)
	GetNotificationCount(ctx context.Context, Req *system.GetNotificationCountReq, callOptions ...callopt.Option) (r *system.GetNotificationCountResp, err error)
	ReadNotification(ctx context.Context, Req *system.ReadNotificationReq, callOptions ...callopt.Option) (r *system.ReadNotificationResp, err error)
	CreateNotification(ctx context.Context, Req *system.CreateNotificationReq, callOptions ...callopt.Option) (r *system.CreateNotificationResp, err error)
	ReadNotifications(ctx context.Context, Req *system.ReadNotificationsReq, callOptions ...callopt.Option) (r *system.ReadNotificationsResp, err error)
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
	return &kSystemServiceClient{
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

type kSystemServiceClient struct {
	*kClient
}

func (p *kSystemServiceClient) GetNotifications(ctx context.Context, Req *system.GetNotificationsReq, callOptions ...callopt.Option) (r *system.GetNotificationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetNotifications(ctx, Req)
}

func (p *kSystemServiceClient) CleanNotification(ctx context.Context, Req *system.CleanNotificationReq, callOptions ...callopt.Option) (r *system.CleanNotificationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CleanNotification(ctx, Req)
}

func (p *kSystemServiceClient) GetNotificationCount(ctx context.Context, Req *system.GetNotificationCountReq, callOptions ...callopt.Option) (r *system.GetNotificationCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetNotificationCount(ctx, Req)
}

func (p *kSystemServiceClient) ReadNotification(ctx context.Context, Req *system.ReadNotificationReq, callOptions ...callopt.Option) (r *system.ReadNotificationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReadNotification(ctx, Req)
}

func (p *kSystemServiceClient) CreateNotification(ctx context.Context, Req *system.CreateNotificationReq, callOptions ...callopt.Option) (r *system.CreateNotificationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateNotification(ctx, Req)
}

func (p *kSystemServiceClient) ReadNotifications(ctx context.Context, Req *system.ReadNotificationsReq, callOptions ...callopt.Option) (r *system.ReadNotificationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReadNotifications(ctx, Req)
}
