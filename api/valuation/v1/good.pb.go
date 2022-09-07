// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: valuation/v1/good.proto

package valuation

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type GoodInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type   string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Unit   string  `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
	Price  float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Tariff float32 `protobuf:"fixed32,6,opt,name=tariff,proto3" json:"tariff,omitempty"`
	Alias  string  `protobuf:"bytes,7,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *GoodInfo) Reset() {
	*x = GoodInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodInfo) ProtoMessage() {}

func (x *GoodInfo) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodInfo.ProtoReflect.Descriptor instead.
func (*GoodInfo) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{0}
}

func (x *GoodInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GoodInfo) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *GoodInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GoodInfo) GetTariff() float32 {
	if x != nil {
		return x.Tariff
	}
	return 0
}

func (x *GoodInfo) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

type CreateGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Good *GoodInfo `protobuf:"bytes,1,opt,name=good,proto3" json:"good,omitempty"`
}

func (x *CreateGoodsRequest) Reset() {
	*x = CreateGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodsRequest) ProtoMessage() {}

func (x *CreateGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodsRequest.ProtoReflect.Descriptor instead.
func (*CreateGoodsRequest) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGoodsRequest) GetGood() *GoodInfo {
	if x != nil {
		return x.Good
	}
	return nil
}

type CreateGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateGoodsReply) Reset() {
	*x = CreateGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodsReply) ProtoMessage() {}

func (x *CreateGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodsReply.ProtoReflect.Descriptor instead.
func (*CreateGoodsReply) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGoodsReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Good *GoodInfo `protobuf:"bytes,1,opt,name=good,proto3" json:"good,omitempty"`
}

func (x *UpdateGoodsRequest) Reset() {
	*x = UpdateGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodsRequest) ProtoMessage() {}

func (x *UpdateGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodsRequest.ProtoReflect.Descriptor instead.
func (*UpdateGoodsRequest) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGoodsRequest) GetGood() *GoodInfo {
	if x != nil {
		return x.Good
	}
	return nil
}

type UpdateGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateGoodsReply) Reset() {
	*x = UpdateGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodsReply) ProtoMessage() {}

func (x *UpdateGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodsReply.ProtoReflect.Descriptor instead.
func (*UpdateGoodsReply) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{4}
}

type DeleteGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGoodsRequest) Reset() {
	*x = DeleteGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodsRequest) ProtoMessage() {}

func (x *DeleteGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodsRequest.ProtoReflect.Descriptor instead.
func (*DeleteGoodsRequest) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteGoodsRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGoodsReply) Reset() {
	*x = DeleteGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodsReply) ProtoMessage() {}

func (x *DeleteGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodsReply.ProtoReflect.Descriptor instead.
func (*DeleteGoodsReply) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{6}
}

type GetGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGoodsRequest) Reset() {
	*x = GetGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsRequest) ProtoMessage() {}

func (x *GetGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsRequest.ProtoReflect.Descriptor instead.
func (*GetGoodsRequest) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{7}
}

func (x *GetGoodsRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetGoodsReply) Reset() {
	*x = GetGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsReply) ProtoMessage() {}

func (x *GetGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsReply.ProtoReflect.Descriptor instead.
func (*GetGoodsReply) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{8}
}

type ListGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  uint64 `protobuf:"varint,1,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize uint64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListGoodsRequest) Reset() {
	*x = ListGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoodsRequest) ProtoMessage() {}

func (x *ListGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoodsRequest.ProtoReflect.Descriptor instead.
func (*ListGoodsRequest) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{9}
}

func (x *ListGoodsRequest) GetPageNum() uint64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListGoodsRequest) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goods    []*GoodInfo `protobuf:"bytes,1,rep,name=goods,proto3" json:"goods,omitempty"`
	PageNum  uint64      `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize uint64      `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Total    uint64      `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListGoodsReply) Reset() {
	*x = ListGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoodsReply) ProtoMessage() {}

