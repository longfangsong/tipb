// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema.proto

package tipb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TableInfo struct {
	TableId              int64         `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id"`
	Columns              []*ColumnInfo `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TableInfo) Reset()         { *m = TableInfo{} }
func (m *TableInfo) String() string { return proto.CompactTextString(m) }
func (*TableInfo) ProtoMessage()    {}
func (*TableInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{0}
}
func (m *TableInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TableInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TableInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TableInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableInfo.Merge(m, src)
}
func (m *TableInfo) XXX_Size() int {
	return m.Size()
}
func (m *TableInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TableInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TableInfo proto.InternalMessageInfo

func (m *TableInfo) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *TableInfo) GetColumns() []*ColumnInfo {
	if m != nil {
		return m.Columns
	}
	return nil
}

type ColumnInfo struct {
	ColumnId             int64    `protobuf:"varint,1,opt,name=column_id,json=columnId" json:"column_id"`
	Tp                   int32    `protobuf:"varint,2,opt,name=tp" json:"tp"`
	Collation            int32    `protobuf:"varint,3,opt,name=collation" json:"collation"`
	ColumnLen            int32    `protobuf:"varint,4,opt,name=columnLen" json:"columnLen"`
	Decimal              int32    `protobuf:"varint,5,opt,name=decimal" json:"decimal"`
	Flag                 int32    `protobuf:"varint,6,opt,name=flag" json:"flag"`
	Elems                []string `protobuf:"bytes,7,rep,name=elems" json:"elems,omitempty"`
	DefaultVal           []byte   `protobuf:"bytes,8,opt,name=default_val,json=defaultVal" json:"default_val,omitempty"`
	PkHandle             bool     `protobuf:"varint,21,opt,name=pk_handle,json=pkHandle" json:"pk_handle"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ColumnInfo) Reset()         { *m = ColumnInfo{} }
func (m *ColumnInfo) String() string { return proto.CompactTextString(m) }
func (*ColumnInfo) ProtoMessage()    {}
func (*ColumnInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{1}
}
func (m *ColumnInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ColumnInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ColumnInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ColumnInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnInfo.Merge(m, src)
}
func (m *ColumnInfo) XXX_Size() int {
	return m.Size()
}
func (m *ColumnInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnInfo proto.InternalMessageInfo

func (m *ColumnInfo) GetColumnId() int64 {
	if m != nil {
		return m.ColumnId
	}
	return 0
}

func (m *ColumnInfo) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

func (m *ColumnInfo) GetCollation() int32 {
	if m != nil {
		return m.Collation
	}
	return 0
}

func (m *ColumnInfo) GetColumnLen() int32 {
	if m != nil {
		return m.ColumnLen
	}
	return 0
}

func (m *ColumnInfo) GetDecimal() int32 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *ColumnInfo) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *ColumnInfo) GetElems() []string {
	if m != nil {
		return m.Elems
	}
	return nil
}

func (m *ColumnInfo) GetDefaultVal() []byte {
	if m != nil {
		return m.DefaultVal
	}
	return nil
}

func (m *ColumnInfo) GetPkHandle() bool {
	if m != nil {
		return m.PkHandle
	}
	return false
}

