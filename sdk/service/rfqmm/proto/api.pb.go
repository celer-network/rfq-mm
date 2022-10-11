// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: service/rfqmm/api.proto

package proto

import (
	common "github.com/celer-network/rfq-mm/sdk/common"
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

type ErrCode int32

const (
	ErrCode_ERROR_UNDEFINED          ErrCode = 0
	ErrCode_ERROR_INVALID_ARGUMENTS  ErrCode = 1
	ErrCode_ERROR_LIQUIDITY_PROVIDER ErrCode = 2
	ErrCode_ERROR_PRICE_PROVIDER     ErrCode = 3
	ErrCode_ERROR_AMOUNT_CALCULATOR  ErrCode = 4
	ErrCode_ERROR_REQUEST_SIGNER     ErrCode = 5
	ErrCode_ERROR_LIQUIDITY_MANAGER  ErrCode = 6
	ErrCode_ERROR_CHAIN_MANAGER      ErrCode = 7
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0: "ERROR_UNDEFINED",
		1: "ERROR_INVALID_ARGUMENTS",
		2: "ERROR_LIQUIDITY_PROVIDER",
		3: "ERROR_PRICE_PROVIDER",
		4: "ERROR_AMOUNT_CALCULATOR",
		5: "ERROR_REQUEST_SIGNER",
		6: "ERROR_LIQUIDITY_MANAGER",
		7: "ERROR_CHAIN_MANAGER",
	}
	ErrCode_value = map[string]int32{
		"ERROR_UNDEFINED":          0,
		"ERROR_INVALID_ARGUMENTS":  1,
		"ERROR_LIQUIDITY_PROVIDER": 2,
		"ERROR_PRICE_PROVIDER":     3,
		"ERROR_AMOUNT_CALCULATOR":  4,
		"ERROR_REQUEST_SIGNER":     5,
		"ERROR_LIQUIDITY_MANAGER":  6,
		"ERROR_CHAIN_MANAGER":      7,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_service_rfqmm_api_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_service_rfqmm_api_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_service_rfqmm_api_proto_rawDescGZIP(), []int{0}
}

type PriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcToken  *common.Token `protobuf:"bytes,1,opt,name=src_token,json=srcToken,proto3" json:"src_token,omitempty"`
	DstToken  *common.Token `protobuf:"bytes,2,opt,name=dst_token,json=dstToken,proto3" json:"dst_token,omitempty"`
	SrcAmount string        `protobuf:"bytes,3,opt,name=src_amount,json=srcAmount,proto3" json:"src_amount,omitempty"`
	DstAmount string        `protobuf:"bytes,4,opt,name=dst_amount,json=dstAmount,proto3" json:"dst_amount,omitempty"`
	// indicates whether the user wants native token on the dst chain (only applicable if the dst token is a native wrap)
	DstNative bool `protobuf:"varint,5,opt,name=dst_native,json=dstNative,proto3" json:"dst_native,omitempty"`
}

func (x *PriceRequest) Reset() {
	*x = PriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfqmm_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceRequest) ProtoMessage() {}

func (x *PriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfqmm_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceRequest.ProtoReflect.Descriptor instead.
func (*PriceRequest) Descriptor() ([]byte, []int) {
	return file_service_rfqmm_api_proto_rawDescGZIP(), []int{0}
}

func (x *PriceRequest) GetSrcToken() *common.Token {
	if x != nil {
		return x.SrcToken
	}
	return nil
}

func (x *PriceRequest) GetDstToken() *common.Token {
	if x != nil {
		return x.DstToken
	}
	return nil
}

func (x *PriceRequest) GetSrcAmount() string {
	if x != nil {
		return x.SrcAmount
	}
	return ""
}

func (x *PriceRequest) GetDstAmount() string {
	if x != nil {
		return x.DstAmount
	}
	return ""
}

func (x *PriceRequest) GetDstNative() bool {
	if x != nil {
		return x.DstNative
	}
	return false
}

type PriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err *common.Err `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	// if receiveAmount is specified in the request, it would be the price of receiveToken in sendToken and vice versa.
	Price *Price `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *PriceResponse) Reset() {
	*x = PriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfqmm_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceResponse) ProtoMessage() {}

