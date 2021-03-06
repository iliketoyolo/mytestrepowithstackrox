// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/compliance_management.proto

package storage

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

// Next Available Tag: 6
type ComplianceRunSchedule struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClusterId            string   `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	StandardId           string   `protobuf:"bytes,3,opt,name=standard_id,json=standardId,proto3" json:"standard_id,omitempty"`
	Suspended            bool     `protobuf:"varint,4,opt,name=suspended,proto3" json:"suspended,omitempty"`
	CrontabSpec          string   `protobuf:"bytes,5,opt,name=crontab_spec,json=crontabSpec,proto3" json:"crontab_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComplianceRunSchedule) Reset()         { *m = ComplianceRunSchedule{} }
func (m *ComplianceRunSchedule) String() string { return proto.CompactTextString(m) }
func (*ComplianceRunSchedule) ProtoMessage()    {}
func (*ComplianceRunSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_1838cc5cee009dca, []int{0}
}
func (m *ComplianceRunSchedule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ComplianceRunSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ComplianceRunSchedule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ComplianceRunSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplianceRunSchedule.Merge(m, src)
}
func (m *ComplianceRunSchedule) XXX_Size() int {
	return m.Size()
}
func (m *ComplianceRunSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplianceRunSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_ComplianceRunSchedule proto.InternalMessageInfo

func (m *ComplianceRunSchedule) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ComplianceRunSchedule) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ComplianceRunSchedule) GetStandardId() string {
	if m != nil {
		return m.StandardId
	}
	return ""
}

func (m *ComplianceRunSchedule) GetSuspended() bool {
	if m != nil {
		return m.Suspended
	}
	return false
}

func (m *ComplianceRunSchedule) GetCrontabSpec() string {
	if m != nil {
		return m.CrontabSpec
	}
	return ""
}

func (m *ComplianceRunSchedule) MessageClone() proto.Message {
	return m.Clone()
}
func (m *ComplianceRunSchedule) Clone() *ComplianceRunSchedule {
	if m == nil {
		return nil
	}
	cloned := new(ComplianceRunSchedule)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterType((*ComplianceRunSchedule)(nil), "storage.ComplianceRunSchedule")
}

func init() {
	proto.RegisterFile("storage/compliance_management.proto", fileDescriptor_1838cc5cee009dca)
}

var fileDescriptor_1838cc5cee009dca = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xcf, 0x3d, 0x4e, 0xc4, 0x30,
	0x10, 0x86, 0x61, 0x1c, 0x7e, 0x33, 0x8b, 0x28, 0x2c, 0x21, 0x05, 0x09, 0xc2, 0x02, 0xcd, 0x56,
	0xa1, 0x80, 0x13, 0x40, 0xb5, 0x6d, 0xb6, 0xa3, 0x89, 0xbc, 0x33, 0xa3, 0xc5, 0x22, 0xb1, 0x2d,
	0xdb, 0x91, 0x38, 0x0a, 0x47, 0xe0, 0x28, 0x94, 0x1c, 0x01, 0x85, 0x8b, 0x20, 0x82, 0xc3, 0xb6,
	0xcf, 0x7c, 0x53, 0xbc, 0x70, 0x13, 0xa2, 0xf5, 0x6a, 0xc3, 0xb7, 0x68, 0x3b, 0xd7, 0x6a, 0x65,
	0x90, 0x9b, 0x4e, 0x19, 0xb5, 0xe1, 0x8e, 0x4d, 0xac, 0x9c, 0xb7, 0xd1, 0xca, 0xc3, 0x34, 0xba,
	0x7e, 0x17, 0x70, 0xfa, 0xf8, 0x3f, 0xac, 0x7b, 0xb3, 0xc2, 0x67, 0xa6, 0xbe, 0x65, 0x79, 0x02,
	0x99, 0xa6, 0x42, 0xcc, 0xc5, 0x22, 0xaf, 0x33, 0x4d, 0xf2, 0x02, 0x00, 0xdb, 0x3e, 0x44, 0xf6,
	0x8d, 0xa6, 0x22, 0x1b, 0x3d, 0x4f, 0xb2, 0x24, 0x79, 0x09, 0xb3, 0x10, 0x95, 0x21, 0xe5, 0xe9,
	0xf7, 0xbe, 0x3b, 0xde, 0x61, 0xa2, 0x25, 0xc9, 0x73, 0xc8, 0x43, 0x1f, 0x1c, 0x1b, 0x62, 0x2a,
	0xf6, 0xe6, 0x62, 0x71, 0x54, 0x6f, 0x41, 0x5e, 0xc1, 0x31, 0x7a, 0x6b, 0xa2, 0x5a, 0x37, 0xc1,
	0x31, 0x16, 0xfb, 0xe3, 0xff, 0x2c, 0xd9, 0xca, 0x31, 0x3e, 0xdc, 0x7f, 0x0c, 0xa5, 0xf8, 0x1c,
	0x4a, 0xf1, 0x35, 0x94, 0xe2, 0xed, 0xbb, 0xdc, 0x81, 0x33, 0x6d, 0xab, 0x10, 0x15, 0xbe, 0x78,
	0xfb, 0xfa, 0x17, 0x56, 0xa5, 0xae, 0xa7, 0x29, 0x70, 0x7d, 0x30, 0xfa, 0xdd, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x77, 0x50, 0xa8, 0x5d, 0x17, 0x01, 0x00, 0x00,
}