func (x *ListGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoodsReply.ProtoReflect.Descriptor instead.
func (*ListGoodsReply) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{10}
}

func (x *ListGoodsReply) GetGoods() []*GoodInfo {
	if x != nil {
		return x.Goods
	}
	return nil
}

func (x *ListGoodsReply) GetPageNum() uint64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListGoodsReply) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListGoodsReply) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListGoodsByWordsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goods []*GoodInfo `protobuf:"bytes,1,rep,name=goods,proto3" json:"goods,omitempty"`
}

func (x *ListGoodsByWordsReply) Reset() {
	*x = ListGoodsByWordsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoodsByWordsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoodsByWordsReply) ProtoMessage() {}

func (x *ListGoodsByWordsReply) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoodsByWordsReply.ProtoReflect.Descriptor instead.
func (*ListGoodsByWordsReply) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{11}
}

func (x *ListGoodsByWordsReply) GetGoods() []*GoodInfo {
	if x != nil {
		return x.Goods
	}
	return nil
}

type ListGoodsByWordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Words string `protobuf:"bytes,1,opt,name=words,proto3" json:"words,omitempty"`
}

func (x *ListGoodsByWordsRequest) Reset() {
	*x = ListGoodsByWordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_valuation_v1_good_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoodsByWordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoodsByWordsRequest) ProtoMessage() {}

func (x *ListGoodsByWordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_valuation_v1_good_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoodsByWordsRequest.ProtoReflect.Descriptor instead.
func (*ListGoodsByWordsRequest) Descriptor() ([]byte, []int) {
	return file_valuation_v1_good_proto_rawDescGZIP(), []int{12}
}

func (x *ListGoodsByWordsRequest) GetWords() string {
	if x != nil {
		return x.Words
	}
	return ""
}

var File_valuation_v1_good_proto protoreflect.FileDescriptor

var file_valuation_v1_good_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xea, 0x01, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x32, 0x04, 0x20,
	0x00, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x14,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x0a, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x0a, 0x52, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x0a, 0x05, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x0a, 0x05, 0x2d, 0x00, 0x00,
	0x00, 0x00, 0x52, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x12, 0x20, 0x0a, 0x05, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05,
	0x10, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x41, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x22,
	0x22, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x67, 0x6f, 0x6f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x5a, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12,
	0x23, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x46, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42,
	0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x2f, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x32, 0xe6, 0x04, 0x0a, 0x04,
	0x47, 0x6f, 0x6f, 0x64, 0x12, 0x62, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x05,
	0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x62, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0a, 0x1a, 0x05, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x51, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x08, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x02, 0x12, 0x00, 0x12, 0x5f, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x67, 0x6f, 0x6f,
	0x64, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x7c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x57, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x12, 0x12, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x7d, 0x42, 0x23, 0x5a, 0x21, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_valuation_v1_good_proto_rawDescOnce sync.Once
	file_valuation_v1_good_proto_rawDescData = file_valuation_v1_good_proto_rawDesc
)

func file_valuation_v1_good_proto_rawDescGZIP() []byte {
	file_valuation_v1_good_proto_rawDescOnce.Do(func() {
		file_valuation_v1_good_proto_rawDescData = protoimpl.X.CompressGZIP(file_valuation_v1_good_proto_rawDescData)
	})
	return file_valuation_v1_good_proto_rawDescData
}