func (x *PriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfqmm_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceResponse.ProtoReflect.Descriptor instead.
func (*PriceResponse) Descriptor() ([]byte, []int) {
	return file_service_rfqmm_api_proto_rawDescGZIP(), []int{1}
}

func (x *PriceResponse) GetErr() *common.Err {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *PriceResponse) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

type QuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price *Price `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	Quote *Quote `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`
	// indicates whether the user wants native token on the dst chain (only applicable if the dst token is a native wrap)
	DstNative bool `protobuf:"varint,3,opt,name=dst_native,json=dstNative,proto3" json:"dst_native,omitempty"`
}

func (x *QuoteRequest) Reset() {
	*x = QuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfqmm_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteRequest) ProtoMessage() {}

func (x *QuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfqmm_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteRequest.ProtoReflect.Descriptor instead.
func (*QuoteRequest) Descriptor() ([]byte, []int) {
	return file_service_rfqmm_api_proto_rawDescGZIP(), []int{2}
}

func (x *QuoteRequest) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *QuoteRequest) GetQuote() *Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *QuoteRequest) GetDstNative() bool {
	if x != nil {
		return x.DstNative
	}
	return false
}

type QuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err      *common.Err `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	QuoteSig string      `protobuf:"bytes,2,opt,name=quote_sig,json=quoteSig,proto3" json:"quote_sig,omitempty"`
}

func (x *QuoteResponse) Reset() {
	*x = QuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfqmm_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteResponse) ProtoMessage() {}

