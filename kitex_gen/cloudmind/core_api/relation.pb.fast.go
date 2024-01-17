// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package core_api

import (
	fmt "fmt"
	basic "github.com/CloudStriver/service-idl-gen-go/kitex_gen/basic"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *CreateRelationReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateRelationReq[number], err)
}

func (x *CreateRelationReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v RelationInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Relation = &v
	return offset, nil
}

func (x *CreateRelationResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *GetFromRelationsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFromRelationsReq[number], err)
}

func (x *GetFromRelationsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FromId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetFromRelationsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FromType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetFromRelationsReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.ToType = TargetType(v)
	return offset, nil
}

func (x *GetFromRelationsReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.RelationType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetFromRelationsReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOptions = &v
	return offset, nil
}

func (x *GetFromRelationsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFromRelationsResp[number], err)
}

func (x *GetFromRelationsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Users = append(x.Users, &v)
	return offset, nil
}

func (x *GetToRelationsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetToRelationsReq[number], err)
}

func (x *GetToRelationsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ToType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetToRelationsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetToRelationsReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.FromType = TargetType(v)
	return offset, nil
}

func (x *GetToRelationsReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.RelationType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetToRelationsReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOptions = &v
	return offset, nil
}

func (x *GetToRelationsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetToRelationsResp[number], err)
}

func (x *GetToRelationsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Users = append(x.Users, &v)
	return offset, nil
}

func (x *GetRelationReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetRelationReq[number], err)
}

func (x *GetRelationReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v RelationInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.RelationInfo = &v
	return offset, nil
}

func (x *GetRelationResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetRelationResp[number], err)
}

func (x *GetRelationResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Ok, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *DeleteRelationReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteRelationReq[number], err)
}

func (x *DeleteRelationReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v RelationInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.RelationInfo = &v
	return offset, nil
}

func (x *DeleteRelationResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *CreateRelationReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CreateRelationReq) fastWriteField1(buf []byte) (offset int) {
	if x.Relation == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetRelation())
	return offset
}

func (x *CreateRelationResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *GetFromRelationsReq) FastWrite(buf []byte) (offset int) {
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

func (x *GetFromRelationsReq) fastWriteField1(buf []byte) (offset int) {
	if x.FromId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetFromId())
	return offset
}

func (x *GetFromRelationsReq) fastWriteField2(buf []byte) (offset int) {
	if x.FromType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetFromType())
	return offset
}

func (x *GetFromRelationsReq) fastWriteField3(buf []byte) (offset int) {
	if x.ToType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, int32(x.GetToType()))
	return offset
}

func (x *GetFromRelationsReq) fastWriteField4(buf []byte) (offset int) {
	if x.RelationType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.GetRelationType())
	return offset
}

func (x *GetFromRelationsReq) fastWriteField5(buf []byte) (offset int) {
	if x.PaginationOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 5, x.GetPaginationOptions())
	return offset
}

func (x *GetFromRelationsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetFromRelationsResp) fastWriteField1(buf []byte) (offset int) {
	if x.Users == nil {
		return offset
	}
	for i := range x.GetUsers() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUsers()[i])
	}
	return offset
}

func (x *GetToRelationsReq) FastWrite(buf []byte) (offset int) {
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

func (x *GetToRelationsReq) fastWriteField1(buf []byte) (offset int) {
	if x.ToType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetToType())
	return offset
}

func (x *GetToRelationsReq) fastWriteField2(buf []byte) (offset int) {
	if x.ToId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetToId())
	return offset
}

func (x *GetToRelationsReq) fastWriteField3(buf []byte) (offset int) {
	if x.FromType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, int32(x.GetFromType()))
	return offset
}

func (x *GetToRelationsReq) fastWriteField4(buf []byte) (offset int) {
	if x.RelationType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.GetRelationType())
	return offset
}

func (x *GetToRelationsReq) fastWriteField5(buf []byte) (offset int) {
	if x.PaginationOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 5, x.GetPaginationOptions())
	return offset
}

func (x *GetToRelationsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetToRelationsResp) fastWriteField1(buf []byte) (offset int) {
	if x.Users == nil {
		return offset
	}
	for i := range x.GetUsers() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUsers()[i])
	}
	return offset
}

