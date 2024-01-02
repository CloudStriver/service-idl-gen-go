// Code generated by Kitex v0.8.0. DO NOT EDIT.

package auth

import (
	"context"
	core_api "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/core_api"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return authServiceInfo
}

var authServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "auth"
	handlerType := (*core_api.Auth)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register":              kitex.NewMethodInfo(registerHandler, newRegisterArgs, newRegisterResult, false),
		"EmailLogin":            kitex.NewMethodInfo(emailLoginHandler, newEmailLoginArgs, newEmailLoginResult, false),
		"GithubLogin":           kitex.NewMethodInfo(githubLoginHandler, newGithubLoginArgs, newGithubLoginResult, false),
		"GiteeLogin":            kitex.NewMethodInfo(giteeLoginHandler, newGiteeLoginArgs, newGiteeLoginResult, false),
		"RefreshToken":          kitex.NewMethodInfo(refreshTokenHandler, newRefreshTokenArgs, newRefreshTokenResult, false),
		"SendEmail":             kitex.NewMethodInfo(sendEmailHandler, newSendEmailArgs, newSendEmailResult, false),
		"GetCaptcha":            kitex.NewMethodInfo(getCaptchaHandler, newGetCaptchaArgs, newGetCaptchaResult, false),
		"SetPasswordByEmail":    kitex.NewMethodInfo(setPasswordByEmailHandler, newSetPasswordByEmailArgs, newSetPasswordByEmailResult, false),
		"SetPasswordByPassword": kitex.NewMethodInfo(setPasswordByPasswordHandler, newSetPasswordByPasswordArgs, newSetPasswordByPasswordResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "cloudmind.core_api",
		"ServiceFilePath": ``,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.RegisterReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).Register(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RegisterArgs:
		success, err := handler.(core_api.Auth).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
	}
	return nil
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *core_api.RegisterReq
}

