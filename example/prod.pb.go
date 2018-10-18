// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/prod.proto

/*
	Package example is a generated protocol buffer package.

	It is generated from these files:
		example/prod.proto

	It has these top-level messages:
		Prod
		Sku
		ProdId
		Empty
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golangper/protoc-gen-rorm/options"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// import "github.com/golangper/protoc-gen-rorm/example/user.proto";
type Prod struct {
	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Details string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	Skus    []*Sku `protobuf:"bytes,4,rep,name=skus" json:"skus,omitempty"`
}

func (m *Prod) Reset()                    { *m = Prod{} }
func (m *Prod) String() string            { return proto.CompactTextString(m) }
func (*Prod) ProtoMessage()               {}
func (*Prod) Descriptor() ([]byte, []int) { return fileDescriptorProd, []int{0} }

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
	SkuId  int64   `protobuf:"varint,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	Price  float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Bn     string  `protobuf:"bytes,3,opt,name=bn,proto3" json:"bn,omitempty"`
	Weight float32 `protobuf:"fixed32,4,opt,name=weight,proto3" json:"weight,omitempty"`
	ProdId int64   `protobuf:"varint,5,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
}

func (m *Sku) Reset()                    { *m = Sku{} }
func (m *Sku) String() string            { return proto.CompactTextString(m) }
func (*Sku) ProtoMessage()               {}
func (*Sku) Descriptor() ([]byte, []int) { return fileDescriptorProd, []int{1} }

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

func (m *ProdId) Reset()                    { *m = ProdId{} }
func (m *ProdId) String() string            { return proto.CompactTextString(m) }
func (*ProdId) ProtoMessage()               {}
func (*ProdId) Descriptor() ([]byte, []int) { return fileDescriptorProd, []int{2} }

func (m *ProdId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptorProd, []int{3} }

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

// Client API for Product service

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
	err := grpc.Invoke(ctx, "/example.product/getProd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) SetProd(ctx context.Context, in *Prod, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/example.product/setProd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Product service

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
	Metadata: "example/prod.proto",
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
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Price))))
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
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Weight))))
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
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovProd(uint64(m.Id))
	}
	return n
}

func (m *Empty) Size() (n int) {
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
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
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
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
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

func init() { proto.RegisterFile("example/prod.proto", fileDescriptorProd) }

var fileDescriptorProd = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x4f, 0x6b, 0xd4, 0x4e,
	0x18, 0xfe, 0x25, 0xd9, 0xdd, 0xb4, 0xd3, 0xfe, 0xaa, 0x0c, 0x46, 0x63, 0x0e, 0xcb, 0x90, 0xcb,
	0x6e, 0x25, 0x4d, 0xb0, 0x1e, 0x04, 0x4b, 0x59, 0x28, 0x14, 0xa9, 0xa7, 0xb2, 0xbd, 0x0a, 0xb2,
	0xd9, 0x99, 0xa6, 0xe3, 0xee, 0xce, 0xc4, 0xcc, 0x4c, 0x5b, 0xc1, 0x9b, 0x17, 0x29, 0x08, 0xe2,
	0x49, 0x04, 0xbf, 0x82, 0xf8, 0x31, 0xf4, 0xe6, 0x2d, 0x57, 0xa9, 0x77, 0x3f, 0x83, 0xcc, 0x9f,
	0x16, 0xba, 0x9e, 0xf2, 0xbe, 0x33, 0xef, 0xfb, 0xbc, 0xcf, 0xf3, 0xcc, 0x1b, 0x00, 0xc9, 0xf9,
	0x64, 0x51, 0xcf, 0x49, 0x51, 0x37, 0x1c, 0xe7, 0x75, 0xc3, 0x25, 0x87, 0xa1, 0x3b, 0x4b, 0x1e,
	0x57, 0x54, 0x9e, 0xa8, 0x32, 0x9f, 0xf2, 0x45, 0x51, 0xf1, 0xf9, 0x84, 0x55, 0x35, 0x69, 0x0a,
	0x53, 0x34, 0xdd, 0xaa, 0x08, 0xdb, 0x6a, 0x78, 0xb3, 0x28, 0x78, 0x2d, 0x29, 0x67, 0xa2, 0xd0,
	0x89, 0x45, 0x48, 0x8f, 0x41, 0xe7, 0xb0, 0xe1, 0x18, 0x6e, 0x00, 0x9f, 0xe2, 0xd8, 0x43, 0xde,
	0x30, 0x18, 0xfb, 0x14, 0x43, 0x08, 0x3a, 0x6c, 0xb2, 0x20, 0xb1, 0x8f, 0xbc, 0xe1, 0xea, 0xd8,
	0xc4, 0x30, 0x06, 0x21, 0x26, 0x72, 0x42, 0xe7, 0x22, 0x0e, 0xcc, 0xf1, 0x55, 0x0a, 0x11, 0xe8,
	0x88, 0x99, 0x12, 0x71, 0x07, 0x05, 0xc3, 0xb5, 0xed, 0xf5, 0xdc, 0xd1, 0xca, 0x8f, 0x66, 0x6a,
	0x6c, 0x6e, 0x52, 0x01, 0x82, 0xa3, 0x99, 0x82, 0x11, 0xe8, 0x89, 0x99, 0x7a, 0x71, 0x3d, 0xaa,
	0x2b, 0x66, 0xea, 0x00, 0xc3, 0x3b, 0xa0, 0x5b, 0x37, 0x74, 0x6a, 0xc7, 0xf9, 0x63, 0x9b, 0x68,
	0x4e, 0x25, 0x73, 0xa3, 0xfc, 0x92, 0xc1, 0xbb, 0xa0, 0x77, 0x46, 0x68, 0x75, 0x22, 0xe3, 0x8e,
	0x29, 0x73, 0x19, 0xbc, 0x07, 0x42, 0xed, 0x89, 0x46, 0xed, 0x1a, 0xd4, 0x9e, 0x4e, 0x0f, 0x70,
	0x1a, 0x83, 0xde, 0xa1, 0x89, 0x96, 0xe5, 0xa5, 0x21, 0xe8, 0xee, 0x2f, 0x6a, 0xf9, 0x7a, 0xfb,
	0x43, 0xc7, 0x36, 0xab, 0xa9, 0x84, 0x3f, 0x3c, 0x10, 0x56, 0x44, 0x1a, 0x3f, 0x6e, 0x5d, 0x6b,
	0xb0, 0x08, 0xc9, 0xff, 0x37, 0x0e, 0xd2, 0x2f, 0xde, 0x45, 0x1b, 0xbd, 0x01, 0x01, 0x57, 0x12,
	0xae, 0x88, 0x57, 0xf3, 0xf3, 0xfc, 0x29, 0x91, 0xc9, 0x60, 0x20, 0xc8, 0x9c, 0x4c, 0x25, 0x7a,
	0x80, 0x8e, 0x1b, 0xbe, 0x40, 0x1a, 0x19, 0x9d, 0x9d, 0x90, 0x86, 0x20, 0x8a, 0xd1, 0x2e, 0x1a,
	0x0d, 0x76, 0x28, 0xcb, 0x29, 0x7e, 0xb2, 0x07, 0x56, 0xb8, 0x92, 0xb9, 0xb6, 0x06, 0xae, 0x99,
	0xf6, 0x23, 0xd3, 0x98, 0x6c, 0x2e, 0x21, 0x88, 0x99, 0x72, 0x00, 0x4e, 0xe2, 0xee, 0x15, 0xc6,
	0xe7, 0x36, 0xba, 0x0f, 0x82, 0x8a, 0x48, 0x08, 0xd3, 0xe2, 0xf4, 0xa1, 0x59, 0x8b, 0xc2, 0xd1,
	0x4f, 0xe1, 0x57, 0x1f, 0x84, 0xc2, 0x69, 0xb9, 0x49, 0x3d, 0xd9, 0xb8, 0x4e, 0x8d, 0x05, 0xe9,
	0x5b, 0xff, 0xa2, 0x8d, 0xfe, 0x78, 0xe9, 0x73, 0xab, 0x66, 0xd5, 0xd0, 0xd9, 0x3f, 0x27, 0xd3,
	0xe4, 0xd9, 0x80, 0x32, 0x41, 0x1a, 0x89, 0x28, 0x93, 0xdc, 0x8a, 0x19, 0x52, 0x9c, 0xe9, 0x65,
	0xc8, 0xdc, 0xd3, 0x6f, 0xa2, 0xd3, 0xc9, 0x5c, 0x11, 0x81, 0x86, 0xa3, 0x6c, 0x94, 0x8d, 0x36,
	0x07, 0x3b, 0x8a, 0x62, 0x4d, 0x50, 0x17, 0xe9, 0xaf, 0xab, 0x4b, 0xdf, 0x7b, 0xff, 0xc0, 0x37,
	0x37, 0xe0, 0xb5, 0xd2, 0xa1, 0x5d, 0x8e, 0xcc, 0xbc, 0x7f, 0x56, 0xb2, 0xcc, 0xbe, 0x70, 0xe6,
	0xb4, 0x2f, 0x0d, 0x73, 0x03, 0x79, 0xf9, 0x32, 0xb7, 0x7d, 0x26, 0x34, 0xbd, 0x26, 0x2a, 0x99,
	0xf9, 0x58, 0x0c, 0x6b, 0x5a, 0x1a, 0x52, 0x66, 0x5c, 0xff, 0xd8, 0x46, 0xeb, 0x20, 0x50, 0x14,
	0xc3, 0xae, 0xb9, 0x48, 0x56, 0xde, 0xb5, 0x91, 0xf7, 0xad, 0x8d, 0xbc, 0xbd, 0xdb, 0xdf, 0x2f,
	0xfb, 0xde, 0xcf, 0xcb, 0xbe, 0xf7, 0xeb, 0xb2, 0xef, 0x7d, 0xfa, 0xdd, 0xff, 0xaf, 0xec, 0x99,
	0x7f, 0xe5, 0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x65, 0xae, 0x29, 0x83, 0x03, 0x00,
	0x00,
}
