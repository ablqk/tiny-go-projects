// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: habits.proto

package api

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

// CreateHabitRequest is the message sent to create a habit.
type CreateHabitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Habit to track.
	Habit *Habit `protobuf:"bytes,1,opt,name=habit,proto3" json:"habit,omitempty"`
}

func (x *CreateHabitRequest) Reset() {
	*x = CreateHabitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habits_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHabitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHabitRequest) ProtoMessage() {}

func (x *CreateHabitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_habits_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHabitRequest.ProtoReflect.Descriptor instead.
func (*CreateHabitRequest) Descriptor() ([]byte, []int) {
	return file_habits_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHabitRequest) GetHabit() *Habit {
	if x != nil {
		return x.Habit
	}
	return nil
}

// Habit represents an objective one wants to complete a given amount of times per week.
type Habit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the habit.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Frequency, expressed in times per week. Defaults to 1.
	Frequency *int32 `protobuf:"varint,2,opt,name=frequency,proto3,oneof" json:"frequency,omitempty"`
}

func (x *Habit) Reset() {
	*x = Habit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_habits_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Habit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Habit) ProtoMessage() {}

func (x *Habit) ProtoReflect() protoreflect.Message {
	mi := &file_habits_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Habit.ProtoReflect.Descriptor instead.
func (*Habit) Descriptor() ([]byte, []int) {
	return file_habits_proto_rawDescGZIP(), []int{1}
}

func (x *Habit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Habit) GetFrequency() int32 {
	if x != nil && x.Frequency != nil {
		return *x.Frequency
	}
	return 0
}

var File_habits_proto protoreflect.FileDescriptor

var file_habits_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x68, 0x61, 0x62, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x68, 0x61, 0x62, 0x69, 0x74, 0x73, 0x22, 0x39, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05,
	0x68, 0x61, 0x62, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x68, 0x61,
	0x62, 0x69, 0x74, 0x73, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x05, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x22, 0x4c, 0x0a, 0x05, 0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x88, 0x01,
	0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x32,
	0x42, 0x0a, 0x06, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x1a, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x73, 0x2e, 0x48, 0x61,
	0x62, 0x69, 0x74, 0x42, 0x1c, 0x5a, 0x1a, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x67, 0x6f, 0x2d, 0x70,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x68, 0x61, 0x62, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_habits_proto_rawDescOnce sync.Once
	file_habits_proto_rawDescData = file_habits_proto_rawDesc
)

func file_habits_proto_rawDescGZIP() []byte {
	file_habits_proto_rawDescOnce.Do(func() {
		file_habits_proto_rawDescData = protoimpl.X.CompressGZIP(file_habits_proto_rawDescData)
	})
	return file_habits_proto_rawDescData
}

var file_habits_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_habits_proto_goTypes = []interface{}{
	(*CreateHabitRequest)(nil), // 0: habits.CreateHabitRequest
	(*Habit)(nil),              // 1: habits.Habit
}
var file_habits_proto_depIdxs = []int32{
	1, // 0: habits.CreateHabitRequest.habit:type_name -> habits.Habit
	0, // 1: habits.Habits.CreateHabit:input_type -> habits.CreateHabitRequest
	1, // 2: habits.Habits.CreateHabit:output_type -> habits.Habit
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_habits_proto_init() }
func file_habits_proto_init() {
	if File_habits_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_habits_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHabitRequest); i {
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
		file_habits_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Habit); i {
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
	file_habits_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_habits_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_habits_proto_goTypes,
		DependencyIndexes: file_habits_proto_depIdxs,
		MessageInfos:      file_habits_proto_msgTypes,
	}.Build()
	File_habits_proto = out.File
	file_habits_proto_rawDesc = nil
	file_habits_proto_goTypes = nil
	file_habits_proto_depIdxs = nil
}
