// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: data-provider/protoadapter/adapter.proto

package protoadapter

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

type JobTypes int32

const (
	JobTypes_DATA_FEED     JobTypes = 0
	JobTypes_RANDOM_NUMBER JobTypes = 1
)

// Enum value maps for JobTypes.
var (
	JobTypes_name = map[int32]string{
		0: "DATA_FEED",
		1: "RANDOM_NUMBER",
	}
	JobTypes_value = map[string]int32{
		"DATA_FEED":     0,
		"RANDOM_NUMBER": 1,
	}
)

func (x JobTypes) Enum() *JobTypes {
	p := new(JobTypes)
	*p = x
	return p
}

func (x JobTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_data_provider_protoadapter_adapter_proto_enumTypes[0].Descriptor()
}

func (JobTypes) Type() protoreflect.EnumType {
	return &file_data_provider_protoadapter_adapter_proto_enumTypes[0]
}

func (x JobTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobTypes.Descriptor instead.
func (JobTypes) EnumDescriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{0}
}

type RequestType int32

const (
	RequestType_GET  RequestType = 0
	RequestType_POST RequestType = 1
)

// Enum value maps for RequestType.
var (
	RequestType_name = map[int32]string{
		0: "GET",
		1: "POST",
	}
	RequestType_value = map[string]int32{
		"GET":  0,
		"POST": 1,
	}
)

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}

func (x RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_data_provider_protoadapter_adapter_proto_enumTypes[1].Descriptor()
}

func (RequestType) Type() protoreflect.EnumType {
	return &file_data_provider_protoadapter_adapter_proto_enumTypes[1]
}

func (x RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestType.Descriptor instead.
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{1}
}

type FeedCategory int32

const (
	FeedCategory_RNG        FeedCategory = 0 // random number generator
	FeedCategory_SPORT      FeedCategory = 1 // live scores, past data e.t.c
	FeedCategory_PRICE_FEED FeedCategory = 2 // data such as spot prices
	FeedCategory_WEATHER    FeedCategory = 3 // transactions or balances happening on other blockchains
	FeedCategory_FLIGHT     FeedCategory = 4 // departure and arrival statuses on-chain
	FeedCategory_BLOCKCHAIN FeedCategory = 5 // transactions or balances happening on other blockchains
	FeedCategory_NFT_DATA   FeedCategory = 6 // floor Prices, Asset Listings, Collection Details
)

// Enum value maps for FeedCategory.
var (
	FeedCategory_name = map[int32]string{
		0: "RNG",
		1: "SPORT",
		2: "PRICE_FEED",
		3: "WEATHER",
		4: "FLIGHT",
		5: "BLOCKCHAIN",
		6: "NFT_DATA",
	}
	FeedCategory_value = map[string]int32{
		"RNG":        0,
		"SPORT":      1,
		"PRICE_FEED": 2,
		"WEATHER":    3,
		"FLIGHT":     4,
		"BLOCKCHAIN": 5,
		"NFT_DATA":   6,
	}
)

func (x FeedCategory) Enum() *FeedCategory {
	p := new(FeedCategory)
	*p = x
	return p
}

func (x FeedCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_data_provider_protoadapter_adapter_proto_enumTypes[2].Descriptor()
}

func (FeedCategory) Type() protoreflect.EnumType {
	return &file_data_provider_protoadapter_adapter_proto_enumTypes[2]
}

func (x FeedCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedCategory.Descriptor instead.
func (FeedCategory) EnumDescriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{2}
}

type DPInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenAddress string `protobuf:"bytes,1,opt,name=listen_address,json=listenAddress,proto3" json:"listen_address,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	KOrgId        string `protobuf:"bytes,3,opt,name=k_org_id,json=kOrgId,proto3" json:"k_org_id,omitempty"`
}

func (x *DPInfo) Reset() {
	*x = DPInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DPInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DPInfo) ProtoMessage() {}

func (x *DPInfo) ProtoReflect() protoreflect.Message {
	mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[0]
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
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{0}
}

func (x *DPInfo) GetListenAddress() string {
	if x != nil {
		return x.ListenAddress
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

// Node sends its IP details, sign `an arbitrary message
// and provides hash of message and signature to proof
// that it can request information from a Data Provider
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP string `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

type ServerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alive   bool   `protobuf:"varint,1,opt,name=alive,proto3" json:"alive,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServerStatus) Reset() {
	*x = ServerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatus) ProtoMessage() {}

