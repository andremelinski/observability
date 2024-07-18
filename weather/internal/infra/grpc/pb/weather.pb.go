// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.25.1
// source: proto/weather.proto

package pb

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

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temp_C float32 `protobuf:"fixed32,1,opt,name=temp_C,json=tempC,proto3" json:"temp_C,omitempty"`
	Temp_F float32 `protobuf:"fixed32,2,opt,name=temp_F,json=tempF,proto3" json:"temp_F,omitempty"`
	Temp_K float32 `protobuf:"fixed32,3,opt,name=temp_K,json=tempK,proto3" json:"temp_K,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderResponse) GetTemp_C() float32 {
	if x != nil {
		return x.Temp_C
	}
	return 0
}

func (x *CreateOrderResponse) GetTemp_F() float32 {
	if x != nil {
		return x.Temp_F
	}
	return 0
}

func (x *CreateOrderResponse) GetTemp_K() float32 {
	if x != nil {
		return x.Temp_K
	}
	return 0
}

type WeatherLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Place string `protobuf:"bytes,1,opt,name=place,proto3" json:"place,omitempty"`
}

func (x *WeatherLocationRequest) Reset() {
	*x = WeatherLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherLocationRequest) ProtoMessage() {}

func (x *WeatherLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherLocationRequest.ProtoReflect.Descriptor instead.
func (*WeatherLocationRequest) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{1}
}

func (x *WeatherLocationRequest) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

var File_proto_weather_proto protoreflect.FileDescriptor

var file_proto_weather_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x5a, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x43, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x65, 0x6d, 0x70, 0x5f,
	0x46, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x65, 0x6d, 0x70, 0x46, 0x12, 0x15,
	0x0a, 0x06, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x4b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x74, 0x65, 0x6d, 0x70, 0x4b, 0x22, 0x2e, 0x0a, 0x16, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x32, 0x5f, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_weather_proto_rawDescOnce sync.Once
	file_proto_weather_proto_rawDescData = file_proto_weather_proto_rawDesc
)

func file_proto_weather_proto_rawDescGZIP() []byte {
	file_proto_weather_proto_rawDescOnce.Do(func() {
		file_proto_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_weather_proto_rawDescData)
	})
	return file_proto_weather_proto_rawDescData
}

var file_proto_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_weather_proto_goTypes = []interface{}{
	(*CreateOrderResponse)(nil),    // 0: pb.CreateOrderResponse
	(*WeatherLocationRequest)(nil), // 1: pb.WeatherLocationRequest
}
var file_proto_weather_proto_depIdxs = []int32{
	1, // 0: pb.WeatherService.GetLocationTemperature:input_type -> pb.WeatherLocationRequest
	0, // 1: pb.WeatherService.GetLocationTemperature:output_type -> pb.CreateOrderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_weather_proto_init() }
func file_proto_weather_proto_init() {
	if File_proto_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_proto_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherLocationRequest); i {
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
			RawDescriptor: file_proto_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_weather_proto_goTypes,
		DependencyIndexes: file_proto_weather_proto_depIdxs,
		MessageInfos:      file_proto_weather_proto_msgTypes,
	}.Build()
	File_proto_weather_proto = out.File
	file_proto_weather_proto_rawDesc = nil
	file_proto_weather_proto_goTypes = nil
	file_proto_weather_proto_depIdxs = nil
}
