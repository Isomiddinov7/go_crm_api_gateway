// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: supper_admin.proto

package users_service

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

type SupperAdminPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *SupperAdminPrimaryKey) Reset() {
	*x = SupperAdminPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supper_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupperAdminPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupperAdminPrimaryKey) ProtoMessage() {}

func (x *SupperAdminPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_supper_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupperAdminPrimaryKey.ProtoReflect.Descriptor instead.
func (*SupperAdminPrimaryKey) Descriptor() ([]byte, []int) {
	return file_supper_admin_proto_rawDescGZIP(), []int{0}
}

func (x *SupperAdminPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SupperAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	FullName  string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password"`
	Login     string `protobuf:"bytes,4,opt,name=login,proto3" json:"login"`
	RoleId    string `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
}

func (x *SupperAdmin) Reset() {
	*x = SupperAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supper_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupperAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupperAdmin) ProtoMessage() {}

func (x *SupperAdmin) ProtoReflect() protoreflect.Message {
	mi := &file_supper_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupperAdmin.ProtoReflect.Descriptor instead.
func (*SupperAdmin) Descriptor() ([]byte, []int) {
	return file_supper_admin_proto_rawDescGZIP(), []int{1}
}

func (x *SupperAdmin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SupperAdmin) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *SupperAdmin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SupperAdmin) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *SupperAdmin) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *SupperAdmin) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SupperAdmin) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateSupperAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
	Login    string `protobuf:"bytes,3,opt,name=login,proto3" json:"login"`
	RoleId   string `protobuf:"bytes,4,opt,name=role_id,json=roleId,proto3" json:"role_id"`
}

func (x *CreateSupperAdmin) Reset() {
	*x = CreateSupperAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supper_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSupperAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupperAdmin) ProtoMessage() {}

func (x *CreateSupperAdmin) ProtoReflect() protoreflect.Message {
	mi := &file_supper_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupperAdmin.ProtoReflect.Descriptor instead.
func (*CreateSupperAdmin) Descriptor() ([]byte, []int) {
	return file_supper_admin_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSupperAdmin) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateSupperAdmin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateSupperAdmin) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CreateSupperAdmin) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

type UpdateSupperAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password"`
	Login    string `protobuf:"bytes,4,opt,name=login,proto3" json:"login"`
	RoleId   string `protobuf:"bytes,5,opt,name=role_id,json=roleId,proto3" json:"role_id"`
}

func (x *UpdateSupperAdmin) Reset() {
	*x = UpdateSupperAdmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supper_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSupperAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSupperAdmin) ProtoMessage() {}

func (x *UpdateSupperAdmin) ProtoReflect() protoreflect.Message {
	mi := &file_supper_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSupperAdmin.ProtoReflect.Descriptor instead.
func (*UpdateSupperAdmin) Descriptor() ([]byte, []int) {
	return file_supper_admin_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSupperAdmin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSupperAdmin) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateSupperAdmin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateSupperAdmin) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UpdateSupperAdmin) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

type GetListSupperAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search"`
}

func (x *GetListSupperAdminRequest) Reset() {
	*x = GetListSupperAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supper_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListSupperAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListSupperAdminRequest) ProtoMessage() {}

