// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v2alpha/certs.proto

package envoy_admin_v2alpha

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Proto representation of certificate details. Admin endpoint uses this wrapper for `/certs` to
// display certificate information. See :ref:`/certs <operations_admin_interface_certs>` for more
// information.
type Certificates struct {
	// List of certificates known to an Envoy.
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Certificates) Reset()         { *m = Certificates{} }
func (m *Certificates) String() string { return proto.CompactTextString(m) }
func (*Certificates) ProtoMessage()    {}
func (*Certificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_certs_7538b07d9e99fb72, []int{0}
}
func (m *Certificates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Certificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Certificates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Certificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificates.Merge(dst, src)
}
func (m *Certificates) XXX_Size() int {
	return m.Size()
}
func (m *Certificates) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificates.DiscardUnknown(m)
}

var xxx_messageInfo_Certificates proto.InternalMessageInfo

func (m *Certificates) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type Certificate struct {
	// Details of CA certificate.
	CaCert []*CertificateDetails `protobuf:"bytes,1,rep,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	// Details of Certificate Chain
	CertChain            []*CertificateDetails `protobuf:"bytes,2,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_certs_7538b07d9e99fb72, []int{1}
}
func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(dst, src)
}
func (m *Certificate) XXX_Size() int {
	return m.Size()
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetCaCert() []*CertificateDetails {
	if m != nil {
		return m.CaCert
	}
	return nil
}

func (m *Certificate) GetCertChain() []*CertificateDetails {
	if m != nil {
		return m.CertChain
	}
	return nil
}

type CertificateDetails struct {
	// Path of the certificate.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Certificate Serial Number.
	SerialNumber string `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	// List of Subject Alternate names.
	SubjectAltNames []*SubjectAlternateName `protobuf:"bytes,3,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	// Minimum of days until expiration of certificate and it's chain.
	DaysUntilExpiration  uint64   `protobuf:"varint,4,opt,name=days_until_expiration,json=daysUntilExpiration,proto3" json:"days_until_expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateDetails) Reset()         { *m = CertificateDetails{} }
func (m *CertificateDetails) String() string { return proto.CompactTextString(m) }
func (*CertificateDetails) ProtoMessage()    {}
func (*CertificateDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_certs_7538b07d9e99fb72, []int{2}
}
func (m *CertificateDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertificateDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertificateDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CertificateDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateDetails.Merge(dst, src)
}
func (m *CertificateDetails) XXX_Size() int {
	return m.Size()
}
func (m *CertificateDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateDetails proto.InternalMessageInfo

func (m *CertificateDetails) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CertificateDetails) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *CertificateDetails) GetSubjectAltNames() []*SubjectAlternateName {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

func (m *CertificateDetails) GetDaysUntilExpiration() uint64 {
	if m != nil {
		return m.DaysUntilExpiration
	}
	return 0
}

type SubjectAlternateName struct {
	// Subject Alternate Name.
	//
	// Types that are valid to be assigned to Name:
	//	*SubjectAlternateName_Dns
	//	*SubjectAlternateName_Uri
	Name                 isSubjectAlternateName_Name `protobuf_oneof:"name"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SubjectAlternateName) Reset()         { *m = SubjectAlternateName{} }
func (m *SubjectAlternateName) String() string { return proto.CompactTextString(m) }
func (*SubjectAlternateName) ProtoMessage()    {}
func (*SubjectAlternateName) Descriptor() ([]byte, []int) {
	return fileDescriptor_certs_7538b07d9e99fb72, []int{3}
}
func (m *SubjectAlternateName) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubjectAlternateName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubjectAlternateName.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SubjectAlternateName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubjectAlternateName.Merge(dst, src)
}
func (m *SubjectAlternateName) XXX_Size() int {
	return m.Size()
}
func (m *SubjectAlternateName) XXX_DiscardUnknown() {
	xxx_messageInfo_SubjectAlternateName.DiscardUnknown(m)
}

