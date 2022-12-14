// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: society.proto

package protorender

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Society struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint32  `protobuf:"fixed32,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NameComplement   string  `protobuf:"bytes,3,opt,name=nameComplement,proto3" json:"nameComplement,omitempty"`
	Logo             string  `protobuf:"bytes,4,opt,name=logo,proto3" json:"logo,omitempty"`
	Abstract         string  `protobuf:"bytes,5,opt,name=abstract,proto3" json:"abstract,omitempty"`
	AbstractPhoto    string  `protobuf:"bytes,6,opt,name=abstractPhoto,proto3" json:"abstractPhoto,omitempty"`
	Country          string  `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`
	Department       string  `protobuf:"bytes,8,opt,name=department,proto3" json:"department,omitempty"`
	FamiliesNumber   int32   `protobuf:"varint,9,opt,name=familiesNumber,proto3" json:"familiesNumber,omitempty"`
	AssociatesNumber int32   `protobuf:"varint,10,opt,name=associatesNumber,proto3" json:"associatesNumber,omitempty"`
	SalesNumber      int32   `protobuf:"varint,11,opt,name=salesNumber,proto3" json:"salesNumber,omitempty"`
	Longitude        float32 `protobuf:"fixed32,12,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude         float32 `protobuf:"fixed32,13,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Web              string  `protobuf:"bytes,14,opt,name=web,proto3" json:"web,omitempty"`
	Video            string  `protobuf:"bytes,15,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *Society) Reset() {
	*x = Society{}
	if protoimpl.UnsafeEnabled {
		mi := &file_society_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Society) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Society) ProtoMessage() {}

func (x *Society) ProtoReflect() protoreflect.Message {
	mi := &file_society_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Society.ProtoReflect.Descriptor instead.
func (*Society) Descriptor() ([]byte, []int) {
	return file_society_proto_rawDescGZIP(), []int{0}
}

func (x *Society) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Society) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Society) GetNameComplement() string {
	if x != nil {
		return x.NameComplement
	}
	return ""
}

func (x *Society) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Society) GetAbstract() string {
	if x != nil {
		return x.Abstract
	}
	return ""
}

func (x *Society) GetAbstractPhoto() string {
	if x != nil {
		return x.AbstractPhoto
	}
	return ""
}

func (x *Society) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Society) GetDepartment() string {
	if x != nil {
		return x.Department
	}
	return ""
}

func (x *Society) GetFamiliesNumber() int32 {
	if x != nil {
		return x.FamiliesNumber
	}
	return 0
}

func (x *Society) GetAssociatesNumber() int32 {
	if x != nil {
		return x.AssociatesNumber
	}
	return 0
}

func (x *Society) GetSalesNumber() int32 {
	if x != nil {
		return x.SalesNumber
	}
	return 0
}

func (x *Society) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Society) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Society) GetWeb() string {
	if x != nil {
		return x.Web
	}
	return ""
}

func (x *Society) GetVideo() string {
	if x != nil {
		return x.Video
	}
	return ""
}

type RequestSocieties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"fixed32,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestSocieties) Reset() {
	*x = RequestSocieties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_society_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestSocieties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestSocieties) ProtoMessage() {}

func (x *RequestSocieties) ProtoReflect() protoreflect.Message {
	mi := &file_society_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestSocieties.ProtoReflect.Descriptor instead.
func (*RequestSocieties) Descriptor() ([]byte, []int) {
	return file_society_proto_rawDescGZIP(), []int{1}
}

func (x *RequestSocieties) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ResponseSocieties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Societies []*Society `protobuf:"bytes,1,rep,name=societies,proto3" json:"societies,omitempty"`
}

func (x *ResponseSocieties) Reset() {
	*x = ResponseSocieties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_society_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseSocieties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseSocieties) ProtoMessage() {}

func (x *ResponseSocieties) ProtoReflect() protoreflect.Message {
	mi := &file_society_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseSocieties.ProtoReflect.Descriptor instead.
func (*ResponseSocieties) Descriptor() ([]byte, []int) {
	return file_society_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseSocieties) GetSocieties() []*Society {
	if x != nil {
		return x.Societies
	}
	return nil
}

var File_society_proto protoreflect.FileDescriptor

var file_society_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0xbd, 0x03, 0x0a,
	0x07, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x07, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x62, 0x73, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x62, 0x73, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x65, 0x73,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x69, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10,
	0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74,
	0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73,
	0x61, 0x6c, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x65, 0x62, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x77, 0x65, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x22, 0x0a, 0x10,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x47, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x6f, 0x63, 0x69,
	0x65, 0x74, 0x69, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x73, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x79, 0x52, 0x09,
	0x73, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x32, 0xe6, 0x02, 0x0a, 0x0e, 0x53, 0x6f,
	0x63, 0x69, 0x65, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x09,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x50, 0x43, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x6f, 0x63, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53,
	0x6f, 0x63, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x16, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x52, 0x50, 0x43, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74,
	0x69, 0x65, 0x73, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5b, 0x0a,
	0x16, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x52, 0x50, 0x43, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x6f, 0x63,
	0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x6f, 0x63,
	0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x56, 0x0a, 0x19, 0x42, 0x69,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x52, 0x50, 0x43, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x6f, 0x63,
	0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x65, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x19, 0x5a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_society_proto_rawDescOnce sync.Once
	file_society_proto_rawDescData = file_society_proto_rawDesc
)

func file_society_proto_rawDescGZIP() []byte {
	file_society_proto_rawDescOnce.Do(func() {
		file_society_proto_rawDescData = protoimpl.X.CompressGZIP(file_society_proto_rawDescData)
	})
	return file_society_proto_rawDescData
}

var file_society_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_society_proto_goTypes = []interface{}{
	(*Society)(nil),           // 0: protorender.Society
	(*RequestSocieties)(nil),  // 1: protorender.RequestSocieties
	(*ResponseSocieties)(nil), // 2: protorender.ResponseSocieties
}
var file_society_proto_depIdxs = []int32{
	0, // 0: protorender.ResponseSocieties.societies:type_name -> protorender.Society
	1, // 1: protorender.SocietyService.SimpleRPC:input_type -> protorender.RequestSocieties
	1, // 2: protorender.SocietyService.ServerSideStreamingRPC:input_type -> protorender.RequestSocieties
	1, // 3: protorender.SocietyService.ClientSideStreamingRPC:input_type -> protorender.RequestSocieties
	1, // 4: protorender.SocietyService.BidirectionalStreamingRPC:input_type -> protorender.RequestSocieties
	2, // 5: protorender.SocietyService.SimpleRPC:output_type -> protorender.ResponseSocieties
	0, // 6: protorender.SocietyService.ServerSideStreamingRPC:output_type -> protorender.Society
	2, // 7: protorender.SocietyService.ClientSideStreamingRPC:output_type -> protorender.ResponseSocieties
	0, // 8: protorender.SocietyService.BidirectionalStreamingRPC:output_type -> protorender.Society
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_society_proto_init() }
func file_society_proto_init() {
	if File_society_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_society_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Society); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_society_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestSocieties); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_society_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseSocieties); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_society_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_society_proto_goTypes,
		DependencyIndexes: file_society_proto_depIdxs,
		MessageInfos:      file_society_proto_msgTypes,
	}.Build()
	File_society_proto = out.File
	file_society_proto_rawDesc = nil
	file_society_proto_goTypes = nil
	file_society_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SocietyServiceClient is the client API for SocietyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SocietyServiceClient interface {
	SimpleRPC(ctx context.Context, in *RequestSocieties, opts ...grpc.CallOption) (*ResponseSocieties, error)
	ServerSideStreamingRPC(ctx context.Context, in *RequestSocieties, opts ...grpc.CallOption) (SocietyService_ServerSideStreamingRPCClient, error)
	ClientSideStreamingRPC(ctx context.Context, opts ...grpc.CallOption) (SocietyService_ClientSideStreamingRPCClient, error)
	BidirectionalStreamingRPC(ctx context.Context, opts ...grpc.CallOption) (SocietyService_BidirectionalStreamingRPCClient, error)
}

type societyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSocietyServiceClient(cc grpc.ClientConnInterface) SocietyServiceClient {
	return &societyServiceClient{cc}
}

func (c *societyServiceClient) SimpleRPC(ctx context.Context, in *RequestSocieties, opts ...grpc.CallOption) (*ResponseSocieties, error) {
	out := new(ResponseSocieties)
	err := c.cc.Invoke(ctx, "/protorender.SocietyService/SimpleRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *societyServiceClient) ServerSideStreamingRPC(ctx context.Context, in *RequestSocieties, opts ...grpc.CallOption) (SocietyService_ServerSideStreamingRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SocietyService_serviceDesc.Streams[0], "/protorender.SocietyService/ServerSideStreamingRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &societyServiceServerSideStreamingRPCClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SocietyService_ServerSideStreamingRPCClient interface {
	Recv() (*Society, error)
	grpc.ClientStream
}

type societyServiceServerSideStreamingRPCClient struct {
	grpc.ClientStream
}

func (x *societyServiceServerSideStreamingRPCClient) Recv() (*Society, error) {
	m := new(Society)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *societyServiceClient) ClientSideStreamingRPC(ctx context.Context, opts ...grpc.CallOption) (SocietyService_ClientSideStreamingRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SocietyService_serviceDesc.Streams[1], "/protorender.SocietyService/ClientSideStreamingRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &societyServiceClientSideStreamingRPCClient{stream}
	return x, nil
}

type SocietyService_ClientSideStreamingRPCClient interface {
	Send(*RequestSocieties) error
	CloseAndRecv() (*ResponseSocieties, error)
	grpc.ClientStream
}

type societyServiceClientSideStreamingRPCClient struct {
	grpc.ClientStream
}

func (x *societyServiceClientSideStreamingRPCClient) Send(m *RequestSocieties) error {
	return x.ClientStream.SendMsg(m)
}

func (x *societyServiceClientSideStreamingRPCClient) CloseAndRecv() (*ResponseSocieties, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResponseSocieties)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *societyServiceClient) BidirectionalStreamingRPC(ctx context.Context, opts ...grpc.CallOption) (SocietyService_BidirectionalStreamingRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SocietyService_serviceDesc.Streams[2], "/protorender.SocietyService/BidirectionalStreamingRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &societyServiceBidirectionalStreamingRPCClient{stream}
	return x, nil
}

type SocietyService_BidirectionalStreamingRPCClient interface {
	Send(*RequestSocieties) error
	Recv() (*Society, error)
	grpc.ClientStream
}

type societyServiceBidirectionalStreamingRPCClient struct {
	grpc.ClientStream
}

func (x *societyServiceBidirectionalStreamingRPCClient) Send(m *RequestSocieties) error {
	return x.ClientStream.SendMsg(m)
}

func (x *societyServiceBidirectionalStreamingRPCClient) Recv() (*Society, error) {
	m := new(Society)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SocietyServiceServer is the server API for SocietyService service.
type SocietyServiceServer interface {
	SimpleRPC(context.Context, *RequestSocieties) (*ResponseSocieties, error)
	ServerSideStreamingRPC(*RequestSocieties, SocietyService_ServerSideStreamingRPCServer) error
	ClientSideStreamingRPC(SocietyService_ClientSideStreamingRPCServer) error
	BidirectionalStreamingRPC(SocietyService_BidirectionalStreamingRPCServer) error
}

// UnimplementedSocietyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSocietyServiceServer struct {
}

func (*UnimplementedSocietyServiceServer) SimpleRPC(context.Context, *RequestSocieties) (*ResponseSocieties, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimpleRPC not implemented")
}
func (*UnimplementedSocietyServiceServer) ServerSideStreamingRPC(*RequestSocieties, SocietyService_ServerSideStreamingRPCServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerSideStreamingRPC not implemented")
}
func (*UnimplementedSocietyServiceServer) ClientSideStreamingRPC(SocietyService_ClientSideStreamingRPCServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientSideStreamingRPC not implemented")
}
func (*UnimplementedSocietyServiceServer) BidirectionalStreamingRPC(SocietyService_BidirectionalStreamingRPCServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamingRPC not implemented")
}

func RegisterSocietyServiceServer(s *grpc.Server, srv SocietyServiceServer) {
	s.RegisterService(&_SocietyService_serviceDesc, srv)
}

func _SocietyService_SimpleRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSocieties)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocietyServiceServer).SimpleRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protorender.SocietyService/SimpleRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocietyServiceServer).SimpleRPC(ctx, req.(*RequestSocieties))
	}
	return interceptor(ctx, in, info, handler)
}

func _SocietyService_ServerSideStreamingRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestSocieties)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SocietyServiceServer).ServerSideStreamingRPC(m, &societyServiceServerSideStreamingRPCServer{stream})
}

type SocietyService_ServerSideStreamingRPCServer interface {
	Send(*Society) error
	grpc.ServerStream
}

type societyServiceServerSideStreamingRPCServer struct {
	grpc.ServerStream
}

func (x *societyServiceServerSideStreamingRPCServer) Send(m *Society) error {
	return x.ServerStream.SendMsg(m)
}

func _SocietyService_ClientSideStreamingRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SocietyServiceServer).ClientSideStreamingRPC(&societyServiceClientSideStreamingRPCServer{stream})
}

type SocietyService_ClientSideStreamingRPCServer interface {
	SendAndClose(*ResponseSocieties) error
	Recv() (*RequestSocieties, error)
	grpc.ServerStream
}

type societyServiceClientSideStreamingRPCServer struct {
	grpc.ServerStream
}

func (x *societyServiceClientSideStreamingRPCServer) SendAndClose(m *ResponseSocieties) error {
	return x.ServerStream.SendMsg(m)
}

func (x *societyServiceClientSideStreamingRPCServer) Recv() (*RequestSocieties, error) {
	m := new(RequestSocieties)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SocietyService_BidirectionalStreamingRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SocietyServiceServer).BidirectionalStreamingRPC(&societyServiceBidirectionalStreamingRPCServer{stream})
}

type SocietyService_BidirectionalStreamingRPCServer interface {
	Send(*Society) error
	Recv() (*RequestSocieties, error)
	grpc.ServerStream
}

type societyServiceBidirectionalStreamingRPCServer struct {
	grpc.ServerStream
}

func (x *societyServiceBidirectionalStreamingRPCServer) Send(m *Society) error {
	return x.ServerStream.SendMsg(m)
}

func (x *societyServiceBidirectionalStreamingRPCServer) Recv() (*RequestSocieties, error) {
	m := new(RequestSocieties)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SocietyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protorender.SocietyService",
	HandlerType: (*SocietyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SimpleRPC",
			Handler:    _SocietyService_SimpleRPC_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerSideStreamingRPC",
			Handler:       _SocietyService_ServerSideStreamingRPC_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientSideStreamingRPC",
			Handler:       _SocietyService_ClientSideStreamingRPC_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamingRPC",
			Handler:       _SocietyService_BidirectionalStreamingRPC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "society.proto",
}