// Code generated by Kitex v0.8.0. DO NOT EDIT.

package auth

import (
	"context"
	core_api "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/core_api"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AskUploadAvatar(ctx context.Context, Req *core_api.AskUploadAvatarReq, callOptions ...callopt.Option) (r *core_api.AskUploadAvatarResp, err error)
	Register(ctx context.Context, Req *core_api.RegisterReq, callOptions ...callopt.Option) (r *core_api.RegisterResp, err error)
	CheckEmail(ctx context.Context, Req *core_api.CheckEmailReq, callOptions ...callopt.Option) (r *core_api.CheckEmailResp, err error)
	EmailLogin(ctx context.Context, Req *core_api.EmailLoginReq, callOptions ...callopt.Option) (r *core_api.EmailLoginResp, err error)
	GithubLogin(ctx context.Context, Req *core_api.GithubLoginReq, callOptions ...callopt.Option) (r *core_api.GithubLoginResp, err error)
	GiteeLogin(ctx context.Context, Req *core_api.GiteeLoginReq, callOptions ...callopt.Option) (r *core_api.GiteeLoginResp, err error)
	RefreshToken(ctx context.Context, Req *core_api.RefreshTokenReq, callOptions ...callopt.Option) (r *core_api.RefreshTokenResp, err error)
	SendEmail(ctx context.Context, Req *core_api.SendEmailReq, callOptions ...callopt.Option) (r *core_api.SendEmailResp, err error)
	SetPasswordByEmail(ctx context.Context, Req *core_api.SetPasswordByEmailReq, callOptions ...callopt.Option) (r *core_api.SetPasswordByEmailResp, err error)
	SetPasswordByPassword(ctx context.Context, Req *core_api.SetPasswordByPasswordReq, callOptions ...callopt.Option) (r *core_api.SetPasswordByPasswordReq, err error)
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
	return &kAuthClient{
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

type kAuthClient struct {
	*kClient
}

func (p *kAuthClient) AskUploadAvatar(ctx context.Context, Req *core_api.AskUploadAvatarReq, callOptions ...callopt.Option) (r *core_api.AskUploadAvatarResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AskUploadAvatar(ctx, Req)
}

func (p *kAuthClient) Register(ctx context.Context, Req *core_api.RegisterReq, callOptions ...callopt.Option) (r *core_api.RegisterResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, Req)
}

func (p *kAuthClient) CheckEmail(ctx context.Context, Req *core_api.CheckEmailReq, callOptions ...callopt.Option) (r *core_api.CheckEmailResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckEmail(ctx, Req)
}

func (p *kAuthClient) EmailLogin(ctx context.Context, Req *core_api.EmailLoginReq, callOptions ...callopt.Option) (r *core_api.EmailLoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EmailLogin(ctx, Req)
}

func (p *kAuthClient) GithubLogin(ctx context.Context, Req *core_api.GithubLoginReq, callOptions ...callopt.Option) (r *core_api.GithubLoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GithubLogin(ctx, Req)
}

func (p *kAuthClient) GiteeLogin(ctx context.Context, Req *core_api.GiteeLoginReq, callOptions ...callopt.Option) (r *core_api.GiteeLoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GiteeLogin(ctx, Req)
}

func (p *kAuthClient) RefreshToken(ctx context.Context, Req *core_api.RefreshTokenReq, callOptions ...callopt.Option) (r *core_api.RefreshTokenResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RefreshToken(ctx, Req)
}

func (p *kAuthClient) SendEmail(ctx context.Context, Req *core_api.SendEmailReq, callOptions ...callopt.Option) (r *core_api.SendEmailResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendEmail(ctx, Req)
}

func (p *kAuthClient) SetPasswordByEmail(ctx context.Context, Req *core_api.SetPasswordByEmailReq, callOptions ...callopt.Option) (r *core_api.SetPasswordByEmailResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetPasswordByEmail(ctx, Req)
}

func (p *kAuthClient) SetPasswordByPassword(ctx context.Context, Req *core_api.SetPasswordByPasswordReq, callOptions ...callopt.Option) (r *core_api.SetPasswordByPasswordReq, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetPasswordByPassword(ctx, Req)
}
