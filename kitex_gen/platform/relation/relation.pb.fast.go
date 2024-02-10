// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package relation

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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateRelationReq[number], err)
}

func (x *CreateRelationReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FromType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateRelationReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FromId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateRelationReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ToType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateRelationReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ToId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateRelationReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.RelationType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateRelationResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateRelationResp[number], err)
}

func (x *CreateRelationResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Ok, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *GetRelationReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetRelationReq[number], err)
}

func (x *GetRelationReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FromType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetRelationReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FromId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetRelationReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ToType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetRelationReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ToId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetRelationReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.RelationType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteRelationReq[number], err)
}

func (x *DeleteRelationReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FromType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteRelationReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FromId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteRelationReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ToType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteRelationReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ToId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteRelationReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.RelationType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
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

func (x *FromFilterOptions) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FromFilterOptions[number], err)
}

func (x *FromFilterOptions) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FromType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FromFilterOptions) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FromId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *FromFilterOptions) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ToType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ToFilterOptions) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ToFilterOptions[number], err)
}

func (x *ToFilterOptions) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ToType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ToFilterOptions) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ToFilterOptions) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.FromType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetRelationsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetRelationsReq[number], err)
}

func (x *GetRelationsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var ov GetRelationsReq_FromFilterOptions
	x.RelationFilterOptions = &ov
	var v FromFilterOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	ov.FromFilterOptions = &v
	return offset, nil
}

func (x *GetRelationsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var ov GetRelationsReq_ToFilterOptions
	x.RelationFilterOptions = &ov
	var v ToFilterOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	ov.ToFilterOptions = &v
	return offset, nil
}

func (x *GetRelationsReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.RelationType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetRelationsReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOptions = &v
	return offset, nil
}

func (x *GetRelationsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetRelationsResp[number], err)
}

func (x *GetRelationsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Relation
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Relations = append(x.Relations, &v)
	return offset, nil
}

func (x *GetRelationsResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetRelationCountReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetRelationCountReq[number], err)
}

func (x *GetRelationCountReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var ov GetRelationCountReq_FromFilterOptions
	x.RelationFilterOptions = &ov
	var v FromFilterOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	ov.FromFilterOptions = &v
	return offset, nil
}

func (x *GetRelationCountReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var ov GetRelationCountReq_ToFilterOptions
	x.RelationFilterOptions = &ov
	var v ToFilterOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	ov.ToFilterOptions = &v
	return offset, nil
}

func (x *GetRelationCountReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.RelationType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetRelationCountResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetRelationCountResp[number], err)
}

func (x *GetRelationCountResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateRelationReq) FastWrite(buf []byte) (offset int) {
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

func (x *CreateRelationReq) fastWriteField1(buf []byte) (offset int) {
	if x.FromType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetFromType())
	return offset
}

func (x *CreateRelationReq) fastWriteField2(buf []byte) (offset int) {
	if x.FromId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFromId())
	return offset
}

func (x *CreateRelationReq) fastWriteField3(buf []byte) (offset int) {
	if x.ToType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetToType())
	return offset
}

func (x *CreateRelationReq) fastWriteField4(buf []byte) (offset int) {
	if x.ToId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetToId())
	return offset
}

func (x *CreateRelationReq) fastWriteField5(buf []byte) (offset int) {
	if x.RelationType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetRelationType())
	return offset
}

func (x *CreateRelationResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CreateRelationResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Ok {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetOk())
	return offset
}

func (x *GetRelationReq) FastWrite(buf []byte) (offset int) {
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

func (x *GetRelationReq) fastWriteField1(buf []byte) (offset int) {
	if x.FromType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetFromType())
	return offset
}

func (x *GetRelationReq) fastWriteField2(buf []byte) (offset int) {
	if x.FromId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFromId())
	return offset
}

func (x *GetRelationReq) fastWriteField3(buf []byte) (offset int) {
	if x.ToType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetToType())
	return offset
}

func (x *GetRelationReq) fastWriteField4(buf []byte) (offset int) {
	if x.ToId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetToId())
	return offset
}

func (x *GetRelationReq) fastWriteField5(buf []byte) (offset int) {
	if x.RelationType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetRelationType())
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
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *DeleteRelationReq) fastWriteField1(buf []byte) (offset int) {
	if x.FromType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetFromType())
	return offset
}

