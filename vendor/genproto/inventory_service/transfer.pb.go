// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: transfer.proto

package inventory_service

import (
	common "genproto/common"
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

type CreateTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DepatureShopId string          `protobuf:"bytes,1,opt,name=depature_shop_id,json=depatureShopId,proto3" json:"depature_shop_id,omitempty"`
	ArrivalShopId  string          `protobuf:"bytes,2,opt,name=arrival_shop_id,json=arrivalShopId,proto3" json:"arrival_shop_id,omitempty"`
	TransferName   string          `protobuf:"bytes,3,opt,name=transfer_name,json=transferName,proto3" json:"transfer_name,omitempty"`
	Request        *common.Request `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *CreateTransferRequest) Reset() {
	*x = CreateTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransferRequest) ProtoMessage() {}

func (x *CreateTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransferRequest.ProtoReflect.Descriptor instead.
func (*CreateTransferRequest) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTransferRequest) GetDepatureShopId() string {
	if x != nil {
		return x.DepatureShopId
	}
	return ""
}

func (x *CreateTransferRequest) GetArrivalShopId() string {
	if x != nil {
		return x.ArrivalShopId
	}
	return ""
}

func (x *CreateTransferRequest) GetTransferName() string {
	if x != nil {
		return x.TransferName
	}
	return ""
}

func (x *CreateTransferRequest) GetRequest() *common.Request {
	if x != nil {
		return x.Request
	}
	return nil
}

type AddItemToTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId string          `protobuf:"bytes,1,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"`
	ProductId  string          `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Amount     float32         `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Request    *common.Request `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *AddItemToTransferRequest) Reset() {
	*x = AddItemToTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemToTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemToTransferRequest) ProtoMessage() {}

func (x *AddItemToTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemToTransferRequest.ProtoReflect.Descriptor instead.
func (*AddItemToTransferRequest) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *AddItemToTransferRequest) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

func (x *AddItemToTransferRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddItemToTransferRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *AddItemToTransferRequest) GetRequest() *common.Request {
	if x != nil {
		return x.Request
	}
	return nil
}

type ShortTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedBy    *common.ShortUser `protobuf:"bytes,2,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt    string            `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	AcceptedBy   string            `protobuf:"bytes,4,opt,name=accepted_by,json=acceptedBy,proto3" json:"accepted_by,omitempty"`
	Status       *common.Status    `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	DepatureShop *common.ShortShop `protobuf:"bytes,6,opt,name=depature_shop,json=depatureShop,proto3" json:"depature_shop,omitempty"`
	ArrivalShop  *common.ShortShop `protobuf:"bytes,7,opt,name=arrival_shop,json=arrivalShop,proto3" json:"arrival_shop,omitempty"`
	TotalPrice   float32           `protobuf:"fixed32,8,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *ShortTransfer) Reset() {
	*x = ShortTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortTransfer) ProtoMessage() {}

func (x *ShortTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortTransfer.ProtoReflect.Descriptor instead.
func (*ShortTransfer) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *ShortTransfer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShortTransfer) GetCreatedBy() *common.ShortUser {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *ShortTransfer) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ShortTransfer) GetAcceptedBy() string {
	if x != nil {
		return x.AcceptedBy
	}
	return ""
}

func (x *ShortTransfer) GetStatus() *common.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ShortTransfer) GetDepatureShop() *common.ShortShop {
	if x != nil {
		return x.DepatureShop
	}
	return nil
}

func (x *ShortTransfer) GetArrivalShop() *common.ShortShop {
	if x != nil {
		return x.ArrivalShop
	}
	return nil
}

func (x *ShortTransfer) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type GetAllTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *common.SearchRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetAllTransferRequest) Reset() {
	*x = GetAllTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTransferRequest) ProtoMessage() {}

