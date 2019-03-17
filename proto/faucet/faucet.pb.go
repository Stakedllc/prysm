// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/proto/faucet.proto

package faucet

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FundingRequest struct {
	WalletAddress        string   `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	RecaptchaSiteKey     string   `protobuf:"bytes,2,opt,name=recaptcha_site_key,json=recaptchaSiteKey,proto3" json:"recaptcha_site_key,omitempty"`
	RecaptchaResponse    string   `protobuf:"bytes,3,opt,name=recaptcha_response,json=recaptchaResponse,proto3" json:"recaptcha_response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FundingRequest) Reset()         { *m = FundingRequest{} }
func (m *FundingRequest) String() string { return proto.CompactTextString(m) }
func (*FundingRequest) ProtoMessage()    {}
func (*FundingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4118183f16331b80, []int{0}
}
func (m *FundingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingRequest.Merge(m, src)
}
func (m *FundingRequest) XXX_Size() int {
	return m.Size()
}
func (m *FundingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FundingRequest proto.InternalMessageInfo

func (m *FundingRequest) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *FundingRequest) GetRecaptchaSiteKey() string {
	if m != nil {
		return m.RecaptchaSiteKey
	}
	return ""
}

func (m *FundingRequest) GetRecaptchaResponse() string {
	if m != nil {
		return m.RecaptchaResponse
	}
	return ""
}

type FundingResponse struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Amount               string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	TransactionHash      string   `protobuf:"bytes,3,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FundingResponse) Reset()         { *m = FundingResponse{} }
func (m *FundingResponse) String() string { return proto.CompactTextString(m) }
func (*FundingResponse) ProtoMessage()    {}
func (*FundingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4118183f16331b80, []int{1}
}
func (m *FundingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FundingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FundingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FundingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FundingResponse.Merge(m, src)
}
func (m *FundingResponse) XXX_Size() int {
	return m.Size()
}
func (m *FundingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FundingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FundingResponse proto.InternalMessageInfo

func (m *FundingResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *FundingResponse) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *FundingResponse) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

func init() {
	proto.RegisterType((*FundingRequest)(nil), "faucet.FundingRequest")
	proto.RegisterType((*FundingResponse)(nil), "faucet.FundingResponse")
}

func init() { proto.RegisterFile("src/proto/faucet.proto", fileDescriptor_4118183f16331b80) }

var fileDescriptor_4118183f16331b80 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x4a, 0xf4, 0x40,
	0x10, 0x86, 0xbf, 0xfe, 0xc4, 0x80, 0xcd, 0xfc, 0x68, 0x23, 0x31, 0xb8, 0x08, 0x32, 0x20, 0xcc,
	0x42, 0x67, 0x40, 0x4f, 0x30, 0x2e, 0x06, 0xc1, 0x5d, 0xe6, 0x00, 0xa1, 0xed, 0x94, 0x4e, 0xe3,
	0xd8, 0x1d, 0xab, 0x2a, 0xca, 0xdc, 0xc3, 0x43, 0xb9, 0xf4, 0x08, 0x92, 0x93, 0x88, 0xe9, 0x26,
	0xf8, 0xb3, 0x7c, 0x9f, 0x7a, 0xa1, 0x5e, 0x1e, 0x99, 0x12, 0x9a, 0x79, 0x8d, 0x9e, 0xfd, 0xfc,
	0x4e, 0x37, 0x06, 0x78, 0xd6, 0x05, 0x95, 0x84, 0x34, 0x79, 0x15, 0x72, 0xb4, 0x6c, 0x5c, 0x65,
	0xdd, 0x7d, 0x01, 0x4f, 0x0d, 0x10, 0xab, 0x53, 0x39, 0x7a, 0xd1, 0x9b, 0x0d, 0x70, 0xa9, 0xab,
	0x0a, 0x81, 0x28, 0x13, 0x27, 0x62, 0xba, 0x57, 0x0c, 0x03, 0x5d, 0x04, 0xa8, 0xce, 0xa4, 0x42,
	0x30, 0xba, 0x66, 0xb3, 0xd6, 0x25, 0x59, 0x86, 0xf2, 0x01, 0xb6, 0xd9, 0xff, 0xae, 0xba, 0xdf,
	0x5f, 0x56, 0x96, 0xe1, 0x06, 0xb6, 0xea, 0xfc, 0x7b, 0x1b, 0x81, 0x6a, 0xef, 0x08, 0xb2, 0x9d,
	0xae, 0x7d, 0xd0, 0x5f, 0x8a, 0x78, 0x98, 0x58, 0x39, 0xee, 0x57, 0x05, 0xa4, 0x0e, 0xe5, 0x2e,
	0x20, 0x7a, 0x8c, 0x6b, 0x42, 0x50, 0xa9, 0x4c, 0xf4, 0xa3, 0x6f, 0x1c, 0xc7, 0xcf, 0x31, 0xa9,
	0xa9, 0x1c, 0x33, 0x6a, 0x47, 0xda, 0xb0, 0xf5, 0xee, 0x5a, 0xd3, 0x3a, 0x3e, 0xfb, 0x8d, 0x2f,
	0x0a, 0x39, 0x5c, 0x76, 0x2e, 0x56, 0x80, 0xcf, 0xd6, 0x80, 0x5a, 0xc8, 0x41, 0x54, 0xf1, 0x35,
	0x81, 0x54, 0x3a, 0x8b, 0xe6, 0x7e, 0x7a, 0x3a, 0x3e, 0xfa, 0xc3, 0xe3, 0xf8, 0x7f, 0x57, 0x83,
	0xb7, 0x36, 0x17, 0xef, 0x6d, 0x2e, 0x3e, 0xda, 0x5c, 0xdc, 0x26, 0x9d, 0xf2, 0xcb, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x9a, 0x27, 0xc6, 0x35, 0x8c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FaucetServiceClient is the client API for FaucetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FaucetServiceClient interface {
	RequestFunds(ctx context.Context, in *FundingRequest, opts ...grpc.CallOption) (*FundingResponse, error)
}

type faucetServiceClient struct {
	cc *grpc.ClientConn
}

func NewFaucetServiceClient(cc *grpc.ClientConn) FaucetServiceClient {
	return &faucetServiceClient{cc}
}

func (c *faucetServiceClient) RequestFunds(ctx context.Context, in *FundingRequest, opts ...grpc.CallOption) (*FundingResponse, error) {
	out := new(FundingResponse)
	err := c.cc.Invoke(ctx, "/faucet.FaucetService/RequestFunds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaucetServiceServer is the server API for FaucetService service.
type FaucetServiceServer interface {
	RequestFunds(context.Context, *FundingRequest) (*FundingResponse, error)
}

func RegisterFaucetServiceServer(s *grpc.Server, srv FaucetServiceServer) {
	s.RegisterService(&_FaucetService_serviceDesc, srv)
}

func _FaucetService_RequestFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaucetServiceServer).RequestFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/faucet.FaucetService/RequestFunds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaucetServiceServer).RequestFunds(ctx, req.(*FundingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FaucetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "faucet.FaucetService",
	HandlerType: (*FaucetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestFunds",
			Handler:    _FaucetService_RequestFunds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/faucet.proto",
}

func (m *FundingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundingRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.WalletAddress) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFaucet(dAtA, i, uint64(len(m.WalletAddress)))
		i += copy(dAtA[i:], m.WalletAddress)
	}
	if len(m.RecaptchaSiteKey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFaucet(dAtA, i, uint64(len(m.RecaptchaSiteKey)))
		i += copy(dAtA[i:], m.RecaptchaSiteKey)
	}
	if len(m.RecaptchaResponse) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFaucet(dAtA, i, uint64(len(m.RecaptchaResponse)))
		i += copy(dAtA[i:], m.RecaptchaResponse)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *FundingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FundingResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFaucet(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	if len(m.Amount) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFaucet(dAtA, i, uint64(len(m.Amount)))
		i += copy(dAtA[i:], m.Amount)
	}
	if len(m.TransactionHash) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFaucet(dAtA, i, uint64(len(m.TransactionHash)))
		i += copy(dAtA[i:], m.TransactionHash)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintFaucet(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FundingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WalletAddress)
	if l > 0 {
		n += 1 + l + sovFaucet(uint64(l))
	}
	l = len(m.RecaptchaSiteKey)
	if l > 0 {
		n += 1 + l + sovFaucet(uint64(l))
	}
	l = len(m.RecaptchaResponse)
	if l > 0 {
		n += 1 + l + sovFaucet(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FundingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovFaucet(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovFaucet(uint64(l))
	}
	l = len(m.TransactionHash)
	if l > 0 {
		n += 1 + l + sovFaucet(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFaucet(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFaucet(x uint64) (n int) {
	return sovFaucet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FundingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFaucet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FundingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaucet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFaucet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFaucet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WalletAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecaptchaSiteKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaucet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFaucet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFaucet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecaptchaSiteKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecaptchaResponse", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaucet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFaucet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFaucet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecaptchaResponse = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFaucet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFaucet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFaucet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FundingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFaucet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FundingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FundingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaucet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFaucet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFaucet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaucet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFaucet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFaucet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransactionHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFaucet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFaucet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFaucet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransactionHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFaucet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFaucet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFaucet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFaucet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFaucet
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFaucet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFaucet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFaucet
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFaucet
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFaucet
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFaucet(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFaucet
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFaucet = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFaucet   = fmt.Errorf("proto: integer overflow")
)
