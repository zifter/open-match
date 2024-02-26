// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0
// source: api/matchfunction.proto

package pb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A MatchProfile defines constraints of Tickets in a Match and shapes the Match proposed by the MatchFunction.
	Profile *MatchProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_matchfunction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_matchfunction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_api_matchfunction_proto_rawDescGZIP(), []int{0}
}

func (x *RunRequest) GetProfile() *MatchProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A Proposal represents a Match candidate that satifies the constraints defined in the input Profile.
	// A valid Proposal response will contain at least one ticket.
	Proposal *Match `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_matchfunction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_matchfunction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_api_matchfunction_proto_rawDescGZIP(), []int{1}
}

func (x *RunResponse) GetProposal() *Match {
	if x != nil {
		return x.Proposal
	}
	return nil
}

var File_api_matchfunction_proto protoreflect.FileDescriptor

var file_api_matchfunction_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x3b, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x32, 0x69, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x15, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x72, 0x75, 0x6e, 0x30, 0x01, 0x42,
	0x91, 0x03, 0x92, 0x41, 0xdf, 0x02, 0x12, 0xb8, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x20, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x49, 0x0a, 0x0a, 0x4f, 0x70, 0x65,
	0x6e, 0x20, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x1a,
	0x23, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x64, 0x69, 0x73, 0x63,
	0x75, 0x73, 0x73, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x56, 0x0a, 0x12, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x20, 0x32,
	0x2e, 0x30, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x66, 0x6f, 0x72, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x2d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x03, 0x31, 0x2e,
	0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34,
	0x12, 0x34, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65,
	0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64,
	0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06,
	0x0a, 0x04, 0x9a, 0x02, 0x01, 0x07, 0x72, 0x3d, 0x0a, 0x18, 0x4f, 0x70, 0x65, 0x6e, 0x20, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x20, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x2d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x2f,
	0x64, 0x6f, 0x63, 0x73, 0x2f, 0x5a, 0x20, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0xaa, 0x02, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_matchfunction_proto_rawDescOnce sync.Once
	file_api_matchfunction_proto_rawDescData = file_api_matchfunction_proto_rawDesc
)

func file_api_matchfunction_proto_rawDescGZIP() []byte {
	file_api_matchfunction_proto_rawDescOnce.Do(func() {
		file_api_matchfunction_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_matchfunction_proto_rawDescData)
	})
	return file_api_matchfunction_proto_rawDescData
}

var file_api_matchfunction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_matchfunction_proto_goTypes = []interface{}{
	(*RunRequest)(nil),   // 0: openmatch.RunRequest
	(*RunResponse)(nil),  // 1: openmatch.RunResponse
	(*MatchProfile)(nil), // 2: openmatch.MatchProfile
	(*Match)(nil),        // 3: openmatch.Match
}
var file_api_matchfunction_proto_depIdxs = []int32{
	2, // 0: openmatch.RunRequest.profile:type_name -> openmatch.MatchProfile
	3, // 1: openmatch.RunResponse.proposal:type_name -> openmatch.Match
	0, // 2: openmatch.MatchFunction.Run:input_type -> openmatch.RunRequest
	1, // 3: openmatch.MatchFunction.Run:output_type -> openmatch.RunResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_matchfunction_proto_init() }
func file_api_matchfunction_proto_init() {
	if File_api_matchfunction_proto != nil {
		return
	}
	file_api_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_matchfunction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunRequest); i {
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
		file_api_matchfunction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
			RawDescriptor: file_api_matchfunction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_matchfunction_proto_goTypes,
		DependencyIndexes: file_api_matchfunction_proto_depIdxs,
		MessageInfos:      file_api_matchfunction_proto_msgTypes,
	}.Build()
	File_api_matchfunction_proto = out.File
	file_api_matchfunction_proto_rawDesc = nil
	file_api_matchfunction_proto_goTypes = nil
	file_api_matchfunction_proto_depIdxs = nil
}
