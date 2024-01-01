// Code generated by Kitex v0.8.0. DO NOT EDIT.

package stsservice

import (
	"context"
	sts "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/sts"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return stsServiceServiceInfo
}

var stsServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "StsService"
	handlerType := (*sts.StsService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateCaptcha": kitex.NewMethodInfo(createCaptchaHandler, newCreateCaptchaArgs, newCreateCaptchaResult, false),
		"CheckCaptcha":  kitex.NewMethodInfo(checkCaptchaHandler, newCheckCaptchaArgs, newCheckCaptchaResult, false),
		"SetPassword":   kitex.NewMethodInfo(setPasswordHandler, newSetPasswordArgs, newSetPasswordResult, false),
		"SendEmail":     kitex.NewMethodInfo(sendEmailHandler, newSendEmailArgs, newSendEmailResult, false),
		"CheckEmail":    kitex.NewMethodInfo(checkEmailHandler, newCheckEmailArgs, newCheckEmailResult, false),
		"CreateAuth":    kitex.NewMethodInfo(createAuthHandler, newCreateAuthArgs, newCreateAuthResult, false),
		"Login":         kitex.NewMethodInfo(loginHandler, newLoginArgs, newLoginResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "cloudmind.sts",
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

func createCaptchaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(sts.CreateCaptchaReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(sts.StsService).CreateCaptcha(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateCaptchaArgs:
		success, err := handler.(sts.StsService).CreateCaptcha(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateCaptchaResult)
		realResult.Success = success
	}
	return nil
}
func newCreateCaptchaArgs() interface{} {
	return &CreateCaptchaArgs{}
}

func newCreateCaptchaResult() interface{} {
	return &CreateCaptchaResult{}
}

type CreateCaptchaArgs struct {
	Req *sts.CreateCaptchaReq
}

func (p *CreateCaptchaArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(sts.CreateCaptchaReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateCaptchaArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateCaptchaArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateCaptchaArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateCaptchaArgs) Unmarshal(in []byte) error {
	msg := new(sts.CreateCaptchaReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateCaptchaArgs_Req_DEFAULT *sts.CreateCaptchaReq

func (p *CreateCaptchaArgs) GetReq() *sts.CreateCaptchaReq {
	if !p.IsSetReq() {
		return CreateCaptchaArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateCaptchaArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateCaptchaArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateCaptchaResult struct {
	Success *sts.CreateCaptchaResp
}

var CreateCaptchaResult_Success_DEFAULT *sts.CreateCaptchaResp

func (p *CreateCaptchaResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(sts.CreateCaptchaResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateCaptchaResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateCaptchaResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateCaptchaResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateCaptchaResult) Unmarshal(in []byte) error {
	msg := new(sts.CreateCaptchaResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateCaptchaResult) GetSuccess() *sts.CreateCaptchaResp {
	if !p.IsSetSuccess() {
		return CreateCaptchaResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateCaptchaResult) SetSuccess(x interface{}) {
	p.Success = x.(*sts.CreateCaptchaResp)
}

func (p *CreateCaptchaResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateCaptchaResult) GetResult() interface{} {
	return p.Success
}

func checkCaptchaHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(sts.CheckCaptchaReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(sts.StsService).CheckCaptcha(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckCaptchaArgs:
		success, err := handler.(sts.StsService).CheckCaptcha(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckCaptchaResult)
		realResult.Success = success
	}
	return nil
}
func newCheckCaptchaArgs() interface{} {
	return &CheckCaptchaArgs{}
}

func newCheckCaptchaResult() interface{} {
	return &CheckCaptchaResult{}
}

type CheckCaptchaArgs struct {
	Req *sts.CheckCaptchaReq
}

func (p *CheckCaptchaArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(sts.CheckCaptchaReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CheckCaptchaArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CheckCaptchaArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CheckCaptchaArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CheckCaptchaArgs) Unmarshal(in []byte) error {
	msg := new(sts.CheckCaptchaReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckCaptchaArgs_Req_DEFAULT *sts.CheckCaptchaReq

func (p *CheckCaptchaArgs) GetReq() *sts.CheckCaptchaReq {
	if !p.IsSetReq() {
		return CheckCaptchaArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckCaptchaArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CheckCaptchaArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CheckCaptchaResult struct {
	Success *sts.CheckCaptchaResp
}

var CheckCaptchaResult_Success_DEFAULT *sts.CheckCaptchaResp

func (p *CheckCaptchaResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(sts.CheckCaptchaResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CheckCaptchaResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CheckCaptchaResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CheckCaptchaResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CheckCaptchaResult) Unmarshal(in []byte) error {
	msg := new(sts.CheckCaptchaResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckCaptchaResult) GetSuccess() *sts.CheckCaptchaResp {
	if !p.IsSetSuccess() {
		return CheckCaptchaResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckCaptchaResult) SetSuccess(x interface{}) {
	p.Success = x.(*sts.CheckCaptchaResp)
}

func (p *CheckCaptchaResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckCaptchaResult) GetResult() interface{} {
	return p.Success
}

func setPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(sts.SetPasswordReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(sts.StsService).SetPassword(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SetPasswordArgs:
		success, err := handler.(sts.StsService).SetPassword(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SetPasswordResult)
		realResult.Success = success
	}
	return nil
}
func newSetPasswordArgs() interface{} {
	return &SetPasswordArgs{}
}

func newSetPasswordResult() interface{} {
	return &SetPasswordResult{}
}

type SetPasswordArgs struct {
	Req *sts.SetPasswordReq
}

func (p *SetPasswordArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(sts.SetPasswordReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SetPasswordArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SetPasswordArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SetPasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SetPasswordArgs) Unmarshal(in []byte) error {
	msg := new(sts.SetPasswordReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SetPasswordArgs_Req_DEFAULT *sts.SetPasswordReq

func (p *SetPasswordArgs) GetReq() *sts.SetPasswordReq {
	if !p.IsSetReq() {
		return SetPasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetPasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SetPasswordArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SetPasswordResult struct {
	Success *sts.SetPasswordResp
}

var SetPasswordResult_Success_DEFAULT *sts.SetPasswordResp

func (p *SetPasswordResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(sts.SetPasswordResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SetPasswordResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SetPasswordResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SetPasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SetPasswordResult) Unmarshal(in []byte) error {
	msg := new(sts.SetPasswordResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetPasswordResult) GetSuccess() *sts.SetPasswordResp {
	if !p.IsSetSuccess() {
		return SetPasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetPasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*sts.SetPasswordResp)
}

func (p *SetPasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetPasswordResult) GetResult() interface{} {
	return p.Success
}

func sendEmailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(sts.SendEmailReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(sts.StsService).SendEmail(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SendEmailArgs:
		success, err := handler.(sts.StsService).SendEmail(ctx, s.Req)
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
	Req *sts.SendEmailReq
}

func (p *SendEmailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(sts.SendEmailReq)
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
	msg := new(sts.SendEmailReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendEmailArgs_Req_DEFAULT *sts.SendEmailReq

func (p *SendEmailArgs) GetReq() *sts.SendEmailReq {
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
	Success *sts.SendEmailResp
}

var SendEmailResult_Success_DEFAULT *sts.SendEmailResp

func (p *SendEmailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(sts.SendEmailResp)
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
	msg := new(sts.SendEmailResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendEmailResult) GetSuccess() *sts.SendEmailResp {
	if !p.IsSetSuccess() {
		return SendEmailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendEmailResult) SetSuccess(x interface{}) {
	p.Success = x.(*sts.SendEmailResp)
}

func (p *SendEmailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendEmailResult) GetResult() interface{} {
	return p.Success
}

func checkEmailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(sts.CheckEmailReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(sts.StsService).CheckEmail(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckEmailArgs:
		success, err := handler.(sts.StsService).CheckEmail(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckEmailResult)
		realResult.Success = success
	}
	return nil
}
func newCheckEmailArgs() interface{} {
	return &CheckEmailArgs{}
}

func newCheckEmailResult() interface{} {
	return &CheckEmailResult{}
}

type CheckEmailArgs struct {
	Req *sts.CheckEmailReq
}

func (p *CheckEmailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(sts.CheckEmailReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CheckEmailArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CheckEmailArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CheckEmailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CheckEmailArgs) Unmarshal(in []byte) error {
	msg := new(sts.CheckEmailReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckEmailArgs_Req_DEFAULT *sts.CheckEmailReq

func (p *CheckEmailArgs) GetReq() *sts.CheckEmailReq {
	if !p.IsSetReq() {
		return CheckEmailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckEmailArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CheckEmailArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CheckEmailResult struct {
	Success *sts.CheckEmailResp
}

var CheckEmailResult_Success_DEFAULT *sts.CheckEmailResp

func (p *CheckEmailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(sts.CheckEmailResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CheckEmailResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CheckEmailResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CheckEmailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CheckEmailResult) Unmarshal(in []byte) error {
	msg := new(sts.CheckEmailResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckEmailResult) GetSuccess() *sts.CheckEmailResp {
	if !p.IsSetSuccess() {
		return CheckEmailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckEmailResult) SetSuccess(x interface{}) {
	p.Success = x.(*sts.CheckEmailResp)
}

func (p *CheckEmailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckEmailResult) GetResult() interface{} {
	return p.Success
}

func createAuthHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(sts.CreateAuthReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(sts.StsService).CreateAuth(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateAuthArgs:
		success, err := handler.(sts.StsService).CreateAuth(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateAuthResult)
		realResult.Success = success
	}
	return nil
}
func newCreateAuthArgs() interface{} {
	return &CreateAuthArgs{}
}

func newCreateAuthResult() interface{} {
	return &CreateAuthResult{}
}

type CreateAuthArgs struct {
	Req *sts.CreateAuthReq
}

func (p *CreateAuthArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(sts.CreateAuthReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateAuthArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateAuthArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateAuthArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateAuthArgs) Unmarshal(in []byte) error {
	msg := new(sts.CreateAuthReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateAuthArgs_Req_DEFAULT *sts.CreateAuthReq

func (p *CreateAuthArgs) GetReq() *sts.CreateAuthReq {
	if !p.IsSetReq() {
		return CreateAuthArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateAuthArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateAuthArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateAuthResult struct {
	Success *sts.CreateAuthResp
}

var CreateAuthResult_Success_DEFAULT *sts.CreateAuthResp

func (p *CreateAuthResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(sts.CreateAuthResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateAuthResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateAuthResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateAuthResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateAuthResult) Unmarshal(in []byte) error {
	msg := new(sts.CreateAuthResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateAuthResult) GetSuccess() *sts.CreateAuthResp {
	if !p.IsSetSuccess() {
		return CreateAuthResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateAuthResult) SetSuccess(x interface{}) {
	p.Success = x.(*sts.CreateAuthResp)
}

func (p *CreateAuthResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateAuthResult) GetResult() interface{} {
	return p.Success
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(sts.LoginReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(sts.StsService).Login(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *LoginArgs:
		success, err := handler.(sts.StsService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
	}
	return nil
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *sts.LoginReq
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(sts.LoginReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(sts.LoginReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *sts.LoginReq

func (p *LoginArgs) GetReq() *sts.LoginReq {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type LoginResult struct {
	Success *sts.LoginResp
}

var LoginResult_Success_DEFAULT *sts.LoginResp

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(sts.LoginResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(sts.LoginResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *sts.LoginResp {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*sts.LoginResp)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginResult) GetResult() interface{} {
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

func (p *kClient) CreateCaptcha(ctx context.Context, Req *sts.CreateCaptchaReq) (r *sts.CreateCaptchaResp, err error) {
	var _args CreateCaptchaArgs
	_args.Req = Req
	var _result CreateCaptchaResult
	if err = p.c.Call(ctx, "CreateCaptcha", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckCaptcha(ctx context.Context, Req *sts.CheckCaptchaReq) (r *sts.CheckCaptchaResp, err error) {
	var _args CheckCaptchaArgs
	_args.Req = Req
	var _result CheckCaptchaResult
	if err = p.c.Call(ctx, "CheckCaptcha", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetPassword(ctx context.Context, Req *sts.SetPasswordReq) (r *sts.SetPasswordResp, err error) {
	var _args SetPasswordArgs
	_args.Req = Req
	var _result SetPasswordResult
	if err = p.c.Call(ctx, "SetPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SendEmail(ctx context.Context, Req *sts.SendEmailReq) (r *sts.SendEmailResp, err error) {
	var _args SendEmailArgs
	_args.Req = Req
	var _result SendEmailResult
	if err = p.c.Call(ctx, "SendEmail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckEmail(ctx context.Context, Req *sts.CheckEmailReq) (r *sts.CheckEmailResp, err error) {
	var _args CheckEmailArgs
	_args.Req = Req
	var _result CheckEmailResult
	if err = p.c.Call(ctx, "CheckEmail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateAuth(ctx context.Context, Req *sts.CreateAuthReq) (r *sts.CreateAuthResp, err error) {
	var _args CreateAuthArgs
	_args.Req = Req
	var _result CreateAuthResult
	if err = p.c.Call(ctx, "CreateAuth", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *sts.LoginReq) (r *sts.LoginResp, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}