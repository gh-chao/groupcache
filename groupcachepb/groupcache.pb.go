//
//Copyright 2012 Google Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: groupcache.proto

package groupcachepb

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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *string `protobuf:"bytes,1,req,name=group" json:"group,omitempty"`
	Key   *string `protobuf:"bytes,2,req,name=key" json:"key,omitempty"` // not actually required/guaranteed to be UTF-8
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupcache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groupcache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_groupcache_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetGroup() string {
	if x != nil && x.Group != nil {
		return *x.Group
	}
	return ""
}

func (x *GetRequest) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value     []byte   `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	MinuteQps *float64 `protobuf:"fixed64,2,opt,name=minute_qps,json=minuteQps" json:"minute_qps,omitempty"`
	Expire    *int64   `protobuf:"varint,3,opt,name=expire" json:"expire,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupcache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groupcache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_groupcache_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *GetResponse) GetMinuteQps() float64 {
	if x != nil && x.MinuteQps != nil {
		return *x.MinuteQps
	}
	return 0
}

func (x *GetResponse) GetExpire() int64 {
	if x != nil && x.Expire != nil {
		return *x.Expire
	}
	return 0
}

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group  *string `protobuf:"bytes,1,req,name=group" json:"group,omitempty"`
	Key    *string `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Value  []byte  `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Expire *int64  `protobuf:"varint,4,opt,name=expire" json:"expire,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupcache_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groupcache_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_groupcache_proto_rawDescGZIP(), []int{2}
}

func (x *SetRequest) GetGroup() string {
	if x != nil && x.Group != nil {
		return *x.Group
	}
	return ""
}

func (x *SetRequest) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *SetRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *SetRequest) GetExpire() int64 {
	if x != nil && x.Expire != nil {
		return *x.Expire
	}
	return 0
}

var File_groupcache_proto protoreflect.FileDescriptor

var file_groupcache_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62,
	0x22, 0x34, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x5a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x5f, 0x71, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x51, 0x70, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x22, 0x62, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x32, 0x4a, 0x0a, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x68, 0x2d, 0x63, 0x68, 0x61, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x62,
}

var (
	file_groupcache_proto_rawDescOnce sync.Once
	file_groupcache_proto_rawDescData = file_groupcache_proto_rawDesc
)

func file_groupcache_proto_rawDescGZIP() []byte {
	file_groupcache_proto_rawDescOnce.Do(func() {
		file_groupcache_proto_rawDescData = protoimpl.X.CompressGZIP(file_groupcache_proto_rawDescData)
	})
	return file_groupcache_proto_rawDescData
}

var file_groupcache_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_groupcache_proto_goTypes = []any{
	(*GetRequest)(nil),  // 0: groupcachepb.GetRequest
	(*GetResponse)(nil), // 1: groupcachepb.GetResponse
	(*SetRequest)(nil),  // 2: groupcachepb.SetRequest
}
var file_groupcache_proto_depIdxs = []int32{
	0, // 0: groupcachepb.GroupCache.Get:input_type -> groupcachepb.GetRequest
	1, // 1: groupcachepb.GroupCache.Get:output_type -> groupcachepb.GetResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_groupcache_proto_init() }
func file_groupcache_proto_init() {
	if File_groupcache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_groupcache_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetRequest); i {
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
		file_groupcache_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetResponse); i {
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
		file_groupcache_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SetRequest); i {
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
			RawDescriptor: file_groupcache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_groupcache_proto_goTypes,
		DependencyIndexes: file_groupcache_proto_depIdxs,
		MessageInfos:      file_groupcache_proto_msgTypes,
	}.Build()
	File_groupcache_proto = out.File
	file_groupcache_proto_rawDesc = nil
	file_groupcache_proto_goTypes = nil
	file_groupcache_proto_depIdxs = nil
}
