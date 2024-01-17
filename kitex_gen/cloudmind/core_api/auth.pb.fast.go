// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package core_api

import (
	fmt "fmt"
	basic "github.com/CloudStriver/service-idl-gen-go/kitex_gen/basic"
	content "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/content"
	sts "github.com/CloudStriver/service-idl-gen-go/kitex_gen/cloudmind/sts"
	http "github.com/CloudStriver/service-idl-gen-go/kitex_gen/http"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *RegisterReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RegisterReq[number], err)
}

func (x *RegisterReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Sex, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RegisterReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RegisterResp[number], err)
}

func (x *RegisterResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ShortToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.LongToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RegisterResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EmailLoginReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_EmailLoginReq[number], err)
}

func (x *EmailLoginReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EmailLoginReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EmailLoginResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_EmailLoginResp[number], err)
}

func (x *EmailLoginResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ShortToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EmailLoginResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.LongToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *EmailLoginResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GithubLoginReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GithubLoginReq[number], err)
}

func (x *GithubLoginReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GithubLoginResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GithubLoginResp[number], err)
}

func (x *GithubLoginResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ShortToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GithubLoginResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.LongToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GithubLoginResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GiteeLoginReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GiteeLoginReq[number], err)
}

func (x *GiteeLoginReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GiteeLoginResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GiteeLoginResp[number], err)
}

func (x *GiteeLoginResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ShortToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GiteeLoginResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.LongToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GiteeLoginResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RefreshTokenReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RefreshTokenReq[number], err)
}

func (x *RefreshTokenReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.LongToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RefreshTokenResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RefreshTokenResp[number], err)
}

func (x *RefreshTokenResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ShortToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RefreshTokenResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.LongToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendEmailReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendEmailReq[number], err)
}

func (x *SendEmailReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendEmailReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Subject, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendEmailResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *GetCaptchaReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *GetCaptchaResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetCaptchaResp[number], err)
}

func (x *GetCaptchaResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OriginalImageBase64, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetCaptchaResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.JigsawImageBase64, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetCaptchaResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Key, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SetPasswordByEmailReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SetPasswordByEmailReq[number], err)
}

func (x *SetPasswordByEmailReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SetPasswordByEmailReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Code, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SetPasswordByEmailReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SetPasswordByEmailResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *SetPasswordByPasswordReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SetPasswordByPasswordReq[number], err)
}

func (x *SetPasswordByPasswordReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OldPassword, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SetPasswordByPasswordReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SetPasswordByPasswordResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *RegisterReq) FastWrite(buf []byte) (offset int) {
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

func (x *RegisterReq) fastWriteField1(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetName())
	return offset
}

func (x *RegisterReq) fastWriteField2(buf []byte) (offset int) {
	if x.Sex == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetSex())
	return offset
}

func (x *RegisterReq) fastWriteField3(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetEmail())
	return offset
}

func (x *RegisterReq) fastWriteField4(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetPassword())
	return offset
}

func (x *RegisterReq) fastWriteField5(buf []byte) (offset int) {
	if x.Code == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetCode())
	return offset
}

func (x *RegisterResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *RegisterResp) fastWriteField1(buf []byte) (offset int) {
	if x.ShortToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetShortToken())
	return offset
}

func (x *RegisterResp) fastWriteField2(buf []byte) (offset int) {
	if x.LongToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetLongToken())
	return offset
}

func (x *RegisterResp) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUserId())
	return offset
}

func (x *EmailLoginReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *EmailLoginReq) fastWriteField1(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetEmail())
	return offset
}

func (x *EmailLoginReq) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *EmailLoginResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *EmailLoginResp) fastWriteField1(buf []byte) (offset int) {
	if x.ShortToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetShortToken())
	return offset
}

func (x *EmailLoginResp) fastWriteField2(buf []byte) (offset int) {
	if x.LongToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetLongToken())
	return offset
}

func (x *EmailLoginResp) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUserId())
	return offset
}

func (x *GithubLoginReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GithubLoginReq) fastWriteField1(buf []byte) (offset int) {
	if x.Code == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *GithubLoginResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GithubLoginResp) fastWriteField1(buf []byte) (offset int) {
	if x.ShortToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetShortToken())
	return offset
}

func (x *GithubLoginResp) fastWriteField2(buf []byte) (offset int) {
	if x.LongToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetLongToken())
	return offset
}

