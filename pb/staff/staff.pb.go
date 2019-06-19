// Code generated by protoc-gen-go. DO NOT EDIT.
// source: staff/staff.proto

package staff

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetStaffMessage struct {
	StaffId              int64    `protobuf:"varint,1,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStaffMessage) Reset()         { *m = GetStaffMessage{} }
func (m *GetStaffMessage) String() string { return proto.CompactTextString(m) }
func (*GetStaffMessage) ProtoMessage()    {}
func (*GetStaffMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0877d72e908c1a13, []int{0}
}

func (m *GetStaffMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStaffMessage.Unmarshal(m, b)
}
func (m *GetStaffMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStaffMessage.Marshal(b, m, deterministic)
}
func (m *GetStaffMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStaffMessage.Merge(m, src)
}
func (m *GetStaffMessage) XXX_Size() int {
	return xxx_messageInfo_GetStaffMessage.Size(m)
}
func (m *GetStaffMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStaffMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetStaffMessage proto.InternalMessageInfo

func (m *GetStaffMessage) GetStaffId() int64 {
	if m != nil {
		return m.StaffId
	}
	return 0
}

type StaffResponse struct {
	StaffId              int64    `protobuf:"varint,1,opt,name=staff_id,json=staffId,proto3" json:"staff_id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Age                  int32    `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaffResponse) Reset()         { *m = StaffResponse{} }
func (m *StaffResponse) String() string { return proto.CompactTextString(m) }
func (*StaffResponse) ProtoMessage()    {}
func (*StaffResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0877d72e908c1a13, []int{1}
}

func (m *StaffResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaffResponse.Unmarshal(m, b)
}
func (m *StaffResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaffResponse.Marshal(b, m, deterministic)
}
func (m *StaffResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaffResponse.Merge(m, src)
}
func (m *StaffResponse) XXX_Size() int {
	return xxx_messageInfo_StaffResponse.Size(m)
}
func (m *StaffResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StaffResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StaffResponse proto.InternalMessageInfo

func (m *StaffResponse) GetStaffId() int64 {
	if m != nil {
		return m.StaffId
	}
	return 0
}

func (m *StaffResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *StaffResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *StaffResponse) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *StaffResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*GetStaffMessage)(nil), "staff.GetStaffMessage")
	proto.RegisterType((*StaffResponse)(nil), "staff.StaffResponse")
}

func init() { proto.RegisterFile("staff/staff.proto", fileDescriptor_0877d72e908c1a13) }

var fileDescriptor_0877d72e908c1a13 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2e, 0x49, 0x4c,
	0x4b, 0xd3, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x92, 0x0e,
	0x17, 0xbf, 0x7b, 0x6a, 0x49, 0x30, 0x88, 0xed, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0x24,
	0xc9, 0xc5, 0x01, 0x96, 0x8b, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0x07,
	0xf3, 0x3d, 0x53, 0x94, 0xfa, 0x18, 0xb9, 0x78, 0xc1, 0x6a, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3,
	0x8a, 0xf1, 0x29, 0x16, 0x92, 0xe5, 0xe2, 0x4a, 0xcb, 0x2c, 0x2a, 0x2e, 0x89, 0xcf, 0x4b, 0xcc,
	0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x04, 0x8b, 0xf8, 0x25, 0xe6, 0xa6, 0x0a,
	0x49, 0x73, 0x71, 0xe6, 0x24, 0xc2, 0x64, 0x99, 0xc1, 0xb2, 0x1c, 0x20, 0x01, 0xb0, 0xa4, 0x00,
	0x17, 0x73, 0x62, 0x7a, 0xaa, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x88, 0x29, 0x24, 0xc2,
	0xc5, 0x9a, 0x9a, 0x9b, 0x98, 0x99, 0x23, 0xc1, 0x0a, 0x56, 0x0a, 0xe1, 0x18, 0xb9, 0x70, 0x71,
	0x80, 0xdd, 0xe3, 0x18, 0xe0, 0x29, 0x64, 0xc1, 0xc5, 0x01, 0xf3, 0x8a, 0x90, 0x98, 0x1e, 0xc4,
	0xaf, 0x68, 0x7e, 0x93, 0x12, 0x81, 0x8a, 0xa3, 0x78, 0xc2, 0x89, 0x3d, 0x0a, 0x12, 0x1a, 0x49,
	0x6c, 0xe0, 0xb0, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x32, 0x9b, 0xf8, 0x86, 0x30, 0x01,
	0x00, 0x00,
}
