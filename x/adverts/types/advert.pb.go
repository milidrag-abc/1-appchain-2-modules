// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: remap/adverts/advert.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Advert struct {
	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AdvertId        string `protobuf:"bytes,2,opt,name=advertId,proto3" json:"advertId,omitempty"`
	State           string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	AssetId         string `protobuf:"bytes,4,opt,name=assetId,proto3" json:"assetId,omitempty"`
	Price           uint64 `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	MinRentalPeriod uint64 `protobuf:"varint,6,opt,name=minRentalPeriod,proto3" json:"minRentalPeriod,omitempty"`
	MaxRentalPeriod uint64 `protobuf:"varint,7,opt,name=maxRentalPeriod,proto3" json:"maxRentalPeriod,omitempty"`
	Conditions      string `protobuf:"bytes,8,opt,name=conditions,proto3" json:"conditions,omitempty"`
	ExpirationDate  uint64 `protobuf:"varint,9,opt,name=expirationDate,proto3" json:"expirationDate,omitempty"`
	Creator         string `protobuf:"bytes,10,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Advert) Reset()         { *m = Advert{} }
func (m *Advert) String() string { return proto.CompactTextString(m) }
func (*Advert) ProtoMessage()    {}
func (*Advert) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1622d600f86955e, []int{0}
}
func (m *Advert) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Advert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Advert.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Advert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Advert.Merge(m, src)
}
func (m *Advert) XXX_Size() int {
	return m.Size()
}
func (m *Advert) XXX_DiscardUnknown() {
	xxx_messageInfo_Advert.DiscardUnknown(m)
}

var xxx_messageInfo_Advert proto.InternalMessageInfo

func (m *Advert) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Advert) GetAdvertId() string {
	if m != nil {
		return m.AdvertId
	}
	return ""
}

func (m *Advert) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Advert) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *Advert) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Advert) GetMinRentalPeriod() uint64 {
	if m != nil {
		return m.MinRentalPeriod
	}
	return 0
}

func (m *Advert) GetMaxRentalPeriod() uint64 {
	if m != nil {
		return m.MaxRentalPeriod
	}
	return 0
}

func (m *Advert) GetConditions() string {
	if m != nil {
		return m.Conditions
	}
	return ""
}

func (m *Advert) GetExpirationDate() uint64 {
	if m != nil {
		return m.ExpirationDate
	}
	return 0
}

func (m *Advert) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Advert)(nil), "remap.adverts.Advert")
}

func init() { proto.RegisterFile("remap/adverts/advert.proto", fileDescriptor_d1622d600f86955e) }

var fileDescriptor_d1622d600f86955e = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xe3, 0xd0, 0xa6, 0xed, 0x93, 0x28, 0x92, 0x05, 0x92, 0xd5, 0xc1, 0xaa, 0x18, 0x50,
	0xa6, 0x76, 0xe0, 0x04, 0x20, 0x16, 0x36, 0x94, 0x91, 0xcd, 0xc4, 0x6f, 0xb0, 0x44, 0xe3, 0xc8,
	0xb6, 0x50, 0xb8, 0x05, 0x47, 0xe1, 0x18, 0x8c, 0x1d, 0x19, 0x51, 0x72, 0x11, 0xe4, 0xe7, 0x50,
	0xd1, 0x4e, 0xf6, 0xf7, 0xff, 0xbf, 0x9f, 0xfc, 0x7e, 0x58, 0x39, 0xdc, 0xa9, 0x76, 0xab, 0xf4,
	0x1b, 0xba, 0xe0, 0xc7, 0x73, 0xd3, 0x3a, 0x1b, 0x2c, 0x3f, 0x27, 0x6f, 0x33, 0x7a, 0xd7, 0x9f,
	0x39, 0x14, 0x77, 0x74, 0xe7, 0x4b, 0xc8, 0x8d, 0x16, 0x6c, 0xcd, 0xca, 0x49, 0x95, 0x1b, 0xcd,
	0x57, 0x30, 0x4f, 0xa9, 0x47, 0x2d, 0xf2, 0x35, 0x2b, 0x17, 0xd5, 0x81, 0xf9, 0x25, 0x4c, 0x7d,
	0x50, 0x01, 0xc5, 0x19, 0x19, 0x09, 0xb8, 0x80, 0x99, 0xf2, 0x1e, 0xe3, 0x83, 0x09, 0xe9, 0x7f,
	0x18, 0xf3, 0xad, 0x33, 0x35, 0x8a, 0x29, 0x8d, 0x4f, 0xc0, 0x4b, 0xb8, 0xd8, 0x99, 0xa6, 0xc2,
	0x26, 0xa8, 0xd7, 0x27, 0x74, 0xc6, 0x6a, 0x51, 0x90, 0x7f, 0x2a, 0x53, 0x52, 0x75, 0x47, 0xc9,
	0xd9, 0x98, 0x3c, 0x96, 0xb9, 0x04, 0xa8, 0x6d, 0xa3, 0x4d, 0x30, 0xb6, 0xf1, 0x62, 0x4e, 0xdf,
	0xf8, 0xa7, 0xf0, 0x1b, 0x58, 0x62, 0xd7, 0x1a, 0xa7, 0x22, 0x3e, 0xc4, 0x15, 0x16, 0x34, 0xe8,
	0x44, 0x8d, 0xbb, 0xd4, 0x0e, 0x55, 0xb0, 0x4e, 0x40, 0xda, 0x65, 0xc4, 0xfb, 0xed, 0x57, 0x2f,
	0xd9, 0xbe, 0x97, 0xec, 0xa7, 0x97, 0xec, 0x63, 0x90, 0xd9, 0x7e, 0x90, 0xd9, 0xf7, 0x20, 0xb3,
	0xe7, 0xab, 0xd4, 0x7b, 0x77, 0x68, 0x3e, 0xbc, 0xb7, 0xe8, 0x5f, 0x0a, 0x6a, 0xfe, 0xf6, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x1a, 0x9e, 0x2f, 0x3a, 0x97, 0x01, 0x00, 0x00,
}

