// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contactpb/contact.proto

package contactpb

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

type Contact struct {
	PhoneNumber          string   `protobuf:"bytes,1,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c204982b934798c, []int{0}
}

func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (m *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(m, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contact) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type InsertRequest struct {
	Contact              *Contact `protobuf:"bytes,1,opt,name=Contact,proto3" json:"Contact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertRequest) Reset()         { *m = InsertRequest{} }
func (m *InsertRequest) String() string { return proto.CompactTextString(m) }
func (*InsertRequest) ProtoMessage()    {}
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c204982b934798c, []int{1}
}

func (m *InsertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertRequest.Unmarshal(m, b)
}
func (m *InsertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertRequest.Marshal(b, m, deterministic)
}
func (m *InsertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertRequest.Merge(m, src)
}
func (m *InsertRequest) XXX_Size() int {
	return xxx_messageInfo_InsertRequest.Size(m)
}
func (m *InsertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InsertRequest proto.InternalMessageInfo

func (m *InsertRequest) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

type InsertResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertResponse) Reset()         { *m = InsertResponse{} }
func (m *InsertResponse) String() string { return proto.CompactTextString(m) }
func (*InsertResponse) ProtoMessage()    {}
func (*InsertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c204982b934798c, []int{2}
}

func (m *InsertResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertResponse.Unmarshal(m, b)
}
func (m *InsertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertResponse.Marshal(b, m, deterministic)
}
func (m *InsertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertResponse.Merge(m, src)
}
func (m *InsertResponse) XXX_Size() int {
	return xxx_messageInfo_InsertResponse.Size(m)
}
func (m *InsertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InsertResponse proto.InternalMessageInfo

func (m *InsertResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *InsertResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReadRequest struct {
	PhoneNumber          string   `protobuf:"bytes,1,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c204982b934798c, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

type ReadResponse struct {
	Contact              *Contact `protobuf:"bytes,1,opt,name=contact,proto3" json:"contact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c204982b934798c, []int{4}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetContact() *Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

type UpdateRequest struct {
	NewContact           *Contact `protobuf:"bytes,1,opt,name=newContact,proto3" json:"newContact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c204982b934798c, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetNewContact() *Contact {
	if m != nil {
		return m.NewContact
	}
	return nil
}

type UpdateResponse struct {
	UpdateContact        *Contact `protobuf:"bytes,1,opt,name=updateContact,proto3" json:"updateContact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c204982b934798c, []int{6}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetUpdateContact() *Contact {
	if m != nil {
		return m.UpdateContact
	}
	return nil
}

type DeleteRequest struct {
	PhoneNumber          string   `protobuf:"bytes,1,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c204982b934798c, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

type DeleteResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c204982b934798c, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *DeleteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SearchRequest struct {
	SearchName           string   `protobuf:"bytes,1,opt,name=searchName,proto3" json:"searchName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c204982b934798c, []int{9}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetSearchName() string {
	if m != nil {
		return m.SearchName
	}
	return ""
}

type SearchResponse struct {
	Results              []*Contact `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c204982b934798c, []int{10}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetResults() []*Contact {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*Contact)(nil), "contact.Contact")
	proto.RegisterType((*InsertRequest)(nil), "contact.InsertRequest")
	proto.RegisterType((*InsertResponse)(nil), "contact.InsertResponse")
	proto.RegisterType((*ReadRequest)(nil), "contact.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "contact.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "contact.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "contact.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "contact.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "contact.DeleteResponse")
	proto.RegisterType((*SearchRequest)(nil), "contact.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "contact.SearchResponse")
}

func init() {
	proto.RegisterFile("contactpb/contact.proto", fileDescriptor_7c204982b934798c)
}

var fileDescriptor_7c204982b934798c = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x4f, 0xe3, 0x30,
	0x10, 0xed, 0xd7, 0xb6, 0xea, 0x64, 0x13, 0xad, 0xac, 0xdd, 0x6d, 0xd4, 0xc3, 0xaa, 0xf2, 0x69,
	0xd5, 0x43, 0xbb, 0x5b, 0x04, 0x07, 0xca, 0x05, 0xca, 0x01, 0x7a, 0xe8, 0x21, 0x15, 0x07, 0xb8,
	0xb9, 0xc9, 0x88, 0x22, 0xb5, 0x49, 0x88, 0x1d, 0xf8, 0x71, 0xfc, 0x39, 0x14, 0x3b, 0x4e, 0x6d,
	0x3e, 0x94, 0x0b, 0x37, 0xfb, 0x65, 0xde, 0xbc, 0xe7, 0x37, 0xa3, 0xc0, 0x20, 0x4c, 0x62, 0xc1,
	0x42, 0x91, 0x6e, 0xa6, 0xe5, 0x69, 0x92, 0x66, 0x89, 0x48, 0x48, 0xaf, 0xbc, 0xd2, 0x5b, 0xe8,
	0x2d, 0xd4, 0x91, 0x8c, 0xc0, 0x49, 0xb7, 0x49, 0x8c, 0xab, 0x7c, 0xbf, 0xc1, 0xcc, 0x6f, 0x8e,
	0x9a, 0x7f, 0xfb, 0x81, 0x09, 0x11, 0x02, 0x9d, 0x98, 0xed, 0xd1, 0x6f, 0xc9, 0x4f, 0xf2, 0x4c,
	0x7c, 0xe8, 0xb1, 0x28, 0xca, 0x90, 0x73, 0xbf, 0x2d, 0x61, 0x7d, 0xa5, 0x73, 0x70, 0xaf, 0x63,
	0x8e, 0x99, 0x08, 0xf0, 0x31, 0x47, 0x2e, 0xc8, 0xb8, 0xd2, 0x92, 0xcd, 0x9d, 0xd9, 0x8f, 0x89,
	0x76, 0x55, 0xe2, 0x81, 0x2e, 0xa0, 0x4b, 0xf0, 0x34, 0x99, 0xa7, 0x49, 0xcc, 0x91, 0xfc, 0x01,
	0xe0, 0x82, 0x89, 0x9c, 0x2f, 0x92, 0x08, 0x65, 0x83, 0x6f, 0x81, 0x81, 0x14, 0x46, 0xf6, 0xc8,
	0x39, 0xbb, 0xd7, 0xfe, 0xf4, 0x95, 0x4e, 0xc1, 0x09, 0x90, 0x45, 0xda, 0x46, 0xed, 0x3b, 0xe9,
	0x29, 0x7c, 0x57, 0x84, 0x52, 0x7a, 0x0c, 0x3a, 0xaf, 0xcf, 0x8d, 0xeb, 0x40, 0xcf, 0xc1, 0xbd,
	0x49, 0x23, 0x26, 0x50, 0xcb, 0xfd, 0x03, 0x88, 0xf1, 0xb9, 0xee, 0xe1, 0x46, 0x0d, 0xbd, 0x02,
	0x4f, 0xb7, 0x28, 0x0d, 0x9c, 0x80, 0x9b, 0x4b, 0xa4, 0xae, 0x8d, 0x5d, 0x46, 0xff, 0x83, 0x7b,
	0x89, 0x3b, 0x3c, 0x98, 0xa9, 0x7f, 0xfb, 0x12, 0x3c, 0x4d, 0xf9, 0x82, 0xe0, 0xdd, 0x35, 0xb2,
	0x2c, 0xdc, 0x6a, 0xf9, 0xa2, 0x95, 0x04, 0x56, 0xc5, 0x1a, 0x29, 0x75, 0x03, 0xa1, 0x67, 0xe0,
	0x69, 0xc2, 0x21, 0xfa, 0x0c, 0x79, 0xbe, 0x13, 0xdc, 0x6f, 0x8e, 0xda, 0x1f, 0x47, 0x5f, 0x16,
	0xcc, 0x5e, 0x5a, 0xe0, 0x95, 0xe0, 0x1a, 0xb3, 0xa7, 0x87, 0x10, 0xc9, 0x1c, 0xba, 0x6a, 0x8d,
	0xc8, 0xef, 0x8a, 0x67, 0x2d, 0xe5, 0x70, 0xf0, 0x0e, 0x57, 0xca, 0xb4, 0x41, 0x8e, 0xa1, 0x53,
	0xac, 0x01, 0xf9, 0x59, 0x95, 0x18, 0x6b, 0x34, 0xfc, 0xf5, 0x06, 0xad, 0x68, 0x73, 0xe8, 0xaa,
	0xf1, 0x19, 0x9a, 0xd6, 0x4a, 0x18, 0x9a, 0xf6, 0x9c, 0x15, 0x59, 0xc5, 0x6f, 0x90, 0xad, 0x11,
	0x1a, 0x64, 0x7b, 0x4e, 0x8a, 0xac, 0xe2, 0x33, 0xc8, 0xd6, 0x00, 0x0c, 0xb2, 0x9d, 0x33, 0x6d,
	0x5c, 0x38, 0x77, 0xfd, 0xea, 0x6f, 0xb1, 0xe9, 0xca, 0xdf, 0xc4, 0xd1, 0x6b, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x64, 0x05, 0x17, 0xda, 0x41, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ContactServiceClient is the client API for ContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContactServiceClient interface {
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type contactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContactServiceClient(cc grpc.ClientConnInterface) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, "/contact.ContactService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/contact.ContactService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/contact.ContactService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/contact.ContactService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/contact.ContactService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServiceServer is the server API for ContactService service.
type ContactServiceServer interface {
	Insert(context.Context, *InsertRequest) (*InsertResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

// UnimplementedContactServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContactServiceServer struct {
}

func (*UnimplementedContactServiceServer) Insert(ctx context.Context, req *InsertRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (*UnimplementedContactServiceServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedContactServiceServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedContactServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedContactServiceServer) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterContactServiceServer(s *grpc.Server, srv ContactServiceServer) {
	s.RegisterService(&_ContactService_serviceDesc, srv)
}

func _ContactService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contact.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _ContactService_Insert_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ContactService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ContactService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ContactService_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _ContactService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contactpb/contact.proto",
}
