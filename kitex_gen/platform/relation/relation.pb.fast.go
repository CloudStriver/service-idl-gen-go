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

func (x *Relation) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Relation[number], err)
}

func (x *Relation) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FromType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Relation) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FromId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Relation) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ToType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Relation) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ToId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Relation) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.RelationType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *Relation) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.CreateTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Relation) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.UpdateTime, offset, err = fastpb.ReadInt64(buf, _type)
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
	x.FromType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RelationInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FromId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RelationInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ToType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *RelationInfo) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ToId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RelationInfo) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.RelationType, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

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

func (x *CreateRelationReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.IsOnly, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
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

func (x *RelationFilterOptions) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationFilterOptions[number], err)
}

func (x *RelationFilterOptions) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt32(buf, _type)
	x.OnlyFromType = &tmp
	return offset, err
}

func (x *RelationFilterOptions) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.OnlyFromId = &tmp
	return offset, err
}

func (x *RelationFilterOptions) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt32(buf, _type)
	x.OnlyToType = &tmp
	return offset, err
}

func (x *RelationFilterOptions) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.OnlyToId = &tmp
	return offset, err
}

func (x *RelationFilterOptions) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt32(buf, _type)
	x.OnlyRelationType = &tmp
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
	var v RelationFilterOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.FilterOptions = &v
	return offset, nil
}

func (x *GetRelationsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
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
	x.Total, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *GetRelationsResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.LastToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
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
	var v Relation
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Relation = &v
	return offset, nil
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
	x.Id, offset, err = fastpb.ReadString(buf, _type)
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

func (x *Relation) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	return offset
}

func (x *Relation) fastWriteField1(buf []byte) (offset int) {
	if x.FromType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetFromType())
	return offset
}

func (x *Relation) fastWriteField2(buf []byte) (offset int) {
	if x.FromId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFromId())
	return offset
}

func (x *Relation) fastWriteField3(buf []byte) (offset int) {
	if x.ToType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetToType())
	return offset
}

func (x *Relation) fastWriteField4(buf []byte) (offset int) {
	if x.ToId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetToId())
	return offset
}

func (x *Relation) fastWriteField5(buf []byte) (offset int) {
	if x.RelationType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, x.GetRelationType())
	return offset
}

func (x *Relation) fastWriteField6(buf []byte) (offset int) {
	if x.CreateTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 6, x.GetCreateTime())
	return offset
}

func (x *Relation) fastWriteField7(buf []byte) (offset int) {
	if x.UpdateTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.GetUpdateTime())
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
	if x.FromType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetFromType())
	return offset
}

func (x *RelationInfo) fastWriteField2(buf []byte) (offset int) {
	if x.FromId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFromId())
	return offset
}

func (x *RelationInfo) fastWriteField3(buf []byte) (offset int) {
	if x.ToType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetToType())
	return offset
}

func (x *RelationInfo) fastWriteField4(buf []byte) (offset int) {
	if x.ToId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetToId())
	return offset
}

func (x *RelationInfo) fastWriteField5(buf []byte) (offset int) {
	if x.RelationType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, x.GetRelationType())
	return offset
}

func (x *CreateRelationReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CreateRelationReq) fastWriteField1(buf []byte) (offset int) {
	if x.Relation == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetRelation())
	return offset
}

func (x *CreateRelationReq) fastWriteField2(buf []byte) (offset int) {
	if !x.IsOnly {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 2, x.GetIsOnly())
	return offset
}

func (x *CreateRelationResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *RelationFilterOptions) FastWrite(buf []byte) (offset int) {
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

func (x *RelationFilterOptions) fastWriteField1(buf []byte) (offset int) {
	if x.OnlyFromType == nil {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetOnlyFromType())
	return offset
}

func (x *RelationFilterOptions) fastWriteField2(buf []byte) (offset int) {
	if x.OnlyFromId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetOnlyFromId())
	return offset
}

func (x *RelationFilterOptions) fastWriteField3(buf []byte) (offset int) {
	if x.OnlyToType == nil {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetOnlyToType())
	return offset
}

func (x *RelationFilterOptions) fastWriteField4(buf []byte) (offset int) {
	if x.OnlyToId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetOnlyToId())
	return offset
}

func (x *RelationFilterOptions) fastWriteField5(buf []byte) (offset int) {
	if x.OnlyRelationType == nil {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, x.GetOnlyRelationType())
	return offset
}

func (x *GetRelationsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetRelationsReq) fastWriteField1(buf []byte) (offset int) {
	if x.FilterOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetFilterOptions())
	return offset
}

func (x *GetRelationsReq) fastWriteField2(buf []byte) (offset int) {
	if x.PaginationOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetPaginationOptions())
	return offset
}

