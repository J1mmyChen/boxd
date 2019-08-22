// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: receipt.proto

package corepb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
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

type Receipt struct {
	TxHash   []byte `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	TxIndex  uint32 `protobuf:"varint,2,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	Deployed bool   `protobuf:"varint,3,opt,name=deployed,proto3" json:"deployed,omitempty"`
	Failed   bool   `protobuf:"varint,4,opt,name=failed,proto3" json:"failed,omitempty"`
	GasUsed  uint64 `protobuf:"varint,5,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	Logs     []*Log `protobuf:"bytes,6,rep,name=logs,proto3" json:"logs,omitempty"`
	Bloom    []byte `protobuf:"bytes,7,opt,name=bloom,proto3" json:"bloom,omitempty"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_ace1d6eb38fad2c8, []int{0}
}
func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(m, src)
}
func (m *Receipt) XXX_Size() int {
	return m.Size()
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Receipt) GetTxIndex() uint32 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

func (m *Receipt) GetDeployed() bool {
	if m != nil {
		return m.Deployed
	}
	return false
}

func (m *Receipt) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

func (m *Receipt) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *Receipt) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Receipt) GetBloom() []byte {
	if m != nil {
		return m.Bloom
	}
	return nil
}

type Receipts struct {
	Receipts []*Receipt `protobuf:"bytes,1,rep,name=receipts,proto3" json:"receipts,omitempty"`
}

func (m *Receipts) Reset()         { *m = Receipts{} }
func (m *Receipts) String() string { return proto.CompactTextString(m) }
func (*Receipts) ProtoMessage()    {}
func (*Receipts) Descriptor() ([]byte, []int) {
	return fileDescriptor_ace1d6eb38fad2c8, []int{1}
}
func (m *Receipts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Receipts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Receipts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Receipts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipts.Merge(m, src)
}
func (m *Receipts) XXX_Size() int {
	return m.Size()
}
func (m *Receipts) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipts.DiscardUnknown(m)
}

var xxx_messageInfo_Receipts proto.InternalMessageInfo

func (m *Receipts) GetReceipts() []*Receipt {
	if m != nil {
		return m.Receipts
	}
	return nil
}