func (x *QuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfqmm_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteResponse.ProtoReflect.Descriptor instead.
func (*QuoteResponse) Descriptor() ([]byte, []int) {
	return file_service_rfqmm_api_proto_rawDescGZIP(), []int{3}
}

func (x *QuoteResponse) GetErr() *common.Err {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *QuoteResponse) GetQuoteSig() string {
	if x != nil {
		return x.QuoteSig
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*common.Token `protobuf:"bytes,2,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfqmm_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfqmm_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_service_rfqmm_api_proto_rawDescGZIP(), []int{4}
}

func (x *Config) GetTokens() []*common.Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcToken *common.Token `protobuf:"bytes,1,opt,name=src_token,json=srcToken,proto3" json:"src_token,omitempty"`
	// src_amount reflects the total amount of src_token the user should deposit in the contract on
	// the src chain it should include rfq protocol fee + msg fee + mm charged fee
	SrcAmount        string        `protobuf:"bytes,2,opt,name=src_amount,json=srcAmount,proto3" json:"src_amount,omitempty"`
	DstToken         *common.Token `protobuf:"bytes,3,opt,name=dst_token,json=dstToken,proto3" json:"dst_token,omitempty"`
	SrcReleaseAmount string        `protobuf:"bytes,4,opt,name=src_release_amount,json=srcReleaseAmount,proto3" json:"src_release_amount,omitempty"`
	DstAmount        string        `protobuf:"bytes,5,opt,name=dst_amount,json=dstAmount,proto3" json:"dst_amount,omitempty"`
	// fee = mm fee + msg fee + src tx gas cost + dst tx gas cost
	FeeAmount string `protobuf:"bytes,6,opt,name=fee_amount,json=feeAmount,proto3" json:"fee_amount,omitempty"`
	// unix epoch milliseconds. the time before which the price response is valid for Quote
	ValidThru int64  `protobuf:"varint,7,opt,name=valid_thru,json=validThru,proto3" json:"valid_thru,omitempty"`
	MmAddr    string `protobuf:"bytes,8,opt,name=mm_addr,json=mmAddr,proto3" json:"mm_addr,omitempty"`
	// sig(hash('rfq price', mm_addr, valid_thru, src_chain_id, token_in, amount_in, dst_chain_id, token_out, amount_out))
	// when calling Quote(), mm uses this signature to verify the price content is agreed by them previously
	// and is not beyond deadline.
	Sig string `protobuf:"bytes,9,opt,name=sig,proto3" json:"sig,omitempty"`
	// the maximum src deposit period that is expected by mm, will be started from the time when mm receives the quote request
	SrcDepositPeriod int64 `protobuf:"varint,10,opt,name=src_deposit_period,json=srcDepositPeriod,proto3" json:"src_deposit_period,omitempty"`
	// the minimum dst transfer period that is expected by mm, will be started from the time when mm receives the quote request
	DstTransferPeriod int64 `protobuf:"varint,11,opt,name=dst_transfer_period,json=dstTransferPeriod,proto3" json:"dst_transfer_period,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfqmm_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfqmm_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_service_rfqmm_api_proto_rawDescGZIP(), []int{5}
}

func (x *Price) GetSrcToken() *common.Token {
	if x != nil {
		return x.SrcToken
	}
	return nil
}

func (x *Price) GetSrcAmount() string {
	if x != nil {
		return x.SrcAmount
	}
	return ""
}

func (x *Price) GetDstToken() *common.Token {
	if x != nil {
		return x.DstToken
	}
	return nil
}

func (x *Price) GetSrcReleaseAmount() string {
	if x != nil {
		return x.SrcReleaseAmount
	}
	return ""
}

func (x *Price) GetDstAmount() string {
	if x != nil {
		return x.DstAmount
	}
	return ""
}

func (x *Price) GetFeeAmount() string {
	if x != nil {
		return x.FeeAmount
	}
	return ""
}

func (x *Price) GetValidThru() int64 {
	if x != nil {
		return x.ValidThru
	}
	return 0
}

func (x *Price) GetMmAddr() string {
	if x != nil {
		return x.MmAddr
	}
	return ""
}

func (x *Price) GetSig() string {
	if x != nil {
		return x.Sig
	}
	return ""
}

func (x *Price) GetSrcDepositPeriod() int64 {
	if x != nil {
		return x.SrcDepositPeriod
	}
	return 0
}

func (x *Price) GetDstTransferPeriod() int64 {
	if x != nil {
		return x.DstTransferPeriod
	}
	return 0
}

// some of the fields map to the Quote struct in the RFQ contract
type Quote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the quote hash
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// the input token amount on the src chain
	SrcToken  *common.Token `protobuf:"bytes,2,opt,name=src_token,json=srcToken,proto3" json:"src_token,omitempty"`
	SrcAmount string        `protobuf:"bytes,3,opt,name=src_amount,json=srcAmount,proto3" json:"src_amount,omitempty"`
	// the token amount (same token as src_token) that the market maker will receive by filling this quote
	SrcReleaseAmount string `protobuf:"bytes,4,opt,name=src_release_amount,json=srcReleaseAmount,proto3" json:"src_release_amount,omitempty"`
	// the token amount out on the dst chain to be received by the user
	DstToken  *common.Token `protobuf:"bytes,5,opt,name=dst_token,json=dstToken,proto3" json:"dst_token,omitempty"`
	DstAmount string        `protobuf:"bytes,6,opt,name=dst_amount,json=dstAmount,proto3" json:"dst_amount,omitempty"`
	// the deadline before which the user can submit on the src chain
	SrcDeadline int64 `protobuf:"varint,7,opt,name=src_deadline,json=srcDeadline,proto3" json:"src_deadline,omitempty"`
	// the time after which the order is eligible for refund
	DstDeadline int64 `protobuf:"varint,8,opt,name=dst_deadline,json=dstDeadline,proto3" json:"dst_deadline,omitempty"`
	// nonce that is determined by the server that is used to dedup quotes
	Nonce uint64 `protobuf:"varint,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// sender of the src tx (msg.sender). it's also the user who deposits the src fund
	Sender string `protobuf:"bytes,10,opt,name=sender,proto3" json:"sender,omitempty"`
	// the receiver of the token on the dst chain
	Receiver string `protobuf:"bytes,11,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// the receiver of the refund (if any) on the src chain
	RefundTo string `protobuf:"bytes,12,opt,name=refund_to,json=refundTo,proto3" json:"refund_to,omitempty"`
	// the address of the liquidity provider who's going to transfer fund to the user on the dst chain
	MmAddr string `protobuf:"bytes,13,opt,name=mm_addr,json=mmAddr,proto3" json:"mm_addr,omitempty"`
}

func (x *Quote) Reset() {
	*x = Quote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_rfqmm_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quote) ProtoMessage() {}

func (x *Quote) ProtoReflect() protoreflect.Message {
	mi := &file_service_rfqmm_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quote.ProtoReflect.Descriptor instead.
func (*Quote) Descriptor() ([]byte, []int) {
	return file_service_rfqmm_api_proto_rawDescGZIP(), []int{6}
}

func (x *Quote) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Quote) GetSrcToken() *common.Token {
	if x != nil {
		return x.SrcToken
	}
	return nil
}

func (x *Quote) GetSrcAmount() string {
	if x != nil {
		return x.SrcAmount
	}
	return ""
}

func (x *Quote) GetSrcReleaseAmount() string {
	if x != nil {
		return x.SrcReleaseAmount
	}
	return ""
}

func (x *Quote) GetDstToken() *common.Token {
	if x != nil {
		return x.DstToken
	}
	return nil
}

func (x *Quote) GetDstAmount() string {
	if x != nil {
		return x.DstAmount
	}
	return ""
}

func (x *Quote) GetSrcDeadline() int64 {
	if x != nil {
		return x.SrcDeadline
	}
	return 0
}

func (x *Quote) GetDstDeadline() int64 {
	if x != nil {
		return x.DstDeadline
	}
	return 0
}

func (x *Quote) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Quote) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Quote) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *Quote) GetRefundTo() string {
	if x != nil {
		return x.RefundTo
	}
	return ""
}

