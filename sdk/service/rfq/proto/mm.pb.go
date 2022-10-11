// API for Market Makers

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: service/rfq/mm.proto

package proto

import (
	proto "github.com/celer-network/rfq-mm/sdk/service/rfqmm/proto"
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

type PendingOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PendingOrdersRequest) Reset() {
	*x = PendingOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfq_mm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingOrdersRequest) ProtoMessage() {}

func (x *PendingOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfq_mm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingOrdersRequest.ProtoReflect.Descriptor instead.
func (*PendingOrdersRequest) Descriptor() ([]byte, []int) {
	return file_service_rfq_mm_proto_rawDescGZIP(), []int{0}
}

type PendingOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*PendingOrder `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *PendingOrdersResponse) Reset() {
	*x = PendingOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfq_mm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingOrdersResponse) ProtoMessage() {}

func (x *PendingOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfq_mm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingOrdersResponse.ProtoReflect.Descriptor instead.
func (*PendingOrdersResponse) Descriptor() ([]byte, []int) {
	return file_service_rfq_mm_proto_rawDescGZIP(), []int{1}
}

func (x *PendingOrdersResponse) GetOrders() []*PendingOrder {
	if x != nil {
		return x.Orders
	}
	return nil
}

type OrderUpdates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuoteHash   string      `protobuf:"bytes,1,opt,name=quote_hash,json=quoteHash,proto3" json:"quote_hash,omitempty"`
	OrderStatus OrderStatus `protobuf:"varint,2,opt,name=order_status,json=orderStatus,proto3,enum=service.rfq.OrderStatus" json:"order_status,omitempty"`
	ExecTxHash  string      `protobuf:"bytes,3,opt,name=exec_tx_hash,json=execTxHash,proto3" json:"exec_tx_hash,omitempty"`
}

func (x *OrderUpdates) Reset() {
	*x = OrderUpdates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfq_mm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUpdates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpdates) ProtoMessage() {}

func (x *OrderUpdates) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfq_mm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpdates.ProtoReflect.Descriptor instead.
func (*OrderUpdates) Descriptor() ([]byte, []int) {
	return file_service_rfq_mm_proto_rawDescGZIP(), []int{2}
}

func (x *OrderUpdates) GetQuoteHash() string {
	if x != nil {
		return x.QuoteHash
	}
	return ""
}

func (x *OrderUpdates) GetOrderStatus() OrderStatus {
	if x != nil {
		return x.OrderStatus
	}
	return OrderStatus_STATUS_PENDING
}

func (x *OrderUpdates) GetExecTxHash() string {
	if x != nil {
		return x.ExecTxHash
	}
	return ""
}

type UpdateOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updates []*OrderUpdates `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
}

func (x *UpdateOrdersRequest) Reset() {
	*x = UpdateOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfq_mm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrdersRequest) ProtoMessage() {}

func (x *UpdateOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfq_mm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrdersRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrdersRequest) Descriptor() ([]byte, []int) {
	return file_service_rfq_mm_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOrdersRequest) GetUpdates() []*OrderUpdates {
	if x != nil {
		return x.Updates
	}
	return nil
}

type UpdateOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateOrdersResponse) Reset() {
	*x = UpdateOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfq_mm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrdersResponse) ProtoMessage() {}

func (x *UpdateOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfq_mm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrdersResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrdersResponse) Descriptor() ([]byte, []int) {
	return file_service_rfq_mm_proto_rawDescGZIP(), []int{4}
}

type UpdateConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *proto.Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *UpdateConfigsRequest) Reset() {
	*x = UpdateConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfq_mm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigsRequest) ProtoMessage() {}

func (x *UpdateConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfq_mm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigsRequest.ProtoReflect.Descriptor instead.
func (*UpdateConfigsRequest) Descriptor() ([]byte, []int) {
	return file_service_rfq_mm_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateConfigsRequest) GetConfig() *proto.Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type UpdateConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateConfigsResponse) Reset() {
	*x = UpdateConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfq_mm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigsResponse) ProtoMessage() {}

