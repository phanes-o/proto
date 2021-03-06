// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/base/primitive.proto

package base

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation errors at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{1}
}

func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (m *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(m, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Strings struct {
	Items                []string `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Strings) Reset()         { *m = Strings{} }
func (m *Strings) String() string { return proto.CompactTextString(m) }
func (*Strings) ProtoMessage()    {}
func (*Strings) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{2}
}

func (m *Strings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Strings.Unmarshal(m, b)
}
func (m *Strings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Strings.Marshal(b, m, deterministic)
}
func (m *Strings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Strings.Merge(m, src)
}
func (m *Strings) XXX_Size() int {
	return xxx_messageInfo_Strings.Size(m)
}
func (m *Strings) XXX_DiscardUnknown() {
	xxx_messageInfo_Strings.DiscardUnknown(m)
}

var xxx_messageInfo_Strings proto.InternalMessageInfo

func (m *Strings) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

type TimeSpanOptions struct {
	Valid                bool     `protobuf:"varint,1,opt,name=Valid,proto3" json:"Valid,omitempty"`
	Start                int64    `protobuf:"varint,2,opt,name=Start,proto3" json:"Start,omitempty"`
	End                  int64    `protobuf:"varint,3,opt,name=End,proto3" json:"End,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeSpanOptions) Reset()         { *m = TimeSpanOptions{} }
func (m *TimeSpanOptions) String() string { return proto.CompactTextString(m) }
func (*TimeSpanOptions) ProtoMessage()    {}
func (*TimeSpanOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{3}
}

func (m *TimeSpanOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSpanOptions.Unmarshal(m, b)
}
func (m *TimeSpanOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSpanOptions.Marshal(b, m, deterministic)
}
func (m *TimeSpanOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSpanOptions.Merge(m, src)
}
func (m *TimeSpanOptions) XXX_Size() int {
	return xxx_messageInfo_TimeSpanOptions.Size(m)
}
func (m *TimeSpanOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSpanOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSpanOptions proto.InternalMessageInfo

func (m *TimeSpanOptions) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *TimeSpanOptions) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *TimeSpanOptions) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type StringOption struct {
	Value                string   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringOption) Reset()         { *m = StringOption{} }
func (m *StringOption) String() string { return proto.CompactTextString(m) }
func (*StringOption) ProtoMessage()    {}
func (*StringOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{4}
}

func (m *StringOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringOption.Unmarshal(m, b)
}
func (m *StringOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringOption.Marshal(b, m, deterministic)
}
func (m *StringOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringOption.Merge(m, src)
}
func (m *StringOption) XXX_Size() int {
	return xxx_messageInfo_StringOption.Size(m)
}
func (m *StringOption) XXX_DiscardUnknown() {
	xxx_messageInfo_StringOption.DiscardUnknown(m)
}

var xxx_messageInfo_StringOption proto.InternalMessageInfo

func (m *StringOption) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *StringOption) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type StringsOption struct {
	Items                []string `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringsOption) Reset()         { *m = StringsOption{} }
func (m *StringsOption) String() string { return proto.CompactTextString(m) }
func (*StringsOption) ProtoMessage()    {}
func (*StringsOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{5}
}

func (m *StringsOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringsOption.Unmarshal(m, b)
}
func (m *StringsOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringsOption.Marshal(b, m, deterministic)
}
func (m *StringsOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringsOption.Merge(m, src)
}
func (m *StringsOption) XXX_Size() int {
	return xxx_messageInfo_StringsOption.Size(m)
}
func (m *StringsOption) XXX_DiscardUnknown() {
	xxx_messageInfo_StringsOption.DiscardUnknown(m)
}

var xxx_messageInfo_StringsOption proto.InternalMessageInfo

func (m *StringsOption) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *StringsOption) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type Int32 struct {
	Value                int32    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32) Reset()         { *m = Int32{} }
func (m *Int32) String() string { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()    {}
func (*Int32) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{6}
}

func (m *Int32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32.Unmarshal(m, b)
}
func (m *Int32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32.Marshal(b, m, deterministic)
}
func (m *Int32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32.Merge(m, src)
}
func (m *Int32) XXX_Size() int {
	return xxx_messageInfo_Int32.Size(m)
}
func (m *Int32) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32.DiscardUnknown(m)
}

var xxx_messageInfo_Int32 proto.InternalMessageInfo

func (m *Int32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int32S struct {
	Items                []int32  `protobuf:"varint,1,rep,packed,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32S) Reset()         { *m = Int32S{} }
func (m *Int32S) String() string { return proto.CompactTextString(m) }
func (*Int32S) ProtoMessage()    {}
func (*Int32S) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{7}
}

func (m *Int32S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32S.Unmarshal(m, b)
}
func (m *Int32S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32S.Marshal(b, m, deterministic)
}
func (m *Int32S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32S.Merge(m, src)
}
func (m *Int32S) XXX_Size() int {
	return xxx_messageInfo_Int32S.Size(m)
}
func (m *Int32S) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32S.DiscardUnknown(m)
}