type HashReceipt struct {
	TxHash   []byte     `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	TxIndex  uint32     `protobuf:"varint,2,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	Deployed bool       `protobuf:"varint,3,opt,name=deployed,proto3" json:"deployed,omitempty"`
	Failed   bool       `protobuf:"varint,4,opt,name=failed,proto3" json:"failed,omitempty"`
	GasUsed  uint64     `protobuf:"varint,5,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	Logs     []*HashLog `protobuf:"bytes,6,rep,name=logs,proto3" json:"logs,omitempty"`
	Bloom    []byte     `protobuf:"bytes,7,opt,name=bloom,proto3" json:"bloom,omitempty"`
}

func (m *HashReceipt) Reset()         { *m = HashReceipt{} }
func (m *HashReceipt) String() string { return proto.CompactTextString(m) }
func (*HashReceipt) ProtoMessage()    {}
func (*HashReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_ace1d6eb38fad2c8, []int{2}
}
func (m *HashReceipt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HashReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HashReceipt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HashReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashReceipt.Merge(m, src)
}
func (m *HashReceipt) XXX_Size() int {
	return m.Size()
}
func (m *HashReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_HashReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_HashReceipt proto.InternalMessageInfo

func (m *HashReceipt) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *HashReceipt) GetTxIndex() uint32 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

func (m *HashReceipt) GetDeployed() bool {
	if m != nil {
		return m.Deployed
	}
	return false
}

func (m *HashReceipt) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

func (m *HashReceipt) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *HashReceipt) GetLogs() []*HashLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *HashReceipt) GetBloom() []byte {
	if m != nil {
		return m.Bloom
	}
	return nil
}

type HashReceipts struct {
	Receipts []*HashReceipt `protobuf:"bytes,1,rep,name=receipts,proto3" json:"receipts,omitempty"`
}

func (m *HashReceipts) Reset()         { *m = HashReceipts{} }
func (m *HashReceipts) String() string { return proto.CompactTextString(m) }
func (*HashReceipts) ProtoMessage()    {}
func (*HashReceipts) Descriptor() ([]byte, []int) {
	return fileDescriptor_ace1d6eb38fad2c8, []int{3}
}
func (m *HashReceipts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HashReceipts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HashReceipts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HashReceipts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashReceipts.Merge(m, src)
}
func (m *HashReceipts) XXX_Size() int {
	return m.Size()
}
func (m *HashReceipts) XXX_DiscardUnknown() {
	xxx_messageInfo_HashReceipts.DiscardUnknown(m)
}

var xxx_messageInfo_HashReceipts proto.InternalMessageInfo

func (m *HashReceipts) GetReceipts() []*HashReceipt {
	if m != nil {
		return m.Receipts
	}
	return nil
}

func init() {
	proto.RegisterType((*Receipt)(nil), "corepb.Receipt")
	proto.RegisterType((*Receipts)(nil), "corepb.Receipts")
	proto.RegisterType((*HashReceipt)(nil), "corepb.HashReceipt")
	proto.RegisterType((*HashReceipts)(nil), "corepb.HashReceipts")
}

func init() { proto.RegisterFile("receipt.proto", fileDescriptor_ace1d6eb38fad2c8) }

var fileDescriptor_ace1d6eb38fad2c8 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0x3f, 0x4b, 0xf4, 0x40,
	0x18, 0xc4, 0x6f, 0xdf, 0xcb, 0x25, 0x79, 0x9f, 0xdc, 0x21, 0xac, 0xa2, 0xeb, 0x15, 0x6b, 0x88,
	0x4d, 0x40, 0x88, 0xa0, 0x85, 0xa5, 0x60, 0xa5, 0x60, 0xb5, 0x60, 0x1d, 0x92, 0xdb, 0x35, 0x09,
	0x44, 0x37, 0x64, 0x57, 0x88, 0xdf, 0xc2, 0xef, 0x64, 0x23, 0xd8, 0x5c, 0x69, 0x29, 0xc9, 0x17,
	0x91, 0xfc, 0x31, 0x9c, 0x5c, 0x61, 0x6b, 0x39, 0xfb, 0x0c, 0xc3, 0xcc, 0x8f, 0x85, 0x45, 0x29,
	0x56, 0x22, 0x2b, 0x74, 0x50, 0x94, 0x52, 0x4b, 0x6c, 0xae, 0x64, 0x29, 0x8a, 0x78, 0xf9, 0x3f,
	0x97, 0x49, 0xff, 0xe4, 0xbd, 0x22, 0xb0, 0x58, 0x6f, 0xc2, 0x07, 0x60, 0xe9, 0x2a, 0x4c, 0x23,
	0x95, 0x12, 0xe4, 0x22, 0x7f, 0xce, 0x4c, 0x5d, 0x5d, 0x47, 0x2a, 0xc5, 0x87, 0x60, 0xeb, 0x2a,
	0xcc, 0x1e, 0xb9, 0xa8, 0xc8, 0x3f, 0x17, 0xf9, 0x0b, 0x66, 0xe9, 0xea, 0xa6, 0x95, 0x78, 0x09,
	0x36, 0x17, 0x45, 0x2e, 0x9f, 0x05, 0x27, 0x53, 0x17, 0xf9, 0x36, 0x1b, 0x35, 0xde, 0x07, 0xf3,
	0x3e, 0xca, 0x72, 0xc1, 0x89, 0xd1, 0x5d, 0x06, 0xd5, 0xc6, 0x25, 0x91, 0x0a, 0x9f, 0x94, 0xe0,
	0x64, 0xe6, 0x22, 0xdf, 0x60, 0x56, 0x12, 0xa9, 0x3b, 0x25, 0x38, 0x3e, 0x02, 0x23, 0x97, 0x89,
	0x22, 0xa6, 0x3b, 0xf5, 0x9d, 0x33, 0x27, 0xe8, 0x0b, 0x07, 0xb7, 0x32, 0x61, 0xdd, 0x01, 0xef,
	0xc1, 0x2c, 0xce, 0xa5, 0x7c, 0x20, 0x56, 0xd7, 0xb0, 0x17, 0xde, 0x05, 0xd8, 0xc3, 0x08, 0x85,
	0x4f, 0xc0, 0x1e, 0x56, 0x2b, 0x82, 0xba, 0x98, 0x9d, 0xef, 0x98, 0xc1, 0xc3, 0x46, 0x83, 0xf7,
	0x8e, 0xc0, 0x69, 0x27, 0xfe, 0x21, 0x04, 0xc7, 0x3f, 0x10, 0x8c, 0xdd, 0xdb, 0x16, 0xbf, 0x61,
	0xb8, 0x84, 0xf9, 0xc6, 0x18, 0x85, 0x4f, 0xb7, 0x50, 0xec, 0x6e, 0xc6, 0x6d, 0xe1, 0xb8, 0x22,
	0x6f, 0x35, 0x45, 0xeb, 0x9a, 0xa2, 0xcf, 0x9a, 0xa2, 0x97, 0x86, 0x4e, 0xd6, 0x0d, 0x9d, 0x7c,
	0x34, 0x74, 0x12, 0x9b, 0xdd, 0x77, 0x39, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x32, 0x19,
	0x3d, 0x52, 0x02, 0x00, 0x00,
}

func (m *Receipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Receipt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.TxHash)))
		i += copy(dAtA[i:], m.TxHash)
	}
	if m.TxIndex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintReceipt(dAtA, i, uint64(m.TxIndex))
	}
	if m.Deployed {
		dAtA[i] = 0x18
		i++
		if m.Deployed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Failed {
		dAtA[i] = 0x20
		i++
		if m.Failed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.GasUsed != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintReceipt(dAtA, i, uint64(m.GasUsed))
	}
	if len(m.Logs) > 0 {
		for _, msg := range m.Logs {
			dAtA[i] = 0x32
			i++
			i = encodeVarintReceipt(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Bloom) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.Bloom)))
		i += copy(dAtA[i:], m.Bloom)
	}
	return i, nil
}

func (m *Receipts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Receipts) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Receipts) > 0 {
		for _, msg := range m.Receipts {
			dAtA[i] = 0xa
			i++
			i = encodeVarintReceipt(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *HashReceipt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashReceipt) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.TxHash)))
		i += copy(dAtA[i:], m.TxHash)
	}
	if m.TxIndex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintReceipt(dAtA, i, uint64(m.TxIndex))
	}
	if m.Deployed {
		dAtA[i] = 0x18
		i++
		if m.Deployed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Failed {
		dAtA[i] = 0x20
		i++
		if m.Failed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.GasUsed != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintReceipt(dAtA, i, uint64(m.GasUsed))
	}
	if len(m.Logs) > 0 {
		for _, msg := range m.Logs {
			dAtA[i] = 0x32
			i++
			i = encodeVarintReceipt(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Bloom) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintReceipt(dAtA, i, uint64(len(m.Bloom)))
		i += copy(dAtA[i:], m.Bloom)
	}
	return i, nil
}

func (m *HashReceipts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashReceipts) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Receipts) > 0 {
		for _, msg := range m.Receipts {
			dAtA[i] = 0xa
			i++
			i = encodeVarintReceipt(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintReceipt(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Receipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	if m.TxIndex != 0 {
		n += 1 + sovReceipt(uint64(m.TxIndex))
	}
	if m.Deployed {
		n += 2
	}
	if m.Failed {
		n += 2
	}
	if m.GasUsed != 0 {
		n += 1 + sovReceipt(uint64(m.GasUsed))
	}
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovReceipt(uint64(l))
		}
	}
	l = len(m.Bloom)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	return n
}

func (m *Receipts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Receipts) > 0 {
		for _, e := range m.Receipts {
			l = e.Size()
			n += 1 + l + sovReceipt(uint64(l))
		}
	}
	return n
}

func (m *HashReceipt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	if m.TxIndex != 0 {
		n += 1 + sovReceipt(uint64(m.TxIndex))
	}
	if m.Deployed {
		n += 2
	}
	if m.Failed {
		n += 2
	}
	if m.GasUsed != 0 {
		n += 1 + sovReceipt(uint64(m.GasUsed))
	}
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovReceipt(uint64(l))
		}
	}
	l = len(m.Bloom)
	if l > 0 {
		n += 1 + l + sovReceipt(uint64(l))
	}
	return n
}

func (m *HashReceipts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Receipts) > 0 {
		for _, e := range m.Receipts {
			l = e.Size()
			n += 1 + l + sovReceipt(uint64(l))
		}
	}
	return n
}

func sovReceipt(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozReceipt(x uint64) (n int) {
	return sovReceipt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Receipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReceipt
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
			return fmt.Errorf("proto: Receipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
			}
			m.TxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxIndex |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deployed = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Failed = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &Log{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bloom", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bloom = append(m.Bloom[:0], dAtA[iNdEx:postIndex]...)
			if m.Bloom == nil {
				m.Bloom = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReceipt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReceipt
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
func (m *Receipts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReceipt
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
			return fmt.Errorf("proto: Receipts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Receipts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receipts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receipts = append(m.Receipts, &Receipt{})
			if err := m.Receipts[len(m.Receipts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReceipt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReceipt
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
func (m *HashReceipt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReceipt
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
			return fmt.Errorf("proto: HashReceipt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashReceipt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
			}
			m.TxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxIndex |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deployed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deployed = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Failed = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &HashLog{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bloom", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bloom = append(m.Bloom[:0], dAtA[iNdEx:postIndex]...)
			if m.Bloom == nil {
				m.Bloom = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReceipt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReceipt
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
func (m *HashReceipts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReceipt
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
			return fmt.Errorf("proto: HashReceipts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashReceipts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receipts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReceipt
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
				return ErrInvalidLengthReceipt
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receipts = append(m.Receipts, &HashReceipt{})
			if err := m.Receipts[len(m.Receipts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReceipt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReceipt
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
func skipReceipt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReceipt
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
					return 0, ErrIntOverflowReceipt
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
					return 0, ErrIntOverflowReceipt
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
				return 0, ErrInvalidLengthReceipt
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowReceipt
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
				next, err := skipReceipt(dAtA[start:])
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
	ErrInvalidLengthReceipt = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReceipt   = fmt.Errorf("proto: integer overflow")
)
