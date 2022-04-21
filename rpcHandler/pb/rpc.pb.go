// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: rpc.proto

package pb

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

type Tokenmsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Tokenmsg) Reset() {
	*x = Tokenmsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tokenmsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokenmsg) ProtoMessage() {}

func (x *Tokenmsg) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tokenmsg.ProtoReflect.Descriptor instead.
func (*Tokenmsg) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *Tokenmsg) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserIDmsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UserIDmsg) Reset() {
	*x = UserIDmsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIDmsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIDmsg) ProtoMessage() {}

func (x *UserIDmsg) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIDmsg.ProtoReflect.Descriptor instead.
func (*UserIDmsg) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *UserIDmsg) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type DeviceIDsmsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceIDs []int64 `protobuf:"varint,1,rep,packed,name=deviceIDs,proto3" json:"deviceIDs,omitempty"`
}

func (x *DeviceIDsmsg) Reset() {
	*x = DeviceIDsmsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceIDsmsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceIDsmsg) ProtoMessage() {}

func (x *DeviceIDsmsg) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceIDsmsg.ProtoReflect.Descriptor instead.
func (*DeviceIDsmsg) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceIDsmsg) GetDeviceIDs() []int64 {
	if x != nil {
		return x.DeviceIDs
	}
	return nil
}

type Didmsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Didmsg) Reset() {
	*x = Didmsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Didmsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Didmsg) ProtoMessage() {}

func (x *Didmsg) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Didmsg.ProtoReflect.Descriptor instead.
func (*Didmsg) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *Didmsg) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DataTypemsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataType int32 `protobuf:"varint,1,opt,name=dataType,proto3" json:"dataType,omitempty"`
}

func (x *DataTypemsg) Reset() {
	*x = DataTypemsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataTypemsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataTypemsg) ProtoMessage() {}

func (x *DataTypemsg) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataTypemsg.ProtoReflect.Descriptor instead.
func (*DataTypemsg) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *DataTypemsg) GetDataType() int32 {
	if x != nil {
		return x.DataType
	}
	return 0
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0x20, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x23, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x6d, 0x73, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2c, 0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x44, 0x73, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x44, 0x73, 0x22, 0x18, 0x0a, 0x06, 0x44, 0x69, 0x64, 0x6d, 0x73, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29,
	0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x6d, 0x73, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x32, 0xa4, 0x01, 0x0a, 0x0c, 0x61, 0x6c,
	0x6c, 0x52, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x55, 0x69, 0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x73, 0x67, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x6d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x64, 0x49, 0x44, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x6d, 0x73, 0x67, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x73, 0x6d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x42, 0x79, 0x64, 0x49,
	0x44, 0x12, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x64, 0x6d, 0x73, 0x67, 0x1a, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x6d, 0x73, 0x67, 0x22, 0x00,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x69, 0x72, 0x6c, 0x61, 0x6e, 0x72, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x74, 0x2d,
	0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_proto_goTypes = []interface{}{
	(*Tokenmsg)(nil),     // 0: pb.Tokenmsg
	(*UserIDmsg)(nil),    // 1: pb.UserIDmsg
	(*DeviceIDsmsg)(nil), // 2: pb.DeviceIDsmsg
	(*Didmsg)(nil),       // 3: pb.Didmsg
	(*DataTypemsg)(nil),  // 4: pb.DataTypemsg
}
var file_rpc_proto_depIdxs = []int32{
	0, // 0: pb.allRpcServer.GetUidByToken:input_type -> pb.Tokenmsg
	0, // 1: pb.allRpcServer.GetdIDByToken:input_type -> pb.Tokenmsg
	3, // 2: pb.allRpcServer.GetDataTypeBydID:input_type -> pb.Didmsg
	1, // 3: pb.allRpcServer.GetUidByToken:output_type -> pb.UserIDmsg
	2, // 4: pb.allRpcServer.GetdIDByToken:output_type -> pb.DeviceIDsmsg
	4, // 5: pb.allRpcServer.GetDataTypeBydID:output_type -> pb.DataTypemsg
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tokenmsg); i {
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
		file_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIDmsg); i {
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
		file_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceIDsmsg); i {
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
		file_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Didmsg); i {
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
		file_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataTypemsg); i {
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
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}