var xxx_messageInfo_Int32S proto.InternalMessageInfo

func (m *Int32S) GetItems() []int32 {
	if m != nil {
		return m.Items
	}
	return nil
}

type Int32Option struct {
	Value                int32    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32Option) Reset()         { *m = Int32Option{} }
func (m *Int32Option) String() string { return proto.CompactTextString(m) }
func (*Int32Option) ProtoMessage()    {}
func (*Int32Option) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{8}
}

func (m *Int32Option) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32Option.Unmarshal(m, b)
}
func (m *Int32Option) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32Option.Marshal(b, m, deterministic)
}
func (m *Int32Option) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32Option.Merge(m, src)
}
func (m *Int32Option) XXX_Size() int {
	return xxx_messageInfo_Int32Option.Size(m)
}
func (m *Int32Option) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32Option.DiscardUnknown(m)
}

var xxx_messageInfo_Int32Option proto.InternalMessageInfo

func (m *Int32Option) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Int32Option) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type Int64 struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64) Reset()         { *m = Int64{} }
func (m *Int64) String() string { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()    {}
func (*Int64) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{9}
}

func (m *Int64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64.Unmarshal(m, b)
}
func (m *Int64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64.Marshal(b, m, deterministic)
}
func (m *Int64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64.Merge(m, src)
}
func (m *Int64) XXX_Size() int {
	return xxx_messageInfo_Int64.Size(m)
}
func (m *Int64) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64.DiscardUnknown(m)
}

var xxx_messageInfo_Int64 proto.InternalMessageInfo

