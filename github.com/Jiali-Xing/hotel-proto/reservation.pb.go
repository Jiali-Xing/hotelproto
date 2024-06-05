// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: reservation.proto

package hotel_proto

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

type FrontendReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelId  string `protobuf:"bytes,1,opt,name=HotelId,proto3" json:"HotelId,omitempty"`
	InDate   string `protobuf:"bytes,2,opt,name=InDate,proto3" json:"InDate,omitempty"`
	OutDate  string `protobuf:"bytes,3,opt,name=OutDate,proto3" json:"OutDate,omitempty"`
	Rooms    int32  `protobuf:"varint,4,opt,name=Rooms,proto3" json:"Rooms,omitempty"`
	Username string `protobuf:"bytes,5,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,6,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *FrontendReservationRequest) Reset() {
	*x = FrontendReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontendReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendReservationRequest) ProtoMessage() {}

func (x *FrontendReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendReservationRequest.ProtoReflect.Descriptor instead.
func (*FrontendReservationRequest) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{0}
}

func (x *FrontendReservationRequest) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *FrontendReservationRequest) GetInDate() string {
	if x != nil {
		return x.InDate
	}
	return ""
}

func (x *FrontendReservationRequest) GetOutDate() string {
	if x != nil {
		return x.OutDate
	}
	return ""
}

func (x *FrontendReservationRequest) GetRooms() int32 {
	if x != nil {
		return x.Rooms
	}
	return 0
}

func (x *FrontendReservationRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FrontendReservationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type FrontendReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *FrontendReservationResponse) Reset() {
	*x = FrontendReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontendReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendReservationResponse) ProtoMessage() {}

func (x *FrontendReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendReservationResponse.ProtoReflect.Descriptor instead.
func (*FrontendReservationResponse) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{1}
}

func (x *FrontendReservationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type MakeReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerName string `protobuf:"bytes,1,opt,name=CustomerName,proto3" json:"CustomerName,omitempty"`
	HotelId      string `protobuf:"bytes,2,opt,name=HotelId,proto3" json:"HotelId,omitempty"`
	InDate       string `protobuf:"bytes,3,opt,name=InDate,proto3" json:"InDate,omitempty"`
	OutDate      string `protobuf:"bytes,4,opt,name=OutDate,proto3" json:"OutDate,omitempty"`
	RoomNumber   int32  `protobuf:"varint,5,opt,name=RoomNumber,proto3" json:"RoomNumber,omitempty"`
}

func (x *MakeReservationRequest) Reset() {
	*x = MakeReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeReservationRequest) ProtoMessage() {}

func (x *MakeReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeReservationRequest.ProtoReflect.Descriptor instead.
func (*MakeReservationRequest) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{2}
}

func (x *MakeReservationRequest) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *MakeReservationRequest) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *MakeReservationRequest) GetInDate() string {
	if x != nil {
		return x.InDate
	}
	return ""
}

func (x *MakeReservationRequest) GetOutDate() string {
	if x != nil {
		return x.OutDate
	}
	return ""
}

func (x *MakeReservationRequest) GetRoomNumber() int32 {
	if x != nil {
		return x.RoomNumber
	}
	return 0
}

type MakeReservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *MakeReservationResponse) Reset() {
	*x = MakeReservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeReservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeReservationResponse) ProtoMessage() {}

func (x *MakeReservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeReservationResponse.ProtoReflect.Descriptor instead.
func (*MakeReservationResponse) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{3}
}

func (x *MakeReservationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CheckAvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerName string   `protobuf:"bytes,1,opt,name=CustomerName,proto3" json:"CustomerName,omitempty"`
	HotelIds     []string `protobuf:"bytes,2,rep,name=HotelIds,proto3" json:"HotelIds,omitempty"`
	InDate       string   `protobuf:"bytes,3,opt,name=InDate,proto3" json:"InDate,omitempty"`
	OutDate      string   `protobuf:"bytes,4,opt,name=OutDate,proto3" json:"OutDate,omitempty"`
	RoomNumber   int32    `protobuf:"varint,5,opt,name=RoomNumber,proto3" json:"RoomNumber,omitempty"`
}

func (x *CheckAvailabilityRequest) Reset() {
	*x = CheckAvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAvailabilityRequest) ProtoMessage() {}

func (x *CheckAvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAvailabilityRequest.ProtoReflect.Descriptor instead.
func (*CheckAvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{4}
}

func (x *CheckAvailabilityRequest) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *CheckAvailabilityRequest) GetHotelIds() []string {
	if x != nil {
		return x.HotelIds
	}
	return nil
}

func (x *CheckAvailabilityRequest) GetInDate() string {
	if x != nil {
		return x.InDate
	}
	return ""
}

func (x *CheckAvailabilityRequest) GetOutDate() string {
	if x != nil {
		return x.OutDate
	}
	return ""
}

func (x *CheckAvailabilityRequest) GetRoomNumber() int32 {
	if x != nil {
		return x.RoomNumber
	}
	return 0
}

type CheckAvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelIds []string `protobuf:"bytes,1,rep,name=HotelIds,proto3" json:"HotelIds,omitempty"`
}

func (x *CheckAvailabilityResponse) Reset() {
	*x = CheckAvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAvailabilityResponse) ProtoMessage() {}

func (x *CheckAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAvailabilityResponse.ProtoReflect.Descriptor instead.
func (*CheckAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{5}
}

func (x *CheckAvailabilityResponse) GetHotelIds() []string {
	if x != nil {
		return x.HotelIds
	}
	return nil
}

var File_reservation_proto protoreflect.FileDescriptor

var file_reservation_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x22, 0xb6, 0x01, 0x0a, 0x1a, 0x46,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x6f, 0x74,
	0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f,
	0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x75,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x37, 0x0a, 0x1b, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xa8, 0x01, 0x0a,
	0x16, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x48, 0x6f,
	0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x52, 0x6f, 0x6f,
	0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x33, 0x0a, 0x17, 0x4d, 0x61, 0x6b, 0x65, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xac, 0x01, 0x0a,
	0x18, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52,
	0x6f, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x19, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x49, 0x64, 0x73, 0x32, 0x9c, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x13, 0x46,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x4d, 0x61, 0x6b,
	0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x11, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x1f, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4a, 0x69, 0x61, 0x6c, 0x69, 0x2d, 0x58, 0x69, 0x6e, 0x67, 0x2f, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reservation_proto_rawDescOnce sync.Once
	file_reservation_proto_rawDescData = file_reservation_proto_rawDesc
)

func file_reservation_proto_rawDescGZIP() []byte {
	file_reservation_proto_rawDescOnce.Do(func() {
		file_reservation_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservation_proto_rawDescData)
	})
	return file_reservation_proto_rawDescData
}

var file_reservation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_reservation_proto_goTypes = []interface{}{
	(*FrontendReservationRequest)(nil),  // 0: hotel.FrontendReservationRequest
	(*FrontendReservationResponse)(nil), // 1: hotel.FrontendReservationResponse
	(*MakeReservationRequest)(nil),      // 2: hotel.MakeReservationRequest
	(*MakeReservationResponse)(nil),     // 3: hotel.MakeReservationResponse
	(*CheckAvailabilityRequest)(nil),    // 4: hotel.CheckAvailabilityRequest
	(*CheckAvailabilityResponse)(nil),   // 5: hotel.CheckAvailabilityResponse
}
var file_reservation_proto_depIdxs = []int32{
	0, // 0: hotel.ReservationService.FrontendReservation:input_type -> hotel.FrontendReservationRequest
	2, // 1: hotel.ReservationService.MakeReservation:input_type -> hotel.MakeReservationRequest
	4, // 2: hotel.ReservationService.CheckAvailability:input_type -> hotel.CheckAvailabilityRequest
	1, // 3: hotel.ReservationService.FrontendReservation:output_type -> hotel.FrontendReservationResponse
	3, // 4: hotel.ReservationService.MakeReservation:output_type -> hotel.MakeReservationResponse
	5, // 5: hotel.ReservationService.CheckAvailability:output_type -> hotel.CheckAvailabilityResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reservation_proto_init() }
func file_reservation_proto_init() {
	if File_reservation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontendReservationRequest); i {
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
		file_reservation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontendReservationResponse); i {
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
		file_reservation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeReservationRequest); i {
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
		file_reservation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeReservationResponse); i {
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
		file_reservation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAvailabilityRequest); i {
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
		file_reservation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckAvailabilityResponse); i {
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
			RawDescriptor: file_reservation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservation_proto_goTypes,
		DependencyIndexes: file_reservation_proto_depIdxs,
		MessageInfos:      file_reservation_proto_msgTypes,
	}.Build()
	File_reservation_proto = out.File
	file_reservation_proto_rawDesc = nil
	file_reservation_proto_goTypes = nil
	file_reservation_proto_depIdxs = nil
}
