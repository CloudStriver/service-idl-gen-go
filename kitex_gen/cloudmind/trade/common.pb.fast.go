// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package trade

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *UpdateBalanceReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateBalanceReq[number], err)
}

func (x *UpdateBalanceReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateBalanceReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Oldbalance, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UpdateBalanceReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Balance, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UpdateBalanceReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.BalanceType = BalanceType(v)
	return offset, nil
}

func (x *UpdateBalanceResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *GetBalanceReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetBalanceReq[number], err)
}

func (x *GetBalanceReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetBalanceResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetBalanceResp[number], err)
}

func (x *GetBalanceResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Flow, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetBalanceResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Memory, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetBalanceResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Point, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateBalanceReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateBalanceReq[number], err)
}

func (x *CreateBalanceReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateBalanceResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *AddStockReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AddStockReq[number], err)
}

func (x *AddStockReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ProductId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AddStockReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Amount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AddStockResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *GetStockReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetStockReq[number], err)
}

func (x *GetStockReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ProductId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetStockResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetStockResp[number], err)
}

func (x *GetStockResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Stock, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetStocksReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetStocksReq[number], err)
}

func (x *GetStocksReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.ProductIds = append(x.ProductIds, v)
	return offset, err
}

func (x *GetStocksResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetStocksResp[number], err)
}

func (x *GetStocksResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int64
			v, offset, err = fastpb.ReadInt64(buf, _type)
			if err != nil {
				return offset, err
			}
			x.Stocks = append(x.Stocks, v)
			return offset, err
		})
	return offset, err
}

func (x *UpdateBalanceReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *UpdateBalanceReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *UpdateBalanceReq) fastWriteField2(buf []byte) (offset int) {
	if x.Oldbalance == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetOldbalance())
	return offset
}

func (x *UpdateBalanceReq) fastWriteField3(buf []byte) (offset int) {
	if x.Balance == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetBalance())
	return offset
}

func (x *UpdateBalanceReq) fastWriteField4(buf []byte) (offset int) {
	if x.BalanceType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, int32(x.GetBalanceType()))
	return offset
}

func (x *UpdateBalanceResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *GetBalanceReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetBalanceReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetBalanceResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetBalanceResp) fastWriteField1(buf []byte) (offset int) {
	if x.Flow == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetFlow())
	return offset
}

func (x *GetBalanceResp) fastWriteField2(buf []byte) (offset int) {
	if x.Memory == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetMemory())
	return offset
}

func (x *GetBalanceResp) fastWriteField3(buf []byte) (offset int) {
	if x.Point == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetPoint())
	return offset
}

func (x *CreateBalanceReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CreateBalanceReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *CreateBalanceResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *AddStockReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *AddStockReq) fastWriteField1(buf []byte) (offset int) {
	if x.ProductId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetProductId())
	return offset
}

func (x *AddStockReq) fastWriteField3(buf []byte) (offset int) {
	if x.Amount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetAmount())
	return offset
}

func (x *AddStockResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *GetStockReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetStockReq) fastWriteField1(buf []byte) (offset int) {
	if x.ProductId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetProductId())
	return offset
}

func (x *GetStockResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetStockResp) fastWriteField1(buf []byte) (offset int) {
	if x.Stock == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetStock())
	return offset
}

func (x *GetStocksReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetStocksReq) fastWriteField1(buf []byte) (offset int) {
	if len(x.ProductIds) == 0 {
		return offset
	}
	for i := range x.GetProductIds() {
		offset += fastpb.WriteString(buf[offset:], 1, x.GetProductIds()[i])
	}
	return offset
}

func (x *GetStocksResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetStocksResp) fastWriteField1(buf []byte) (offset int) {
	if len(x.Stocks) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 1, len(x.GetStocks()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt64(buf[offset:], numTagOrKey, x.GetStocks()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *UpdateBalanceReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *UpdateBalanceReq) sizeField1() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserId())
	return n
}

func (x *UpdateBalanceReq) sizeField2() (n int) {
	if x.Oldbalance == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetOldbalance())
	return n
}

func (x *UpdateBalanceReq) sizeField3() (n int) {
	if x.Balance == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetBalance())
	return n
}

func (x *UpdateBalanceReq) sizeField4() (n int) {
	if x.BalanceType == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, int32(x.GetBalanceType()))
	return n
}

func (x *UpdateBalanceResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *GetBalanceReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetBalanceReq) sizeField1() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserId())
	return n
}

func (x *GetBalanceResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetBalanceResp) sizeField1() (n int) {
	if x.Flow == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetFlow())
	return n
}

func (x *GetBalanceResp) sizeField2() (n int) {
	if x.Memory == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetMemory())
	return n
}

func (x *GetBalanceResp) sizeField3() (n int) {
	if x.Point == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetPoint())
	return n
}

func (x *CreateBalanceReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CreateBalanceReq) sizeField1() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserId())
	return n
}

func (x *CreateBalanceResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *AddStockReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField3()
	return n
}

func (x *AddStockReq) sizeField1() (n int) {
	if x.ProductId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetProductId())
	return n
}

func (x *AddStockReq) sizeField3() (n int) {
	if x.Amount == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetAmount())
	return n
}

func (x *AddStockResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *GetStockReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetStockReq) sizeField1() (n int) {
	if x.ProductId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetProductId())
	return n
}

func (x *GetStockResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetStockResp) sizeField1() (n int) {
	if x.Stock == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetStock())
	return n
}

func (x *GetStocksReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetStocksReq) sizeField1() (n int) {
	if len(x.ProductIds) == 0 {
		return n
	}
	for i := range x.GetProductIds() {
		n += fastpb.SizeString(1, x.GetProductIds()[i])
	}
	return n
}

func (x *GetStocksResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetStocksResp) sizeField1() (n int) {
	if len(x.Stocks) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(1, len(x.GetStocks()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt64(numTagOrKey, x.GetStocks()[numIdxOrVal])
			return n
		})
	return n
}

var fieldIDToName_UpdateBalanceReq = map[int32]string{
	1: "UserId",
	2: "Oldbalance",
	3: "Balance",
	4: "BalanceType",
}

var fieldIDToName_UpdateBalanceResp = map[int32]string{}

var fieldIDToName_GetBalanceReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetBalanceResp = map[int32]string{
	1: "Flow",
	2: "Memory",
	3: "Point",
}

var fieldIDToName_CreateBalanceReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_CreateBalanceResp = map[int32]string{}

var fieldIDToName_AddStockReq = map[int32]string{
	1: "ProductId",
	3: "Amount",
}

var fieldIDToName_AddStockResp = map[int32]string{}

var fieldIDToName_GetStockReq = map[int32]string{
	1: "ProductId",
}

var fieldIDToName_GetStockResp = map[int32]string{
	1: "Stock",
}

var fieldIDToName_GetStocksReq = map[int32]string{
	1: "ProductIds",
}

var fieldIDToName_GetStocksResp = map[int32]string{
	1: "Stocks",
}
