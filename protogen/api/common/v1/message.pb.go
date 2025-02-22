// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/common/v1/message.proto

package common

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// (-- api-linter: core::0123::resource-annotation=disabled --)
type Region struct {
	// E.g., aws, gcp, azure.
	CloudProvider string `protobuf:"bytes,1,opt,name=cloud_provider,json=cloudProvider,proto3" json:"cloud_provider,omitempty"`
	// Cloud-specific region name. E.g., us-west-2 for AWS and europe-west1 for GCP.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The flag indicates if the region supports global namespace.
	SupportGlobalNamespace bool `protobuf:"varint,3,opt,name=support_global_namespace,json=supportGlobalNamespace,proto3" json:"support_global_namespace,omitempty"` // Deprecated: Do not use.
	// The allow list of connection between the current region with a target region.
	ConnectableRegions []*RegionID `protobuf:"bytes,4,rep,name=connectable_regions,json=connectableRegions,proto3" json:"connectable_regions,omitempty"`
}

func (m *Region) Reset()      { *m = Region{} }
func (*Region) ProtoMessage() {}
func (*Region) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4c2c740e1454e5, []int{0}
}
func (m *Region) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Region) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Region.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Region) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Region.Merge(m, src)
}
func (m *Region) XXX_Size() int {
	return m.Size()
}
func (m *Region) XXX_DiscardUnknown() {
	xxx_messageInfo_Region.DiscardUnknown(m)
}

var xxx_messageInfo_Region proto.InternalMessageInfo

func (m *Region) GetCloudProvider() string {
	if m != nil {
		return m.CloudProvider
	}
	return ""
}

func (m *Region) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Deprecated: Do not use.
func (m *Region) GetSupportGlobalNamespace() bool {
	if m != nil {
		return m.SupportGlobalNamespace
	}
	return false
}

func (m *Region) GetConnectableRegions() []*RegionID {
	if m != nil {
		return m.ConnectableRegions
	}
	return nil
}

type RegionID struct {
	// E.g., aws, gcp, azure.
	CloudProvider string `protobuf:"bytes,1,opt,name=cloud_provider,json=cloudProvider,proto3" json:"cloud_provider,omitempty"`
	// Cloud-specific region name. E.g., us-west-2 for AWS and europe-west1 for GCP.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *RegionID) Reset()      { *m = RegionID{} }
func (*RegionID) ProtoMessage() {}
func (*RegionID) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4c2c740e1454e5, []int{1}
}
func (m *RegionID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegionID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegionID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegionID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionID.Merge(m, src)
}
func (m *RegionID) XXX_Size() int {
	return m.Size()
}
func (m *RegionID) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionID.DiscardUnknown(m)
}

var xxx_messageInfo_RegionID proto.InternalMessageInfo

func (m *RegionID) GetCloudProvider() string {
	if m != nil {
		return m.CloudProvider
	}
	return ""
}

func (m *RegionID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Region)(nil), "api.common.v1.Region")
	proto.RegisterType((*RegionID)(nil), "api.common.v1.RegionID")
}

func init() { proto.RegisterFile("api/common/v1/message.proto", fileDescriptor_df4c2c740e1454e5) }

var fileDescriptor_df4c2c740e1454e5 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x31, 0x4b, 0x03, 0x31,
	0x1c, 0xc5, 0xef, 0xdf, 0x96, 0x52, 0x23, 0x75, 0x88, 0xa0, 0x01, 0xe1, 0xcf, 0x51, 0x10, 0x3a,
	0xe5, 0xa8, 0x8e, 0x3a, 0x15, 0x45, 0x5d, 0x44, 0x6e, 0x70, 0x70, 0x39, 0xd2, 0x6b, 0x28, 0x07,
	0x77, 0x97, 0x90, 0x5c, 0x3b, 0xfb, 0x11, 0xfc, 0x18, 0x7e, 0x14, 0xc1, 0xa5, 0x63, 0x47, 0x9b,
	0x5b, 0x1c, 0xfb, 0x11, 0xc4, 0x5c, 0x0b, 0x76, 0x76, 0x7b, 0xfc, 0xde, 0x4b, 0x78, 0x7f, 0x1e,
	0x39, 0x13, 0x3a, 0x8b, 0x52, 0x55, 0x14, 0xaa, 0x8c, 0x16, 0xa3, 0xa8, 0x90, 0xd6, 0x8a, 0x99,
	0xe4, 0xda, 0xa8, 0x4a, 0xd1, 0xbe, 0xd0, 0x19, 0x6f, 0x4c, 0xbe, 0x18, 0x0d, 0x3e, 0x81, 0x74,
	0x63, 0x39, 0xcb, 0x54, 0x49, 0xcf, 0xc9, 0x51, 0x9a, 0xab, 0xf9, 0x34, 0xd1, 0x46, 0x2d, 0xb2,
	0xa9, 0x34, 0x0c, 0x42, 0x18, 0x1e, 0xc4, 0x7d, 0x4f, 0x9f, 0xb6, 0x90, 0x52, 0xd2, 0x29, 0x45,
	0x21, 0x59, 0xcb, 0x9b, 0x5e, 0xd3, 0x6b, 0xc2, 0xec, 0x5c, 0x6b, 0x65, 0xaa, 0x64, 0x96, 0xab,
	0x89, 0xc8, 0x93, 0x5f, 0x6c, 0xb5, 0x48, 0x25, 0x6b, 0x87, 0x30, 0xec, 0x8d, 0x5b, 0x0c, 0xe2,
	0x93, 0x6d, 0xe6, 0xce, 0x47, 0x1e, 0x77, 0x09, 0x7a, 0x4f, 0x8e, 0x53, 0x55, 0x96, 0x32, 0xad,
	0xc4, 0x24, 0x97, 0x89, 0xf1, 0x75, 0x2c, 0xeb, 0x84, 0xed, 0xe1, 0xe1, 0xc5, 0x29, 0xdf, 0x2b,
	0xcc, 0x9b, 0xb2, 0x0f, 0x37, 0x31, 0xfd, 0xf3, 0xa6, 0x81, 0x76, 0x70, 0x4b, 0x7a, 0x3b, 0xff,
	0x1f, 0xe7, 0x8c, 0x9f, 0x97, 0x6b, 0x0c, 0x56, 0x6b, 0x0c, 0x36, 0x6b, 0x84, 0x57, 0x87, 0xf0,
	0xee, 0x10, 0x3e, 0x1c, 0xc2, 0xd2, 0x21, 0x7c, 0x39, 0x84, 0x6f, 0x87, 0xc1, 0xc6, 0x21, 0xbc,
	0xd5, 0x18, 0x2c, 0x6b, 0x0c, 0x56, 0x35, 0x06, 0x2f, 0x61, 0x55, 0x68, 0x93, 0x73, 0xff, 0x7d,
	0xb4, 0x37, 0xc1, 0x55, 0xa3, 0x26, 0x5d, 0x3f, 0xc1, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x53, 0x59, 0x1e, 0x0f, 0xa1, 0x01, 0x00, 0x00,
}

