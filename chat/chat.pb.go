// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: chat.proto

package chat

import (
	context "context"
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

type Propuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Propuesta) Reset() {
	*x = Propuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Propuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Propuesta) ProtoMessage() {}

func (x *Propuesta) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Propuesta.ProtoReflect.Descriptor instead.
func (*Propuesta) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Propuesta) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Respuesta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Respuesta) Reset() {
	*x = Respuesta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Respuesta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Respuesta) ProtoMessage() {}

func (x *Respuesta) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
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
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Respuesta) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre    string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Parte     string `protobuf:"bytes,2,opt,name=parte,proto3" json:"parte,omitempty"`
	NumPartes uint64 `protobuf:"varint,3,opt,name=numPartes,proto3" json:"numPartes,omitempty"`
	Buffer    []byte `protobuf:"bytes,4,opt,name=buffer,proto3" json:"buffer,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Message) GetParte() string {
	if x != nil {
		return x.Parte
	}
	return ""
}

func (x *Message) GetNumPartes() uint64 {
	if x != nil {
		return x.NumPartes
	}
	return 0
}

func (x *Message) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre    string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Parte     string `protobuf:"bytes,2,opt,name=parte,proto3" json:"parte,omitempty"`
	NumPartes uint64 `protobuf:"varint,3,opt,name=numPartes,proto3" json:"numPartes,omitempty"`
	Buffer    []byte `protobuf:"bytes,4,opt,name=buffer,proto3" json:"buffer,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *Chunk) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Chunk) GetParte() string {
	if x != nil {
		return x.Parte
	}
	return ""
}

func (x *Chunk) GetNumPartes() uint64 {
	if x != nil {
		return x.NumPartes
	}
	return 0
}

func (x *Chunk) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0x1f, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x22, 0x1f, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x22, 0x6d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x6e, 0x75, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x22, 0x6b, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75,
	0x6d, 0x50, 0x61, 0x72, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e,
	0x75, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x32, 0x88, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2d, 0x0a, 0x07, 0x53, 0x61, 0x79, 0x48, 0x6f, 0x6c, 0x61, 0x12, 0x0f, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x22, 0x00, 0x12,
	0x2c, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x69, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x0b, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0f, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x0f, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x12, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x1a, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73,
	0x74, 0x61, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0b, 0x50, 0x65, 0x64, 0x69, 0x72, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75,
	0x65, 0x73, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x75, 0x65, 0x73, 0x74, 0x61, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x61,
	0x72, 0x67, 0x61, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chat_proto_goTypes = []interface{}{
	(*Propuesta)(nil), // 0: chat.Propuesta
	(*Respuesta)(nil), // 1: chat.Respuesta
	(*Message)(nil),   // 2: chat.Message
	(*Chunk)(nil),     // 3: chat.Chunk
}
var file_chat_proto_depIdxs = []int32{
	1, // 0: chat.ChatService.SayHola:input_type -> chat.Respuesta
	3, // 1: chat.ChatService.SubirChunk:input_type -> chat.Chunk
	0, // 2: chat.ChatService.EnviarPropuesta:input_type -> chat.Propuesta
	1, // 3: chat.ChatService.PedirChunks:input_type -> chat.Respuesta
	1, // 4: chat.ChatService.DescargarChunk:input_type -> chat.Respuesta
	1, // 5: chat.ChatService.SayHola:output_type -> chat.Respuesta
	1, // 6: chat.ChatService.SubirChunk:output_type -> chat.Respuesta
	1, // 7: chat.ChatService.EnviarPropuesta:output_type -> chat.Respuesta
	1, // 8: chat.ChatService.PedirChunks:output_type -> chat.Respuesta
	2, // 9: chat.ChatService.DescargarChunk:output_type -> chat.Message
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Propuesta); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	SayHola(ctx context.Context, in *Respuesta, opts ...grpc.CallOption) (*Respuesta, error)
	SubirChunk(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Respuesta, error)
	//rpc ConfirmarPropuestas(Propuesta)returns (Respuesta){}
	EnviarPropuesta(ctx context.Context, in *Propuesta, opts ...grpc.CallOption) (*Respuesta, error)
	PedirChunks(ctx context.Context, in *Respuesta, opts ...grpc.CallOption) (*Respuesta, error)
	DescargarChunk(ctx context.Context, in *Respuesta, opts ...grpc.CallOption) (*Message, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SayHola(ctx context.Context, in *Respuesta, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SayHola", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SubirChunk(ctx context.Context, in *Chunk, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SubirChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) EnviarPropuesta(ctx context.Context, in *Propuesta, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/chat.ChatService/EnviarPropuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) PedirChunks(ctx context.Context, in *Respuesta, opts ...grpc.CallOption) (*Respuesta, error) {
	out := new(Respuesta)
	err := c.cc.Invoke(ctx, "/chat.ChatService/PedirChunks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DescargarChunk(ctx context.Context, in *Respuesta, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/DescargarChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	SayHola(context.Context, *Respuesta) (*Respuesta, error)
	SubirChunk(context.Context, *Chunk) (*Respuesta, error)
	//rpc ConfirmarPropuestas(Propuesta)returns (Respuesta){}
	EnviarPropuesta(context.Context, *Propuesta) (*Respuesta, error)
	PedirChunks(context.Context, *Respuesta) (*Respuesta, error)
	DescargarChunk(context.Context, *Respuesta) (*Message, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) SayHola(context.Context, *Respuesta) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHola not implemented")
}
func (*UnimplementedChatServiceServer) SubirChunk(context.Context, *Chunk) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubirChunk not implemented")
}
func (*UnimplementedChatServiceServer) EnviarPropuesta(context.Context, *Propuesta) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarPropuesta not implemented")
}
func (*UnimplementedChatServiceServer) PedirChunks(context.Context, *Respuesta) (*Respuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirChunks not implemented")
}
func (*UnimplementedChatServiceServer) DescargarChunk(context.Context, *Respuesta) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescargarChunk not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_SayHola_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Respuesta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SayHola(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SayHola",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SayHola(ctx, req.(*Respuesta))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SubirChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SubirChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SubirChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SubirChunk(ctx, req.(*Chunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_EnviarPropuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Propuesta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).EnviarPropuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/EnviarPropuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).EnviarPropuesta(ctx, req.(*Propuesta))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_PedirChunks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Respuesta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).PedirChunks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/PedirChunks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).PedirChunks(ctx, req.(*Respuesta))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DescargarChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Respuesta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DescargarChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/DescargarChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DescargarChunk(ctx, req.(*Respuesta))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHola",
			Handler:    _ChatService_SayHola_Handler,
		},
		{
			MethodName: "SubirChunk",
			Handler:    _ChatService_SubirChunk_Handler,
		},
		{
			MethodName: "EnviarPropuesta",
			Handler:    _ChatService_EnviarPropuesta_Handler,
		},
		{
			MethodName: "PedirChunks",
			Handler:    _ChatService_PedirChunks_Handler,
		},
		{
			MethodName: "DescargarChunk",
			Handler:    _ChatService_DescargarChunk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}