func (m *Int64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int64S struct {
	Items                []int64  `protobuf:"varint,1,rep,packed,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64S) Reset()         { *m = Int64S{} }
func (m *Int64S) String() string { return proto.CompactTextString(m) }
func (*Int64S) ProtoMessage()    {}
func (*Int64S) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{10}
}

func (m *Int64S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64S.Unmarshal(m, b)
}
func (m *Int64S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64S.Marshal(b, m, deterministic)
}
func (m *Int64S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64S.Merge(m, src)
}
func (m *Int64S) XXX_Size() int {
	return xxx_messageInfo_Int64S.Size(m)
}
func (m *Int64S) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64S.DiscardUnknown(m)
}

var xxx_messageInfo_Int64S proto.InternalMessageInfo

func (m *Int64S) GetItems() []int64 {
	if m != nil {
		return m.Items
	}
	return nil
}

type Int64Option struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Option) Reset()         { *m = Int64Option{} }
func (m *Int64Option) String() string { return proto.CompactTextString(m) }
func (*Int64Option) ProtoMessage()    {}
func (*Int64Option) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{11}
}

func (m *Int64Option) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64Option.Unmarshal(m, b)
}
func (m *Int64Option) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64Option.Marshal(b, m, deterministic)
}
func (m *Int64Option) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Option.Merge(m, src)
}
func (m *Int64Option) XXX_Size() int {
	return xxx_messageInfo_Int64Option.Size(m)
}
func (m *Int64Option) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Option.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Option proto.InternalMessageInfo

func (m *Int64Option) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Int64Option) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type Float struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Float) Reset()         { *m = Float{} }
func (m *Float) String() string { return proto.CompactTextString(m) }
func (*Float) ProtoMessage()    {}
func (*Float) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{12}
}

func (m *Float) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Float.Unmarshal(m, b)
}
func (m *Float) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Float.Marshal(b, m, deterministic)
}
func (m *Float) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Float.Merge(m, src)
}
func (m *Float) XXX_Size() int {
	return xxx_messageInfo_Float.Size(m)
}
func (m *Float) XXX_DiscardUnknown() {
	xxx_messageInfo_Float.DiscardUnknown(m)
}

var xxx_messageInfo_Float proto.InternalMessageInfo

func (m *Float) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Floats struct {
	Items                []float32 `protobuf:"fixed32,1,rep,packed,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Floats) Reset()         { *m = Floats{} }
func (m *Floats) String() string { return proto.CompactTextString(m) }
func (*Floats) ProtoMessage()    {}
func (*Floats) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{13}
}

func (m *Floats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Floats.Unmarshal(m, b)
}
func (m *Floats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Floats.Marshal(b, m, deterministic)
}
func (m *Floats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Floats.Merge(m, src)
}
func (m *Floats) XXX_Size() int {
	return xxx_messageInfo_Floats.Size(m)
}
func (m *Floats) XXX_DiscardUnknown() {
	xxx_messageInfo_Floats.DiscardUnknown(m)
}

var xxx_messageInfo_Floats proto.InternalMessageInfo

func (m *Floats) GetItems() []float32 {
	if m != nil {
		return m.Items
	}
	return nil
}

type FloatOption struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FloatOption) Reset()         { *m = FloatOption{} }
func (m *FloatOption) String() string { return proto.CompactTextString(m) }
func (*FloatOption) ProtoMessage()    {}
func (*FloatOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{14}
}

func (m *FloatOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatOption.Unmarshal(m, b)
}
func (m *FloatOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatOption.Marshal(b, m, deterministic)
}
func (m *FloatOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatOption.Merge(m, src)
}
func (m *FloatOption) XXX_Size() int {
	return xxx_messageInfo_FloatOption.Size(m)
}
func (m *FloatOption) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatOption.DiscardUnknown(m)
}

var xxx_messageInfo_FloatOption proto.InternalMessageInfo

func (m *FloatOption) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *FloatOption) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type Double struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Double) Reset()         { *m = Double{} }
func (m *Double) String() string { return proto.CompactTextString(m) }
func (*Double) ProtoMessage()    {}
func (*Double) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{15}
}

func (m *Double) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Double.Unmarshal(m, b)
}
func (m *Double) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Double.Marshal(b, m, deterministic)
}
func (m *Double) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Double.Merge(m, src)
}
func (m *Double) XXX_Size() int {
	return xxx_messageInfo_Double.Size(m)
}
func (m *Double) XXX_DiscardUnknown() {
	xxx_messageInfo_Double.DiscardUnknown(m)
}

var xxx_messageInfo_Double proto.InternalMessageInfo

func (m *Double) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Doubles struct {
	Items                []float64 `protobuf:"fixed64,1,rep,packed,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Doubles) Reset()         { *m = Doubles{} }
func (m *Doubles) String() string { return proto.CompactTextString(m) }
func (*Doubles) ProtoMessage()    {}
func (*Doubles) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{16}
}

func (m *Doubles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Doubles.Unmarshal(m, b)
}
func (m *Doubles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Doubles.Marshal(b, m, deterministic)
}
func (m *Doubles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Doubles.Merge(m, src)
}
func (m *Doubles) XXX_Size() int {
	return xxx_messageInfo_Doubles.Size(m)
}
func (m *Doubles) XXX_DiscardUnknown() {
	xxx_messageInfo_Doubles.DiscardUnknown(m)
}

var xxx_messageInfo_Doubles proto.InternalMessageInfo

func (m *Doubles) GetItems() []float64 {
	if m != nil {
		return m.Items
	}
	return nil
}

type DoubleOption struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoubleOption) Reset()         { *m = DoubleOption{} }
func (m *DoubleOption) String() string { return proto.CompactTextString(m) }
func (*DoubleOption) ProtoMessage()    {}
func (*DoubleOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{17}
}

func (m *DoubleOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleOption.Unmarshal(m, b)
}
func (m *DoubleOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleOption.Marshal(b, m, deterministic)
}
func (m *DoubleOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleOption.Merge(m, src)
}
func (m *DoubleOption) XXX_Size() int {
	return xxx_messageInfo_DoubleOption.Size(m)
}
func (m *DoubleOption) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleOption.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleOption proto.InternalMessageInfo

func (m *DoubleOption) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *DoubleOption) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type Bool struct {
	Value                bool     `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bool) Reset()         { *m = Bool{} }