func (this *Region) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Region)
	if !ok {
		that2, ok := that.(Region)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CloudProvider != that1.CloudProvider {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.SupportGlobalNamespace != that1.SupportGlobalNamespace {
		return false
	}
	if len(this.ConnectableRegions) != len(that1.ConnectableRegions) {
		return false
	}
	for i := range this.ConnectableRegions {
		if !this.ConnectableRegions[i].Equal(that1.ConnectableRegions[i]) {
			return false
		}
	}
	return true
}
func (this *RegionID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegionID)
	if !ok {
		that2, ok := that.(RegionID)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CloudProvider != that1.CloudProvider {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *Region) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&common.Region{")
	s = append(s, "CloudProvider: "+fmt.Sprintf("%#v", this.CloudProvider)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "SupportGlobalNamespace: "+fmt.Sprintf("%#v", this.SupportGlobalNamespace)+",\n")
	if this.ConnectableRegions != nil {
		s = append(s, "ConnectableRegions: "+fmt.Sprintf("%#v", this.ConnectableRegions)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RegionID) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&common.RegionID{")
	s = append(s, "CloudProvider: "+fmt.Sprintf("%#v", this.CloudProvider)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Region) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Region) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Region) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConnectableRegions) > 0 {
		for iNdEx := len(m.ConnectableRegions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ConnectableRegions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.SupportGlobalNamespace {
		i--
		if m.SupportGlobalNamespace {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CloudProvider) > 0 {
		i -= len(m.CloudProvider)
		copy(dAtA[i:], m.CloudProvider)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.CloudProvider)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegionID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegionID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegionID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CloudProvider) > 0 {
		i -= len(m.CloudProvider)
		copy(dAtA[i:], m.CloudProvider)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.CloudProvider)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Region) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CloudProvider)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.SupportGlobalNamespace {
		n += 2
	}
	if len(m.ConnectableRegions) > 0 {
		for _, e := range m.ConnectableRegions {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	return n
}

func (m *RegionID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CloudProvider)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Region) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForConnectableRegions := "[]*RegionID{"
	for _, f := range this.ConnectableRegions {
		repeatedStringForConnectableRegions += strings.Replace(f.String(), "RegionID", "RegionID", 1) + ","
	}
	repeatedStringForConnectableRegions += "}"
	s := strings.Join([]string{`&Region{`,
		`CloudProvider:` + fmt.Sprintf("%v", this.CloudProvider) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`SupportGlobalNamespace:` + fmt.Sprintf("%v", this.SupportGlobalNamespace) + `,`,
		`ConnectableRegions:` + repeatedStringForConnectableRegions + `,`,
		`}`,
	}, "")
	return s
}
func (this *RegionID) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RegionID{`,
		`CloudProvider:` + fmt.Sprintf("%v", this.CloudProvider) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Region) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Region: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Region: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CloudProvider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CloudProvider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupportGlobalNamespace", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
			m.SupportGlobalNamespace = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectableRegions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectableRegions = append(m.ConnectableRegions, &RegionID{})
			if err := m.ConnectableRegions[len(m.ConnectableRegions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *RegionID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: RegionID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegionID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CloudProvider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CloudProvider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
