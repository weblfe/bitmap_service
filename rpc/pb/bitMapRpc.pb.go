// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bitMapRpc.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

// 留存类型(0:日期为截止日,1:日期为开始日)
type RpcSingleRequest_DataType int32

const (
	RpcSingleRequest_END   RpcSingleRequest_DataType = 0
	RpcSingleRequest_START RpcSingleRequest_DataType = 1
)

var RpcSingleRequest_DataType_name = map[int32]string{
	0: "END",
	1: "START",
}

var RpcSingleRequest_DataType_value = map[string]int32{
	"END":   0,
	"START": 1,
}

func (x RpcSingleRequest_DataType) String() string {
	return proto.EnumName(RpcSingleRequest_DataType_name, int32(x))
}

func (RpcSingleRequest_DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{0, 0}
}

// 留存类型(0:日期为截止日,1:日期为开始日)
type RpcMulitRequest_DataType int32

const (
	RpcMulitRequest_END   RpcMulitRequest_DataType = 0
	RpcMulitRequest_START RpcMulitRequest_DataType = 1
)

var RpcMulitRequest_DataType_name = map[int32]string{
	0: "END",
	1: "START",
}

var RpcMulitRequest_DataType_value = map[string]int32{
	"END":   0,
	"START": 1,
}

func (x RpcMulitRequest_DataType) String() string {
	return proto.EnumName(RpcMulitRequest_DataType_name, int32(x))
}

func (RpcMulitRequest_DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{1, 0}
}

// 留存类型(0:日期为截止日,1:日期为开始日)
type RpcMulitChannelsRequest_DataType int32

const (
	RpcMulitChannelsRequest_END   RpcMulitChannelsRequest_DataType = 0
	RpcMulitChannelsRequest_START RpcMulitChannelsRequest_DataType = 1
)

var RpcMulitChannelsRequest_DataType_name = map[int32]string{
	0: "END",
	1: "START",
}

var RpcMulitChannelsRequest_DataType_value = map[string]int32{
	"END":   0,
	"START": 1,
}

func (x RpcMulitChannelsRequest_DataType) String() string {
	return proto.EnumName(RpcMulitChannelsRequest_DataType_name, int32(x))
}

func (RpcMulitChannelsRequest_DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{2, 0}
}

