// Code generated by Kitex v0.8.0. DO NOT EDIT.

package system

import (
	"context"
	core_api "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/core_api"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return systemServiceInfo
}

var systemServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "system"
	handlerType := (*core_api.System)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetNotifications":     kitex.NewMethodInfo(getNotificationsHandler, newGetNotificationsArgs, newGetNotificationsResult, false),
		"GetNotificationCount": kitex.NewMethodInfo(getNotificationCountHandler, newGetNotificationCountArgs, newGetNotificationCountResult, false),
		"UpdateSlider":         kitex.NewMethodInfo(updateSliderHandler, newUpdateSliderArgs, newUpdateSliderResult, false),
		"DeleteSlider":         kitex.NewMethodInfo(deleteSliderHandler, newDeleteSliderArgs, newDeleteSliderResult, false),
		"CreateSlider":         kitex.NewMethodInfo(createSliderHandler, newCreateSliderArgs, newCreateSliderResult, false),
		"GetSliders":           kitex.NewMethodInfo(getSlidersHandler, newGetSlidersArgs, newGetSlidersResult, false),
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

func getNotificationsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetNotificationsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.System).GetNotifications(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetNotificationsArgs:
		success, err := handler.(core_api.System).GetNotifications(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetNotificationsResult)
		realResult.Success = success
	}
	return nil
}
func newGetNotificationsArgs() interface{} {
	return &GetNotificationsArgs{}
}

func newGetNotificationsResult() interface{} {
	return &GetNotificationsResult{}
}

type GetNotificationsArgs struct {
	Req *core_api.GetNotificationsReq
}

func (p *GetNotificationsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetNotificationsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetNotificationsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetNotificationsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetNotificationsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetNotificationsArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetNotificationsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetNotificationsArgs_Req_DEFAULT *core_api.GetNotificationsReq

func (p *GetNotificationsArgs) GetReq() *core_api.GetNotificationsReq {
	if !p.IsSetReq() {
		return GetNotificationsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetNotificationsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetNotificationsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetNotificationsResult struct {
	Success *core_api.GetNotificationsResp
}

var GetNotificationsResult_Success_DEFAULT *core_api.GetNotificationsResp

func (p *GetNotificationsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetNotificationsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetNotificationsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetNotificationsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetNotificationsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetNotificationsResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetNotificationsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNotificationsResult) GetSuccess() *core_api.GetNotificationsResp {
	if !p.IsSetSuccess() {
		return GetNotificationsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetNotificationsResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetNotificationsResp)
}

func (p *GetNotificationsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetNotificationsResult) GetResult() interface{} {
	return p.Success
}

func getNotificationCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetNotificationCountReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.System).GetNotificationCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetNotificationCountArgs:
		success, err := handler.(core_api.System).GetNotificationCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetNotificationCountResult)
		realResult.Success = success
	}
	return nil
}
func newGetNotificationCountArgs() interface{} {
	return &GetNotificationCountArgs{}
}

func newGetNotificationCountResult() interface{} {
	return &GetNotificationCountResult{}
}

type GetNotificationCountArgs struct {
	Req *core_api.GetNotificationCountReq
}

