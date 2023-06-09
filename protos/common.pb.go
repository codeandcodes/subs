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

type SubscriptionFrequency_Cadence int32

const (
	SubscriptionFrequency_DAILY             SubscriptionFrequency_Cadence = 0
	SubscriptionFrequency_WEEKLY            SubscriptionFrequency_Cadence = 1
	SubscriptionFrequency_EVERY_TWO_WEEKS   SubscriptionFrequency_Cadence = 2
	SubscriptionFrequency_THIRTY_DAYS       SubscriptionFrequency_Cadence = 3
	SubscriptionFrequency_SIXTY_DAYS        SubscriptionFrequency_Cadence = 4
	SubscriptionFrequency_NINETY_DAYS       SubscriptionFrequency_Cadence = 5
	SubscriptionFrequency_MONTHLY           SubscriptionFrequency_Cadence = 6
	SubscriptionFrequency_EVERY_TWO_MONTHS  SubscriptionFrequency_Cadence = 7
	SubscriptionFrequency_QUARTERLY         SubscriptionFrequency_Cadence = 8
	SubscriptionFrequency_EVERY_FOUR_MONTHS SubscriptionFrequency_Cadence = 9
	SubscriptionFrequency_EVERY_SIX_MONTHS  SubscriptionFrequency_Cadence = 10
	SubscriptionFrequency_ANNUAL            SubscriptionFrequency_Cadence = 11
	SubscriptionFrequency_EVERY_TWO_YEARS   SubscriptionFrequency_Cadence = 12
)

// Enum value maps for SubscriptionFrequency_Cadence.
var (
	SubscriptionFrequency_Cadence_name = map[int32]string{
		0:  "DAILY",
		1:  "WEEKLY",
		2:  "EVERY_TWO_WEEKS",
		3:  "THIRTY_DAYS",
		4:  "SIXTY_DAYS",
		5:  "NINETY_DAYS",
		6:  "MONTHLY",
		7:  "EVERY_TWO_MONTHS",
		8:  "QUARTERLY",
		9:  "EVERY_FOUR_MONTHS",
		10: "EVERY_SIX_MONTHS",
		11: "ANNUAL",
		12: "EVERY_TWO_YEARS",
	}
	SubscriptionFrequency_Cadence_value = map[string]int32{
		"DAILY":             0,
		"WEEKLY":            1,
		"EVERY_TWO_WEEKS":   2,
		"THIRTY_DAYS":       3,
		"SIXTY_DAYS":        4,
		"NINETY_DAYS":       5,
		"MONTHLY":           6,
		"EVERY_TWO_MONTHS":  7,
		"QUARTERLY":         8,
		"EVERY_FOUR_MONTHS": 9,
		"EVERY_SIX_MONTHS":  10,
		"ANNUAL":            11,
		"EVERY_TWO_YEARS":   12,
	}
)

func (x SubscriptionFrequency_Cadence) Enum() *SubscriptionFrequency_Cadence {
	p := new(SubscriptionFrequency_Cadence)
	*p = x
	return p
}