type IndexInfo struct {
	TableId              int64         `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id"`
	IndexId              int64         `protobuf:"varint,2,opt,name=index_id,json=indexId" json:"index_id"`
	Columns              []*ColumnInfo `protobuf:"bytes,3,rep,name=columns" json:"columns,omitempty"`
	Unique               bool          `protobuf:"varint,4,opt,name=unique" json:"unique"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IndexInfo) Reset()         { *m = IndexInfo{} }
func (m *IndexInfo) String() string { return proto.CompactTextString(m) }
func (*IndexInfo) ProtoMessage()    {}
func (*IndexInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{2}
}
func (m *IndexInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexInfo.Merge(m, src)
}
func (m *IndexInfo) XXX_Size() int {
	return m.Size()
}
func (m *IndexInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IndexInfo proto.InternalMessageInfo

func (m *IndexInfo) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *IndexInfo) GetIndexId() int64 {
	if m != nil {
		return m.IndexId
	}
	return 0
}

func (m *IndexInfo) GetColumns() []*ColumnInfo {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *IndexInfo) GetUnique() bool {
	if m != nil {
		return m.Unique
	}
	return false
}

// KeyRange is the encoded index key range, low is closed, high is open. (low <= x < high)
type KeyRange struct {
	Low                  []byte   `protobuf:"bytes,1,opt,name=low" json:"low,omitempty"`
	High                 []byte   `protobuf:"bytes,2,opt,name=high" json:"high,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyRange) Reset()         { *m = KeyRange{} }
func (m *KeyRange) String() string { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()    {}
func (*KeyRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c5fb4d8cc22d66a, []int{3}
}
func (m *KeyRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRange.Merge(m, src)
}
func (m *KeyRange) XXX_Size() int {
	return m.Size()
}
func (m *KeyRange) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRange.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRange proto.InternalMessageInfo

func (m *KeyRange) GetLow() []byte {
	if m != nil {
		return m.Low
	}
	return nil
}

func (m *KeyRange) GetHigh() []byte {
	if m != nil {
		return m.High
	}
	return nil
}

func init() {
	proto.RegisterType((*TableInfo)(nil), "tipb.TableInfo")
	proto.RegisterType((*ColumnInfo)(nil), "tipb.ColumnInfo")
	proto.RegisterType((*IndexInfo)(nil), "tipb.IndexInfo")
	proto.RegisterType((*KeyRange)(nil), "tipb.KeyRange")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor_1c5fb4d8cc22d66a) }

var fileDescriptor_1c5fb4d8cc22d66a = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xd7, 0x49, 0xba, 0x4d, 0x67, 0x7b, 0x58, 0x59, 0x5d, 0xc9, 0x42, 0x28, 0x0d, 0x39,
	0x05, 0x0e, 0x01, 0xf1, 0x08, 0xcb, 0x85, 0x0a, 0x0e, 0x28, 0x42, 0x88, 0x5b, 0xe5, 0xc6, 0x6e,
	0x62, 0xad, 0x63, 0x1b, 0xea, 0xf0, 0xe7, 0x49, 0xe0, 0xca, 0xdb, 0xec, 0x91, 0x27, 0x40, 0xa8,
	0xbc, 0x08, 0xb2, 0x93, 0x68, 0x43, 0x2f, 0xec, 0x6d, 0xfc, 0xfb, 0x3e, 0x8f, 0xbe, 0x99, 0x81,
	0xe5, 0xa1, 0x6a, 0x78, 0x4b, 0x0b, 0xf3, 0x51, 0x5b, 0x8d, 0x23, 0x2b, 0xcc, 0xee, 0xc1, 0xaa,
	0xd6, 0xb5, 0xf6, 0xe0, 0xa9, 0xab, 0x7a, 0x2d, 0x7b, 0x0f, 0x8b, 0xb7, 0x74, 0x27, 0xf9, 0x46,
	0xed, 0x35, 0x5e, 0x43, 0x6c, 0xdd, 0x63, 0x2b, 0x18, 0x41, 0x29, 0xca, 0xc3, 0xeb, 0xe8, 0xf6,
	0xd7, 0xfa, 0xac, 0x9c, 0x7b, 0xba, 0x61, 0xf8, 0x09, 0xcc, 0x2b, 0x2d, 0xbb, 0x56, 0x1d, 0x48,
	0x90, 0x86, 0xf9, 0xc5, 0xf3, 0xcb, 0xc2, 0xf5, 0x2e, 0x5e, 0x78, 0xe8, 0x7a, 0x94, 0xa3, 0x21,
	0xfb, 0x11, 0x00, 0xdc, 0x71, 0xfc, 0x08, 0x16, 0xbd, 0x72, 0xda, 0x3c, 0xee, 0xf1, 0x86, 0xe1,
	0x15, 0x04, 0xd6, 0x90, 0x20, 0x45, 0xf9, 0x6c, 0xd0, 0x02, 0x6b, 0x70, 0xe6, 0x3f, 0x4a, 0x6a,
	0x85, 0x56, 0x24, 0x9c, 0x88, 0x77, 0x78, 0xf0, 0x74, 0xad, 0x7a, 0xcd, 0x15, 0x89, 0x4e, 0x3c,
	0x3d, 0xc6, 0x09, 0xcc, 0x19, 0xaf, 0x44, 0x4b, 0x25, 0x99, 0x4d, 0x1c, 0x23, 0xc4, 0x04, 0xa2,
	0xbd, 0xa4, 0x35, 0x39, 0x9f, 0x88, 0x9e, 0xe0, 0x15, 0xcc, 0xb8, 0xe4, 0xed, 0x81, 0xcc, 0xd3,
	0x30, 0x5f, 0x94, 0xfd, 0x03, 0xaf, 0xe1, 0x82, 0xf1, 0x3d, 0xed, 0xa4, 0xdd, 0x7e, 0xa2, 0x92,
	0xc4, 0x29, 0xca, 0x97, 0x25, 0x0c, 0xe8, 0x1d, 0x95, 0x6e, 0x62, 0x73, 0xb3, 0x6d, 0xa8, 0x62,
	0x92, 0x93, 0xab, 0x14, 0xe5, 0xf1, 0x38, 0xb1, 0xb9, 0x79, 0xe9, 0x69, 0xf6, 0x0d, 0xc1, 0x62,
	0xa3, 0x18, 0xff, 0x72, 0xbf, 0xf5, 0xaf, 0x21, 0x16, 0xce, 0xed, 0x0c, 0xc1, 0xd4, 0xe0, 0xe9,
	0xbf, 0xf7, 0x09, 0xff, 0x73, 0x1f, 0xfc, 0x10, 0xce, 0x3b, 0x25, 0x3e, 0x74, 0xdc, 0x2f, 0x6c,
	0xcc, 0x36, 0xb0, 0xec, 0x19, 0xc4, 0xaf, 0xf8, 0xd7, 0x92, 0xaa, 0x9a, 0xe3, 0x4b, 0x08, 0xa5,
	0xfe, 0xec, 0x23, 0x2d, 0x4b, 0x57, 0x62, 0x0c, 0x51, 0x23, 0xea, 0xc6, 0x87, 0x58, 0x96, 0xbe,
	0xbe, 0x7e, 0x7c, 0x7b, 0x4c, 0xd0, 0xcf, 0x63, 0x82, 0x7e, 0x1f, 0x13, 0xf4, 0xfd, 0x4f, 0x72,
	0x06, 0x57, 0x95, 0x6e, 0x0b, 0x23, 0x54, 0x5d, 0x51, 0x53, 0x58, 0xc1, 0x76, 0x3e, 0xcc, 0x1b,
	0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x75, 0x06, 0xaa, 0x3a, 0x9f, 0x02, 0x00, 0x00,
}

