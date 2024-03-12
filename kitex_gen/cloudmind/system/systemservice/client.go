// Code generated by Kitex v0.9.0. DO NOT EDIT.

package systemservice

import (
	"context"
	system "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/system"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetSliders(ctx context.Context, Req *system.GetSlidersReq, callOptions ...callopt.Option) (r *system.GetSlidersResp, err error)
	CreateSlider(ctx context.Context, Req *system.CreateSliderReq, callOptions ...callopt.Option) (r *system.CreateSliderResp, err error)
	UpdateSlider(ctx context.Context, Req *system.UpdateSliderReq, callOptions ...callopt.Option) (r *system.UpdateSliderResp, err error)
	DeleteSlider(ctx context.Context, Req *system.DeleteSliderReq, callOptions ...callopt.Option) (r *system.DeleteSliderResp, err error)
	GetNotifications(ctx context.Context, Req *system.GetNotificationsReq, callOptions ...callopt.Option) (r *system.GetNotificationsResp, err error)
	GetNotificationCount(ctx context.Context, Req *system.GetNotificationCountReq, callOptions ...callopt.Option) (r *system.GetNotificationCountResp, err error)
	CreateNotifications(ctx context.Context, Req *system.CreateNotificationsReq, callOptions ...callopt.Option) (r *system.CreateNotificationsResp, err error)
	ReadNotifications(ctx context.Context, Req *system.ReadNotificationsReq, callOptions ...callopt.Option) (r *system.ReadNotificationsResp, err error)
	InsertNotificationCount(ctx context.Context, Req *system.InsertNotificationCountReq, callOptions ...callopt.Option) (r *system.InsertNotificationCountResp, err error)
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

func (p *kSystemServiceClient) GetSliders(ctx context.Context, Req *system.GetSlidersReq, callOptions ...callopt.Option) (r *system.GetSlidersResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetSliders(ctx, Req)
}

func (p *kSystemServiceClient) CreateSlider(ctx context.Context, Req *system.CreateSliderReq, callOptions ...callopt.Option) (r *system.CreateSliderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateSlider(ctx, Req)
}

func (p *kSystemServiceClient) UpdateSlider(ctx context.Context, Req *system.UpdateSliderReq, callOptions ...callopt.Option) (r *system.UpdateSliderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateSlider(ctx, Req)
}

func (p *kSystemServiceClient) DeleteSlider(ctx context.Context, Req *system.DeleteSliderReq, callOptions ...callopt.Option) (r *system.DeleteSliderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteSlider(ctx, Req)
}

func (p *kSystemServiceClient) GetNotifications(ctx context.Context, Req *system.GetNotificationsReq, callOptions ...callopt.Option) (r *system.GetNotificationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetNotifications(ctx, Req)
}

func (p *kSystemServiceClient) GetNotificationCount(ctx context.Context, Req *system.GetNotificationCountReq, callOptions ...callopt.Option) (r *system.GetNotificationCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetNotificationCount(ctx, Req)
}

func (p *kSystemServiceClient) CreateNotifications(ctx context.Context, Req *system.CreateNotificationsReq, callOptions ...callopt.Option) (r *system.CreateNotificationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateNotifications(ctx, Req)
}

func (p *kSystemServiceClient) ReadNotifications(ctx context.Context, Req *system.ReadNotificationsReq, callOptions ...callopt.Option) (r *system.ReadNotificationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReadNotifications(ctx, Req)
}

func (p *kSystemServiceClient) InsertNotificationCount(ctx context.Context, Req *system.InsertNotificationCountReq, callOptions ...callopt.Option) (r *system.InsertNotificationCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InsertNotificationCount(ctx, Req)
}
