// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.3
// source: common.proto

package protos

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

type DayOfWeek int32

const (
	DayOfWeek_SUNDAY    DayOfWeek = 0
	DayOfWeek_MONDAY    DayOfWeek = 1
	DayOfWeek_TUESDAY   DayOfWeek = 2
	DayOfWeek_WEDNESDAY DayOfWeek = 3
	DayOfWeek_THURSDAY  DayOfWeek = 4
	DayOfWeek_FRIDAY    DayOfWeek = 5
	DayOfWeek_SATURDAY  DayOfWeek = 6
)

// Enum value maps for DayOfWeek.
var (
	DayOfWeek_name = map[int32]string{
		0: "SUNDAY",
		1: "MONDAY",
		2: "TUESDAY",
		3: "WEDNESDAY",
		4: "THURSDAY",
		5: "FRIDAY",
		6: "SATURDAY",
	}
	DayOfWeek_value = map[string]int32{
		"SUNDAY":    0,
		"MONDAY":    1,
		"TUESDAY":   2,
		"WEDNESDAY": 3,
		"THURSDAY":  4,
		"FRIDAY":    5,
		"SATURDAY":  6,
	}
)

func (x DayOfWeek) Enum() *DayOfWeek {
	p := new(DayOfWeek)
	*p = x
	return p
}

