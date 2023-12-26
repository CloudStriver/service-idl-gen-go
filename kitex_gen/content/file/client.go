// Code generated by Kitex v0.8.0. DO NOT EDIT.

package file

import (
	"context"
	content "github.com/CloudStriver/service-idl-gen-go/kitex_gen/content"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetCosInfo(ctx context.Context, Req *content.GetCosInfoReq, callOptions ...callopt.Option) (r *content.GetCosInfoResp, err error)
	GetFileByMd5(ctx context.Context, Req *content.GetFileByMd5Req, callOptions ...callopt.Option) (r *content.GetFileByMd5Resp, err error)
	GetFileInfoByFileId(ctx context.Context, Req *content.GetFileInfoByFileIdReq, callOptions ...callopt.Option) (r *content.GetFileInfoByFileIdResp, err error)
	GetFileByFileIdWithUserId(ctx context.Context, Req *content.GetFileByFileIdWithUserIdReq, callOptions ...callopt.Option) (r *content.GetFileByFileIdWithUserIdResp, err error)
	GetFileInfoByFileIds(ctx context.Context, Req *content.GetFileInfoByFileIdsReq, callOptions ...callopt.Option) (r *content.GetFileInfoByFileIdsResp, err error)
	GetFileInfoByFileIdsWithUserId(ctx context.Context, Req *content.GetFileInfoByFileIdsWithUserIdReq, callOptions ...callopt.Option) (r *content.GetFileInfoByFileIdsWithUserIdResp, err error)
	GetFilesByFatherId(ctx context.Context, Req *content.GetFilesByFatherIdReq, callOptions ...callopt.Option) (r *content.GetFilesByFatherIdResp, err error)
	GetFilesByFatherIdWithUserId(ctx context.Context, Req *content.GetFilesByFatherIdWithUserIdReq, callOptions ...callopt.Option) (r *content.GetFilesByFatherIdWithUserIdResp, err error)
	GetAllFileByFatherId(ctx context.Context, Req *content.GetAllFileByFatherIdReq, callOptions ...callopt.Option) (r *content.GetAllFileByFatherIdResp, err error)
	GetFileCount(ctx context.Context, Req *content.GetFileCountReq, callOptions ...callopt.Option) (r *content.GetFileCountResp, err error)
	GetFileType(ctx context.Context, Req *content.GetFileTypeReq, callOptions ...callopt.Option) (r *content.GetFileTypeReq, err error)
	GetFileInfoInPublicByFileId(ctx context.Context, Req *content.GetFileInfoInPublicByFileIdReq, callOptions ...callopt.Option) (r *content.GetFileInfoInPublicByFileIdResp, err error)
	GetPublicFileByFileIdWithUserId(ctx context.Context, Req *content.GetPublicFileByFileIdWithUserIdReq, callOptions ...callopt.Option) (r *content.GetPublicFileByFileIdWithUserIdResp, err error)
	GetPublicByFileIds(ctx context.Context, Req *content.GetPublicByFileIdsReq, callOptions ...callopt.Option) (r *content.GetPublicByFileIdsResp, err error)
	GetPublicFilesByUserId(ctx context.Context, Req *content.GetPublicFilesByUserIdReq, callOptions ...callopt.Option) (r *content.GetPublicFilesByUserIdResp, err error)
	GetPublicFilesByFatherId(ctx context.Context, Req *content.GetPublicFilesByFatherIdReq, callOptions ...callopt.Option) (r *content.GetPublicFilesByFatherIdResp, err error)
	GetAllPublicFileByFatherId(ctx context.Context, Req *content.GetAllPublicFileByFatherIdReq, callOptions ...callopt.Option) (r *content.GetAllPublicFileByFatherIdResp, err error)
	GetFileBySharingCode(ctx context.Context, Req *content.GetFileBySharingCodeReq, callOptions ...callopt.Option) (r *content.GetFileBySharingCodeResp, err error)
	CalFileSize(ctx context.Context, Req *content.CalFileSizeReq, callOptions ...callopt.Option) (r *content.CalFileSizeResp, err error)
	CalPublicFileSize(ctx context.Context, Req *content.CalPublicFileSizeReq, callOptions ...callopt.Option) (r *content.CalPublicFileSizeResp, err error)
	UploadFile(ctx context.Context, Req *content.UploadFileReq, callOptions ...callopt.Option) (r *content.UploadFileResp, err error)
	AskUploadFile(ctx context.Context, Req *content.AskUploadFileReq, callOptions ...callopt.Option) (r *content.AskUploadFileResp, err error)
	AskUploadFileRollback(ctx context.Context, Req *content.AskUploadFileReq, callOptions ...callopt.Option) (r *content.AskUploadFileResp, err error)
	DeleteExpiredFiles(ctx context.Context, Req *content.DeleteExpiredFilesReq, callOptions ...callopt.Option) (r *content.DeleteExpiredFilesResp, err error)
	DeleteExpiredShareCodes(ctx context.Context, Req *content.DeleteExpiredShareCodesReq, callOptions ...callopt.Option) (r *content.DeleteExpiredShareCodesResp, err error)
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
	return &kFileClient{
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

type kFileClient struct {
	*kClient
}

func (p *kFileClient) GetCosInfo(ctx context.Context, Req *content.GetCosInfoReq, callOptions ...callopt.Option) (r *content.GetCosInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCosInfo(ctx, Req)
}

func (p *kFileClient) GetFileByMd5(ctx context.Context, Req *content.GetFileByMd5Req, callOptions ...callopt.Option) (r *content.GetFileByMd5Resp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileByMd5(ctx, Req)
}

func (p *kFileClient) GetFileInfoByFileId(ctx context.Context, Req *content.GetFileInfoByFileIdReq, callOptions ...callopt.Option) (r *content.GetFileInfoByFileIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileInfoByFileId(ctx, Req)
}

func (p *kFileClient) GetFileByFileIdWithUserId(ctx context.Context, Req *content.GetFileByFileIdWithUserIdReq, callOptions ...callopt.Option) (r *content.GetFileByFileIdWithUserIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileByFileIdWithUserId(ctx, Req)
}

func (p *kFileClient) GetFileInfoByFileIds(ctx context.Context, Req *content.GetFileInfoByFileIdsReq, callOptions ...callopt.Option) (r *content.GetFileInfoByFileIdsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileInfoByFileIds(ctx, Req)
}

func (p *kFileClient) GetFileInfoByFileIdsWithUserId(ctx context.Context, Req *content.GetFileInfoByFileIdsWithUserIdReq, callOptions ...callopt.Option) (r *content.GetFileInfoByFileIdsWithUserIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileInfoByFileIdsWithUserId(ctx, Req)
}

func (p *kFileClient) GetFilesByFatherId(ctx context.Context, Req *content.GetFilesByFatherIdReq, callOptions ...callopt.Option) (r *content.GetFilesByFatherIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFilesByFatherId(ctx, Req)
}

func (p *kFileClient) GetFilesByFatherIdWithUserId(ctx context.Context, Req *content.GetFilesByFatherIdWithUserIdReq, callOptions ...callopt.Option) (r *content.GetFilesByFatherIdWithUserIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFilesByFatherIdWithUserId(ctx, Req)
}

func (p *kFileClient) GetAllFileByFatherId(ctx context.Context, Req *content.GetAllFileByFatherIdReq, callOptions ...callopt.Option) (r *content.GetAllFileByFatherIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllFileByFatherId(ctx, Req)
}

func (p *kFileClient) GetFileCount(ctx context.Context, Req *content.GetFileCountReq, callOptions ...callopt.Option) (r *content.GetFileCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileCount(ctx, Req)
}

func (p *kFileClient) GetFileType(ctx context.Context, Req *content.GetFileTypeReq, callOptions ...callopt.Option) (r *content.GetFileTypeReq, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileType(ctx, Req)
}

func (p *kFileClient) GetFileInfoInPublicByFileId(ctx context.Context, Req *content.GetFileInfoInPublicByFileIdReq, callOptions ...callopt.Option) (r *content.GetFileInfoInPublicByFileIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileInfoInPublicByFileId(ctx, Req)
}

func (p *kFileClient) GetPublicFileByFileIdWithUserId(ctx context.Context, Req *content.GetPublicFileByFileIdWithUserIdReq, callOptions ...callopt.Option) (r *content.GetPublicFileByFileIdWithUserIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublicFileByFileIdWithUserId(ctx, Req)
}

func (p *kFileClient) GetPublicByFileIds(ctx context.Context, Req *content.GetPublicByFileIdsReq, callOptions ...callopt.Option) (r *content.GetPublicByFileIdsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublicByFileIds(ctx, Req)
}

func (p *kFileClient) GetPublicFilesByUserId(ctx context.Context, Req *content.GetPublicFilesByUserIdReq, callOptions ...callopt.Option) (r *content.GetPublicFilesByUserIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublicFilesByUserId(ctx, Req)
}

func (p *kFileClient) GetPublicFilesByFatherId(ctx context.Context, Req *content.GetPublicFilesByFatherIdReq, callOptions ...callopt.Option) (r *content.GetPublicFilesByFatherIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPublicFilesByFatherId(ctx, Req)
}

func (p *kFileClient) GetAllPublicFileByFatherId(ctx context.Context, Req *content.GetAllPublicFileByFatherIdReq, callOptions ...callopt.Option) (r *content.GetAllPublicFileByFatherIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllPublicFileByFatherId(ctx, Req)
}

func (p *kFileClient) GetFileBySharingCode(ctx context.Context, Req *content.GetFileBySharingCodeReq, callOptions ...callopt.Option) (r *content.GetFileBySharingCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFileBySharingCode(ctx, Req)
}

func (p *kFileClient) CalFileSize(ctx context.Context, Req *content.CalFileSizeReq, callOptions ...callopt.Option) (r *content.CalFileSizeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CalFileSize(ctx, Req)
}

func (p *kFileClient) CalPublicFileSize(ctx context.Context, Req *content.CalPublicFileSizeReq, callOptions ...callopt.Option) (r *content.CalPublicFileSizeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CalPublicFileSize(ctx, Req)
}

func (p *kFileClient) UploadFile(ctx context.Context, Req *content.UploadFileReq, callOptions ...callopt.Option) (r *content.UploadFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadFile(ctx, Req)
}

func (p *kFileClient) AskUploadFile(ctx context.Context, Req *content.AskUploadFileReq, callOptions ...callopt.Option) (r *content.AskUploadFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AskUploadFile(ctx, Req)
}

func (p *kFileClient) AskUploadFileRollback(ctx context.Context, Req *content.AskUploadFileReq, callOptions ...callopt.Option) (r *content.AskUploadFileResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AskUploadFileRollback(ctx, Req)
}

func (p *kFileClient) DeleteExpiredFiles(ctx context.Context, Req *content.DeleteExpiredFilesReq, callOptions ...callopt.Option) (r *content.DeleteExpiredFilesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteExpiredFiles(ctx, Req)
}

func (p *kFileClient) DeleteExpiredShareCodes(ctx context.Context, Req *content.DeleteExpiredShareCodesReq, callOptions ...callopt.Option) (r *content.DeleteExpiredShareCodesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteExpiredShareCodes(ctx, Req)
}
