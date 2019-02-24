// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: prod.proto

package example

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golangper/protoc-gen-rorm/options"
import _ "github.com/lyft/protoc-gen-validate/validate"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Prod struct {
	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Details string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	Skus    []*Sku `protobuf:"bytes,4,rep,name=skus" json:"skus,omitempty"`
}

func (m *Prod) Reset()         { *m = Prod{} }
func (m *Prod) String() string { return proto.CompactTextString(m) }
func (*Prod) ProtoMessage()    {}
func (*Prod) Descriptor() ([]byte, []int) {
	return fileDescriptor_prod_5e0166586b3bf9b4, []int{0}
}
func (m *Prod) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Prod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Prod.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Prod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prod.Merge(dst, src)
}
func (m *Prod) XXX_Size() int {
	return m.Size()
}
func (m *Prod) XXX_DiscardUnknown() {
	xxx_messageInfo_Prod.DiscardUnknown(m)
}

var xxx_messageInfo_Prod proto.InternalMessageInfo

func (m *Prod) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Prod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Prod) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *Prod) GetSkus() []*Sku {
	if m != nil {
		return m.Skus
	}
	return nil
}

type Sku struct {
	SkuId  int64   `protobuf:"varint,1,opt,name=skuId,proto3" json:"skuId,omitempty"`
	Price  float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Bn     string  `protobuf:"bytes,3,opt,name=bn,proto3" json:"bn,omitempty"`
	Weight float32 `protobuf:"fixed32,4,opt,name=weight,proto3" json:"weight,omitempty"`
	ProdId int64   `protobuf:"varint,5,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
}

func (m *Sku) Reset()         { *m = Sku{} }
func (m *Sku) String() string { return proto.CompactTextString(m) }
func (*Sku) ProtoMessage()    {}
func (*Sku) Descriptor() ([]byte, []int) {
	return fileDescriptor_prod_5e0166586b3bf9b4, []int{1}
}
func (m *Sku) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sku) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sku.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Sku) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sku.Merge(dst, src)
}
func (m *Sku) XXX_Size() int {
	return m.Size()
}
func (m *Sku) XXX_DiscardUnknown() {
	xxx_messageInfo_Sku.DiscardUnknown(m)
}

var xxx_messageInfo_Sku proto.InternalMessageInfo

func (m *Sku) GetSkuId() int64 {
	if m != nil {
		return m.SkuId
	}
	return 0
}

func (m *Sku) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Sku) GetBn() string {
	if m != nil {
		return m.Bn
	}
	return ""
}

func (m *Sku) GetWeight() float32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Sku) GetProdId() int64 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

type ProdId struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *ProdId) Reset()         { *m = ProdId{} }
func (m *ProdId) String() string { return proto.CompactTextString(m) }
func (*ProdId) ProtoMessage()    {}
func (*ProdId) Descriptor() ([]byte, []int) {
	return fileDescriptor_prod_5e0166586b3bf9b4, []int{2}
}
func (m *ProdId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProdId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProdId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ProdId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdId.Merge(dst, src)
}
func (m *ProdId) XXX_Size() int {
	return m.Size()
}
func (m *ProdId) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdId.DiscardUnknown(m)
}

var xxx_messageInfo_ProdId proto.InternalMessageInfo

func (m *ProdId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Empty struct {
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_prod_5e0166586b3bf9b4, []int{3}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return m.Size()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Prod)(nil), "example.Prod")
	proto.RegisterType((*Sku)(nil), "example.Sku")
	proto.RegisterType((*ProdId)(nil), "example.ProdId")
	proto.RegisterType((*Empty)(nil), "example.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductClient interface {
	GetProd(ctx context.Context, in *ProdId, opts ...grpc.CallOption) (*Prod, error)
	SetProd(ctx context.Context, in *Prod, opts ...grpc.CallOption) (*Empty, error)
}

type productClient struct {
	cc *grpc.ClientConn
}

func NewProductClient(cc *grpc.ClientConn) ProductClient {
	return &productClient{cc}
}

func (c *productClient) GetProd(ctx context.Context, in *ProdId, opts ...grpc.CallOption) (*Prod, error) {
	out := new(Prod)
	err := c.cc.Invoke(ctx, "/example.product/getProd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) SetProd(ctx context.Context, in *Prod, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/example.product/setProd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
type ProductServer interface {
	GetProd(context.Context, *ProdId) (*Prod, error)
	SetProd(context.Context, *Prod) (*Empty, error)
}

func RegisterProductServer(s *grpc.Server, srv ProductServer) {
	s.RegisterService(&_Product_serviceDesc, srv)
}

func _Product_GetProd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetProd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.product/GetProd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetProd(ctx, req.(*ProdId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_SetProd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Prod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).SetProd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.product/SetProd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).SetProd(ctx, req.(*Prod))
	}
	return interceptor(ctx, in, info, handler)
}

var _Product_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getProd",
			Handler:    _Product_GetProd_Handler,
		},
		{
			MethodName: "setProd",
			Handler:    _Product_SetProd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prod.proto",
}

// Product2Client is the client API for Product2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Product2Client interface {
	GetProd(ctx context.Context, in *ProdId, opts ...grpc.CallOption) (*Prod, error)
}

type product2Client struct {
	cc *grpc.ClientConn
}

func NewProduct2Client(cc *grpc.ClientConn) Product2Client {
	return &product2Client{cc}
}

func (c *product2Client) GetProd(ctx context.Context, in *ProdId, opts ...grpc.CallOption) (*Prod, error) {
	out := new(Prod)
	err := c.cc.Invoke(ctx, "/example.product2/getProd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Product2Server is the server API for Product2 service.
type Product2Server interface {
	GetProd(context.Context, *ProdId) (*Prod, error)
}

func RegisterProduct2Server(s *grpc.Server, srv Product2Server) {
	s.RegisterService(&_Product2_serviceDesc, srv)
}

func _Product2_GetProd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Product2Server).GetProd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.product2/GetProd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Product2Server).GetProd(ctx, req.(*ProdId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Product2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.product2",
	HandlerType: (*Product2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getProd",
			Handler:    _Product2_GetProd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prod.proto",
}

func (m *Prod) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Prod) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProd(dAtA, i, uint64(m.Id))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProd(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Details) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProd(dAtA, i, uint64(len(m.Details)))
		i += copy(dAtA[i:], m.Details)
	}
	if len(m.Skus) > 0 {
		for _, msg := range m.Skus {
			dAtA[i] = 0x22
			i++
			i = encodeVarintProd(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Sku) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sku) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SkuId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProd(dAtA, i, uint64(m.SkuId))
	}
	if m.Price != 0 {
		dAtA[i] = 0x15
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Price))))
		i += 4
	}
	if len(m.Bn) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProd(dAtA, i, uint64(len(m.Bn)))
		i += copy(dAtA[i:], m.Bn)
	}
	if m.Weight != 0 {
		dAtA[i] = 0x25
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Weight))))
		i += 4
	}
	if m.ProdId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintProd(dAtA, i, uint64(m.ProdId))
	}
	return i, nil
}

func (m *ProdId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProdId) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProd(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintProd(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Prod) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovProd(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProd(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovProd(uint64(l))
	}
	if len(m.Skus) > 0 {
		for _, e := range m.Skus {
			l = e.Size()
			n += 1 + l + sovProd(uint64(l))
		}
	}
	return n
}

func (m *Sku) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SkuId != 0 {
		n += 1 + sovProd(uint64(m.SkuId))
	}
	if m.Price != 0 {
		n += 5
	}
	l = len(m.Bn)
	if l > 0 {
		n += 1 + l + sovProd(uint64(l))
	}
	if m.Weight != 0 {
		n += 5
	}
	if m.ProdId != 0 {
		n += 1 + sovProd(uint64(m.ProdId))
	}
	return n
}

func (m *ProdId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovProd(uint64(m.Id))
	}
	return n
}

func (m *Empty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovProd(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProd(x uint64) (n int) {
	return sovProd(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Prod) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProd
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Prod: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Prod: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProd
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProd
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Skus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProd
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Skus = append(m.Skus, &Sku{})
			if err := m.Skus[len(m.Skus)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProd
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Sku) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProd
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Sku: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sku: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkuId", wireType)
			}
			m.SkuId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SkuId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Price = float32(math.Float32frombits(v))
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProd
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Weight = float32(math.Float32frombits(v))
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProdId", wireType)
			}
			m.ProdId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProdId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProd
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProdId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProd
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProdId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProdId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProd
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProd
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProd
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProd(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProd
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
					return 0, ErrIntOverflowProd
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
					return 0, ErrIntOverflowProd
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProd
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProd
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
				next, err := skipProd(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthProd = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProd   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("prod.proto", fileDescriptor_prod_5e0166586b3bf9b4) }

var fileDescriptor_prod_5e0166586b3bf9b4 = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0x13, 0x27, 0x6e, 0xa7, 0xfd, 0xfa, 0xa1, 0x51, 0xdd, 0x5a, 0x5e, 0x44, 0x23, 0x4b,
	0x88, 0x14, 0xb9, 0xb6, 0x08, 0x0b, 0xa4, 0x56, 0x55, 0x04, 0x52, 0x81, 0x4a, 0x2c, 0x4a, 0xba,
	0x63, 0x53, 0xd9, 0x99, 0x89, 0x33, 0xc4, 0xf6, 0x58, 0xf6, 0xb8, 0x3f, 0x12, 0x0f, 0x80, 0xca,
	0x06, 0xb1, 0xa4, 0x2f, 0xc1, 0x23, 0x20, 0x16, 0xc0, 0xb2, 0x3b, 0xb3, 0x44, 0xed, 0x8e, 0x27,
	0x60, 0x89, 0x66, 0xc6, 0xa9, 0xd2, 0x76, 0xc3, 0x9a, 0xd5, 0xdc, 0x3b, 0xb9, 0xe7, 0xdc, 0x7b,
	0xce, 0x9d, 0x18, 0x80, 0x2c, 0x67, 0xd8, 0xcb, 0x72, 0xc6, 0x19, 0x34, 0xc8, 0x71, 0x90, 0x64,
	0x31, 0xb1, 0x1f, 0x45, 0x94, 0x8f, 0xcb, 0xd0, 0x1b, 0xb2, 0xc4, 0x8f, 0x58, 0x1c, 0xa4, 0x51,
	0x46, 0x72, 0x5f, 0x16, 0x0d, 0x37, 0x22, 0x92, 0x6e, 0xe4, 0x2c, 0x4f, 0x7c, 0x96, 0x71, 0xca,
	0xd2, 0xc2, 0x17, 0x89, 0x62, 0xb0, 0xb7, 0x66, 0x80, 0xf1, 0xc9, 0x88, 0xcf, 0x62, 0x0e, 0x83,
	0x98, 0xe2, 0x80, 0x13, 0xff, 0x56, 0xa0, 0xc0, 0xce, 0x08, 0xe8, 0x7b, 0x39, 0xc3, 0x70, 0x19,
	0x34, 0x28, 0xb6, 0x34, 0xa4, 0x75, 0x9b, 0x83, 0x06, 0xc5, 0x10, 0x02, 0x3d, 0x0d, 0x12, 0x62,
	0x35, 0x90, 0xd6, 0x5d, 0x18, 0xc8, 0x18, 0x5a, 0xc0, 0xc0, 0x84, 0x07, 0x34, 0x2e, 0xac, 0xa6,
	0xbc, 0x9e, 0xa6, 0x10, 0x01, 0xbd, 0x98, 0x94, 0x85, 0xa5, 0xa3, 0x66, 0x77, 0xb1, 0xb7, 0xe4,
	0xd5, 0x9a, 0xbc, 0xfd, 0x49, 0x39, 0x90, 0xbf, 0x38, 0x39, 0x68, 0xee, 0x4f, 0x4a, 0xb8, 0x02,
	0x5a, 0xc5, 0xa4, 0xdc, 0x9d, 0x76, 0x52, 0x89, 0xb8, 0xcd, 0x72, 0x3a, 0x54, 0xdd, 0x1a, 0x03,
	0x95, 0x88, 0x91, 0xc2, 0xb4, 0xee, 0xd4, 0x08, 0x53, 0xb8, 0x0a, 0xda, 0x47, 0x84, 0x46, 0x63,
	0x6e, 0xe9, 0xb2, 0xac, 0xce, 0xe0, 0x1a, 0x30, 0x84, 0x9f, 0x07, 0x14, 0x5b, 0x2d, 0xc9, 0xda,
	0x16, 0xe9, 0x2e, 0x76, 0x2c, 0xd0, 0xde, 0x93, 0xd1, 0x4d, 0x75, 0x8e, 0x01, 0x5a, 0x3b, 0x49,
	0xc6, 0x4f, 0x7a, 0x9f, 0x75, 0x05, 0x2e, 0x87, 0x1c, 0x7e, 0xd3, 0x80, 0x11, 0x11, 0x2e, 0xed,
	0xf8, 0xff, 0x4a, 0x82, 0x62, 0xb0, 0xff, 0xbb, 0x76, 0xe1, 0x9c, 0x69, 0xa7, 0x95, 0xf9, 0x06,
	0x34, 0x59, 0xc9, 0xe1, 0xe2, 0xb1, 0xd8, 0xc5, 0xfe, 0xcb, 0x17, 0xcf, 0x08, 0xb7, 0xef, 0x16,
	0x24, 0x26, 0x43, 0x8e, 0xee, 0xa3, 0x51, 0xce, 0x12, 0x24, 0xb8, 0xd1, 0xd1, 0x98, 0xe4, 0x04,
	0x51, 0x8c, 0xb6, 0x51, 0x7f, 0x8b, 0xa6, 0x1e, 0xc5, 0x9b, 0x8f, 0xc1, 0x3c, 0x2b, 0xb9, 0x27,
	0xac, 0x81, 0x4b, 0x53, 0xfc, 0x53, 0x9a, 0x62, 0xfb, 0xde, 0x75, 0x82, 0x62, 0x52, 0xd6, 0xf8,
	0x5a, 0xe3, 0x76, 0x4d, 0xf1, 0xb1, 0x32, 0xd7, 0x40, 0x33, 0x22, 0x1c, 0xde, 0xf1, 0x0f, 0x1f,
	0x88, 0x9d, 0x63, 0x7f, 0x3a, 0xfd, 0xd7, 0x06, 0x30, 0x8a, 0x3a, 0xbe, 0x3e, 0xb8, 0xbd, 0x7c,
	0x95, 0x4a, 0x03, 0x9c, 0xb3, 0xc6, 0x69, 0x65, 0xfe, 0xd2, 0x9c, 0x57, 0x4a, 0xcb, 0x82, 0x9c,
	0x65, 0xe7, 0x98, 0x0c, 0xed, 0xe7, 0x34, 0x2d, 0x48, 0xce, 0x11, 0x4d, 0x39, 0x53, 0x3a, 0xba,
	0x14, 0xbb, 0xe2, 0x21, 0xb8, 0xf5, 0xda, 0xd7, 0xd1, 0x61, 0x10, 0x97, 0xa4, 0x40, 0xdd, 0xbe,
	0xdb, 0x77, 0xfb, 0xeb, 0x5b, 0x25, 0xc5, 0x62, 0x38, 0x51, 0x23, 0xce, 0xba, 0xcc, 0x79, 0xa7,
	0xdd, 0x22, 0xcf, 0x66, 0xc9, 0x85, 0xc6, 0x6e, 0x31, 0x29, 0x0f, 0x28, 0x76, 0xe5, 0xea, 0xdd,
	0x30, 0x75, 0xd5, 0x72, 0xdd, 0x5a, 0xf5, 0x8d, 0x56, 0xaa, 0x1d, 0x0b, 0x5f, 0x7b, 0x0a, 0x26,
	0x43, 0x09, 0x95, 0x51, 0x98, 0xca, 0x43, 0x51, 0x28, 0xbb, 0x1c, 0x83, 0xa6, 0xd2, 0xee, 0x0f,
	0x95, 0xb9, 0x04, 0x9a, 0x25, 0xc5, 0xb0, 0x35, 0xf5, 0xd1, 0x02, 0x7a, 0xc6, 0x8a, 0x59, 0x23,
	0x6b, 0xf3, 0x6c, 0xe7, 0x6d, 0x65, 0x6a, 0x9f, 0x2a, 0x53, 0xfb, 0x52, 0x99, 0xab, 0x60, 0x65,
	0xcc, 0x79, 0x56, 0x6c, 0xfa, 0x7e, 0x72, 0x32, 0x62, 0x0c, 0x53, 0x22, 0xfe, 0x72, 0xbd, 0xdf,
	0x1a, 0x98, 0xaf, 0x9f, 0x50, 0xef, 0xdf, 0x79, 0x43, 0x7f, 0x23, 0xfd, 0x89, 0xf5, 0xfd, 0xa2,
	0xa3, 0x9d, 0x5f, 0x74, 0xb4, 0x9f, 0x17, 0x1d, 0xed, 0xfd, 0x65, 0x67, 0xee, 0xfc, 0xb2, 0x33,
	0xf7, 0xe3, 0xb2, 0x33, 0x17, 0xb6, 0xe5, 0xd7, 0xe5, 0xe1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x9d, 0x68, 0x51, 0x2b, 0xea, 0x04, 0x00, 0x00,
}
