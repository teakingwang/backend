package medical

import (
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetMedicalRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Id            int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMedicalRecordRequest) Reset()         { *x = GetMedicalRecordRequest{} }
func (x *GetMedicalRecordRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetMedicalRecordRequest) ProtoMessage()    {}
func (x *GetMedicalRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_medical_medical_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*GetMedicalRecordRequest) Descriptor() ([]byte, []int) {
	return file_medical_medical_proto_rawDesc, []int{0}
}
func (x *GetMedicalRecordRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMedicalRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetMedicalRecordResponse) Reset()         { *x = GetMedicalRecordResponse{} }
func (x *GetMedicalRecordResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*GetMedicalRecordResponse) ProtoMessage()    {}
func (x *GetMedicalRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_medical_medical_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*GetMedicalRecordResponse) Descriptor() ([]byte, []int) {
	return file_medical_medical_proto_rawDesc, []int{1}
}
func (x *GetMedicalRecordResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}
func (x *GetMedicalRecordResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}
func (x *GetMedicalRecordResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ListMedicalRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	UserId        int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListMedicalRecordsRequest) Reset()         { *x = ListMedicalRecordsRequest{} }
func (x *ListMedicalRecordsRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*ListMedicalRecordsRequest) ProtoMessage()    {}
func (x *ListMedicalRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_medical_medical_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*ListMedicalRecordsRequest) Descriptor() ([]byte, []int) {
	return file_medical_medical_proto_rawDesc, []int{2}
}
func (x *ListMedicalRecordsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type MedicalRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId      int64  `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	UserId        int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Description   string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	HospitalName  string `protobuf:"bytes,6,opt,name=hospital_name,json=hospitalName,proto3" json:"hospital_name,omitempty"`
	VisitDate     string `protobuf:"bytes,7,opt,name=visit_date,json=visitDate,proto3" json:"visit_date,omitempty"`
}

func (x *MedicalRecord) Reset()         { *x = MedicalRecord{} }
func (x *MedicalRecord) String() string { return protoimpl.X.MessageStringOf(x) }
func (*MedicalRecord) ProtoMessage()    {}
func (x *MedicalRecord) ProtoReflect() protoreflect.Message {
	mi := &file_medical_medical_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MedicalRecord) Descriptor() ([]byte, []int) {
	return file_medical_medical_proto_rawDesc, []int{3}
}
func (x *MedicalRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}
func (x *MedicalRecord) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}
func (x *MedicalRecord) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}
func (x *MedicalRecord) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}
func (x *MedicalRecord) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}
func (x *MedicalRecord) GetHospitalName() string {
	if x != nil {
		return x.HospitalName
	}
	return ""
}
func (x *MedicalRecord) GetVisitDate() string {
	if x != nil {
		return x.VisitDate
	}
	return ""
}

type ListMedicalRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Records       []*MedicalRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *ListMedicalRecordsResponse) Reset()         { *x = ListMedicalRecordsResponse{} }
func (x *ListMedicalRecordsResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*ListMedicalRecordsResponse) ProtoMessage()    {}
func (x *ListMedicalRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_medical_medical_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*ListMedicalRecordsResponse) Descriptor() ([]byte, []int) {
	return file_medical_medical_proto_rawDesc, []int{4}
}
func (x *ListMedicalRecordsResponse) GetRecords() []*MedicalRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

type DeleteMedicalRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Id            int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMedicalRecordRequest) Reset()         { *x = DeleteMedicalRecordRequest{} }
func (x *DeleteMedicalRecordRequest) String() string { return protoimpl.X.MessageStringOf(x) }
func (*DeleteMedicalRecordRequest) ProtoMessage()    {}
func (x *DeleteMedicalRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_medical_medical_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DeleteMedicalRecordRequest) Descriptor() ([]byte, []int) {
	return file_medical_medical_proto_rawDesc, []int{5}
}
func (x *DeleteMedicalRecordRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteMedicalRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMedicalRecordResponse) Reset()         { *x = DeleteMedicalRecordResponse{} }
func (x *DeleteMedicalRecordResponse) String() string { return protoimpl.X.MessageStringOf(x) }
func (*DeleteMedicalRecordResponse) ProtoMessage()    {}
func (x *DeleteMedicalRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_medical_medical_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DeleteMedicalRecordResponse) Descriptor() ([]byte, []int) {
	return file_medical_medical_proto_rawDesc, []int{6}
}

var File_medical_medical_proto protoreflect.FileDescriptor

var file_medical_medical_proto_rawDesc = buildRawDesc()

func buildRawDesc() []byte {
	fd := &descriptorpb.FileDescriptorProto{
		Name:    strPtr("medical/medical.proto"),
		Package: strPtr("medical"),
		MessageType: []*descriptorpb.DescriptorProto{
			{
				Name: strPtr("GetMedicalRecordRequest"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("id"), Number: int32Ptr(1), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_INT64), JsonName: strPtr("id")},
				},
			},
			{
				Name: strPtr("GetMedicalRecordResponse"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("id"), Number: int32Ptr(1), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_INT64), JsonName: strPtr("id")},
					{Name: strPtr("user_id"), Number: int32Ptr(2), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_INT64), JsonName: strPtr("userId")},
					{Name: strPtr("title"), Number: int32Ptr(3), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_STRING), JsonName: strPtr("title")},
				},
			},
			{
				Name: strPtr("ListMedicalRecordsRequest"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("user_id"), Number: int32Ptr(1), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_INT64), JsonName: strPtr("userId")},
				},
			},
			{
				Name: strPtr("MedicalRecord"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("id"), Number: int32Ptr(1), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_INT64), JsonName: strPtr("id")},
					{Name: strPtr("parent_id"), Number: int32Ptr(2), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_INT64), JsonName: strPtr("parentId")},
					{Name: strPtr("user_id"), Number: int32Ptr(3), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_INT64), JsonName: strPtr("userId")},
					{Name: strPtr("title"), Number: int32Ptr(4), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_STRING), JsonName: strPtr("title")},
					{Name: strPtr("description"), Number: int32Ptr(5), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_STRING), JsonName: strPtr("description")},
					{Name: strPtr("hospital_name"), Number: int32Ptr(6), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_STRING), JsonName: strPtr("hospitalName")},
					{Name: strPtr("visit_date"), Number: int32Ptr(7), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_STRING), JsonName: strPtr("visitDate")},
				},
			},
			{
				Name: strPtr("ListMedicalRecordsResponse"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("records"), Number: int32Ptr(1), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_REPEATED), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_MESSAGE), TypeName: strPtr(".medical.MedicalRecord"), JsonName: strPtr("records")},
				},
			},
			{
				Name: strPtr("DeleteMedicalRecordRequest"),
				Field: []*descriptorpb.FieldDescriptorProto{
					{Name: strPtr("id"), Number: int32Ptr(1), Label: labelPtr(descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL), Type: typePtr(descriptorpb.FieldDescriptorProto_TYPE_INT64), JsonName: strPtr("id")},
				},
			},
			{
				Name:  strPtr("DeleteMedicalRecordResponse"),
				Field: []*descriptorpb.FieldDescriptorProto{},
			},
		},
		Service: []*descriptorpb.ServiceDescriptorProto{
			{
				Name: strPtr("MedicalService"),
				Method: []*descriptorpb.MethodDescriptorProto{
					{Name: strPtr("GetMedicalRecord"), InputType: strPtr(".medical.GetMedicalRecordRequest"), OutputType: strPtr(".medical.GetMedicalRecordResponse")},
					{Name: strPtr("ListMedicalRecords"), InputType: strPtr(".medical.ListMedicalRecordsRequest"), OutputType: strPtr(".medical.ListMedicalRecordsResponse")},
					{Name: strPtr("DeleteMedicalRecord"), InputType: strPtr(".medical.DeleteMedicalRecordRequest"), OutputType: strPtr(".medical.DeleteMedicalRecordResponse")},
				},
			},
		},
	}
	b, _ := proto.Marshal(fd)
	return b
}

func strPtr(s string) *string { return &s }
func int32Ptr(i int) *int32   { v := int32(i); return &v }
func labelPtr(l descriptorpb.FieldDescriptorProto_Label) *descriptorpb.FieldDescriptorProto_Label {
	return &l
}
func typePtr(t descriptorpb.FieldDescriptorProto_Type) *descriptorpb.FieldDescriptorProto_Type {
	return &t
}

var file_medical_medical_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_medical_medical_proto_goTypes = []interface{}{
	(*GetMedicalRecordRequest)(nil),
	(*GetMedicalRecordResponse)(nil),
	(*ListMedicalRecordsRequest)(nil),
	(*MedicalRecord)(nil),
	(*ListMedicalRecordsResponse)(nil),
	(*DeleteMedicalRecordRequest)(nil),
	(*DeleteMedicalRecordResponse)(nil),
}
var file_medical_medical_proto_depIdxs = []int32{
	3,
	0,
	2,
	4,
	5,
	6,
	6,
}

func init() {
	file_medical_medical_proto_init()
}

func file_medical_medical_proto_init() {
	if File_medical_medical_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_medical_medical_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMedicalRecordRequest); i {
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
		file_medical_medical_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMedicalRecordResponse); i {
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
		file_medical_medical_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMedicalRecordsRequest); i {
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
		file_medical_medical_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedicalRecord); i {
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
		file_medical_medical_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMedicalRecordsResponse); i {
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
		file_medical_medical_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMedicalRecordRequest); i {
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
		file_medical_medical_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMedicalRecordResponse); i {
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
			GoPackagePath: reflectTypePath(x{}),
			RawDescriptor: file_medical_medical_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_medical_medical_proto_goTypes,
		DependencyIndexes: file_medical_medical_proto_depIdxs,
		MessageInfos:      file_medical_medical_proto_msgTypes,
	}.Build()
	File_medical_medical_proto = out.File
	file_medical_medical_proto_rawDesc = nil
	file_medical_medical_proto_goTypes = nil
	file_medical_medical_proto_depIdxs = nil

	_ = protoregistry.GlobalFiles.RegisterFile(File_medical_medical_proto)
}

func reflectTypePath(x interface{}) string {
	return "github.com/teakingwang/backend/api/medical"
}