func (x *Quote) GetMmAddr() string {
	if x != nil {
		return x.MmAddr
	}
	return ""
}

var File_service_rfqmm_api_proto protoreflect.FileDescriptor

var file_service_rfqmm_api_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x66, 0x71, 0x6d, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x72, 0x66, 0x71, 0x6d, 0x6d, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3,
	0x01, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2a, 0x0a, 0x09, 0x73, 0x72, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x08, 0x73, 0x72, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x09, 0x64,
	0x73, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x08, 0x64,
	0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x72, 0x63, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x72, 0x63,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x73, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x73, 0x74, 0x4e, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x22, 0x5a, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x66,
	0x71, 0x6d, 0x6d, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x85, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x66, 0x71, 0x6d, 0x6d,
	0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a,
	0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x66, 0x71, 0x6d, 0x6d, 0x2e, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64,
	0x73, 0x74, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x22, 0x4b, 0x0a, 0x0d, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x5f, 0x73, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x53, 0x69, 0x67, 0x22, 0x2f, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x25, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x06,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x92, 0x03, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x2a, 0x0a, 0x09, 0x73, 0x72, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x08, 0x73, 0x72, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x72, 0x63, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x72, 0x63, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x64,
	0x73, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x08, 0x64,
	0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x72, 0x63, 0x5f, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x72, 0x63, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x73, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x65, 0x65, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x74, 0x68, 0x72,
	0x75, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x54, 0x68,
	0x72, 0x75, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x67, 0x12, 0x2c, 0x0a,
	0x12, 0x73, 0x72, 0x63, 0x5f, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x5f, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x72, 0x63, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x64,
	0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x64, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0xa5, 0x03, 0x0a, 0x05,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x09, 0x73, 0x72, 0x63,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x08, 0x73, 0x72, 0x63,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x72, 0x63, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x72, 0x63, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x72, 0x63, 0x5f, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x73, 0x72, 0x63, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x64, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x08, 0x64, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x73, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x73, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x72, 0x63, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x72, 0x63, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x73, 0x74, 0x44, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x54, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x6d,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6d, 0x41,
	0x64, 0x64, 0x72, 0x2a, 0xe0, 0x01, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x13, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10,
	0x01, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4c, 0x49, 0x51, 0x55, 0x49,
	0x44, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12,
	0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x50,
	0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x4c, 0x43, 0x55, 0x4c,
	0x41, 0x54, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x52, 0x10, 0x05,
	0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4c, 0x49, 0x51, 0x55, 0x49, 0x44,
	0x49, 0x54, 0x59, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x06, 0x12, 0x17, 0x0a,
	0x13, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x41, 0x4e,
	0x41, 0x47, 0x45, 0x52, 0x10, 0x07, 0x32, 0xc5, 0x01, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x12, 0x5e,
	0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x72, 0x66, 0x71, 0x6d, 0x6d, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72,
	0x66, 0x71, 0x6d, 0x6d, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x66, 0x71, 0x6d, 0x6d, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5e,
	0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x72, 0x66, 0x71, 0x6d, 0x6d, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72,
	0x66, 0x71, 0x6d, 0x6d, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x66, 0x71, 0x6d, 0x6d, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x39,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x6c,
	0x65, 0x72, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x66, 0x71, 0x2d, 0x6d,
	0x6d, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x66,
	0x71, 0x6d, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_rfqmm_api_proto_rawDescOnce sync.Once
	file_service_rfqmm_api_proto_rawDescData = file_service_rfqmm_api_proto_rawDesc
)

