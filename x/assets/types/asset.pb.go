// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: remap/assets/asset.proto

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

type Asset struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AssetId     string `protobuf:"bytes,2,opt,name=assetId,proto3" json:"assetId,omitempty"`
	Version     uint64 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	State       string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	VcAvailable bool   `protobuf:"varint,5,opt,name=vcAvailable,proto3" json:"vcAvailable,omitempty"`
	AssetType   string `protobuf:"bytes,6,opt,name=assetType,proto3" json:"assetType,omitempty"`
	Owner       string `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	Address     string `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	AssetSize   uint64 `protobuf:"varint,9,opt,name=assetSize,proto3" json:"assetSize,omitempty"`
	Properties  string `protobuf:"bytes,10,opt,name=properties,proto3" json:"properties,omitempty"`
	Creator     string `protobuf:"bytes,11,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_c010015e7ba440f0, []int{0}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return m.Size()
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Asset) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *Asset) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Asset) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Asset) GetVcAvailable() bool {
	if m != nil {
		return m.VcAvailable
	}
	return false
}

func (m *Asset) GetAssetType() string {
	if m != nil {
		return m.AssetType
	}
	return ""
}

func (m *Asset) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Asset) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Asset) GetAssetSize() uint64 {
	if m != nil {
		return m.AssetSize
	}
	return 0
}

func (m *Asset) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

func (m *Asset) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Asset)(nil), "remap.assets.Asset")
}

func init() { proto.RegisterFile("remap/assets/asset.proto", fileDescriptor_c010015e7ba440f0) }

var fileDescriptor_c010015e7ba440f0 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe3, 0xd0, 0xb4, 0x8d, 0x8b, 0x18, 0xac, 0x0e, 0x37, 0x20, 0x2b, 0x62, 0xca, 0x14,
	0x06, 0x9e, 0xa0, 0x6c, 0xac, 0x81, 0x89, 0xcd, 0x6d, 0x6e, 0xb0, 0x54, 0x6a, 0xcb, 0xb6, 0x02,
	0xe5, 0x29, 0x78, 0x0d, 0xde, 0x84, 0xb1, 0x23, 0x23, 0x4a, 0x5e, 0x04, 0xe5, 0xdc, 0xb4, 0x9d,
	0xec, 0xef, 0xbf, 0xbb, 0x4f, 0xa7, 0xe3, 0xe0, 0xf0, 0x4d, 0xd9, 0x7b, 0xe5, 0x3d, 0x06, 0x1f,
	0x9f, 0xca, 0x3a, 0x13, 0x8c, 0xb8, 0xa6, 0x4a, 0x15, 0x2b, 0x77, 0xdf, 0x29, 0xcf, 0x56, 0xc3,
	0x57, 0xdc, 0xf0, 0x54, 0x37, 0xc0, 0x0a, 0x56, 0x4e, 0xea, 0x54, 0x37, 0x02, 0xf8, 0x8c, 0x7a,
	0x9e, 0x1a, 0x48, 0x0b, 0x56, 0xe6, 0xf5, 0x88, 0x43, 0xa5, 0x45, 0xe7, 0xb5, 0xd9, 0xc1, 0x15,
	0xb5, 0x8f, 0x28, 0x96, 0x3c, 0xf3, 0x41, 0x05, 0x84, 0x09, 0x4d, 0x44, 0x10, 0x05, 0x5f, 0xb4,
	0x9b, 0x55, 0xab, 0xf4, 0x56, 0xad, 0xb7, 0x08, 0x59, 0xc1, 0xca, 0x79, 0x7d, 0x19, 0x89, 0x5b,
	0x9e, 0x93, 0xfc, 0x65, 0x6f, 0x11, 0xa6, 0x34, 0x7b, 0x0e, 0x06, 0xab, 0x79, 0xdf, 0xa1, 0x83,
	0x59, 0xb4, 0x12, 0xd0, 0x7e, 0x4d, 0xe3, 0xd0, 0x7b, 0x98, 0x1f, 0xf7, 0x8b, 0x78, 0xb2, 0x3d,
	0xeb, 0x4f, 0x84, 0x9c, 0x36, 0x3c, 0x07, 0x42, 0x72, 0x6e, 0x9d, 0xb1, 0xe8, 0x82, 0x46, 0x0f,
	0x9c, 0x46, 0x2f, 0x92, 0xc1, 0xbb, 0x71, 0xa8, 0x82, 0x71, 0xb0, 0x88, 0xde, 0x23, 0x3e, 0x56,
	0x3f, 0x9d, 0x64, 0x87, 0x4e, 0xb2, 0xbf, 0x4e, 0xb2, 0xaf, 0x5e, 0x26, 0x87, 0x5e, 0x26, 0xbf,
	0xbd, 0x4c, 0x5e, 0x97, 0xf1, 0xda, 0x1f, 0xe3, 0xbd, 0xc3, 0xde, 0xa2, 0x5f, 0x4f, 0xe9, 0xe0,
	0x0f, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x25, 0xc5, 0xd4, 0x44, 0x8c, 0x01, 0x00, 0x00,
}

func (m *Asset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Asset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Asset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Properties) > 0 {
		i -= len(m.Properties)
		copy(dAtA[i:], m.Properties)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Properties)))
		i--
		dAtA[i] = 0x52
	}
	if m.AssetSize != 0 {
		i = encodeVarintAsset(dAtA, i, uint64(m.AssetSize))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.AssetType) > 0 {
		i -= len(m.AssetType)
		copy(dAtA[i:], m.AssetType)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.AssetType)))
		i--
		dAtA[i] = 0x32
	}
	if m.VcAvailable {
		i--
		if m.VcAvailable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x22
	}
	if m.Version != 0 {
		i = encodeVarintAsset(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if len(m.AssetId) > 0 {
		i -= len(m.AssetId)
		copy(dAtA[i:], m.AssetId)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.AssetId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAsset(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAsset(dAtA []byte, offset int, v uint64) int {
	offset -= sovAsset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Asset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAsset(uint64(m.Id))
	}
	l = len(m.AssetId)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovAsset(uint64(m.Version))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if m.VcAvailable {
		n += 2
	}
	l = len(m.AssetType)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if m.AssetSize != 0 {
		n += 1 + sovAsset(uint64(m.AssetSize))
	}
	l = len(m.Properties)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	return n
}

func sovAsset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAsset(x uint64) (n int) {
	return sovAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Asset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsset
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
			return fmt.Errorf("proto: Asset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Asset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VcAvailable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
			m.VcAvailable = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetSize", wireType)
			}
			m.AssetSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Properties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAsset
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
func skipAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
				return 0, ErrInvalidLengthAsset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAsset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAsset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAsset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAsset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAsset = fmt.Errorf("proto: unexpected end of group")
)