func (x *ServerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatus.ProtoReflect.Descriptor instead.
func (*ServerStatus) Descriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{2}
}

func (x *ServerStatus) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

func (x *ServerStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Null struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Null) Reset() {
	*x = Null{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Null) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Null) ProtoMessage() {}

func (x *Null) ProtoReflect() protoreflect.Message {
	mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Null.ProtoReflect.Descriptor instead.
func (*Null) Descriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{3}
}

// Contains field that doesn't change throughout
// the lifetime of an adapter.
type AdapterHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdapterId     string `protobuf:"bytes,1,opt,name=adapter_id,json=adapterId,proto3" json:"adapter_id,omitempty"`
	OracleAddress string `protobuf:"bytes,2,opt,name=oracle_address,json=oracleAddress,proto3" json:"oracle_address,omitempty"`
}

func (x *AdapterHead) Reset() {
	*x = AdapterHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdapterHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdapterHead) ProtoMessage() {}

func (x *AdapterHead) ProtoReflect() protoreflect.Message {
	mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdapterHead.ProtoReflect.Descriptor instead.
func (*AdapterHead) Descriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{4}
}

func (x *AdapterHead) GetAdapterId() string {
	if x != nil {
		return x.AdapterId
	}
	return ""
}

func (x *AdapterHead) GetOracleAddress() string {
	if x != nil {
		return x.OracleAddress
	}
	return ""
}

type Adapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active        bool         `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	Name          string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	JobType       JobTypes     `protobuf:"varint,3,opt,name=job_type,json=jobType,proto3,enum=protoadapter.JobTypes" json:"job_type,omitempty"`
	AdapterId     string       `protobuf:"bytes,4,opt,name=adapter_id,json=adapterId,proto3" json:"adapter_id,omitempty"`
	OracleAddress string       `protobuf:"bytes,5,opt,name=oracle_address,json=oracleAddress,proto3" json:"oracle_address,omitempty"`
	Feeds         []*Feed      `protobuf:"bytes,6,rep,name=feeds,proto3" json:"feeds,omitempty"`
	Category      FeedCategory `protobuf:"varint,7,opt,name=category,proto3,enum=protoadapter.FeedCategory" json:"category,omitempty"`
}

func (x *Adapter) Reset() {
	*x = Adapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adapter) ProtoMessage() {}

func (x *Adapter) ProtoReflect() protoreflect.Message {
	mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adapter.ProtoReflect.Descriptor instead.
func (*Adapter) Descriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{5}
}

func (x *Adapter) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Adapter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Adapter) GetJobType() JobTypes {
	if x != nil {
		return x.JobType
	}
	return JobTypes_DATA_FEED
}

func (x *Adapter) GetAdapterId() string {
	if x != nil {
		return x.AdapterId
	}
	return ""
}

func (x *Adapter) GetOracleAddress() string {
	if x != nil {
		return x.OracleAddress
	}
	return ""
}

func (x *Adapter) GetFeeds() []*Feed {
	if x != nil {
		return x.Feeds
	}
	return nil
}

func (x *Adapter) GetCategory() FeedCategory {
	if x != nil {
		return x.Category
	}
	return FeedCategory_RNG
}

type Adapters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adapters []*Adapter `protobuf:"bytes,1,rep,name=adapters,proto3" json:"adapters,omitempty"`
}

func (x *Adapters) Reset() {
	*x = Adapters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Adapters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Adapters) ProtoMessage() {}

func (x *Adapters) ProtoReflect() protoreflect.Message {
	mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Adapters.ProtoReflect.Descriptor instead.
func (*Adapters) Descriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{6}
}

func (x *Adapters) GetAdapters() []*Adapter {
	if x != nil {
		return x.Adapters
	}
	return nil
}

type Reducer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Function string   `protobuf:"bytes,1,opt,name=function,proto3" json:"function,omitempty"`
	Args     []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *Reducer) Reset() {
	*x = Reducer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reducer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reducer) ProtoMessage() {}

func (x *Reducer) ProtoReflect() protoreflect.Message {
	mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reducer.ProtoReflect.Descriptor instead.
func (*Reducer) Descriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{7}
}

func (x *Reducer) GetFunction() string {
	if x != nil {
		return x.Function
	}
	return ""
}

func (x *Reducer) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type Feed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string         `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Headers     []*Feed_Header `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	RequestType RequestType    `protobuf:"varint,3,opt,name=request_type,json=requestType,proto3,enum=protoadapter.RequestType" json:"request_type,omitempty"`
	Reducers    []*Reducer     `protobuf:"bytes,4,rep,name=reducers,proto3" json:"reducers,omitempty"`
}

func (x *Feed) Reset() {
	*x = Feed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feed) ProtoMessage() {}

func (x *Feed) ProtoReflect() protoreflect.Message {
	mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feed.ProtoReflect.Descriptor instead.
func (*Feed) Descriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{8}
}

func (x *Feed) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Feed) GetHeaders() []*Feed_Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Feed) GetRequestType() RequestType {
	if x != nil {
		return x.RequestType
	}
	return RequestType_GET
}

func (x *Feed) GetReducers() []*Reducer {
	if x != nil {
		return x.Reducers
	}
	return nil
}

type Feed_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field map[string]string `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Feed_Header) Reset() {
	*x = Feed_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feed_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feed_Header) ProtoMessage() {}

