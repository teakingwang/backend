package medical

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const _ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

type GetMedicalRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Id            int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMedicalRecordRequest) Reset()                             { *x = GetMedicalRecordRequest{} }
func (x *GetMedicalRecordRequest) ProtoReflect() protoreflect.Message { return nil }

type GetMedicalRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetMedicalRecordResponse) Reset()                             { *x = GetMedicalRecordResponse{} }
func (x *GetMedicalRecordResponse) ProtoReflect() protoreflect.Message { return nil }
