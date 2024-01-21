// Code generated by Kitex v0.8.0. DO NOT EDIT.

package contentservice

import (
	"context"
	content "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/content"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetFileIsExist(ctx context.Context, Req *content.GetFileIsExistReq, callOptions ...callopt.Option) (r *content.GetFileIsExistResp, err error)
	GetFile(ctx context.Context, Req *content.GetFileReq, callOptions ...callopt.Option) (r *content.GetFileResp, err error)
	GetFileList(ctx context.Context, Req *content.GetFileListReq, callOptions ...callopt.Option) (r *content.GetFileListResp, err error)
	GetFileCount(ctx context.Context, Req *content.GetFileCountReq, callOptions ...callopt.Option) (r *content.GetFileCountResp, err error)
	GetFileBySharingCode(ctx context.Context, Req *content.GetFileBySharingCodeReq, callOptions ...callopt.Option) (r *content.GetFileBySharingCodeResp, err error)
	GetFolderSize(ctx context.Context, Req *content.GetFolderSizeReq, callOptions ...callopt.Option) (r *content.GetFolderSizeResp, err error)
	CreateFolder(ctx context.Context, Req *content.CreateFolderReq, callOptions ...callopt.Option) (r *content.CreateFolderResp, err error)
	UpdateFile(ctx context.Context, Req *content.UpdateFileReq, callOptions ...callopt.Option) (r *content.UpdateFileResp, err error)
	MoveFile(ctx context.Context, Req *content.MoveFileReq, callOptions ...callopt.Option) (r *content.MoveFileResp, err error)
	SaveFileToPrivateSpace(ctx context.Context, Req *content.SaveFileToPrivateSpaceReq, callOptions ...callopt.Option) (r *content.SaveFileToPrivateSpaceResp, err error)
	AddFileToPublicSpace(ctx context.Context, Req *content.AddFileToPublicSpaceReq, callOptions ...callopt.Option) (r *content.AddFileToPublicSpaceResp, err error)
	DeleteFile(ctx context.Context, Req *content.DeleteFileReq, callOptions ...callopt.Option) (r *content.DeleteFileResp, err error)
	RecoverRecycleBinFile(ctx context.Context, Req *content.RecoverRecycleBinFileReq, callOptions ...callopt.Option) (r *content.RecoverRecycleBinFileResp, err error)
	GetLabel(ctx context.Context, Req *content.GetLabelReq, callOptions ...callopt.Option) (r *content.GetLabelResp, err error)
	CreateLabel(ctx context.Context, Req *content.CreateLabelReq, callOptions ...callopt.Option) (r *content.CreateLabelResp, err error)
	UpdateLabel(ctx context.Context, Req *content.UpdateLabelReq, callOptions ...callopt.Option) (r *content.UpdateLabelResp, err error)
	DeleteLabel(ctx context.Context, Req *content.DeleteLabelReq, callOptions ...callopt.Option) (r *content.DeleteLabelResp, err error)
	GetShareList(ctx context.Context, Req *content.GetShareListReq, callOptions ...callopt.Option) (r *content.GetShareListResp, err error)
	CreateShareCode(ctx context.Context, Req *content.CreateShareCodeReq, callOptions ...callopt.Option) (r *content.CreateShareCodeResp, err error)
	UpdateShareCode(ctx context.Context, Req *content.UpdateShareCodeReq, callOptions ...callopt.Option) (r *content.UpdateShareCodeResp, err error)
	DeleteShareCode(ctx context.Context, Req *content.DeleteShareCodeReq, callOptions ...callopt.Option) (r *content.DeleteShareCodeResp, err error)
	ParsingShareCode(ctx context.Context, Req *content.ParsingShareCodeReq, callOptions ...callopt.Option) (r *content.ParsingShareCodeResp, err error)
	UpdateUser(ctx context.Context, Req *content.UpdateUserReq, callOptions ...callopt.Option) (r *content.UpdateUserResp, err error)
	GetUser(ctx context.Context, Req *content.GetUserReq, callOptions ...callopt.Option) (r *content.GetUserResp, err error)
	SearchUser(ctx context.Context, Req *content.SearchUserReq, callOptions ...callopt.Option) (r *content.SearchUserResp, err error)
	CreateUser(ctx context.Context, Req *content.CreateUserReq, callOptions ...callopt.Option) (r *content.CreateUserResp, err error)
	DeleteUser(ctx context.Context, Req *content.DeleteUserReq, callOptions ...callopt.Option) (r *content.DeleteUserResp, err error)
	CreatePost(ctx context.Context, Req *content.CreatePostReq, callOptions ...callopt.Option) (r *content.CreatePostResp, err error)
	DeletePost(ctx context.Context, Req *content.DeletePostReq, callOptions ...callopt.Option) (r *content.DeletePostResp, err error)
	UpdatePost(ctx context.Context, Req *content.UpdatePostReq, callOptions ...callopt.Option) (r *content.UpdatePostResp, err error)
	GetPost(ctx context.Context, Req *content.GetPostReq, callOptions ...callopt.Option) (r *content.GetPostResp, err error)
	GetPosts(ctx context.Context, Req *content.GetPostsReq, callOptions ...callopt.Option) (r *content.GetPostsResp, err error)
	CreateProduct(ctx context.Context, Req *content.CreateProductReq, callOptions ...callopt.Option) (r *content.CreateProductResp, err error)
	DeleteProduct(ctx context.Context, Req *content.DeleteProductReq, callOptions ...callopt.Option) (r *content.DeleteProductResp, err error)
	UpdateProduct(ctx context.Context, Req *content.UpdateProductReq, callOptions ...callopt.Option) (r *content.UpdateProductResp, err error)
	GetProduct(ctx context.Context, Req *content.GetProductReq, callOptions ...callopt.Option) (r *content.GetProductResp, err error)
	GetProducts(ctx context.Context, Req *content.GetProductsReq, callOptions ...callopt.Option) (r *content.GetProductsResp, err error)
	CreateCoupon(ctx context.Context, Req *content.CreateCouponReq, callOptions ...callopt.Option) (r *content.CreateCouponResp, err error)
	DeleteCoupon(ctx context.Context, Req *content.DeleteCouponReq, callOptions ...callopt.Option) (r *content.DeleteCouponResp, err error)
	UpdateCoupon(ctx context.Context, Req *content.UpdateCouponReq, callOptions ...callopt.Option) (r *content.UpdateCouponResp, err error)
	GetCoupon(ctx context.Context, Req *content.GetCouponReq, callOptions ...callopt.Option) (r *content.GetCouponResp, err error)
	GetCoupons(ctx context.Context, Req *content.GetCouponsReq, callOptions ...callopt.Option) (r *content.GetCouponsResp, err error)
	CreateOrder(ctx context.Context, Req *content.CreateOrderReq, callOptions ...callopt.Option) (r *content.CreateOrderResp, err error)
	DeleteOrder(ctx context.Context, Req *content.DeleteOrderReq, callOptions ...callopt.Option) (r *content.DeleteOrderResp, err error)
	UpdateOrder(ctx context.Context, Req *content.UpdateOrderReq, callOptions ...callopt.Option) (r *content.UpdateOrderResp, err error)
	GetOrder(ctx context.Context, Req *content.GetOrderReq, callOptions ...callopt.Option) (r *content.GetOrderResp, err error)
	GetOrders(ctx context.Context, Req *content.GetOrdersReq, callOptions ...callopt.Option) (r *content.GetOrdersResp, err error)
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
	return &kContentServiceClient{
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

type kContentServiceClient struct {
	*kClient
}

func (p *kContentServiceClient) GetFileIsExist(ctx context.Context, Req *content.GetFileIsExistReq, callOptions ...callopt.Option) (r *content.GetFileIsExistResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileIsExist(ctx, Req)
}

func (p *kContentServiceClient) GetFile(ctx context.Context, Req *content.GetFileReq, callOptions ...callopt.Option) (r *content.GetFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFile(ctx, Req)
}

func (p *kContentServiceClient) GetFileList(ctx context.Context, Req *content.GetFileListReq, callOptions ...callopt.Option) (r *content.GetFileListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileList(ctx, Req)
}

func (p *kContentServiceClient) GetFileCount(ctx context.Context, Req *content.GetFileCountReq, callOptions ...callopt.Option) (r *content.GetFileCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileCount(ctx, Req)
}

func (p *kContentServiceClient) GetFileBySharingCode(ctx context.Context, Req *content.GetFileBySharingCodeReq, callOptions ...callopt.Option) (r *content.GetFileBySharingCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileBySharingCode(ctx, Req)
}

func (p *kContentServiceClient) GetFolderSize(ctx context.Context, Req *content.GetFolderSizeReq, callOptions ...callopt.Option) (r *content.GetFolderSizeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFolderSize(ctx, Req)
}

func (p *kContentServiceClient) CreateFolder(ctx context.Context, Req *content.CreateFolderReq, callOptions ...callopt.Option) (r *content.CreateFolderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateFolder(ctx, Req)
}

func (p *kContentServiceClient) UpdateFile(ctx context.Context, Req *content.UpdateFileReq, callOptions ...callopt.Option) (r *content.UpdateFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateFile(ctx, Req)
}

func (p *kContentServiceClient) MoveFile(ctx context.Context, Req *content.MoveFileReq, callOptions ...callopt.Option) (r *content.MoveFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MoveFile(ctx, Req)
}

func (p *kContentServiceClient) SaveFileToPrivateSpace(ctx context.Context, Req *content.SaveFileToPrivateSpaceReq, callOptions ...callopt.Option) (r *content.SaveFileToPrivateSpaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SaveFileToPrivateSpace(ctx, Req)
}

func (p *kContentServiceClient) AddFileToPublicSpace(ctx context.Context, Req *content.AddFileToPublicSpaceReq, callOptions ...callopt.Option) (r *content.AddFileToPublicSpaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddFileToPublicSpace(ctx, Req)
}

func (p *kContentServiceClient) DeleteFile(ctx context.Context, Req *content.DeleteFileReq, callOptions ...callopt.Option) (r *content.DeleteFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteFile(ctx, Req)
}

func (p *kContentServiceClient) RecoverRecycleBinFile(ctx context.Context, Req *content.RecoverRecycleBinFileReq, callOptions ...callopt.Option) (r *content.RecoverRecycleBinFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RecoverRecycleBinFile(ctx, Req)
}

func (p *kContentServiceClient) GetLabel(ctx context.Context, Req *content.GetLabelReq, callOptions ...callopt.Option) (r *content.GetLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLabel(ctx, Req)
}

func (p *kContentServiceClient) CreateLabel(ctx context.Context, Req *content.CreateLabelReq, callOptions ...callopt.Option) (r *content.CreateLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateLabel(ctx, Req)
}

func (p *kContentServiceClient) UpdateLabel(ctx context.Context, Req *content.UpdateLabelReq, callOptions ...callopt.Option) (r *content.UpdateLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateLabel(ctx, Req)
}

func (p *kContentServiceClient) DeleteLabel(ctx context.Context, Req *content.DeleteLabelReq, callOptions ...callopt.Option) (r *content.DeleteLabelResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteLabel(ctx, Req)
}

func (p *kContentServiceClient) GetShareList(ctx context.Context, Req *content.GetShareListReq, callOptions ...callopt.Option) (r *content.GetShareListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetShareList(ctx, Req)
}

func (p *kContentServiceClient) CreateShareCode(ctx context.Context, Req *content.CreateShareCodeReq, callOptions ...callopt.Option) (r *content.CreateShareCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateShareCode(ctx, Req)
}

func (p *kContentServiceClient) UpdateShareCode(ctx context.Context, Req *content.UpdateShareCodeReq, callOptions ...callopt.Option) (r *content.UpdateShareCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateShareCode(ctx, Req)
}

func (p *kContentServiceClient) DeleteShareCode(ctx context.Context, Req *content.DeleteShareCodeReq, callOptions ...callopt.Option) (r *content.DeleteShareCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteShareCode(ctx, Req)
}

func (p *kContentServiceClient) ParsingShareCode(ctx context.Context, Req *content.ParsingShareCodeReq, callOptions ...callopt.Option) (r *content.ParsingShareCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ParsingShareCode(ctx, Req)
}

func (p *kContentServiceClient) UpdateUser(ctx context.Context, Req *content.UpdateUserReq, callOptions ...callopt.Option) (r *content.UpdateUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUser(ctx, Req)
}

func (p *kContentServiceClient) GetUser(ctx context.Context, Req *content.GetUserReq, callOptions ...callopt.Option) (r *content.GetUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUser(ctx, Req)
}

func (p *kContentServiceClient) SearchUser(ctx context.Context, Req *content.SearchUserReq, callOptions ...callopt.Option) (r *content.SearchUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchUser(ctx, Req)
}

func (p *kContentServiceClient) CreateUser(ctx context.Context, Req *content.CreateUserReq, callOptions ...callopt.Option) (r *content.CreateUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, Req)
}

func (p *kContentServiceClient) DeleteUser(ctx context.Context, Req *content.DeleteUserReq, callOptions ...callopt.Option) (r *content.DeleteUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteUser(ctx, Req)
}

func (p *kContentServiceClient) CreatePost(ctx context.Context, Req *content.CreatePostReq, callOptions ...callopt.Option) (r *content.CreatePostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreatePost(ctx, Req)
}

func (p *kContentServiceClient) DeletePost(ctx context.Context, Req *content.DeletePostReq, callOptions ...callopt.Option) (r *content.DeletePostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeletePost(ctx, Req)
}

func (p *kContentServiceClient) UpdatePost(ctx context.Context, Req *content.UpdatePostReq, callOptions ...callopt.Option) (r *content.UpdatePostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePost(ctx, Req)
}

func (p *kContentServiceClient) GetPost(ctx context.Context, Req *content.GetPostReq, callOptions ...callopt.Option) (r *content.GetPostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPost(ctx, Req)
}

func (p *kContentServiceClient) GetPosts(ctx context.Context, Req *content.GetPostsReq, callOptions ...callopt.Option) (r *content.GetPostsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPosts(ctx, Req)
}

func (p *kContentServiceClient) CreateProduct(ctx context.Context, Req *content.CreateProductReq, callOptions ...callopt.Option) (r *content.CreateProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateProduct(ctx, Req)
}

func (p *kContentServiceClient) DeleteProduct(ctx context.Context, Req *content.DeleteProductReq, callOptions ...callopt.Option) (r *content.DeleteProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteProduct(ctx, Req)
}

func (p *kContentServiceClient) UpdateProduct(ctx context.Context, Req *content.UpdateProductReq, callOptions ...callopt.Option) (r *content.UpdateProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateProduct(ctx, Req)
}

func (p *kContentServiceClient) GetProduct(ctx context.Context, Req *content.GetProductReq, callOptions ...callopt.Option) (r *content.GetProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProduct(ctx, Req)
}

func (p *kContentServiceClient) GetProducts(ctx context.Context, Req *content.GetProductsReq, callOptions ...callopt.Option) (r *content.GetProductsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProducts(ctx, Req)
}

func (p *kContentServiceClient) CreateCoupon(ctx context.Context, Req *content.CreateCouponReq, callOptions ...callopt.Option) (r *content.CreateCouponResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateCoupon(ctx, Req)
}

func (p *kContentServiceClient) DeleteCoupon(ctx context.Context, Req *content.DeleteCouponReq, callOptions ...callopt.Option) (r *content.DeleteCouponResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCoupon(ctx, Req)
}

func (p *kContentServiceClient) UpdateCoupon(ctx context.Context, Req *content.UpdateCouponReq, callOptions ...callopt.Option) (r *content.UpdateCouponResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateCoupon(ctx, Req)
}

func (p *kContentServiceClient) GetCoupon(ctx context.Context, Req *content.GetCouponReq, callOptions ...callopt.Option) (r *content.GetCouponResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCoupon(ctx, Req)
}

func (p *kContentServiceClient) GetCoupons(ctx context.Context, Req *content.GetCouponsReq, callOptions ...callopt.Option) (r *content.GetCouponsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCoupons(ctx, Req)
}

func (p *kContentServiceClient) CreateOrder(ctx context.Context, Req *content.CreateOrderReq, callOptions ...callopt.Option) (r *content.CreateOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateOrder(ctx, Req)
}

func (p *kContentServiceClient) DeleteOrder(ctx context.Context, Req *content.DeleteOrderReq, callOptions ...callopt.Option) (r *content.DeleteOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteOrder(ctx, Req)
}

func (p *kContentServiceClient) UpdateOrder(ctx context.Context, Req *content.UpdateOrderReq, callOptions ...callopt.Option) (r *content.UpdateOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateOrder(ctx, Req)
}

func (p *kContentServiceClient) GetOrder(ctx context.Context, Req *content.GetOrderReq, callOptions ...callopt.Option) (r *content.GetOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOrder(ctx, Req)
}

func (p *kContentServiceClient) GetOrders(ctx context.Context, Req *content.GetOrdersReq, callOptions ...callopt.Option) (r *content.GetOrdersResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOrders(ctx, Req)
}