func (x *GithubLoginResp) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUserId())
	return offset
}

func (x *GiteeLoginReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GiteeLoginReq) fastWriteField1(buf []byte) (offset int) {
	if x.Code == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetCode())
	return offset
}

func (x *GiteeLoginResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GiteeLoginResp) fastWriteField1(buf []byte) (offset int) {
	if x.ShortToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetShortToken())
	return offset
}

func (x *GiteeLoginResp) fastWriteField2(buf []byte) (offset int) {
	if x.LongToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetLongToken())
	return offset
}

func (x *GiteeLoginResp) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUserId())
	return offset
}

func (x *RefreshTokenReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *RefreshTokenReq) fastWriteField1(buf []byte) (offset int) {
	if x.LongToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetLongToken())
	return offset
}

func (x *RefreshTokenResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RefreshTokenResp) fastWriteField1(buf []byte) (offset int) {
	if x.ShortToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetShortToken())
	return offset
}

func (x *RefreshTokenResp) fastWriteField2(buf []byte) (offset int) {
	if x.LongToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetLongToken())
	return offset
}

func (x *SendEmailReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SendEmailReq) fastWriteField1(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetEmail())
	return offset
}

func (x *SendEmailReq) fastWriteField2(buf []byte) (offset int) {
	if x.Subject == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSubject())
	return offset
}

func (x *SendEmailResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *GetCaptchaReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *GetCaptchaResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetCaptchaResp) fastWriteField1(buf []byte) (offset int) {
	if x.OriginalImageBase64 == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetOriginalImageBase64())
	return offset
}

func (x *GetCaptchaResp) fastWriteField2(buf []byte) (offset int) {
	if x.JigsawImageBase64 == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetJigsawImageBase64())
	return offset
}

func (x *GetCaptchaResp) fastWriteField3(buf []byte) (offset int) {
	if x.Key == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetKey())
	return offset
}

func (x *SetPasswordByEmailReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *SetPasswordByEmailReq) fastWriteField1(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetEmail())
	return offset
}

func (x *SetPasswordByEmailReq) fastWriteField2(buf []byte) (offset int) {
	if x.Code == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetCode())
	return offset
}

func (x *SetPasswordByEmailReq) fastWriteField3(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPassword())
	return offset
}

func (x *SetPasswordByEmailResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *SetPasswordByPasswordReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SetPasswordByPasswordReq) fastWriteField1(buf []byte) (offset int) {
	if x.OldPassword == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetOldPassword())
	return offset
}

func (x *SetPasswordByPasswordReq) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *SetPasswordByPasswordResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *RegisterReq) Size() (n int) {
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

func (x *RegisterReq) sizeField1() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetName())
	return n
}

func (x *RegisterReq) sizeField2() (n int) {
	if x.Sex == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetSex())
	return n
}

func (x *RegisterReq) sizeField3() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetEmail())
	return n
}

func (x *RegisterReq) sizeField4() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetPassword())
	return n
}

func (x *RegisterReq) sizeField5() (n int) {
	if x.Code == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetCode())
	return n
}

func (x *RegisterResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *RegisterResp) sizeField1() (n int) {
	if x.ShortToken == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetShortToken())
	return n
}

func (x *RegisterResp) sizeField2() (n int) {
	if x.LongToken == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetLongToken())
	return n
}

func (x *RegisterResp) sizeField3() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUserId())
	return n
}

func (x *EmailLoginReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *EmailLoginReq) sizeField1() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetEmail())
	return n
}

func (x *EmailLoginReq) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *EmailLoginResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *EmailLoginResp) sizeField1() (n int) {
	if x.ShortToken == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetShortToken())
	return n
}

func (x *EmailLoginResp) sizeField2() (n int) {
	if x.LongToken == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetLongToken())
	return n
}

func (x *EmailLoginResp) sizeField3() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUserId())
	return n
}

func (x *GithubLoginReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GithubLoginReq) sizeField1() (n int) {
	if x.Code == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetCode())
	return n
}

func (x *GithubLoginResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GithubLoginResp) sizeField1() (n int) {
	if x.ShortToken == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetShortToken())
	return n
}

func (x *GithubLoginResp) sizeField2() (n int) {
	if x.LongToken == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetLongToken())
	return n
}

func (x *GithubLoginResp) sizeField3() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUserId())
	return n
}