var xxx_messageInfo_SubjectAlternateName proto.InternalMessageInfo

type isSubjectAlternateName_Name interface {
	isSubjectAlternateName_Name()
	MarshalTo([]byte) (int, error)
	Size() int
}

type SubjectAlternateName_Dns struct {
	Dns string `protobuf:"bytes,1,opt,name=dns,proto3,oneof"`
}
type SubjectAlternateName_Uri struct {
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof"`
}

func (*SubjectAlternateName_Dns) isSubjectAlternateName_Name() {}
func (*SubjectAlternateName_Uri) isSubjectAlternateName_Name() {}

func (m *SubjectAlternateName) GetName() isSubjectAlternateName_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SubjectAlternateName) GetDns() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Dns); ok {
		return x.Dns
	}
	return ""
}

func (m *SubjectAlternateName) GetUri() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Uri); ok {
		return x.Uri
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SubjectAlternateName) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SubjectAlternateName_OneofMarshaler, _SubjectAlternateName_OneofUnmarshaler, _SubjectAlternateName_OneofSizer, []interface{}{
		(*SubjectAlternateName_Dns)(nil),
		(*SubjectAlternateName_Uri)(nil),
	}
}

func _SubjectAlternateName_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SubjectAlternateName)
	// name
	switch x := m.Name.(type) {
	case *SubjectAlternateName_Dns:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Dns)
	case *SubjectAlternateName_Uri:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Uri)
	case nil:
	default:
		return fmt.Errorf("SubjectAlternateName.Name has unexpected type %T", x)
	}
	return nil
}

func _SubjectAlternateName_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SubjectAlternateName)
	switch tag {
	case 1: // name.dns
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Name = &SubjectAlternateName_Dns{x}
		return true, err
	case 2: // name.uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Name = &SubjectAlternateName_Uri{x}
		return true, err
	default:
		return false, nil
	}
}