func (x *Feed_Header) ProtoReflect() protoreflect.Message {
	mi := &file_data_provider_protoadapter_adapter_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feed_Header.ProtoReflect.Descriptor instead.
func (*Feed_Header) Descriptor() ([]byte, []int) {
	return file_data_provider_protoadapter_adapter_proto_rawDescGZIP(), []int{8, 0}
}

func (x *Feed_Header) GetField() map[string]string {
	if x != nil {
		return x.Field
	}
	return nil
}

var File_data_provider_protoadapter_adapter_proto protoreflect.FileDescriptor

var file_data_provider_protoadapter_adapter_proto_rawDesc = []byte{
	0x0a, 0x28, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x22, 0x5d, 0x0a, 0x06, 0x44, 0x50, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x08, 0x6b, 0x5f, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6b, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x22,
	0x3e, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x61, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x06, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x22, 0x53, 0x0a, 0x0b, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x48, 0x65, 0x61, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x90, 0x02, 0x0a,
	0x07, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x07,
	0x6a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a,
	0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x65, 0x64,
	0x52, 0x05, 0x66, 0x65, 0x65, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22,
	0x3d, 0x0a, 0x08, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x52, 0x08, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x22, 0x39,
	0x0a, 0x07, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0xbe, 0x02, 0x0a, 0x04, 0x46, 0x65,
	0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x33, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x72,
	0x52, 0x08, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x72, 0x73, 0x1a, 0x7e, 0x0a, 0x06, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x1a, 0x38, 0x0a, 0x0a, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x2c, 0x0a, 0x08, 0x4a, 0x6f,
	0x62, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x46,
	0x45, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x5f,
	0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x01, 0x2a, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x2a, 0x69, 0x0a, 0x0c, 0x46, 0x65,
	0x65, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x4e,
	0x47, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x57, 0x45, 0x41, 0x54, 0x48, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x4c, 0x49, 0x47, 0x48, 0x54, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x4c, 0x4f, 0x43, 0x4b,
	0x43, 0x48, 0x41, 0x49, 0x4e, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x46, 0x54, 0x5f, 0x44,
	0x41, 0x54, 0x41, 0x10, 0x06, 0x32, 0x88, 0x02, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x6f,
	0x4b, 0x6e, 0x6f, 0x77, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x50, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x4e,
	0x75, 0x6c, 0x6c, 0x12, 0x3e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x2e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x48, 0x65, 0x61, 0x64, 0x1a, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x30, 0x01,
	0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x6c, 0x61, 0x79, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2f, 0x6b, 0x6c, 0x61, 0x79, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_provider_protoadapter_adapter_proto_rawDescOnce sync.Once
	file_data_provider_protoadapter_adapter_proto_rawDescData = file_data_provider_protoadapter_adapter_proto_rawDesc
)

func file_data_provider_protoadapter_adapter_proto_rawDescGZIP() []byte {
	file_data_provider_protoadapter_adapter_proto_rawDescOnce.Do(func() {
		file_data_provider_protoadapter_adapter_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_provider_protoadapter_adapter_proto_rawDescData)
	})
	return file_data_provider_protoadapter_adapter_proto_rawDescData
}

