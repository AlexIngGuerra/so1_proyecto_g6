// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: protoc/comunicacion.proto

package protoc

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

// The request message containing the user's name.
type IngresoSolicitud struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Team1 string `protobuf:"bytes,1,opt,name=team1,proto3" json:"team1,omitempty"`
	Team2 string `protobuf:"bytes,2,opt,name=team2,proto3" json:"team2,omitempty"`
	Score string `protobuf:"bytes,3,opt,name=score,proto3" json:"score,omitempty"`
	Phase string `protobuf:"bytes,4,opt,name=phase,proto3" json:"phase,omitempty"`
}

func (x *IngresoSolicitud) Reset() {
	*x = IngresoSolicitud{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_comunicacion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngresoSolicitud) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngresoSolicitud) ProtoMessage() {}

func (x *IngresoSolicitud) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_comunicacion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngresoSolicitud.ProtoReflect.Descriptor instead.
func (*IngresoSolicitud) Descriptor() ([]byte, []int) {
	return file_protoc_comunicacion_proto_rawDescGZIP(), []int{0}
}

func (x *IngresoSolicitud) GetTeam1() string {
	if x != nil {
		return x.Team1
	}
	return ""
}

func (x *IngresoSolicitud) GetTeam2() string {
	if x != nil {
		return x.Team2
	}
	return ""
}

func (x *IngresoSolicitud) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *IngresoSolicitud) GetPhase() string {
	if x != nil {
		return x.Phase
	}
	return ""
}

// The response message containing the greetings
type Respuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codigo  string `protobuf:"bytes,1,opt,name=codigo,proto3" json:"codigo,omitempty"`
	Mensaje string `protobuf:"bytes,2,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
}

func (x *Respuesta) Reset() {
	*x = Respuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_comunicacion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Respuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Respuesta) ProtoMessage() {}

func (x *Respuesta) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_comunicacion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Respuesta.ProtoReflect.Descriptor instead.
func (*Respuesta) Descriptor() ([]byte, []int) {
	return file_protoc_comunicacion_proto_rawDescGZIP(), []int{1}
}

func (x *Respuesta) GetCodigo() string {
	if x != nil {
		return x.Codigo
	}
	return ""
}

func (x *Respuesta) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

var File_protoc_comunicacion_proto protoreflect.FileDescriptor

var file_protoc_comunicacion_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x63, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x6a, 0x0a, 0x10, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x6f, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x65, 0x61, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x61, 0x6d,
	0x31, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x61, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x73,
	0x61, 0x6a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x73, 0x61,
	0x6a, 0x65, 0x32, 0x50, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x45, 0x0a,
	0x0c, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x6f, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x1c, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x6f, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x1a, 0x15, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73,
	0x74, 0x61, 0x22, 0x00, 0x42, 0x40, 0x0a, 0x0e, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x63, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x42, 0x0d, 0x50, 0x72, 0x6f, 0x79, 0x65, 0x63, 0x74, 0x6f,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x72, 0x63, 0x70, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoc_comunicacion_proto_rawDescOnce sync.Once
	file_protoc_comunicacion_proto_rawDescData = file_protoc_comunicacion_proto_rawDesc
)

func file_protoc_comunicacion_proto_rawDescGZIP() []byte {
	file_protoc_comunicacion_proto_rawDescOnce.Do(func() {
		file_protoc_comunicacion_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoc_comunicacion_proto_rawDescData)
	})
	return file_protoc_comunicacion_proto_rawDescData
}

var file_protoc_comunicacion_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protoc_comunicacion_proto_goTypes = []interface{}{
	(*IngresoSolicitud)(nil), // 0: helloworld.IngresoSolicitud
	(*Respuesta)(nil),        // 1: helloworld.Respuesta
}
var file_protoc_comunicacion_proto_depIdxs = []int32{
	0, // 0: helloworld.Greeter.IngresoDatos:input_type -> helloworld.IngresoSolicitud
	1, // 1: helloworld.Greeter.IngresoDatos:output_type -> helloworld.Respuesta
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoc_comunicacion_proto_init() }
func file_protoc_comunicacion_proto_init() {
	if File_protoc_comunicacion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoc_comunicacion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngresoSolicitud); i {
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
		file_protoc_comunicacion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Respuesta); i {
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
			RawDescriptor: file_protoc_comunicacion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoc_comunicacion_proto_goTypes,
		DependencyIndexes: file_protoc_comunicacion_proto_depIdxs,
		MessageInfos:      file_protoc_comunicacion_proto_msgTypes,
	}.Build()
	File_protoc_comunicacion_proto = out.File
	file_protoc_comunicacion_proto_rawDesc = nil
	file_protoc_comunicacion_proto_goTypes = nil
	file_protoc_comunicacion_proto_depIdxs = nil
}