func (x DayOfWeek) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DayOfWeek) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (DayOfWeek) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x DayOfWeek) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DayOfWeek.Descriptor instead.
func (DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type SubscriptionFrequency_FrequencyType int32

const (
	SubscriptionFrequency_DAILY   SubscriptionFrequency_FrequencyType = 0
	SubscriptionFrequency_WEEKLY  SubscriptionFrequency_FrequencyType = 1
	SubscriptionFrequency_MONTHLY SubscriptionFrequency_FrequencyType = 2
	SubscriptionFrequency_YEARLY  SubscriptionFrequency_FrequencyType = 3
)

// Enum value maps for SubscriptionFrequency_FrequencyType.
var (
	SubscriptionFrequency_FrequencyType_name = map[int32]string{
		0: "DAILY",
		1: "WEEKLY",
		2: "MONTHLY",
		3: "YEARLY",
	}
	SubscriptionFrequency_FrequencyType_value = map[string]int32{
		"DAILY":   0,
		"WEEKLY":  1,
		"MONTHLY": 2,
		"YEARLY":  3,
	}
)

func (x SubscriptionFrequency_FrequencyType) Enum() *SubscriptionFrequency_FrequencyType {
	p := new(SubscriptionFrequency_FrequencyType)
	*p = x
	return p
}

func (x SubscriptionFrequency_FrequencyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubscriptionFrequency_FrequencyType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (SubscriptionFrequency_FrequencyType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x SubscriptionFrequency_FrequencyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubscriptionFrequency_FrequencyType.Descriptor instead.
func (SubscriptionFrequency_FrequencyType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1, 0}
}

// User
// User can either be pulled from the database
// Or if only an email address is available, uses that instead
// Preference should be:
// 1) ID (stored user in DB, has an account with onlysubs)
// 2) email address
// 3) Phone number or other contact methods (e.g. in the future)
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EmailAddress string `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

// Adapted from Google schedule
type SubscriptionFrequency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type SubscriptionFrequency_FrequencyType `protobuf:"varint,1,opt,name=type,proto3,enum=subs.SubscriptionFrequency_FrequencyType" json:"type,omitempty"`
	// The interval at which the event recurs.
	// For example, if type is WEEKLY and interval is 2, the event occurs every 2 weeks.
	Interval int32 `protobuf:"varint,2,opt,name=interval,proto3" json:"interval,omitempty"`
	// The date of the first occurrence
	StartDate string `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// An optional end date for the recurrence
	EndDate string `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// An optional maximum number of occurrences.
	// If this field is set, the event will stop recurring after this many occurrences.
	MaxOccurrences int32 `protobuf:"varint,5,opt,name=max_occurrences,json=maxOccurrences,proto3" json:"max_occurrences,omitempty"`
	// If the type is MONTHLY or YEARLY,
	// this could represent the day of the month that the event should occur
	DayOfMonth int32 `protobuf:"varint,6,opt,name=day_of_month,json=dayOfMonth,proto3" json:"day_of_month,omitempty"`
	// If the type is WEEKLY,
	// this could represent the days of the week that the event should occur.
	DaysOfWeek []DayOfWeek `protobuf:"varint,7,rep,packed,name=days_of_week,json=daysOfWeek,proto3,enum=subs.DayOfWeek" json:"days_of_week,omitempty"`
	// If true, the event recurs indefinitely.
	IsOngoing bool `protobuf:"varint,8,opt,name=is_ongoing,json=isOngoing,proto3" json:"is_ongoing,omitempty"`
}

func (x *SubscriptionFrequency) Reset() {
	*x = SubscriptionFrequency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionFrequency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionFrequency) ProtoMessage() {}

func (x *SubscriptionFrequency) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionFrequency.ProtoReflect.Descriptor instead.
func (*SubscriptionFrequency) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriptionFrequency) GetType() SubscriptionFrequency_FrequencyType {
	if x != nil {
		return x.Type
	}
	return SubscriptionFrequency_DAILY
}

func (x *SubscriptionFrequency) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *SubscriptionFrequency) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *SubscriptionFrequency) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *SubscriptionFrequency) GetMaxOccurrences() int32 {
	if x != nil {
		return x.MaxOccurrences
	}
	return 0
}

func (x *SubscriptionFrequency) GetDayOfMonth() int32 {
	if x != nil {
		return x.DayOfMonth
	}
	return 0
}

func (x *SubscriptionFrequency) GetDaysOfWeek() []DayOfWeek {
	if x != nil {
		return x.DaysOfWeek
	}
	return nil
}

func (x *SubscriptionFrequency) GetIsOngoing() bool {
	if x != nil {
		return x.IsOngoing
	}
	return false
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x73, 0x75, 0x62, 0x73, 0x22, 0x3b, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x8a, 0x03, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3d, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x4f, 0x63,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x61, 0x79,
	0x5f, 0x6f, 0x66, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x64, 0x61, 0x79, 0x4f, 0x66, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x31, 0x0a, 0x0c, 0x64,
	0x61, 0x79, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x2e, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65,
	0x65, 0x6b, 0x52, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x22, 0x3f, 0x0a,
	0x0d, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x45, 0x45,
	0x4b, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x45, 0x41, 0x52, 0x4c, 0x59, 0x10, 0x03, 0x2a, 0x67,
	0x0a, 0x09, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x0a, 0x0a, 0x06, 0x53,
	0x55, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x44, 0x41,
	0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x55, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x02,
	0x12, 0x0d, 0x0a, 0x09, 0x57, 0x45, 0x44, 0x4e, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12,
	0x0c, 0x0a, 0x08, 0x54, 0x48, 0x55, 0x52, 0x53, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x52, 0x49, 0x44, 0x41, 0x59, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x41, 0x54,
	0x55, 0x52, 0x44, 0x41, 0x59, 0x10, 0x06, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x61, 0x6e, 0x64, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_proto_goTypes = []interface{}{
	(DayOfWeek)(0),                           // 0: subs.DayOfWeek
	(SubscriptionFrequency_FrequencyType)(0), // 1: subs.SubscriptionFrequency.FrequencyType
	(*User)(nil),                             // 2: subs.User
	(*SubscriptionFrequency)(nil),            // 3: subs.SubscriptionFrequency
}
var file_common_proto_depIdxs = []int32{
	1, // 0: subs.SubscriptionFrequency.type:type_name -> subs.SubscriptionFrequency.FrequencyType
	0, // 1: subs.SubscriptionFrequency.days_of_week:type_name -> subs.DayOfWeek
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionFrequency); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
