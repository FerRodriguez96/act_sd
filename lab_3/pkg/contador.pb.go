// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pkg/contador.proto

package pkg

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

type Vacio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Vacio) Reset() {
	*x = Vacio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_contador_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vacio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vacio) ProtoMessage() {}

func (x *Vacio) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_contador_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vacio.ProtoReflect.Descriptor instead.
func (*Vacio) Descriptor() ([]byte, []int) {
	return file_pkg_contador_proto_rawDescGZIP(), []int{0}
}

type Valor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valor string `protobuf:"bytes,1,opt,name=Valor,proto3" json:"Valor,omitempty"`
}

func (x *Valor) Reset() {
	*x = Valor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_contador_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Valor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Valor) ProtoMessage() {}

func (x *Valor) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_contador_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Valor.ProtoReflect.Descriptor instead.
func (*Valor) Descriptor() ([]byte, []int) {
	return file_pkg_contador_proto_rawDescGZIP(), []int{1}
}

func (x *Valor) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

var File_pkg_contador_proto protoreflect.FileDescriptor

var file_pkg_contador_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x64, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x64, 0x6f, 0x72, 0x22, 0x07,
	0x0a, 0x05, 0x56, 0x61, 0x63, 0x69, 0x6f, 0x22, 0x1d, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x32, 0x0a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x64,
	0x6f, 0x72, 0x42, 0x0e, 0x5a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x64, 0x6f, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_contador_proto_rawDescOnce sync.Once
	file_pkg_contador_proto_rawDescData = file_pkg_contador_proto_rawDesc
)

func file_pkg_contador_proto_rawDescGZIP() []byte {
	file_pkg_contador_proto_rawDescOnce.Do(func() {
		file_pkg_contador_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_contador_proto_rawDescData)
	})
	return file_pkg_contador_proto_rawDescData
}

var file_pkg_contador_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_contador_proto_goTypes = []interface{}{
	(*Vacio)(nil), // 0: contador.Vacio
	(*Valor)(nil), // 1: contador.Valor
}
var file_pkg_contador_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_contador_proto_init() }
func file_pkg_contador_proto_init() {
	if File_pkg_contador_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_contador_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vacio); i {
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
		file_pkg_contador_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Valor); i {
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
			RawDescriptor: file_pkg_contador_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_contador_proto_goTypes,
		DependencyIndexes: file_pkg_contador_proto_depIdxs,
		MessageInfos:      file_pkg_contador_proto_msgTypes,
	}.Build()
	File_pkg_contador_proto = out.File
	file_pkg_contador_proto_rawDesc = nil
	file_pkg_contador_proto_goTypes = nil
	file_pkg_contador_proto_depIdxs = nil
}
