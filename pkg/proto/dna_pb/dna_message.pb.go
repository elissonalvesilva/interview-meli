// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: dna_message.proto

package dna_pb

import (
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

type DNARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dna []string `protobuf:"bytes,1,rep,name=dna,proto3" json:"dna,omitempty"`
}

func (x *DNARequest) Reset() {
	*x = DNARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dna_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNARequest) ProtoMessage() {}

func (x *DNARequest) ProtoReflect() protoreflect.Message {
	mi := &file_dna_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNARequest.ProtoReflect.Descriptor instead.
func (*DNARequest) Descriptor() ([]byte, []int) {
	return file_dna_message_proto_rawDescGZIP(), []int{0}
}

func (x *DNARequest) GetDna() []string {
	if x != nil {
		return x.Dna
	}
	return nil
}

type DNAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeDNA  string `protobuf:"bytes,1,opt,name=typeDNA,proto3" json:"typeDNA,omitempty"`
	DnaError string `protobuf:"bytes,2,opt,name=dnaError,proto3" json:"dnaError,omitempty"`
}

func (x *DNAResponse) Reset() {
	*x = DNAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dna_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNAResponse) ProtoMessage() {}

func (x *DNAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dna_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNAResponse.ProtoReflect.Descriptor instead.
func (*DNAResponse) Descriptor() ([]byte, []int) {
	return file_dna_message_proto_rawDescGZIP(), []int{1}
}

func (x *DNAResponse) GetTypeDNA() string {
	if x != nil {
		return x.TypeDNA
	}
	return ""
}

func (x *DNAResponse) GetDnaError() string {
	if x != nil {
		return x.DnaError
	}
	return ""
}

var File_dna_message_proto protoreflect.FileDescriptor

var file_dna_message_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x6e, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x6e, 0x61, 0x5f, 0x70, 0x62, 0x22, 0x1e, 0x0a, 0x0a, 0x44,
	0x4e, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6e, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6e, 0x61, 0x22, 0x43, 0x0a, 0x0b, 0x44,
	0x4e, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x79,
	0x70, 0x65, 0x44, 0x4e, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70,
	0x65, 0x44, 0x4e, 0x41, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6e, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6e, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0x4a, 0x0a, 0x12, 0x44, 0x4e, 0x41, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x12, 0x12, 0x2e, 0x64, 0x6e, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x44, 0x4e, 0x41, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x64, 0x6e, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x44,
	0x4e, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6e, 0x61, 0x5f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dna_message_proto_rawDescOnce sync.Once
	file_dna_message_proto_rawDescData = file_dna_message_proto_rawDesc
)

func file_dna_message_proto_rawDescGZIP() []byte {
	file_dna_message_proto_rawDescOnce.Do(func() {
		file_dna_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_dna_message_proto_rawDescData)
	})
	return file_dna_message_proto_rawDescData
}

var file_dna_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dna_message_proto_goTypes = []interface{}{
	(*DNARequest)(nil),  // 0: dna_pb.DNARequest
	(*DNAResponse)(nil), // 1: dna_pb.DNAResponse
}
var file_dna_message_proto_depIdxs = []int32{
	0, // 0: dna_pb.DNAServiceAnalyzer.Analyze:input_type -> dna_pb.DNARequest
	1, // 1: dna_pb.DNAServiceAnalyzer.Analyze:output_type -> dna_pb.DNAResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dna_message_proto_init() }
func file_dna_message_proto_init() {
	if File_dna_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dna_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNARequest); i {
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
		file_dna_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNAResponse); i {
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
			RawDescriptor: file_dna_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dna_message_proto_goTypes,
		DependencyIndexes: file_dna_message_proto_depIdxs,
		MessageInfos:      file_dna_message_proto_msgTypes,
	}.Build()
	File_dna_message_proto = out.File
	file_dna_message_proto_rawDesc = nil
	file_dna_message_proto_goTypes = nil
	file_dna_message_proto_depIdxs = nil
}
