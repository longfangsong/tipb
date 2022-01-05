// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: binlog.proto

package binlog

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

type MutationType int32

const (
	MutationType_Insert    MutationType = 0
	MutationType_Update    MutationType = 1
	MutationType_DeleteID  MutationType = 2
	MutationType_DeletePK  MutationType = 3
	MutationType_DeleteRow MutationType = 4
)

var MutationType_name = map[int32]string{
	0: "Insert",
	1: "Update",
	2: "DeleteID",
	3: "DeletePK",
	4: "DeleteRow",
}

var MutationType_value = map[string]int32{
	"Insert":    0,
	"Update":    1,
	"DeleteID":  2,
	"DeletePK":  3,
	"DeleteRow": 4,
}

func (x MutationType) Enum() *MutationType {
	p := new(MutationType)
	*p = x
	return p
}

func (x MutationType) String() string {
	return proto.EnumName(MutationType_name, int32(x))
}

func (x *MutationType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MutationType_value, data, "MutationType")
	if err != nil {
		return err
	}
	*x = MutationType(value)
	return nil
}

func (MutationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_07ff5a43190a1b8c, []int{0}
}

type BinlogType int32

const (
	BinlogType_Prewrite BinlogType = 0
	BinlogType_Commit   BinlogType = 1
	BinlogType_Rollback BinlogType = 2
	BinlogType_PreDDL   BinlogType = 3
	BinlogType_PostDDL  BinlogType = 4
)

var BinlogType_name = map[int32]string{
	0: "Prewrite",
	1: "Commit",
	2: "Rollback",
	3: "PreDDL",
	4: "PostDDL",
}

var BinlogType_value = map[string]int32{
	"Prewrite": 0,
	"Commit":   1,
	"Rollback": 2,
	"PreDDL":   3,
	"PostDDL":  4,
}

func (x BinlogType) Enum() *BinlogType {
	p := new(BinlogType)
	*p = x
	return p
}

func (x BinlogType) String() string {
	return proto.EnumName(BinlogType_name, int32(x))
}

func (x *BinlogType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BinlogType_value, data, "BinlogType")
	if err != nil {
		return err
	}
	*x = BinlogType(value)
	return nil
}

func (BinlogType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_07ff5a43190a1b8c, []int{1}
}

// TableMutation contains mutations in a table.
type TableMutation struct {
	TableId int64 `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id"`
	// The inserted row contains all column values.
	InsertedRows [][]byte `protobuf:"bytes,2,rep,name=inserted_rows,json=insertedRows" json:"inserted_rows,omitempty"`
	// The updated row contains old values and new values of the row.
	UpdatedRows [][]byte `protobuf:"bytes,3,rep,name=updated_rows,json=updatedRows" json:"updated_rows,omitempty"`
	// Obsolete field.
	DeletedIds []int64 `protobuf:"varint,4,rep,name=deleted_ids,json=deletedIds" json:"deleted_ids,omitempty"`
	// Obsolete field.
	DeletedPks [][]byte `protobuf:"bytes,5,rep,name=deleted_pks,json=deletedPks" json:"deleted_pks,omitempty"`
	// The row value of the deleted row.
	DeletedRows [][]byte `protobuf:"bytes,6,rep,name=deleted_rows,json=deletedRows" json:"deleted_rows,omitempty"`
	// Used to apply table mutations in original sequence.
	Sequence             []MutationType `protobuf:"varint,7,rep,name=sequence,enum=binlog.MutationType" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TableMutation) Reset()         { *m = TableMutation{} }
func (m *TableMutation) String() string { return proto.CompactTextString(m) }
func (*TableMutation) ProtoMessage()    {}
func (*TableMutation) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ff5a43190a1b8c, []int{0}
}
func (m *TableMutation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TableMutation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TableMutation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TableMutation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableMutation.Merge(m, src)
}
func (m *TableMutation) XXX_Size() int {
	return m.Size()
}
func (m *TableMutation) XXX_DiscardUnknown() {
	xxx_messageInfo_TableMutation.DiscardUnknown(m)
}

var xxx_messageInfo_TableMutation proto.InternalMessageInfo

func (m *TableMutation) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *TableMutation) GetInsertedRows() [][]byte {
	if m != nil {
		return m.InsertedRows
	}
	return nil
}

func (m *TableMutation) GetUpdatedRows() [][]byte {
	if m != nil {
		return m.UpdatedRows
	}
	return nil
}

func (m *TableMutation) GetDeletedIds() []int64 {
	if m != nil {
		return m.DeletedIds
	}
	return nil
}

func (m *TableMutation) GetDeletedPks() [][]byte {
	if m != nil {
		return m.DeletedPks
	}
	return nil
}

func (m *TableMutation) GetDeletedRows() [][]byte {
	if m != nil {
		return m.DeletedRows
	}
	return nil
}

func (m *TableMutation) GetSequence() []MutationType {
	if m != nil {
		return m.Sequence
	}
	return nil
}

type PrewriteValue struct {
	SchemaVersion        int64           `protobuf:"varint,1,opt,name=schema_version,json=schemaVersion" json:"schema_version"`
	Mutations            []TableMutation `protobuf:"bytes,2,rep,name=mutations" json:"mutations"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *PrewriteValue) Reset()         { *m = PrewriteValue{} }
