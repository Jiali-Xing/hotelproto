// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: search.proto

package hotelproto

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

type NearbyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InDate   string `protobuf:"bytes,1,opt,name=InDate,proto3" json:"InDate,omitempty"`
	OutDate  string `protobuf:"bytes,2,opt,name=OutDate,proto3" json:"OutDate,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=Location,proto3" json:"Location,omitempty"`
}

func (x *NearbyRequest) Reset() {
	*x = NearbyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NearbyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NearbyRequest) ProtoMessage() {}

func (x *NearbyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NearbyRequest.ProtoReflect.Descriptor instead.
func (*NearbyRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

func (x *NearbyRequest) GetInDate() string {
	if x != nil {
		return x.InDate
	}
	return ""
}

func (x *NearbyRequest) GetOutDate() string {
	if x != nil {
		return x.OutDate
	}
	return ""
}

func (x *NearbyRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type NearbyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rates []*Rate `protobuf:"bytes,1,rep,name=Rates,proto3" json:"Rates,omitempty"`
}

func (x *NearbyResponse) Reset() {
	*x = NearbyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NearbyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NearbyResponse) ProtoMessage() {}

func (x *NearbyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NearbyResponse.ProtoReflect.Descriptor instead.
func (*NearbyResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

func (x *NearbyResponse) GetRates() []*Rate {
	if x != nil {
		return x.Rates
	}
	return nil
}

type StoreHotelLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelId  string `protobuf:"bytes,1,opt,name=HotelId,proto3" json:"HotelId,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=Location,proto3" json:"Location,omitempty"`
}

func (x *StoreHotelLocationRequest) Reset() {
	*x = StoreHotelLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreHotelLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreHotelLocationRequest) ProtoMessage() {}

func (x *StoreHotelLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreHotelLocationRequest.ProtoReflect.Descriptor instead.
func (*StoreHotelLocationRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{2}
}

func (x *StoreHotelLocationRequest) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *StoreHotelLocationRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type StoreHotelLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelId string `protobuf:"bytes,1,opt,name=HotelId,proto3" json:"HotelId,omitempty"`
}

func (x *StoreHotelLocationResponse) Reset() {
	*x = StoreHotelLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreHotelLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreHotelLocationResponse) ProtoMessage() {}

func (x *StoreHotelLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreHotelLocationResponse.ProtoReflect.Descriptor instead.
func (*StoreHotelLocationResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{3}
}

func (x *StoreHotelLocationResponse) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

var File_search_proto protoreflect.FileDescriptor

var file_search_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x68, 0x6f, 0x74, 0x65, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0d, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x0e, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x52, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x05, 0x52, 0x61, 0x74, 0x65, 0x73, 0x22,
	0x51, 0x0a, 0x19, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x48,
	0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x1a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x49, 0x64, 0x32, 0xb5, 0x01, 0x0a, 0x0d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06,
	0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x12, 0x19, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e,
	0x65, 0x61, 0x72, 0x62, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a,
	0x12, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x48, 0x6f, 0x74,
	0x65, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4a, 0x69, 0x61, 0x6c, 0x69, 0x2d, 0x58, 0x69, 0x6e, 0x67, 0x2f, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_proto_rawDescOnce sync.Once
	file_search_proto_rawDescData = file_search_proto_rawDesc
)

func file_search_proto_rawDescGZIP() []byte {
	file_search_proto_rawDescOnce.Do(func() {
		file_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_proto_rawDescData)
	})
	return file_search_proto_rawDescData
}

var file_search_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_search_proto_goTypes = []interface{}{
	(*NearbyRequest)(nil),              // 0: hotelproto.NearbyRequest
	(*NearbyResponse)(nil),             // 1: hotelproto.NearbyResponse
	(*StoreHotelLocationRequest)(nil),  // 2: hotelproto.StoreHotelLocationRequest
	(*StoreHotelLocationResponse)(nil), // 3: hotelproto.StoreHotelLocationResponse
	(*Rate)(nil),                       // 4: hotelproto.Rate
}
var file_search_proto_depIdxs = []int32{
	4, // 0: hotelproto.NearbyResponse.Rates:type_name -> hotelproto.Rate
	0, // 1: hotelproto.SearchService.Nearby:input_type -> hotelproto.NearbyRequest
	2, // 2: hotelproto.SearchService.StoreHotelLocation:input_type -> hotelproto.StoreHotelLocationRequest
	1, // 3: hotelproto.SearchService.Nearby:output_type -> hotelproto.NearbyResponse
	3, // 4: hotelproto.SearchService.StoreHotelLocation:output_type -> hotelproto.StoreHotelLocationResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
	file_rate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NearbyRequest); i {
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
		file_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NearbyResponse); i {
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
		file_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreHotelLocationRequest); i {
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
		file_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreHotelLocationResponse); i {
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
			RawDescriptor: file_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_proto_goTypes,
		DependencyIndexes: file_search_proto_depIdxs,
		MessageInfos:      file_search_proto_msgTypes,
	}.Build()
	File_search_proto = out.File
	file_search_proto_rawDesc = nil
	file_search_proto_goTypes = nil
	file_search_proto_depIdxs = nil
}
