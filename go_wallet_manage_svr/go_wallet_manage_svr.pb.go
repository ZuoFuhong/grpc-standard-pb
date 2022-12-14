// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: go_wallet_manage_svr.proto

// package 格式为服务进程名

package go_wallet_manage_svr

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

type CreateWalletReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateWalletReq) Reset() {
	*x = CreateWalletReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_wallet_manage_svr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletReq) ProtoMessage() {}

func (x *CreateWalletReq) ProtoReflect() protoreflect.Message {
	mi := &file_go_wallet_manage_svr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletReq.ProtoReflect.Descriptor instead.
func (*CreateWalletReq) Descriptor() ([]byte, []int) {
	return file_go_wallet_manage_svr_proto_rawDescGZIP(), []int{0}
}

type CreateWalletRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"` // 钱包地址
}

func (x *CreateWalletRsp) Reset() {
	*x = CreateWalletRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_wallet_manage_svr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRsp) ProtoMessage() {}

func (x *CreateWalletRsp) ProtoReflect() protoreflect.Message {
	mi := &file_go_wallet_manage_svr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRsp.ProtoReflect.Descriptor instead.
func (*CreateWalletRsp) Descriptor() ([]byte, []int) {
	return file_go_wallet_manage_svr_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWalletRsp) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ImportWalletReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKey string `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"` // 私钥
}

func (x *ImportWalletReq) Reset() {
	*x = ImportWalletReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_wallet_manage_svr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportWalletReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportWalletReq) ProtoMessage() {}

func (x *ImportWalletReq) ProtoReflect() protoreflect.Message {
	mi := &file_go_wallet_manage_svr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportWalletReq.ProtoReflect.Descriptor instead.
func (*ImportWalletReq) Descriptor() ([]byte, []int) {
	return file_go_wallet_manage_svr_proto_rawDescGZIP(), []int{2}
}

func (x *ImportWalletReq) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

type ImportWalletRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"` // 钱包地址
}

func (x *ImportWalletRsp) Reset() {
	*x = ImportWalletRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_wallet_manage_svr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportWalletRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportWalletRsp) ProtoMessage() {}

func (x *ImportWalletRsp) ProtoReflect() protoreflect.Message {
	mi := &file_go_wallet_manage_svr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportWalletRsp.ProtoReflect.Descriptor instead.
func (*ImportWalletRsp) Descriptor() ([]byte, []int) {
	return file_go_wallet_manage_svr_proto_rawDescGZIP(), []int{3}
}

func (x *ImportWalletRsp) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_go_wallet_manage_svr_proto protoreflect.FileDescriptor

var file_go_wallet_manage_svr_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x6f, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x76, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x6f,
	0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x76, 0x72, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x22, 0x2b, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x32, 0x0a, 0x0f, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x2b, 0x0a, 0x0f, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x32, 0xd2, 0x01, 0x0a, 0x14, 0x67, 0x6f, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x76, 0x72, 0x12, 0x5c, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x67,
	0x6f, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x76, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x76, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x73, 0x70, 0x12, 0x5c, 0x0a, 0x0c, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x5f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x76,
	0x72, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x25, 0x2e, 0x67, 0x6f, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x76, 0x72, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x73, 0x70, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x75, 0x6f, 0x46, 0x75, 0x68, 0x6f, 0x6e, 0x67,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x2d, 0x70,
	0x62, 0x2f, 0x67, 0x6f, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x76, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_wallet_manage_svr_proto_rawDescOnce sync.Once
	file_go_wallet_manage_svr_proto_rawDescData = file_go_wallet_manage_svr_proto_rawDesc
)

func file_go_wallet_manage_svr_proto_rawDescGZIP() []byte {
	file_go_wallet_manage_svr_proto_rawDescOnce.Do(func() {
		file_go_wallet_manage_svr_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_wallet_manage_svr_proto_rawDescData)
	})
	return file_go_wallet_manage_svr_proto_rawDescData
}

var file_go_wallet_manage_svr_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_go_wallet_manage_svr_proto_goTypes = []interface{}{
	(*CreateWalletReq)(nil), // 0: go_wallet_manage_svr.CreateWalletReq
	(*CreateWalletRsp)(nil), // 1: go_wallet_manage_svr.CreateWalletRsp
	(*ImportWalletReq)(nil), // 2: go_wallet_manage_svr.ImportWalletReq
	(*ImportWalletRsp)(nil), // 3: go_wallet_manage_svr.ImportWalletRsp
}
var file_go_wallet_manage_svr_proto_depIdxs = []int32{
	0, // 0: go_wallet_manage_svr.go_wallet_manage_svr.CreateWallet:input_type -> go_wallet_manage_svr.CreateWalletReq
	2, // 1: go_wallet_manage_svr.go_wallet_manage_svr.ImportWallet:input_type -> go_wallet_manage_svr.ImportWalletReq
	1, // 2: go_wallet_manage_svr.go_wallet_manage_svr.CreateWallet:output_type -> go_wallet_manage_svr.CreateWalletRsp
	3, // 3: go_wallet_manage_svr.go_wallet_manage_svr.ImportWallet:output_type -> go_wallet_manage_svr.ImportWalletRsp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_go_wallet_manage_svr_proto_init() }
func file_go_wallet_manage_svr_proto_init() {
	if File_go_wallet_manage_svr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_wallet_manage_svr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletReq); i {
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
		file_go_wallet_manage_svr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletRsp); i {
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
		file_go_wallet_manage_svr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportWalletReq); i {
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
		file_go_wallet_manage_svr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportWalletRsp); i {
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
			RawDescriptor: file_go_wallet_manage_svr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_go_wallet_manage_svr_proto_goTypes,
		DependencyIndexes: file_go_wallet_manage_svr_proto_depIdxs,
		MessageInfos:      file_go_wallet_manage_svr_proto_msgTypes,
	}.Build()
	File_go_wallet_manage_svr_proto = out.File
	file_go_wallet_manage_svr_proto_rawDesc = nil
	file_go_wallet_manage_svr_proto_goTypes = nil
	file_go_wallet_manage_svr_proto_depIdxs = nil
}