func (m *TableInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TableInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Columns) > 0 {
		for iNdEx := len(m.Columns) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Columns[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSchema(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	i = encodeVarintSchema(dAtA, i, uint64(m.TableId))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *ColumnInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ColumnInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ColumnInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	i--
	if m.PkHandle {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0xa8
	if m.DefaultVal != nil {
		i -= len(m.DefaultVal)
		copy(dAtA[i:], m.DefaultVal)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.DefaultVal)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Elems) > 0 {
		for iNdEx := len(m.Elems) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Elems[iNdEx])
			copy(dAtA[i:], m.Elems[iNdEx])
			i = encodeVarintSchema(dAtA, i, uint64(len(m.Elems[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	i = encodeVarintSchema(dAtA, i, uint64(m.Flag))
	i--
	dAtA[i] = 0x30
	i = encodeVarintSchema(dAtA, i, uint64(m.Decimal))
	i--
	dAtA[i] = 0x28
	i = encodeVarintSchema(dAtA, i, uint64(m.ColumnLen))
	i--
	dAtA[i] = 0x20
	i = encodeVarintSchema(dAtA, i, uint64(m.Collation))
	i--
	dAtA[i] = 0x18
	i = encodeVarintSchema(dAtA, i, uint64(m.Tp))
	i--
	dAtA[i] = 0x10
	i = encodeVarintSchema(dAtA, i, uint64(m.ColumnId))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *IndexInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	i--
	if m.Unique {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x20
	if len(m.Columns) > 0 {
		for iNdEx := len(m.Columns) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Columns[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSchema(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	i = encodeVarintSchema(dAtA, i, uint64(m.IndexId))
	i--
	dAtA[i] = 0x10
	i = encodeVarintSchema(dAtA, i, uint64(m.TableId))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *KeyRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.High != nil {
		i -= len(m.High)
		copy(dAtA[i:], m.High)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.High)))
		i--
		dAtA[i] = 0x12
	}
	if m.Low != nil {
		i -= len(m.Low)
		copy(dAtA[i:], m.Low)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Low)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TableInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovSchema(uint64(m.TableId))
	if len(m.Columns) > 0 {
		for _, e := range m.Columns {
			l = e.Size()
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ColumnInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovSchema(uint64(m.ColumnId))
	n += 1 + sovSchema(uint64(m.Tp))
	n += 1 + sovSchema(uint64(m.Collation))
	n += 1 + sovSchema(uint64(m.ColumnLen))
	n += 1 + sovSchema(uint64(m.Decimal))
	n += 1 + sovSchema(uint64(m.Flag))
	if len(m.Elems) > 0 {
		for _, s := range m.Elems {
			l = len(s)
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	if m.DefaultVal != nil {
		l = len(m.DefaultVal)
		n += 1 + l + sovSchema(uint64(l))
	}
	n += 3
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IndexInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovSchema(uint64(m.TableId))
	n += 1 + sovSchema(uint64(m.IndexId))
	if len(m.Columns) > 0 {
		for _, e := range m.Columns {
			l = e.Size()
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	n += 2
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *KeyRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Low != nil {
		l = len(m.Low)
		n += 1 + l + sovSchema(uint64(l))
	}
	if m.High != nil {
		l = len(m.High)
		n += 1 + l + sovSchema(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSchema(x uint64) (n int) {
	return sovSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TableInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TableInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableId", wireType)
			}
			m.TableId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TableId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, &ColumnInfo{})
			if err := m.Columns[len(m.Columns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ColumnInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ColumnInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ColumnInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnId", wireType)
			}
			m.ColumnId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ColumnId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collation", wireType)
			}
			m.Collation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Collation |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnLen", wireType)
			}
			m.ColumnLen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ColumnLen |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimal |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag", wireType)
			}
			m.Flag = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Flag |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elems", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elems = append(m.Elems, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultVal", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultVal = append(m.DefaultVal[:0], dAtA[iNdEx:postIndex]...)
			if m.DefaultVal == nil {
				m.DefaultVal = []byte{}
			}
			iNdEx = postIndex
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PkHandle", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PkHandle = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IndexInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableId", wireType)
			}
			m.TableId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TableId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexId", wireType)
			}
			m.IndexId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, &ColumnInfo{})
			if err := m.Columns[len(m.Columns)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unique", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Unique = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *KeyRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KeyRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Low", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Low = append(m.Low[:0], dAtA[iNdEx:postIndex]...)
			if m.Low == nil {
				m.Low = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field High", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.High = append(m.High[:0], dAtA[iNdEx:postIndex]...)
			if m.High == nil {
				m.High = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSchema = fmt.Errorf("proto: unexpected end of group")
)