func (p *RegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.RegisterReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(core_api.RegisterReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *core_api.RegisterReq

func (p *RegisterArgs) GetReq() *core_api.RegisterReq {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RegisterResult struct {
	Success *core_api.RegisterResp
}

var RegisterResult_Success_DEFAULT *core_api.RegisterResp

func (p *RegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.RegisterResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(core_api.RegisterResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *core_api.RegisterResp {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.RegisterResp)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RegisterResult) GetResult() interface{} {
	return p.Success
}

func emailLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.EmailLoginReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).EmailLogin(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *EmailLoginArgs:
		success, err := handler.(core_api.Auth).EmailLogin(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*EmailLoginResult)
		realResult.Success = success
	}
	return nil
}
func newEmailLoginArgs() interface{} {
	return &EmailLoginArgs{}
}

func newEmailLoginResult() interface{} {
	return &EmailLoginResult{}
}

type EmailLoginArgs struct {
	Req *core_api.EmailLoginReq
}

func (p *EmailLoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.EmailLoginReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *EmailLoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *EmailLoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *EmailLoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *EmailLoginArgs) Unmarshal(in []byte) error {
	msg := new(core_api.EmailLoginReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var EmailLoginArgs_Req_DEFAULT *core_api.EmailLoginReq

func (p *EmailLoginArgs) GetReq() *core_api.EmailLoginReq {
	if !p.IsSetReq() {
		return EmailLoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *EmailLoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *EmailLoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type EmailLoginResult struct {
	Success *core_api.EmailLoginResp
}

var EmailLoginResult_Success_DEFAULT *core_api.EmailLoginResp

func (p *EmailLoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.EmailLoginResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *EmailLoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *EmailLoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *EmailLoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *EmailLoginResult) Unmarshal(in []byte) error {
	msg := new(core_api.EmailLoginResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *EmailLoginResult) GetSuccess() *core_api.EmailLoginResp {
	if !p.IsSetSuccess() {
		return EmailLoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *EmailLoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.EmailLoginResp)
}

func (p *EmailLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EmailLoginResult) GetResult() interface{} {
	return p.Success
}

func githubLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GithubLoginReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).GithubLogin(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GithubLoginArgs:
		success, err := handler.(core_api.Auth).GithubLogin(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GithubLoginResult)
		realResult.Success = success
	}
	return nil
}
func newGithubLoginArgs() interface{} {
	return &GithubLoginArgs{}
}

func newGithubLoginResult() interface{} {
	return &GithubLoginResult{}
}

type GithubLoginArgs struct {
	Req *core_api.GithubLoginReq
}

func (p *GithubLoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GithubLoginReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GithubLoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GithubLoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GithubLoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GithubLoginArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GithubLoginReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GithubLoginArgs_Req_DEFAULT *core_api.GithubLoginReq

func (p *GithubLoginArgs) GetReq() *core_api.GithubLoginReq {
	if !p.IsSetReq() {
		return GithubLoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GithubLoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GithubLoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GithubLoginResult struct {
	Success *core_api.GithubLoginResp
}

var GithubLoginResult_Success_DEFAULT *core_api.GithubLoginResp

func (p *GithubLoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GithubLoginResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GithubLoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GithubLoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GithubLoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GithubLoginResult) Unmarshal(in []byte) error {
	msg := new(core_api.GithubLoginResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GithubLoginResult) GetSuccess() *core_api.GithubLoginResp {
	if !p.IsSetSuccess() {
		return GithubLoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GithubLoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GithubLoginResp)
}

func (p *GithubLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GithubLoginResult) GetResult() interface{} {
	return p.Success
}

func giteeLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GiteeLoginReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).GiteeLogin(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GiteeLoginArgs:
		success, err := handler.(core_api.Auth).GiteeLogin(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GiteeLoginResult)
		realResult.Success = success
	}
	return nil
}
func newGiteeLoginArgs() interface{} {
	return &GiteeLoginArgs{}
}

func newGiteeLoginResult() interface{} {
	return &GiteeLoginResult{}
}

type GiteeLoginArgs struct {
	Req *core_api.GiteeLoginReq
}

func (p *GiteeLoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GiteeLoginReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GiteeLoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GiteeLoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GiteeLoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GiteeLoginArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GiteeLoginReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GiteeLoginArgs_Req_DEFAULT *core_api.GiteeLoginReq

func (p *GiteeLoginArgs) GetReq() *core_api.GiteeLoginReq {
	if !p.IsSetReq() {
		return GiteeLoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GiteeLoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GiteeLoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GiteeLoginResult struct {
	Success *core_api.GiteeLoginResp
}

var GiteeLoginResult_Success_DEFAULT *core_api.GiteeLoginResp

func (p *GiteeLoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GiteeLoginResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GiteeLoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GiteeLoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GiteeLoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GiteeLoginResult) Unmarshal(in []byte) error {
	msg := new(core_api.GiteeLoginResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GiteeLoginResult) GetSuccess() *core_api.GiteeLoginResp {
	if !p.IsSetSuccess() {
		return GiteeLoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GiteeLoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GiteeLoginResp)
}

func (p *GiteeLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GiteeLoginResult) GetResult() interface{} {
	return p.Success
}

func refreshTokenHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.RefreshTokenReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).RefreshToken(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RefreshTokenArgs:
		success, err := handler.(core_api.Auth).RefreshToken(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RefreshTokenResult)
		realResult.Success = success
	}
	return nil
}
func newRefreshTokenArgs() interface{} {
	return &RefreshTokenArgs{}
}

func newRefreshTokenResult() interface{} {
	return &RefreshTokenResult{}
}

type RefreshTokenArgs struct {
	Req *core_api.RefreshTokenReq
}

func (p *RefreshTokenArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.RefreshTokenReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RefreshTokenArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RefreshTokenArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RefreshTokenArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RefreshTokenArgs) Unmarshal(in []byte) error {
	msg := new(core_api.RefreshTokenReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RefreshTokenArgs_Req_DEFAULT *core_api.RefreshTokenReq

func (p *RefreshTokenArgs) GetReq() *core_api.RefreshTokenReq {
	if !p.IsSetReq() {
		return RefreshTokenArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RefreshTokenArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RefreshTokenArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RefreshTokenResult struct {
	Success *core_api.RefreshTokenResp
}

var RefreshTokenResult_Success_DEFAULT *core_api.RefreshTokenResp

func (p *RefreshTokenResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.RefreshTokenResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RefreshTokenResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RefreshTokenResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RefreshTokenResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RefreshTokenResult) Unmarshal(in []byte) error {
	msg := new(core_api.RefreshTokenResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RefreshTokenResult) GetSuccess() *core_api.RefreshTokenResp {
	if !p.IsSetSuccess() {
		return RefreshTokenResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RefreshTokenResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.RefreshTokenResp)
}

func (p *RefreshTokenResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RefreshTokenResult) GetResult() interface{} {
	return p.Success
}

func sendEmailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.SendEmailReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).SendEmail(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SendEmailArgs:
		success, err := handler.(core_api.Auth).SendEmail(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SendEmailResult)
		realResult.Success = success
	}
	return nil
}
func newSendEmailArgs() interface{} {
	return &SendEmailArgs{}
}

func newSendEmailResult() interface{} {
	return &SendEmailResult{}
}

type SendEmailArgs struct {
	Req *core_api.SendEmailReq
}

func (p *SendEmailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.SendEmailReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SendEmailArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SendEmailArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SendEmailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SendEmailArgs) Unmarshal(in []byte) error {
	msg := new(core_api.SendEmailReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendEmailArgs_Req_DEFAULT *core_api.SendEmailReq

func (p *SendEmailArgs) GetReq() *core_api.SendEmailReq {
	if !p.IsSetReq() {
		return SendEmailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendEmailArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SendEmailArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SendEmailResult struct {
	Success *core_api.SendEmailResp
}

var SendEmailResult_Success_DEFAULT *core_api.SendEmailResp

func (p *SendEmailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.SendEmailResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SendEmailResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SendEmailResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SendEmailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SendEmailResult) Unmarshal(in []byte) error {
	msg := new(core_api.SendEmailResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendEmailResult) GetSuccess() *core_api.SendEmailResp {
	if !p.IsSetSuccess() {
		return SendEmailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendEmailResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.SendEmailResp)
}

func (p *SendEmailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendEmailResult) GetResult() interface{} {
	return p.Success
}

func getCaptchaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetCaptchaReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).GetCaptcha(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetCaptchaArgs:
		success, err := handler.(core_api.Auth).GetCaptcha(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCaptchaResult)
		realResult.Success = success
	}
	return nil
}
func newGetCaptchaArgs() interface{} {
	return &GetCaptchaArgs{}
}

func newGetCaptchaResult() interface{} {
	return &GetCaptchaResult{}
}

type GetCaptchaArgs struct {
	Req *core_api.GetCaptchaReq
}

func (p *GetCaptchaArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetCaptchaReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetCaptchaArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetCaptchaArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetCaptchaArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetCaptchaArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetCaptchaReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCaptchaArgs_Req_DEFAULT *core_api.GetCaptchaReq

func (p *GetCaptchaArgs) GetReq() *core_api.GetCaptchaReq {
	if !p.IsSetReq() {
		return GetCaptchaArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCaptchaArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetCaptchaArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetCaptchaResult struct {
	Success *core_api.GetCaptchaResp
}

var GetCaptchaResult_Success_DEFAULT *core_api.GetCaptchaResp

func (p *GetCaptchaResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetCaptchaResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetCaptchaResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetCaptchaResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetCaptchaResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetCaptchaResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetCaptchaResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCaptchaResult) GetSuccess() *core_api.GetCaptchaResp {
	if !p.IsSetSuccess() {
		return GetCaptchaResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCaptchaResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetCaptchaResp)
}

func (p *GetCaptchaResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetCaptchaResult) GetResult() interface{} {
	return p.Success
}

func setPasswordByEmailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.SetPasswordByEmailReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).SetPasswordByEmail(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SetPasswordByEmailArgs:
		success, err := handler.(core_api.Auth).SetPasswordByEmail(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SetPasswordByEmailResult)
		realResult.Success = success
	}
	return nil
}
func newSetPasswordByEmailArgs() interface{} {
	return &SetPasswordByEmailArgs{}
}

func newSetPasswordByEmailResult() interface{} {
	return &SetPasswordByEmailResult{}
}

type SetPasswordByEmailArgs struct {
	Req *core_api.SetPasswordByEmailReq
}

func (p *SetPasswordByEmailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.SetPasswordByEmailReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SetPasswordByEmailArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SetPasswordByEmailArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SetPasswordByEmailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SetPasswordByEmailArgs) Unmarshal(in []byte) error {
	msg := new(core_api.SetPasswordByEmailReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SetPasswordByEmailArgs_Req_DEFAULT *core_api.SetPasswordByEmailReq

func (p *SetPasswordByEmailArgs) GetReq() *core_api.SetPasswordByEmailReq {
	if !p.IsSetReq() {
		return SetPasswordByEmailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetPasswordByEmailArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SetPasswordByEmailArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SetPasswordByEmailResult struct {
	Success *core_api.SetPasswordByEmailResp
}

var SetPasswordByEmailResult_Success_DEFAULT *core_api.SetPasswordByEmailResp

func (p *SetPasswordByEmailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.SetPasswordByEmailResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SetPasswordByEmailResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SetPasswordByEmailResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SetPasswordByEmailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SetPasswordByEmailResult) Unmarshal(in []byte) error {
	msg := new(core_api.SetPasswordByEmailResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetPasswordByEmailResult) GetSuccess() *core_api.SetPasswordByEmailResp {
	if !p.IsSetSuccess() {
		return SetPasswordByEmailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetPasswordByEmailResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.SetPasswordByEmailResp)
}

func (p *SetPasswordByEmailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetPasswordByEmailResult) GetResult() interface{} {
	return p.Success
}

func setPasswordByPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.SetPasswordByPasswordReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).SetPasswordByPassword(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SetPasswordByPasswordArgs:
		success, err := handler.(core_api.Auth).SetPasswordByPassword(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SetPasswordByPasswordResult)
		realResult.Success = success
	}
	return nil
}
func newSetPasswordByPasswordArgs() interface{} {
	return &SetPasswordByPasswordArgs{}
}

func newSetPasswordByPasswordResult() interface{} {
	return &SetPasswordByPasswordResult{}
}

type SetPasswordByPasswordArgs struct {
	Req *core_api.SetPasswordByPasswordReq
}

func (p *SetPasswordByPasswordArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.SetPasswordByPasswordReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SetPasswordByPasswordArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SetPasswordByPasswordArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SetPasswordByPasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SetPasswordByPasswordArgs) Unmarshal(in []byte) error {
	msg := new(core_api.SetPasswordByPasswordReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SetPasswordByPasswordArgs_Req_DEFAULT *core_api.SetPasswordByPasswordReq

func (p *SetPasswordByPasswordArgs) GetReq() *core_api.SetPasswordByPasswordReq {
	if !p.IsSetReq() {
		return SetPasswordByPasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetPasswordByPasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SetPasswordByPasswordArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SetPasswordByPasswordResult struct {
	Success *core_api.SetPasswordByPasswordReq
}

var SetPasswordByPasswordResult_Success_DEFAULT *core_api.SetPasswordByPasswordReq

func (p *SetPasswordByPasswordResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.SetPasswordByPasswordReq)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SetPasswordByPasswordResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SetPasswordByPasswordResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SetPasswordByPasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SetPasswordByPasswordResult) Unmarshal(in []byte) error {
	msg := new(core_api.SetPasswordByPasswordReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetPasswordByPasswordResult) GetSuccess() *core_api.SetPasswordByPasswordReq {
	if !p.IsSetSuccess() {
		return SetPasswordByPasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetPasswordByPasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.SetPasswordByPasswordReq)
}

func (p *SetPasswordByPasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetPasswordByPasswordResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, Req *core_api.RegisterReq) (r *core_api.RegisterResp, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EmailLogin(ctx context.Context, Req *core_api.EmailLoginReq) (r *core_api.EmailLoginResp, err error) {
	var _args EmailLoginArgs
	_args.Req = Req
	var _result EmailLoginResult
	if err = p.c.Call(ctx, "EmailLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GithubLogin(ctx context.Context, Req *core_api.GithubLoginReq) (r *core_api.GithubLoginResp, err error) {
	var _args GithubLoginArgs
	_args.Req = Req
	var _result GithubLoginResult
	if err = p.c.Call(ctx, "GithubLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GiteeLogin(ctx context.Context, Req *core_api.GiteeLoginReq) (r *core_api.GiteeLoginResp, err error) {
	var _args GiteeLoginArgs
	_args.Req = Req
	var _result GiteeLoginResult
	if err = p.c.Call(ctx, "GiteeLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RefreshToken(ctx context.Context, Req *core_api.RefreshTokenReq) (r *core_api.RefreshTokenResp, err error) {
	var _args RefreshTokenArgs
	_args.Req = Req
	var _result RefreshTokenResult
	if err = p.c.Call(ctx, "RefreshToken", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SendEmail(ctx context.Context, Req *core_api.SendEmailReq) (r *core_api.SendEmailResp, err error) {
	var _args SendEmailArgs
	_args.Req = Req
	var _result SendEmailResult
	if err = p.c.Call(ctx, "SendEmail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCaptcha(ctx context.Context, Req *core_api.GetCaptchaReq) (r *core_api.GetCaptchaResp, err error) {
	var _args GetCaptchaArgs
	_args.Req = Req
	var _result GetCaptchaResult
	if err = p.c.Call(ctx, "GetCaptcha", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetPasswordByEmail(ctx context.Context, Req *core_api.SetPasswordByEmailReq) (r *core_api.SetPasswordByEmailResp, err error) {
	var _args SetPasswordByEmailArgs
	_args.Req = Req
	var _result SetPasswordByEmailResult
	if err = p.c.Call(ctx, "SetPasswordByEmail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetPasswordByPassword(ctx context.Context, Req *core_api.SetPasswordByPasswordReq) (r *core_api.SetPasswordByPasswordReq, err error) {
	var _args SetPasswordByPasswordArgs
	_args.Req = Req
	var _result SetPasswordByPasswordResult
	if err = p.c.Call(ctx, "SetPasswordByPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
