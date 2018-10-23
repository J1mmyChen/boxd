// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sync.proto

package netsyncpb

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
	// n, n-1, ... n-k, n-k-2, n-k-5, n-k-10, ... n-k-(2^m+m), ... genesis
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes" json:"hashes,omitempty"`
}

func (m *LocateHeaders) Reset()         { *m = LocateHeaders{} }
func (m *LocateHeaders) String() string { return proto.CompactTextString(m) }
func (*LocateHeaders) ProtoMessage()    {}
func (*LocateHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_a2297e25a33a1406, []int{0}
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
	return fileDescriptor_sync_a2297e25a33a1406, []int{1}
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
	return fileDescriptor_sync_a2297e25a33a1406, []int{2}
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
	return fileDescriptor_sync_a2297e25a33a1406, []int{3}
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
	BeginHash []byte `protobuf:"bytes,1,opt,name=begin_hash,json=beginHash,proto3" json:"begin_hash,omitempty"`
	Length    int32  `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (m *FetchBlockHeaders) Reset()         { *m = FetchBlockHeaders{} }
func (m *FetchBlockHeaders) String() string { return proto.CompactTextString(m) }
func (*FetchBlockHeaders) ProtoMessage()    {}
func (*FetchBlockHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_a2297e25a33a1406, []int{4}
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
	Blocks []*pb.Block `protobuf:"bytes,1,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *SyncBlocks) Reset()         { *m = SyncBlocks{} }
func (m *SyncBlocks) String() string { return proto.CompactTextString(m) }
func (*SyncBlocks) ProtoMessage()    {}
func (*SyncBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_sync_a2297e25a33a1406, []int{5}
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

func (m *SyncBlocks) GetBlocks() []*pb.Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*LocateHeaders)(nil), "netsyncpb.LocateHeaders")
	proto.RegisterType((*SyncHeaders)(nil), "netsyncpb.SyncHeaders")
	proto.RegisterType((*CheckHash)(nil), "netsyncpb.CheckHash")
	proto.RegisterType((*SyncCheckHash)(nil), "netsyncpb.SyncCheckHash")
	proto.RegisterType((*FetchBlockHeaders)(nil), "netsyncpb.FetchBlockHeaders")
	proto.RegisterType((*SyncBlocks)(nil), "netsyncpb.SyncBlocks")
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
	if len(m.Blocks) > 0 {
		for _, msg := range m.Blocks {
			dAtA[i] = 0xa
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

func init() { proto.RegisterFile("sync.proto", fileDescriptor_sync_a2297e25a33a1406) }

var fileDescriptor_sync_a2297e25a33a1406 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0x9b, 0xff, 0xc7, 0x62, 0xef, 0x4c, 0x17, 0x76, 0x31, 0x14, 0xc5, 0x50, 0x0a, 0x83,
	0x5d, 0x48, 0x8b, 0xce, 0x1b, 0x54, 0x18, 0x06, 0x11, 0x84, 0xba, 0x71, 0x27, 0x4d, 0x1a, 0x9a,
	0x32, 0x63, 0x52, 0x9a, 0x0c, 0x38, 0x6f, 0xe1, 0x63, 0xb9, 0x9c, 0xa5, 0x4b, 0x69, 0x5f, 0x44,
	0x92, 0xa9, 0xe8, 0x4e, 0x70, 0x79, 0xee, 0x97, 0x9c, 0x73, 0xef, 0x01, 0x50, 0x3b, 0x41, 0xd3,
	0xb6, 0x93, 0x5a, 0x06, 0x9e, 0x60, 0xda, 0xc8, 0x96, 0x9c, 0x5e, 0xd5, 0x8d, 0xe6, 0x5b, 0x92,
	0x52, 0xf9, 0x9c, 0xe5, 0xf7, 0x8f, 0x4b, 0xb9, 0x15, 0x55, 0xa9, 0x1b, 0x29, 0x32, 0x22, 0x5f,
	0xaa, 0x8c, 0xca, 0x8e, 0x65, 0x2d, 0xc9, 0xc8, 0x46, 0xd2, 0xf5, 0xe1, 0x77, 0x7c, 0x01, 0xfe,
	0x9d, 0xa4, 0xa5, 0x66, 0x2b, 0x56, 0x56, 0xac, 0x53, 0xc1, 0x0c, 0x5c, 0x5e, 0x2a, 0xce, 0x54,
	0x88, 0xa2, 0xff, 0xc9, 0xb4, 0x18, 0x55, 0x3c, 0x87, 0xc9, 0xc3, 0x4e, 0xd0, 0xdf, 0x9e, 0xe5,
	0xe0, 0xdd, 0x70, 0x46, 0xd7, 0xab, 0x52, 0xf1, 0xe0, 0x1c, 0x80, 0xb0, 0xba, 0x11, 0x4f, 0x06,
	0x86, 0x28, 0x42, 0xc9, 0xb4, 0xf0, 0xec, 0xc4, 0xe2, 0x19, 0xb8, 0x1b, 0x26, 0x6a, 0xcd, 0xc3,
	0x7f, 0x11, 0x4a, 0x8e, 0x8a, 0x51, 0xc5, 0x97, 0xe0, 0x9b, 0xa8, 0x6f, 0x9f, 0x33, 0xf0, 0x3a,
	0x29, 0xf5, 0x4f, 0x9b, 0x63, 0x33, 0x30, 0x30, 0xbe, 0x85, 0x93, 0x25, 0xd3, 0x94, 0xe7, 0xe6,
	0xaa, 0xaf, 0xf5, 0xfe, 0x98, 0xbc, 0x00, 0x30, 0xc9, 0xd6, 0x4a, 0x05, 0x73, 0x70, 0x6d, 0x55,
	0x87, 0x1b, 0x27, 0xd7, 0x7e, 0x6a, 0x1a, 0x6c, 0x49, 0x6a, 0x79, 0x31, 0xc2, 0x3c, 0x7c, 0xeb,
	0x31, 0xda, 0xf7, 0x18, 0x7d, 0xf4, 0x18, 0xbd, 0x0e, 0xd8, 0xd9, 0x0f, 0xd8, 0x79, 0x1f, 0xb0,
	0x43, 0x5c, 0xdb, 0xf1, 0xe2, 0x33, 0x00, 0x00, 0xff, 0xff, 0x53, 0x74, 0xb5, 0x82, 0xaf, 0x01,
	0x00, 0x00,
}