func (m *ComplianceRunSchedule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ComplianceRunSchedule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ComplianceRunSchedule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.CrontabSpec) > 0 {
		i -= len(m.CrontabSpec)
		copy(dAtA[i:], m.CrontabSpec)
		i = encodeVarintComplianceManagement(dAtA, i, uint64(len(m.CrontabSpec)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Suspended {
		i--
		if m.Suspended {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.StandardId) > 0 {
		i -= len(m.StandardId)
		copy(dAtA[i:], m.StandardId)
		i = encodeVarintComplianceManagement(dAtA, i, uint64(len(m.StandardId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClusterId) > 0 {
		i -= len(m.ClusterId)
		copy(dAtA[i:], m.ClusterId)
		i = encodeVarintComplianceManagement(dAtA, i, uint64(len(m.ClusterId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintComplianceManagement(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintComplianceManagement(dAtA []byte, offset int, v uint64) int {
	offset -= sovComplianceManagement(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ComplianceRunSchedule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovComplianceManagement(uint64(l))
	}
	l = len(m.ClusterId)
	if l > 0 {
		n += 1 + l + sovComplianceManagement(uint64(l))
	}
	l = len(m.StandardId)
	if l > 0 {
		n += 1 + l + sovComplianceManagement(uint64(l))
	}
	if m.Suspended {
		n += 2
	}
	l = len(m.CrontabSpec)
	if l > 0 {
		n += 1 + l + sovComplianceManagement(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovComplianceManagement(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComplianceManagement(x uint64) (n int) {
	return sovComplianceManagement(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ComplianceRunSchedule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComplianceManagement
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
			return fmt.Errorf("proto: ComplianceRunSchedule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ComplianceRunSchedule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceManagement
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
				return ErrInvalidLengthComplianceManagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceManagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceManagement
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
				return ErrInvalidLengthComplianceManagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceManagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StandardId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceManagement
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
				return ErrInvalidLengthComplianceManagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceManagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StandardId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suspended", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceManagement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Suspended = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrontabSpec", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComplianceManagement
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
				return ErrInvalidLengthComplianceManagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComplianceManagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CrontabSpec = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComplianceManagement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComplianceManagement
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
func skipComplianceManagement(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComplianceManagement
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
					return 0, ErrIntOverflowComplianceManagement
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
					return 0, ErrIntOverflowComplianceManagement
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
				return 0, ErrInvalidLengthComplianceManagement
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComplianceManagement
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComplianceManagement
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComplianceManagement        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComplianceManagement          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComplianceManagement = fmt.Errorf("proto: unexpected end of group")
)
