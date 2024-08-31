// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.2
// source: programmableLayer/programmableLayer.proto

package programmableLayer

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_programmableLayer_programmableLayer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_programmableLayer_programmableLayer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_programmableLayer_programmableLayer_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_programmableLayer_programmableLayer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_programmableLayer_programmableLayer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_programmableLayer_programmableLayer_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash          string `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
	BtcAmount       string `protobuf:"bytes,2,opt,name=btcAmount,proto3" json:"btcAmount,omitempty"`
	RuneId          string `protobuf:"bytes,3,opt,name=runeId,proto3" json:"runeId,omitempty"`
	RuneAmount      uint64 `protobuf:"varint,4,opt,name=runeAmount,proto3" json:"runeAmount,omitempty"`
	UserAddress     string `protobuf:"bytes,5,opt,name=userAddress,proto3" json:"userAddress,omitempty"`
	Publickey       string `protobuf:"bytes,6,opt,name=publickey,proto3" json:"publickey,omitempty"`
	EthereumAddress string `protobuf:"bytes,7,opt,name=ethereumAddress,proto3" json:"ethereumAddress,omitempty"`
}

func (x *FundRequest) Reset() {
	*x = FundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_programmableLayer_programmableLayer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundRequest) ProtoMessage() {}

func (x *FundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_programmableLayer_programmableLayer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundRequest.ProtoReflect.Descriptor instead.
func (*FundRequest) Descriptor() ([]byte, []int) {
	return file_programmableLayer_programmableLayer_proto_rawDescGZIP(), []int{2}
}

func (x *FundRequest) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *FundRequest) GetBtcAmount() string {
	if x != nil {
		return x.BtcAmount
	}
	return ""
}

func (x *FundRequest) GetRuneId() string {
	if x != nil {
		return x.RuneId
	}
	return ""
}

func (x *FundRequest) GetRuneAmount() uint64 {
	if x != nil {
		return x.RuneAmount
	}
	return 0
}

func (x *FundRequest) GetUserAddress() string {
	if x != nil {
		return x.UserAddress
	}
	return ""
}

func (x *FundRequest) GetPublickey() string {
	if x != nil {
		return x.Publickey
	}
	return ""
}

func (x *FundRequest) GetEthereumAddress() string {
	if x != nil {
		return x.EthereumAddress
	}
	return ""
}

type FundRequestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	TxHash  string `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

func (x *FundRequestReply) Reset() {
	*x = FundRequestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_programmableLayer_programmableLayer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundRequestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundRequestReply) ProtoMessage() {}