func (x *GiteeLoginReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GiteeLoginReq) sizeField1() (n int) {
	if x.Code == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetCode())
	return n
}

func (x *GiteeLoginResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GiteeLoginResp) sizeField1() (n int) {
	if x.ShortToken == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetShortToken())
	return n
}

func (x *GiteeLoginResp) sizeField2() (n int) {
	if x.LongToken == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetLongToken())
	return n
}

func (x *GiteeLoginResp) sizeField3() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUserId())
	return n
}

func (x *RefreshTokenReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *RefreshTokenReq) sizeField1() (n int) {
	if x.LongToken == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetLongToken())
	return n
}

func (x *RefreshTokenResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RefreshTokenResp) sizeField1() (n int) {
	if x.ShortToken == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetShortToken())
	return n
}

func (x *RefreshTokenResp) sizeField2() (n int) {
	if x.LongToken == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetLongToken())
	return n
}

func (x *SendEmailReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SendEmailReq) sizeField1() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetEmail())
	return n
}

func (x *SendEmailReq) sizeField2() (n int) {
	if x.Subject == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSubject())
	return n
}

func (x *SendEmailResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *GetCaptchaReq) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *GetCaptchaResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetCaptchaResp) sizeField1() (n int) {
	if x.OriginalImageBase64 == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetOriginalImageBase64())
	return n
}

func (x *GetCaptchaResp) sizeField2() (n int) {
	if x.JigsawImageBase64 == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetJigsawImageBase64())
	return n
}

func (x *GetCaptchaResp) sizeField3() (n int) {
	if x.Key == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetKey())
	return n
}

func (x *SetPasswordByEmailReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *SetPasswordByEmailReq) sizeField1() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetEmail())
	return n
}

func (x *SetPasswordByEmailReq) sizeField2() (n int) {
	if x.Code == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetCode())
	return n
}

func (x *SetPasswordByEmailReq) sizeField3() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPassword())
	return n
}

func (x *SetPasswordByEmailResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *SetPasswordByPasswordReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SetPasswordByPasswordReq) sizeField1() (n int) {
	if x.OldPassword == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetOldPassword())
	return n
}

func (x *SetPasswordByPasswordReq) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *SetPasswordByPasswordResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_RegisterReq = map[int32]string{
	1: "Name",
	2: "Sex",
	3: "Email",
	4: "Password",
	5: "Code",
}

var fieldIDToName_RegisterResp = map[int32]string{
	1: "ShortToken",
	2: "LongToken",
	3: "UserId",
}

var fieldIDToName_EmailLoginReq = map[int32]string{
	1: "Email",
	2: "Password",
}

var fieldIDToName_EmailLoginResp = map[int32]string{
	1: "ShortToken",
	2: "LongToken",
	3: "UserId",
}

var fieldIDToName_GithubLoginReq = map[int32]string{
	1: "Code",
}

var fieldIDToName_GithubLoginResp = map[int32]string{
	1: "ShortToken",
	2: "LongToken",
	3: "UserId",
}

var fieldIDToName_GiteeLoginReq = map[int32]string{
	1: "Code",
}

var fieldIDToName_GiteeLoginResp = map[int32]string{
	1: "ShortToken",
	2: "LongToken",
	3: "UserId",
}

var fieldIDToName_RefreshTokenReq = map[int32]string{
	1: "LongToken",
}

var fieldIDToName_RefreshTokenResp = map[int32]string{
	1: "ShortToken",
	2: "LongToken",
}

var fieldIDToName_SendEmailReq = map[int32]string{
	1: "Email",
	2: "Subject",
}

var fieldIDToName_SendEmailResp = map[int32]string{}

var fieldIDToName_GetCaptchaReq = map[int32]string{}

var fieldIDToName_GetCaptchaResp = map[int32]string{
	1: "OriginalImageBase64",
	2: "JigsawImageBase64",
	3: "Key",
}

var fieldIDToName_SetPasswordByEmailReq = map[int32]string{
	1: "Email",
	2: "Code",
	3: "Password",
}

var fieldIDToName_SetPasswordByEmailResp = map[int32]string{}

var fieldIDToName_SetPasswordByPasswordReq = map[int32]string{
	1: "OldPassword",
	2: "Password",
}

var fieldIDToName_SetPasswordByPasswordResp = map[int32]string{}

var _ = sts.File_cloudmind_sts_common_proto
var _ = content.File_cloudmind_content_common_proto
var _ = basic.File_basic_pagination_proto
var _ = http.File_http_http_proto
