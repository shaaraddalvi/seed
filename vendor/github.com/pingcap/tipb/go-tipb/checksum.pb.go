// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checksum.proto

package tipb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ChecksumScanOn int32

const (
	ChecksumScanOn_Table ChecksumScanOn = 0
	ChecksumScanOn_Index ChecksumScanOn = 1
)

var ChecksumScanOn_name = map[int32]string{
	0: "Table",
	1: "Index",
}
var ChecksumScanOn_value = map[string]int32{
	"Table": 0,
	"Index": 1,
}

func (x ChecksumScanOn) Enum() *ChecksumScanOn {
	p := new(ChecksumScanOn)
	*p = x
	return p
}
func (x ChecksumScanOn) String() string {
	return proto.EnumName(ChecksumScanOn_name, int32(x))
}
func (x *ChecksumScanOn) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ChecksumScanOn_value, data, "ChecksumScanOn")
	if err != nil {
		return err
	}
	*x = ChecksumScanOn(value)
	return nil
}
func (ChecksumScanOn) EnumDescriptor() ([]byte, []int) { return fileDescriptorChecksum, []int{0} }

type ChecksumAlgorithm int32

const (
	ChecksumAlgorithm_Crc64_Xor ChecksumAlgorithm = 0
)

var ChecksumAlgorithm_name = map[int32]string{
	0: "Crc64_Xor",
}
var ChecksumAlgorithm_value = map[string]int32{
	"Crc64_Xor": 0,
}

func (x ChecksumAlgorithm) Enum() *ChecksumAlgorithm {
	p := new(ChecksumAlgorithm)
	*p = x
	return p
}
func (x ChecksumAlgorithm) String() string {
	return proto.EnumName(ChecksumAlgorithm_name, int32(x))
}
func (x *ChecksumAlgorithm) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ChecksumAlgorithm_value, data, "ChecksumAlgorithm")
	if err != nil {
		return err
	}
	*x = ChecksumAlgorithm(value)
	return nil
}
func (ChecksumAlgorithm) EnumDescriptor() ([]byte, []int) { return fileDescriptorChecksum, []int{1} }