func (x *FundRequestReply) ProtoReflect() protoreflect.Message {
	mi := &file_programmableLayer_programmableLayer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundRequestReply.ProtoReflect.Descriptor instead.
func (*FundRequestReply) Descriptor() ([]byte, []int) {
	return file_programmableLayer_programmableLayer_proto_rawDescGZIP(), []int{3}
}

func (x *FundRequestReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *FundRequestReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FundRequestReply) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type MintRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferSignature string `protobuf:"bytes,1,opt,name=transferSignature,proto3" json:"transferSignature,omitempty"`
	CalldataSignature string `protobuf:"bytes,2,opt,name=calldataSignature,proto3" json:"calldataSignature,omitempty"`
	MintAmount        uint64 `protobuf:"varint,3,opt,name=mintAmount,proto3" json:"mintAmount,omitempty"`
	PublicKey         string `protobuf:"bytes,4,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	UserAddress       string `protobuf:"bytes,5,opt,name=userAddress,proto3" json:"userAddress,omitempty"`
	ExpiredAt         uint64 `protobuf:"varint,6,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
	Nonce             uint64 `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	EthereumAddress   string `protobuf:"bytes,8,opt,name=ethereumAddress,proto3" json:"ethereumAddress,omitempty"`
}

func (x *MintRequest) Reset() {
	*x = MintRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_programmableLayer_programmableLayer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MintRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MintRequest) ProtoMessage() {}

func (x *MintRequest) ProtoReflect() protoreflect.Message {
	mi := &file_programmableLayer_programmableLayer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MintRequest.ProtoReflect.Descriptor instead.
func (*MintRequest) Descriptor() ([]byte, []int) {
	return file_programmableLayer_programmableLayer_proto_rawDescGZIP(), []int{4}
}

func (x *MintRequest) GetTransferSignature() string {
	if x != nil {
		return x.TransferSignature
	}
	return ""
}

func (x *MintRequest) GetCalldataSignature() string {
	if x != nil {
		return x.CalldataSignature
	}
	return ""
}

func (x *MintRequest) GetMintAmount() uint64 {
	if x != nil {
		return x.MintAmount
	}
	return 0
}

func (x *MintRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *MintRequest) GetUserAddress() string {
	if x != nil {
		return x.UserAddress
	}
	return ""
}

func (x *MintRequest) GetExpiredAt() uint64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

func (x *MintRequest) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *MintRequest) GetEthereumAddress() string {
	if x != nil {
		return x.EthereumAddress
	}
	return ""
}

type MintRequestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	TxHash  string `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

func (x *MintRequestReply) Reset() {
	*x = MintRequestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_programmableLayer_programmableLayer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MintRequestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MintRequestReply) ProtoMessage() {}

func (x *MintRequestReply) ProtoReflect() protoreflect.Message {
	mi := &file_programmableLayer_programmableLayer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MintRequestReply.ProtoReflect.Descriptor instead.
func (*MintRequestReply) Descriptor() ([]byte, []int) {
	return file_programmableLayer_programmableLayer_proto_rawDescGZIP(), []int{5}
}

func (x *MintRequestReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *MintRequestReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *MintRequestReply) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type ClaimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CalldataSignature string `protobuf:"bytes,1,opt,name=calldataSignature,proto3" json:"calldataSignature,omitempty"`
	RoundId           uint64 `protobuf:"varint,2,opt,name=roundId,proto3" json:"roundId,omitempty"`
	PublicKey         string `protobuf:"bytes,3,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	UserAddress       string `protobuf:"bytes,4,opt,name=userAddress,proto3" json:"userAddress,omitempty"`
	ExpiredAt         string `protobuf:"bytes,5,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
	Nonce             uint64 `protobuf:"varint,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
	EthereumAddress   string `protobuf:"bytes,7,opt,name=ethereumAddress,proto3" json:"ethereumAddress,omitempty"`
}

func (x *ClaimRequest) Reset() {
	*x = ClaimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_programmableLayer_programmableLayer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimRequest) ProtoMessage() {}

func (x *ClaimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_programmableLayer_programmableLayer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimRequest.ProtoReflect.Descriptor instead.
func (*ClaimRequest) Descriptor() ([]byte, []int) {
	return file_programmableLayer_programmableLayer_proto_rawDescGZIP(), []int{6}
}

func (x *ClaimRequest) GetCalldataSignature() string {
	if x != nil {
		return x.CalldataSignature
	}
	return ""
}

func (x *ClaimRequest) GetRoundId() uint64 {
	if x != nil {
		return x.RoundId
	}
	return 0
}

func (x *ClaimRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *ClaimRequest) GetUserAddress() string {
	if x != nil {
		return x.UserAddress
	}
	return ""
}

func (x *ClaimRequest) GetExpiredAt() string {
	if x != nil {
		return x.ExpiredAt
	}
	return ""
}

func (x *ClaimRequest) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *ClaimRequest) GetEthereumAddress() string {
	if x != nil {
		return x.EthereumAddress
	}
	return ""
}

type ClaimRequestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	TxHash  string `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

func (x *ClaimRequestReply) Reset() {
	*x = ClaimRequestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_programmableLayer_programmableLayer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimRequestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimRequestReply) ProtoMessage() {}

func (x *ClaimRequestReply) ProtoReflect() protoreflect.Message {
	mi := &file_programmableLayer_programmableLayer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimRequestReply.ProtoReflect.Descriptor instead.
func (*ClaimRequestReply) Descriptor() ([]byte, []int) {
	return file_programmableLayer_programmableLayer_proto_rawDescGZIP(), []int{7}
}

func (x *ClaimRequestReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ClaimRequestReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ClaimRequestReply) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

var File_programmableLayer_programmableLayer_proto protoreflect.FileDescriptor

var file_programmableLayer_programmableLayer_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x22,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x0b, 0x46,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x74, 0x63, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x74, 0x63, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x75, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x72, 0x75,
	0x6e, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x5a, 0x0a, 0x10, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0xa7,
	0x02, 0x0a, 0x0b, 0x4d, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x11,
	0x63, 0x61, 0x6c, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x61, 0x6c, 0x6c, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69,
	0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x6d, 0x69, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75,
	0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5a, 0x0a, 0x10, 0x4d, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78,
	0x48, 0x61, 0x73, 0x68, 0x22, 0xf4, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x61, 0x6c, 0x6c, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x63, 0x61, 0x6c, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5b, 0x0a, 0x11, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x32, 0x55, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32,
	0x69, 0x0a, 0x0d, 0x46, 0x75, 0x6e, 0x64, 0x54, 0x78, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x12, 0x58, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d,
	0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d,
	0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x69, 0x0a, 0x0d, 0x4d, 0x69,
	0x6e, 0x74, 0x54, 0x78, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x11, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x4d, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x6d, 0x0a, 0x0e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x78,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x5b, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x79, 0x65,
	0x72, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x39, 0x5a, 0x37, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x3b, 0x70, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_programmableLayer_programmableLayer_proto_rawDescOnce sync.Once
	file_programmableLayer_programmableLayer_proto_rawDescData = file_programmableLayer_programmableLayer_proto_rawDesc
)

func file_programmableLayer_programmableLayer_proto_rawDescGZIP() []byte {
	file_programmableLayer_programmableLayer_proto_rawDescOnce.Do(func() {
		file_programmableLayer_programmableLayer_proto_rawDescData = protoimpl.X.CompressGZIP(file_programmableLayer_programmableLayer_proto_rawDescData)
	})
	return file_programmableLayer_programmableLayer_proto_rawDescData
}

var file_programmableLayer_programmableLayer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_programmableLayer_programmableLayer_proto_goTypes = []any{
	(*HelloRequest)(nil),      // 0: programmableLayer.HelloRequest
	(*HelloReply)(nil),        // 1: programmableLayer.HelloReply
	(*FundRequest)(nil),       // 2: programmableLayer.FundRequest
	(*FundRequestReply)(nil),  // 3: programmableLayer.FundRequestReply
	(*MintRequest)(nil),       // 4: programmableLayer.MintRequest
	(*MintRequestReply)(nil),  // 5: programmableLayer.MintRequestReply
	(*ClaimRequest)(nil),      // 6: programmableLayer.ClaimRequest
	(*ClaimRequestReply)(nil), // 7: programmableLayer.ClaimRequestReply
}
var file_programmableLayer_programmableLayer_proto_depIdxs = []int32{
	0, // 0: programmableLayer.Greeter.SayHello:input_type -> programmableLayer.HelloRequest
	2, // 1: programmableLayer.FundTxHandler.SubmitFundRequest:input_type -> programmableLayer.FundRequest
	4, // 2: programmableLayer.MintTxHandler.SubmitMintRequest:input_type -> programmableLayer.MintRequest
	6, // 3: programmableLayer.ClaimTxHandler.SubmitClaimRequest:input_type -> programmableLayer.ClaimRequest
	1, // 4: programmableLayer.Greeter.SayHello:output_type -> programmableLayer.HelloReply
	3, // 5: programmableLayer.FundTxHandler.SubmitFundRequest:output_type -> programmableLayer.FundRequestReply
	5, // 6: programmableLayer.MintTxHandler.SubmitMintRequest:output_type -> programmableLayer.MintRequestReply
	7, // 7: programmableLayer.ClaimTxHandler.SubmitClaimRequest:output_type -> programmableLayer.ClaimRequestReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_programmableLayer_programmableLayer_proto_init() }
func file_programmableLayer_programmableLayer_proto_init() {
	if File_programmableLayer_programmableLayer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_programmableLayer_programmableLayer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HelloRequest); i {
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
		file_programmableLayer_programmableLayer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HelloReply); i {
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
		file_programmableLayer_programmableLayer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FundRequest); i {
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
		file_programmableLayer_programmableLayer_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*FundRequestReply); i {
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
		file_programmableLayer_programmableLayer_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MintRequest); i {
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
		file_programmableLayer_programmableLayer_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*MintRequestReply); i {
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
		file_programmableLayer_programmableLayer_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ClaimRequest); i {
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
		file_programmableLayer_programmableLayer_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ClaimRequestReply); i {
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
			RawDescriptor: file_programmableLayer_programmableLayer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_programmableLayer_programmableLayer_proto_goTypes,
		DependencyIndexes: file_programmableLayer_programmableLayer_proto_depIdxs,
		MessageInfos:      file_programmableLayer_programmableLayer_proto_msgTypes,
	}.Build()
	File_programmableLayer_programmableLayer_proto = out.File
	file_programmableLayer_programmableLayer_proto_rawDesc = nil
	file_programmableLayer_programmableLayer_proto_goTypes = nil
	file_programmableLayer_programmableLayer_proto_depIdxs = nil
}
