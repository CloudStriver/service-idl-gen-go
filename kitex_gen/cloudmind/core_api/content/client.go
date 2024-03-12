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
	AskUploadAvatar(ctx context.Context, Req *core_api.AskUploadAvatarReq, callOptions ...callopt.Option) (r *core_api.AskUploadAvatarResp, err error)
	GetPublicFile(ctx context.Context, Req *core_api.GetPublicFilesReq, callOptions ...callopt.Option) (r *core_api.GetPublicFilesResp, err error)
	GetPrivateFile(ctx context.Context, Req *core_api.GetPrivateFileReq, callOptions ...callopt.Option) (r *core_api.GetPrivateFileResp, err error)
	GetPrivateFiles(ctx context.Context, Req *core_api.GetPrivateFilesReq, callOptions ...callopt.Option) (r *core_api.GetPrivateFilesResp, err error)
	GetPublicFiles(ctx context.Context, Req *core_api.GetPublicFilesReq, callOptions ...callopt.Option) (r *core_api.GetPublicFilesResp, err error)
	GetRecycleBinFiles(ctx context.Context, Req *core_api.GetRecycleBinFilesReq, callOptions ...callopt.Option) (r *core_api.GetRecycleBinFilesResp, err error)
	GetFileBySharingCode(ctx context.Context, Req *core_api.GetFileBySharingCodeReq, callOptions ...callopt.Option) (r *core_api.GetFileBySharingCodeResp, err error)
	CreateFile(ctx context.Context, Req *core_api.CreateFileReq, callOptions ...callopt.Option) (r *core_api.CreateFileResp, err error)
	UpdateFile(ctx context.Context, Req *core_api.UpdateFileReq, callOptions ...callopt.Option) (r *core_api.UpdateFileResp, err error)
	MoveFile(ctx context.Context, Req *core_api.MoveFileReq, callOptions ...callopt.Option) (r *core_api.MoveFileResp, err error)
	SaveFileToPrivateSpace(ctx context.Context, Req *core_api.SaveFileToPrivateSpaceReq, callOptions ...callopt.Option) (r *core_api.SaveFileToPrivateSpaceResp, err error)
	AddFileToPublicSpace(ctx context.Context, Req *core_api.AddFileToPublicSpaceReq, callOptions ...callopt.Option) (r *core_api.AddFileToPublicSpaceResp, err error)
	CompletelyRemoveFile(ctx context.Context, Req *core_api.CompletelyRemoveFileReq, callOptions ...callopt.Option) (r *core_api.CompletelyRemoveFileReq, err error)
	DeleteFile(ctx context.Context, Req *core_api.DeleteFileReq, callOptions ...callopt.Option) (r *core_api.DeleteFileResp, err error)
	RecoverRecycleBinFile(ctx context.Context, Req *core_api.RecoverRecycleBinFileReq, callOptions ...callopt.Option) (r *core_api.RecoverRecycleBinFileResp, err error)
	CreateZone(ctx context.Context, Req *core_api.CreateZoneReq, callOptions ...callopt.Option) (r *core_api.CreateZoneResp, err error)
	UpdateZone(ctx context.Context, Req *core_api.UpdateZoneReq, callOptions ...callopt.Option) (r *core_api.UpdateZoneResp, err error)
	GetZone(ctx context.Context, Req *core_api.GetZoneReq, callOptions ...callopt.Option) (r *core_api.GetZoneResp, err error)
	DeleteZone(ctx context.Context, Req *core_api.DeleteZoneReq, callOptions ...callopt.Option) (r *core_api.DeleteZoneResp, err error)
	CreateShareCode(ctx context.Context, Req *core_api.CreateShareCodeReq, callOptions ...callopt.Option) (r *core_api.CreateShareCodeResp, err error)
	GetShareList(ctx context.Context, Req *core_api.GetShareListReq, callOptions ...callopt.Option) (r *core_api.GetShareListResp, err error)
	DeleteShareCode(ctx context.Context, Req *core_api.DeleteShareCodeReq, callOptions ...callopt.Option) (r *core_api.DeleteShareCodeResp, err error)
	ParsingShareCode(ctx context.Context, Req *core_api.ParsingShareCodeReq, callOptions ...callopt.Option) (r *core_api.ParsingShareCodeResp, err error)
	AskUploadFile(ctx context.Context, Req *core_api.AskUploadFileReq, callOptions ...callopt.Option) (r *core_api.AskUploadFileResp, err error)
	AskDownloadFile(ctx context.Context, Req *core_api.AskDownloadFileReq, callOptions ...callopt.Option) (r *core_api.AskDownloadFileResp, err error)
	CreatePost(ctx context.Context, Req *core_api.CreatePostReq, callOptions ...callopt.Option) (r *core_api.CreatePostResp, err error)
	DeletePost(ctx context.Context, Req *core_api.DeletePostReq, callOptions ...callopt.Option) (r *core_api.DeletePostResp, err error)
	UpdatePost(ctx context.Context, Req *core_api.UpdatePostReq, callOptions ...callopt.Option) (r *core_api.UpdatePostResp, err error)
	GetPosts(ctx context.Context, Req *core_api.GetPostsReq, callOptions ...callopt.Option) (r *core_api.GetPostsResp, err error)
	GetPost(ctx context.Context, Req *core_api.GetPostReq, callOptions ...callopt.Option) (r *core_api.GetPostResp, err error)
	GetRecommendByUser(ctx context.Context, Req *core_api.GetRecommendByUserReq, callOptions ...callopt.Option) (r *core_api.GetRecommendByUserResp, err error)
	GetRecommendByItem(ctx context.Context, Req *core_api.GetRecommendByItemReq, callOptions ...callopt.Option) (r *core_api.GetRecommendByItemResp, err error)
	CreateFeedBack(ctx context.Context, Req *core_api.CreateFeedBackReq, callOptions ...callopt.Option) (r *core_api.CreateFeedBackResp, err error)
	GetPopularRecommend(ctx context.Context, Req *core_api.GetPopularRecommendReq, callOptions ...callopt.Option) (r *core_api.GetPopularRecommendResp, err error)
	GetLatestRecommend(ctx context.Context, Req *core_api.GetLatestRecommendReq, callOptions ...callopt.Option) (r *core_api.GetLatestRecommendResp, err error)
	CreateProduct(ctx context.Context, Req *core_api.CreateProductReq, callOptions ...callopt.Option) (r *core_api.CreateProductResp, err error)
	GetProduct(ctx context.Context, Req *core_api.GetProductReq, callOptions ...callopt.Option) (r *core_api.GetProductResp, err error)
	GetProducts(ctx context.Context, Req *core_api.GetProductsReq, callOptions ...callopt.Option) (r *core_api.GetProductsResp, err error)
	UpdateProduct(ctx context.Context, Req *core_api.UpdateProductReq, callOptions ...callopt.Option) (r *core_api.UpdateProductResp, err error)
	DeleteProduct(ctx context.Context, Req *core_api.DeleteProductReq, callOptions ...callopt.Option) (r *core_api.DeleteProductResp, err error)
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