func (x *UpdateConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfq_mm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigsResponse.ProtoReflect.Descriptor instead.
func (*UpdateConfigsResponse) Descriptor() ([]byte, []int) {
	return file_service_rfq_mm_proto_rawDescGZIP(), []int{6}
}

var File_service_rfq_mm_proto protoreflect.FileDescriptor

var file_service_rfq_mm_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x66, 0x71, 0x2f, 0x6d, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x72, 0x66, 0x71, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x66, 0x71, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x72, 0x66, 0x71, 0x6d, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x15, 0x50,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72,
	0x66, 0x71, 0x2e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x3b, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x66, 0x71, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x74, 0x78, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x65, 0x63,
	0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0x4a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x66, 0x71, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x66, 0x71,
	0x6d, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf1, 0x02, 0x0a, 0x05, 0x4d,
	0x4d, 0x41, 0x70, 0x69, 0x12, 0x7a, 0x0a, 0x0d, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x72, 0x66, 0x71, 0x2e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x72, 0x66, 0x71, 0x2e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6d, 0x2f, 0x67, 0x65, 0x74,
	0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x73, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x66, 0x71, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x66, 0x71,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x6d, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x77, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x72, 0x66, 0x71, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x72, 0x66, 0x71, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6d, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x6c,
	0x65, 0x72, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x66, 0x71, 0x2d, 0x6d,
	0x6d, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x66,
	0x71, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_rfq_mm_proto_rawDescOnce sync.Once
	file_service_rfq_mm_proto_rawDescData = file_service_rfq_mm_proto_rawDesc
)

func file_service_rfq_mm_proto_rawDescGZIP() []byte {
	file_service_rfq_mm_proto_rawDescOnce.Do(func() {
		file_service_rfq_mm_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_rfq_mm_proto_rawDescData)
	})
	return file_service_rfq_mm_proto_rawDescData
}

var file_service_rfq_mm_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_rfq_mm_proto_goTypes = []interface{}{
	(*PendingOrdersRequest)(nil),  // 0: service.rfq.PendingOrdersRequest
	(*PendingOrdersResponse)(nil), // 1: service.rfq.PendingOrdersResponse
	(*OrderUpdates)(nil),          // 2: service.rfq.OrderUpdates
	(*UpdateOrdersRequest)(nil),   // 3: service.rfq.UpdateOrdersRequest
	(*UpdateOrdersResponse)(nil),  // 4: service.rfq.UpdateOrdersResponse
	(*UpdateConfigsRequest)(nil),  // 5: service.rfq.UpdateConfigsRequest
	(*UpdateConfigsResponse)(nil), // 6: service.rfq.UpdateConfigsResponse
	(*PendingOrder)(nil),          // 7: service.rfq.PendingOrder
	(OrderStatus)(0),              // 8: service.rfq.OrderStatus
	(*proto.Config)(nil),          // 9: service.rfqmm.Config
}
var file_service_rfq_mm_proto_depIdxs = []int32{
	7, // 0: service.rfq.PendingOrdersResponse.orders:type_name -> service.rfq.PendingOrder
	8, // 1: service.rfq.OrderUpdates.order_status:type_name -> service.rfq.OrderStatus
	2, // 2: service.rfq.UpdateOrdersRequest.updates:type_name -> service.rfq.OrderUpdates
	9, // 3: service.rfq.UpdateConfigsRequest.config:type_name -> service.rfqmm.Config
	0, // 4: service.rfq.MMApi.PendingOrders:input_type -> service.rfq.PendingOrdersRequest
	3, // 5: service.rfq.MMApi.UpdateOrders:input_type -> service.rfq.UpdateOrdersRequest
	5, // 6: service.rfq.MMApi.UpdateConfigs:input_type -> service.rfq.UpdateConfigsRequest
	1, // 7: service.rfq.MMApi.PendingOrders:output_type -> service.rfq.PendingOrdersResponse
	4, // 8: service.rfq.MMApi.UpdateOrders:output_type -> service.rfq.UpdateOrdersResponse
	6, // 9: service.rfq.MMApi.UpdateConfigs:output_type -> service.rfq.UpdateConfigsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_rfq_mm_proto_init() }
func file_service_rfq_mm_proto_init() {
	if File_service_rfq_mm_proto != nil {
		return
	}
	file_service_rfq_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_rfq_mm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingOrdersRequest); i {
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
		file_service_rfq_mm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingOrdersResponse); i {
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
		file_service_rfq_mm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderUpdates); i {
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
		file_service_rfq_mm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrdersRequest); i {
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
		file_service_rfq_mm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrdersResponse); i {
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
		file_service_rfq_mm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigsRequest); i {
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
		file_service_rfq_mm_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigsResponse); i {
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
			RawDescriptor: file_service_rfq_mm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_rfq_mm_proto_goTypes,
		DependencyIndexes: file_service_rfq_mm_proto_depIdxs,
		MessageInfos:      file_service_rfq_mm_proto_msgTypes,
	}.Build()
	File_service_rfq_mm_proto = out.File
	file_service_rfq_mm_proto_rawDesc = nil
	file_service_rfq_mm_proto_goTypes = nil
	file_service_rfq_mm_proto_depIdxs = nil
}