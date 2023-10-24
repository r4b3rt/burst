// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/burst.proto

package api

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

type PortMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientPort int64  `protobuf:"varint,1,opt,name=clientPort,proto3" json:"clientPort,omitempty"`
	ServerPort int64  `protobuf:"varint,2,opt,name=serverPort,proto3" json:"serverPort,omitempty"`
	Protocol   string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (x *PortMapping) Reset() {
	*x = PortMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_burst_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortMapping) ProtoMessage() {}

func (x *PortMapping) ProtoReflect() protoreflect.Message {
	mi := &file_api_burst_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortMapping.ProtoReflect.Descriptor instead.
func (*PortMapping) Descriptor() ([]byte, []int) {
	return file_api_burst_proto_rawDescGZIP(), []int{0}
}

func (x *PortMapping) GetClientPort() int64 {
	if x != nil {
		return x.ClientPort
	}
	return 0
}

func (x *PortMapping) GetServerPort() int64 {
	if x != nil {
		return x.ServerPort
	}
	return 0
}

func (x *PortMapping) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

type ExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientWsPort int64          `protobuf:"varint,1,opt,name=clientWsPort,proto3" json:"clientWsPort,omitempty"`
	PortMapping  []*PortMapping `protobuf:"bytes,2,rep,name=portMapping,proto3" json:"portMapping,omitempty"`
}

func (x *ExportRequest) Reset() {
	*x = ExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_burst_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportRequest) ProtoMessage() {}

func (x *ExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_burst_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportRequest.ProtoReflect.Descriptor instead.
func (*ExportRequest) Descriptor() ([]byte, []int) {
	return file_api_burst_proto_rawDescGZIP(), []int{1}
}

func (x *ExportRequest) GetClientWsPort() int64 {
	if x != nil {
		return x.ClientWsPort
	}
	return 0
}

func (x *ExportRequest) GetPortMapping() []*PortMapping {
	if x != nil {
		return x.PortMapping
	}
	return nil
}

type PortMappingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mapping      *PortMapping `protobuf:"bytes,1,opt,name=mapping,proto3" json:"mapping,omitempty"`
	Error        string       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	ConnectionId string       `protobuf:"bytes,3,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
}

func (x *PortMappingResp) Reset() {
	*x = PortMappingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_burst_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortMappingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortMappingResp) ProtoMessage() {}

func (x *PortMappingResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_burst_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortMappingResp.ProtoReflect.Descriptor instead.
func (*PortMappingResp) Descriptor() ([]byte, []int) {
	return file_api_burst_proto_rawDescGZIP(), []int{2}
}

func (x *PortMappingResp) GetMapping() *PortMapping {
	if x != nil {
		return x.Mapping
	}
	return nil
}

func (x *PortMappingResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *PortMappingResp) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

type ExportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*PortMappingResp `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ExportResponse) Reset() {
	*x = ExportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_burst_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportResponse) ProtoMessage() {}

func (x *ExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_burst_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportResponse.ProtoReflect.Descriptor instead.
func (*ExportResponse) Descriptor() ([]byte, []int) {
	return file_api_burst_proto_rawDescGZIP(), []int{3}
}

func (x *ExportResponse) GetItems() []*PortMappingResp {
	if x != nil {
		return x.Items
	}
	return nil
}

type TransFromData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionId     string `protobuf:"bytes,1,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
	UserConnectionId string `protobuf:"bytes,2,opt,name=userConnectionId,proto3" json:"userConnectionId,omitempty"`
	Data             []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TransFromData) Reset() {
	*x = TransFromData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_burst_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransFromData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransFromData) ProtoMessage() {}

func (x *TransFromData) ProtoReflect() protoreflect.Message {
	mi := &file_api_burst_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransFromData.ProtoReflect.Descriptor instead.
func (*TransFromData) Descriptor() ([]byte, []int) {
	return file_api_burst_proto_rawDescGZIP(), []int{4}
}

func (x *TransFromData) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *TransFromData) GetUserConnectionId() string {
	if x != nil {
		return x.UserConnectionId
	}
	return ""
}

func (x *TransFromData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_burst_proto protoreflect.FileDescriptor

var file_api_burst_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x75, 0x72, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x69, 0x0a, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x63, 0x0a, 0x0d,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x57, 0x73, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x57, 0x73, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x22, 0x73, 0x0a, 0x0f, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x73, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x63, 0x0a, 0x05, 0x62, 0x75, 0x72, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x46, 0x72,
	0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x46, 0x72,
	0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x28, 0x01, 0x30, 0x01, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x7a, 0x64, 0x77, 0x78, 0x2f, 0x62,
	0x75, 0x72, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_burst_proto_rawDescOnce sync.Once
	file_api_burst_proto_rawDescData = file_api_burst_proto_rawDesc
)

func file_api_burst_proto_rawDescGZIP() []byte {
	file_api_burst_proto_rawDescOnce.Do(func() {
		file_api_burst_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_burst_proto_rawDescData)
	})
	return file_api_burst_proto_rawDescData
}

var file_api_burst_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_burst_proto_goTypes = []interface{}{
	(*PortMapping)(nil),     // 0: PortMapping
	(*ExportRequest)(nil),   // 1: ExportRequest
	(*PortMappingResp)(nil), // 2: PortMappingResp
	(*ExportResponse)(nil),  // 3: ExportResponse
	(*TransFromData)(nil),   // 4: TransFromData
}
var file_api_burst_proto_depIdxs = []int32{
	0, // 0: ExportRequest.portMapping:type_name -> PortMapping
	0, // 1: PortMappingResp.mapping:type_name -> PortMapping
	2, // 2: ExportResponse.items:type_name -> PortMappingResp
	1, // 3: burst.Export:input_type -> ExportRequest
	4, // 4: burst.Transform:input_type -> TransFromData
	3, // 5: burst.Export:output_type -> ExportResponse
	4, // 6: burst.Transform:output_type -> TransFromData
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_burst_proto_init() }
func file_api_burst_proto_init() {
	if File_api_burst_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_burst_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortMapping); i {
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
		file_api_burst_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportRequest); i {
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
		file_api_burst_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortMappingResp); i {
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
		file_api_burst_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportResponse); i {
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
		file_api_burst_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransFromData); i {
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
			RawDescriptor: file_api_burst_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_burst_proto_goTypes,
		DependencyIndexes: file_api_burst_proto_depIdxs,
		MessageInfos:      file_api_burst_proto_msgTypes,
	}.Build()
	File_api_burst_proto = out.File
	file_api_burst_proto_rawDesc = nil
	file_api_burst_proto_goTypes = nil
	file_api_burst_proto_depIdxs = nil
}