func (m *Bool) String() string { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()    {}
func (*Bool) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{18}
}

func (m *Bool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bool.Unmarshal(m, b)
}
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bool.Marshal(b, m, deterministic)
}
func (m *Bool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bool.Merge(m, src)
}
func (m *Bool) XXX_Size() int {
	return xxx_messageInfo_Bool.Size(m)
}
func (m *Bool) XXX_DiscardUnknown() {
	xxx_messageInfo_Bool.DiscardUnknown(m)
}

var xxx_messageInfo_Bool proto.InternalMessageInfo

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type Bools struct {
	Items                []bool   `protobuf:"varint,1,rep,packed,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bools) Reset()         { *m = Bools{} }
func (m *Bools) String() string { return proto.CompactTextString(m) }
func (*Bools) ProtoMessage()    {}
func (*Bools) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{19}
}

func (m *Bools) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bools.Unmarshal(m, b)
}
func (m *Bools) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bools.Marshal(b, m, deterministic)
}
func (m *Bools) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bools.Merge(m, src)
}
func (m *Bools) XXX_Size() int {
	return xxx_messageInfo_Bools.Size(m)
}
func (m *Bools) XXX_DiscardUnknown() {
	xxx_messageInfo_Bools.DiscardUnknown(m)
}

var xxx_messageInfo_Bools proto.InternalMessageInfo

func (m *Bools) GetItems() []bool {
	if m != nil {
		return m.Items
	}
	return nil
}

type BoolOption struct {
	Value                bool     `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolOption) Reset()         { *m = BoolOption{} }
func (m *BoolOption) String() string { return proto.CompactTextString(m) }
func (*BoolOption) ProtoMessage()    {}
func (*BoolOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{20}
}

func (m *BoolOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolOption.Unmarshal(m, b)
}
func (m *BoolOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolOption.Marshal(b, m, deterministic)
}
func (m *BoolOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolOption.Merge(m, src)
}
func (m *BoolOption) XXX_Size() int {
	return xxx_messageInfo_BoolOption.Size(m)
}
func (m *BoolOption) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolOption.DiscardUnknown(m)
}

var xxx_messageInfo_BoolOption proto.InternalMessageInfo

func (m *BoolOption) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func (m *BoolOption) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type Bytes struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bytes) Reset()         { *m = Bytes{} }
func (m *Bytes) String() string { return proto.CompactTextString(m) }
func (*Bytes) ProtoMessage()    {}
func (*Bytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{21}
}

func (m *Bytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bytes.Unmarshal(m, b)
}
func (m *Bytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bytes.Marshal(b, m, deterministic)
}
func (m *Bytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bytes.Merge(m, src)
}
func (m *Bytes) XXX_Size() int {
	return xxx_messageInfo_Bytes.Size(m)
}
func (m *Bytes) XXX_DiscardUnknown() {
	xxx_messageInfo_Bytes.DiscardUnknown(m)
}

var xxx_messageInfo_Bytes proto.InternalMessageInfo

