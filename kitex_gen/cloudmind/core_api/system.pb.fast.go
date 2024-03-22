// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package core_api

import (
	fmt "fmt"
	basic "github.com/CloudStriver/service-idl-gen-go/kitex_gen/basic"
	http "github.com/CloudStriver/service-idl-gen-go/kitex_gen/http"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *GetNotificationsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetNotificationsReq[number], err)
}

func (x *GetNotificationsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.OnlyType = &tmp
	return offset, err
}

func (x *GetNotificationsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.Limit = &tmp
	return offset, err
}

func (x *GetNotificationsReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.LastToken = &tmp
	return offset, err
}

func (x *GetNotificationsReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadBool(buf, _type)
	x.Backward = &tmp
	return offset, err
}

func (x *GetNotificationsReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.Offset = &tmp
	return offset, err
}

func (x *GetNotificationsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetNotificationsResp[number], err)
}

func (x *GetNotificationsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Notification
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Notifications = append(x.Notifications, &v)
	return offset, nil
}

func (x *GetNotificationsResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetNotificationCountReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetNotificationCountReq[number], err)
}

func (x *GetNotificationCountReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.OnlyType = &tmp
	return offset, err
}

func (x *GetNotificationCountResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetNotificationCountResp[number], err)
}

func (x *GetNotificationCountResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteNotificationsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteNotificationsReq[number], err)
}

func (x *DeleteNotificationsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.NotificationIds = append(x.NotificationIds, v)
	return offset, err
}

func (x *DeleteNotificationsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.OnlyType = &tmp
	return offset, err
}

func (x *DeleteNotificationsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *CreateSliderReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateSliderReq[number], err)
}

func (x *CreateSliderReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ImageUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateSliderReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.LinkUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateSliderReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Type, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateSliderReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.IsPublic, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateSliderResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *DeleteSliderReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteSliderReq[number], err)
}

func (x *DeleteSliderReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SliderId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteSliderResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *UpdateSliderReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateSliderReq[number], err)
}

func (x *UpdateSliderReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SliderId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateSliderReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ImageUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateSliderReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.LinkUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateSliderReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Type, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UpdateSliderReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.IsPublic, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UpdateSliderResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *GetSlidersReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *GetSlidersResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetSlidersResp[number], err)
}

func (x *GetSlidersResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Slider
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Sliders = append(x.Sliders, &v)
	return offset, nil
}

func (x *GetSlidersResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetNotificationsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *GetNotificationsReq) fastWriteField1(buf []byte) (offset int) {
	if x.OnlyType == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetOnlyType())
	return offset
}

func (x *GetNotificationsReq) fastWriteField2(buf []byte) (offset int) {
	if x.Limit == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetLimit())
	return offset
}

func (x *GetNotificationsReq) fastWriteField3(buf []byte) (offset int) {
	if x.LastToken == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetLastToken())
	return offset
}

func (x *GetNotificationsReq) fastWriteField4(buf []byte) (offset int) {
	if x.Backward == nil {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 4, x.GetBackward())
	return offset
}

func (x *GetNotificationsReq) fastWriteField5(buf []byte) (offset int) {
	if x.Offset == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetOffset())
	return offset
}

func (x *GetNotificationsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetNotificationsResp) fastWriteField1(buf []byte) (offset int) {
	if x.Notifications == nil {
		return offset
	}
	for i := range x.GetNotifications() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetNotifications()[i])
	}
	return offset
}

func (x *GetNotificationsResp) fastWriteField2(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetToken())
	return offset
}

func (x *GetNotificationCountReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetNotificationCountReq) fastWriteField1(buf []byte) (offset int) {
	if x.OnlyType == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetOnlyType())
	return offset
}

func (x *GetNotificationCountResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetNotificationCountResp) fastWriteField1(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetTotal())
	return offset
}

func (x *DeleteNotificationsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DeleteNotificationsReq) fastWriteField1(buf []byte) (offset int) {
	if len(x.NotificationIds) == 0 {
		return offset
	}
	for i := range x.GetNotificationIds() {
		offset += fastpb.WriteString(buf[offset:], 1, x.GetNotificationIds()[i])
	}
	return offset
}

func (x *DeleteNotificationsReq) fastWriteField2(buf []byte) (offset int) {
	if x.OnlyType == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetOnlyType())
	return offset
}

func (x *DeleteNotificationsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *CreateSliderReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *CreateSliderReq) fastWriteField1(buf []byte) (offset int) {
	if x.ImageUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetImageUrl())
	return offset
}

func (x *CreateSliderReq) fastWriteField2(buf []byte) (offset int) {
	if x.LinkUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetLinkUrl())
	return offset
}

func (x *CreateSliderReq) fastWriteField3(buf []byte) (offset int) {
	if x.Type == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetType())
	return offset
}

func (x *CreateSliderReq) fastWriteField4(buf []byte) (offset int) {
	if x.IsPublic == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetIsPublic())
	return offset
}

func (x *CreateSliderResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *DeleteSliderReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteSliderReq) fastWriteField1(buf []byte) (offset int) {
	if x.SliderId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetSliderId())
	return offset
}

func (x *DeleteSliderResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *UpdateSliderReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *UpdateSliderReq) fastWriteField1(buf []byte) (offset int) {
	if x.SliderId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetSliderId())
	return offset
}

func (x *UpdateSliderReq) fastWriteField2(buf []byte) (offset int) {
	if x.ImageUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetImageUrl())
	return offset
}

func (x *UpdateSliderReq) fastWriteField3(buf []byte) (offset int) {
	if x.LinkUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetLinkUrl())
	return offset
}

