// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: model.proto

package model

import (
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

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty" form:"code" query:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty" form:"msg" query:"msg"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BaseResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type MedalRank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details []*Medal `protobuf:"bytes,1,rep,name=details,proto3" json:"details,omitempty" form:"details" query:"details"`
}

func (x *MedalRank) Reset() {
	*x = MedalRank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedalRank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedalRank) ProtoMessage() {}

func (x *MedalRank) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedalRank.ProtoReflect.Descriptor instead.
func (*MedalRank) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

func (x *MedalRank) GetDetails() []*Medal {
	if x != nil {
		return x.Details
	}
	return nil
}

type Medal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryName string  `protobuf:"bytes,1,opt,name=countryName,proto3" json:"countryName,omitempty" form:"countryName" query:"countryName"`
	Flag        string  `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty" form:"flag" query:"flag"`
	List        *Medals `protobuf:"bytes,3,opt,name=list,proto3" json:"list,omitempty" form:"list" query:"list"`
}

func (x *Medal) Reset() {
	*x = Medal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Medal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Medal) ProtoMessage() {}

func (x *Medal) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Medal.ProtoReflect.Descriptor instead.
func (*Medal) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{2}
}

func (x *Medal) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *Medal) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

func (x *Medal) GetList() *Medals {
	if x != nil {
		return x.List
	}
	return nil
}

type Medals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gold   int64 `protobuf:"varint,1,opt,name=gold,proto3" json:"gold,omitempty" form:"gold" query:"gold"`
	Sliver int64 `protobuf:"varint,2,opt,name=sliver,proto3" json:"sliver,omitempty" form:"sliver" query:"sliver"`
	Bronze int64 `protobuf:"varint,3,opt,name=bronze,proto3" json:"bronze,omitempty" form:"bronze" query:"bronze"`
	Total  int64 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty" form:"total" query:"total"`
}

func (x *Medals) Reset() {
	*x = Medals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Medals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Medals) ProtoMessage() {}

func (x *Medals) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Medals.ProtoReflect.Descriptor instead.
func (*Medals) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{3}
}

func (x *Medals) GetGold() int64 {
	if x != nil {
		return x.Gold
	}
	return 0
}

func (x *Medals) GetSliver() int64 {
	if x != nil {
		return x.Sliver
	}
	return 0
}

func (x *Medals) GetBronze() int64 {
	if x != nil {
		return x.Bronze
	}
	return 0
}

func (x *Medals) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DailyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty" form:"events" query:"events"`
}

func (x *DailyEvent) Reset() {
	*x = DailyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyEvent) ProtoMessage() {}

func (x *DailyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyEvent.ProtoReflect.Descriptor instead.
func (*DailyEvent) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{4}
}

func (x *DailyEvent) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	Time      string     `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty" form:"time" query:"time"`
	Event     string     `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty" form:"event" query:"event"`
	Group     string     `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty" form:"group" query:"group"`
	Countries []*Country `protobuf:"bytes,5,rep,name=countries,proto3" json:"countries,omitempty" form:"countries" query:"countries"`
	Winner    int64      `protobuf:"varint,6,opt,name=winner,proto3" json:"winner,omitempty" form:"winner" query:"winner"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{5}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Event) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *Event) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *Event) GetCountries() []*Country {
	if x != nil {
		return x.Countries
	}
	return nil
}

func (x *Event) GetWinner() int64 {
	if x != nil {
		return x.Winner
	}
	return 0
}

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" form:"name" query:"name"`
	Flag   string `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty" form:"flag" query:"flag"`
	Rating string `protobuf:"bytes,3,opt,name=rating,proto3" json:"rating,omitempty" form:"rating" query:"rating"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{6}
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Country) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

func (x *Country) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

type EventTypeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types []*EventType `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty" form:"types" query:"types"`
}

func (x *EventTypeList) Reset() {
	*x = EventTypeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTypeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTypeList) ProtoMessage() {}

func (x *EventTypeList) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTypeList.ProtoReflect.Descriptor instead.
func (*EventTypeList) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{7}
}

func (x *EventTypeList) GetTypes() []*EventType {
	if x != nil {
		return x.Types
	}
	return nil
}

type EventType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" form:"name" query:"name"`
}

func (x *EventType) Reset() {
	*x = EventType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventType) ProtoMessage() {}

func (x *EventType) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventType.ProtoReflect.Descriptor instead.
func (*EventType) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{8}
}

func (x *EventType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EventTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tables []*Table `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty" form:"tables" query:"tables"`
}

func (x *EventTable) Reset() {
	*x = EventTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventTable) ProtoMessage() {}

func (x *EventTable) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventTable.ProtoReflect.Descriptor instead.
func (*EventTable) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{9}
}

