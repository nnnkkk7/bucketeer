// Copyright 2025 The Bucketeer Authors.
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
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: proto/eventcounter/mau_summary.proto

package eventcounter

import (
	client "github.com/bucketeer-io/bucketeer/proto/event/client"
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

type MAUSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yearmonth       string          `protobuf:"bytes,1,opt,name=yearmonth,proto3" json:"yearmonth"`
	EnvironmentId   string          `protobuf:"bytes,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id"`
	SourceId        client.SourceId `protobuf:"varint,3,opt,name=source_id,json=sourceId,proto3,enum=bucketeer.event.client.SourceId" json:"source_id"`
	UserCount       int64           `protobuf:"varint,4,opt,name=user_count,json=userCount,proto3" json:"user_count"`
	RequestCount    int64           `protobuf:"varint,5,opt,name=request_count,json=requestCount,proto3" json:"request_count"`
	EvaluationCount int64           `protobuf:"varint,6,opt,name=evaluation_count,json=evaluationCount,proto3" json:"evaluation_count"`
	GoalCount       int64           `protobuf:"varint,7,opt,name=goal_count,json=goalCount,proto3" json:"goal_count"`
	IsAll           bool            `protobuf:"varint,8,opt,name=is_all,json=isAll,proto3" json:"is_all"`
	IsFinished      bool            `protobuf:"varint,9,opt,name=is_finished,json=isFinished,proto3" json:"is_finished"`
	CreatedAt       int64           `protobuf:"varint,10,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt       int64           `protobuf:"varint,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *MAUSummary) Reset() {
	*x = MAUSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_eventcounter_mau_summary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MAUSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MAUSummary) ProtoMessage() {}

func (x *MAUSummary) ProtoReflect() protoreflect.Message {
	mi := &file_proto_eventcounter_mau_summary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MAUSummary.ProtoReflect.Descriptor instead.
func (*MAUSummary) Descriptor() ([]byte, []int) {
	return file_proto_eventcounter_mau_summary_proto_rawDescGZIP(), []int{0}
}

func (x *MAUSummary) GetYearmonth() string {
	if x != nil {
		return x.Yearmonth
	}
	return ""
}

func (x *MAUSummary) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *MAUSummary) GetSourceId() client.SourceId {
	if x != nil {
		return x.SourceId
	}
	return client.SourceId_UNKNOWN
}

func (x *MAUSummary) GetUserCount() int64 {
	if x != nil {
		return x.UserCount
	}
	return 0
}

func (x *MAUSummary) GetRequestCount() int64 {
	if x != nil {
		return x.RequestCount
	}
	return 0
}

func (x *MAUSummary) GetEvaluationCount() int64 {
	if x != nil {
		return x.EvaluationCount
	}
	return 0
}

func (x *MAUSummary) GetGoalCount() int64 {
	if x != nil {
		return x.GoalCount
	}
	return 0
}

func (x *MAUSummary) GetIsAll() bool {
	if x != nil {
		return x.IsAll
	}
	return false
}

func (x *MAUSummary) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

func (x *MAUSummary) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *MAUSummary) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_proto_eventcounter_mau_summary_proto protoreflect.FileDescriptor

var file_proto_eventcounter_mau_summary_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x75, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65,
	0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x1e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x03, 0x0a, 0x0a, 0x4d, 0x41, 0x55, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x79, 0x65, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x79, 0x65, 0x61, 0x72, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x3d, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65,
	0x72, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x67, 0x6f, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x69, 0x73, 0x41, 0x6c, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2d, 0x69, 0x6f,
	0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_eventcounter_mau_summary_proto_rawDescOnce sync.Once
	file_proto_eventcounter_mau_summary_proto_rawDescData = file_proto_eventcounter_mau_summary_proto_rawDesc
)

func file_proto_eventcounter_mau_summary_proto_rawDescGZIP() []byte {
	file_proto_eventcounter_mau_summary_proto_rawDescOnce.Do(func() {
		file_proto_eventcounter_mau_summary_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_eventcounter_mau_summary_proto_rawDescData)
	})
	return file_proto_eventcounter_mau_summary_proto_rawDescData
}

var file_proto_eventcounter_mau_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_eventcounter_mau_summary_proto_goTypes = []interface{}{
	(*MAUSummary)(nil),   // 0: bucketeer.eventcounter.MAUSummary
	(client.SourceId)(0), // 1: bucketeer.event.client.SourceId
}
var file_proto_eventcounter_mau_summary_proto_depIdxs = []int32{
	1, // 0: bucketeer.eventcounter.MAUSummary.source_id:type_name -> bucketeer.event.client.SourceId
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_eventcounter_mau_summary_proto_init() }
func file_proto_eventcounter_mau_summary_proto_init() {
	if File_proto_eventcounter_mau_summary_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_eventcounter_mau_summary_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MAUSummary); i {
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
			RawDescriptor: file_proto_eventcounter_mau_summary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_eventcounter_mau_summary_proto_goTypes,
		DependencyIndexes: file_proto_eventcounter_mau_summary_proto_depIdxs,
		MessageInfos:      file_proto_eventcounter_mau_summary_proto_msgTypes,
	}.Build()
	File_proto_eventcounter_mau_summary_proto = out.File
	file_proto_eventcounter_mau_summary_proto_rawDesc = nil
	file_proto_eventcounter_mau_summary_proto_goTypes = nil
	file_proto_eventcounter_mau_summary_proto_depIdxs = nil
}
