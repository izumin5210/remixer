// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: starwars/starship.proto

package starwars_pb

import (
	proto "github.com/golang/protobuf/proto"
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

type Starship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the starship
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the starship
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Starship) Reset() {
	*x = Starship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starship_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Starship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Starship) ProtoMessage() {}

func (x *Starship) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starship_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Starship.ProtoReflect.Descriptor instead.
func (*Starship) Descriptor() ([]byte, []int) {
	return file_starwars_starship_proto_rawDescGZIP(), []int{0}
}

func (x *Starship) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Starship) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetStarshipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetStarshipRequest) Reset() {
	*x = GetStarshipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_starship_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStarshipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStarshipRequest) ProtoMessage() {}

func (x *GetStarshipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_starship_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStarshipRequest.ProtoReflect.Descriptor instead.
func (*GetStarshipRequest) Descriptor() ([]byte, []int) {
	return file_starwars_starship_proto_rawDescGZIP(), []int{1}
}

func (x *GetStarshipRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_starwars_starship_proto protoreflect.FileDescriptor

var file_starwars_starship_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x53,
	0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x32, 0x62, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x12, 0x24, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x42, 0x1e, 0x5a, 0x1c, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x3b, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61,
	0x72, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_starwars_starship_proto_rawDescOnce sync.Once
	file_starwars_starship_proto_rawDescData = file_starwars_starship_proto_rawDesc
)

func file_starwars_starship_proto_rawDescGZIP() []byte {
	file_starwars_starship_proto_rawDescOnce.Do(func() {
		file_starwars_starship_proto_rawDescData = protoimpl.X.CompressGZIP(file_starwars_starship_proto_rawDescData)
	})
	return file_starwars_starship_proto_rawDescData
}

var file_starwars_starship_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_starwars_starship_proto_goTypes = []interface{}{
	(*Starship)(nil),           // 0: testapi.starwars.Starship
	(*GetStarshipRequest)(nil), // 1: testapi.starwars.GetStarshipRequest
}
var file_starwars_starship_proto_depIdxs = []int32{
	1, // 0: testapi.starwars.StarshipService.GetStarship:input_type -> testapi.starwars.GetStarshipRequest
	0, // 1: testapi.starwars.StarshipService.GetStarship:output_type -> testapi.starwars.Starship
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_starwars_starship_proto_init() }
func file_starwars_starship_proto_init() {
	if File_starwars_starship_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_starwars_starship_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Starship); i {
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
		file_starwars_starship_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStarshipRequest); i {
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
			RawDescriptor: file_starwars_starship_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_starwars_starship_proto_goTypes,
		DependencyIndexes: file_starwars_starship_proto_depIdxs,
		MessageInfos:      file_starwars_starship_proto_msgTypes,
	}.Build()
	File_starwars_starship_proto = out.File
	file_starwars_starship_proto_rawDesc = nil
	file_starwars_starship_proto_goTypes = nil
	file_starwars_starship_proto_depIdxs = nil
}