type ChecksumRewriteRule struct {
	OldPrefix        []byte `protobuf:"bytes,1,opt,name=old_prefix,json=oldPrefix" json:"old_prefix,omitempty"`
	NewPrefix        []byte `protobuf:"bytes,2,opt,name=new_prefix,json=newPrefix" json:"new_prefix,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChecksumRewriteRule) Reset()                    { *m = ChecksumRewriteRule{} }
func (m *ChecksumRewriteRule) String() string            { return proto.CompactTextString(m) }
func (*ChecksumRewriteRule) ProtoMessage()               {}
func (*ChecksumRewriteRule) Descriptor() ([]byte, []int) { return fileDescriptorChecksum, []int{0} }

func (m *ChecksumRewriteRule) GetOldPrefix() []byte {
	if m != nil {
		return m.OldPrefix
	}
	return nil
}

func (m *ChecksumRewriteRule) GetNewPrefix() []byte {
	if m != nil {
		return m.NewPrefix
	}
	return nil
}

type ChecksumRequest struct {
	// Deprecated. Start Ts has been moved to coprocessor.Request.
	StartTsFallback  *uint64              `protobuf:"varint,1,opt,name=start_ts_fallback,json=startTsFallback" json:"start_ts_fallback,omitempty"`
	ScanOn           ChecksumScanOn       `protobuf:"varint,2,opt,name=scan_on,json=scanOn,enum=tipb.ChecksumScanOn" json:"scan_on"`
	Algorithm        ChecksumAlgorithm    `protobuf:"varint,3,opt,name=algorithm,enum=tipb.ChecksumAlgorithm" json:"algorithm"`
	Rule             *ChecksumRewriteRule `protobuf:"bytes,4,opt,name=rule" json:"rule,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ChecksumRequest) Reset()                    { *m = ChecksumRequest{} }
func (m *ChecksumRequest) String() string            { return proto.CompactTextString(m) }
func (*ChecksumRequest) ProtoMessage()               {}
func (*ChecksumRequest) Descriptor() ([]byte, []int) { return fileDescriptorChecksum, []int{1} }

func (m *ChecksumRequest) GetStartTsFallback() uint64 {
	if m != nil && m.StartTsFallback != nil {
		return *m.StartTsFallback
	}
	return 0
}

func (m *ChecksumRequest) GetScanOn() ChecksumScanOn {
	if m != nil {
		return m.ScanOn
	}
	return ChecksumScanOn_Table
}

func (m *ChecksumRequest) GetAlgorithm() ChecksumAlgorithm {
	if m != nil {
		return m.Algorithm
	}
	return ChecksumAlgorithm_Crc64_Xor
}

func (m *ChecksumRequest) GetRule() *ChecksumRewriteRule {
	if m != nil {
		return m.Rule
	}
	return nil
}

type ChecksumResponse struct {
	Checksum         uint64 `protobuf:"varint,1,opt,name=checksum" json:"checksum"`
	TotalKvs         uint64 `protobuf:"varint,2,opt,name=total_kvs,json=totalKvs" json:"total_kvs"`
	TotalBytes       uint64 `protobuf:"varint,3,opt,name=total_bytes,json=totalBytes" json:"total_bytes"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChecksumResponse) Reset()                    { *m = ChecksumResponse{} }
func (m *ChecksumResponse) String() string            { return proto.CompactTextString(m) }
func (*ChecksumResponse) ProtoMessage()               {}
func (*ChecksumResponse) Descriptor() ([]byte, []int) { return fileDescriptorChecksum, []int{2} }

func (m *ChecksumResponse) GetChecksum() uint64 {
	if m != nil {
		return m.Checksum
	}
	return 0
}

func (m *ChecksumResponse) GetTotalKvs() uint64 {
	if m != nil {
		return m.TotalKvs
	}
	return 0
}

func (m *ChecksumResponse) GetTotalBytes() uint64 {
	if m != nil {
		return m.TotalBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*ChecksumRewriteRule)(nil), "tipb.ChecksumRewriteRule")
	proto.RegisterType((*ChecksumRequest)(nil), "tipb.ChecksumRequest")
	proto.RegisterType((*ChecksumResponse)(nil), "tipb.ChecksumResponse")
	proto.RegisterEnum("tipb.ChecksumScanOn", ChecksumScanOn_name, ChecksumScanOn_value)
	proto.RegisterEnum("tipb.ChecksumAlgorithm", ChecksumAlgorithm_name, ChecksumAlgorithm_value)
}
func (m *ChecksumRewriteRule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChecksumRewriteRule) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OldPrefix != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintChecksum(dAtA, i, uint64(len(m.OldPrefix)))
		i += copy(dAtA[i:], m.OldPrefix)
	}
	if m.NewPrefix != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintChecksum(dAtA, i, uint64(len(m.NewPrefix)))
		i += copy(dAtA[i:], m.NewPrefix)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ChecksumRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChecksumRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.StartTsFallback != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintChecksum(dAtA, i, uint64(*m.StartTsFallback))
	}
	dAtA[i] = 0x10
	i++
	i = encodeVarintChecksum(dAtA, i, uint64(m.ScanOn))
	dAtA[i] = 0x18
	i++
	i = encodeVarintChecksum(dAtA, i, uint64(m.Algorithm))
	if m.Rule != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintChecksum(dAtA, i, uint64(m.Rule.Size()))
		n1, err := m.Rule.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ChecksumResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChecksumResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintChecksum(dAtA, i, uint64(m.Checksum))
	dAtA[i] = 0x10
	i++
	i = encodeVarintChecksum(dAtA, i, uint64(m.TotalKvs))
	dAtA[i] = 0x18
	i++
	i = encodeVarintChecksum(dAtA, i, uint64(m.TotalBytes))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintChecksum(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ChecksumRewriteRule) Size() (n int) {
	var l int
	_ = l
	if m.OldPrefix != nil {
		l = len(m.OldPrefix)
		n += 1 + l + sovChecksum(uint64(l))
	}
	if m.NewPrefix != nil {
		l = len(m.NewPrefix)
		n += 1 + l + sovChecksum(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ChecksumRequest) Size() (n int) {
	var l int
	_ = l
	if m.StartTsFallback != nil {
		n += 1 + sovChecksum(uint64(*m.StartTsFallback))
	}
	n += 1 + sovChecksum(uint64(m.ScanOn))
	n += 1 + sovChecksum(uint64(m.Algorithm))
	if m.Rule != nil {
		l = m.Rule.Size()
		n += 1 + l + sovChecksum(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ChecksumResponse) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovChecksum(uint64(m.Checksum))
	n += 1 + sovChecksum(uint64(m.TotalKvs))
	n += 1 + sovChecksum(uint64(m.TotalBytes))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovChecksum(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozChecksum(x uint64) (n int) {
	return sovChecksum(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChecksumRewriteRule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChecksum
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
			return fmt.Errorf("proto: ChecksumRewriteRule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChecksumRewriteRule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldPrefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChecksum
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OldPrefix = append(m.OldPrefix[:0], dAtA[iNdEx:postIndex]...)
			if m.OldPrefix == nil {
				m.OldPrefix = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewPrefix", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChecksum
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewPrefix = append(m.NewPrefix[:0], dAtA[iNdEx:postIndex]...)
			if m.NewPrefix == nil {
				m.NewPrefix = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChecksum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChecksum
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
func (m *ChecksumRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChecksum
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
			return fmt.Errorf("proto: ChecksumRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChecksumRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTsFallback", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StartTsFallback = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScanOn", wireType)
			}
			m.ScanOn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScanOn |= (ChecksumScanOn(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Algorithm", wireType)
			}
			m.Algorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Algorithm |= (ChecksumAlgorithm(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rule", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
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
				return ErrInvalidLengthChecksum
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rule == nil {
				m.Rule = &ChecksumRewriteRule{}
			}
			if err := m.Rule.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChecksum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChecksum
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
func (m *ChecksumResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChecksum
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
			return fmt.Errorf("proto: ChecksumResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChecksumResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checksum", wireType)
			}
			m.Checksum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Checksum |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalKvs", wireType)
			}
			m.TotalKvs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalKvs |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBytes", wireType)
			}
			m.TotalBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChecksum
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalBytes |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChecksum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChecksum
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
func skipChecksum(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChecksum
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
					return 0, ErrIntOverflowChecksum
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
					return 0, ErrIntOverflowChecksum
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
				return 0, ErrInvalidLengthChecksum
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowChecksum
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
				next, err := skipChecksum(dAtA[start:])
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
	ErrInvalidLengthChecksum = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChecksum   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("checksum.proto", fileDescriptorChecksum) }

var fileDescriptorChecksum = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x8e, 0xd3, 0x30,
	0x18, 0xc4, 0x63, 0x08, 0x7f, 0xf2, 0x2d, 0x74, 0xb3, 0x66, 0x11, 0x05, 0x89, 0x50, 0x22, 0x81,
	0x96, 0x4a, 0x04, 0x69, 0x41, 0x5c, 0x38, 0xd1, 0x95, 0x90, 0x10, 0x07, 0x56, 0xd9, 0x1e, 0xb8,
	0x45, 0x8e, 0xe3, 0xa6, 0x51, 0x5d, 0x3b, 0xd8, 0x4e, 0x5b, 0x2e, 0x3c, 0x07, 0x8f, 0xd4, 0x23,
	0x4f, 0x50, 0xa1, 0xf2, 0x22, 0x28, 0x4e, 0xd2, 0x52, 0xed, 0xcd, 0x99, 0xdf, 0x7c, 0x13, 0xcd,
	0x40, 0x8f, 0x4e, 0x19, 0x9d, 0xe9, 0x6a, 0x1e, 0x95, 0x4a, 0x1a, 0x89, 0x5d, 0x53, 0x94, 0xe9,
	0x93, 0xd3, 0x5c, 0xe6, 0xd2, 0x0a, 0x6f, 0xea, 0x57, 0xc3, 0xc2, 0x2b, 0x78, 0x70, 0xd1, 0xba,
	0x63, 0xb6, 0x54, 0x85, 0x61, 0x71, 0xc5, 0x19, 0x7e, 0x0a, 0x20, 0x79, 0x96, 0x94, 0x8a, 0x4d,
	0x8a, 0x55, 0x1f, 0x0d, 0xd0, 0xd9, 0xbd, 0xd8, 0x93, 0x3c, 0xbb, 0xb4, 0x42, 0x8d, 0x05, 0x5b,
	0x76, 0xf8, 0x46, 0x83, 0x05, 0x5b, 0x36, 0x38, 0xdc, 0x20, 0x38, 0xde, 0xa7, 0x7e, 0xaf, 0x98,
	0x36, 0x78, 0x08, 0x27, 0xda, 0x10, 0x65, 0x12, 0xa3, 0x93, 0x09, 0xe1, 0x3c, 0x25, 0x74, 0x66,
	0x83, 0xdd, 0xf8, 0xd8, 0x82, 0xb1, 0xfe, 0xd4, 0xca, 0xf8, 0x2d, 0xdc, 0xd1, 0x94, 0x88, 0x44,
	0x0a, 0x9b, 0xdd, 0x3b, 0x3f, 0x8d, 0xea, 0x0a, 0x51, 0x97, 0x79, 0x45, 0x89, 0xf8, 0x2a, 0x46,
	0xee, 0x7a, 0xf3, 0xcc, 0x89, 0x6f, 0x6b, 0xfb, 0x85, 0x3f, 0x80, 0x47, 0x78, 0x2e, 0x55, 0x61,
	0xa6, 0xf3, 0xfe, 0x4d, 0x7b, 0xf6, 0xe8, 0xf0, 0xec, 0x63, 0x87, 0xdb, 0xcb, 0xbd, 0x1f, 0xbf,
	0x06, 0x57, 0x55, 0x9c, 0xf5, 0xdd, 0x01, 0x3a, 0x3b, 0x3a, 0x7f, 0x7c, 0x78, 0xf7, 0xdf, 0x30,
	0xb1, 0xb5, 0x85, 0x3f, 0xc1, 0xdf, 0x43, 0x5d, 0x4a, 0xa1, 0x19, 0x1e, 0xc0, 0xdd, 0x6e, 0xf7,
	0xa6, 0x57, 0xfb, 0x97, 0x9d, 0x8a, 0x9f, 0x83, 0x67, 0xa4, 0x21, 0x3c, 0x99, 0x2d, 0xb4, 0x2d,
	0xb6, 0xb3, 0x58, 0xf9, 0xcb, 0x42, 0xe3, 0x17, 0x70, 0xd4, 0x58, 0xd2, 0x1f, 0x86, 0x69, 0x5b,
	0xa3, 0x33, 0x81, 0x05, 0xa3, 0x5a, 0x1f, 0xbe, 0x84, 0xde, 0xe1, 0x16, 0xd8, 0x83, 0x5b, 0x63,
	0x92, 0x72, 0xe6, 0x3b, 0xf5, 0xf3, 0xb3, 0xc8, 0xd8, 0xca, 0x47, 0xc3, 0x10, 0x4e, 0xae, 0x95,
	0xc7, 0xf7, 0xc1, 0xbb, 0x50, 0xf4, 0xfd, 0xbb, 0xe4, 0x9b, 0x54, 0xbe, 0x33, 0x7a, 0xb5, 0xde,
	0x06, 0xe8, 0xf7, 0x36, 0x40, 0x7f, 0xb6, 0x01, 0xfa, 0xf5, 0x37, 0x70, 0xe0, 0x21, 0x95, 0xf3,
	0xa8, 0x2c, 0x44, 0x4e, 0x49, 0x19, 0x99, 0x22, 0x4b, 0xed, 0x1c, 0x97, 0xe8, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x1d, 0x58, 0xa2, 0x7a, 0x59, 0x02, 0x00, 0x00,
}