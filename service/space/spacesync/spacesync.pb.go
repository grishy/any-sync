// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/space/spacesync/protos/spacesync.proto

package spacesync

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

type ErrCodes int32

const (
	ErrCodes_Unexpected ErrCodes = 0
)

var ErrCodes_name = map[int32]string{
	0: "Unexpected",
}

var ErrCodes_value = map[string]int32{
	"Unexpected": 0,
}

func (x ErrCodes) String() string {
	return proto.EnumName(ErrCodes_name, int32(x))
}

func (ErrCodes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_11d78c2fb7c80384, []int{0}
}

// HeadSyncRange presenting a request for one range
type HeadSyncRange struct {
	From  uint64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To    uint64 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *HeadSyncRange) Reset()         { *m = HeadSyncRange{} }
func (m *HeadSyncRange) String() string { return proto.CompactTextString(m) }
func (*HeadSyncRange) ProtoMessage()    {}
func (*HeadSyncRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_11d78c2fb7c80384, []int{0}
}
func (m *HeadSyncRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeadSyncRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeadSyncRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeadSyncRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadSyncRange.Merge(m, src)
}
func (m *HeadSyncRange) XXX_Size() int {
	return m.Size()
}
func (m *HeadSyncRange) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadSyncRange.DiscardUnknown(m)
}

var xxx_messageInfo_HeadSyncRange proto.InternalMessageInfo

func (m *HeadSyncRange) GetFrom() uint64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *HeadSyncRange) GetTo() uint64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *HeadSyncRange) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// HeadSyncResult presenting a response for one range
type HeadSyncResult struct {
	Hash     []byte                   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Elements []*HeadSyncResultElement `protobuf:"bytes,2,rep,name=elements,proto3" json:"elements,omitempty"`
	Count    uint32                   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *HeadSyncResult) Reset()         { *m = HeadSyncResult{} }
func (m *HeadSyncResult) String() string { return proto.CompactTextString(m) }
func (*HeadSyncResult) ProtoMessage()    {}
func (*HeadSyncResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_11d78c2fb7c80384, []int{1}
}
func (m *HeadSyncResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeadSyncResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeadSyncResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeadSyncResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadSyncResult.Merge(m, src)
}
func (m *HeadSyncResult) XXX_Size() int {
	return m.Size()
}
func (m *HeadSyncResult) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadSyncResult.DiscardUnknown(m)
}

var xxx_messageInfo_HeadSyncResult proto.InternalMessageInfo

func (m *HeadSyncResult) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *HeadSyncResult) GetElements() []*HeadSyncResultElement {
	if m != nil {
		return m.Elements
	}
	return nil
}

func (m *HeadSyncResult) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

// HeadSyncResultElement presenting state of one object
type HeadSyncResultElement struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Head string `protobuf:"bytes,2,opt,name=head,proto3" json:"head,omitempty"`
}

func (m *HeadSyncResultElement) Reset()         { *m = HeadSyncResultElement{} }
func (m *HeadSyncResultElement) String() string { return proto.CompactTextString(m) }
func (*HeadSyncResultElement) ProtoMessage()    {}
func (*HeadSyncResultElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_11d78c2fb7c80384, []int{2}
}
func (m *HeadSyncResultElement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeadSyncResultElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeadSyncResultElement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeadSyncResultElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadSyncResultElement.Merge(m, src)
}
func (m *HeadSyncResultElement) XXX_Size() int {
	return m.Size()
}
func (m *HeadSyncResultElement) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadSyncResultElement.DiscardUnknown(m)
}

var xxx_messageInfo_HeadSyncResultElement proto.InternalMessageInfo

func (m *HeadSyncResultElement) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HeadSyncResultElement) GetHead() string {
	if m != nil {
		return m.Head
	}
	return ""
}