func (m *PrewriteValue) String() string { return proto.CompactTextString(m) }
func (*PrewriteValue) ProtoMessage()    {}
func (*PrewriteValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ff5a43190a1b8c, []int{1}
}
func (m *PrewriteValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrewriteValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PrewriteValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PrewriteValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrewriteValue.Merge(m, src)
}
func (m *PrewriteValue) XXX_Size() int {
	return m.Size()
}
func (m *PrewriteValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PrewriteValue.DiscardUnknown(m)
}

var xxx_messageInfo_PrewriteValue proto.InternalMessageInfo

func (m *PrewriteValue) GetSchemaVersion() int64 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func (m *PrewriteValue) GetMutations() []TableMutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

// Binlog contains all the changes in a transaction, which can be used to reconstruct SQL statement, then export to
// other systems.
type Binlog struct {
	Tp BinlogType `protobuf:"varint,1,opt,name=tp,enum=binlog.BinlogType" json:"tp"`
	// start_ts is used in Prewrite, Commit and Rollback binlog Type.
	// It is used for pairing prewrite log to commit log or rollback log.
	StartTs int64 `protobuf:"varint,2,opt,name=start_ts,json=startTs" json:"start_ts"`
	// commit_ts is used only in binlog type Commit.
	CommitTs int64 `protobuf:"varint,3,opt,name=commit_ts,json=commitTs" json:"commit_ts"`
	// prewrite key is used only in Prewrite binlog type.
	// It is the primary key of the transaction, is used to check that the transaction is
	// commited or not if it failed to pair to commit log or rollback log within a time window.
	PrewriteKey []byte `protobuf:"bytes,4,opt,name=prewrite_key,json=prewriteKey" json:"prewrite_key,omitempty"`
	// prewrite_data is marshalled from PrewriteData type,
	// we do not need to unmarshal prewrite data before the binlog have been successfully paired.
	PrewriteValue []byte `protobuf:"bytes,5,opt,name=prewrite_value,json=prewriteValue" json:"prewrite_value,omitempty"`
	// ddl_query is the original DDL statement query.
	DdlQuery []byte `protobuf:"bytes,6,opt,name=ddl_query,json=ddlQuery" json:"ddl_query,omitempty"`
	// ddl_job_id is used for DDL Binlog.
	// If ddl_job_id is setted, this is a DDL Binlog and ddl_query contains the DDL query, we can query the informations about this job from TiKV.
	DdlJobId int64 `protobuf:"varint,7,opt,name=ddl_job_id,json=ddlJobId" json:"ddl_job_id"`
	// ddl_schema_state is used for DDL Binlog.
	DdlSchemaState       int32    `protobuf:"varint,8,opt,name=ddl_schema_state,json=ddlSchemaState" json:"ddl_schema_state"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Binlog) Reset()         { *m = Binlog{} }
func (m *Binlog) String() string { return proto.CompactTextString(m) }
func (*Binlog) ProtoMessage()    {}
func (*Binlog) Descriptor() ([]byte, []int) {
	return fileDescriptor_07ff5a43190a1b8c, []int{2}
}
func (m *Binlog) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Binlog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Binlog.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Binlog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Binlog.Merge(m, src)
}
func (m *Binlog) XXX_Size() int {
	return m.Size()
}
func (m *Binlog) XXX_DiscardUnknown() {
	xxx_messageInfo_Binlog.DiscardUnknown(m)
}

var xxx_messageInfo_Binlog proto.InternalMessageInfo

func (m *Binlog) GetTp() BinlogType {
	if m != nil {
		return m.Tp
	}
	return BinlogType_Prewrite
}

func (m *Binlog) GetStartTs() int64 {
	if m != nil {
		return m.StartTs
	}
	return 0
}

func (m *Binlog) GetCommitTs() int64 {
	if m != nil {
		return m.CommitTs
	}
	return 0
}

func (m *Binlog) GetPrewriteKey() []byte {
	if m != nil {
		return m.PrewriteKey
	}
	return nil
}

func (m *Binlog) GetPrewriteValue() []byte {
	if m != nil {
		return m.PrewriteValue
	}
	return nil
}

func (m *Binlog) GetDdlQuery() []byte {
	if m != nil {
		return m.DdlQuery
	}
	return nil
}

func (m *Binlog) GetDdlJobId() int64 {
	if m != nil {
		return m.DdlJobId
	}
	return 0
}

func (m *Binlog) GetDdlSchemaState() int32 {
	if m != nil {
		return m.DdlSchemaState
	}
	return 0
}

func init() {
	proto.RegisterEnum("binlog.MutationType", MutationType_name, MutationType_value)
	proto.RegisterEnum("binlog.BinlogType", BinlogType_name, BinlogType_value)
	proto.RegisterType((*TableMutation)(nil), "binlog.TableMutation")
	proto.RegisterType((*PrewriteValue)(nil), "binlog.PrewriteValue")
	proto.RegisterType((*Binlog)(nil), "binlog.Binlog")
}

func init() { proto.RegisterFile("binlog.proto", fileDescriptor_07ff5a43190a1b8c) }

var fileDescriptor_07ff5a43190a1b8c = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xdd, 0x8e, 0xd2, 0x40,
	0x14, 0xc7, 0x69, 0xcb, 0x47, 0x39, 0xb4, 0xa4, 0x99, 0xac, 0x49, 0xa3, 0x09, 0x54, 0x8c, 0x49,
	0xb3, 0x26, 0x68, 0xb8, 0xf3, 0x16, 0xb9, 0xc1, 0x75, 0x13, 0xb6, 0x8b, 0x7b, 0xdb, 0x14, 0x66,
	0x82, 0x95, 0xc2, 0x74, 0x3b, 0xc3, 0x36, 0x3c, 0x85, 0xb7, 0xbe, 0x86, 0x6f, 0xb1, 0x97, 0x3e,
	0x81, 0x31, 0xf8, 0x22, 0x66, 0x3e, 0xca, 0xc7, 0x5d, 0xcf, 0xff, 0xff, 0x9b, 0x73, 0x3a, 0xff,
	0x39, 0xe0, 0x2c, 0xd2, 0x6d, 0x46, 0x57, 0xc3, 0xbc, 0xa0, 0x9c, 0xa2, 0xa6, 0xaa, 0x5e, 0x5e,
	0xad, 0xe8, 0x8a, 0x4a, 0xe9, 0xbd, 0xf8, 0x52, 0xee, 0xe0, 0x87, 0x09, 0xee, 0x3c, 0x59, 0x64,
	0xe4, 0x76, 0xc7, 0x13, 0x9e, 0xd2, 0x2d, 0xea, 0x83, 0xcd, 0x85, 0x10, 0xa7, 0xd8, 0x37, 0x02,
	0x23, 0xb4, 0xc6, 0xf5, 0xe7, 0x3f, 0xfd, 0x5a, 0xd4, 0x92, 0xea, 0x14, 0xa3, 0x37, 0xe0, 0xa6,
	0x5b, 0x46, 0x0a, 0x4e, 0x70, 0x5c, 0xd0, 0x92, 0xf9, 0x66, 0x60, 0x85, 0x4e, 0xe4, 0x54, 0x62,
	0x44, 0x4b, 0x86, 0x5e, 0x83, 0xb3, 0xcb, 0x71, 0x72, 0x64, 0x2c, 0xc9, 0x74, 0xb4, 0x26, 0x91,
	0x3e, 0x74, 0x30, 0xc9, 0x88, 0x40, 0x52, 0xcc, 0xfc, 0x7a, 0x60, 0x85, 0x56, 0x04, 0x5a, 0x9a,
	0xe2, 0x0b, 0x20, 0x5f, 0x33, 0xbf, 0x21, 0x5b, 0x54, 0xc0, 0x6c, 0x2d, 0x87, 0x54, 0x80, 0x1c,
	0xd2, 0x54, 0x43, 0xb4, 0x26, 0x87, 0x7c, 0x00, 0x9b, 0x91, 0xc7, 0x1d, 0xd9, 0x2e, 0x89, 0xdf,
	0x0a, 0xac, 0xb0, 0x3b, 0xba, 0x1a, 0xea, 0x78, 0xaa, 0x1b, 0xcf, 0xf7, 0x39, 0x89, 0x8e, 0xd4,
	0xa0, 0x04, 0x77, 0x56, 0x90, 0xb2, 0x48, 0x39, 0x79, 0x48, 0xb2, 0x1d, 0x41, 0xef, 0xa0, 0xcb,
	0x96, 0xdf, 0xc8, 0x26, 0x89, 0x9f, 0x48, 0xc1, 0x52, 0xba, 0xbd, 0x88, 0xc5, 0x55, 0xde, 0x83,
	0xb2, 0xd0, 0x47, 0x68, 0x6f, 0x74, 0x5f, 0x15, 0x4c, 0x67, 0xf4, 0xa2, 0x1a, 0x78, 0x91, 0xb3,
	0x3e, 0x7e, 0xa2, 0x07, 0xbf, 0x4c, 0x68, 0x8e, 0x25, 0x89, 0x42, 0x30, 0x79, 0x2e, 0xc7, 0x74,
	0x47, 0xa8, 0x3a, 0xae, 0x3c, 0xf1, 0xb7, 0xfa, 0xac, 0xc9, 0x73, 0xf1, 0x5a, 0x8c, 0x27, 0x05,
	0x8f, 0xb9, 0x18, 0x77, 0xf6, 0x5a, 0x52, 0x9d, 0x8b, 0x8c, 0xda, 0x4b, 0xba, 0xd9, 0xa4, 0x92,
	0xb0, 0xce, 0x08, 0x5b, 0xc9, 0x12, 0x71, 0x72, 0x7d, 0xe3, 0x78, 0x4d, 0xf6, 0x7e, 0x3d, 0x30,
	0x44, 0x8c, 0x95, 0x76, 0x43, 0xf6, 0xe8, 0x2d, 0x74, 0x8f, 0xc8, 0x93, 0x48, 0xc5, 0x6f, 0x48,
	0xc8, 0xcd, 0x2f, 0xa2, 0x7a, 0x05, 0x6d, 0x8c, 0xb3, 0xf8, 0x71, 0x47, 0x8a, 0xbd, 0xdf, 0x94,
	0x84, 0x8d, 0x71, 0x76, 0x27, 0x6a, 0x34, 0x00, 0x10, 0xe6, 0x77, 0xba, 0x10, 0xab, 0xd5, 0x3a,
	0xff, 0x15, 0x8c, 0xb3, 0xcf, 0x74, 0x31, 0xc5, 0x68, 0x08, 0x9e, 0x60, 0x74, 0xde, 0x8c, 0x27,
	0x9c, 0xf8, 0x76, 0x60, 0x84, 0x0d, 0x4d, 0x76, 0x31, 0xce, 0xee, 0xa5, 0x79, 0x2f, 0xbc, 0xeb,
	0x3b, 0x70, 0xce, 0x9f, 0x11, 0x01, 0x34, 0xa7, 0x72, 0x0d, 0xbd, 0x9a, 0xf8, 0xfe, 0x2a, 0xd7,
	0xcd, 0x33, 0x90, 0x03, 0xf6, 0x44, 0x6e, 0xc5, 0x74, 0xe2, 0x99, 0xa7, 0x6a, 0x76, 0xe3, 0x59,
	0xc8, 0x85, 0xb6, 0xaa, 0x22, 0x5a, 0x7a, 0xf5, 0xeb, 0x5b, 0x80, 0x53, 0xd2, 0x02, 0xad, 0xb6,
	0x41, 0xb5, 0xfc, 0x24, 0x53, 0x53, 0x2d, 0x23, 0x9a, 0x65, 0x8b, 0x64, 0xb9, 0xf6, 0x4c, 0xe1,
	0xcc, 0x0a, 0x32, 0x99, 0x7c, 0xf1, 0x2c, 0xd4, 0x81, 0xd6, 0x8c, 0x32, 0x2e, 0x8a, 0xfa, 0xd8,
	0x7b, 0x3e, 0xf4, 0x8c, 0xdf, 0x87, 0x9e, 0xf1, 0xf7, 0xd0, 0x33, 0x7e, 0xfe, 0xeb, 0xd5, 0xfe,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x38, 0xc6, 0x04, 0x2b, 0x9f, 0x03, 0x00, 0x00,
}

func (m *TableMutation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableMutation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TableMutation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Sequence) > 0 {
		for iNdEx := len(m.Sequence) - 1; iNdEx >= 0; iNdEx-- {
			i = encodeVarintBinlog(dAtA, i, uint64(m.Sequence[iNdEx]))
			i--
			dAtA[i] = 0x38
		}
	}
	if len(m.DeletedRows) > 0 {
		for iNdEx := len(m.DeletedRows) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DeletedRows[iNdEx])
			copy(dAtA[i:], m.DeletedRows[iNdEx])
			i = encodeVarintBinlog(dAtA, i, uint64(len(m.DeletedRows[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.DeletedPks) > 0 {
		for iNdEx := len(m.DeletedPks) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DeletedPks[iNdEx])
			copy(dAtA[i:], m.DeletedPks[iNdEx])
			i = encodeVarintBinlog(dAtA, i, uint64(len(m.DeletedPks[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.DeletedIds) > 0 {
		for iNdEx := len(m.DeletedIds) - 1; iNdEx >= 0; iNdEx-- {
			i = encodeVarintBinlog(dAtA, i, uint64(m.DeletedIds[iNdEx]))
			i--
			dAtA[i] = 0x20
		}
	}
	if len(m.UpdatedRows) > 0 {
		for iNdEx := len(m.UpdatedRows) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.UpdatedRows[iNdEx])
			copy(dAtA[i:], m.UpdatedRows[iNdEx])
			i = encodeVarintBinlog(dAtA, i, uint64(len(m.UpdatedRows[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.InsertedRows) > 0 {
		for iNdEx := len(m.InsertedRows) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.InsertedRows[iNdEx])
			copy(dAtA[i:], m.InsertedRows[iNdEx])
			i = encodeVarintBinlog(dAtA, i, uint64(len(m.InsertedRows[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	i = encodeVarintBinlog(dAtA, i, uint64(m.TableId))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *PrewriteValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrewriteValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrewriteValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Mutations) > 0 {
		for iNdEx := len(m.Mutations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Mutations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBinlog(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	i = encodeVarintBinlog(dAtA, i, uint64(m.SchemaVersion))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *Binlog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Binlog) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Binlog) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	i = encodeVarintBinlog(dAtA, i, uint64(m.DdlSchemaState))
	i--
	dAtA[i] = 0x40
	i = encodeVarintBinlog(dAtA, i, uint64(m.DdlJobId))
	i--
	dAtA[i] = 0x38
	if m.DdlQuery != nil {
		i -= len(m.DdlQuery)
		copy(dAtA[i:], m.DdlQuery)
		i = encodeVarintBinlog(dAtA, i, uint64(len(m.DdlQuery)))
		i--
		dAtA[i] = 0x32
	}
	if m.PrewriteValue != nil {
		i -= len(m.PrewriteValue)
		copy(dAtA[i:], m.PrewriteValue)
		i = encodeVarintBinlog(dAtA, i, uint64(len(m.PrewriteValue)))
		i--
		dAtA[i] = 0x2a
	}
	if m.PrewriteKey != nil {
		i -= len(m.PrewriteKey)
		copy(dAtA[i:], m.PrewriteKey)
		i = encodeVarintBinlog(dAtA, i, uint64(len(m.PrewriteKey)))
		i--
		dAtA[i] = 0x22
	}
	i = encodeVarintBinlog(dAtA, i, uint64(m.CommitTs))
	i--
	dAtA[i] = 0x18
	i = encodeVarintBinlog(dAtA, i, uint64(m.StartTs))
	i--
	dAtA[i] = 0x10
	i = encodeVarintBinlog(dAtA, i, uint64(m.Tp))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func encodeVarintBinlog(dAtA []byte, offset int, v uint64) int {
	offset -= sovBinlog(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TableMutation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovBinlog(uint64(m.TableId))
	if len(m.InsertedRows) > 0 {
		for _, b := range m.InsertedRows {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if len(m.UpdatedRows) > 0 {
		for _, b := range m.UpdatedRows {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if len(m.DeletedIds) > 0 {
		for _, e := range m.DeletedIds {
			n += 1 + sovBinlog(uint64(e))
		}
	}
	if len(m.DeletedPks) > 0 {
		for _, b := range m.DeletedPks {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if len(m.DeletedRows) > 0 {
		for _, b := range m.DeletedRows {
			l = len(b)
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if len(m.Sequence) > 0 {
		for _, e := range m.Sequence {
			n += 1 + sovBinlog(uint64(e))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PrewriteValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovBinlog(uint64(m.SchemaVersion))
	if len(m.Mutations) > 0 {
		for _, e := range m.Mutations {
			l = e.Size()
			n += 1 + l + sovBinlog(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Binlog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovBinlog(uint64(m.Tp))
	n += 1 + sovBinlog(uint64(m.StartTs))
	n += 1 + sovBinlog(uint64(m.CommitTs))
	if m.PrewriteKey != nil {
		l = len(m.PrewriteKey)
		n += 1 + l + sovBinlog(uint64(l))
	}
	if m.PrewriteValue != nil {
		l = len(m.PrewriteValue)
		n += 1 + l + sovBinlog(uint64(l))
	}
	if m.DdlQuery != nil {
		l = len(m.DdlQuery)
		n += 1 + l + sovBinlog(uint64(l))
	}
	n += 1 + sovBinlog(uint64(m.DdlJobId))
	n += 1 + sovBinlog(uint64(m.DdlSchemaState))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBinlog(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBinlog(x uint64) (n int) {
	return sovBinlog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TableMutation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinlog
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
			return fmt.Errorf("proto: TableMutation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableMutation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableId", wireType)
			}
			m.TableId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
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
				return fmt.Errorf("proto: wrong wireType = %d for field InsertedRows", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
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
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBinlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InsertedRows = append(m.InsertedRows, make([]byte, postIndex-iNdEx))
			copy(m.InsertedRows[len(m.InsertedRows)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedRows", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
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
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBinlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedRows = append(m.UpdatedRows, make([]byte, postIndex-iNdEx))
			copy(m.UpdatedRows[len(m.UpdatedRows)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBinlog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.DeletedIds = append(m.DeletedIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBinlog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBinlog
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthBinlog
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.DeletedIds) == 0 {
					m.DeletedIds = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBinlog
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.DeletedIds = append(m.DeletedIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedIds", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedPks", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
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
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBinlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeletedPks = append(m.DeletedPks, make([]byte, postIndex-iNdEx))
			copy(m.DeletedPks[len(m.DeletedPks)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedRows", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
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
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBinlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeletedRows = append(m.DeletedRows, make([]byte, postIndex-iNdEx))
			copy(m.DeletedRows[len(m.DeletedRows)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType == 0 {
				var v MutationType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBinlog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= MutationType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Sequence = append(m.Sequence, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBinlog
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBinlog
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthBinlog
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Sequence) == 0 {
					m.Sequence = make([]MutationType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v MutationType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBinlog
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= MutationType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Sequence = append(m.Sequence, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBinlog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBinlog
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
func (m *PrewriteValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinlog
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
			return fmt.Errorf("proto: PrewriteValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrewriteValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaVersion", wireType)
			}
			m.SchemaVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SchemaVersion |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mutations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
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
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBinlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mutations = append(m.Mutations, TableMutation{})
			if err := m.Mutations[len(m.Mutations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBinlog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBinlog
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
func (m *Binlog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBinlog
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
			return fmt.Errorf("proto: Binlog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Binlog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= BinlogType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTs", wireType)
			}
			m.StartTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitTs", wireType)
			}
			m.CommitTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitTs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrewriteKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
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
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBinlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrewriteKey = append(m.PrewriteKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PrewriteKey == nil {
				m.PrewriteKey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrewriteValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
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
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBinlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrewriteValue = append(m.PrewriteValue[:0], dAtA[iNdEx:postIndex]...)
			if m.PrewriteValue == nil {
				m.PrewriteValue = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DdlQuery", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
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
				return ErrInvalidLengthBinlog
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBinlog
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DdlQuery = append(m.DdlQuery[:0], dAtA[iNdEx:postIndex]...)
			if m.DdlQuery == nil {
				m.DdlQuery = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DdlJobId", wireType)
			}
			m.DdlJobId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DdlJobId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DdlSchemaState", wireType)
			}
			m.DdlSchemaState = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBinlog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DdlSchemaState |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBinlog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBinlog
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
func skipBinlog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBinlog
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
					return 0, ErrIntOverflowBinlog
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
					return 0, ErrIntOverflowBinlog
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
				return 0, ErrInvalidLengthBinlog
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBinlog
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBinlog
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBinlog        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBinlog          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBinlog = fmt.Errorf("proto: unexpected end of group")
)