func (x *UpdateSliderReq) fastWriteField4(buf []byte) (offset int) {
	if x.Type == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetType())
	return offset
}

func (x *UpdateSliderReq) fastWriteField5(buf []byte) (offset int) {
	if x.IsPublic == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetIsPublic())
	return offset
}

func (x *UpdateSliderResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *GetSlidersReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *GetSlidersResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetSlidersResp) fastWriteField1(buf []byte) (offset int) {
	if x.Sliders == nil {
		return offset
	}
	for i := range x.GetSliders() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetSliders()[i])
	}
	return offset
}

func (x *GetSlidersResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *GetNotificationsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *GetNotificationsReq) sizeField1() (n int) {
	if x.OnlyType == nil {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetOnlyType())
	return n
}

func (x *GetNotificationsReq) sizeField2() (n int) {
	if x.Limit == nil {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetLimit())
	return n
}

func (x *GetNotificationsReq) sizeField3() (n int) {
	if x.LastToken == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetLastToken())
	return n
}

func (x *GetNotificationsReq) sizeField4() (n int) {
	if x.Backward == nil {
		return n
	}
	n += fastpb.SizeBool(4, x.GetBackward())
	return n
}

func (x *GetNotificationsReq) sizeField5() (n int) {
	if x.Offset == nil {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetOffset())
	return n
}

func (x *GetNotificationsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetNotificationsResp) sizeField1() (n int) {
	if x.Notifications == nil {
		return n
	}
	for i := range x.GetNotifications() {
		n += fastpb.SizeMessage(1, x.GetNotifications()[i])
	}
	return n
}

func (x *GetNotificationsResp) sizeField2() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetToken())
	return n
}

func (x *GetNotificationCountReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetNotificationCountReq) sizeField1() (n int) {
	if x.OnlyType == nil {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetOnlyType())
	return n
}

func (x *GetNotificationCountResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetNotificationCountResp) sizeField1() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetTotal())
	return n
}

func (x *DeleteNotificationsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DeleteNotificationsReq) sizeField1() (n int) {
	if len(x.NotificationIds) == 0 {
		return n
	}
	for i := range x.GetNotificationIds() {
		n += fastpb.SizeString(1, x.GetNotificationIds()[i])
	}
	return n
}

func (x *DeleteNotificationsReq) sizeField2() (n int) {
	if x.OnlyType == nil {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetOnlyType())
	return n
}

func (x *DeleteNotificationsResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *CreateSliderReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *CreateSliderReq) sizeField1() (n int) {
	if x.ImageUrl == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetImageUrl())
	return n
}

func (x *CreateSliderReq) sizeField2() (n int) {
	if x.LinkUrl == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetLinkUrl())
	return n
}

func (x *CreateSliderReq) sizeField3() (n int) {
	if x.Type == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetType())
	return n
}

func (x *CreateSliderReq) sizeField4() (n int) {
	if x.IsPublic == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetIsPublic())
	return n
}

func (x *CreateSliderResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *DeleteSliderReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteSliderReq) sizeField1() (n int) {
	if x.SliderId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetSliderId())
	return n
}

func (x *DeleteSliderResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *UpdateSliderReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *UpdateSliderReq) sizeField1() (n int) {
	if x.SliderId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetSliderId())
	return n
}

func (x *UpdateSliderReq) sizeField2() (n int) {
	if x.ImageUrl == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetImageUrl())
	return n
}

func (x *UpdateSliderReq) sizeField3() (n int) {
	if x.LinkUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetLinkUrl())
	return n
}

func (x *UpdateSliderReq) sizeField4() (n int) {
	if x.Type == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetType())
	return n
}

func (x *UpdateSliderReq) sizeField5() (n int) {
	if x.IsPublic == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetIsPublic())
	return n
}

func (x *UpdateSliderResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *GetSlidersReq) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *GetSlidersResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetSlidersResp) sizeField1() (n int) {
	if x.Sliders == nil {
		return n
	}
	for i := range x.GetSliders() {
		n += fastpb.SizeMessage(1, x.GetSliders()[i])
	}
	return n
}

func (x *GetSlidersResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

var fieldIDToName_GetNotificationsReq = map[int32]string{
	1: "OnlyType",
	2: "Limit",
	3: "LastToken",
	4: "Backward",
	5: "Offset",
}

var fieldIDToName_GetNotificationsResp = map[int32]string{
	1: "Notifications",
	2: "Token",
}

var fieldIDToName_GetNotificationCountReq = map[int32]string{
	1: "OnlyType",
}

var fieldIDToName_GetNotificationCountResp = map[int32]string{
	1: "Total",
}

var fieldIDToName_DeleteNotificationsReq = map[int32]string{
	1: "NotificationIds",
	2: "OnlyType",
}

var fieldIDToName_DeleteNotificationsResp = map[int32]string{}

var fieldIDToName_CreateSliderReq = map[int32]string{
	1: "ImageUrl",
	2: "LinkUrl",
	3: "Type",
	4: "IsPublic",
}

var fieldIDToName_CreateSliderResp = map[int32]string{}

var fieldIDToName_DeleteSliderReq = map[int32]string{
	1: "SliderId",
}

var fieldIDToName_DeleteSliderResp = map[int32]string{}

var fieldIDToName_UpdateSliderReq = map[int32]string{
	1: "SliderId",
	2: "ImageUrl",
	3: "LinkUrl",
	4: "Type",
	5: "IsPublic",
}

var fieldIDToName_UpdateSliderResp = map[int32]string{}

var fieldIDToName_GetSlidersReq = map[int32]string{}

var fieldIDToName_GetSlidersResp = map[int32]string{
	1: "Sliders",
	2: "Total",
}

var _ = basic.File_basic_pagination_proto
var _ = http.File_http_http_proto
