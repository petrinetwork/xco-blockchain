// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kaiju/legacy/did/did.proto

package legacydid

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Digital identity credential issued to an kaiju DID
type DidCredential struct {
	CredType []string `protobuf:"bytes,1,rep,name=cred_type,json=credType,proto3" json:"type" yaml:"type"`
	Issuer   string   `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty" yaml:"issuer"`
	Issued   string   `protobuf:"bytes,3,opt,name=issued,proto3" json:"issued,omitempty" yaml:"issued"`
	Claim    *Claim   `protobuf:"bytes,4,opt,name=claim,proto3" json:"claim,omitempty" yaml:"claim"`
}

func (m *DidCredential) Reset()         { *m = DidCredential{} }
func (m *DidCredential) String() string { return proto.CompactTextString(m) }
func (*DidCredential) ProtoMessage()    {}
func (*DidCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_515f25a9bd100c9a, []int{0}
}
func (m *DidCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DidCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DidCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DidCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DidCredential.Merge(m, src)
}
func (m *DidCredential) XXX_Size() int {
	return m.Size()
}
func (m *DidCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_DidCredential.DiscardUnknown(m)
}

var xxx_messageInfo_DidCredential proto.InternalMessageInfo

func (m *DidCredential) GetCredType() []string {
	if m != nil {
		return m.CredType
	}
	return nil
}

func (m *DidCredential) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *DidCredential) GetIssued() string {
	if m != nil {
		return m.Issued
	}
	return ""
}

func (m *DidCredential) GetClaim() *Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

// The claim section of a credential, indicating if the DID is KYC validated
type Claim struct {
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	KYCValidated bool   `protobuf:"varint,2,opt,name=KYC_validated,json=KYCValidated,proto3" json:"KYCValidated" yaml:"KYCValidated"`
}

func (m *Claim) Reset()         { *m = Claim{} }
func (m *Claim) String() string { return proto.CompactTextString(m) }
func (*Claim) ProtoMessage()    {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_515f25a9bd100c9a, []int{1}
}
func (m *Claim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return m.Size()
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

func (m *Claim) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Claim) GetKYCValidated() bool {
	if m != nil {
		return m.KYCValidated
	}
	return false
}

// An kaiju DID with public and private keys, based on the Sovrin DID spec
type KaijuDid struct {
	Did                 string  `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty" yaml:"did"`
	VerifyKey           string  `protobuf:"bytes,2,opt,name=verify_key,json=verifyKey,proto3" json:"verifyKey" yaml:"verifyKey"`
	EncryptionPublicKey string  `protobuf:"bytes,3,opt,name=encryption_public_key,json=encryptionPublicKey,proto3" json:"encryptionPublicKey" yaml:"encryptionPublicKey"`
	Secret              *Secret `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty" yaml:"secret"`
}

func (m *KaijuDid) Reset()         { *m = KaijuDid{} }
func (m *KaijuDid) String() string { return proto.CompactTextString(m) }
func (*KaijuDid) ProtoMessage()    {}
func (*KaijuDid) Descriptor() ([]byte, []int) {
	return fileDescriptor_515f25a9bd100c9a, []int{2}
}
func (m *KaijuDid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KaijuDid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KaijuDid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KaijuDid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KaijuDid.Merge(m, src)
}
func (m *KaijuDid) XXX_Size() int {
	return m.Size()
}
func (m *KaijuDid) XXX_DiscardUnknown() {
	xxx_messageInfo_KaijuDid.DiscardUnknown(m)
}

var xxx_messageInfo_KaijuDid proto.InternalMessageInfo

func (m *KaijuDid) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *KaijuDid) GetVerifyKey() string {
	if m != nil {
		return m.VerifyKey
	}
	return ""
}

func (m *KaijuDid) GetEncryptionPublicKey() string {
	if m != nil {
		return m.EncryptionPublicKey
	}
	return ""
}

func (m *KaijuDid) GetSecret() *Secret {
	if m != nil {
		return m.Secret
	}
	return nil
}