var file_data_provider_protoadapter_adapter_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_data_provider_protoadapter_adapter_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_data_provider_protoadapter_adapter_proto_goTypes = []interface{}{
	(JobTypes)(0),        // 0: protoadapter.JobTypes
	(RequestType)(0),     // 1: protoadapter.RequestType
	(FeedCategory)(0),    // 2: protoadapter.FeedCategory
	(*DPInfo)(nil),       // 3: protoadapter.DPInfo
	(*Node)(nil),         // 4: protoadapter.Node
	(*ServerStatus)(nil), // 5: protoadapter.ServerStatus
	(*Null)(nil),         // 6: protoadapter.Null
	(*AdapterHead)(nil),  // 7: protoadapter.AdapterHead
	(*Adapter)(nil),      // 8: protoadapter.Adapter
	(*Adapters)(nil),     // 9: protoadapter.Adapters
	(*Reducer)(nil),      // 10: protoadapter.Reducer
	(*Feed)(nil),         // 11: protoadapter.Feed
	(*Feed_Header)(nil),  // 12: protoadapter.Feed.Header
	nil,                  // 13: protoadapter.Feed.Header.FieldEntry
}
var file_data_provider_protoadapter_adapter_proto_depIdxs = []int32{
	0,  // 0: protoadapter.Adapter.job_type:type_name -> protoadapter.JobTypes
	11, // 1: protoadapter.Adapter.feeds:type_name -> protoadapter.Feed
	2,  // 2: protoadapter.Adapter.category:type_name -> protoadapter.FeedCategory
	8,  // 3: protoadapter.Adapters.adapters:type_name -> protoadapter.Adapter
	12, // 4: protoadapter.Feed.headers:type_name -> protoadapter.Feed.Header
	1,  // 5: protoadapter.Feed.request_type:type_name -> protoadapter.RequestType
	10, // 6: protoadapter.Feed.reducers:type_name -> protoadapter.Reducer
	13, // 7: protoadapter.Feed.Header.field:type_name -> protoadapter.Feed.Header.FieldEntry
	4,  // 8: protoadapter.DataProviderService.Ping:input_type -> protoadapter.Node
	3,  // 9: protoadapter.DataProviderService.AddToKnowPeers:input_type -> protoadapter.DPInfo
	7,  // 10: protoadapter.DataProviderService.GetAdapter:input_type -> protoadapter.AdapterHead
	4,  // 11: protoadapter.DataProviderService.ListAdapters:input_type -> protoadapter.Node
	5,  // 12: protoadapter.DataProviderService.Ping:output_type -> protoadapter.ServerStatus
	6,  // 13: protoadapter.DataProviderService.AddToKnowPeers:output_type -> protoadapter.Null
	8,  // 14: protoadapter.DataProviderService.GetAdapter:output_type -> protoadapter.Adapter
	8,  // 15: protoadapter.DataProviderService.ListAdapters:output_type -> protoadapter.Adapter
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_data_provider_protoadapter_adapter_proto_init() }
func file_data_provider_protoadapter_adapter_proto_init() {
	if File_data_provider_protoadapter_adapter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_provider_protoadapter_adapter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_data_provider_protoadapter_adapter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_data_provider_protoadapter_adapter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatus); i {
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
		file_data_provider_protoadapter_adapter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Null); i {
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
		file_data_provider_protoadapter_adapter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdapterHead); i {
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
		file_data_provider_protoadapter_adapter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adapter); i {
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
		file_data_provider_protoadapter_adapter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Adapters); i {
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
		file_data_provider_protoadapter_adapter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reducer); i {
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
		file_data_provider_protoadapter_adapter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feed); i {
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
		file_data_provider_protoadapter_adapter_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feed_Header); i {
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
			RawDescriptor: file_data_provider_protoadapter_adapter_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_provider_protoadapter_adapter_proto_goTypes,
		DependencyIndexes: file_data_provider_protoadapter_adapter_proto_depIdxs,
		EnumInfos:         file_data_provider_protoadapter_adapter_proto_enumTypes,
		MessageInfos:      file_data_provider_protoadapter_adapter_proto_msgTypes,
	}.Build()
	File_data_provider_protoadapter_adapter_proto = out.File
	file_data_provider_protoadapter_adapter_proto_rawDesc = nil
	file_data_provider_protoadapter_adapter_proto_goTypes = nil
	file_data_provider_protoadapter_adapter_proto_depIdxs = nil
}