func (x SubscriptionFrequency_Cadence) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubscriptionFrequency_Cadence) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (SubscriptionFrequency_Cadence) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x SubscriptionFrequency_Cadence) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubscriptionFrequency_Cadence.Descriptor instead.
func (SubscriptionFrequency_Cadence) EnumDescriptor() ([]byte, []int) {
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

	// os user id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// non-optional
	EmailAddress string `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	// first name
	GivenName string `protobuf:"bytes,3,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	// last name
	FamilyName string `protobuf:"bytes,4,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	// populated from Square Customer API
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// populated from Square Customer API
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Unique square-assigned ID (not the same as idempotency key)
	SquareId string `protobuf:"bytes,7,opt,name=square_id,json=squareId,proto3" json:"square_id,omitempty"`
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

func (x *User) GetGivenName() string {
	if x != nil {
		return x.GivenName
	}
	return ""
}

func (x *User) GetFamilyName() string {
	if x != nil {
		return x.FamilyName
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *User) GetSquareId() string {
	if x != nil {
		return x.SquareId
	}
	return ""
}

// Maps onto the SubscriptionPhase from Square APIs
type SubscriptionFrequency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cadence SubscriptionFrequency_Cadence `protobuf:"varint,1,opt,name=cadence,proto3,enum=subs.SubscriptionFrequency_Cadence" json:"cadence,omitempty"`
	// The date of the first occurrence `YYYY-MM-DD`-formatted date
	StartDate string `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// This is the number of periods the subscription is active for. It will start on start date.
	Periods int32 `protobuf:"varint,3,opt,name=periods,proto3" json:"periods,omitempty"`
	// If is_ongoing is set, then this is a continuous subscription
	// In the square API, this means that period should be removed
	IsOngoing bool `protobuf:"varint,4,opt,name=is_ongoing,json=isOngoing,proto3" json:"is_ongoing,omitempty"`
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

func (x *SubscriptionFrequency) GetCadence() SubscriptionFrequency_Cadence {
	if x != nil {
		return x.Cadence
	}
	return SubscriptionFrequency_DAILY
}

func (x *SubscriptionFrequency) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *SubscriptionFrequency) GetPeriods() int32 {
	if x != nil {
		return x.Periods
	}
	return 0
}

func (x *SubscriptionFrequency) GetIsOngoing() bool {
	if x != nil {
		return x.IsOngoing
	}
	return false
}

// Maps to a square location
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationId  string `protobuf:"bytes,1,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	CountryCode string `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Location) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Location) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Location) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x73, 0x75, 0x62, 0x73, 0x22, 0xd6, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x49, 0x64, 0x22, 0x98, 0x03,
	0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x61, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x63,
	0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x22, 0xe7,
	0x01, 0x0a, 0x07, 0x43, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x41,
	0x49, 0x4c, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x45, 0x45, 0x4b, 0x4c, 0x59, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x57, 0x4f, 0x5f, 0x57,
	0x45, 0x45, 0x4b, 0x53, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x48, 0x49, 0x52, 0x54, 0x59,
	0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x49, 0x58, 0x54, 0x59,
	0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x49, 0x4e, 0x45, 0x54,
	0x59, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x54,
	0x48, 0x4c, 0x59, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x54,
	0x57, 0x4f, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x53, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x51,
	0x55, 0x41, 0x52, 0x54, 0x45, 0x52, 0x4c, 0x59, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56,
	0x45, 0x52, 0x59, 0x5f, 0x46, 0x4f, 0x55, 0x52, 0x5f, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x53, 0x10,
	0x09, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x49, 0x58, 0x5f, 0x4d,
	0x4f, 0x4e, 0x54, 0x48, 0x53, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4e, 0x4e, 0x55, 0x41,
	0x4c, 0x10, 0x0b, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x57, 0x4f,
	0x5f, 0x59, 0x45, 0x41, 0x52, 0x53, 0x10, 0x0c, 0x22, 0x7c, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x2a, 0x67, 0x0a, 0x09, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57,
	0x65, 0x65, 0x6b, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54,
	0x55, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x45, 0x44, 0x4e,
	0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x48, 0x55, 0x52, 0x53,
	0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x49, 0x44, 0x41, 0x59, 0x10,
	0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x41, 0x54, 0x55, 0x52, 0x44, 0x41, 0x59, 0x10, 0x06, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x61, 0x6e, 0x64, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_proto_goTypes = []interface{}{
	(DayOfWeek)(0),                     // 0: subs.DayOfWeek
	(SubscriptionFrequency_Cadence)(0), // 1: subs.SubscriptionFrequency.Cadence
	(*User)(nil),                       // 2: subs.User
	(*SubscriptionFrequency)(nil),      // 3: subs.SubscriptionFrequency
	(*Location)(nil),                   // 4: subs.Location
}
var file_common_proto_depIdxs = []int32{
	1, // 0: subs.SubscriptionFrequency.cadence:type_name -> subs.SubscriptionFrequency.Cadence
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
			NumMessages:   3,
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
