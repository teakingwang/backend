package medical

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion7

type MedicalServiceClient interface {
	GetMedicalRecord(ctx context.Context, in *GetMedicalRecordRequest, opts ...grpc.CallOption) (*GetMedicalRecordResponse, error)
	ListMedicalRecords(ctx context.Context, in *ListMedicalRecordsRequest, opts ...grpc.CallOption) (*ListMedicalRecordsResponse, error)
	DeleteMedicalRecord(ctx context.Context, in *DeleteMedicalRecordRequest, opts ...grpc.CallOption) (*DeleteMedicalRecordResponse, error)
}

type medicalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMedicalServiceClient(cc grpc.ClientConnInterface) MedicalServiceClient {
	return &medicalServiceClient{cc}
}

func (c *medicalServiceClient) GetMedicalRecord(ctx context.Context, in *GetMedicalRecordRequest, opts ...grpc.CallOption) (*GetMedicalRecordResponse, error) {
	out := new(GetMedicalRecordResponse)
	err := c.cc.Invoke(ctx, "/medical.MedicalService/GetMedicalRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalServiceClient) ListMedicalRecords(ctx context.Context, in *ListMedicalRecordsRequest, opts ...grpc.CallOption) (*ListMedicalRecordsResponse, error) {
	out := new(ListMedicalRecordsResponse)
	err := c.cc.Invoke(ctx, "/medical.MedicalService/ListMedicalRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicalServiceClient) DeleteMedicalRecord(ctx context.Context, in *DeleteMedicalRecordRequest, opts ...grpc.CallOption) (*DeleteMedicalRecordResponse, error) {
	out := new(DeleteMedicalRecordResponse)
	err := c.cc.Invoke(ctx, "/medical.MedicalService/DeleteMedicalRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type MedicalServiceServer interface {
	GetMedicalRecord(context.Context, *GetMedicalRecordRequest) (*GetMedicalRecordResponse, error)
	ListMedicalRecords(context.Context, *ListMedicalRecordsRequest) (*ListMedicalRecordsResponse, error)
	DeleteMedicalRecord(context.Context, *DeleteMedicalRecordRequest) (*DeleteMedicalRecordResponse, error)
}

type UnimplementedMedicalServiceServer struct{}

func (UnimplementedMedicalServiceServer) GetMedicalRecord(context.Context, *GetMedicalRecordRequest) (*GetMedicalRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedicalRecord not implemented")
}
func (UnimplementedMedicalServiceServer) ListMedicalRecords(context.Context, *ListMedicalRecordsRequest) (*ListMedicalRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMedicalRecords not implemented")
}
func (UnimplementedMedicalServiceServer) DeleteMedicalRecord(context.Context, *DeleteMedicalRecordRequest) (*DeleteMedicalRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMedicalRecord not implemented")
}

func RegisterMedicalServiceServer(s grpc.ServiceRegistrar, srv MedicalServiceServer) {
	s.RegisterService(&MedicalService_ServiceDesc, srv)
}

func _MedicalService_GetMedicalRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedicalRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalServiceServer).GetMedicalRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical.MedicalService/GetMedicalRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalServiceServer).GetMedicalRecord(ctx, req.(*GetMedicalRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalService_ListMedicalRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMedicalRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalServiceServer).ListMedicalRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical.MedicalService/ListMedicalRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalServiceServer).ListMedicalRecords(ctx, req.(*ListMedicalRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicalService_DeleteMedicalRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMedicalRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicalServiceServer).DeleteMedicalRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/medical.MedicalService/DeleteMedicalRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicalServiceServer).DeleteMedicalRecord(ctx, req.(*DeleteMedicalRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var MedicalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "medical.MedicalService",
	HandlerType: (*MedicalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "GetMedicalRecord", Handler: _MedicalService_GetMedicalRecord_Handler},
		{MethodName: "ListMedicalRecords", Handler: _MedicalService_ListMedicalRecords_Handler},
		{MethodName: "DeleteMedicalRecord", Handler: _MedicalService_DeleteMedicalRecord_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medical/medical.proto",
}