func (x *EventTable) GetTables() []*Table {
	if x != nil {
		return x.Tables
	}
	return nil
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" form:"title" query:"title"`
	Period    string     `protobuf:"bytes,2,opt,name=period,proto3" json:"period,omitempty" form:"period" query:"period"`
	Countries []*Country `protobuf:"bytes,3,rep,name=countries,proto3" json:"countries,omitempty" form:"countries" query:"countries"`
	Winner    int64      `protobuf:"varint,4,opt,name=winner,proto3" json:"winner,omitempty" form:"winner" query:"winner"`
	Special   string     `protobuf:"bytes,5,opt,name=special,proto3" json:"special,omitempty" form:"special" query:"special"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{10}
}

func (x *Table) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Table) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Table) GetCountries() []*Country {
	if x != nil {
		return x.Countries
	}
	return nil
}

func (x *Table) GetWinner() int64 {
	if x != nil {
		return x.Winner
	}
	return 0
}

func (x *Table) GetSpecial() string {
	if x != nil {
		return x.Special
	}
	return ""
}

type EventList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId  string     `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty" form:"group_id" query:"group_id"`
	Contests []*Contest `protobuf:"bytes,2,rep,name=contests,proto3" json:"contests,omitempty" form:"contests" query:"contests"`
}

func (x *EventList) Reset() {
	*x = EventList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventList) ProtoMessage() {}

func (x *EventList) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventList.ProtoReflect.Descriptor instead.
func (*EventList) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{11}
}

func (x *EventList) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *EventList) GetContests() []*Contest {
	if x != nil {
		return x.Contests
	}
	return nil
}

type Contest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId string     `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty" form:"contest_id" query:"contest_id"`
	Date      string     `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty" form:"date" query:"date"`
	Countries []*Country `protobuf:"bytes,3,rep,name=countries,proto3" json:"countries,omitempty" form:"countries" query:"countries"`
	Status    int64      `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty" form:"status" query:"status"`
}

func (x *Contest) Reset() {
	*x = Contest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contest) ProtoMessage() {}

func (x *Contest) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contest.ProtoReflect.Descriptor instead.
func (*Contest) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{12}
}

func (x *Contest) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *Contest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Contest) GetCountries() []*Country {
	if x != nil {
		return x.Countries
	}
	return nil
}

func (x *Contest) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_model_proto protoreflect.FileDescriptor

var file_model_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x30, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x33, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x52,
	0x61, 0x6e, 0x6b, 0x12, 0x26, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x65, 0x64,
	0x61, 0x6c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x60, 0x0a, 0x05, 0x6d,
	0x65, 0x64, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x21, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x73, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x62, 0x0a,
	0x06, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x6f, 0x6e, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x72, 0x6f, 0x6e, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x32, 0x0a, 0x0a, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x24, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x2c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x49, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x37, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x0a, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x95,
	0x01, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x2c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x22, 0x52, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0x0e, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42,
	0x19, 0x5a, 0x17, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_model_proto_rawDescOnce sync.Once
	file_model_proto_rawDescData = file_model_proto_rawDesc
)

func file_model_proto_rawDescGZIP() []byte {
	file_model_proto_rawDescOnce.Do(func() {
		file_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_proto_rawDescData)
	})
	return file_model_proto_rawDescData
}

var file_model_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_model_proto_goTypes = []interface{}{
	(*BaseResp)(nil),      // 0: model.baseResp
	(*MedalRank)(nil),     // 1: model.medalRank
	(*Medal)(nil),         // 2: model.medal
	(*Medals)(nil),        // 3: model.medals
	(*DailyEvent)(nil),    // 4: model.dailyEvent
	(*Event)(nil),         // 5: model.event
	(*Country)(nil),       // 6: model.country
	(*EventTypeList)(nil), // 7: model.eventTypeList
	(*EventType)(nil),     // 8: model.eventType
	(*EventTable)(nil),    // 9: model.eventTable
	(*Table)(nil),         // 10: model.table
	(*EventList)(nil),     // 11: model.eventList
	(*Contest)(nil),       // 12: model.contest
}
var file_model_proto_depIdxs = []int32{
	2,  // 0: model.medalRank.details:type_name -> model.medal
	3,  // 1: model.medal.list:type_name -> model.medals
	5,  // 2: model.dailyEvent.events:type_name -> model.event
	6,  // 3: model.event.countries:type_name -> model.country
	8,  // 4: model.eventTypeList.types:type_name -> model.eventType
	10, // 5: model.eventTable.tables:type_name -> model.table
	6,  // 6: model.table.countries:type_name -> model.country
	12, // 7: model.eventList.contests:type_name -> model.contest
	6,  // 8: model.contest.countries:type_name -> model.country
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_model_proto_init() }
func file_model_proto_init() {
	if File_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedalRank); i {
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
		file_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Medal); i {
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
		file_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Medals); i {
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
		file_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyEvent); i {
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
		file_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
		file_model_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTypeList); i {
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
		file_model_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventType); i {
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
		file_model_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventTable); i {
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
		file_model_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_model_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventList); i {
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
		file_model_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contest); i {
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
			RawDescriptor: file_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_model_proto_goTypes,
		DependencyIndexes: file_model_proto_depIdxs,
		MessageInfos:      file_model_proto_msgTypes,
	}.Build()
	File_model_proto = out.File
	file_model_proto_rawDesc = nil
	file_model_proto_goTypes = nil
	file_model_proto_depIdxs = nil
}
