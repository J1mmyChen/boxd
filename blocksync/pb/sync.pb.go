// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sync.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import pb "github.com/BOXFoundation/boxd/core/pb"

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

type LocateHeaders struct {
	// locate hashes, normally it is as following:
	// n, n-1, ... n-k, n-k-2, n-k-5, n-k-10, ... n-k-(2^m+m-1), ... genesis
	// n is tail height, k is sequence part length, m is distance factor
	// n-k-(2^m+m-1) is the (k+m)th element
	// to ensure hash elements that the more near to genesis the more looser
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes" json:"hashes,omitempty"`
}

func (m *LocateHeaders) Reset()         { *m = LocateHeaders{} }
func (m *LocateHeaders) String() string { return proto.CompactTextString(m) }
func (*LocateHeaders) ProtoMessage()    {}
func (*LocateHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5e9f013941885146, []int{0}
}
func (m *LocateHeaders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocateHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocateHeaders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *LocateHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocateHeaders.Merge(dst, src)
}
func (m *LocateHeaders) XXX_Size() int {
	return m.Size()
}
func (m *LocateHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_LocateHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_LocateHeaders proto.InternalMessageInfo

func (m *LocateHeaders) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type SyncHeaders struct {
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes" json:"hashes,omitempty"`
}

func (m *SyncHeaders) Reset()         { *m = SyncHeaders{} }
func (m *SyncHeaders) String() string { return proto.CompactTextString(m) }
func (*SyncHeaders) ProtoMessage()    {}
func (*SyncHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5e9f013941885146, []int{1}
}
func (m *SyncHeaders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncHeaders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SyncHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncHeaders.Merge(dst, src)
}
func (m *SyncHeaders) XXX_Size() int {
	return m.Size()
}
func (m *SyncHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_SyncHeaders proto.InternalMessageInfo

func (m *SyncHeaders) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type CheckHash struct {
	BeginHash []byte `protobuf:"bytes,1,opt,name=begin_hash,json=beginHash,proto3" json:"begin_hash,omitempty"`
	Length    int32  `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (m *CheckHash) Reset()         { *m = CheckHash{} }
func (m *CheckHash) String() string { return proto.CompactTextString(m) }
func (*CheckHash) ProtoMessage()    {}
func (*CheckHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5e9f013941885146, []int{2}
}
func (m *CheckHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CheckHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CheckHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CheckHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckHash.Merge(dst, src)
}
func (m *CheckHash) XXX_Size() int {
	return m.Size()
}
func (m *CheckHash) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckHash.DiscardUnknown(m)
}

var xxx_messageInfo_CheckHash proto.InternalMessageInfo

func (m *CheckHash) GetBeginHash() []byte {
	if m != nil {
		return m.BeginHash
	}
	return nil
}

func (m *CheckHash) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type SyncCheckHash struct {
	// it is a root hash for headers between start header and end header
	RootHash []byte `protobuf:"bytes,1,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
}

func (m *SyncCheckHash) Reset()         { *m = SyncCheckHash{} }
func (m *SyncCheckHash) String() string { return proto.CompactTextString(m) }
func (*SyncCheckHash) ProtoMessage()    {}
func (*SyncCheckHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5e9f013941885146, []int{3}
}
func (m *SyncCheckHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncCheckHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncCheckHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SyncCheckHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncCheckHash.Merge(dst, src)
}
func (m *SyncCheckHash) XXX_Size() int {
	return m.Size()
}
func (m *SyncCheckHash) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncCheckHash.DiscardUnknown(m)
}

var xxx_messageInfo_SyncCheckHash proto.InternalMessageInfo

func (m *SyncCheckHash) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

type FetchBlockHeaders struct {
	Idx       int32  `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	BeginHash []byte `protobuf:"bytes,2,opt,name=begin_hash,json=beginHash,proto3" json:"begin_hash,omitempty"`
	Length    int32  `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
}

func (m *FetchBlockHeaders) Reset()         { *m = FetchBlockHeaders{} }
func (m *FetchBlockHeaders) String() string { return proto.CompactTextString(m) }
func (*FetchBlockHeaders) ProtoMessage()    {}
func (*FetchBlockHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5e9f013941885146, []int{4}
}
func (m *FetchBlockHeaders) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchBlockHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchBlockHeaders.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *FetchBlockHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchBlockHeaders.Merge(dst, src)
}
func (m *FetchBlockHeaders) XXX_Size() int {
	return m.Size()
}
func (m *FetchBlockHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchBlockHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_FetchBlockHeaders proto.InternalMessageInfo

func (m *FetchBlockHeaders) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *FetchBlockHeaders) GetBeginHash() []byte {
	if m != nil {
		return m.BeginHash
	}
	return nil
}

func (m *FetchBlockHeaders) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

