// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: idl/task.proto

package task

import (
	_ "backend/biz/model/api"
	model "backend/biz/model/model"
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AllMedalsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MedalSorts string `protobuf:"bytes,1,opt,name=medalSorts,proto3" json:"medalSorts,omitempty" form:"medalSorts" query:"medalSorts"`
}

func (x *AllMedalsReq) Reset() {
	*x = AllMedalsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllMedalsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllMedalsReq) ProtoMessage() {}

func (x *AllMedalsReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllMedalsReq.ProtoReflect.Descriptor instead.
func (*AllMedalsReq) Descriptor() ([]byte, []int) {
	return file_idl_task_proto_rawDescGZIP(), []int{0}
}

func (x *AllMedalsReq) GetMedalSorts() string {
	if x != nil {
		return x.MedalSorts
	}
	return ""
}

type AllMedalsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *model.BaseResp  `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty" form:"base" query:"base"`
	Data *model.MedalRank `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" form:"data" query:"data"`
}

func (x *AllMedalsResp) Reset() {
	*x = AllMedalsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllMedalsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllMedalsResp) ProtoMessage() {}

func (x *AllMedalsResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllMedalsResp.ProtoReflect.Descriptor instead.
func (*AllMedalsResp) Descriptor() ([]byte, []int) {
	return file_idl_task_proto_rawDescGZIP(), []int{1}
}

func (x *AllMedalsResp) GetBase() *model.BaseResp {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *AllMedalsResp) GetData() *model.MedalRank {
	if x != nil {
		return x.Data
	}
	return nil
}

type DailyEventReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty" form:"date" query:"date"`
}

func (x *DailyEventReq) Reset() {
	*x = DailyEventReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyEventReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyEventReq) ProtoMessage() {}

func (x *DailyEventReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyEventReq.ProtoReflect.Descriptor instead.
func (*DailyEventReq) Descriptor() ([]byte, []int) {
	return file_idl_task_proto_rawDescGZIP(), []int{2}
}

func (x *DailyEventReq) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type DailyEventResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *model.BaseResp   `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty" form:"base" query:"base"`
	Data *model.DailyEvent `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" form:"data" query:"data"`
}

func (x *DailyEventResp) Reset() {
	*x = DailyEventResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyEventResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyEventResp) ProtoMessage() {}