func (x *DeleteRelationReq) fastWriteField2(buf []byte) (offset int) {
	if x.FromId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFromId())
	return offset
}

func (x *DeleteRelationReq) fastWriteField3(buf []byte) (offset int) {
	if x.ToType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetToType())
	return offset
}

func (x *DeleteRelationReq) fastWriteField4(buf []byte) (offset int) {
	if x.ToId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetToId())
	return offset
}

func (x *DeleteRelationReq) fastWriteField5(buf []byte) (offset int) {
	if x.RelationType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetRelationType())
	return offset
}

func (x *DeleteRelationResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *FromFilterOptions) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *FromFilterOptions) fastWriteField1(buf []byte) (offset int) {
	if x.FromType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetFromType())
	return offset
}

func (x *FromFilterOptions) fastWriteField2(buf []byte) (offset int) {
	if x.FromId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFromId())
	return offset
}

func (x *FromFilterOptions) fastWriteField3(buf []byte) (offset int) {
	if x.ToType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetToType())
	return offset
}

func (x *ToFilterOptions) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ToFilterOptions) fastWriteField1(buf []byte) (offset int) {
	if x.ToType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetToType())
	return offset
}

func (x *ToFilterOptions) fastWriteField2(buf []byte) (offset int) {
	if x.ToId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetToId())
	return offset
}

func (x *ToFilterOptions) fastWriteField3(buf []byte) (offset int) {
	if x.FromType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetFromType())
	return offset
}

func (x *GetRelationsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *GetRelationsReq) fastWriteField1(buf []byte) (offset int) {
	if x.GetFromFilterOptions() == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetFromFilterOptions())
	return offset
}

func (x *GetRelationsReq) fastWriteField2(buf []byte) (offset int) {
	if x.GetToFilterOptions() == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetToFilterOptions())
	return offset
}

func (x *GetRelationsReq) fastWriteField3(buf []byte) (offset int) {
	if x.RelationType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetRelationType())
	return offset
}

func (x *GetRelationsReq) fastWriteField4(buf []byte) (offset int) {
	if x.PaginationOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 4, x.GetPaginationOptions())
	return offset
}

func (x *GetRelationsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetRelationsResp) fastWriteField1(buf []byte) (offset int) {
	if x.Relations == nil {
		return offset
	}
	for i := range x.GetRelations() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetRelations()[i])
	}
	return offset
}

func (x *GetRelationsResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *GetRelationCountReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetRelationCountReq) fastWriteField1(buf []byte) (offset int) {
	if x.GetFromFilterOptions() == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetFromFilterOptions())
	return offset
}

func (x *GetRelationCountReq) fastWriteField2(buf []byte) (offset int) {
	if x.GetToFilterOptions() == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetToFilterOptions())
	return offset
}

func (x *GetRelationCountReq) fastWriteField3(buf []byte) (offset int) {
	if x.RelationType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetRelationType())
	return offset
}

func (x *GetRelationCountResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetRelationCountResp) fastWriteField1(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetTotal())
	return offset
}

func (x *CreateRelationReq) Size() (n int) {
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

func (x *CreateRelationReq) sizeField1() (n int) {
	if x.FromType == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetFromType())
	return n
}

func (x *CreateRelationReq) sizeField2() (n int) {
	if x.FromId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFromId())
	return n
}

func (x *CreateRelationReq) sizeField3() (n int) {
	if x.ToType == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetToType())
	return n
}

func (x *CreateRelationReq) sizeField4() (n int) {
	if x.ToId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetToId())
	return n
}

func (x *CreateRelationReq) sizeField5() (n int) {
	if x.RelationType == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetRelationType())
	return n
}

func (x *CreateRelationResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CreateRelationResp) sizeField1() (n int) {
	if !x.Ok {
		return n
	}
	n += fastpb.SizeBool(1, x.GetOk())
	return n
}

func (x *GetRelationReq) Size() (n int) {
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

func (x *GetRelationReq) sizeField1() (n int) {
	if x.FromType == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetFromType())
	return n
}

func (x *GetRelationReq) sizeField2() (n int) {
	if x.FromId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFromId())
	return n
}

func (x *GetRelationReq) sizeField3() (n int) {
	if x.ToType == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetToType())
	return n
}

func (x *GetRelationReq) sizeField4() (n int) {
	if x.ToId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetToId())
	return n
}

func (x *GetRelationReq) sizeField5() (n int) {
	if x.RelationType == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetRelationType())
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
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *DeleteRelationReq) sizeField1() (n int) {
	if x.FromType == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetFromType())
	return n
}