type SyncBlocks struct {
	Idx    int32       `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	Blocks []*pb.Block `protobuf:"bytes,2,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *SyncBlocks) Reset()         { *m = SyncBlocks{} }
func (m *SyncBlocks) String() string { return proto.CompactTextString(m) }
func (*SyncBlocks) ProtoMessage()    {}
func (*SyncBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_5e9f013941885146, []int{5}
}
func (m *SyncBlocks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncBlocks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SyncBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncBlocks.Merge(dst, src)
}
func (m *SyncBlocks) XXX_Size() int {
	return m.Size()
}
func (m *SyncBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_SyncBlocks proto.InternalMessageInfo

func (m *SyncBlocks) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *SyncBlocks) GetBlocks() []*pb.Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*LocateHeaders)(nil), "pb.LocateHeaders")
	proto.RegisterType((*SyncHeaders)(nil), "pb.SyncHeaders")
	proto.RegisterType((*CheckHash)(nil), "pb.CheckHash")
	proto.RegisterType((*SyncCheckHash)(nil), "pb.SyncCheckHash")
	proto.RegisterType((*FetchBlockHeaders)(nil), "pb.FetchBlockHeaders")
	proto.RegisterType((*SyncBlocks)(nil), "pb.SyncBlocks")
}
func (m *LocateHeaders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocateHeaders) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, b := range m.Hashes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSync(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func (m *SyncHeaders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncHeaders) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, b := range m.Hashes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSync(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func (m *CheckHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckHash) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.BeginHash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSync(dAtA, i, uint64(len(m.BeginHash)))
		i += copy(dAtA[i:], m.BeginHash)
	}
	if m.Length != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSync(dAtA, i, uint64(m.Length))
	}
	return i, nil
}

func (m *SyncCheckHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncCheckHash) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RootHash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSync(dAtA, i, uint64(len(m.RootHash)))
		i += copy(dAtA[i:], m.RootHash)
	}
	return i, nil
}

func (m *FetchBlockHeaders) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchBlockHeaders) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Idx != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSync(dAtA, i, uint64(m.Idx))
	}
	if len(m.BeginHash) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSync(dAtA, i, uint64(len(m.BeginHash)))
		i += copy(dAtA[i:], m.BeginHash)
	}
	if m.Length != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSync(dAtA, i, uint64(m.Length))
	}
	return i, nil
}

func (m *SyncBlocks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncBlocks) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Idx != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSync(dAtA, i, uint64(m.Idx))
	}
	if len(m.Blocks) > 0 {
		for _, msg := range m.Blocks {
			dAtA[i] = 0x12
			i++
			i = encodeVarintSync(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintSync(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *LocateHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, b := range m.Hashes {
			l = len(b)
			n += 1 + l + sovSync(uint64(l))
		}
	}
	return n
}

func (m *SyncHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, b := range m.Hashes {
			l = len(b)
			n += 1 + l + sovSync(uint64(l))
		}
	}
	return n
}

func (m *CheckHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BeginHash)
	if l > 0 {
		n += 1 + l + sovSync(uint64(l))
	}
	if m.Length != 0 {
		n += 1 + sovSync(uint64(m.Length))
	}
	return n
}

func (m *SyncCheckHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RootHash)
	if l > 0 {
		n += 1 + l + sovSync(uint64(l))
	}
	return n
}

func (m *FetchBlockHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Idx != 0 {
		n += 1 + sovSync(uint64(m.Idx))
	}
	l = len(m.BeginHash)
	if l > 0 {
		n += 1 + l + sovSync(uint64(l))
	}
	if m.Length != 0 {
		n += 1 + sovSync(uint64(m.Length))
	}
	return n
}

func (m *SyncBlocks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Idx != 0 {
		n += 1 + sovSync(uint64(m.Idx))
	}
	if len(m.Blocks) > 0 {
		for _, e := range m.Blocks {
			l = e.Size()
			n += 1 + l + sovSync(uint64(l))
		}
	}
	return n
}