// The private section of an kaiju DID, based on the Sovrin DID spec
type Secret struct {
	Seed                 string `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty" yaml:"seed"`
	SignKey              string `protobuf:"bytes,2,opt,name=sign_key,json=signKey,proto3" json:"signKey" yaml:"signKey"`
	EncryptionPrivateKey string `protobuf:"bytes,3,opt,name=encryption_private_key,json=encryptionPrivateKey,proto3" json:"encryptionPrivateKey" yaml:"encryptionPrivateKey"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_515f25a9bd100c9a, []int{3}
}
func (m *Secret) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return m.Size()
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *Secret) GetSignKey() string {
	if m != nil {
		return m.SignKey
	}
	return ""
}

func (m *Secret) GetEncryptionPrivateKey() string {
	if m != nil {
		return m.EncryptionPrivateKey
	}
	return ""
}

func init() {
	proto.RegisterType((*DidCredential)(nil), "legacydid.DidCredential")
	proto.RegisterType((*Claim)(nil), "legacydid.Claim")
	proto.RegisterType((*KaijuDid)(nil), "legacydid.KaijuDid")
	proto.RegisterType((*Secret)(nil), "legacydid.Secret")
}

func init() { proto.RegisterFile("kaiju/legacy/did/did.proto", fileDescriptor_515f25a9bd100c9a) }

var fileDescriptor_515f25a9bd100c9a = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x1c, 0xc5, 0x97, 0x6e, 0x2b, 0x8d, 0xb7, 0x8e, 0xce, 0x1b, 0x10, 0x0d, 0x2d, 0x2e, 0xe6, 0x40,
	0x39, 0xd0, 0x48, 0xc0, 0xc4, 0x84, 0x38, 0xa0, 0x76, 0x17, 0x54, 0x24, 0x90, 0x41, 0x48, 0xe3,
	0x52, 0xa5, 0xb1, 0xd7, 0x59, 0x4b, 0xe3, 0x2a, 0x4d, 0xab, 0xe6, 0x5b, 0x70, 0xe1, 0x3b, 0x71,
	0xdc, 0x05, 0xc4, 0xc9, 0x42, 0xed, 0x2d, 0xc7, 0x7c, 0x02, 0xe4, 0x38, 0x4d, 0x83, 0xe8, 0xa1,
	0x52, 0xfd, 0xde, 0xef, 0xfd, 0xff, 0xc9, 0x73, 0x0b, 0x2c, 0x3e, 0x17, 0x8e, 0xcf, 0x86, 0xae,
	0x17, 0x3b, 0x94, 0x53, 0xf5, 0x69, 0x8f, 0x43, 0x11, 0x09, 0x68, 0x6a, 0x95, 0x72, 0x7a, 0x72,
	0x3c, 0x14, 0x43, 0x91, 0xa9, 0x8e, 0xfa, 0xa6, 0x01, 0xfc, 0xd3, 0x00, 0xf5, 0x0b, 0x4e, 0xbb,
	0x21, 0xa3, 0x2c, 0x88, 0xb8, 0xeb, 0xc3, 0x97, 0xc0, 0xf4, 0x42, 0x46, 0xfb, 0x51, 0x3c, 0x66,
	0x96, 0xd1, 0xdc, 0x6e, 0x99, 0x9d, 0x07, 0x89, 0x44, 0x3b, 0xea, 0x9c, 0x4a, 0xb4, 0x17, 0xbb,
	0x23, 0xff, 0x35, 0x56, 0x27, 0x4c, 0x6a, 0x8a, 0xfc, 0x1c, 0x8f, 0x19, 0x7c, 0x0a, 0xaa, 0x7c,
	0x32, 0x99, 0xb2, 0xd0, 0xaa, 0x34, 0x8d, 0x96, 0xd9, 0x39, 0x4c, 0x25, 0xaa, 0x6b, 0x54, 0xeb,
	0x98, 0xe4, 0x40, 0x81, 0x52, 0x6b, 0x7b, 0x23, 0x4a, 0x57, 0x28, 0x85, 0xe7, 0x60, 0xd7, 0xf3,
	0x5d, 0x3e, 0xb2, 0x76, 0x9a, 0x46, 0x6b, 0xef, 0x79, 0xa3, 0x5d, 0xbc, 0x4e, 0xbb, 0xab, 0xf4,
	0x4e, 0x23, 0x95, 0x68, 0x5f, 0x67, 0x33, 0x10, 0x13, 0x1d, 0xc0, 0x11, 0xd8, 0xcd, 0x08, 0x78,
	0x0a, 0x2a, 0x9c, 0x5a, 0x46, 0xb6, 0xa9, 0x9e, 0x4a, 0x64, 0xe6, 0x9b, 0x28, 0x26, 0x15, 0x4e,
	0xe1, 0x7b, 0x50, 0xef, 0x5d, 0x76, 0xfb, 0x33, 0xd7, 0xe7, 0xd4, 0x8d, 0x18, 0xcd, 0x1e, 0xbf,
	0xd6, 0x79, 0x92, 0x48, 0xb4, 0xdf, 0xbb, 0xec, 0x7e, 0x59, 0xe9, 0xa9, 0x44, 0x47, 0x3a, 0x59,
	0x56, 0x31, 0xf9, 0x07, 0xc2, 0xdf, 0x2b, 0xa0, 0xfa, 0x6e, 0x2e, 0x2e, 0x38, 0x85, 0x4d, 0xb0,
	0x4d, 0x8b, 0xc5, 0x07, 0xa9, 0x44, 0x40, 0xc7, 0xa9, 0xda, 0xac, 0x2c, 0xf8, 0x16, 0x80, 0x19,
	0x0b, 0xf9, 0x55, 0xdc, 0xbf, 0x61, 0x71, 0x5e, 0xdb, 0xa3, 0x44, 0x22, 0x53, 0xab, 0x3d, 0x16,
	0xa7, 0x12, 0x35, 0x74, 0xaa, 0x90, 0x30, 0x59, 0xdb, 0x90, 0x83, 0x7b, 0x2c, 0xf0, 0xc2, 0x78,
	0x1c, 0x71, 0x11, 0xf4, 0xc7, 0xd3, 0x81, 0xcf, 0xbd, 0x6c, 0x98, 0x2e, 0xf6, 0x2c, 0x91, 0xe8,
	0x68, 0x0d, 0x7c, 0xcc, 0x7c, 0x3d, 0xf6, 0x44, 0x8f, 0xdd, 0x60, 0x62, 0xb2, 0x29, 0x02, 0xdf,
	0x80, 0xea, 0x84, 0x79, 0x21, 0x8b, 0xf2, 0xab, 0x38, 0x2c, 0x5d, 0xc5, 0xa7, 0xcc, 0x28, 0xdf,
	0xa3, 0x46, 0x31, 0xc9, 0x33, 0xf8, 0x97, 0x01, 0xaa, 0x9a, 0x82, 0x8f, 0xc1, 0xce, 0x84, 0xb1,
	0x55, 0x31, 0x77, 0xd7, 0xbf, 0x28, 0xa5, 0x62, 0x92, 0x99, 0xf0, 0x1c, 0xd4, 0x26, 0x7c, 0x18,
	0x94, 0x8a, 0x39, 0x4d, 0x24, 0xba, 0xa3, 0x34, 0xfd, 0xfc, 0x07, 0x79, 0x46, 0x0b, 0x98, 0xac,
	0x2c, 0x38, 0x02, 0xf7, 0xcb, 0x95, 0x84, 0x7c, 0xe6, 0x46, 0xac, 0xd4, 0xc9, 0xab, 0x44, 0xa2,
	0xe3, 0xd2, 0x0b, 0x6a, 0x40, 0x0f, 0x7d, 0xf8, 0x5f, 0x29, 0x85, 0x8b, 0xc9, 0xc6, 0x50, 0xe7,
	0xc3, 0x8f, 0x85, 0x6d, 0xdc, 0x2e, 0x6c, 0xe3, 0xcf, 0xc2, 0x36, 0xbe, 0x2d, 0xed, 0xad, 0xdb,
	0xa5, 0xbd, 0xf5, 0x7b, 0x69, 0x6f, 0x7d, 0x3d, 0x1b, 0xf2, 0xe8, 0x7a, 0x3a, 0x68, 0x7b, 0x62,
	0xe4, 0xf0, 0xb9, 0xb8, 0x12, 0xd3, 0x80, 0xba, 0x2a, 0xad, 0x4e, 0xcf, 0x06, 0xbe, 0xf0, 0x6e,
	0xbc, 0x6b, 0x97, 0x07, 0x8e, 0xcf, 0x07, 0x4e, 0xd1, 0xe5, 0xa0, 0x9a, 0xfd, 0x2d, 0x5f, 0xfc,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xbf, 0x23, 0x55, 0xd3, 0x03, 0x00, 0x00,
}

func (m *DidCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DidCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DidCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDid(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Issued) > 0 {
		i -= len(m.Issued)
		copy(dAtA[i:], m.Issued)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Issued)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CredType) > 0 {
		for iNdEx := len(m.CredType) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CredType[iNdEx])
			copy(dAtA[i:], m.CredType[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.CredType[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Claim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Claim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Claim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.KYCValidated {
		i--
		if m.KYCValidated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KaijuDid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KaijuDid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KaijuDid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Secret != nil {
		{
			size, err := m.Secret.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDid(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.EncryptionPublicKey) > 0 {
		i -= len(m.EncryptionPublicKey)
		copy(dAtA[i:], m.EncryptionPublicKey)
		i = encodeVarintDid(dAtA, i, uint64(len(m.EncryptionPublicKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VerifyKey) > 0 {
		i -= len(m.VerifyKey)
		copy(dAtA[i:], m.VerifyKey)
		i = encodeVarintDid(dAtA, i, uint64(len(m.VerifyKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Secret) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Secret) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Secret) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptionPrivateKey) > 0 {
		i -= len(m.EncryptionPrivateKey)
		copy(dAtA[i:], m.EncryptionPrivateKey)
		i = encodeVarintDid(dAtA, i, uint64(len(m.EncryptionPrivateKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SignKey) > 0 {
		i -= len(m.SignKey)
		copy(dAtA[i:], m.SignKey)
		i = encodeVarintDid(dAtA, i, uint64(len(m.SignKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Seed) > 0 {
		i -= len(m.Seed)
		copy(dAtA[i:], m.Seed)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Seed)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDid(dAtA []byte, offset int, v uint64) int {
	offset -= sovDid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DidCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CredType) > 0 {
		for _, s := range m.CredType {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.Issued)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovDid(uint64(l))
	}
	return n
}

func (m *Claim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if m.KYCValidated {
		n += 2
	}
	return n
}

func (m *KaijuDid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.VerifyKey)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.EncryptionPublicKey)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if m.Secret != nil {
		l = m.Secret.Size()
		n += 1 + l + sovDid(uint64(l))
	}
	return n
}

func (m *Secret) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Seed)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.SignKey)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.EncryptionPrivateKey)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	return n
}

func sovDid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDid(x uint64) (n int) {
	return sovDid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DidCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
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
			return fmt.Errorf("proto: DidCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DidCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CredType = append(m.CredType, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issued", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issued = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &Claim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDid
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
func (m *Claim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
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
			return fmt.Errorf("proto: Claim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Claim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KYCValidated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
			m.KYCValidated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDid
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
func (m *KaijuDid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
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
			return fmt.Errorf("proto: KaijuDid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KaijuDid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerifyKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptionPublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptionPublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Secret == nil {
				m.Secret = &Secret{}
			}
			if err := m.Secret.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDid
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
func (m *Secret) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
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
			return fmt.Errorf("proto: Secret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Secret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptionPrivateKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
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
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptionPrivateKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDid
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
func skipDid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDid
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
					return 0, ErrIntOverflowDid
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
					return 0, ErrIntOverflowDid
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
				return 0, ErrInvalidLengthDid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDid = fmt.Errorf("proto: unexpected end of group")
)
