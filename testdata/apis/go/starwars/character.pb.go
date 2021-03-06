// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: starwars/character.proto

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

// A character from the Star Wars universe
type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Character:
	//	*Character_Human
	//	*Character_Droid
	Character isCharacter_Character `protobuf_oneof:"character"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_character_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_character_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_starwars_character_proto_rawDescGZIP(), []int{0}
}

func (m *Character) GetCharacter() isCharacter_Character {
	if m != nil {
		return m.Character
	}
	return nil
}

func (x *Character) GetHuman() *Human {
	if x, ok := x.GetCharacter().(*Character_Human); ok {
		return x.Human
	}
	return nil
}

func (x *Character) GetDroid() *Droid {
	if x, ok := x.GetCharacter().(*Character_Droid); ok {
		return x.Droid
	}
	return nil
}

type isCharacter_Character interface {
	isCharacter_Character()
}

type Character_Human struct {
	Human *Human `protobuf:"bytes,1,opt,name=human,proto3,oneof"`
}

type Character_Droid struct {
	Droid *Droid `protobuf:"bytes,2,opt,name=droid,proto3,oneof"`
}

func (*Character_Human) isCharacter_Character() {}

func (*Character_Droid) isCharacter_Character() {}

// A humanoid creature from the Star Wars universe
type Human struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the human
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// What this human calls themselves
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The home planet of the human, or null if unknown
	HomePlanet string `protobuf:"bytes,3,opt,name=homePlanet,proto3" json:"homePlanet,omitempty"`
	// Mass in kilograms, or null if unknown
	Mass float32 `protobuf:"fixed32,5,opt,name=mass,proto3" json:"mass,omitempty"`
	// This human's friends, or an empty list if they have none
	Friends []*Character `protobuf:"bytes,6,rep,name=friends,proto3" json:"friends,omitempty"`
	// The movies this human appears in
	AppearsIn []Episode `protobuf:"varint,8,rep,packed,name=appears_in,json=appearsIn,proto3,enum=testapi.starwars.Episode" json:"appears_in,omitempty"`
	// A list of starships this person has piloted, or an empty list if none
	Tarships []*Starship `protobuf:"bytes,9,rep,name=tarships,proto3" json:"tarships,omitempty"`
}

func (x *Human) Reset() {
	*x = Human{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_character_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Human) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Human) ProtoMessage() {}

func (x *Human) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_character_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Human.ProtoReflect.Descriptor instead.
func (*Human) Descriptor() ([]byte, []int) {
	return file_starwars_character_proto_rawDescGZIP(), []int{1}
}

func (x *Human) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Human) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Human) GetHomePlanet() string {
	if x != nil {
		return x.HomePlanet
	}
	return ""
}

func (x *Human) GetMass() float32 {
	if x != nil {
		return x.Mass
	}
	return 0
}

func (x *Human) GetFriends() []*Character {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *Human) GetAppearsIn() []Episode {
	if x != nil {
		return x.AppearsIn
	}
	return nil
}

func (x *Human) GetTarships() []*Starship {
	if x != nil {
		return x.Tarships
	}
	return nil
}

// An autonomous mechanical character in the Star Wars universe
type Droid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the droid
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// What others call this droid
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// This droid's friends, or an empty list if they have none
	Friends []*Character `protobuf:"bytes,3,rep,name=friends,proto3" json:"friends,omitempty"`
	// The movies this droid appears in
	AppearsIn []Episode `protobuf:"varint,5,rep,packed,name=appears_in,json=appearsIn,proto3,enum=testapi.starwars.Episode" json:"appears_in,omitempty"`
	// This droid's primary function
	PrimaryFunction string `protobuf:"bytes,6,opt,name=primary_function,json=primaryFunction,proto3" json:"primary_function,omitempty"`
}

func (x *Droid) Reset() {
	*x = Droid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_character_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Droid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Droid) ProtoMessage() {}

func (x *Droid) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_character_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Droid.ProtoReflect.Descriptor instead.
func (*Droid) Descriptor() ([]byte, []int) {
	return file_starwars_character_proto_rawDescGZIP(), []int{2}
}

func (x *Droid) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Droid) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Droid) GetFriends() []*Character {
	if x != nil {
		return x.Friends
	}
	return nil
}

func (x *Droid) GetAppearsIn() []Episode {
	if x != nil {
		return x.AppearsIn
	}
	return nil
}

func (x *Droid) GetPrimaryFunction() string {
	if x != nil {
		return x.PrimaryFunction
	}
	return ""
}

type GetCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the character
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCharacterRequest) Reset() {
	*x = GetCharacterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_character_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharacterRequest) ProtoMessage() {}

func (x *GetCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_character_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharacterRequest.ProtoReflect.Descriptor instead.
func (*GetCharacterRequest) Descriptor() ([]byte, []int) {
	return file_starwars_character_proto_rawDescGZIP(), []int{3}
}

func (x *GetCharacterRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetHumanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the human
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHumanRequest) Reset() {
	*x = GetHumanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_character_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHumanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHumanRequest) ProtoMessage() {}

func (x *GetHumanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_character_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHumanRequest.ProtoReflect.Descriptor instead.
func (*GetHumanRequest) Descriptor() ([]byte, []int) {
	return file_starwars_character_proto_rawDescGZIP(), []int{4}
}

func (x *GetHumanRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDroidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the droid
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDroidRequest) Reset() {
	*x = GetDroidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_starwars_character_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDroidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDroidRequest) ProtoMessage() {}

func (x *GetDroidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_starwars_character_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDroidRequest.ProtoReflect.Descriptor instead.
func (*GetDroidRequest) Descriptor() ([]byte, []int) {
	return file_starwars_character_proto_rawDescGZIP(), []int{5}
}

func (x *GetDroidRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_starwars_character_proto protoreflect.FileDescriptor

var file_starwars_character_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x1a, 0x17, 0x73, 0x74,
	0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x09, 0x43, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x05, 0x68, 0x75, 0x6d, 0x61, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x48,
	0x00, 0x52, 0x05, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x12, 0x2f, 0x0a, 0x05, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x44, 0x72, 0x6f, 0x69, 0x64,
	0x48, 0x00, 0x52, 0x05, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x22, 0x88, 0x02, 0x0a, 0x05, 0x48, 0x75, 0x6d, 0x61, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12,
	0x38, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x52, 0x09,
	0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x73, 0x49, 0x6e, 0x12, 0x36, 0x0a, 0x08, 0x74, 0x61, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x08, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x22, 0xc7, 0x01, 0x0a, 0x05, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77,
	0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72,
	0x73, 0x5f, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x45, 0x70,
	0x69, 0x73, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x73, 0x49, 0x6e,
	0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x72, 0x6f, 0x69,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0xf6, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x25, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x12, 0x46, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x12, 0x21, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77,
	0x61, 0x72, 0x73, 0x2e, 0x44, 0x72, 0x6f, 0x69, 0x64, 0x12, 0x46, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x48, 0x75, 0x6d, 0x61, 0x6e, 0x12, 0x21, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x75, 0x6d, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x48, 0x75, 0x6d, 0x61,
	0x6e, 0x42, 0x1e, 0x5a, 0x1c, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x61,
	0x72, 0x77, 0x61, 0x72, 0x73, 0x3b, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x5f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_starwars_character_proto_rawDescOnce sync.Once
	file_starwars_character_proto_rawDescData = file_starwars_character_proto_rawDesc
)

func file_starwars_character_proto_rawDescGZIP() []byte {
	file_starwars_character_proto_rawDescOnce.Do(func() {
		file_starwars_character_proto_rawDescData = protoimpl.X.CompressGZIP(file_starwars_character_proto_rawDescData)
	})
	return file_starwars_character_proto_rawDescData
}

var file_starwars_character_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_starwars_character_proto_goTypes = []interface{}{
	(*Character)(nil),           // 0: testapi.starwars.Character
	(*Human)(nil),               // 1: testapi.starwars.Human
	(*Droid)(nil),               // 2: testapi.starwars.Droid
	(*GetCharacterRequest)(nil), // 3: testapi.starwars.GetCharacterRequest
	(*GetHumanRequest)(nil),     // 4: testapi.starwars.GetHumanRequest
	(*GetDroidRequest)(nil),     // 5: testapi.starwars.GetDroidRequest
	(Episode)(0),                // 6: testapi.starwars.Episode
	(*Starship)(nil),            // 7: testapi.starwars.Starship
}
var file_starwars_character_proto_depIdxs = []int32{
	1,  // 0: testapi.starwars.Character.human:type_name -> testapi.starwars.Human
	2,  // 1: testapi.starwars.Character.droid:type_name -> testapi.starwars.Droid
	0,  // 2: testapi.starwars.Human.friends:type_name -> testapi.starwars.Character
	6,  // 3: testapi.starwars.Human.appears_in:type_name -> testapi.starwars.Episode
	7,  // 4: testapi.starwars.Human.tarships:type_name -> testapi.starwars.Starship
	0,  // 5: testapi.starwars.Droid.friends:type_name -> testapi.starwars.Character
	6,  // 6: testapi.starwars.Droid.appears_in:type_name -> testapi.starwars.Episode
	3,  // 7: testapi.starwars.CharacterService.GetCharacter:input_type -> testapi.starwars.GetCharacterRequest
	5,  // 8: testapi.starwars.CharacterService.GetDroid:input_type -> testapi.starwars.GetDroidRequest
	4,  // 9: testapi.starwars.CharacterService.GetHuman:input_type -> testapi.starwars.GetHumanRequest
	0,  // 10: testapi.starwars.CharacterService.GetCharacter:output_type -> testapi.starwars.Character
	2,  // 11: testapi.starwars.CharacterService.GetDroid:output_type -> testapi.starwars.Droid
	1,  // 12: testapi.starwars.CharacterService.GetHuman:output_type -> testapi.starwars.Human
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_starwars_character_proto_init() }
func file_starwars_character_proto_init() {
	if File_starwars_character_proto != nil {
		return
	}
	file_starwars_starship_proto_init()
	file_starwars_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_starwars_character_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_starwars_character_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Human); i {
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
		file_starwars_character_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Droid); i {
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
		file_starwars_character_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCharacterRequest); i {
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
		file_starwars_character_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHumanRequest); i {
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
		file_starwars_character_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDroidRequest); i {
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
	file_starwars_character_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Character_Human)(nil),
		(*Character_Droid)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_starwars_character_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_starwars_character_proto_goTypes,
		DependencyIndexes: file_starwars_character_proto_depIdxs,
		MessageInfos:      file_starwars_character_proto_msgTypes,
	}.Build()
	File_starwars_character_proto = out.File
	file_starwars_character_proto_rawDesc = nil
	file_starwars_character_proto_goTypes = nil
	file_starwars_character_proto_depIdxs = nil
}