func (m *Advert) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Advert) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Advert) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAdvert(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x52
	}
	if m.ExpirationDate != 0 {
		i = encodeVarintAdvert(dAtA, i, uint64(m.ExpirationDate))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Conditions) > 0 {
		i -= len(m.Conditions)
		copy(dAtA[i:], m.Conditions)
		i = encodeVarintAdvert(dAtA, i, uint64(len(m.Conditions)))
		i--
		dAtA[i] = 0x42
	}
	if m.MaxRentalPeriod != 0 {
		i = encodeVarintAdvert(dAtA, i, uint64(m.MaxRentalPeriod))
		i--
		dAtA[i] = 0x38
	}
	if m.MinRentalPeriod != 0 {
		i = encodeVarintAdvert(dAtA, i, uint64(m.MinRentalPeriod))
		i--
		dAtA[i] = 0x30
	}
	if m.Price != 0 {
		i = encodeVarintAdvert(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x28
	}
	if len(m.AssetId) > 0 {
		i -= len(m.AssetId)
		copy(dAtA[i:], m.AssetId)
		i = encodeVarintAdvert(dAtA, i, uint64(len(m.AssetId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintAdvert(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AdvertId) > 0 {
		i -= len(m.AdvertId)
		copy(dAtA[i:], m.AdvertId)
		i = encodeVarintAdvert(dAtA, i, uint64(len(m.AdvertId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAdvert(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAdvert(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdvert(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Advert) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAdvert(uint64(m.Id))
	}
	l = len(m.AdvertId)
	if l > 0 {
		n += 1 + l + sovAdvert(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovAdvert(uint64(l))
	}
	l = len(m.AssetId)
	if l > 0 {
		n += 1 + l + sovAdvert(uint64(l))
	}
	if m.Price != 0 {
		n += 1 + sovAdvert(uint64(m.Price))
	}
	if m.MinRentalPeriod != 0 {
		n += 1 + sovAdvert(uint64(m.MinRentalPeriod))
	}
	if m.MaxRentalPeriod != 0 {
		n += 1 + sovAdvert(uint64(m.MaxRentalPeriod))
	}
	l = len(m.Conditions)
	if l > 0 {
		n += 1 + l + sovAdvert(uint64(l))
	}
	if m.ExpirationDate != 0 {
		n += 1 + sovAdvert(uint64(m.ExpirationDate))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAdvert(uint64(l))
	}
	return n
}

func sovAdvert(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdvert(x uint64) (n int) {
	return sovAdvert(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Advert) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdvert
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
			return fmt.Errorf("proto: Advert: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Advert: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdvert
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdvertId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdvert
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
				return ErrInvalidLengthAdvert
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdvert
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdvertId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdvert
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
				return ErrInvalidLengthAdvert
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdvert
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdvert
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
				return ErrInvalidLengthAdvert
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdvert
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdvert
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinRentalPeriod", wireType)
			}
			m.MinRentalPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdvert
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinRentalPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRentalPeriod", wireType)
			}
			m.MaxRentalPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdvert
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRentalPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdvert
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
				return ErrInvalidLengthAdvert
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdvert
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Conditions = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationDate", wireType)
			}
			m.ExpirationDate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdvert
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationDate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdvert
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
				return ErrInvalidLengthAdvert
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdvert
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdvert(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdvert
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
func skipAdvert(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdvert
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
					return 0, ErrIntOverflowAdvert
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
					return 0, ErrIntOverflowAdvert
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
				return 0, ErrInvalidLengthAdvert
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdvert
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdvert
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdvert        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdvert          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdvert = fmt.Errorf("proto: unexpected end of group")
)
