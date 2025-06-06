// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: aliasreader/aliasreader.proto

package aliasreader

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

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	mi := &file_aliasreader_aliasreader_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_aliasreader_aliasreader_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_aliasreader_aliasreader_proto_rawDescGZIP(), []int{0}
}

func (x *ID) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type Alias struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alias string `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *Alias) Reset() {
	*x = Alias{}
	mi := &file_aliasreader_aliasreader_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Alias) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alias) ProtoMessage() {}

func (x *Alias) ProtoReflect() protoreflect.Message {
	mi := &file_aliasreader_aliasreader_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alias.ProtoReflect.Descriptor instead.
func (*Alias) Descriptor() ([]byte, []int) {
	return file_aliasreader_aliasreader_proto_rawDescGZIP(), []int{1}
}

func (x *Alias) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

type AliasList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aliases []string `protobuf:"bytes,1,rep,name=aliases,proto3" json:"aliases,omitempty"`
}

func (x *AliasList) Reset() {
	*x = AliasList{}
	mi := &file_aliasreader_aliasreader_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AliasList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AliasList) ProtoMessage() {}

func (x *AliasList) ProtoReflect() protoreflect.Message {
	mi := &file_aliasreader_aliasreader_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AliasList.ProtoReflect.Descriptor instead.
func (*AliasList) Descriptor() ([]byte, []int) {
	return file_aliasreader_aliasreader_proto_rawDescGZIP(), []int{2}
}

func (x *AliasList) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

var File_aliasreader_aliasreader_proto protoreflect.FileDescriptor

var file_aliasreader_aliasreader_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x14, 0x0a, 0x02,
	0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x1d, 0x0a, 0x05, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x22, 0x25, 0x0a, 0x09, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x32, 0xa5, 0x01, 0x0a, 0x0b, 0x41, 0x6c, 0x69,
	0x61, 0x73, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x12, 0x12, 0x2e, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x2e, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x1a, 0x0f, 0x2e, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x72, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x0f, 0x2e, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x72,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x44, 0x1a, 0x12, 0x2e, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x32, 0x0a, 0x07,
	0x41, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x72,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x76, 0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x68,
	0x65, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aliasreader_aliasreader_proto_rawDescOnce sync.Once
	file_aliasreader_aliasreader_proto_rawDescData = file_aliasreader_aliasreader_proto_rawDesc
)

func file_aliasreader_aliasreader_proto_rawDescGZIP() []byte {
	file_aliasreader_aliasreader_proto_rawDescOnce.Do(func() {
		file_aliasreader_aliasreader_proto_rawDescData = protoimpl.X.CompressGZIP(file_aliasreader_aliasreader_proto_rawDescData)
	})
	return file_aliasreader_aliasreader_proto_rawDescData
}

var file_aliasreader_aliasreader_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_aliasreader_aliasreader_proto_goTypes = []any{
	(*ID)(nil),        // 0: aliasreader.ID
	(*Alias)(nil),     // 1: aliasreader.Alias
	(*AliasList)(nil), // 2: aliasreader.AliasList
}
var file_aliasreader_aliasreader_proto_depIdxs = []int32{
	1, // 0: aliasreader.AliasReader.Lookup:input_type -> aliasreader.Alias
	0, // 1: aliasreader.AliasReader.PrimaryAlias:input_type -> aliasreader.ID
	0, // 2: aliasreader.AliasReader.Aliases:input_type -> aliasreader.ID
	0, // 3: aliasreader.AliasReader.Lookup:output_type -> aliasreader.ID
	1, // 4: aliasreader.AliasReader.PrimaryAlias:output_type -> aliasreader.Alias
	2, // 5: aliasreader.AliasReader.Aliases:output_type -> aliasreader.AliasList
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aliasreader_aliasreader_proto_init() }
func file_aliasreader_aliasreader_proto_init() {
	if File_aliasreader_aliasreader_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aliasreader_aliasreader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aliasreader_aliasreader_proto_goTypes,
		DependencyIndexes: file_aliasreader_aliasreader_proto_depIdxs,
		MessageInfos:      file_aliasreader_aliasreader_proto_msgTypes,
	}.Build()
	File_aliasreader_aliasreader_proto = out.File
	file_aliasreader_aliasreader_proto_rawDesc = nil
	file_aliasreader_aliasreader_proto_goTypes = nil
	file_aliasreader_aliasreader_proto_depIdxs = nil
}