type RpcSingleRequest struct {
	Date                 string                    `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Type                 RpcSingleRequest_DataType `protobuf:"varint,2,opt,name=type,proto3,enum=rpc.RpcSingleRequest_DataType" json:"type,omitempty"`
	Day                  int32                     `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	Channel              string                    `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	Role                 string                    `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	DataSource           string                    `protobuf:"bytes,6,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RpcSingleRequest) Reset()         { *m = RpcSingleRequest{} }
func (m *RpcSingleRequest) String() string { return proto.CompactTextString(m) }
func (*RpcSingleRequest) ProtoMessage()    {}
func (*RpcSingleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{0}
}

func (m *RpcSingleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcSingleRequest.Unmarshal(m, b)
}
func (m *RpcSingleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcSingleRequest.Marshal(b, m, deterministic)
}
func (m *RpcSingleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcSingleRequest.Merge(m, src)
}
func (m *RpcSingleRequest) XXX_Size() int {
	return xxx_messageInfo_RpcSingleRequest.Size(m)
}
func (m *RpcSingleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcSingleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RpcSingleRequest proto.InternalMessageInfo

func (m *RpcSingleRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *RpcSingleRequest) GetType() RpcSingleRequest_DataType {
	if m != nil {
		return m.Type
	}
	return RpcSingleRequest_END
}

func (m *RpcSingleRequest) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *RpcSingleRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *RpcSingleRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *RpcSingleRequest) GetDataSource() string {
	if m != nil {
		return m.DataSource
	}
	return ""
}

type RpcMulitRequest struct {
	Date                 string                   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Type                 RpcMulitRequest_DataType `protobuf:"varint,2,opt,name=type,proto3,enum=rpc.RpcMulitRequest_DataType" json:"type,omitempty"`
	Days                 []int32                  `protobuf:"varint,3,rep,packed,name=days,proto3" json:"days,omitempty"`
	Channel              string                   `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	Role                 string                   `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	DataSource           string                   `protobuf:"bytes,6,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RpcMulitRequest) Reset()         { *m = RpcMulitRequest{} }
func (m *RpcMulitRequest) String() string { return proto.CompactTextString(m) }
func (*RpcMulitRequest) ProtoMessage()    {}
func (*RpcMulitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{1}
}

func (m *RpcMulitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcMulitRequest.Unmarshal(m, b)
}
func (m *RpcMulitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcMulitRequest.Marshal(b, m, deterministic)
}
func (m *RpcMulitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcMulitRequest.Merge(m, src)
}
func (m *RpcMulitRequest) XXX_Size() int {
	return xxx_messageInfo_RpcMulitRequest.Size(m)
}
func (m *RpcMulitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcMulitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RpcMulitRequest proto.InternalMessageInfo

func (m *RpcMulitRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *RpcMulitRequest) GetType() RpcMulitRequest_DataType {
	if m != nil {
		return m.Type
	}
	return RpcMulitRequest_END
}

func (m *RpcMulitRequest) GetDays() []int32 {
	if m != nil {
		return m.Days
	}
	return nil
}

func (m *RpcMulitRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *RpcMulitRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *RpcMulitRequest) GetDataSource() string {
	if m != nil {
		return m.DataSource
	}
	return ""
}

type RpcMulitChannelsRequest struct {
	Date                 string                           `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Type                 RpcMulitChannelsRequest_DataType `protobuf:"varint,2,opt,name=type,proto3,enum=rpc.RpcMulitChannelsRequest_DataType" json:"type,omitempty"`
	Days                 []int32                          `protobuf:"varint,3,rep,packed,name=days,proto3" json:"days,omitempty"`
	Channels             []string                         `protobuf:"bytes,4,rep,name=channels,proto3" json:"channels,omitempty"`
	Role                 string                           `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	DataSource           string                           `protobuf:"bytes,6,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RpcMulitChannelsRequest) Reset()         { *m = RpcMulitChannelsRequest{} }
func (m *RpcMulitChannelsRequest) String() string { return proto.CompactTextString(m) }
func (*RpcMulitChannelsRequest) ProtoMessage()    {}
func (*RpcMulitChannelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{2}
}

func (m *RpcMulitChannelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcMulitChannelsRequest.Unmarshal(m, b)
}
func (m *RpcMulitChannelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcMulitChannelsRequest.Marshal(b, m, deterministic)
}
func (m *RpcMulitChannelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcMulitChannelsRequest.Merge(m, src)
}
func (m *RpcMulitChannelsRequest) XXX_Size() int {
	return xxx_messageInfo_RpcMulitChannelsRequest.Size(m)
}
func (m *RpcMulitChannelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcMulitChannelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RpcMulitChannelsRequest proto.InternalMessageInfo

func (m *RpcMulitChannelsRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *RpcMulitChannelsRequest) GetType() RpcMulitChannelsRequest_DataType {
	if m != nil {
		return m.Type
	}
	return RpcMulitChannelsRequest_END
}

func (m *RpcMulitChannelsRequest) GetDays() []int32 {
	if m != nil {
		return m.Days
	}
	return nil
}

func (m *RpcMulitChannelsRequest) GetChannels() []string {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *RpcMulitChannelsRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *RpcMulitChannelsRequest) GetDataSource() string {
	if m != nil {
		return m.DataSource
	}
	return ""
}

type RpcErrorResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	ErrorNo              int32    `protobuf:"varint,3,opt,name=error_no,json=errorNo,proto3" json:"error_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcErrorResponse) Reset()         { *m = RpcErrorResponse{} }
func (m *RpcErrorResponse) String() string { return proto.CompactTextString(m) }
func (*RpcErrorResponse) ProtoMessage()    {}
func (*RpcErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{3}
}

func (m *RpcErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcErrorResponse.Unmarshal(m, b)
}
func (m *RpcErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcErrorResponse.Marshal(b, m, deterministic)
}
func (m *RpcErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcErrorResponse.Merge(m, src)
}
func (m *RpcErrorResponse) XXX_Size() int {
	return xxx_messageInfo_RpcErrorResponse.Size(m)
}
func (m *RpcErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RpcErrorResponse proto.InternalMessageInfo

func (m *RpcErrorResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RpcErrorResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RpcErrorResponse) GetErrorNo() int32 {
	if m != nil {
		return m.ErrorNo
	}
	return 0
}

type KeeperStruct struct {
	Keep                 string   `protobuf:"bytes,1,opt,name=keep,proto3" json:"keep,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Register             int64    `protobuf:"varint,3,opt,name=register,proto3" json:"register,omitempty"`
	Left                 int64    `protobuf:"varint,4,opt,name=left,proto3" json:"left,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeeperStruct) Reset()         { *m = KeeperStruct{} }
func (m *KeeperStruct) String() string { return proto.CompactTextString(m) }
func (*KeeperStruct) ProtoMessage()    {}
func (*KeeperStruct) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{4}
}

func (m *KeeperStruct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeeperStruct.Unmarshal(m, b)
}
func (m *KeeperStruct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeeperStruct.Marshal(b, m, deterministic)
}
func (m *KeeperStruct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeeperStruct.Merge(m, src)
}
func (m *KeeperStruct) XXX_Size() int {
	return xxx_messageInfo_KeeperStruct.Size(m)
}
func (m *KeeperStruct) XXX_DiscardUnknown() {
	xxx_messageInfo_KeeperStruct.DiscardUnknown(m)
}

var xxx_messageInfo_KeeperStruct proto.InternalMessageInfo

func (m *KeeperStruct) GetKeep() string {
	if m != nil {
		return m.Keep
	}
	return ""
}

func (m *KeeperStruct) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KeeperStruct) GetRegister() int64 {
	if m != nil {
		return m.Register
	}
	return 0
}

func (m *KeeperStruct) GetLeft() int64 {
	if m != nil {
		return m.Left
	}
	return 0
}

type RpcChannelResponse struct {
	Channel              string          `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Items                []*KeeperStruct `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RpcChannelResponse) Reset()         { *m = RpcChannelResponse{} }
func (m *RpcChannelResponse) String() string { return proto.CompactTextString(m) }
func (*RpcChannelResponse) ProtoMessage()    {}
func (*RpcChannelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{5}
}

func (m *RpcChannelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcChannelResponse.Unmarshal(m, b)
}
func (m *RpcChannelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcChannelResponse.Marshal(b, m, deterministic)
}
func (m *RpcChannelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcChannelResponse.Merge(m, src)
}
func (m *RpcChannelResponse) XXX_Size() int {
	return xxx_messageInfo_RpcChannelResponse.Size(m)
}
func (m *RpcChannelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcChannelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RpcChannelResponse proto.InternalMessageInfo

func (m *RpcChannelResponse) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *RpcChannelResponse) GetItems() []*KeeperStruct {
	if m != nil {
		return m.Items
	}
	return nil
}

type RpcSingleResponse struct {
	Code                 int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	ErrorNo              int32         `protobuf:"varint,3,opt,name=error_no,json=errorNo,proto3" json:"error_no,omitempty"`
	Data                 *KeeperStruct `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RpcSingleResponse) Reset()         { *m = RpcSingleResponse{} }
func (m *RpcSingleResponse) String() string { return proto.CompactTextString(m) }
func (*RpcSingleResponse) ProtoMessage()    {}
func (*RpcSingleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{6}
}

func (m *RpcSingleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcSingleResponse.Unmarshal(m, b)
}
func (m *RpcSingleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcSingleResponse.Marshal(b, m, deterministic)
}
func (m *RpcSingleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcSingleResponse.Merge(m, src)
}
func (m *RpcSingleResponse) XXX_Size() int {
	return xxx_messageInfo_RpcSingleResponse.Size(m)
}
func (m *RpcSingleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcSingleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RpcSingleResponse proto.InternalMessageInfo

func (m *RpcSingleResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RpcSingleResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RpcSingleResponse) GetErrorNo() int32 {
	if m != nil {
		return m.ErrorNo
	}
	return 0
}

func (m *RpcSingleResponse) GetData() *KeeperStruct {
	if m != nil {
		return m.Data
	}
	return nil
}

type RpcMulitResponse struct {
	Code                 int32           `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	ErrorNo              int32           `protobuf:"varint,3,opt,name=error_no,json=errorNo,proto3" json:"error_no,omitempty"`
	Data                 []*KeeperStruct `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RpcMulitResponse) Reset()         { *m = RpcMulitResponse{} }
func (m *RpcMulitResponse) String() string { return proto.CompactTextString(m) }
func (*RpcMulitResponse) ProtoMessage()    {}
func (*RpcMulitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{7}
}

func (m *RpcMulitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcMulitResponse.Unmarshal(m, b)
}
func (m *RpcMulitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcMulitResponse.Marshal(b, m, deterministic)
}
func (m *RpcMulitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcMulitResponse.Merge(m, src)
}
func (m *RpcMulitResponse) XXX_Size() int {
	return xxx_messageInfo_RpcMulitResponse.Size(m)
}
func (m *RpcMulitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcMulitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RpcMulitResponse proto.InternalMessageInfo

func (m *RpcMulitResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RpcMulitResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RpcMulitResponse) GetErrorNo() int32 {
	if m != nil {
		return m.ErrorNo
	}
	return 0
}

func (m *RpcMulitResponse) GetData() []*KeeperStruct {
	if m != nil {
		return m.Data
	}
	return nil
}

type RpcMulitChannelsResponse struct {
	Code                 int32                 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string                `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	ErrorNo              int32                 `protobuf:"varint,3,opt,name=error_no,json=errorNo,proto3" json:"error_no,omitempty"`
	Data                 []*RpcChannelResponse `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RpcMulitChannelsResponse) Reset()         { *m = RpcMulitChannelsResponse{} }
func (m *RpcMulitChannelsResponse) String() string { return proto.CompactTextString(m) }
func (*RpcMulitChannelsResponse) ProtoMessage()    {}
func (*RpcMulitChannelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{8}
}

func (m *RpcMulitChannelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcMulitChannelsResponse.Unmarshal(m, b)
}
func (m *RpcMulitChannelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcMulitChannelsResponse.Marshal(b, m, deterministic)
}
func (m *RpcMulitChannelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcMulitChannelsResponse.Merge(m, src)
}
func (m *RpcMulitChannelsResponse) XXX_Size() int {
	return xxx_messageInfo_RpcMulitChannelsResponse.Size(m)
}
func (m *RpcMulitChannelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcMulitChannelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RpcMulitChannelsResponse proto.InternalMessageInfo

func (m *RpcMulitChannelsResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RpcMulitChannelsResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RpcMulitChannelsResponse) GetErrorNo() int32 {
	if m != nil {
		return m.ErrorNo
	}
	return 0
}

func (m *RpcMulitChannelsResponse) GetData() []*RpcChannelResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

type PingRequest struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Ack                  int32    `protobuf:"varint,2,opt,name=ack,proto3" json:"ack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{9}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PingRequest) GetAck() int32 {
	if m != nil {
		return m.Ack
	}
	return 0
}

type PingResponse struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Ack                  int32    `protobuf:"varint,2,opt,name=ack,proto3" json:"ack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c4f0e9ae13493f, []int{10}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PingResponse) GetAck() int32 {
	if m != nil {
		return m.Ack
	}
	return 0
}

func init() {
	proto.RegisterEnum("rpc.RpcSingleRequest_DataType", RpcSingleRequest_DataType_name, RpcSingleRequest_DataType_value)
	proto.RegisterEnum("rpc.RpcMulitRequest_DataType", RpcMulitRequest_DataType_name, RpcMulitRequest_DataType_value)
	proto.RegisterEnum("rpc.RpcMulitChannelsRequest_DataType", RpcMulitChannelsRequest_DataType_name, RpcMulitChannelsRequest_DataType_value)
	proto.RegisterType((*RpcSingleRequest)(nil), "rpc.RpcSingleRequest")
	proto.RegisterType((*RpcMulitRequest)(nil), "rpc.RpcMulitRequest")
	proto.RegisterType((*RpcMulitChannelsRequest)(nil), "rpc.RpcMulitChannelsRequest")
	proto.RegisterType((*RpcErrorResponse)(nil), "rpc.RpcErrorResponse")
	proto.RegisterType((*KeeperStruct)(nil), "rpc.KeeperStruct")
	proto.RegisterType((*RpcChannelResponse)(nil), "rpc.RpcChannelResponse")
	proto.RegisterType((*RpcSingleResponse)(nil), "rpc.RpcSingleResponse")
	proto.RegisterType((*RpcMulitResponse)(nil), "rpc.RpcMulitResponse")
	proto.RegisterType((*RpcMulitChannelsResponse)(nil), "rpc.RpcMulitChannelsResponse")
	proto.RegisterType((*PingRequest)(nil), "rpc.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "rpc.PingResponse")
}

func init() { proto.RegisterFile("bitMapRpc.proto", fileDescriptor_39c4f0e9ae13493f) }

var fileDescriptor_39c4f0e9ae13493f = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xd1, 0x6a, 0x13, 0x4d,
	0x14, 0xee, 0x66, 0xb2, 0x6d, 0x73, 0x52, 0xfe, 0xa6, 0xf3, 0x5b, 0xbb, 0x86, 0xb6, 0x86, 0x85,
	0x62, 0x40, 0x0c, 0x18, 0xaf, 0x04, 0x2b, 0x58, 0xdb, 0x2b, 0x69, 0x91, 0x49, 0x44, 0xf0, 0xa6,
	0x4c, 0x37, 0xa7, 0x71, 0x69, 0xb2, 0x3b, 0xce, 0x4c, 0x84, 0xe8, 0x0b, 0x78, 0xe5, 0x0b, 0xfa,
	0x0c, 0xbe, 0x83, 0xcc, 0xd9, 0xdd, 0x74, 0x13, 0x6b, 0x50, 0xa8, 0xde, 0x7d, 0x73, 0x66, 0xcf,
	0x39, 0xdf, 0x77, 0xe6, 0x9b, 0x59, 0xd8, 0xbc, 0x88, 0xed, 0xa9, 0x54, 0x42, 0x45, 0x1d, 0xa5,
	0x53, 0x9b, 0x72, 0xa6, 0x55, 0x14, 0x7e, 0xf3, 0xa0, 0x21, 0x54, 0xd4, 0x8b, 0x93, 0xe1, 0x08,
	0x05, 0x7e, 0x98, 0xa0, 0xb1, 0x9c, 0x43, 0x75, 0x20, 0x2d, 0x06, 0x5e, 0xcb, 0x6b, 0xd7, 0x04,
	0x61, 0xde, 0x85, 0xaa, 0x9d, 0x2a, 0x0c, 0x2a, 0x2d, 0xaf, 0xfd, 0x5f, 0x77, 0xbf, 0xa3, 0x55,
	0xd4, 0x59, 0x4c, 0xec, 0x1c, 0x4b, 0x2b, 0xfb, 0x53, 0x85, 0x82, 0xbe, 0xe5, 0x0d, 0x60, 0x03,
	0x39, 0x0d, 0x58, 0xcb, 0x6b, 0xfb, 0xc2, 0x41, 0x1e, 0xc0, 0x5a, 0xf4, 0x5e, 0x26, 0x09, 0x8e,
	0x82, 0x2a, 0x15, 0x2f, 0x96, 0xae, 0xa7, 0x4e, 0x47, 0x18, 0xf8, 0x59, 0x4f, 0x87, 0xf9, 0x7d,
	0xa8, 0x0f, 0xa4, 0x95, 0xe7, 0x26, 0x9d, 0xe8, 0x08, 0x83, 0x55, 0xda, 0x02, 0x17, 0xea, 0x51,
	0x24, 0xdc, 0x87, 0xf5, 0xa2, 0x25, 0x5f, 0x03, 0x76, 0x72, 0x76, 0xdc, 0x58, 0xe1, 0x35, 0xf0,
	0x7b, 0xfd, 0x17, 0xa2, 0xdf, 0xf0, 0x9c, 0xba, 0x4d, 0xa1, 0xa2, 0xd3, 0xc9, 0x28, 0xb6, 0xcb,
	0xc4, 0x3d, 0x9e, 0x13, 0xb7, 0x57, 0x88, 0x2b, 0xe7, 0x2d, 0x6a, 0xa3, 0x32, 0x53, 0x13, 0xb0,
	0x16, 0x6b, 0xfb, 0x82, 0xf0, 0xbf, 0x56, 0xf7, 0xdd, 0x83, 0x9d, 0x82, 0xe5, 0xcb, 0xac, 0x91,
	0x59, 0xa6, 0xf2, 0xe9, 0x9c, 0xca, 0x83, 0x39, 0x95, 0x0b, 0xf9, 0xbf, 0xa3, 0xb6, 0x09, 0xeb,
	0xb9, 0x3c, 0x13, 0x54, 0x5b, 0xac, 0x5d, 0x13, 0xb3, 0xf5, 0xdf, 0xd1, 0xdb, 0x23, 0xab, 0x9e,
	0x68, 0x9d, 0x6a, 0x81, 0x46, 0xa5, 0x89, 0x21, 0x62, 0x51, 0x3a, 0xc8, 0x74, 0xfa, 0x82, 0xb0,
	0xb3, 0xdd, 0xd8, 0x0c, 0x49, 0x66, 0x4d, 0x38, 0xc8, 0xef, 0xc1, 0x3a, 0xba, 0xb4, 0xf3, 0x24,
	0xcd, 0xdd, 0xb8, 0x46, 0xeb, 0xb3, 0x34, 0xbc, 0x84, 0x8d, 0x57, 0x88, 0x0a, 0x75, 0xcf, 0xea,
	0x49, 0x44, 0x83, 0xbb, 0x42, 0x54, 0xc5, 0xe0, 0x1c, 0x76, 0xb1, 0x44, 0x8e, 0x31, 0xaf, 0x48,
	0xd8, 0xa9, 0xd7, 0x38, 0x8c, 0x8d, 0x45, 0x4d, 0x25, 0x99, 0x98, 0xad, 0xdd, 0xf7, 0x23, 0xbc,
	0xb4, 0x64, 0x02, 0x26, 0x08, 0x87, 0x6f, 0x81, 0x0b, 0x15, 0xe5, 0x63, 0x9e, 0xd1, 0x2f, 0x39,
	0xc6, 0x9b, 0x77, 0xcc, 0x03, 0xf0, 0x63, 0x8b, 0x63, 0x13, 0x54, 0x5a, 0xac, 0x5d, 0xef, 0x6e,
	0xd1, 0x69, 0x95, 0x99, 0x8a, 0x6c, 0x3f, 0xfc, 0x0c, 0x5b, 0xa5, 0x7b, 0x78, 0x4b, 0x63, 0xe1,
	0x07, 0xe4, 0x1f, 0x49, 0x12, 0x6e, 0xec, 0x4e, 0xdb, 0xe1, 0x27, 0x3a, 0x92, 0xfc, 0x9e, 0xdc,
	0x7e, 0x6f, 0xb6, 0xac, 0xf7, 0x17, 0x0f, 0x82, 0x9f, 0xed, 0x7b, 0x5b, 0x24, 0x1e, 0xce, 0x91,
	0xd8, 0x29, 0x2e, 0xcb, 0xc2, 0x01, 0xe6, 0x54, 0x0e, 0xa1, 0xfe, 0x3a, 0x4e, 0x86, 0xc5, 0xe5,
	0xdb, 0x85, 0x9a, 0x8d, 0xc7, 0x68, 0xac, 0x1c, 0x67, 0x46, 0x62, 0xe2, 0x3a, 0xe0, 0x68, 0xc8,
	0xe8, 0x8a, 0x68, 0xf8, 0xc2, 0xc1, 0xf0, 0x39, 0x6c, 0x64, 0xe9, 0x39, 0xf9, 0x3f, 0xcc, 0xef,
	0x7e, 0xad, 0x00, 0x1c, 0xd1, 0xeb, 0xde, 0x43, 0xfd, 0x91, 0x3f, 0x82, 0xaa, 0x2b, 0xc7, 0x1b,
	0x44, 0xba, 0x44, 0xac, 0xb9, 0x55, 0x8a, 0x64, 0xbd, 0xc2, 0x15, 0x7e, 0x08, 0xf0, 0xc6, 0xa0,
	0xce, 0x26, 0xcc, 0xb7, 0x6f, 0x7c, 0xd9, 0x9b, 0x77, 0x17, 0xc3, 0xb3, 0xf4, 0x67, 0x50, 0xbf,
	0x4e, 0x37, 0xfc, 0xce, 0x4d, 0x8f, 0x67, 0x73, 0x7b, 0x21, 0x3a, 0xcb, 0xee, 0xc3, 0xff, 0xc5,
	0xd9, 0x95, 0xab, 0xec, 0x2e, 0x7b, 0x9c, 0x9a, 0x7b, 0xbf, 0xd8, 0x2d, 0xaa, 0x1e, 0x55, 0xdf,
	0x55, 0xd4, 0xc5, 0xc5, 0x2a, 0xfd, 0xe7, 0x9e, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xa3,
	0x22, 0x61, 0xfa, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BitMapServClient is the client API for BitMapServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BitMapServClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	UserKeeper(ctx context.Context, in *RpcSingleRequest, opts ...grpc.CallOption) (*RpcSingleResponse, error)
	UserKeepers(ctx context.Context, in *RpcMulitRequest, opts ...grpc.CallOption) (*RpcMulitResponse, error)
	ChannelsUserKeepers(ctx context.Context, in *RpcMulitChannelsRequest, opts ...grpc.CallOption) (*RpcMulitChannelsResponse, error)
}

type bitMapServClient struct {
	cc *grpc.ClientConn
}

func NewBitMapServClient(cc *grpc.ClientConn) BitMapServClient {
	return &bitMapServClient{cc}
}

func (c *bitMapServClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/rpc.BitMapServ/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitMapServClient) UserKeeper(ctx context.Context, in *RpcSingleRequest, opts ...grpc.CallOption) (*RpcSingleResponse, error) {
	out := new(RpcSingleResponse)
	err := c.cc.Invoke(ctx, "/rpc.BitMapServ/UserKeeper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitMapServClient) UserKeepers(ctx context.Context, in *RpcMulitRequest, opts ...grpc.CallOption) (*RpcMulitResponse, error) {
	out := new(RpcMulitResponse)
	err := c.cc.Invoke(ctx, "/rpc.BitMapServ/UserKeepers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitMapServClient) ChannelsUserKeepers(ctx context.Context, in *RpcMulitChannelsRequest, opts ...grpc.CallOption) (*RpcMulitChannelsResponse, error) {
	out := new(RpcMulitChannelsResponse)
	err := c.cc.Invoke(ctx, "/rpc.BitMapServ/ChannelsUserKeepers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BitMapServServer is the server API for BitMapServ service.
type BitMapServServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	UserKeeper(context.Context, *RpcSingleRequest) (*RpcSingleResponse, error)
	UserKeepers(context.Context, *RpcMulitRequest) (*RpcMulitResponse, error)
	ChannelsUserKeepers(context.Context, *RpcMulitChannelsRequest) (*RpcMulitChannelsResponse, error)
}

// UnimplementedBitMapServServer can be embedded to have forward compatible implementations.
type UnimplementedBitMapServServer struct {
}

func (*UnimplementedBitMapServServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedBitMapServServer) UserKeeper(ctx context.Context, req *RpcSingleRequest) (*RpcSingleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserKeeper not implemented")
}
func (*UnimplementedBitMapServServer) UserKeepers(ctx context.Context, req *RpcMulitRequest) (*RpcMulitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserKeepers not implemented")
}
func (*UnimplementedBitMapServServer) ChannelsUserKeepers(ctx context.Context, req *RpcMulitChannelsRequest) (*RpcMulitChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelsUserKeepers not implemented")
}

func RegisterBitMapServServer(s *grpc.Server, srv BitMapServServer) {
	s.RegisterService(&_BitMapServ_serviceDesc, srv)
}

func _BitMapServ_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitMapServServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.BitMapServ/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitMapServServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitMapServ_UserKeeper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcSingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitMapServServer).UserKeeper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.BitMapServ/UserKeeper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitMapServServer).UserKeeper(ctx, req.(*RpcSingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitMapServ_UserKeepers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcMulitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitMapServServer).UserKeepers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.BitMapServ/UserKeepers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitMapServServer).UserKeepers(ctx, req.(*RpcMulitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitMapServ_ChannelsUserKeepers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcMulitChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitMapServServer).ChannelsUserKeepers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.BitMapServ/ChannelsUserKeepers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitMapServServer).ChannelsUserKeepers(ctx, req.(*RpcMulitChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BitMapServ_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.BitMapServ",
	HandlerType: (*BitMapServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BitMapServ_Ping_Handler,
		},
		{
			MethodName: "UserKeeper",
			Handler:    _BitMapServ_UserKeeper_Handler,
		},
		{
			MethodName: "UserKeepers",
			Handler:    _BitMapServ_UserKeepers_Handler,
		},
		{
			MethodName: "ChannelsUserKeepers",
			Handler:    _BitMapServ_ChannelsUserKeepers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bitMapRpc.proto",
}