func (p *GetNotificationCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetNotificationCountReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetNotificationCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetNotificationCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetNotificationCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetNotificationCountArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetNotificationCountReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetNotificationCountArgs_Req_DEFAULT *core_api.GetNotificationCountReq

func (p *GetNotificationCountArgs) GetReq() *core_api.GetNotificationCountReq {
	if !p.IsSetReq() {
		return GetNotificationCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetNotificationCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetNotificationCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetNotificationCountResult struct {
	Success *core_api.GetNotificationCountResp
}

var GetNotificationCountResult_Success_DEFAULT *core_api.GetNotificationCountResp

func (p *GetNotificationCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetNotificationCountResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetNotificationCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetNotificationCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetNotificationCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetNotificationCountResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetNotificationCountResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetNotificationCountResult) GetSuccess() *core_api.GetNotificationCountResp {
	if !p.IsSetSuccess() {
		return GetNotificationCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetNotificationCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetNotificationCountResp)
}

func (p *GetNotificationCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetNotificationCountResult) GetResult() interface{} {
	return p.Success
}

func updateSliderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.UpdateSliderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.System).UpdateSlider(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdateSliderArgs:
		success, err := handler.(core_api.System).UpdateSlider(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateSliderResult)
		realResult.Success = success
	}
	return nil
}
func newUpdateSliderArgs() interface{} {
	return &UpdateSliderArgs{}
}

func newUpdateSliderResult() interface{} {
	return &UpdateSliderResult{}
}

type UpdateSliderArgs struct {
	Req *core_api.UpdateSliderReq
}

func (p *UpdateSliderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.UpdateSliderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateSliderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateSliderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateSliderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateSliderArgs) Unmarshal(in []byte) error {
	msg := new(core_api.UpdateSliderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateSliderArgs_Req_DEFAULT *core_api.UpdateSliderReq

func (p *UpdateSliderArgs) GetReq() *core_api.UpdateSliderReq {
	if !p.IsSetReq() {
		return UpdateSliderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateSliderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateSliderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateSliderResult struct {
	Success *core_api.UpdateSliderResp
}

var UpdateSliderResult_Success_DEFAULT *core_api.UpdateSliderResp

func (p *UpdateSliderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.UpdateSliderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateSliderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateSliderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateSliderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateSliderResult) Unmarshal(in []byte) error {
	msg := new(core_api.UpdateSliderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateSliderResult) GetSuccess() *core_api.UpdateSliderResp {
	if !p.IsSetSuccess() {
		return UpdateSliderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateSliderResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.UpdateSliderResp)
}

func (p *UpdateSliderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateSliderResult) GetResult() interface{} {
	return p.Success
}

func deleteSliderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.DeleteSliderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.System).DeleteSlider(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteSliderArgs:
		success, err := handler.(core_api.System).DeleteSlider(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteSliderResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteSliderArgs() interface{} {
	return &DeleteSliderArgs{}
}

func newDeleteSliderResult() interface{} {
	return &DeleteSliderResult{}
}

type DeleteSliderArgs struct {
	Req *core_api.DeleteSliderReq
}

func (p *DeleteSliderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.DeleteSliderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteSliderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteSliderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteSliderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteSliderArgs) Unmarshal(in []byte) error {
	msg := new(core_api.DeleteSliderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteSliderArgs_Req_DEFAULT *core_api.DeleteSliderReq

func (p *DeleteSliderArgs) GetReq() *core_api.DeleteSliderReq {
	if !p.IsSetReq() {
		return DeleteSliderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteSliderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteSliderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteSliderResult struct {
	Success *core_api.DeleteSliderResp
}

var DeleteSliderResult_Success_DEFAULT *core_api.DeleteSliderResp

func (p *DeleteSliderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.DeleteSliderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteSliderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteSliderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteSliderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteSliderResult) Unmarshal(in []byte) error {
	msg := new(core_api.DeleteSliderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteSliderResult) GetSuccess() *core_api.DeleteSliderResp {
	if !p.IsSetSuccess() {
		return DeleteSliderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteSliderResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.DeleteSliderResp)
}

func (p *DeleteSliderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteSliderResult) GetResult() interface{} {
	return p.Success
}

func createSliderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.CreateSliderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.System).CreateSlider(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateSliderArgs:
		success, err := handler.(core_api.System).CreateSlider(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateSliderResult)
		realResult.Success = success
	}
	return nil
}
func newCreateSliderArgs() interface{} {
	return &CreateSliderArgs{}
}

func newCreateSliderResult() interface{} {
	return &CreateSliderResult{}
}

type CreateSliderArgs struct {
	Req *core_api.CreateSliderReq
}

func (p *CreateSliderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.CreateSliderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateSliderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateSliderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateSliderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateSliderArgs) Unmarshal(in []byte) error {
	msg := new(core_api.CreateSliderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateSliderArgs_Req_DEFAULT *core_api.CreateSliderReq

func (p *CreateSliderArgs) GetReq() *core_api.CreateSliderReq {
	if !p.IsSetReq() {
		return CreateSliderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateSliderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateSliderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateSliderResult struct {
	Success *core_api.CreateSliderResp
}

var CreateSliderResult_Success_DEFAULT *core_api.CreateSliderResp

func (p *CreateSliderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.CreateSliderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateSliderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateSliderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateSliderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateSliderResult) Unmarshal(in []byte) error {
	msg := new(core_api.CreateSliderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateSliderResult) GetSuccess() *core_api.CreateSliderResp {
	if !p.IsSetSuccess() {
		return CreateSliderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateSliderResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.CreateSliderResp)
}

func (p *CreateSliderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateSliderResult) GetResult() interface{} {
	return p.Success
}

func getSlidersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetSlidersReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.System).GetSliders(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetSlidersArgs:
		success, err := handler.(core_api.System).GetSliders(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetSlidersResult)
		realResult.Success = success
	}
	return nil
}
func newGetSlidersArgs() interface{} {
	return &GetSlidersArgs{}
}

func newGetSlidersResult() interface{} {
	return &GetSlidersResult{}
}

type GetSlidersArgs struct {
	Req *core_api.GetSlidersReq
}

func (p *GetSlidersArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetSlidersReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetSlidersArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetSlidersArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetSlidersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetSlidersArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetSlidersReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetSlidersArgs_Req_DEFAULT *core_api.GetSlidersReq

func (p *GetSlidersArgs) GetReq() *core_api.GetSlidersReq {
	if !p.IsSetReq() {
		return GetSlidersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetSlidersArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetSlidersArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetSlidersResult struct {
	Success *core_api.GetSlidersResp
}

var GetSlidersResult_Success_DEFAULT *core_api.GetSlidersResp

func (p *GetSlidersResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetSlidersResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetSlidersResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetSlidersResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetSlidersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetSlidersResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetSlidersResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetSlidersResult) GetSuccess() *core_api.GetSlidersResp {
	if !p.IsSetSuccess() {
		return GetSlidersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetSlidersResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetSlidersResp)
}

func (p *GetSlidersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetSlidersResult) GetResult() interface{} {
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

func (p *kClient) GetNotifications(ctx context.Context, Req *core_api.GetNotificationsReq) (r *core_api.GetNotificationsResp, err error) {
	var _args GetNotificationsArgs
	_args.Req = Req
	var _result GetNotificationsResult
	if err = p.c.Call(ctx, "GetNotifications", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetNotificationCount(ctx context.Context, Req *core_api.GetNotificationCountReq) (r *core_api.GetNotificationCountResp, err error) {
	var _args GetNotificationCountArgs
	_args.Req = Req
	var _result GetNotificationCountResult
	if err = p.c.Call(ctx, "GetNotificationCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateSlider(ctx context.Context, Req *core_api.UpdateSliderReq) (r *core_api.UpdateSliderResp, err error) {
	var _args UpdateSliderArgs
	_args.Req = Req
	var _result UpdateSliderResult
	if err = p.c.Call(ctx, "UpdateSlider", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteSlider(ctx context.Context, Req *core_api.DeleteSliderReq) (r *core_api.DeleteSliderResp, err error) {
	var _args DeleteSliderArgs
	_args.Req = Req
	var _result DeleteSliderResult
	if err = p.c.Call(ctx, "DeleteSlider", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateSlider(ctx context.Context, Req *core_api.CreateSliderReq) (r *core_api.CreateSliderResp, err error) {
	var _args CreateSliderArgs
	_args.Req = Req
	var _result CreateSliderResult
	if err = p.c.Call(ctx, "CreateSlider", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetSliders(ctx context.Context, Req *core_api.GetSlidersReq) (r *core_api.GetSlidersResp, err error) {
	var _args GetSlidersArgs
	_args.Req = Req
	var _result GetSlidersResult
	if err = p.c.Call(ctx, "GetSliders", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
