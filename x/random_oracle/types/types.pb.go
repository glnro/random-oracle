// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: glnro/random_oracle/v1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type RandomResult struct {
	Round      uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	Randomness string `protobuf:"bytes,2,opt,name=randomness,proto3" json:"randomness,omitempty"`
	Signature  string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *RandomResult) Reset()         { *m = RandomResult{} }
func (m *RandomResult) String() string { return proto.CompactTextString(m) }
func (*RandomResult) ProtoMessage()    {}
func (*RandomResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef189d0d74ef6a89, []int{0}
}
func (m *RandomResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RandomResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RandomResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RandomResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomResult.Merge(m, src)
}
func (m *RandomResult) XXX_Size() int {
	return m.Size()
}
func (m *RandomResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomResult.DiscardUnknown(m)
}

var xxx_messageInfo_RandomResult proto.InternalMessageInfo

func (m *RandomResult) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *RandomResult) GetRandomness() string {
	if m != nil {
		return m.Randomness
	}
	return ""
}

func (m *RandomResult) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*RandomResult)(nil), "glnro.random_oracle.v1.RandomResult")
}

func init() {
	proto.RegisterFile("glnro/random_oracle/v1/types.proto", fileDescriptor_ef189d0d74ef6a89)
}

var fileDescriptor_ef189d0d74ef6a89 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x6a, 0xc3, 0x30,
	0x10, 0x87, 0xad, 0xfe, 0x83, 0x88, 0x2e, 0x35, 0xa1, 0xb8, 0xa1, 0x88, 0x10, 0x3a, 0x64, 0xa9,
	0x0f, 0xd3, 0xbd, 0x43, 0x1f, 0xc1, 0x63, 0x97, 0x20, 0x39, 0x42, 0x35, 0xd8, 0xba, 0xa0, 0x93,
	0x02, 0x7d, 0x8b, 0x3e, 0x56, 0xc7, 0x8c, 0x1d, 0x8b, 0xfd, 0x22, 0x25, 0x92, 0xa1, 0xc9, 0x22,
	0x74, 0xdf, 0x7d, 0xc7, 0x1d, 0x3f, 0xbe, 0x32, 0x9d, 0x75, 0x08, 0x4e, 0xda, 0x2d, 0xf6, 0x1b,
	0x74, 0xb2, 0xe9, 0x34, 0xec, 0x2b, 0xf0, 0x9f, 0x3b, 0x4d, 0xe5, 0xce, 0xa1, 0xc7, 0xfc, 0x3e,
	0x3a, 0xe5, 0x99, 0x53, 0xee, 0xab, 0xc5, 0x43, 0x83, 0xd4, 0x23, 0x6d, 0xa2, 0x05, 0xa9, 0x48,
	0x23, 0x8b, 0xb9, 0x41, 0x83, 0x89, 0x1f, 0x7f, 0x13, 0x15, 0xc9, 0x01, 0x25, 0xe9, 0xb8, 0x44,
	0x69, 0x2f, 0x2b, 0x68, 0xb0, 0xb5, 0x53, 0xff, 0x4e, 0xf6, 0xad, 0x45, 0x88, 0x6f, 0x42, 0x2b,
	0xc5, 0x6f, 0xeb, 0xb8, 0xb7, 0xd6, 0x14, 0x3a, 0x9f, 0xcf, 0xf9, 0xb5, 0xc3, 0x60, 0xb7, 0x05,
	0x5b, 0xb2, 0xf5, 0x55, 0x9d, 0x8a, 0x5c, 0x70, 0x9e, 0xae, 0xb3, 0x9a, 0xa8, 0xb8, 0x58, 0xb2,
	0xf5, 0xac, 0x3e, 0x21, 0xf9, 0x23, 0x9f, 0x51, 0x6b, 0xac, 0xf4, 0xc1, 0xe9, 0xe2, 0x32, 0xb6,
	0xff, 0xc1, 0xdb, 0xeb, 0xf7, 0x20, 0xd8, 0x61, 0x10, 0xec, 0x77, 0x10, 0xec, 0x6b, 0x14, 0xd9,
	0x61, 0x14, 0xd9, 0xcf, 0x28, 0xb2, 0xf7, 0x27, 0xd3, 0xfa, 0x8f, 0xa0, 0xca, 0x06, 0x7b, 0x38,
	0x0d, 0xea, 0x79, 0x0a, 0x2a, 0xa6, 0xa4, 0x6e, 0xe2, 0xa9, 0x2f, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x54, 0x3d, 0xf7, 0xaf, 0x4c, 0x01, 0x00, 0x00,
}

func (m *RandomResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RandomResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RandomResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Randomness) > 0 {
		i -= len(m.Randomness)
		copy(dAtA[i:], m.Randomness)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Randomness)))
		i--
		dAtA[i] = 0x12
	}
	if m.Round != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RandomResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovTypes(uint64(m.Round))
	}
	l = len(m.Randomness)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RandomResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RandomResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RandomResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Randomness", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Randomness = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)