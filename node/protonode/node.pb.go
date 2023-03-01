// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: node/protonode/node.proto

package protonode

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DPInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenIP string `protobuf:"bytes,1,opt,name=ListenIP,proto3" json:"ListenIP,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	KOrgId   string `protobuf:"bytes,3,opt,name=k_org_id,json=kOrgId,proto3" json:"k_org_id,omitempty"`
}

func (x *DPInfo) Reset() {
	*x = DPInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_protonode_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DPInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DPInfo) ProtoMessage() {}

func (x *DPInfo) ProtoReflect() protoreflect.Message {
	mi := &file_node_protonode_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DPInfo.ProtoReflect.Descriptor instead.
func (*DPInfo) Descriptor() ([]byte, []int) {
	return file_node_protonode_node_proto_rawDescGZIP(), []int{0}
}

func (x *DPInfo) GetListenIP() string {
	if x != nil {
		return x.ListenIP
	}
	return ""
}

func (x *DPInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DPInfo) GetKOrgId() string {
	if x != nil {
		return x.KOrgId
	}
	return ""
}

type HandShakeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ErrorMsg string `protobuf:"bytes,2,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
}

func (x *HandShakeStatus) Reset() {
	*x = HandShakeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_protonode_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandShakeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandShakeStatus) ProtoMessage() {}

func (x *HandShakeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_node_protonode_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandShakeStatus.ProtoReflect.Descriptor instead.
func (*HandShakeStatus) Descriptor() ([]byte, []int) {
	return file_node_protonode_node_proto_rawDescGZIP(), []int{1}
}

func (x *HandShakeStatus) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *HandShakeStatus) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

var File_node_protonode_node_proto protoreflect.FileDescriptor

var file_node_protonode_node_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x6f, 0x64, 0x65,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x06, 0x44, 0x50, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x08, 0x6b, 0x5f, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6b, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x0f, 0x48, 0x61,
	0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x73, 0x67, 0x32, 0x49, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x50, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x48, 0x61,
	0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6c, 0x61, 0x79,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2f, 0x6b, 0x6c, 0x61, 0x79, 0x6f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_node_protonode_node_proto_rawDescOnce sync.Once
	file_node_protonode_node_proto_rawDescData = file_node_protonode_node_proto_rawDesc
)

func file_node_protonode_node_proto_rawDescGZIP() []byte {
	file_node_protonode_node_proto_rawDescOnce.Do(func() {
		file_node_protonode_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_protonode_node_proto_rawDescData)
	})
	return file_node_protonode_node_proto_rawDescData
}

var file_node_protonode_node_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_node_protonode_node_proto_goTypes = []interface{}{
	(*DPInfo)(nil),          // 0: protonode.DPInfo
	(*HandShakeStatus)(nil), // 1: protonode.HandShakeStatus
}
var file_node_protonode_node_proto_depIdxs = []int32{
	0, // 0: protonode.NodeService.HandShake:input_type -> protonode.DPInfo
	1, // 1: protonode.NodeService.HandShake:output_type -> protonode.HandShakeStatus
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_node_protonode_node_proto_init() }
func file_node_protonode_node_proto_init() {
	if File_node_protonode_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_protonode_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DPInfo); i {
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
		file_node_protonode_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandShakeStatus); i {
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
			RawDescriptor: file_node_protonode_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_protonode_node_proto_goTypes,
		DependencyIndexes: file_node_protonode_node_proto_depIdxs,
		MessageInfos:      file_node_protonode_node_proto_msgTypes,
	}.Build()
	File_node_protonode_node_proto = out.File
	file_node_protonode_node_proto_rawDesc = nil
	file_node_protonode_node_proto_goTypes = nil
	file_node_protonode_node_proto_depIdxs = nil
}