func (p *kContentClient) AskUploadAvatar(ctx context.Context, Req *core_api.AskUploadAvatarReq, callOptions ...callopt.Option) (r *core_api.AskUploadAvatarResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AskUploadAvatar(ctx, Req)
}

func (p *kContentClient) GetPublicFile(ctx context.Context, Req *core_api.GetPublicFilesReq, callOptions ...callopt.Option) (r *core_api.GetPublicFilesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublicFile(ctx, Req)
}

func (p *kContentClient) GetPrivateFile(ctx context.Context, Req *core_api.GetPrivateFileReq, callOptions ...callopt.Option) (r *core_api.GetPrivateFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPrivateFile(ctx, Req)
}

func (p *kContentClient) GetPrivateFiles(ctx context.Context, Req *core_api.GetPrivateFilesReq, callOptions ...callopt.Option) (r *core_api.GetPrivateFilesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPrivateFiles(ctx, Req)
}

func (p *kContentClient) GetPublicFiles(ctx context.Context, Req *core_api.GetPublicFilesReq, callOptions ...callopt.Option) (r *core_api.GetPublicFilesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublicFiles(ctx, Req)
}

func (p *kContentClient) GetRecycleBinFiles(ctx context.Context, Req *core_api.GetRecycleBinFilesReq, callOptions ...callopt.Option) (r *core_api.GetRecycleBinFilesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRecycleBinFiles(ctx, Req)
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

func (p *kContentClient) CompletelyRemoveFile(ctx context.Context, Req *core_api.CompletelyRemoveFileReq, callOptions ...callopt.Option) (r *core_api.CompletelyRemoveFileReq, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CompletelyRemoveFile(ctx, Req)
}

func (p *kContentClient) DeleteFile(ctx context.Context, Req *core_api.DeleteFileReq, callOptions ...callopt.Option) (r *core_api.DeleteFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteFile(ctx, Req)
}

func (p *kContentClient) RecoverRecycleBinFile(ctx context.Context, Req *core_api.RecoverRecycleBinFileReq, callOptions ...callopt.Option) (r *core_api.RecoverRecycleBinFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RecoverRecycleBinFile(ctx, Req)
}

func (p *kContentClient) CreateZone(ctx context.Context, Req *core_api.CreateZoneReq, callOptions ...callopt.Option) (r *core_api.CreateZoneResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateZone(ctx, Req)
}

func (p *kContentClient) UpdateZone(ctx context.Context, Req *core_api.UpdateZoneReq, callOptions ...callopt.Option) (r *core_api.UpdateZoneResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateZone(ctx, Req)
}

func (p *kContentClient) GetZone(ctx context.Context, Req *core_api.GetZoneReq, callOptions ...callopt.Option) (r *core_api.GetZoneResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetZone(ctx, Req)
}

func (p *kContentClient) DeleteZone(ctx context.Context, Req *core_api.DeleteZoneReq, callOptions ...callopt.Option) (r *core_api.DeleteZoneResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteZone(ctx, Req)
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

func (p *kContentClient) AskUploadFile(ctx context.Context, Req *core_api.AskUploadFileReq, callOptions ...callopt.Option) (r *core_api.AskUploadFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AskUploadFile(ctx, Req)
}

func (p *kContentClient) AskDownloadFile(ctx context.Context, Req *core_api.AskDownloadFileReq, callOptions ...callopt.Option) (r *core_api.AskDownloadFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AskDownloadFile(ctx, Req)
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

func (p *kContentClient) GetRecommendByUser(ctx context.Context, Req *core_api.GetRecommendByUserReq, callOptions ...callopt.Option) (r *core_api.GetRecommendByUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRecommendByUser(ctx, Req)
}

func (p *kContentClient) GetRecommendByItem(ctx context.Context, Req *core_api.GetRecommendByItemReq, callOptions ...callopt.Option) (r *core_api.GetRecommendByItemResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRecommendByItem(ctx, Req)
}

func (p *kContentClient) CreateFeedBack(ctx context.Context, Req *core_api.CreateFeedBackReq, callOptions ...callopt.Option) (r *core_api.CreateFeedBackResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateFeedBack(ctx, Req)
}

func (p *kContentClient) GetPopularRecommend(ctx context.Context, Req *core_api.GetPopularRecommendReq, callOptions ...callopt.Option) (r *core_api.GetPopularRecommendResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPopularRecommend(ctx, Req)
}

func (p *kContentClient) GetLatestRecommend(ctx context.Context, Req *core_api.GetLatestRecommendReq, callOptions ...callopt.Option) (r *core_api.GetLatestRecommendResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLatestRecommend(ctx, Req)
}

func (p *kContentClient) CreateProduct(ctx context.Context, Req *core_api.CreateProductReq, callOptions ...callopt.Option) (r *core_api.CreateProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateProduct(ctx, Req)
}

func (p *kContentClient) GetProduct(ctx context.Context, Req *core_api.GetProductReq, callOptions ...callopt.Option) (r *core_api.GetProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProduct(ctx, Req)
}

func (p *kContentClient) GetProducts(ctx context.Context, Req *core_api.GetProductsReq, callOptions ...callopt.Option) (r *core_api.GetProductsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProducts(ctx, Req)
}

func (p *kContentClient) UpdateProduct(ctx context.Context, Req *core_api.UpdateProductReq, callOptions ...callopt.Option) (r *core_api.UpdateProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateProduct(ctx, Req)
}

func (p *kContentClient) DeleteProduct(ctx context.Context, Req *core_api.DeleteProductReq, callOptions ...callopt.Option) (r *core_api.DeleteProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteProduct(ctx, Req)
}