// HeadSyncRequest is a request for HeadSync
type HeadSyncRequest struct {
	SpaceId string           `protobuf:"bytes,1,opt,name=spaceId,proto3" json:"spaceId,omitempty"`
	Ranges  []*HeadSyncRange `protobuf:"bytes,2,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (m *HeadSyncRequest) Reset()         { *m = HeadSyncRequest{} }
func (m *HeadSyncRequest) String() string { return proto.CompactTextString(m) }
func (*HeadSyncRequest) ProtoMessage()    {}
func (*HeadSyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11d78c2fb7c80384, []int{3}
}
func (m *HeadSyncRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeadSyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeadSyncRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeadSyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadSyncRequest.Merge(m, src)
}
func (m *HeadSyncRequest) XXX_Size() int {
	return m.Size()
}
func (m *HeadSyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadSyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeadSyncRequest proto.InternalMessageInfo

func (m *HeadSyncRequest) GetSpaceId() string {
	if m != nil {
		return m.SpaceId
	}
	return ""
}

func (m *HeadSyncRequest) GetRanges() []*HeadSyncRange {
	if m != nil {
		return m.Ranges
	}
	return nil
}

// HeadSyncResponse is a response for HeadSync
type HeadSyncResponse struct {
	Results []*HeadSyncResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (m *HeadSyncResponse) Reset()         { *m = HeadSyncResponse{} }
func (m *HeadSyncResponse) String() string { return proto.CompactTextString(m) }
func (*HeadSyncResponse) ProtoMessage()    {}
func (*HeadSyncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11d78c2fb7c80384, []int{4}
}
func (m *HeadSyncResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeadSyncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeadSyncResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeadSyncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeadSyncResponse.Merge(m, src)
}
func (m *HeadSyncResponse) XXX_Size() int {
	return m.Size()
}
func (m *HeadSyncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeadSyncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeadSyncResponse proto.InternalMessageInfo

func (m *HeadSyncResponse) GetResults() []*HeadSyncResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterEnum("anySpace.ErrCodes", ErrCodes_name, ErrCodes_value)
	proto.RegisterType((*HeadSyncRange)(nil), "anySpace.HeadSyncRange")
	proto.RegisterType((*HeadSyncResult)(nil), "anySpace.HeadSyncResult")
	proto.RegisterType((*HeadSyncResultElement)(nil), "anySpace.HeadSyncResultElement")
	proto.RegisterType((*HeadSyncRequest)(nil), "anySpace.HeadSyncRequest")
	proto.RegisterType((*HeadSyncResponse)(nil), "anySpace.HeadSyncResponse")
}

func init() {
	proto.RegisterFile("service/space/spacesync/protos/spacesync.proto", fileDescriptor_11d78c2fb7c80384)
}

var fileDescriptor_11d78c2fb7c80384 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbf, 0x6e, 0xe2, 0x40,
	0x10, 0xc6, 0x6d, 0xf3, 0xcf, 0xcc, 0x1d, 0x1c, 0x5a, 0xdd, 0x09, 0x1f, 0x85, 0x0f, 0xb9, 0x42,
	0x57, 0x18, 0x85, 0x94, 0x54, 0x49, 0x44, 0x14, 0x52, 0x2e, 0x4a, 0x13, 0xa5, 0x71, 0xec, 0x49,
	0xb0, 0x04, 0xbb, 0x8e, 0x77, 0x49, 0xc2, 0x5b, 0xe4, 0xb1, 0x52, 0x52, 0xa6, 0x8c, 0xe0, 0x45,
	0x22, 0x0f, 0x31, 0x24, 0x92, 0x69, 0xac, 0x99, 0xf1, 0x7c, 0xdf, 0xfc, 0x76, 0x34, 0xe0, 0x2b,
	0x4c, 0x1f, 0xe3, 0x10, 0xfb, 0x2a, 0x09, 0xf2, 0xaf, 0x5a, 0x8a, 0xb0, 0x9f, 0xa4, 0x52, 0x4b,
	0xb5, 0x2f, 0xf8, 0x54, 0x60, 0x76, 0x20, 0x96, 0x93, 0xac, 0xe6, 0x8d, 0xa1, 0x71, 0x81, 0x41,
	0x34, 0x59, 0x8a, 0x90, 0x07, 0xe2, 0x1e, 0x19, 0x83, 0xf2, 0x5d, 0x2a, 0xe7, 0x8e, 0xd9, 0x35,
	0x7b, 0x65, 0x4e, 0x31, 0x6b, 0x82, 0xa5, 0xa5, 0x63, 0x51, 0xc5, 0xd2, 0x92, 0xfd, 0x86, 0xca,
	0x2c, 0x9e, 0xc7, 0xda, 0x29, 0x75, 0xcd, 0x5e, 0x83, 0x6f, 0x13, 0xef, 0x09, 0x9a, 0x3b, 0x2b,
	0x54, 0x8b, 0x99, 0xce, 0xbc, 0xa6, 0x81, 0x9a, 0x92, 0xd7, 0x4f, 0x4e, 0x31, 0x1b, 0x82, 0x8d,
	0x33, 0x9c, 0xa3, 0xd0, 0xca, 0xb1, 0xba, 0xa5, 0xde, 0x8f, 0xc1, 0x3f, 0x3f, 0xa7, 0xf1, 0xbf,
	0xeb, 0x47, 0xdb, 0x3e, 0xbe, 0x13, 0x64, 0x83, 0x43, 0xb9, 0x10, 0xbb, 0xc1, 0x94, 0x78, 0x43,
	0xf8, 0x53, 0x28, 0xcc, 0xb8, 0xe3, 0x88, 0xa6, 0xd7, 0xb9, 0x15, 0x47, 0xc4, 0x83, 0x41, 0x44,
	0x2f, 0xa9, 0x73, 0x8a, 0xbd, 0x1b, 0xf8, 0xb5, 0x17, 0x3f, 0x2c, 0x50, 0x69, 0xe6, 0x40, 0x8d,
	0x16, 0x36, 0xce, 0xb5, 0x79, 0xca, 0xfa, 0x50, 0x4d, 0xb3, 0x2d, 0xe5, 0xe8, 0xed, 0x02, 0xf4,
	0xec, 0x3f, 0xff, 0x6c, 0xf3, 0xce, 0xa1, 0xf5, 0x05, 0x2d, 0x91, 0x42, 0x21, 0x1b, 0x40, 0x2d,
	0x25, 0x4c, 0xe5, 0x98, 0xe4, 0xe2, 0x1c, 0x5a, 0x00, 0xcf, 0x1b, 0xff, 0x77, 0xc0, 0x1e, 0xa5,
	0xe9, 0x99, 0x8c, 0x50, 0xb1, 0x26, 0xc0, 0x95, 0xc0, 0xe7, 0x04, 0x43, 0x8d, 0x51, 0xcb, 0x18,
	0x5c, 0x42, 0x85, 0xc4, 0xec, 0x04, 0xec, 0x5c, 0xcf, 0xfe, 0x16, 0x79, 0xd2, 0xf3, 0x3a, 0x9d,
	0xc2, 0x71, 0xc4, 0x76, 0x7a, 0xf4, 0xba, 0x76, 0xcd, 0xd5, 0xda, 0x35, 0xdf, 0xd7, 0xae, 0xf9,
	0xb2, 0x71, 0x8d, 0xd5, 0xc6, 0x35, 0xde, 0x36, 0xae, 0x71, 0xdd, 0x3e, 0x70, 0x62, 0xb7, 0x55,
	0x3a, 0xa9, 0xe3, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x7e, 0xc8, 0x77, 0x84, 0x02, 0x00,
	0x00,
}

func (m *HeadSyncRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeadSyncRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeadSyncRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Limit != 0 {
		i = encodeVarintSpacesync(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x18
	}
	if m.To != 0 {
		i = encodeVarintSpacesync(dAtA, i, uint64(m.To))
		i--
		dAtA[i] = 0x10
	}
	if m.From != 0 {
		i = encodeVarintSpacesync(dAtA, i, uint64(m.From))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HeadSyncResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeadSyncResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeadSyncResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		i = encodeVarintSpacesync(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Elements) > 0 {
		for iNdEx := len(m.Elements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Elements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSpacesync(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintSpacesync(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HeadSyncResultElement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeadSyncResultElement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeadSyncResultElement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Head) > 0 {
		i -= len(m.Head)
		copy(dAtA[i:], m.Head)
		i = encodeVarintSpacesync(dAtA, i, uint64(len(m.Head)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintSpacesync(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HeadSyncRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeadSyncRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeadSyncRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ranges) > 0 {
		for iNdEx := len(m.Ranges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Ranges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSpacesync(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.SpaceId) > 0 {
		i -= len(m.SpaceId)
		copy(dAtA[i:], m.SpaceId)
		i = encodeVarintSpacesync(dAtA, i, uint64(len(m.SpaceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HeadSyncResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeadSyncResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeadSyncResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Results[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSpacesync(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintSpacesync(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpacesync(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HeadSyncRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.From != 0 {
		n += 1 + sovSpacesync(uint64(m.From))
	}
	if m.To != 0 {
		n += 1 + sovSpacesync(uint64(m.To))
	}
	if m.Limit != 0 {
		n += 1 + sovSpacesync(uint64(m.Limit))
	}
	return n
}

func (m *HeadSyncResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovSpacesync(uint64(l))
	}
	if len(m.Elements) > 0 {
		for _, e := range m.Elements {
			l = e.Size()
			n += 1 + l + sovSpacesync(uint64(l))
		}
	}
	if m.Count != 0 {
		n += 1 + sovSpacesync(uint64(m.Count))
	}
	return n
}

func (m *HeadSyncResultElement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSpacesync(uint64(l))
	}
	l = len(m.Head)
	if l > 0 {
		n += 1 + l + sovSpacesync(uint64(l))
	}
	return n
}

func (m *HeadSyncRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpaceId)
	if l > 0 {
		n += 1 + l + sovSpacesync(uint64(l))
	}
	if len(m.Ranges) > 0 {
		for _, e := range m.Ranges {
			l = e.Size()
			n += 1 + l + sovSpacesync(uint64(l))
		}
	}
	return n
}

func (m *HeadSyncResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovSpacesync(uint64(l))
		}
	}
	return n
}

func sovSpacesync(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpacesync(x uint64) (n int) {
	return sovSpacesync(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HeadSyncRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpacesync
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
			return fmt.Errorf("proto: HeadSyncRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeadSyncRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			m.From = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpacesync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.From |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			m.To = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpacesync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.To |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpacesync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSpacesync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpacesync
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
func (m *HeadSyncResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpacesync
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
			return fmt.Errorf("proto: HeadSyncResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeadSyncResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpacesync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSpacesync
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpacesync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpacesync
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
				return ErrInvalidLengthSpacesync
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpacesync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elements = append(m.Elements, &HeadSyncResultElement{})
			if err := m.Elements[len(m.Elements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpacesync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSpacesync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpacesync
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
func (m *HeadSyncResultElement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpacesync
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
			return fmt.Errorf("proto: HeadSyncResultElement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeadSyncResultElement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpacesync
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
				return ErrInvalidLengthSpacesync
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpacesync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Head", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpacesync
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
				return ErrInvalidLengthSpacesync
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpacesync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Head = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpacesync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpacesync
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
func (m *HeadSyncRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpacesync
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
			return fmt.Errorf("proto: HeadSyncRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeadSyncRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpacesync
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
				return ErrInvalidLengthSpacesync
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpacesync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ranges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpacesync
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
				return ErrInvalidLengthSpacesync
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpacesync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ranges = append(m.Ranges, &HeadSyncRange{})
			if err := m.Ranges[len(m.Ranges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpacesync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpacesync
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
func (m *HeadSyncResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpacesync
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
			return fmt.Errorf("proto: HeadSyncResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeadSyncResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpacesync
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
				return ErrInvalidLengthSpacesync
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpacesync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &HeadSyncResult{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpacesync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpacesync
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
func skipSpacesync(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpacesync
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
					return 0, ErrIntOverflowSpacesync
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
					return 0, ErrIntOverflowSpacesync
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
				return 0, ErrInvalidLengthSpacesync
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpacesync
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpacesync
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpacesync        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpacesync          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpacesync = fmt.Errorf("proto: unexpected end of group")
)
