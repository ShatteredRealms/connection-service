// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: sro/gameserver/connection.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ConnectGameServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port         uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	ConnectionId string `protobuf:"bytes,3,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
}

func (x *ConnectGameServerResponse) Reset() {
	*x = ConnectGameServerResponse{}
	mi := &file_sro_gameserver_connection_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectGameServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectGameServerResponse) ProtoMessage() {}

func (x *ConnectGameServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_connection_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectGameServerResponse.ProtoReflect.Descriptor instead.
func (*ConnectGameServerResponse) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_connection_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectGameServerResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ConnectGameServerResponse) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ConnectGameServerResponse) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

type VerifyConnectRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionId string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	ServerName   string `protobuf:"bytes,2,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
}

func (x *VerifyConnectRequestMessage) Reset() {
	*x = VerifyConnectRequestMessage{}
	mi := &file_sro_gameserver_connection_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyConnectRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyConnectRequestMessage) ProtoMessage() {}

func (x *VerifyConnectRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_connection_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyConnectRequestMessage.ProtoReflect.Descriptor instead.
func (*VerifyConnectRequestMessage) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_connection_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyConnectRequestMessage) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *VerifyConnectRequestMessage) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

type ConnectionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Online bool `protobuf:"varint,1,opt,name=online,proto3" json:"online,omitempty"`
}

func (x *ConnectionStatus) Reset() {
	*x = ConnectionStatus{}
	mi := &file_sro_gameserver_connection_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionStatus) ProtoMessage() {}

func (x *ConnectionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_connection_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionStatus.ProtoReflect.Descriptor instead.
func (*ConnectionStatus) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_connection_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectionStatus) GetOnline() bool {
	if x != nil {
		return x.Online
	}
	return false
}

type TransferPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Character string    `protobuf:"bytes,1,opt,name=character,proto3" json:"character,omitempty"`
	Location  *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *TransferPlayerRequest) Reset() {
	*x = TransferPlayerRequest{}
	mi := &file_sro_gameserver_connection_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferPlayerRequest) ProtoMessage() {}

func (x *TransferPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sro_gameserver_connection_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferPlayerRequest.ProtoReflect.Descriptor instead.
func (*TransferPlayerRequest) Descriptor() ([]byte, []int) {
	return file_sro_gameserver_connection_proto_rawDescGZIP(), []int{3}
}

func (x *TransferPlayerRequest) GetCharacter() string {
	if x != nil {
		return x.Character
	}
	return ""
}

func (x *TransferPlayerRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_sro_gameserver_connection_proto protoreflect.FileDescriptor

var file_sro_gameserver_connection_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x72, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x73, 0x72, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x63, 0x0a, 0x1b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0x60, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73,
	0x72, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xc0, 0x03, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x11,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x14, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x44, 0x5a, 0x23, 0x12, 0x21, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12,
	0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x78,
	0x0a, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x8d, 0x01, 0x0a, 0x09, 0x49, 0x73, 0x50,
	0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x73, 0x72, 0x6f, 0x2e, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x20, 0x2e, 0x73,
	0x72, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x48,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x42, 0x5a, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2f, 0x6e,
	0x61, 0x6d, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x2f, 0x69, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sro_gameserver_connection_proto_rawDescOnce sync.Once
	file_sro_gameserver_connection_proto_rawDescData = file_sro_gameserver_connection_proto_rawDesc
)

func file_sro_gameserver_connection_proto_rawDescGZIP() []byte {
	file_sro_gameserver_connection_proto_rawDescOnce.Do(func() {
		file_sro_gameserver_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_sro_gameserver_connection_proto_rawDescData)
	})
	return file_sro_gameserver_connection_proto_rawDescData
}

var file_sro_gameserver_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sro_gameserver_connection_proto_goTypes = []any{
	(*ConnectGameServerResponse)(nil),   // 0: sro.gameserver.ConnectGameServerResponse
	(*VerifyConnectRequestMessage)(nil), // 1: sro.gameserver.VerifyConnectRequestMessage
	(*ConnectionStatus)(nil),            // 2: sro.gameserver.ConnectionStatus
	(*TransferPlayerRequest)(nil),       // 3: sro.gameserver.TransferPlayerRequest
	(*Location)(nil),                    // 4: sro.Location
	(*CharacterTarget)(nil),             // 5: sro.CharacterTarget
}
var file_sro_gameserver_connection_proto_depIdxs = []int32{
	4, // 0: sro.gameserver.TransferPlayerRequest.location:type_name -> sro.Location
	5, // 1: sro.gameserver.ConnectionService.ConnectGameServer:input_type -> sro.CharacterTarget
	1, // 2: sro.gameserver.ConnectionService.VerifyConnectRequest:input_type -> sro.gameserver.VerifyConnectRequestMessage
	5, // 3: sro.gameserver.ConnectionService.IsPlaying:input_type -> sro.CharacterTarget
	0, // 4: sro.gameserver.ConnectionService.ConnectGameServer:output_type -> sro.gameserver.ConnectGameServerResponse
	5, // 5: sro.gameserver.ConnectionService.VerifyConnectRequest:output_type -> sro.CharacterTarget
	2, // 6: sro.gameserver.ConnectionService.IsPlaying:output_type -> sro.gameserver.ConnectionStatus
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sro_gameserver_connection_proto_init() }
func file_sro_gameserver_connection_proto_init() {
	if File_sro_gameserver_connection_proto != nil {
		return
	}
	file_sro_globals_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sro_gameserver_connection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sro_gameserver_connection_proto_goTypes,
		DependencyIndexes: file_sro_gameserver_connection_proto_depIdxs,
		MessageInfos:      file_sro_gameserver_connection_proto_msgTypes,
	}.Build()
	File_sro_gameserver_connection_proto = out.File
	file_sro_gameserver_connection_proto_rawDesc = nil
	file_sro_gameserver_connection_proto_goTypes = nil
	file_sro_gameserver_connection_proto_depIdxs = nil
}