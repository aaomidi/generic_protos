// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: message.proto

package message

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

type Pet int32

const (
	Pet_CAT Pet = 0
	Pet_DOG Pet = 1
)

// Enum value maps for Pet.
var (
	Pet_name = map[int32]string{
		0: "CAT",
		1: "DOG",
	}
	Pet_value = map[string]int32{
		"CAT": 0,
		"DOG": 1,
	}
)

func (x Pet) Enum() *Pet {
	p := new(Pet)
	*p = x
	return p
}

func (x Pet) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Pet) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (Pet) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x Pet) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Pet.Descriptor instead.
func (Pet) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type MediaPlayer int32

const (
	MediaPlayer_PLEX     MediaPlayer = 0
	MediaPlayer_JELLYFIN MediaPlayer = 1
)

// Enum value maps for MediaPlayer.
var (
	MediaPlayer_name = map[int32]string{
		0: "PLEX",
		1: "JELLYFIN",
	}
	MediaPlayer_value = map[string]int32{
		"PLEX":     0,
		"JELLYFIN": 1,
	}
)

func (x MediaPlayer) Enum() *MediaPlayer {
	p := new(MediaPlayer)
	*p = x
	return p
}

func (x MediaPlayer) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaPlayer) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[1].Descriptor()
}

func (MediaPlayer) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[1]
}

func (x MediaPlayer) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaPlayer.Descriptor instead.
func (MediaPlayer) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x61, 0x6d, 0x69, 0x72, 0x5f, 0x70, 0x6b, 0x67, 0x1a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x2b, 0x0a, 0x03, 0x50, 0x65,
	0x74, 0x12, 0x11, 0x0a, 0x03, 0x43, 0x41, 0x54, 0x10, 0x00, 0x1a, 0x08, 0x9a, 0x80, 0xce, 0x04,
	0x03, 0x63, 0x61, 0x74, 0x12, 0x11, 0x0a, 0x03, 0x44, 0x4f, 0x47, 0x10, 0x01, 0x1a, 0x08, 0x9a,
	0x80, 0xce, 0x04, 0x03, 0x64, 0x6f, 0x67, 0x2a, 0x3e, 0x0a, 0x0b, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x4c, 0x45, 0x58, 0x10, 0x00,
	0x1a, 0x08, 0x9a, 0x80, 0xce, 0x04, 0x03, 0x70, 0x6c, 0x78, 0x12, 0x1b, 0x0a, 0x08, 0x4a, 0x45,
	0x4c, 0x4c, 0x59, 0x46, 0x49, 0x4e, 0x10, 0x01, 0x1a, 0x0d, 0x9a, 0x80, 0xce, 0x04, 0x08, 0x6e,
	0x6f, 0x74, 0x4a, 0x65, 0x6c, 0x6c, 0x79, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_message_proto_goTypes = []interface{}{
	(Pet)(0),         // 0: amir_pkg.Pet
	(MediaPlayer)(0), // 1: amir_pkg.MediaPlayer
}
var file_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	file_extension_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
