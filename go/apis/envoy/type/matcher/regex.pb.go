// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/matcher/regex.proto

package envoy_type_matcher

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A regex matcher designed for safety when used with untrusted input.
type RegexMatcher struct {
	// Types that are valid to be assigned to EngineType:
	//	*RegexMatcher_GoogleRe2
	EngineType isRegexMatcher_EngineType `protobuf_oneof:"engine_type"`
	// The regex match string. The string must be supported by the configured engine.
	Regex                string   `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegexMatcher) Reset()         { *m = RegexMatcher{} }
func (m *RegexMatcher) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher) ProtoMessage()    {}
func (*RegexMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ba289439e2572f3, []int{0}
}
func (m *RegexMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegexMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegexMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegexMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher.Merge(m, src)
}
func (m *RegexMatcher) XXX_Size() int {
	return m.Size()
}
func (m *RegexMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher proto.InternalMessageInfo

type isRegexMatcher_EngineType interface {
	isRegexMatcher_EngineType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type RegexMatcher_GoogleRe2 struct {
	GoogleRe2 *RegexMatcher_GoogleRE2 `protobuf:"bytes,1,opt,name=google_re2,json=googleRe2,proto3,oneof"`
}

func (*RegexMatcher_GoogleRe2) isRegexMatcher_EngineType() {}

func (m *RegexMatcher) GetEngineType() isRegexMatcher_EngineType {
	if m != nil {
		return m.EngineType
	}
	return nil
}

func (m *RegexMatcher) GetGoogleRe2() *RegexMatcher_GoogleRE2 {
	if x, ok := m.GetEngineType().(*RegexMatcher_GoogleRe2); ok {
		return x.GoogleRe2
	}
	return nil
}

func (m *RegexMatcher) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RegexMatcher) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RegexMatcher_OneofMarshaler, _RegexMatcher_OneofUnmarshaler, _RegexMatcher_OneofSizer, []interface{}{
		(*RegexMatcher_GoogleRe2)(nil),
	}
}

func _RegexMatcher_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RegexMatcher)
	// engine_type
	switch x := m.EngineType.(type) {
	case *RegexMatcher_GoogleRe2:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GoogleRe2); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RegexMatcher.EngineType has unexpected type %T", x)
	}
	return nil
}

func _RegexMatcher_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RegexMatcher)
	switch tag {
	case 1: // engine_type.google_re2
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RegexMatcher_GoogleRE2)
		err := b.DecodeMessage(msg)
		m.EngineType = &RegexMatcher_GoogleRe2{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RegexMatcher_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RegexMatcher)
	// engine_type
	switch x := m.EngineType.(type) {
	case *RegexMatcher_GoogleRe2:
		s := proto.Size(x.GoogleRe2)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Google's `RE2 <https://github.com/google/re2>`_ regex engine. The regex string must adhere to
// the documented `syntax <https://github.com/google/re2/wiki/Syntax>`_. The engine is designed
// to complete execution in linear time as well as limit the amount of memory used.
type RegexMatcher_GoogleRE2 struct {
	// This field controls the RE2 "program size" which is a rough estimate of how complex a
	// compiled regex is to evaluate. A regex that has a program size greater than the configured
	// value will fail to compile. In this case, the configured max program size can be increased
	// or the regex can be simplified. If not specified, the default is 100.
	MaxProgramSize       *types.UInt32Value `protobuf:"bytes,1,opt,name=max_program_size,json=maxProgramSize,proto3" json:"max_program_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RegexMatcher_GoogleRE2) Reset()         { *m = RegexMatcher_GoogleRE2{} }
func (m *RegexMatcher_GoogleRE2) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher_GoogleRE2) ProtoMessage()    {}
func (*RegexMatcher_GoogleRE2) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ba289439e2572f3, []int{0, 0}
}
func (m *RegexMatcher_GoogleRE2) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegexMatcher_GoogleRE2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegexMatcher_GoogleRE2.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegexMatcher_GoogleRE2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher_GoogleRE2.Merge(m, src)
}
func (m *RegexMatcher_GoogleRE2) XXX_Size() int {
	return m.Size()
}
func (m *RegexMatcher_GoogleRE2) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher_GoogleRE2.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher_GoogleRE2 proto.InternalMessageInfo

func (m *RegexMatcher_GoogleRE2) GetMaxProgramSize() *types.UInt32Value {
	if m != nil {
		return m.MaxProgramSize
	}
	return nil
}

func init() {
	proto.RegisterType((*RegexMatcher)(nil), "envoy.type.matcher.RegexMatcher")
	proto.RegisterType((*RegexMatcher_GoogleRE2)(nil), "envoy.type.matcher.RegexMatcher.GoogleRE2")
}

func init() { proto.RegisterFile("envoy/type/matcher/regex.proto", fileDescriptor_0ba289439e2572f3) }

var fileDescriptor_0ba289439e2572f3 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x9d, 0xd6, 0x0a, 0xd9, 0x8a, 0x94, 0x45, 0xb0, 0x14, 0x89, 0xc1, 0x53, 0xf1, 0xb0,
	0x0b, 0xe9, 0xd5, 0x53, 0xc0, 0x7f, 0x07, 0xa1, 0x6c, 0x51, 0x8f, 0x61, 0xab, 0x63, 0x0c, 0x34,
	0xd9, 0x65, 0x9b, 0xd6, 0xb4, 0x8f, 0xe0, 0x13, 0x89, 0xa7, 0x1e, 0x3d, 0xfa, 0x08, 0xd2, 0x5b,
	0xcf, 0xbe, 0x80, 0x64, 0xb7, 0x15, 0x41, 0x6f, 0xcb, 0x7e, 0xf3, 0xfd, 0x66, 0x7e, 0xc4, 0xc7,
	0x7c, 0xaa, 0x66, 0xbc, 0x98, 0x69, 0xe4, 0x99, 0x2c, 0xee, 0x9f, 0xd0, 0x70, 0x83, 0x09, 0x96,
	0x4c, 0x1b, 0x55, 0x28, 0x4a, 0x6d, 0xce, 0xaa, 0x9c, 0xad, 0xf3, 0x8e, 0x9f, 0x28, 0x95, 0x8c,
	0x90, 0xdb, 0x89, 0xe1, 0xe4, 0x91, 0x3f, 0x1b, 0xa9, 0x35, 0x9a, 0xb1, 0xeb, 0x74, 0x0e, 0xa6,
	0x72, 0x94, 0x3e, 0xc8, 0x02, 0xf9, 0xe6, 0xe1, 0x82, 0xe3, 0x2f, 0x20, 0xbb, 0xa2, 0x82, 0x5f,
	0x3b, 0x12, 0xbd, 0x23, 0xc4, 0xb1, 0x62, 0x83, 0x61, 0x1b, 0x02, 0xe8, 0x36, 0xc3, 0x13, 0xf6,
	0x77, 0x25, 0xfb, 0xdd, 0x62, 0x17, 0xb6, 0x22, 0xce, 0xc2, 0x88, 0xbc, 0xad, 0x16, 0xf5, 0xc6,
	0x0b, 0xd4, 0x5a, 0x70, 0xb9, 0x25, 0x3c, 0xc7, 0x12, 0x18, 0xd2, 0x23, 0xd2, 0xb0, 0x16, 0xed,
	0x5a, 0x00, 0x5d, 0x2f, 0xf2, 0xaa, 0xb9, 0x6d, 0x53, 0x0b, 0x40, 0xb8, 0xff, 0xce, 0x80, 0x78,
	0x3f, 0x18, 0x7a, 0x4e, 0x5a, 0x99, 0x2c, 0x63, 0x6d, 0x54, 0x62, 0x64, 0x16, 0x8f, 0xd3, 0x39,
	0xae, 0x8f, 0x39, 0x64, 0x8e, 0xc9, 0x36, 0xae, 0xec, 0xe6, 0x2a, 0x2f, 0x7a, 0xe1, 0xad, 0x1c,
	0x4d, 0x50, 0xec, 0x65, 0xb2, 0xec, 0xbb, 0xd2, 0x20, 0x9d, 0x63, 0xb4, 0x4f, 0x9a, 0x98, 0x27,
	0x69, 0x8e, 0x71, 0x75, 0x3c, 0x6d, 0xbc, 0xae, 0x16, 0x75, 0x88, 0x4e, 0xdf, 0x97, 0x3e, 0x7c,
	0x2c, 0x7d, 0xf8, 0x5c, 0xfa, 0x40, 0x82, 0x54, 0x39, 0x41, 0x6d, 0x54, 0x39, 0xfb, 0xc7, 0x35,
	0x22, 0x56, 0xb6, 0x5f, 0x2d, 0xec, 0xc3, 0x70, 0xc7, 0x6e, 0xee, 0x7d, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xfe, 0xdf, 0x4f, 0x5d, 0xa9, 0x01, 0x00, 0x00,
}

func (m *RegexMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegexMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegexMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Regex) > 0 {
		i -= len(m.Regex)
		copy(dAtA[i:], m.Regex)
		i = encodeVarintRegex(dAtA, i, uint64(len(m.Regex)))
		i--
		dAtA[i] = 0x12
	}
	if m.EngineType != nil {
		{
			size := m.EngineType.Size()
			i -= size
			if _, err := m.EngineType.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *RegexMatcher_GoogleRe2) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *RegexMatcher_GoogleRe2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.GoogleRe2 != nil {
		{
			size, err := m.GoogleRe2.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRegex(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *RegexMatcher_GoogleRE2) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegexMatcher_GoogleRE2) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegexMatcher_GoogleRE2) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MaxProgramSize != nil {
		{
			size, err := m.MaxProgramSize.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRegex(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRegex(dAtA []byte, offset int, v uint64) int {
	offset -= sovRegex(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegexMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EngineType != nil {
		n += m.EngineType.Size()
	}
	l = len(m.Regex)
	if l > 0 {
		n += 1 + l + sovRegex(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegexMatcher_GoogleRe2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GoogleRe2 != nil {
		l = m.GoogleRe2.Size()
		n += 1 + l + sovRegex(uint64(l))
	}
	return n
}
func (m *RegexMatcher_GoogleRE2) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxProgramSize != nil {
		l = m.MaxProgramSize.Size()
		n += 1 + l + sovRegex(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRegex(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRegex(x uint64) (n int) {
	return sovRegex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegexMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegex
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
			return fmt.Errorf("proto: RegexMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegexMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleRe2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegex
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
				return ErrInvalidLengthRegex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRegex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RegexMatcher_GoogleRE2{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.EngineType = &RegexMatcher_GoogleRe2{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegex
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
				return ErrInvalidLengthRegex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRegex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Regex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegex
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRegex
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
func (m *RegexMatcher_GoogleRE2) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRegex
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
			return fmt.Errorf("proto: GoogleRE2: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoogleRE2: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxProgramSize", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRegex
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
				return ErrInvalidLengthRegex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRegex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxProgramSize == nil {
				m.MaxProgramSize = &types.UInt32Value{}
			}
			if err := m.MaxProgramSize.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRegex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRegex
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRegex
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
func skipRegex(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRegex
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
					return 0, ErrIntOverflowRegex
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
					return 0, ErrIntOverflowRegex
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
				return 0, ErrInvalidLengthRegex
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthRegex
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRegex
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
				next, err := skipRegex(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthRegex
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
	ErrInvalidLengthRegex = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRegex   = fmt.Errorf("proto: integer overflow")
)