func _SubjectAlternateName_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SubjectAlternateName)
	// name
	switch x := m.Name.(type) {
	case *SubjectAlternateName_Dns:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Dns)))
		n += len(x.Dns)
	case *SubjectAlternateName_Uri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Uri)))
		n += len(x.Uri)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Certificates)(nil), "envoy.admin.v2alpha.Certificates")
	proto.RegisterType((*Certificate)(nil), "envoy.admin.v2alpha.Certificate")
	proto.RegisterType((*CertificateDetails)(nil), "envoy.admin.v2alpha.CertificateDetails")
	proto.RegisterType((*SubjectAlternateName)(nil), "envoy.admin.v2alpha.SubjectAlternateName")
}
func (m *Certificates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Certificates) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Certificates) > 0 {
		for _, msg := range m.Certificates {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCerts(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Certificate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Certificate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.CaCert) > 0 {
		for _, msg := range m.CaCert {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCerts(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.CertChain) > 0 {
		for _, msg := range m.CertChain {
			dAtA[i] = 0x12
			i++
			i = encodeVarintCerts(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CertificateDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertificateDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCerts(dAtA, i, uint64(len(m.Path)))
		i += copy(dAtA[i:], m.Path)
	}
	if len(m.SerialNumber) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCerts(dAtA, i, uint64(len(m.SerialNumber)))
		i += copy(dAtA[i:], m.SerialNumber)
	}
	if len(m.SubjectAltNames) > 0 {
		for _, msg := range m.SubjectAltNames {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintCerts(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.DaysUntilExpiration != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintCerts(dAtA, i, uint64(m.DaysUntilExpiration))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SubjectAlternateName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubjectAlternateName) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Name != nil {
		nn1, err := m.Name.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SubjectAlternateName_Dns) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintCerts(dAtA, i, uint64(len(m.Dns)))
	i += copy(dAtA[i:], m.Dns)
	return i, nil
}
func (m *SubjectAlternateName_Uri) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x12
	i++
	i = encodeVarintCerts(dAtA, i, uint64(len(m.Uri)))
	i += copy(dAtA[i:], m.Uri)
	return i, nil
}
func encodeVarintCerts(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Certificates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Certificates) > 0 {
		for _, e := range m.Certificates {
			l = e.Size()
			n += 1 + l + sovCerts(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Certificate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CaCert) > 0 {
		for _, e := range m.CaCert {
			l = e.Size()
			n += 1 + l + sovCerts(uint64(l))
		}
	}
	if len(m.CertChain) > 0 {
		for _, e := range m.CertChain {
			l = e.Size()
			n += 1 + l + sovCerts(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CertificateDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovCerts(uint64(l))
	}
	l = len(m.SerialNumber)
	if l > 0 {
		n += 1 + l + sovCerts(uint64(l))
	}
	if len(m.SubjectAltNames) > 0 {
		for _, e := range m.SubjectAltNames {
			l = e.Size()
			n += 1 + l + sovCerts(uint64(l))
		}
	}
	if m.DaysUntilExpiration != 0 {
		n += 1 + sovCerts(uint64(m.DaysUntilExpiration))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SubjectAlternateName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Name != nil {
		n += m.Name.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SubjectAlternateName_Dns) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Dns)
	n += 1 + l + sovCerts(uint64(l))
	return n
}
func (m *SubjectAlternateName_Uri) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uri)
	n += 1 + l + sovCerts(uint64(l))
	return n
}

func sovCerts(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCerts(x uint64) (n int) {
	return sovCerts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Certificates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCerts
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
			return fmt.Errorf("proto: Certificates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Certificates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificates = append(m.Certificates, &Certificate{})
			if err := m.Certificates[len(m.Certificates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCerts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCerts
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
func (m *Certificate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCerts
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
			return fmt.Errorf("proto: Certificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Certificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CaCert", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CaCert = append(m.CaCert, &CertificateDetails{})
			if err := m.CaCert[len(m.CaCert)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertChain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertChain = append(m.CertChain, &CertificateDetails{})
			if err := m.CertChain[len(m.CertChain)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCerts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCerts
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
func (m *CertificateDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCerts
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
			return fmt.Errorf("proto: CertificateDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerialNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerialNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectAltNames", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectAltNames = append(m.SubjectAltNames, &SubjectAlternateName{})
			if err := m.SubjectAltNames[len(m.SubjectAltNames)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaysUntilExpiration", wireType)
			}
			m.DaysUntilExpiration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DaysUntilExpiration |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCerts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCerts
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
func (m *SubjectAlternateName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCerts
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
			return fmt.Errorf("proto: SubjectAlternateName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubjectAlternateName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dns", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = &SubjectAlternateName_Dns{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCerts
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
				return ErrInvalidLengthCerts
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = &SubjectAlternateName_Uri{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCerts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCerts
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
func skipCerts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCerts
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
					return 0, ErrIntOverflowCerts
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
					return 0, ErrIntOverflowCerts
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
				return 0, ErrInvalidLengthCerts
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCerts
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
				next, err := skipCerts(dAtA[start:])
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
	ErrInvalidLengthCerts = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCerts   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/certs.proto", fileDescriptor_certs_7538b07d9e99fb72)
}

var fileDescriptor_certs_7538b07d9e99fb72 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6a, 0xea, 0x40,
	0x14, 0x86, 0xef, 0x68, 0xf0, 0xe2, 0x31, 0x97, 0xcb, 0x1d, 0x6f, 0x21, 0x2b, 0x1b, 0xd2, 0x45,
	0xed, 0x26, 0x82, 0x7d, 0x81, 0x56, 0xad, 0x74, 0xe5, 0x22, 0xad, 0xeb, 0x70, 0x1c, 0xa7, 0x38,
	0x25, 0x99, 0x84, 0x99, 0x89, 0xd4, 0x27, 0xe9, 0x2b, 0x75, 0xd9, 0x65, 0x97, 0xc5, 0x27, 0x29,
	0x33, 0x11, 0xb1, 0x54, 0x28, 0xdd, 0xcd, 0xf9, 0xff, 0xff, 0xfb, 0x39, 0x07, 0x06, 0x4e, 0xb9,
	0x5c, 0x17, 0x9b, 0x01, 0x2e, 0x73, 0x21, 0x07, 0xeb, 0x21, 0x66, 0xe5, 0x0a, 0x07, 0x8c, 0x2b,
	0xa3, 0xe3, 0x52, 0x15, 0xa6, 0xa0, 0x5d, 0x17, 0x88, 0x5d, 0x20, 0xde, 0x05, 0xa2, 0x7b, 0xf0,
	0xc7, 0x5c, 0x19, 0xf1, 0x20, 0x18, 0x1a, 0xae, 0xe9, 0x04, 0x7c, 0x76, 0x30, 0x07, 0x24, 0x6c,
	0xf6, 0x3b, 0xc3, 0x30, 0x3e, 0xc2, 0xc6, 0x07, 0x60, 0xf2, 0x89, 0x8a, 0x9e, 0x09, 0x74, 0x0e,
	0x5c, 0x7a, 0x05, 0xbf, 0x19, 0xa6, 0x36, 0xb2, 0x2b, 0x3c, 0xff, 0xae, 0x70, 0xc2, 0x0d, 0x8a,
	0x4c, 0x27, 0x2d, 0x86, 0x56, 0xa5, 0x53, 0x00, 0x8b, 0xa7, 0x6c, 0x85, 0x42, 0x06, 0x8d, 0x9f,
	0x95, 0xb4, 0x2d, 0x3a, 0xb6, 0x64, 0xf4, 0x46, 0x80, 0x7e, 0x4d, 0x50, 0x0a, 0x5e, 0x89, 0x66,
	0x15, 0x90, 0x90, 0xf4, 0xdb, 0x89, 0x7b, 0xd3, 0x33, 0xf8, 0xa3, 0xb9, 0x12, 0x98, 0xa5, 0xb2,
	0xca, 0x17, 0x5c, 0x05, 0x0d, 0x67, 0xfa, 0xb5, 0x38, 0x73, 0x1a, 0x9d, 0xc3, 0x3f, 0x5d, 0x2d,
	0x1e, 0x39, 0x33, 0x29, 0x66, 0x26, 0x95, 0x98, 0x73, 0x1d, 0x34, 0xdd, 0x7a, 0x17, 0x47, 0xd7,
	0xbb, 0xab, 0xd3, 0xd7, 0x99, 0xe1, 0x4a, 0xa2, 0xe1, 0x33, 0xcc, 0x79, 0xf2, 0x57, 0xef, 0x55,
	0x3b, 0x6b, 0x3a, 0x84, 0x93, 0x25, 0x6e, 0x74, 0x5a, 0x49, 0x23, 0xb2, 0x94, 0x3f, 0x95, 0x42,
	0xa1, 0x11, 0x85, 0x0c, 0xbc, 0x90, 0xf4, 0xbd, 0xa4, 0x6b, 0xcd, 0xb9, 0xf5, 0x6e, 0xf6, 0x56,
	0x34, 0x85, 0xff, 0xc7, 0xca, 0x29, 0x85, 0xe6, 0x52, 0xea, 0xfa, 0xb4, 0xdb, 0x5f, 0x89, 0x1d,
	0xac, 0x56, 0x29, 0x51, 0x5f, 0x64, 0xb5, 0x4a, 0x89, 0x51, 0x0b, 0x3c, 0xbb, 0xfe, 0xc8, 0x7f,
	0xd9, 0xf6, 0xc8, 0xeb, 0xb6, 0x47, 0xde, 0xb7, 0x3d, 0xb2, 0x68, 0xb9, 0xcf, 0x73, 0xf9, 0x11,
	0x00, 0x00, 0xff, 0xff, 0x65, 0xb9, 0x56, 0xe4, 0x5f, 0x02, 0x00, 0x00,
}
