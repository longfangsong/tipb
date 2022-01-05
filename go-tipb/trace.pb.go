// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trace.proto

package tipb

import (
	fmt "fmt"
	math "math"

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

type Event int32

const (
	Event_Unknown                       Event = 0
	Event_TiKvCoprGetRequest            Event = 1000
	Event_TiKvCoprHandleRequest         Event = 1001
	Event_TiKvCoprScheduleTask          Event = 1002
	Event_TiKvCoprGetSnapshot           Event = 1003
	Event_TiKvCoprExecuteDagRunner      Event = 1004
	Event_TiKvCoprExecuteBatchDagRunner Event = 1005
)

var Event_name = map[int32]string{
	0:    "Unknown",
	1000: "TiKvCoprGetRequest",
	1001: "TiKvCoprHandleRequest",
	1002: "TiKvCoprScheduleTask",
	1003: "TiKvCoprGetSnapshot",
	1004: "TiKvCoprExecuteDagRunner",
	1005: "TiKvCoprExecuteBatchDagRunner",
}

var Event_value = map[string]int32{
	"Unknown":                       0,
	"TiKvCoprGetRequest":            1000,
	"TiKvCoprHandleRequest":         1001,
	"TiKvCoprScheduleTask":          1002,
	"TiKvCoprGetSnapshot":           1003,
	"TiKvCoprExecuteDagRunner":      1004,
	"TiKvCoprExecuteBatchDagRunner": 1005,
}

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}

func (x *Event) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Event_value, data, "Event")
	if err != nil {
		return err
	}
	*x = Event(value)
	return nil
}

func (Event) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0571941a1d628a80, []int{0}
}

func init() {
	proto.RegisterEnum("tipb.Event", Event_name, Event_value)
}

func init() { proto.RegisterFile("trace.proto", fileDescriptor_0571941a1d628a80) }

var fileDescriptor_0571941a1d628a80 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x29, 0x4a, 0x4c,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0xc9, 0x2c, 0x48, 0x92, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0x0b, 0xe8, 0x83, 0x58, 0x10, 0x39, 0xad, 0x83, 0x8c, 0x5c, 0xac, 0xae,
	0x65, 0xa9, 0x79, 0x25, 0x42, 0xdc, 0x5c, 0xec, 0xa1, 0x79, 0xd9, 0x79, 0xf9, 0xe5, 0x79, 0x02,
	0x0c, 0x42, 0xe2, 0x5c, 0x42, 0x21, 0x99, 0xde, 0x65, 0xce, 0xf9, 0x05, 0x45, 0xee, 0xa9, 0x25,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x02, 0x2f, 0xd8, 0x85, 0xa4, 0xb8, 0x44, 0x61, 0x12,
	0x1e, 0x89, 0x79, 0x29, 0x39, 0xa9, 0x30, 0xb9, 0x97, 0xec, 0x42, 0x92, 0x5c, 0x22, 0x30, 0xb9,
	0xe0, 0xe4, 0x8c, 0xd4, 0x94, 0xd2, 0x9c, 0xd4, 0x90, 0xc4, 0xe2, 0x6c, 0x81, 0x57, 0xec, 0x42,
	0x12, 0x5c, 0xc2, 0x48, 0xe6, 0x05, 0xe7, 0x25, 0x16, 0x14, 0x67, 0xe4, 0x97, 0x08, 0xbc, 0x66,
	0x17, 0x92, 0xe5, 0x92, 0x80, 0xc9, 0xb8, 0x56, 0xa4, 0x26, 0x97, 0x96, 0xa4, 0xba, 0x24, 0xa6,
	0x07, 0x95, 0xe6, 0xe5, 0xa5, 0x16, 0x09, 0xbc, 0x61, 0x17, 0x52, 0xe2, 0x92, 0x45, 0x93, 0x76,
	0x4a, 0x2c, 0x49, 0xce, 0x40, 0xa8, 0x79, 0xcb, 0xee, 0xa4, 0x79, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0xc0, 0x25, 0x9a, 0x9c, 0x9f,
	0xab, 0x57, 0x90, 0x99, 0x97, 0x9e, 0x9c, 0x58, 0xa0, 0x57, 0x92, 0x99, 0x92, 0xa4, 0x07, 0x0a,
	0x82, 0x00, 0x46, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0xee, 0x61, 0x2c, 0x18, 0x01, 0x00,
	0x00,
}