func (x *DeleteRelationReq) sizeField2() (n int) {
	if x.FromId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFromId())
	return n
}

func (x *DeleteRelationReq) sizeField3() (n int) {
	if x.ToType == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetToType())
	return n
}

func (x *DeleteRelationReq) sizeField4() (n int) {
	if x.ToId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetToId())
	return n
}

func (x *DeleteRelationReq) sizeField5() (n int) {
	if x.RelationType == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetRelationType())
	return n
}

func (x *DeleteRelationResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *FromFilterOptions) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *FromFilterOptions) sizeField1() (n int) {
	if x.FromType == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetFromType())
	return n
}

func (x *FromFilterOptions) sizeField2() (n int) {
	if x.FromId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFromId())
	return n
}

func (x *FromFilterOptions) sizeField3() (n int) {
	if x.ToType == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetToType())
	return n
}

func (x *ToFilterOptions) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ToFilterOptions) sizeField1() (n int) {
	if x.ToType == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetToType())
	return n
}

func (x *ToFilterOptions) sizeField2() (n int) {
	if x.ToId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetToId())
	return n
}

func (x *ToFilterOptions) sizeField3() (n int) {
	if x.FromType == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetFromType())
	return n
}

func (x *GetRelationsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *GetRelationsReq) sizeField1() (n int) {
	if x.GetFromFilterOptions() == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetFromFilterOptions())
	return n
}

func (x *GetRelationsReq) sizeField2() (n int) {
	if x.GetToFilterOptions() == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetToFilterOptions())
	return n
}

func (x *GetRelationsReq) sizeField3() (n int) {
	if x.RelationType == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetRelationType())
	return n
}

func (x *GetRelationsReq) sizeField4() (n int) {
	if x.PaginationOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(4, x.GetPaginationOptions())
	return n
}

func (x *GetRelationsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetRelationsResp) sizeField1() (n int) {
	if x.Relations == nil {
		return n
	}
	for i := range x.GetRelations() {
		n += fastpb.SizeMessage(1, x.GetRelations()[i])
	}
	return n
}

func (x *GetRelationsResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *GetRelationCountReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetRelationCountReq) sizeField1() (n int) {
	if x.GetFromFilterOptions() == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetFromFilterOptions())
	return n
}

func (x *GetRelationCountReq) sizeField2() (n int) {
	if x.GetToFilterOptions() == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetToFilterOptions())
	return n
}

func (x *GetRelationCountReq) sizeField3() (n int) {
	if x.RelationType == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetRelationType())
	return n
}

func (x *GetRelationCountResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetRelationCountResp) sizeField1() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetTotal())
	return n
}

var fieldIDToName_CreateRelationReq = map[int32]string{
	1: "FromType",
	2: "FromId",
	3: "ToType",
	4: "ToId",
	5: "RelationType",
}

var fieldIDToName_CreateRelationResp = map[int32]string{
	1: "Ok",
}

var fieldIDToName_GetRelationReq = map[int32]string{
	1: "FromType",
	2: "FromId",
	3: "ToType",
	4: "ToId",
	5: "RelationType",
}

var fieldIDToName_GetRelationResp = map[int32]string{
	1: "Ok",
}

var fieldIDToName_DeleteRelationReq = map[int32]string{
	1: "FromType",
	2: "FromId",
	3: "ToType",
	4: "ToId",
	5: "RelationType",
}

var fieldIDToName_DeleteRelationResp = map[int32]string{}

var fieldIDToName_FromFilterOptions = map[int32]string{
	1: "FromType",
	2: "FromId",
	3: "ToType",
}

var fieldIDToName_ToFilterOptions = map[int32]string{
	1: "ToType",
	2: "ToId",
	3: "FromType",
}

var fieldIDToName_GetRelationsReq = map[int32]string{
	1: "FromFilterOptions",
	2: "ToFilterOptions",
	3: "RelationType",
	4: "PaginationOptions",
}

var fieldIDToName_GetRelationsResp = map[int32]string{
	1: "Relations",
	2: "Total",
}

var fieldIDToName_GetRelationCountReq = map[int32]string{
	1: "FromFilterOptions",
	2: "ToFilterOptions",
	3: "RelationType",
}

var fieldIDToName_GetRelationCountResp = map[int32]string{
	1: "Total",
}

var _ = basic.File_basic_pagination_proto