func (m *Bytes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Byteses struct {
	Items                [][]byte `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Byteses) Reset()         { *m = Byteses{} }
func (m *Byteses) String() string { return proto.CompactTextString(m) }
func (*Byteses) ProtoMessage()    {}
func (*Byteses) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{22}
}

func (m *Byteses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Byteses.Unmarshal(m, b)
}
func (m *Byteses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Byteses.Marshal(b, m, deterministic)
}
func (m *Byteses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Byteses.Merge(m, src)
}
func (m *Byteses) XXX_Size() int {
	return xxx_messageInfo_Byteses.Size(m)
}
func (m *Byteses) XXX_DiscardUnknown() {
	xxx_messageInfo_Byteses.DiscardUnknown(m)
}

var xxx_messageInfo_Byteses proto.InternalMessageInfo

func (m *Byteses) GetItems() [][]byte {
	if m != nil {
		return m.Items
	}
	return nil
}

type BytesOption struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesOption) Reset()         { *m = BytesOption{} }
func (m *BytesOption) String() string { return proto.CompactTextString(m) }
func (*BytesOption) ProtoMessage()    {}
func (*BytesOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{23}
}

func (m *BytesOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesOption.Unmarshal(m, b)
}
func (m *BytesOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesOption.Marshal(b, m, deterministic)
}
func (m *BytesOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesOption.Merge(m, src)
}
func (m *BytesOption) XXX_Size() int {
	return xxx_messageInfo_BytesOption.Size(m)
}
func (m *BytesOption) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesOption.DiscardUnknown(m)
}

var xxx_messageInfo_BytesOption proto.InternalMessageInfo

func (m *BytesOption) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *BytesOption) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

type KVStringAtString struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVStringAtString) Reset()         { *m = KVStringAtString{} }
func (m *KVStringAtString) String() string { return proto.CompactTextString(m) }
func (*KVStringAtString) ProtoMessage()    {}
func (*KVStringAtString) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{24}
}

func (m *KVStringAtString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVStringAtString.Unmarshal(m, b)
}
func (m *KVStringAtString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVStringAtString.Marshal(b, m, deterministic)
}
func (m *KVStringAtString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVStringAtString.Merge(m, src)
}
func (m *KVStringAtString) XXX_Size() int {
	return xxx_messageInfo_KVStringAtString.Size(m)
}
func (m *KVStringAtString) XXX_DiscardUnknown() {
	xxx_messageInfo_KVStringAtString.DiscardUnknown(m)
}

var xxx_messageInfo_KVStringAtString proto.InternalMessageInfo

func (m *KVStringAtString) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVStringAtString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KVInt64AtString struct {
	Key                  int64    `protobuf:"varint,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVInt64AtString) Reset()         { *m = KVInt64AtString{} }
func (m *KVInt64AtString) String() string { return proto.CompactTextString(m) }
func (*KVInt64AtString) ProtoMessage()    {}
func (*KVInt64AtString) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{25}
}

func (m *KVInt64AtString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVInt64AtString.Unmarshal(m, b)
}
func (m *KVInt64AtString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVInt64AtString.Marshal(b, m, deterministic)
}
func (m *KVInt64AtString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVInt64AtString.Merge(m, src)
}
func (m *KVInt64AtString) XXX_Size() int {
	return xxx_messageInfo_KVInt64AtString.Size(m)
}
func (m *KVInt64AtString) XXX_DiscardUnknown() {
	xxx_messageInfo_KVInt64AtString.DiscardUnknown(m)
}

var xxx_messageInfo_KVInt64AtString proto.InternalMessageInfo

func (m *KVInt64AtString) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *KVInt64AtString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KVInt64AtInt64S struct {
	Key                  int64    `protobuf:"varint,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Values               []int64  `protobuf:"varint,2,rep,packed,name=Values,proto3" json:"Values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVInt64AtInt64S) Reset()         { *m = KVInt64AtInt64S{} }
func (m *KVInt64AtInt64S) String() string { return proto.CompactTextString(m) }
func (*KVInt64AtInt64S) ProtoMessage()    {}
func (*KVInt64AtInt64S) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6aab5286ac87312, []int{26}
}

func (m *KVInt64AtInt64S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVInt64AtInt64S.Unmarshal(m, b)
}
func (m *KVInt64AtInt64S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVInt64AtInt64S.Marshal(b, m, deterministic)
}
func (m *KVInt64AtInt64S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVInt64AtInt64S.Merge(m, src)
}
func (m *KVInt64AtInt64S) XXX_Size() int {
	return xxx_messageInfo_KVInt64AtInt64S.Size(m)
}
func (m *KVInt64AtInt64S) XXX_DiscardUnknown() {
	xxx_messageInfo_KVInt64AtInt64S.DiscardUnknown(m)
}

var xxx_messageInfo_KVInt64AtInt64S proto.InternalMessageInfo

func (m *KVInt64AtInt64S) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *KVInt64AtInt64S) GetValues() []int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "base.Empty")
	proto.RegisterType((*String)(nil), "base.String")
	proto.RegisterType((*Strings)(nil), "base.Strings")
	proto.RegisterType((*TimeSpanOptions)(nil), "base.TimeSpanOptions")
	proto.RegisterType((*StringOption)(nil), "base.StringOption")
	proto.RegisterType((*StringsOption)(nil), "base.StringsOption")
	proto.RegisterType((*Int32)(nil), "base.Int32")
	proto.RegisterType((*Int32S)(nil), "base.Int32s")
	proto.RegisterType((*Int32Option)(nil), "base.Int32Option")
	proto.RegisterType((*Int64)(nil), "base.Int64")
	proto.RegisterType((*Int64S)(nil), "base.Int64s")
	proto.RegisterType((*Int64Option)(nil), "base.Int64Option")
	proto.RegisterType((*Float)(nil), "base.Float")
	proto.RegisterType((*Floats)(nil), "base.Floats")
	proto.RegisterType((*FloatOption)(nil), "base.FloatOption")
	proto.RegisterType((*Double)(nil), "base.Double")
	proto.RegisterType((*Doubles)(nil), "base.Doubles")
	proto.RegisterType((*DoubleOption)(nil), "base.DoubleOption")
	proto.RegisterType((*Bool)(nil), "base.Bool")
	proto.RegisterType((*Bools)(nil), "base.Bools")
	proto.RegisterType((*BoolOption)(nil), "base.BoolOption")
	proto.RegisterType((*Bytes)(nil), "base.Bytes")
	proto.RegisterType((*Byteses)(nil), "base.Byteses")
	proto.RegisterType((*BytesOption)(nil), "base.BytesOption")
	proto.RegisterType((*KVStringAtString)(nil), "base.KVStringAtString")
	proto.RegisterType((*KVInt64AtString)(nil), "base.KVInt64AtString")
	proto.RegisterType((*KVInt64AtInt64S)(nil), "base.KVInt64AtInt64s")
}