var file_valuation_v1_good_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_valuation_v1_good_proto_goTypes = []interface{}{
	(*GoodInfo)(nil),                // 0: api.valuation.GoodInfo
	(*CreateGoodsRequest)(nil),      // 1: api.valuation.CreateGoodsRequest
	(*CreateGoodsReply)(nil),        // 2: api.valuation.CreateGoodsReply
	(*UpdateGoodsRequest)(nil),      // 3: api.valuation.UpdateGoodsRequest
	(*UpdateGoodsReply)(nil),        // 4: api.valuation.UpdateGoodsReply
	(*DeleteGoodsRequest)(nil),      // 5: api.valuation.DeleteGoodsRequest
	(*DeleteGoodsReply)(nil),        // 6: api.valuation.DeleteGoodsReply
	(*GetGoodsRequest)(nil),         // 7: api.valuation.GetGoodsRequest
	(*GetGoodsReply)(nil),           // 8: api.valuation.GetGoodsReply
	(*ListGoodsRequest)(nil),        // 9: api.valuation.ListGoodsRequest
	(*ListGoodsReply)(nil),          // 10: api.valuation.ListGoodsReply
	(*ListGoodsByWordsReply)(nil),   // 11: api.valuation.ListGoodsByWordsReply
	(*ListGoodsByWordsRequest)(nil), // 12: api.valuation.ListGoodsByWordsRequest
}
var file_valuation_v1_good_proto_depIdxs = []int32{
	0,  // 0: api.valuation.CreateGoodsRequest.good:type_name -> api.valuation.GoodInfo
	0,  // 1: api.valuation.UpdateGoodsRequest.good:type_name -> api.valuation.GoodInfo
	0,  // 2: api.valuation.ListGoodsReply.goods:type_name -> api.valuation.GoodInfo
	0,  // 3: api.valuation.ListGoodsByWordsReply.goods:type_name -> api.valuation.GoodInfo
	1,  // 4: api.valuation.Good.CreateGood:input_type -> api.valuation.CreateGoodsRequest
	3,  // 5: api.valuation.Good.UpdateGood:input_type -> api.valuation.UpdateGoodsRequest
	5,  // 6: api.valuation.Good.DeleteGood:input_type -> api.valuation.DeleteGoodsRequest
	7,  // 7: api.valuation.Good.GetGood:input_type -> api.valuation.GetGoodsRequest
	9,  // 8: api.valuation.Good.ListGoods:input_type -> api.valuation.ListGoodsRequest
	12, // 9: api.valuation.Good.ListGoodsByWords:input_type -> api.valuation.ListGoodsByWordsRequest
	2,  // 10: api.valuation.Good.CreateGood:output_type -> api.valuation.CreateGoodsReply
	4,  // 11: api.valuation.Good.UpdateGood:output_type -> api.valuation.UpdateGoodsReply
	6,  // 12: api.valuation.Good.DeleteGood:output_type -> api.valuation.DeleteGoodsReply
	8,  // 13: api.valuation.Good.GetGood:output_type -> api.valuation.GetGoodsReply
	10, // 14: api.valuation.Good.ListGoods:output_type -> api.valuation.ListGoodsReply
	11, // 15: api.valuation.Good.ListGoodsByWords:output_type -> api.valuation.ListGoodsByWordsReply
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_valuation_v1_good_proto_init() }
func file_valuation_v1_good_proto_init() {
	if File_valuation_v1_good_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_valuation_v1_good_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodInfo); i {
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
		file_valuation_v1_good_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodsRequest); i {
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
		file_valuation_v1_good_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodsReply); i {
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
		file_valuation_v1_good_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodsRequest); i {
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
		file_valuation_v1_good_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodsReply); i {
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
		file_valuation_v1_good_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoodsRequest); i {
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
		file_valuation_v1_good_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoodsReply); i {
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
		file_valuation_v1_good_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsRequest); i {
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
		file_valuation_v1_good_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsReply); i {
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
		file_valuation_v1_good_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoodsRequest); i {
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
		file_valuation_v1_good_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoodsReply); i {
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
		file_valuation_v1_good_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoodsByWordsReply); i {
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
		file_valuation_v1_good_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoodsByWordsRequest); i {
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
			RawDescriptor: file_valuation_v1_good_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_valuation_v1_good_proto_goTypes,
		DependencyIndexes: file_valuation_v1_good_proto_depIdxs,
		MessageInfos:      file_valuation_v1_good_proto_msgTypes,
	}.Build()
	File_valuation_v1_good_proto = out.File
	file_valuation_v1_good_proto_rawDesc = nil
	file_valuation_v1_good_proto_goTypes = nil
	file_valuation_v1_good_proto_depIdxs = nil
}
