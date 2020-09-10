// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data_catalog_service.proto

package connectors

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("data_catalog_service.proto", fileDescriptor_505bb950a1c4f82c)
}

var fileDescriptor_505bb950a1c4f82c = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x49, 0x2c, 0x49,
	0x8c, 0x4f, 0x4e, 0x2c, 0x49, 0xcc, 0xc9, 0x4f, 0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0xce, 0xcf, 0xcb, 0x4b, 0x4d, 0x2e, 0xc9,
	0x2f, 0x2a, 0x96, 0x42, 0x55, 0x57, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x02, 0x51, 0x27, 0x25,
	0x8d, 0x26, 0x57, 0x5c, 0x90, 0x9f, 0x57, 0x0c, 0x35, 0xc4, 0x28, 0x9b, 0x4b, 0xc8, 0x25, 0xb1,
	0x24, 0xd1, 0x19, 0x22, 0x1b, 0x0c, 0xb1, 0x40, 0x28, 0x94, 0x8b, 0xcf, 0x3d, 0xb5, 0x04, 0x24,
	0x51, 0x9c, 0x5a, 0xe2, 0x99, 0x97, 0x96, 0x2f, 0xa4, 0xa8, 0x87, 0xb0, 0x4d, 0x0f, 0xaa, 0x1a,
	0x2a, 0x1f, 0x04, 0xb1, 0x4d, 0x4a, 0x0e, 0xb7, 0x12, 0x90, 0x11, 0x4a, 0x0c, 0x4e, 0xbc, 0x5c,
	0xdc, 0xc9, 0xf9, 0xb9, 0x7a, 0x29, 0x89, 0x25, 0xb9, 0xa9, 0xc5, 0x19, 0x49, 0x6c, 0x60, 0x27,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x6b, 0xa4, 0x2c, 0xe5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataCatalogServiceClient is the client API for DataCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataCatalogServiceClient interface {
	GetDatasetInfo(ctx context.Context, in *CatalogDatasetRequest, opts ...grpc.CallOption) (*CatalogDatasetInfo, error)
}

type dataCatalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataCatalogServiceClient(cc grpc.ClientConnInterface) DataCatalogServiceClient {
	return &dataCatalogServiceClient{cc}
}

func (c *dataCatalogServiceClient) GetDatasetInfo(ctx context.Context, in *CatalogDatasetRequest, opts ...grpc.CallOption) (*CatalogDatasetInfo, error) {
	out := new(CatalogDatasetInfo)
	err := c.cc.Invoke(ctx, "/connectors.DataCatalogService/GetDatasetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataCatalogServiceServer is the server API for DataCatalogService service.
type DataCatalogServiceServer interface {
	GetDatasetInfo(context.Context, *CatalogDatasetRequest) (*CatalogDatasetInfo, error)
}

// UnimplementedDataCatalogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataCatalogServiceServer struct {
}

func (*UnimplementedDataCatalogServiceServer) GetDatasetInfo(ctx context.Context, req *CatalogDatasetRequest) (*CatalogDatasetInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetInfo not implemented")
}

func RegisterDataCatalogServiceServer(s *grpc.Server, srv DataCatalogServiceServer) {
	s.RegisterService(&_DataCatalogService_serviceDesc, srv)
}

func _DataCatalogService_GetDatasetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CatalogDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataCatalogServiceServer).GetDatasetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.DataCatalogService/GetDatasetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataCatalogServiceServer).GetDatasetInfo(ctx, req.(*CatalogDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataCatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connectors.DataCatalogService",
	HandlerType: (*DataCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDatasetInfo",
			Handler:    _DataCatalogService_GetDatasetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data_catalog_service.proto",
}
