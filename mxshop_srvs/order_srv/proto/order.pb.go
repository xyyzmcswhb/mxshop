// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: order.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CartItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     int32   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	GoodsId    int32   `protobuf:"varint,3,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	Nums       int32   `protobuf:"varint,4,opt,name=nums,proto3" json:"nums,omitempty"`
	Checked    bool    `protobuf:"varint,5,opt,name=checked,proto3" json:"checked,omitempty"`
	GoodsName  string  `protobuf:"bytes,6,opt,name=goodsName,proto3" json:"goodsName,omitempty"`
	GoodsImage string  `protobuf:"bytes,7,opt,name=goodsImage,proto3" json:"goodsImage,omitempty"`
	GoodsPrice float32 `protobuf:"fixed32,8,opt,name=goodsPrice,proto3" json:"goodsPrice,omitempty"`
}

func (x *CartItemReq) Reset() {
	*x = CartItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemReq) ProtoMessage() {}

func (x *CartItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemReq.ProtoReflect.Descriptor instead.
func (*CartItemReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *CartItemReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CartItemReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CartItemReq) GetGoodsId() int32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *CartItemReq) GetNums() int32 {
	if x != nil {
		return x.Nums
	}
	return 0
}

func (x *CartItemReq) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

func (x *CartItemReq) GetGoodsName() string {
	if x != nil {
		return x.GoodsName
	}
	return ""
}

func (x *CartItemReq) GetGoodsImage() string {
	if x != nil {
		return x.GoodsImage
	}
	return ""
}

func (x *CartItemReq) GetGoodsPrice() float32 {
	if x != nil {
		return x.GoodsPrice
	}
	return 0
}

type ShopCartInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId  int32 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	GoodsId int32 `protobuf:"varint,3,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	Nums    int32 `protobuf:"varint,4,opt,name=nums,proto3" json:"nums,omitempty"`
	Checked bool  `protobuf:"varint,5,opt,name=checked,proto3" json:"checked,omitempty"`
}

func (x *ShopCartInfoResponse) Reset() {
	*x = ShopCartInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopCartInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopCartInfoResponse) ProtoMessage() {}

func (x *ShopCartInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopCartInfoResponse.ProtoReflect.Descriptor instead.
func (*ShopCartInfoResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *ShopCartInfoResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShopCartInfoResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ShopCartInfoResponse) GetGoodsId() int32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *ShopCartInfoResponse) GetNums() int32 {
	if x != nil {
		return x.Nums
	}
	return 0
}

func (x *ShopCartInfoResponse) GetChecked() bool {
	if x != nil {
		return x.Checked
	}
	return false
}

type CartItemListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32                   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*ShopCartInfoResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CartItemListResponse) Reset() {
	*x = CartItemListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemListResponse) ProtoMessage() {}

func (x *CartItemListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemListResponse.ProtoReflect.Descriptor instead.
func (*CartItemListResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *CartItemListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CartItemListResponse) GetData() []*ShopCartInfoResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId  int32  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Name    string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Mobile  string `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Post    string `protobuf:"bytes,6,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *OrderReq) Reset() {
	*x = OrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReq) ProtoMessage() {}

func (x *OrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReq.ProtoReflect.Descriptor instead.
func (*OrderReq) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *OrderReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderReq) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrderReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *OrderReq) GetPost() string {
	if x != nil {
		return x.Post
	}
	return ""
}

type OrderInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId  int32   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	OrderSn string  `protobuf:"bytes,3,opt,name=orderSn,proto3" json:"orderSn,omitempty"`
	PayType string  `protobuf:"bytes,4,opt,name=payType,proto3" json:"payType,omitempty"`
	Status  string  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Post    string  `protobuf:"bytes,6,opt,name=post,proto3" json:"post,omitempty"`
	Total   float32 `protobuf:"fixed32,7,opt,name=total,proto3" json:"total,omitempty"`
	Address string  `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	Name    string  `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	Mobile  string  `protobuf:"bytes,10,opt,name=mobile,proto3" json:"mobile,omitempty"`
	AddTime string  `protobuf:"bytes,11,opt,name=addTime,proto3" json:"addTime,omitempty"`
}

func (x *OrderInfoRsp) Reset() {
	*x = OrderInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfoRsp) ProtoMessage() {}

func (x *OrderInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfoRsp.ProtoReflect.Descriptor instead.
func (*OrderInfoRsp) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderInfoRsp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderInfoRsp) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderInfoRsp) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *OrderInfoRsp) GetPayType() string {
	if x != nil {
		return x.PayType
	}
	return ""
}

func (x *OrderInfoRsp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *OrderInfoRsp) GetPost() string {
	if x != nil {
		return x.Post
	}
	return ""
}

func (x *OrderInfoRsp) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OrderInfoRsp) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OrderInfoRsp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderInfoRsp) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *OrderInfoRsp) GetAddTime() string {
	if x != nil {
		return x.AddTime
	}
	return ""
}

type OrderFilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Page    int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerNums int32 `protobuf:"varint,3,opt,name=PerNums,proto3" json:"PerNums,omitempty"`
}

func (x *OrderFilterRequest) Reset() {
	*x = OrderFilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderFilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderFilterRequest) ProtoMessage() {}

func (x *OrderFilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderFilterRequest.ProtoReflect.Descriptor instead.
func (*OrderFilterRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *OrderFilterRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderFilterRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OrderFilterRequest) GetPerNums() int32 {
	if x != nil {
		return x.PerNums
	}
	return 0
}

type OrderListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32           `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*OrderInfoRsp `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OrderListRsp) Reset() {
	*x = OrderListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListRsp) ProtoMessage() {}

func (x *OrderListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListRsp.ProtoReflect.Descriptor instead.
func (*OrderListRsp) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

func (x *OrderListRsp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *OrderListRsp) GetData() []*OrderInfoRsp {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrderItemRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId    int32   `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	GoodsId    int32   `protobuf:"varint,3,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
	GoodName   string  `protobuf:"bytes,4,opt,name=goodName,proto3" json:"goodName,omitempty"`
	GoodImage  string  `protobuf:"bytes,5,opt,name=goodImage,proto3" json:"goodImage,omitempty"`
	GoodsPrice float32 `protobuf:"fixed32,6,opt,name=goodsPrice,proto3" json:"goodsPrice,omitempty"`
	Nums       int32   `protobuf:"varint,7,opt,name=nums,proto3" json:"nums,omitempty"`
}

func (x *OrderItemRsp) Reset() {
	*x = OrderItemRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemRsp) ProtoMessage() {}

func (x *OrderItemRsp) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemRsp.ProtoReflect.Descriptor instead.
func (*OrderItemRsp) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{8}
}

func (x *OrderItemRsp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderItemRsp) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderItemRsp) GetGoodsId() int32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

func (x *OrderItemRsp) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *OrderItemRsp) GetGoodImage() string {
	if x != nil {
		return x.GoodImage
	}
	return ""
}

func (x *OrderItemRsp) GetGoodsPrice() float32 {
	if x != nil {
		return x.GoodsPrice
	}
	return 0
}

func (x *OrderItemRsp) GetNums() int32 {
	if x != nil {
		return x.Nums
	}
	return 0
}

type OrderInfoDetalRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderInfo *OrderInfoRsp   `protobuf:"bytes,1,opt,name=orderInfo,proto3" json:"orderInfo,omitempty"`
	Goods     []*OrderItemRsp `protobuf:"bytes,2,rep,name=goods,proto3" json:"goods,omitempty"`
}

func (x *OrderInfoDetalRsp) Reset() {
	*x = OrderInfoDetalRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderInfoDetalRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderInfoDetalRsp) ProtoMessage() {}

func (x *OrderInfoDetalRsp) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderInfoDetalRsp.ProtoReflect.Descriptor instead.
func (*OrderInfoDetalRsp) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{9}
}

func (x *OrderInfoDetalRsp) GetOrderInfo() *OrderInfoRsp {
	if x != nil {
		return x.OrderInfo
	}
	return nil
}

func (x *OrderInfoDetalRsp) GetGoods() []*OrderItemRsp {
	if x != nil {
		return x.Goods
	}
	return nil
}

type OrderStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderSn string `protobuf:"bytes,2,opt,name=orderSn,proto3" json:"orderSn,omitempty"`
	Status  string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *OrderStatus) Reset() {
	*x = OrderStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatus) ProtoMessage() {}

func (x *OrderStatus) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatus.ProtoReflect.Descriptor instead.
func (*OrderStatus) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{10}
}

func (x *OrderStatus) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderStatus) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *OrderStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdb, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x14, 0x53, 0x68, 0x6f, 0x70, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e,
	0x75, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x57, 0x0a,
	0x14, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x53, 0x68, 0x6f, 0x70,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8c, 0x01, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x8c, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x4e, 0x75, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x50, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x73,
	0x22, 0x47, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x73, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc0, 0x01, 0x0a, 0x0c, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x6f,
	0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67,
	0x6f, 0x6f, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x65, 0x0a, 0x11,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x6c, 0x52, 0x73,
	0x70, 0x12, 0x2b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x73, 0x70, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23,
	0x0a, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x73, 0x70, 0x52, 0x05, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x22, 0x4f, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x32, 0xa3, 0x03, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x30,
	0x0a, 0x0c, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x09,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0c, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x1a, 0x15, 0x2e, 0x53, 0x68, 0x6f, 0x70, 0x43, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0c, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x36, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x0c, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70,
	0x12, 0x2f, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73,
	0x70, 0x12, 0x2c, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x09, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x6c, 0x52, 0x73, 0x70, 0x12,
	0x39, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_order_proto_goTypes = []any{
	(*UserInfo)(nil),             // 0: UserInfo
	(*CartItemReq)(nil),          // 1: CartItemReq
	(*ShopCartInfoResponse)(nil), // 2: ShopCartInfoResponse
	(*CartItemListResponse)(nil), // 3: CartItemListResponse
	(*OrderReq)(nil),             // 4: OrderReq
	(*OrderInfoRsp)(nil),         // 5: OrderInfoRsp
	(*OrderFilterRequest)(nil),   // 6: OrderFilterRequest
	(*OrderListRsp)(nil),         // 7: OrderListRsp
	(*OrderItemRsp)(nil),         // 8: OrderItemRsp
	(*OrderInfoDetalRsp)(nil),    // 9: OrderInfoDetalRsp
	(*OrderStatus)(nil),          // 10: OrderStatus
	(*emptypb.Empty)(nil),        // 11: google.protobuf.Empty
}
var file_order_proto_depIdxs = []int32{
	2,  // 0: CartItemListResponse.data:type_name -> ShopCartInfoResponse
	5,  // 1: OrderListRsp.data:type_name -> OrderInfoRsp
	5,  // 2: OrderInfoDetalRsp.orderInfo:type_name -> OrderInfoRsp
	8,  // 3: OrderInfoDetalRsp.goods:type_name -> OrderItemRsp
	0,  // 4: Order.CartItemList:input_type -> UserInfo
	1,  // 5: Order.CreateCartItem:input_type -> CartItemReq
	1,  // 6: Order.UpdateCartItem:input_type -> CartItemReq
	1,  // 7: Order.DeleteCartItem:input_type -> CartItemReq
	4,  // 8: Order.CreateOrder:input_type -> OrderReq
	6,  // 9: Order.OrderList:input_type -> OrderFilterRequest
	4,  // 10: Order.OrderDetail:input_type -> OrderReq
	10, // 11: Order.UpdateOrderStatus:input_type -> OrderStatus
	3,  // 12: Order.CartItemList:output_type -> CartItemListResponse
	2,  // 13: Order.CreateCartItem:output_type -> ShopCartInfoResponse
	11, // 14: Order.UpdateCartItem:output_type -> google.protobuf.Empty
	11, // 15: Order.DeleteCartItem:output_type -> google.protobuf.Empty
	5,  // 16: Order.CreateOrder:output_type -> OrderInfoRsp
	7,  // 17: Order.OrderList:output_type -> OrderListRsp
	9,  // 18: Order.OrderDetail:output_type -> OrderInfoDetalRsp
	11, // 19: Order.UpdateOrderStatus:output_type -> google.protobuf.Empty
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserInfo); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CartItemReq); i {
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
		file_order_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ShopCartInfoResponse); i {
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
		file_order_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CartItemListResponse); i {
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
		file_order_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*OrderReq); i {
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
		file_order_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*OrderInfoRsp); i {
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
		file_order_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*OrderFilterRequest); i {
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
		file_order_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*OrderListRsp); i {
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
		file_order_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*OrderItemRsp); i {
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
		file_order_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*OrderInfoDetalRsp); i {
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
		file_order_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*OrderStatus); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}