func init() { proto.RegisterFile("proto/base/primitive.proto", fileDescriptor_c6aab5286ac87312) }

var fileDescriptor_c6aab5286ac87312 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xc1, 0x8b, 0x9b, 0x40,
	0x14, 0xc6, 0xd1, 0x59, 0x5d, 0xfb, 0x6a, 0xd9, 0x10, 0x4a, 0x91, 0xd2, 0xb5, 0xe2, 0x69, 0x2f,
	0x5d, 0xa1, 0x1b, 0x42, 0x37, 0x7b, 0x6a, 0xe8, 0x16, 0x96, 0x1c, 0x16, 0x92, 0x92, 0x43, 0x6f,
	0xda, 0x0c, 0xc9, 0x80, 0x3a, 0xa2, 0x93, 0x42, 0xfe, 0xfb, 0x32, 0x33, 0x8a, 0x33, 0x66, 0x6a,
	0x6e, 0xf3, 0xe6, 0x7b, 0xdf, 0xfb, 0x3d, 0x1f, 0x6f, 0x84, 0x8f, 0x55, 0x4d, 0x19, 0x4d, 0xb2,
	0xb4, 0xc1, 0x49, 0x55, 0x93, 0x82, 0x30, 0xf2, 0x17, 0xdf, 0x8b, 0xcb, 0xe9, 0x15, 0xbf, 0x8d,
	0xaf, 0xc1, 0x79, 0x2e, 0x2a, 0x76, 0x8a, 0x43, 0x70, 0x37, 0xac, 0x26, 0xe5, 0x7e, 0xfa, 0x1e,
	0x9c, 0x6d, 0x9a, 0x1f, 0x71, 0x60, 0x45, 0xd6, 0xdd, 0x9b, 0xb5, 0x0c, 0xe2, 0xcf, 0x70, 0x2d,
	0xf5, 0x86, 0x27, 0xbc, 0x30, 0x5c, 0x34, 0x81, 0x15, 0x21, 0x9e, 0x20, 0x82, 0xf8, 0x15, 0x6e,
	0x7e, 0x91, 0x02, 0x6f, 0xaa, 0xb4, 0x7c, 0xad, 0x18, 0xa1, 0x65, 0xd3, 0x56, 0x22, 0x3b, 0x51,
	0xc9, 0x5b, 0xcb, 0x80, 0xdf, 0x6e, 0x58, 0x5a, 0xb3, 0xc0, 0x8e, 0xac, 0x3b, 0xb4, 0x96, 0xc1,
	0x74, 0x02, 0xe8, 0xb9, 0xdc, 0x05, 0x48, 0xdc, 0xf1, 0x63, 0xbc, 0x00, 0x5f, 0x12, 0x65, 0x39,
	0x73, 0x5f, 0x3d, 0xc3, 0x56, 0x18, 0xf1, 0x13, 0xbc, 0x6b, 0xbb, 0xed, 0xcd, 0xe7, 0x3d, 0xff,
	0xc7, 0x7c, 0x0b, 0xce, 0x4b, 0xc9, 0x1e, 0xbe, 0xea, 0x44, 0xa7, 0x9b, 0x44, 0x08, 0xae, 0x90,
	0x07, 0x83, 0x70, 0xba, 0x41, 0x3c, 0xc2, 0x5b, 0xa1, 0x9b, 0xda, 0x76, 0xc6, 0xdb, 0x96, 0xe4,
	0xf9, 0x4c, 0x37, 0x21, 0x9d, 0x3c, 0x9f, 0x0d, 0xc8, 0x48, 0x27, 0xcf, 0x67, 0x26, 0x32, 0xba,
	0x48, 0xfe, 0x99, 0xd3, 0x94, 0xe9, 0x26, 0x5b, 0x21, 0x0b, 0x79, 0x40, 0xb6, 0x15, 0xb2, 0xd0,
	0x4d, 0x64, 0x7b, 0x9c, 0x1c, 0x82, 0xfb, 0x83, 0x1e, 0xb3, 0x1c, 0xeb, 0x2e, 0x4b, 0x59, 0x3c,
	0xa9, 0x0f, 0xd8, 0x56, 0xc7, 0x5e, 0x80, 0x2f, 0x13, 0x4c, 0x70, 0x6b, 0x1c, 0xfe, 0x09, 0xae,
	0x96, 0x94, 0xe6, 0xba, 0xc7, 0xeb, 0xd0, 0xb7, 0xe0, 0x70, 0x75, 0x00, 0xf6, 0x3a, 0xf0, 0x37,
	0x00, 0x2e, 0x9b, 0xb0, 0xde, 0xc5, 0x69, 0x2f, 0x4f, 0x0c, 0x37, 0xba, 0xc9, 0x57, 0x3e, 0x59,
	0xc8, 0xc3, 0x4f, 0xf6, 0x95, 0x71, 0x8b, 0x04, 0x13, 0xda, 0x1f, 0x47, 0x2f, 0x60, 0xb2, 0xda,
	0xca, 0xb7, 0xf1, 0x9d, 0xb5, 0x2f, 0x7e, 0x02, 0x68, 0x85, 0x4f, 0xed, 0xbb, 0xe2, 0xc7, 0xbe,
	0xa2, 0xad, 0xfe, 0x03, 0x1e, 0xe1, 0x66, 0xb5, 0x15, 0x1b, 0x66, 0xb2, 0xa2, 0x31, 0xeb, 0x93,
	0x62, 0x6d, 0x77, 0xf8, 0xdc, 0xfa, 0x01, 0x5c, 0x91, 0xdd, 0x04, 0xb6, 0x58, 0xeb, 0x36, 0x5a,
	0x46, 0xbf, 0xc3, 0x3d, 0x61, 0x87, 0x63, 0x76, 0xff, 0x87, 0x16, 0x49, 0x75, 0x48, 0x4b, 0xdc,
	0x7c, 0xa1, 0x49, 0xff, 0x73, 0xcb, 0x5c, 0x71, 0x7e, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0x60,
	0x76, 0xea, 0xfb, 0xf1, 0x04, 0x00, 0x00,
}