func file_service_rfqmm_api_proto_rawDescGZIP() []byte {
	file_service_rfqmm_api_proto_rawDescOnce.Do(func() {
		file_service_rfqmm_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_rfqmm_api_proto_rawDescData)
	})
	return file_service_rfqmm_api_proto_rawDescData
}

var file_service_rfqmm_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_rfqmm_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_rfqmm_api_proto_goTypes = []interface{}{
	(ErrCode)(0),          // 0: service.rfqmm.ErrCode
	(*PriceRequest)(nil),  // 1: service.rfqmm.PriceRequest
	(*PriceResponse)(nil), // 2: service.rfqmm.PriceResponse
	(*QuoteRequest)(nil),  // 3: service.rfqmm.QuoteRequest
	(*QuoteResponse)(nil), // 4: service.rfqmm.QuoteResponse
	(*Config)(nil),        // 5: service.rfqmm.Config
	(*Price)(nil),         // 6: service.rfqmm.Price
	(*Quote)(nil),         // 7: service.rfqmm.Quote
	(*common.Token)(nil),  // 8: common.Token
	(*common.Err)(nil),    // 9: common.Err
}
var file_service_rfqmm_api_proto_depIdxs = []int32{
	8,  // 0: service.rfqmm.PriceRequest.src_token:type_name -> common.Token
	8,  // 1: service.rfqmm.PriceRequest.dst_token:type_name -> common.Token
	9,  // 2: service.rfqmm.PriceResponse.err:type_name -> common.Err
	6,  // 3: service.rfqmm.PriceResponse.price:type_name -> service.rfqmm.Price
	6,  // 4: service.rfqmm.QuoteRequest.price:type_name -> service.rfqmm.Price
	7,  // 5: service.rfqmm.QuoteRequest.quote:type_name -> service.rfqmm.Quote
	9,  // 6: service.rfqmm.QuoteResponse.err:type_name -> common.Err
	8,  // 7: service.rfqmm.Config.tokens:type_name -> common.Token
	8,  // 8: service.rfqmm.Price.src_token:type_name -> common.Token
	8,  // 9: service.rfqmm.Price.dst_token:type_name -> common.Token
	8,  // 10: service.rfqmm.Quote.src_token:type_name -> common.Token
	8,  // 11: service.rfqmm.Quote.dst_token:type_name -> common.Token
	1,  // 12: service.rfqmm.api.Price:input_type -> service.rfqmm.PriceRequest
	3,  // 13: service.rfqmm.api.Quote:input_type -> service.rfqmm.QuoteRequest
	2,  // 14: service.rfqmm.api.Price:output_type -> service.rfqmm.PriceResponse
	4,  // 15: service.rfqmm.api.Quote:output_type -> service.rfqmm.QuoteResponse
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_service_rfqmm_api_proto_init() }
func file_service_rfqmm_api_proto_init() {
	if File_service_rfqmm_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_rfqmm_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceRequest); i {
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
		file_service_rfqmm_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceResponse); i {
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
		file_service_rfqmm_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteRequest); i {
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
		file_service_rfqmm_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteResponse); i {
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
		file_service_rfqmm_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_service_rfqmm_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_service_rfqmm_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quote); i {
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
			RawDescriptor: file_service_rfqmm_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_rfqmm_api_proto_goTypes,
		DependencyIndexes: file_service_rfqmm_api_proto_depIdxs,
		EnumInfos:         file_service_rfqmm_api_proto_enumTypes,
		MessageInfos:      file_service_rfqmm_api_proto_msgTypes,
	}.Build()
	File_service_rfqmm_api_proto = out.File
	file_service_rfqmm_api_proto_rawDesc = nil
	file_service_rfqmm_api_proto_goTypes = nil
	file_service_rfqmm_api_proto_depIdxs = nil
}