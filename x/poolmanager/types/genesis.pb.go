// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/poolmanager/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params holds parameters for the poolmanager module
type Params struct {
	PoolCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=pool_creation_fee,json=poolCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pool_creation_fee" yaml:"pool_creation_fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa099d9fbdf68b35, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetPoolCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PoolCreationFee
	}
	return nil
}

// GenesisState defines the poolmanager module's genesis state.
type GenesisState struct {
	// the next_pool_id
	NextPoolId uint64 `protobuf:"varint,1,opt,name=next_pool_id,json=nextPoolId,proto3" json:"next_pool_id,omitempty"`
	// params is the container of poolmanager parameters.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa099d9fbdf68b35, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetNextPoolId() uint64 {
	if m != nil {
		return m.NextPoolId
	}
	return 0
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.poolmanager.v1beta1.Params")
	proto.RegisterType((*GenesisState)(nil), "osmosis.poolmanager.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("osmosis/poolmanager/v1beta1/genesis.proto", fileDescriptor_aa099d9fbdf68b35)
}

var fileDescriptor_aa099d9fbdf68b35 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x3d, 0x6a, 0xe3, 0x40,
	0x14, 0xd6, 0xec, 0x1a, 0x17, 0xb2, 0x61, 0x59, 0xb1, 0x85, 0xec, 0x05, 0x59, 0x68, 0x1b, 0x6d,
	0xe1, 0x19, 0xec, 0x5d, 0x58, 0xd8, 0x2e, 0x36, 0x24, 0x04, 0x52, 0x38, 0x4e, 0x97, 0x46, 0x8c,
	0xa4, 0xb1, 0x22, 0x22, 0xe9, 0x09, 0xcd, 0xd8, 0xd8, 0xb7, 0x08, 0xa4, 0xcf, 0x01, 0x72, 0x12,
	0x97, 0x2e, 0x53, 0x39, 0xc1, 0xbe, 0x41, 0x4e, 0x10, 0x34, 0x1a, 0x05, 0x27, 0x81, 0x54, 0xd2,
	0x9b, 0xef, 0xe7, 0xbd, 0xef, 0xcd, 0xe8, 0xbf, 0x81, 0xa7, 0xc0, 0x63, 0x4e, 0x72, 0x80, 0x24,
	0xa5, 0x19, 0x8d, 0x58, 0x41, 0x16, 0x03, 0x9f, 0x09, 0x3a, 0x20, 0x11, 0xcb, 0x18, 0x8f, 0x39,
	0xce, 0x0b, 0x10, 0x60, 0xfc, 0x54, 0x54, 0x7c, 0x40, 0xc5, 0x8a, 0xda, 0xfd, 0x11, 0x41, 0x04,
	0x92, 0x47, 0xca, 0xbf, 0x4a, 0xd2, 0xed, 0x44, 0x00, 0x51, 0xc2, 0x88, 0xac, 0xfc, 0xf9, 0x8c,
	0xd0, 0x6c, 0x55, 0x43, 0x81, 0xb4, 0xf3, 0x2a, 0x4d, 0x55, 0x28, 0xc8, 0x7a, 0xaf, 0x0a, 0xe7,
	0x05, 0x15, 0x31, 0x64, 0x35, 0x5e, 0xb1, 0x89, 0x4f, 0x39, 0x7b, 0x9d, 0x35, 0x80, 0x58, 0xe1,
	0xce, 0x1d, 0xd2, 0x9b, 0x13, 0x5a, 0xd0, 0x94, 0x1b, 0xb7, 0x48, 0xff, 0x5e, 0x8e, 0xeb, 0x05,
	0x05, 0x93, 0x16, 0xde, 0x8c, 0x31, 0x13, 0xd9, 0x5f, 0xdd, 0xd6, 0xb0, 0x83, 0x55, 0xd7, 0xd2,
	0xa7, 0x0e, 0x82, 0xc7, 0x10, 0x67, 0xa3, 0xb3, 0xf5, 0xb6, 0xa7, 0x3d, 0x6f, 0x7b, 0xe6, 0x8a,
	0xa6, 0xc9, 0x7f, 0xe7, 0x83, 0x83, 0x73, 0xff, 0xd8, 0x73, 0xa3, 0x58, 0x5c, 0xcd, 0x7d, 0x1c,
	0x40, 0xaa, 0xc6, 0x57, 0x9f, 0x3e, 0x0f, 0xaf, 0x89, 0x58, 0xe5, 0x8c, 0x4b, 0x33, 0x3e, 0xfd,
	0x56, 0xea, 0xc7, 0x4a, 0x7e, 0xcc, 0x98, 0xc3, 0xf5, 0xf6, 0x49, 0xb5, 0xda, 0x0b, 0x41, 0x05,
	0x33, 0x6c, 0xbd, 0x9d, 0xb1, 0xa5, 0xf0, 0x64, 0x9f, 0x38, 0x34, 0x91, 0x8d, 0xdc, 0xc6, 0x54,
	0x2f, 0xcf, 0x26, 0x00, 0xc9, 0x69, 0x68, 0x1c, 0xe9, 0xcd, 0x5c, 0x26, 0x32, 0xbf, 0xd8, 0xc8,
	0x6d, 0x0d, 0x7f, 0xe1, 0x4f, 0x2e, 0x03, 0x57, 0xe1, 0x47, 0x8d, 0x32, 0xc5, 0x54, 0x09, 0x47,
	0xe7, 0xeb, 0x9d, 0x85, 0x36, 0x3b, 0x0b, 0x3d, 0xed, 0x2c, 0x74, 0xb3, 0xb7, 0xb4, 0xcd, 0xde,
	0xd2, 0x1e, 0xf6, 0x96, 0x76, 0xf9, 0xef, 0x20, 0x89, 0xb2, 0xed, 0x27, 0xd4, 0xe7, 0x75, 0x41,
	0x16, 0x83, 0xbf, 0x64, 0xf9, 0xe6, 0x85, 0xc8, 0x78, 0x7e, 0x53, 0xee, 0xfb, 0xcf, 0x4b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x27, 0x68, 0x2d, 0xfc, 0x45, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolCreationFee) > 0 {
		for iNdEx := len(m.PoolCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.NextPoolId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextPoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PoolCreationFee) > 0 {
		for _, e := range m.PoolCreationFee {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NextPoolId != 0 {
		n += 1 + sovGenesis(uint64(m.NextPoolId))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolCreationFee = append(m.PoolCreationFee, types.Coin{})
			if err := m.PoolCreationFee[len(m.PoolCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextPoolId", wireType)
			}
			m.NextPoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextPoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
