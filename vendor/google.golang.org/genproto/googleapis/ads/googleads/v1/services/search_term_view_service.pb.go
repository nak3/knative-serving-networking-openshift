// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/search_term_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for
// [SearchTermViewService.GetSearchTermView][google.ads.googleads.v1.services.SearchTermViewService.GetSearchTermView].
type GetSearchTermViewRequest struct {
	// The resource name of the search term view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSearchTermViewRequest) Reset()         { *m = GetSearchTermViewRequest{} }
func (m *GetSearchTermViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetSearchTermViewRequest) ProtoMessage()    {}
func (*GetSearchTermViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05af83b1baea038c, []int{0}
}

func (m *GetSearchTermViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSearchTermViewRequest.Unmarshal(m, b)
}
func (m *GetSearchTermViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSearchTermViewRequest.Marshal(b, m, deterministic)
}
func (m *GetSearchTermViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSearchTermViewRequest.Merge(m, src)
}
func (m *GetSearchTermViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetSearchTermViewRequest.Size(m)
}
func (m *GetSearchTermViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSearchTermViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSearchTermViewRequest proto.InternalMessageInfo

func (m *GetSearchTermViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSearchTermViewRequest)(nil), "google.ads.googleads.v1.services.GetSearchTermViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/search_term_view_service.proto", fileDescriptor_05af83b1baea038c)
}

var fileDescriptor_05af83b1baea038c = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x6a, 0xe3, 0x40,
	0x14, 0x45, 0x5a, 0x58, 0x58, 0xb1, 0x5b, 0xac, 0x60, 0xc1, 0x88, 0x2d, 0x8c, 0xd7, 0xc5, 0xe2,
	0x62, 0x06, 0xc5, 0x4d, 0x32, 0x21, 0x18, 0xb9, 0x71, 0xaa, 0x60, 0xec, 0xa0, 0x22, 0x08, 0xc4,
	0x44, 0xfa, 0x28, 0x02, 0x4b, 0xe3, 0xcc, 0x97, 0xe5, 0x22, 0xa4, 0x48, 0xae, 0x90, 0x1b, 0xa4,
	0xcc, 0x1d, 0x72, 0x81, 0xb4, 0x29, 0x72, 0x81, 0x54, 0x39, 0x45, 0x90, 0x47, 0x23, 0x30, 0xb1,
	0x70, 0xf7, 0x98, 0xff, 0xde, 0xfb, 0xef, 0x3f, 0xc9, 0x1a, 0x25, 0x42, 0x24, 0x0b, 0xa0, 0x3c,
	0x46, 0xaa, 0x60, 0x85, 0x4a, 0x97, 0x22, 0xc8, 0x32, 0x8d, 0x00, 0x29, 0x02, 0x97, 0xd1, 0x55,
	0x58, 0x80, 0xcc, 0xc2, 0x32, 0x85, 0x75, 0x58, 0x4f, 0xc8, 0x52, 0x8a, 0x42, 0xd8, 0x5d, 0xa5,
	0x22, 0x3c, 0x46, 0xd2, 0x18, 0x90, 0xd2, 0x25, 0xda, 0xc0, 0x39, 0x6c, 0x5b, 0x21, 0x01, 0xc5,
	0x4a, 0xee, 0xda, 0xa1, 0xbc, 0x9d, 0xbf, 0x5a, 0xb9, 0x4c, 0x29, 0xcf, 0x73, 0x51, 0xf0, 0x22,
	0x15, 0x39, 0xaa, 0x69, 0x6f, 0x64, 0x75, 0x26, 0x50, 0xcc, 0x37, 0xd2, 0x73, 0x90, 0x99, 0x9f,
	0xc2, 0x7a, 0x06, 0xd7, 0x2b, 0xc0, 0xc2, 0xfe, 0x67, 0xfd, 0xd2, 0xee, 0x61, 0xce, 0x33, 0xe8,
	0x18, 0x5d, 0xe3, 0xff, 0x8f, 0xd9, 0x4f, 0xfd, 0x78, 0xc6, 0x33, 0x38, 0x78, 0x33, 0xac, 0x3f,
	0xdb, 0xf2, 0xb9, 0xca, 0x6c, 0x3f, 0x1b, 0xd6, 0xef, 0x2f, 0xde, 0x36, 0x23, 0xfb, 0x6e, 0x25,
	0x6d, 0x81, 0x1c, 0xb7, 0x55, 0xdb, 0xb4, 0x40, 0xb6, 0x95, 0xbd, 0xa3, 0xfb, 0xd7, 0xf7, 0x07,
	0x73, 0x68, 0xbb, 0x55, 0x57, 0x37, 0x5b, 0xe7, 0x9c, 0x44, 0x2b, 0x2c, 0x44, 0x06, 0x12, 0xe9,
	0xa0, 0x2e, 0x4f, 0xcb, 0x90, 0x0e, 0x6e, 0xc7, 0x77, 0xa6, 0xd5, 0x8f, 0x44, 0xb6, 0x37, 0xef,
	0xd8, 0xd9, 0x79, 0xff, 0xb4, 0xea, 0x77, 0x6a, 0x5c, 0x9c, 0xd6, 0xfa, 0x44, 0x2c, 0x78, 0x9e,
	0x10, 0x21, 0x13, 0x9a, 0x40, 0xbe, 0x69, 0x5f, 0x7f, 0xc9, 0x65, 0x8a, 0xed, 0xff, 0xce, 0xb1,
	0x06, 0x8f, 0xe6, 0xb7, 0x89, 0xe7, 0x3d, 0x99, 0xdd, 0x89, 0x32, 0xf4, 0x62, 0x24, 0x0a, 0x56,
	0xc8, 0x77, 0x49, 0xbd, 0x18, 0x5f, 0x34, 0x25, 0xf0, 0x62, 0x0c, 0x1a, 0x4a, 0xe0, 0xbb, 0x81,
	0xa6, 0x7c, 0x98, 0x7d, 0xf5, 0xce, 0x98, 0x17, 0x23, 0x63, 0x0d, 0x89, 0x31, 0xdf, 0x65, 0x4c,
	0xd3, 0x2e, 0xbf, 0x6f, 0x72, 0x0e, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x3e, 0x6e, 0xc9,
	0xe2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchTermViewServiceClient is the client API for SearchTermViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchTermViewServiceClient interface {
	// Returns the attributes of the requested search term view.
	GetSearchTermView(ctx context.Context, in *GetSearchTermViewRequest, opts ...grpc.CallOption) (*resources.SearchTermView, error)
}

type searchTermViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchTermViewServiceClient(cc *grpc.ClientConn) SearchTermViewServiceClient {
	return &searchTermViewServiceClient{cc}
}

func (c *searchTermViewServiceClient) GetSearchTermView(ctx context.Context, in *GetSearchTermViewRequest, opts ...grpc.CallOption) (*resources.SearchTermView, error) {
	out := new(resources.SearchTermView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.SearchTermViewService/GetSearchTermView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchTermViewServiceServer is the server API for SearchTermViewService service.
type SearchTermViewServiceServer interface {
	// Returns the attributes of the requested search term view.
	GetSearchTermView(context.Context, *GetSearchTermViewRequest) (*resources.SearchTermView, error)
}

// UnimplementedSearchTermViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchTermViewServiceServer struct {
}

func (*UnimplementedSearchTermViewServiceServer) GetSearchTermView(ctx context.Context, req *GetSearchTermViewRequest) (*resources.SearchTermView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchTermView not implemented")
}

func RegisterSearchTermViewServiceServer(s *grpc.Server, srv SearchTermViewServiceServer) {
	s.RegisterService(&_SearchTermViewService_serviceDesc, srv)
}

func _SearchTermViewService_GetSearchTermView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSearchTermViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchTermViewServiceServer).GetSearchTermView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.SearchTermViewService/GetSearchTermView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchTermViewServiceServer).GetSearchTermView(ctx, req.(*GetSearchTermViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchTermViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.SearchTermViewService",
	HandlerType: (*SearchTermViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSearchTermView",
			Handler:    _SearchTermViewService_GetSearchTermView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/search_term_view_service.proto",
}
