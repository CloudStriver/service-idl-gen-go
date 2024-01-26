// Code generated by Kitex v0.8.0. DO NOT EDIT.

package content

import (
	"context"
	core_api "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/core_api"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UpdateUser(ctx context.Context, Req *core_api.UpdateUserReq, callOptions ...callopt.Option) (r *core_api.UpdateUserResp, err error)
	SearchUser(ctx context.Context, Req *core_api.SearchUserReq, callOptions ...callopt.Option) (r *core_api.SearchUserResp, err error)
	GetUser(ctx context.Context, Req *core_api.GetUserReq, callOptions ...callopt.Option) (r *core_api.GetUserResp, err error)
	GetUserDetail(ctx context.Context, Req *core_api.GetUserDetailReq, callOptions ...callopt.Option) (r *core_api.GetUserDetailResp, err error)
	GetFile(ctx context.Context, Req *core_api.GetFileReq, callOptions ...callopt.Option) (r *core_api.GetFileResp, err error)
	GetFileList(ctx context.Context, Req *core_api.GetFileListReq, callOptions ...callopt.Option) (r *core_api.GetFileListResp, err error)
	GetFileBySharingCode(ctx context.Context, Req *core_api.GetFileBySharingCodeReq, callOptions ...callopt.Option) (r *core_api.GetFileBySharingCodeResp, err error)
	CreateFile(ctx context.Context, Req *core_api.CreateFileReq, callOptions ...callopt.Option) (r *core_api.CreateFileResp, err error)
	UpdateFile(ctx context.Context, Req *core_api.UpdateFileReq, callOptions ...callopt.Option) (r *core_api.UpdateFileResp, err error)
	MoveFile(ctx context.Context, Req *core_api.MoveFileReq, callOptions ...callopt.Option) (r *core_api.MoveFileResp, err error)
	SaveFileToPrivateSpace(ctx context.Context, Req *core_api.SaveFileToPrivateSpaceReq, callOptions ...callopt.Option) (r *core_api.SaveFileToPrivateSpaceResp, err error)
	AddFileToPublicSpace(ctx context.Context, Req *core_api.AddFileToPublicSpaceReq, callOptions ...callopt.Option) (r *core_api.AddFileToPublicSpaceResp, err error)
	DeleteFile(ctx context.Context, Req *core_api.DeleteFileReq, callOptions ...callopt.Option) (r *core_api.DeleteFileResp, err error)
	RecoverRecycleBinFile(ctx context.Context, Req *core_api.RecoverRecycleBinFileReq, callOptions ...callopt.Option) (r *core_api.RecoverRecycleBinFileResp, err error)
	CreateLabel(ctx context.Context, Req *core_api.CreateLabelReq, callOptions ...callopt.Option) (r *core_api.CreateLabelResp, err error)
	UpdateLabel(ctx context.Context, Req *core_api.UpdateLabelReq, callOptions ...callopt.Option) (r *core_api.UpdateLabelResp, err error)
	GetLabel(ctx context.Context, Req *core_api.GetLabelReq, callOptions ...callopt.Option) (r *core_api.GetLabelResp, err error)
	DeleteLabel(ctx context.Context, Req *core_api.DeleteLabelReq, callOptions ...callopt.Option) (r *core_api.DeleteLabelResp, err error)
	CreateShareCode(ctx context.Context, Req *core_api.CreateShareCodeReq, callOptions ...callopt.Option) (r *core_api.CreateShareCodeResp, err error)
	GetShareList(ctx context.Context, Req *core_api.GetShareListReq, callOptions ...callopt.Option) (r *core_api.GetShareListResp, err error)
	DeleteShareCode(ctx context.Context, Req *core_api.DeleteShareCodeReq, callOptions ...callopt.Option) (r *core_api.DeleteShareCodeResp, err error)
	ParsingShareCode(ctx context.Context, Req *core_api.ParsingShareCodeReq, callOptions ...callopt.Option) (r *core_api.ParsingShareCodeResp, err error)
	CreatePost(ctx context.Context, Req *core_api.CreatePostReq, callOptions ...callopt.Option) (r *core_api.CreatePostResp, err error)
	DeletePost(ctx context.Context, Req *core_api.DeletePostReq, callOptions ...callopt.Option) (r *core_api.DeletePostResp, err error)
	UpdatePost(ctx context.Context, Req *core_api.UpdatePostReq, callOptions ...callopt.Option) (r *core_api.UpdatePostResp, err error)
	GetPosts(ctx context.Context, Req *core_api.GetPostsReq, callOptions ...callopt.Option) (r *core_api.GetPostsResp, err error)
	GetPost(ctx context.Context, Req *core_api.GetPostReq, callOptions ...callopt.Option) (r *core_api.GetPostResp, err error)
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
	return &kContentClient{
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

type kContentClient struct {
	*kClient
}

func (p *kContentClient) UpdateUser(ctx context.Context, Req *core_api.UpdateUserReq, callOptions ...callopt.Option) (r *core_api.UpdateUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUser(ctx, Req)
}

func (p *kContentClient) SearchUser(ctx context.Context, Req *core_api.SearchUserReq, callOptions ...callopt.Option) (r *core_api.SearchUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchUser(ctx, Req)
}

func (p *kContentClient) GetUser(ctx context.Context, Req *core_api.GetUserReq, callOptions ...callopt.Option) (r *core_api.GetUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUser(ctx, Req)
}

func (p *kContentClient) GetUserDetail(ctx context.Context, Req *core_api.GetUserDetailReq, callOptions ...callopt.Option) (r *core_api.GetUserDetailResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserDetail(ctx, Req)
}

func (p *kContentClient) GetFile(ctx context.Context, Req *core_api.GetFileReq, callOptions ...callopt.Option) (r *core_api.GetFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFile(ctx, Req)
}

func (p *kContentClient) GetFileList(ctx context.Context, Req *core_api.GetFileListReq, callOptions ...callopt.Option) (r *core_api.GetFileListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileList(ctx, Req)
}

func (p *kContentClient) GetFileBySharingCode(ctx context.Context, Req *core_api.GetFileBySharingCodeReq, callOptions ...callopt.Option) (r *core_api.GetFileBySharingCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileBySharingCode(ctx, Req)
}

func (p *kContentClient) CreateFile(ctx context.Context, Req *core_api.CreateFileReq, callOptions ...callopt.Option) (r *core_api.CreateFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateFile(ctx, Req)
}

func (p *kContentClient) UpdateFile(ctx context.Context, Req *core_api.UpdateFileReq, callOptions ...callopt.Option) (r *core_api.UpdateFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateFile(ctx, Req)
}

func (p *kContentClient) MoveFile(ctx context.Context, Req *core_api.MoveFileReq, callOptions ...callopt.Option) (r *core_api.MoveFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MoveFile(ctx, Req)
}

func (p *kContentClient) SaveFileToPrivateSpace(ctx context.Context, Req *core_api.SaveFileToPrivateSpaceReq, callOptions ...callopt.Option) (r *core_api.SaveFileToPrivateSpaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SaveFileToPrivateSpace(ctx, Req)
}

func (p *kContentClient) AddFileToPublicSpace(ctx context.Context, Req *core_api.AddFileToPublicSpaceReq, callOptions ...callopt.Option) (r *core_api.AddFileToPublicSpaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddFileToPublicSpace(ctx, Req)
}

func (p *kContentClient) DeleteFile(ctx context.Context, Req *core_api.DeleteFileReq, callOptions ...callopt.Option) (r *core_api.DeleteFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteFile(ctx, Req)
}

func (p *kContentClient) RecoverRecycleBinFile(ctx context.Context, Req *core_api.RecoverRecycleBinFileReq, callOptions ...callopt.Option) (r *core_api.RecoverRecycleBinFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RecoverRecycleBinFile(ctx, Req)
}

func (p *kContentClient) CreateLabel(ctx context.Context, Req *core_api.CreateLabelReq, callOptions ...callopt.Option) (r *core_api.CreateLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateLabel(ctx, Req)
}

func (p *kContentClient) UpdateLabel(ctx context.Context, Req *core_api.UpdateLabelReq, callOptions ...callopt.Option) (r *core_api.UpdateLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateLabel(ctx, Req)
}

func (p *kContentClient) GetLabel(ctx context.Context, Req *core_api.GetLabelReq, callOptions ...callopt.Option) (r *core_api.GetLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLabel(ctx, Req)
}

func (p *kContentClient) DeleteLabel(ctx context.Context, Req *core_api.DeleteLabelReq, callOptions ...callopt.Option) (r *core_api.DeleteLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteLabel(ctx, Req)
}

func (p *kContentClient) CreateShareCode(ctx context.Context, Req *core_api.CreateShareCodeReq, callOptions ...callopt.Option) (r *core_api.CreateShareCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateShareCode(ctx, Req)
}

func (p *kContentClient) GetShareList(ctx context.Context, Req *core_api.GetShareListReq, callOptions ...callopt.Option) (r *core_api.GetShareListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetShareList(ctx, Req)
}

func (p *kContentClient) DeleteShareCode(ctx context.Context, Req *core_api.DeleteShareCodeReq, callOptions ...callopt.Option) (r *core_api.DeleteShareCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteShareCode(ctx, Req)
}

func (p *kContentClient) ParsingShareCode(ctx context.Context, Req *core_api.ParsingShareCodeReq, callOptions ...callopt.Option) (r *core_api.ParsingShareCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ParsingShareCode(ctx, Req)
}

func (p *kContentClient) CreatePost(ctx context.Context, Req *core_api.CreatePostReq, callOptions ...callopt.Option) (r *core_api.CreatePostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreatePost(ctx, Req)
}

func (p *kContentClient) DeletePost(ctx context.Context, Req *core_api.DeletePostReq, callOptions ...callopt.Option) (r *core_api.DeletePostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeletePost(ctx, Req)
}

func (p *kContentClient) UpdatePost(ctx context.Context, Req *core_api.UpdatePostReq, callOptions ...callopt.Option) (r *core_api.UpdatePostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePost(ctx, Req)
}

func (p *kContentClient) GetPosts(ctx context.Context, Req *core_api.GetPostsReq, callOptions ...callopt.Option) (r *core_api.GetPostsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPosts(ctx, Req)
}

func (p *kContentClient) GetPost(ctx context.Context, Req *core_api.GetPostReq, callOptions ...callopt.Option) (r *core_api.GetPostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPost(ctx, Req)
}