func (x *GetRelationsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
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
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *GetRelationsResp) fastWriteField3(buf []byte) (offset int) {
	if x.LastToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetLastToken())
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
	if x.Relation == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetRelation())
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
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *DeleteRelationResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *Relation) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	return n
}

func (x *Relation) sizeField1() (n int) {
	if x.FromType == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetFromType())
	return n
}

func (x *Relation) sizeField2() (n int) {
	if x.FromId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFromId())
	return n
}

func (x *Relation) sizeField3() (n int) {
	if x.ToType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetToType())
	return n
}

func (x *Relation) sizeField4() (n int) {
	if x.ToId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetToId())
	return n
}

func (x *Relation) sizeField5() (n int) {
	if x.RelationType == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, x.GetRelationType())
	return n
}

func (x *Relation) sizeField6() (n int) {
	if x.CreateTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(6, x.GetCreateTime())
	return n
}

func (x *Relation) sizeField7() (n int) {
	if x.UpdateTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(7, x.GetUpdateTime())
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
	if x.FromType == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetFromType())
	return n
}

func (x *RelationInfo) sizeField2() (n int) {
	if x.FromId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFromId())
	return n
}

func (x *RelationInfo) sizeField3() (n int) {
	if x.ToType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetToType())
	return n
}

func (x *RelationInfo) sizeField4() (n int) {
	if x.ToId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetToId())
	return n
}

func (x *RelationInfo) sizeField5() (n int) {
	if x.RelationType == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, x.GetRelationType())
	return n
}

func (x *CreateRelationReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CreateRelationReq) sizeField1() (n int) {
	if x.Relation == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetRelation())
	return n
}

func (x *CreateRelationReq) sizeField2() (n int) {
	if !x.IsOnly {
		return n
	}
	n += fastpb.SizeBool(2, x.GetIsOnly())
	return n
}

func (x *CreateRelationResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *RelationFilterOptions) Size() (n int) {
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

func (x *RelationFilterOptions) sizeField1() (n int) {
	if x.OnlyFromType == nil {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetOnlyFromType())
	return n
}

func (x *RelationFilterOptions) sizeField2() (n int) {
	if x.OnlyFromId == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetOnlyFromId())
	return n
}

func (x *RelationFilterOptions) sizeField3() (n int) {
	if x.OnlyToType == nil {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetOnlyToType())
	return n
}

func (x *RelationFilterOptions) sizeField4() (n int) {
	if x.OnlyToId == nil {
		return n
	}
	n += fastpb.SizeString(4, x.GetOnlyToId())
	return n
}

func (x *RelationFilterOptions) sizeField5() (n int) {
	if x.OnlyRelationType == nil {
		return n
	}
	n += fastpb.SizeInt32(5, x.GetOnlyRelationType())
	return n
}

func (x *GetRelationsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetRelationsReq) sizeField1() (n int) {
	if x.FilterOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetFilterOptions())
	return n
}

func (x *GetRelationsReq) sizeField2() (n int) {
	if x.PaginationOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetPaginationOptions())
	return n
}

func (x *GetRelationsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
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
	n += fastpb.SizeInt32(2, x.GetTotal())
	return n
}

func (x *GetRelationsResp) sizeField3() (n int) {
	if x.LastToken == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetLastToken())
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
	if x.Relation == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetRelation())
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
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *DeleteRelationResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_Relation = map[int32]string{
	1: "FromType",
	2: "FromId",
	3: "ToType",
	4: "ToId",
	5: "RelationType",
	6: "CreateTime",
	7: "UpdateTime",
}

var fieldIDToName_RelationInfo = map[int32]string{
	1: "FromType",
	2: "FromId",
	3: "ToType",
	4: "ToId",
	5: "RelationType",
}

var fieldIDToName_CreateRelationReq = map[int32]string{
	1: "Relation",
	2: "IsOnly",
}

var fieldIDToName_CreateRelationResp = map[int32]string{}

var fieldIDToName_RelationFilterOptions = map[int32]string{
	1: "OnlyFromType",
	2: "OnlyFromId",
	3: "OnlyToType",
	4: "OnlyToId",
	5: "OnlyRelationType",
}

var fieldIDToName_GetRelationsReq = map[int32]string{
	1: "FilterOptions",
	2: "PaginationOptions",
}

var fieldIDToName_GetRelationsResp = map[int32]string{
	1: "Relations",
	2: "Total",
	3: "LastToken",
}

var fieldIDToName_GetRelationReq = map[int32]string{
	1: "RelationInfo",
}

var fieldIDToName_GetRelationResp = map[int32]string{
	1: "Relation",
}

var fieldIDToName_DeleteRelationReq = map[int32]string{
	1: "Id",
}

var fieldIDToName_DeleteRelationResp = map[int32]string{}

var _ = basic.File_basic_pagination_proto