func (x *GetRelationReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetRelationReq) fastWriteField1(buf []byte) (offset int) {
	if x.RelationInfo == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetRelationInfo())
	return offset
}

func (x *GetRelationResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetRelationResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Ok {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetOk())
	return offset
}

func (x *DeleteRelationReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteRelationReq) fastWriteField1(buf []byte) (offset int) {
	if x.RelationInfo == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetRelationInfo())
	return offset
}

func (x *DeleteRelationResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *CreateRelationReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CreateRelationReq) sizeField1() (n int) {
	if x.Relation == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetRelation())
	return n
}

func (x *CreateRelationResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *GetFromRelationsReq) Size() (n int) {
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

func (x *GetFromRelationsReq) sizeField1() (n int) {
	if x.FromId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetFromId())
	return n
}

func (x *GetFromRelationsReq) sizeField2() (n int) {
	if x.FromType == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetFromType())
	return n
}

func (x *GetFromRelationsReq) sizeField3() (n int) {
	if x.ToType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, int32(x.GetToType()))
	return n
}

func (x *GetFromRelationsReq) sizeField4() (n int) {
	if x.RelationType == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.GetRelationType())
	return n
}

func (x *GetFromRelationsReq) sizeField5() (n int) {
	if x.PaginationOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(5, x.GetPaginationOptions())
	return n
}

func (x *GetFromRelationsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetFromRelationsResp) sizeField1() (n int) {
	if x.Users == nil {
		return n
	}
	for i := range x.GetUsers() {
		n += fastpb.SizeMessage(1, x.GetUsers()[i])
	}
	return n
}

func (x *GetToRelationsReq) Size() (n int) {
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

func (x *GetToRelationsReq) sizeField1() (n int) {
	if x.ToType == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetToType())
	return n
}

func (x *GetToRelationsReq) sizeField2() (n int) {
	if x.ToId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetToId())
	return n
}

func (x *GetToRelationsReq) sizeField3() (n int) {
	if x.FromType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, int32(x.GetFromType()))
	return n
}

func (x *GetToRelationsReq) sizeField4() (n int) {
	if x.RelationType == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.GetRelationType())
	return n
}

func (x *GetToRelationsReq) sizeField5() (n int) {
	if x.PaginationOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(5, x.GetPaginationOptions())
	return n
}

func (x *GetToRelationsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetToRelationsResp) sizeField1() (n int) {
	if x.Users == nil {
		return n
	}
	for i := range x.GetUsers() {
		n += fastpb.SizeMessage(1, x.GetUsers()[i])
	}
	return n
}

func (x *GetRelationReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetRelationReq) sizeField1() (n int) {
	if x.RelationInfo == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetRelationInfo())
	return n
}

func (x *GetRelationResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetRelationResp) sizeField1() (n int) {
	if !x.Ok {
		return n
	}
	n += fastpb.SizeBool(1, x.GetOk())
	return n
}

func (x *DeleteRelationReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteRelationReq) sizeField1() (n int) {
	if x.RelationInfo == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetRelationInfo())
	return n
}

func (x *DeleteRelationResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_CreateRelationReq = map[int32]string{
	1: "Relation",
}

var fieldIDToName_CreateRelationResp = map[int32]string{}

var fieldIDToName_GetFromRelationsReq = map[int32]string{
	1: "FromId",
	2: "FromType",
	3: "ToType",
	4: "RelationType",
	5: "PaginationOptions",
}

var fieldIDToName_GetFromRelationsResp = map[int32]string{
	1: "Users",
}

var fieldIDToName_GetToRelationsReq = map[int32]string{
	1: "ToType",
	2: "ToId",
	3: "FromType",
	4: "RelationType",
	5: "PaginationOptions",
}

var fieldIDToName_GetToRelationsResp = map[int32]string{
	1: "Users",
}

var fieldIDToName_GetRelationReq = map[int32]string{
	1: "RelationInfo",
}

var fieldIDToName_GetRelationResp = map[int32]string{
	1: "Ok",
}

var fieldIDToName_DeleteRelationReq = map[int32]string{
	1: "RelationInfo",
}

var fieldIDToName_DeleteRelationResp = map[int32]string{}

var _ = basic.File_basic_pagination_proto