func (x *DailyEventResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyEventResp.ProtoReflect.Descriptor instead.
func (*DailyEventResp) Descriptor() ([]byte, []int) {
	return file_idl_task_proto_rawDescGZIP(), []int{3}
}

func (x *DailyEventResp) GetBase() *model.BaseResp {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *DailyEventResp) GetData() *model.DailyEvent {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventTypeListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EventTypeListReq) Reset() {
	*x = EventTypeListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypeListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypeListReq) ProtoMessage() {}

func (x *EventTypeListReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypeListReq.ProtoReflect.Descriptor instead.
func (*EventTypeListReq) Descriptor() ([]byte, []int) {
	return file_idl_task_proto_rawDescGZIP(), []int{4}
}

type EventTypeListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *model.BaseResp       `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty" form:"base" query:"base"`
	Data *model.EventTypeLists `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" form:"data" query:"data"`
}

func (x *EventTypeListResp) Reset() {
	*x = EventTypeListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypeListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypeListResp) ProtoMessage() {}

func (x *EventTypeListResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypeListResp.ProtoReflect.Descriptor instead.
func (*EventTypeListResp) Descriptor() ([]byte, []int) {
	return file_idl_task_proto_rawDescGZIP(), []int{5}
}

func (x *EventTypeListResp) GetBase() *model.BaseResp {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *EventTypeListResp) GetData() *model.EventTypeLists {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventTableReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventID string `protobuf:"bytes,1,opt,name=eventID,proto3" json:"eventID,omitempty" form:"eventID" query:"eventID"`
}

func (x *EventTableReq) Reset() {
	*x = EventTableReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTableReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTableReq) ProtoMessage() {}

func (x *EventTableReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTableReq.ProtoReflect.Descriptor instead.
func (*EventTableReq) Descriptor() ([]byte, []int) {
	return file_idl_task_proto_rawDescGZIP(), []int{6}
}

func (x *EventTableReq) GetEventID() string {
	if x != nil {
		return x.EventID
	}
	return ""
}

type EventTableResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *model.BaseResp   `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty" form:"base" query:"base"`
	Data *model.EventTable `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" form:"data" query:"data"`
}

func (x *EventTableResp) Reset() {
	*x = EventTableResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTableResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTableResp) ProtoMessage() {}

func (x *EventTableResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTableResp.ProtoReflect.Descriptor instead.
func (*EventTableResp) Descriptor() ([]byte, []int) {
	return file_idl_task_proto_rawDescGZIP(), []int{7}
}

func (x *EventTableResp) GetBase() *model.BaseResp {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *EventTableResp) GetData() *model.EventTable {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventID string `protobuf:"bytes,1,opt,name=eventID,proto3" json:"eventID,omitempty" form:"eventID" query:"eventID"`
}

func (x *EventListReq) Reset() {
	*x = EventListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventListReq) ProtoMessage() {}

func (x *EventListReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventListReq.ProtoReflect.Descriptor instead.
func (*EventListReq) Descriptor() ([]byte, []int) {
	return file_idl_task_proto_rawDescGZIP(), []int{8}
}

func (x *EventListReq) GetEventID() string {
	if x != nil {
		return x.EventID
	}
	return ""
}

type EventListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *model.BaseResp   `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty" form:"base" query:"base"`
	Data *model.EventLists `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" form:"data" query:"data"`
}

func (x *EventListResp) Reset() {
	*x = EventListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_task_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventListResp) ProtoMessage() {}

func (x *EventListResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_task_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventListResp.ProtoReflect.Descriptor instead.
func (*EventListResp) Descriptor() ([]byte, []int) {
	return file_idl_task_proto_rawDescGZIP(), []int{9}
}

func (x *EventListResp) GetBase() *model.BaseResp {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *EventListResp) GetData() *model.EventLists {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_idl_task_proto protoreflect.FileDescriptor

var file_idl_task_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e,
	0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x4d, 0x65, 0x64, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x53, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x53, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x5a,
	0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x4d, 0x65, 0x64, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x23, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x65, 0x64, 0x61, 0x6c,
	0x52, 0x61, 0x6e, 0x6b, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x0d, 0x64, 0x61,
	0x69, 0x6c, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x5c, 0x0a, 0x0e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x64, 0x61, 0x69,
	0x6c, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x12, 0x0a,
	0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x22, 0x63, 0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x73,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x22, 0x5c, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x28, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x5b, 0x0a, 0x0d, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x73,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x96, 0x03, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x64,
	0x61, 0x6c, 0x73, 0x12, 0x12, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x6c, 0x6c, 0x4d, 0x65,
	0x64, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61,
	0x6c, 0x6c, 0x4d, 0x65, 0x64, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x22, 0x12, 0xca, 0xc1, 0x18,
	0x0e, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x73, 0x2f, 0x61, 0x6c, 0x6c, 0x12,
	0x4c, 0x0a, 0x0a, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x13, 0xca, 0xc1, 0x18, 0x0f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x12, 0x58, 0x0a,
	0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x16, 0xca, 0xc1, 0x18, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x13, 0xca, 0xc1, 0x18, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x12, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x12, 0xca, 0xc1, 0x18,
	0x0e, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42,
	0x18, 0x5a, 0x16, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_idl_task_proto_rawDescOnce sync.Once
	file_idl_task_proto_rawDescData = file_idl_task_proto_rawDesc
)

func file_idl_task_proto_rawDescGZIP() []byte {
	file_idl_task_proto_rawDescOnce.Do(func() {
		file_idl_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_task_proto_rawDescData)
	})
	return file_idl_task_proto_rawDescData
}

var file_idl_task_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_idl_task_proto_goTypes = []interface{}{
	(*AllMedalsReq)(nil),         // 0: task.allMedalsReq
	(*AllMedalsResp)(nil),        // 1: task.allMedalsResp
	(*DailyEventReq)(nil),        // 2: task.dailyEventReq
	(*DailyEventResp)(nil),       // 3: task.dailyEventResp
	(*EventTypeListReq)(nil),     // 4: task.eventTypeListReq
	(*EventTypeListResp)(nil),    // 5: task.eventTypeListResp
	(*EventTableReq)(nil),        // 6: task.eventTableReq
	(*EventTableResp)(nil),       // 7: task.eventTableResp
	(*EventListReq)(nil),         // 8: task.eventListReq
	(*EventListResp)(nil),        // 9: task.eventListResp
	(*model.BaseResp)(nil),       // 10: model.baseResp
	(*model.MedalRank)(nil),      // 11: model.medalRank
	(*model.DailyEvent)(nil),     // 12: model.dailyEvent
	(*model.EventTypeLists)(nil), // 13: model.eventTypeLists
	(*model.EventTable)(nil),     // 14: model.eventTable
	(*model.EventLists)(nil),     // 15: model.eventLists
}
var file_idl_task_proto_depIdxs = []int32{
	10, // 0: task.allMedalsResp.base:type_name -> model.baseResp
	11, // 1: task.allMedalsResp.data:type_name -> model.medalRank
	10, // 2: task.dailyEventResp.base:type_name -> model.baseResp
	12, // 3: task.dailyEventResp.data:type_name -> model.dailyEvent
	10, // 4: task.eventTypeListResp.base:type_name -> model.baseResp
	13, // 5: task.eventTypeListResp.data:type_name -> model.eventTypeLists
	10, // 6: task.eventTableResp.base:type_name -> model.baseResp
	14, // 7: task.eventTableResp.data:type_name -> model.eventTable
	10, // 8: task.eventListResp.base:type_name -> model.baseResp
	15, // 9: task.eventListResp.data:type_name -> model.eventLists
	0,  // 10: task.TaskService.AllMedals:input_type -> task.allMedalsReq
	2,  // 11: task.TaskService.DailyEvent:input_type -> task.dailyEventReq
	4,  // 12: task.TaskService.EventTypeList:input_type -> task.eventTypeListReq
	6,  // 13: task.TaskService.EventTable:input_type -> task.eventTableReq
	8,  // 14: task.TaskService.EventList:input_type -> task.eventListReq
	0,  // 15: task.TaskService.AllMedals:output_type -> task.allMedalsReq
	3,  // 16: task.TaskService.DailyEvent:output_type -> task.dailyEventResp
	5,  // 17: task.TaskService.EventTypeList:output_type -> task.eventTypeListResp
	7,  // 18: task.TaskService.EventTable:output_type -> task.eventTableResp
	9,  // 19: task.TaskService.EventList:output_type -> task.eventListResp
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_idl_task_proto_init() }
func file_idl_task_proto_init() {
	if File_idl_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllMedalsReq); i {
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
		file_idl_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllMedalsResp); i {
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
		file_idl_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyEventReq); i {
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
		file_idl_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyEventResp); i {
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
		file_idl_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypeListReq); i {
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
		file_idl_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypeListResp); i {
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
		file_idl_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTableReq); i {
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
		file_idl_task_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTableResp); i {
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
		file_idl_task_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventListReq); i {
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
		file_idl_task_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventListResp); i {
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
			RawDescriptor: file_idl_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_task_proto_goTypes,
		DependencyIndexes: file_idl_task_proto_depIdxs,
		MessageInfos:      file_idl_task_proto_msgTypes,
	}.Build()
	File_idl_task_proto = out.File
	file_idl_task_proto_rawDesc = nil
	file_idl_task_proto_goTypes = nil
	file_idl_task_proto_depIdxs = nil
}
