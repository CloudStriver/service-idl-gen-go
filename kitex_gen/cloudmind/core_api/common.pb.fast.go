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

func (x *User) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User[number], err)
}

func (x *User) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Url, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserDetail) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UserDetail[number], err)
}

func (x *UserDetail) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserDetail) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Sex, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *UserDetail) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.FullName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserDetail) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.IdCard, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserDetail) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Description, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RelationInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationInfo[number], err)
}

func (x *RelationInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FromId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RelationInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FromType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ToId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RelationInfo) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ToType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationInfo) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.RelationType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *User) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *User) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *User) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *User) fastWriteField3(buf []byte) (offset int) {
	if x.Url == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUrl())
	return offset
}

func (x *UserDetail) FastWrite(buf []byte) (offset int) {
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

func (x *UserDetail) fastWriteField1(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetName())
	return offset
}

func (x *UserDetail) fastWriteField2(buf []byte) (offset int) {
	if x.Sex == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetSex())
	return offset
}

func (x *UserDetail) fastWriteField3(buf []byte) (offset int) {
	if x.FullName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetFullName())
	return offset
}

func (x *UserDetail) fastWriteField4(buf []byte) (offset int) {
	if x.IdCard == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetIdCard())
	return offset
}

func (x *UserDetail) fastWriteField5(buf []byte) (offset int) {
	if x.Description == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetDescription())
	return offset
}

func (x *RelationInfo) FastWrite(buf []byte) (offset int) {
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

func (x *RelationInfo) fastWriteField1(buf []byte) (offset int) {
	if x.FromId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetFromId())
	return offset
}

func (x *RelationInfo) fastWriteField2(buf []byte) (offset int) {
	if x.FromType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetFromType())
	return offset
}

func (x *RelationInfo) fastWriteField3(buf []byte) (offset int) {
	if x.ToId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetToId())
	return offset
}

func (x *RelationInfo) fastWriteField4(buf []byte) (offset int) {
	if x.ToType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetToType())
	return offset
}

func (x *RelationInfo) fastWriteField5(buf []byte) (offset int) {
	if x.RelationType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, x.GetRelationType())
	return offset
}

func (x *User) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *User) sizeField1() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserId())
	return n
}

func (x *User) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *User) sizeField3() (n int) {
	if x.Url == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUrl())
	return n
}

func (x *UserDetail) Size() (n int) {
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

func (x *UserDetail) sizeField1() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetName())
	return n
}

func (x *UserDetail) sizeField2() (n int) {
	if x.Sex == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetSex())
	return n
}

func (x *UserDetail) sizeField3() (n int) {
	if x.FullName == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetFullName())
	return n
}

func (x *UserDetail) sizeField4() (n int) {
	if x.IdCard == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetIdCard())
	return n
}

func (x *UserDetail) sizeField5() (n int) {
	if x.Description == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetDescription())
	return n
}

func (x *RelationInfo) Size() (n int) {
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

func (x *RelationInfo) sizeField1() (n int) {
	if x.FromId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetFromId())
	return n
}

func (x *RelationInfo) sizeField2() (n int) {
	if x.FromType == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetFromType())
	return n
}

func (x *RelationInfo) sizeField3() (n int) {
	if x.ToId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetToId())
	return n
}

func (x *RelationInfo) sizeField4() (n int) {
	if x.ToType == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetToType())
	return n
}

func (x *RelationInfo) sizeField5() (n int) {
	if x.RelationType == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, x.GetRelationType())
	return n
}

var fieldIDToName_User = map[int32]string{
	1: "UserId",
	2: "Name",
	3: "Url",
}

var fieldIDToName_UserDetail = map[int32]string{
	1: "Name",
	2: "Sex",
	3: "FullName",
	4: "IdCard",
	5: "Description",
}

var fieldIDToName_RelationInfo = map[int32]string{
	1: "FromId",
	2: "FromType",
	3: "ToId",
	4: "ToType",
	5: "RelationType",
}

var _ = sts.File_cloudmind_sts_common_proto
var _ = content.File_cloudmind_content_common_proto
var _ = basic.File_basic_pagination_proto
var _ = http.File_http_http_proto