func sovSync(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSync(x uint64) (n int) {
	return sovSync(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LocateHeaders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
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
			return fmt.Errorf("proto: LocateHeaders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocateHeaders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
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
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashes = append(m.Hashes, make([]byte, postIndex-iNdEx))
			copy(m.Hashes[len(m.Hashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func (m *SyncHeaders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
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
			return fmt.Errorf("proto: SyncHeaders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncHeaders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
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
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashes = append(m.Hashes, make([]byte, postIndex-iNdEx))
			copy(m.Hashes[len(m.Hashes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func (m *CheckHash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
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
			return fmt.Errorf("proto: CheckHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
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
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeginHash = append(m.BeginHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BeginHash == nil {
				m.BeginHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func (m *SyncCheckHash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
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
			return fmt.Errorf("proto: SyncCheckHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncCheckHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
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
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootHash = append(m.RootHash[:0], dAtA[iNdEx:postIndex]...)
			if m.RootHash == nil {
				m.RootHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func (m *FetchBlockHeaders) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
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
			return fmt.Errorf("proto: FetchBlockHeaders: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchBlockHeaders: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Idx", wireType)
			}
			m.Idx = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Idx |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
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
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeginHash = append(m.BeginHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BeginHash == nil {
				m.BeginHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func (m *SyncBlocks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSync
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
			return fmt.Errorf("proto: SyncBlocks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncBlocks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Idx", wireType)
			}
			m.Idx = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Idx |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSync
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
				return ErrInvalidLengthSync
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocks = append(m.Blocks, &pb.Block{})
			if err := m.Blocks[len(m.Blocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSync
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
func skipSync(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSync
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
					return 0, ErrIntOverflowSync
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
					return 0, ErrIntOverflowSync
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
				return 0, ErrInvalidLengthSync
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSync
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
				next, err := skipSync(dAtA[start:])
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
	ErrInvalidLengthSync = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSync   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sync.proto", fileDescriptor_sync_5e9f013941885146) }

var fileDescriptor_sync_5e9f013941885146 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x97, 0x96, 0x15, 0xf7, 0x6e, 0x03, 0xdd, 0x61, 0x14, 0xc5, 0x50, 0x0a, 0xc3, 0x1e,
	0xa4, 0x45, 0xfd, 0x06, 0x15, 0xc7, 0x0e, 0x82, 0x50, 0x2f, 0x1e, 0x04, 0x69, 0xd2, 0xb0, 0x94,
	0xcd, 0xa4, 0x34, 0x19, 0x6c, 0xdf, 0xc2, 0x8f, 0xe5, 0x71, 0x47, 0x8f, 0xb2, 0x7e, 0x11, 0x49,
	0xac, 0xf8, 0x07, 0xc5, 0x5b, 0xde, 0xf7, 0x79, 0xdf, 0x27, 0xbf, 0x3c, 0x01, 0x50, 0x1b, 0x41,
	0xe3, 0xaa, 0x96, 0x5a, 0x8e, 0x9c, 0x8a, 0x1c, 0x9e, 0xcd, 0x4b, 0xcd, 0x57, 0x24, 0xa6, 0xf2,
	0x31, 0x49, 0x6f, 0xee, 0xa6, 0x72, 0x25, 0x8a, 0x5c, 0x97, 0x52, 0x24, 0x44, 0xae, 0x8b, 0x84,
	0xca, 0x9a, 0x25, 0x15, 0x49, 0xc8, 0x52, 0xd2, 0xc5, 0xfb, 0x5a, 0x78, 0x02, 0xc3, 0x6b, 0x49,
	0x73, 0xcd, 0x66, 0x2c, 0x2f, 0x58, 0xad, 0x46, 0x63, 0xf0, 0x78, 0xae, 0x38, 0x53, 0x3e, 0x0a,
	0xdc, 0x68, 0x90, 0xb5, 0x55, 0x38, 0x81, 0xfe, 0xed, 0x46, 0xd0, 0xff, 0xc6, 0x52, 0xe8, 0x5d,
	0x72, 0x46, 0x17, 0xb3, 0x5c, 0xf1, 0xd1, 0x31, 0x00, 0x61, 0xf3, 0x52, 0x3c, 0x18, 0xd1, 0x47,
	0x01, 0x8a, 0x06, 0x59, 0xcf, 0x76, 0xac, 0x3c, 0x06, 0x6f, 0xc9, 0xc4, 0x5c, 0x73, 0xdf, 0x09,
	0x50, 0xd4, 0xcd, 0xda, 0x2a, 0x3c, 0x85, 0xa1, 0xb9, 0xea, 0xd3, 0xe7, 0x08, 0x7a, 0xb5, 0x94,
	0xfa, 0xab, 0xcd, 0x9e, 0x69, 0x18, 0x31, 0xbc, 0x87, 0x83, 0x29, 0xd3, 0x94, 0xa7, 0xe6, 0x55,
	0x1f, 0x78, 0xfb, 0xe0, 0x96, 0xc5, 0xda, 0xce, 0x76, 0x33, 0x73, 0xfc, 0xc1, 0xe2, 0xfc, 0xcd,
	0xe2, 0x7e, 0x63, 0xb9, 0x02, 0x30, 0x2c, 0xd6, 0xfc, 0x37, 0xdb, 0x09, 0x78, 0x36, 0x4e, 0xe5,
	0x3b, 0x81, 0x1b, 0xf5, 0xcf, 0x87, 0xb1, 0x49, 0xb9, 0x22, 0xb1, 0xdd, 0xc8, 0x5a, 0x31, 0xf5,
	0x9f, 0x77, 0x18, 0x6d, 0x77, 0x18, 0xbd, 0xee, 0x30, 0x7a, 0x6a, 0x70, 0x67, 0xdb, 0xe0, 0xce,
	0x4b, 0x83, 0x3b, 0xc4, 0xb3, 0xff, 0x70, 0xf1, 0x16, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xf3, 0x5c,
	0xe4, 0xcc, 0x01, 0x00, 0x00,
}