func (x *GetAllTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTransferRequest.ProtoReflect.Descriptor instead.
func (*GetAllTransferRequest) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllTransferRequest) GetRequest() *common.SearchRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetAllTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*ShortTransfer `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int64            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetAllTransferResponse) Reset() {
	*x = GetAllTransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTransferResponse) ProtoMessage() {}

func (x *GetAllTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTransferResponse.ProtoReflect.Descriptor instead.
func (*GetAllTransferResponse) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllTransferResponse) GetData() []*ShortTransfer {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetAllTransferResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type TransferItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProdcutId string   `protobuf:"bytes,2,opt,name=prodcut_id,json=prodcutId,proto3" json:"prodcut_id,omitempty"`
	Sku       string   `protobuf:"bytes,3,opt,name=sku,proto3" json:"sku,omitempty"`
	Barcodes  []string `protobuf:"bytes,4,rep,name=barcodes,proto3" json:"barcodes,omitempty"`
	Name      string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TransferItem) Reset() {
	*x = TransferItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferItem) ProtoMessage() {}

func (x *TransferItem) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferItem.ProtoReflect.Descriptor instead.
func (*TransferItem) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{5}
}

func (x *TransferItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TransferItem) GetProdcutId() string {
	if x != nil {
		return x.ProdcutId
	}
	return ""
}

func (x *TransferItem) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *TransferItem) GetBarcodes() []string {
	if x != nil {
		return x.Barcodes
	}
	return nil
}

func (x *TransferItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAllTransferItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *common.SearchRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetAllTransferItemsRequest) Reset() {
	*x = GetAllTransferItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTransferItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTransferItemsRequest) ProtoMessage() {}

func (x *GetAllTransferItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTransferItemsRequest.ProtoReflect.Descriptor instead.
func (*GetAllTransferItemsRequest) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllTransferItemsRequest) GetRequest() *common.SearchRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetAllTransferItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data                []*TransferItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	TotalProducts       int64           `protobuf:"varint,2,opt,name=total_products,json=totalProducts,proto3" json:"total_products,omitempty"`
	TotalDepaturedPrice float32         `protobuf:"fixed32,3,opt,name=total_depatured_price,json=totalDepaturedPrice,proto3" json:"total_depatured_price,omitempty"`
	TotalArrivedPrice   float32         `protobuf:"fixed32,4,opt,name=total_arrived_price,json=totalArrivedPrice,proto3" json:"total_arrived_price,omitempty"`
}

func (x *GetAllTransferItemsResponse) Reset() {
	*x = GetAllTransferItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTransferItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTransferItemsResponse) ProtoMessage() {}

func (x *GetAllTransferItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTransferItemsResponse.ProtoReflect.Descriptor instead.
func (*GetAllTransferItemsResponse) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllTransferItemsResponse) GetData() []*TransferItem {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetAllTransferItemsResponse) GetTotalProducts() int64 {
	if x != nil {
		return x.TotalProducts
	}
	return 0
}

func (x *GetAllTransferItemsResponse) GetTotalDepaturedPrice() float32 {
	if x != nil {
		return x.TotalDepaturedPrice
	}
	return 0
}

func (x *GetAllTransferItemsResponse) GetTotalArrivedPrice() float32 {
	if x != nil {
		return x.TotalArrivedPrice
	}
	return 0
}

var File_transfer_proto protoreflect.FileDescriptor

var file_transfer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x68, 0x6f,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x72,
	0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x53, 0x68, 0x6f, 0x70,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x18,
	0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x22, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xac, 0x02, 0x0a, 0x0d, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2f, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x73,
	0x68, 0x6f, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53,
	0x68, 0x6f, 0x70, 0x12, 0x2d, 0x0a, 0x0c, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x73,
	0x68, 0x6f, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x0b, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x53, 0x68,
	0x6f, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x22, 0x41, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x7f, 0x0a, 0x0c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x63, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x63, 0x75, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x1a, 0x0a, 0x08, 0x62,
	0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x62,
	0x61, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xcb, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x32, 0x0a,
	0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x61, 0x74, 0x75, 0x72, 0x65, 0x64,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x44, 0x65, 0x70, 0x61, 0x74, 0x75, 0x72, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x72, 0x72, 0x69, 0x76,
	0x65, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x72, 0x72, 0x69, 0x76, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transfer_proto_rawDescOnce sync.Once
	file_transfer_proto_rawDescData = file_transfer_proto_rawDesc
)

func file_transfer_proto_rawDescGZIP() []byte {
	file_transfer_proto_rawDescOnce.Do(func() {
		file_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfer_proto_rawDescData)
	})
	return file_transfer_proto_rawDescData
}

var file_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_transfer_proto_goTypes = []interface{}{
	(*CreateTransferRequest)(nil),       // 0: CreateTransferRequest
	(*AddItemToTransferRequest)(nil),    // 1: AddItemToTransferRequest
	(*ShortTransfer)(nil),               // 2: ShortTransfer
	(*GetAllTransferRequest)(nil),       // 3: GetAllTransferRequest
	(*GetAllTransferResponse)(nil),      // 4: GetAllTransferResponse
	(*TransferItem)(nil),                // 5: TransferItem
	(*GetAllTransferItemsRequest)(nil),  // 6: GetAllTransferItemsRequest
	(*GetAllTransferItemsResponse)(nil), // 7: GetAllTransferItemsResponse
	(*common.Request)(nil),              // 8: Request
	(*common.ShortUser)(nil),            // 9: ShortUser
	(*common.Status)(nil),               // 10: Status
	(*common.ShortShop)(nil),            // 11: ShortShop
	(*common.SearchRequest)(nil),        // 12: SearchRequest
}
var file_transfer_proto_depIdxs = []int32{
	8,  // 0: CreateTransferRequest.request:type_name -> Request
	8,  // 1: AddItemToTransferRequest.request:type_name -> Request
	9,  // 2: ShortTransfer.created_by:type_name -> ShortUser
	10, // 3: ShortTransfer.status:type_name -> Status
	11, // 4: ShortTransfer.depature_shop:type_name -> ShortShop
	11, // 5: ShortTransfer.arrival_shop:type_name -> ShortShop
	12, // 6: GetAllTransferRequest.request:type_name -> SearchRequest
	2,  // 7: GetAllTransferResponse.data:type_name -> ShortTransfer
	12, // 8: GetAllTransferItemsRequest.request:type_name -> SearchRequest
	5,  // 9: GetAllTransferItemsResponse.data:type_name -> TransferItem
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_transfer_proto_init() }
func file_transfer_proto_init() {
	if File_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransferRequest); i {
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
		file_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemToTransferRequest); i {
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
		file_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortTransfer); i {
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
		file_transfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTransferRequest); i {
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
		file_transfer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTransferResponse); i {
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
		file_transfer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferItem); i {
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
		file_transfer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTransferItemsRequest); i {
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
		file_transfer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTransferItemsResponse); i {
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
			RawDescriptor: file_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transfer_proto_goTypes,
		DependencyIndexes: file_transfer_proto_depIdxs,
		MessageInfos:      file_transfer_proto_msgTypes,
	}.Build()
	File_transfer_proto = out.File
	file_transfer_proto_rawDesc = nil
	file_transfer_proto_goTypes = nil
	file_transfer_proto_depIdxs = nil
}
