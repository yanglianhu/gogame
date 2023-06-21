// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sshead.proto

package ss

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SSHead struct {
	CmdId                uint32   `protobuf:"varint,1,opt,name=CmdId,proto3" json:"CmdId,omitempty"`
	SeqId                uint32   `protobuf:"varint,2,opt,name=SeqId,proto3" json:"SeqId,omitempty"`
	RetCode              uint32   `protobuf:"varint,3,opt,name=RetCode,proto3" json:"RetCode,omitempty"`
	Src                  uint64   `protobuf:"varint,4,opt,name=Src,proto3" json:"Src,omitempty"`
	Dst                  uint64   `protobuf:"varint,5,opt,name=Dst,proto3" json:"Dst,omitempty"`
	Uid                  uint64   `protobuf:"varint,6,opt,name=Uid,proto3" json:"Uid,omitempty"`
	SubCmdId             uint32   `protobuf:"varint,7,opt,name=SubCmdId,proto3" json:"SubCmdId,omitempty"`
	Flag                 uint32   `protobuf:"varint,8,opt,name=Flag,proto3" json:"Flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSHead) Reset()         { *m = SSHead{} }
func (m *SSHead) String() string { return proto.CompactTextString(m) }
func (*SSHead) ProtoMessage()    {}
func (*SSHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aa821ff079f83e9, []int{0}
}
func (m *SSHead) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SSHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SSHead.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SSHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSHead.Merge(m, src)
}
func (m *SSHead) XXX_Size() int {
	return m.Size()
}
func (m *SSHead) XXX_DiscardUnknown() {
	xxx_messageInfo_SSHead.DiscardUnknown(m)
}

var xxx_messageInfo_SSHead proto.InternalMessageInfo

func (m *SSHead) GetCmdId() uint32 {
	if m != nil {
		return m.CmdId
	}
	return 0
}

func (m *SSHead) GetSeqId() uint32 {
	if m != nil {
		return m.SeqId
	}
	return 0
}

func (m *SSHead) GetRetCode() uint32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *SSHead) GetSrc() uint64 {
	if m != nil {
		return m.Src
	}
	return 0
}

func (m *SSHead) GetDst() uint64 {
	if m != nil {
		return m.Dst
	}
	return 0
}

func (m *SSHead) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SSHead) GetSubCmdId() uint32 {
	if m != nil {
		return m.SubCmdId
	}
	return 0
}

func (m *SSHead) GetFlag() uint32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func init() {
	proto.RegisterType((*SSHead)(nil), "ss.SSHead")
}

func init() { proto.RegisterFile("sshead.proto", fileDescriptor_4aa821ff079f83e9) }

var fileDescriptor_4aa821ff079f83e9 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0xce, 0x48,
	0x4d, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x2e, 0x56, 0xda, 0xc2, 0xc8,
	0xc5, 0x16, 0x1c, 0xec, 0x91, 0x9a, 0x98, 0x22, 0x24, 0xc2, 0xc5, 0xea, 0x9c, 0x9b, 0xe2, 0x99,
	0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1b, 0x04, 0xe1, 0x80, 0x44, 0x83, 0x53, 0x0b, 0x3d, 0x53,
	0x24, 0x98, 0x20, 0xa2, 0x60, 0x8e, 0x90, 0x04, 0x17, 0x7b, 0x50, 0x6a, 0x89, 0x73, 0x7e, 0x4a,
	0xaa, 0x04, 0x33, 0x58, 0x1c, 0xc6, 0x15, 0x12, 0xe0, 0x62, 0x0e, 0x2e, 0x4a, 0x96, 0x60, 0x51,
	0x60, 0xd4, 0x60, 0x09, 0x02, 0x31, 0x41, 0x22, 0x2e, 0xc5, 0x25, 0x12, 0xac, 0x10, 0x11, 0x97,
	0xe2, 0x12, 0x90, 0x48, 0x68, 0x66, 0x8a, 0x04, 0x1b, 0x44, 0x24, 0x34, 0x33, 0x45, 0x48, 0x8a,
	0x8b, 0x23, 0xb8, 0x34, 0x09, 0x62, 0x3d, 0x3b, 0xd8, 0x40, 0x38, 0x5f, 0x48, 0x88, 0x8b, 0xc5,
	0x2d, 0x27, 0x31, 0x5d, 0x82, 0x03, 0x2c, 0x0e, 0x66, 0x3b, 0x89, 0x9f, 0x78, 0x24, 0xc7, 0x78,
	0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0xb1, 0xea, 0xe9,
	0x5b, 0x17, 0x17, 0x27, 0xb1, 0x81, 0xbd, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x74,
	0x33, 0xd4, 0xea, 0x00, 0x00, 0x00,
}

func (m *SSHead) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SSHead) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SSHead) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Flag != 0 {
		i = encodeVarintSshead(dAtA, i, uint64(m.Flag))
		i--
		dAtA[i] = 0x40
	}
	if m.SubCmdId != 0 {
		i = encodeVarintSshead(dAtA, i, uint64(m.SubCmdId))
		i--
		dAtA[i] = 0x38
	}
	if m.Uid != 0 {
		i = encodeVarintSshead(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x30
	}
	if m.Dst != 0 {
		i = encodeVarintSshead(dAtA, i, uint64(m.Dst))
		i--
		dAtA[i] = 0x28
	}
	if m.Src != 0 {
		i = encodeVarintSshead(dAtA, i, uint64(m.Src))
		i--
		dAtA[i] = 0x20
	}
	if m.RetCode != 0 {
		i = encodeVarintSshead(dAtA, i, uint64(m.RetCode))
		i--
		dAtA[i] = 0x18
	}
	if m.SeqId != 0 {
		i = encodeVarintSshead(dAtA, i, uint64(m.SeqId))
		i--
		dAtA[i] = 0x10
	}
	if m.CmdId != 0 {
		i = encodeVarintSshead(dAtA, i, uint64(m.CmdId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSshead(dAtA []byte, offset int, v uint64) int {
	offset -= sovSshead(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SSHead) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CmdId != 0 {
		n += 1 + sovSshead(uint64(m.CmdId))
	}
	if m.SeqId != 0 {
		n += 1 + sovSshead(uint64(m.SeqId))
	}
	if m.RetCode != 0 {
		n += 1 + sovSshead(uint64(m.RetCode))
	}
	if m.Src != 0 {
		n += 1 + sovSshead(uint64(m.Src))
	}
	if m.Dst != 0 {
		n += 1 + sovSshead(uint64(m.Dst))
	}
	if m.Uid != 0 {
		n += 1 + sovSshead(uint64(m.Uid))
	}
	if m.SubCmdId != 0 {
		n += 1 + sovSshead(uint64(m.SubCmdId))
	}
	if m.Flag != 0 {
		n += 1 + sovSshead(uint64(m.Flag))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSshead(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSshead(x uint64) (n int) {
	return sovSshead(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SSHead) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSshead
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
			return fmt.Errorf("proto: SSHead: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SSHead: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CmdId", wireType)
			}
			m.CmdId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSshead
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CmdId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeqId", wireType)
			}
			m.SeqId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSshead
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SeqId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetCode", wireType)
			}
			m.RetCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSshead
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RetCode |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Src", wireType)
			}
			m.Src = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSshead
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Src |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dst", wireType)
			}
			m.Dst = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSshead
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dst |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSshead
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubCmdId", wireType)
			}
			m.SubCmdId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSshead
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubCmdId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag", wireType)
			}
			m.Flag = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSshead
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Flag |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSshead(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSshead
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
func skipSshead(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSshead
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
					return 0, ErrIntOverflowSshead
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
					return 0, ErrIntOverflowSshead
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
				return 0, ErrInvalidLengthSshead
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSshead
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSshead
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSshead        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSshead          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSshead = fmt.Errorf("proto: unexpected end of group")
)
