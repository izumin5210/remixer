// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: starwars/type.proto

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

// The episodes in the Star Wars trilogy
type Episode int32

const (
	Episode_EPISODE_UNSPECIFIED Episode = 0
	// Star Wars Episode IV: A New Hope, released in 1977.
	Episode_NEWHOPE Episode = 1
	// Star Wars Episode V: The Empire Strikes Back, released in 1980.
	Episode_EMPIRE Episode = 2
	// Star Wars Episode VI: Return of the Jedi, released in 1983.
	Episode_JEDI Episode = 3
)

// Enum value maps for Episode.
var (
	Episode_name = map[int32]string{
		0: "EPISODE_UNSPECIFIED",
		1: "NEWHOPE",
		2: "EMPIRE",
		3: "JEDI",
	}
	Episode_value = map[string]int32{
		"EPISODE_UNSPECIFIED": 0,
		"NEWHOPE":             1,
		"EMPIRE":              2,
		"JEDI":                3,
	}
)

func (x Episode) Enum() *Episode {
	p := new(Episode)
	*p = x
	return p
}

func (x Episode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Episode) Descriptor() protoreflect.EnumDescriptor {
	return file_starwars_type_proto_enumTypes[0].Descriptor()
}

func (Episode) Type() protoreflect.EnumType {
	return &file_starwars_type_proto_enumTypes[0]
}

func (x Episode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Episode.Descriptor instead.
func (Episode) EnumDescriptor() ([]byte, []int) {
	return file_starwars_type_proto_rawDescGZIP(), []int{0}
}

// Units of height
type LengthUnit int32

const (
	LengthUnit_LENGTH_UNIT_UNSPECIFIED LengthUnit = 0
	// The standard unit around the world
	LengthUnit_METER LengthUnit = 1
	// Primarily used in the United States
	LengthUnit_FOOT LengthUnit = 2
	// DEPRECATED. Test deprecated enum case.
	// Ancient unit used during the Middle Ages
	//
	// Deprecated: Do not use.
	LengthUnit_CUBIT LengthUnit = 3
)

// Enum value maps for LengthUnit.
var (
	LengthUnit_name = map[int32]string{
		0: "LENGTH_UNIT_UNSPECIFIED",
		1: "METER",
		2: "FOOT",
		3: "CUBIT",
	}
	LengthUnit_value = map[string]int32{
		"LENGTH_UNIT_UNSPECIFIED": 0,
		"METER":                   1,
		"FOOT":                    2,
		"CUBIT":                   3,
	}
)

func (x LengthUnit) Enum() *LengthUnit {
	p := new(LengthUnit)
	*p = x
	return p
}

func (x LengthUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LengthUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_starwars_type_proto_enumTypes[1].Descriptor()
}

func (LengthUnit) Type() protoreflect.EnumType {
	return &file_starwars_type_proto_enumTypes[1]
}

func (x LengthUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LengthUnit.Descriptor instead.
func (LengthUnit) EnumDescriptor() ([]byte, []int) {
	return file_starwars_type_proto_rawDescGZIP(), []int{1}
}

var File_starwars_type_proto protoreflect.FileDescriptor

var file_starwars_type_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2a, 0x45, 0x0a, 0x07, 0x45, 0x70, 0x69, 0x73, 0x6f,
	0x64, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x50, 0x49, 0x53, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4e,
	0x45, 0x57, 0x48, 0x4f, 0x50, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x4d, 0x50, 0x49,
	0x52, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x45, 0x44, 0x49, 0x10, 0x03, 0x2a, 0x4d,
	0x0a, 0x0a, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x17,
	0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x45, 0x54,
	0x45, 0x52, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x4f, 0x4f, 0x54, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x05, 0x43, 0x55, 0x42, 0x49, 0x54, 0x10, 0x03, 0x1a, 0x02, 0x08, 0x01, 0x42, 0x1e, 0x5a,
	0x1c, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72,
	0x73, 0x3b, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_starwars_type_proto_rawDescOnce sync.Once
	file_starwars_type_proto_rawDescData = file_starwars_type_proto_rawDesc
)

func file_starwars_type_proto_rawDescGZIP() []byte {
	file_starwars_type_proto_rawDescOnce.Do(func() {
		file_starwars_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_starwars_type_proto_rawDescData)
	})
	return file_starwars_type_proto_rawDescData
}

var file_starwars_type_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_starwars_type_proto_goTypes = []interface{}{
	(Episode)(0),    // 0: testapi.starwars.Episode
	(LengthUnit)(0), // 1: testapi.starwars.LengthUnit
}
var file_starwars_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_starwars_type_proto_init() }
func file_starwars_type_proto_init() {
	if File_starwars_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_starwars_type_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_starwars_type_proto_goTypes,
		DependencyIndexes: file_starwars_type_proto_depIdxs,
		EnumInfos:         file_starwars_type_proto_enumTypes,
	}.Build()
	File_starwars_type_proto = out.File
	file_starwars_type_proto_rawDesc = nil
	file_starwars_type_proto_goTypes = nil
	file_starwars_type_proto_depIdxs = nil
}