func (x *GetListSupperAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supper_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListSupperAdminRequest.ProtoReflect.Descriptor instead.
func (*GetListSupperAdminRequest) Descriptor() ([]byte, []int) {
	return file_supper_admin_proto_rawDescGZIP(), []int{4}
}

func (x *GetListSupperAdminRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListSupperAdminRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListSupperAdminRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListSupperAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       int64          `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	SpperAdmins []*SupperAdmin `protobuf:"bytes,2,rep,name=spper_admins,json=spperAdmins,proto3" json:"spper_admins"`
}

func (x *GetListSupperAdminResponse) Reset() {
	*x = GetListSupperAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supper_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListSupperAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListSupperAdminResponse) ProtoMessage() {}

func (x *GetListSupperAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supper_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListSupperAdminResponse.ProtoReflect.Descriptor instead.
func (*GetListSupperAdminResponse) Descriptor() ([]byte, []int) {
	return file_supper_admin_proto_rawDescGZIP(), []int{5}
}

func (x *GetListSupperAdminResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListSupperAdminResponse) GetSpperAdmins() []*SupperAdmin {
	if x != nil {
		return x.SpperAdmins
	}
	return nil
}

var File_supper_admin_proto protoreflect.FileDescriptor

var file_supper_admin_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x15, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x15, 0x53, 0x75,
	0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xc3, 0x01, 0x0a, 0x0b, 0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7b, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f,
	0x6c, 0x65, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x71, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x73,
	0x70, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x0b, 0x73,
	0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x32, 0xa1, 0x03, 0x0a, 0x12, 0x53,
	0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x48, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75,
	0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1a, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x70,
	0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70,
	0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70,
	0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x75, 0x70, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x18,
	0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_supper_admin_proto_rawDescOnce sync.Once
	file_supper_admin_proto_rawDescData = file_supper_admin_proto_rawDesc
)

func file_supper_admin_proto_rawDescGZIP() []byte {
	file_supper_admin_proto_rawDescOnce.Do(func() {
		file_supper_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_supper_admin_proto_rawDescData)
	})
	return file_supper_admin_proto_rawDescData
}

var file_supper_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_supper_admin_proto_goTypes = []interface{}{
	(*SupperAdminPrimaryKey)(nil),      // 0: users_service.SupperAdminPrimaryKey
	(*SupperAdmin)(nil),                // 1: users_service.SupperAdmin
	(*CreateSupperAdmin)(nil),          // 2: users_service.CreateSupperAdmin
	(*UpdateSupperAdmin)(nil),          // 3: users_service.UpdateSupperAdmin
	(*GetListSupperAdminRequest)(nil),  // 4: users_service.GetListSupperAdminRequest
	(*GetListSupperAdminResponse)(nil), // 5: users_service.GetListSupperAdminResponse
	(*Empty)(nil),                      // 6: users_service.Empty
}
var file_supper_admin_proto_depIdxs = []int32{
	1, // 0: users_service.GetListSupperAdminResponse.spper_admins:type_name -> users_service.SupperAdmin
	2, // 1: users_service.SupperAdminService.Create:input_type -> users_service.CreateSupperAdmin
	0, // 2: users_service.SupperAdminService.GetById:input_type -> users_service.SupperAdminPrimaryKey
	4, // 3: users_service.SupperAdminService.GetList:input_type -> users_service.GetListSupperAdminRequest
	3, // 4: users_service.SupperAdminService.Update:input_type -> users_service.UpdateSupperAdmin
	0, // 5: users_service.SupperAdminService.Delete:input_type -> users_service.SupperAdminPrimaryKey
	1, // 6: users_service.SupperAdminService.Create:output_type -> users_service.SupperAdmin
	1, // 7: users_service.SupperAdminService.GetById:output_type -> users_service.SupperAdmin
	5, // 8: users_service.SupperAdminService.GetList:output_type -> users_service.GetListSupperAdminResponse
	1, // 9: users_service.SupperAdminService.Update:output_type -> users_service.SupperAdmin
	6, // 10: users_service.SupperAdminService.Delete:output_type -> users_service.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_supper_admin_proto_init() }
func file_supper_admin_proto_init() {
	if File_supper_admin_proto != nil {
		return
	}
	file_teacher_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_supper_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupperAdminPrimaryKey); i {
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
		file_supper_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupperAdmin); i {
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
		file_supper_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSupperAdmin); i {
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
		file_supper_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSupperAdmin); i {
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
		file_supper_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListSupperAdminRequest); i {
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
		file_supper_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListSupperAdminResponse); i {
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
			RawDescriptor: file_supper_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_supper_admin_proto_goTypes,
		DependencyIndexes: file_supper_admin_proto_depIdxs,
		MessageInfos:      file_supper_admin_proto_msgTypes,
	}.Build()
	File_supper_admin_proto = out.File
	file_supper_admin_proto_rawDesc = nil
	file_supper_admin_proto_goTypes = nil
	file_supper_admin_proto_depIdxs = nil
}
