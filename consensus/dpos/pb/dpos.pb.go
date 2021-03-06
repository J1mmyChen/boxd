// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dpos.proto

package dpospb

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

type PeriodContext struct {
	Period     []*Period `protobuf:"bytes,1,rep,name=period" json:"period,omitempty"`
	NextPeriod []*Period `protobuf:"bytes,2,rep,name=next_period,json=nextPeriod" json:"next_period,omitempty"`
}

func (m *PeriodContext) Reset()         { *m = PeriodContext{} }
func (m *PeriodContext) String() string { return proto.CompactTextString(m) }
func (*PeriodContext) ProtoMessage()    {}
func (*PeriodContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_dpos_4f5a372154c1a2c0, []int{0}
}
func (m *PeriodContext) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeriodContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeriodContext.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *PeriodContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeriodContext.Merge(dst, src)
}
func (m *PeriodContext) XXX_Size() int {
	return m.Size()
}
func (m *PeriodContext) XXX_DiscardUnknown() {
	xxx_messageInfo_PeriodContext.DiscardUnknown(m)
}

var xxx_messageInfo_PeriodContext proto.InternalMessageInfo

func (m *PeriodContext) GetPeriod() []*Period {
	if m != nil {
		return m.Period
	}
	return nil
}

func (m *PeriodContext) GetNextPeriod() []*Period {
	if m != nil {
		return m.NextPeriod
	}
	return nil
}

type Period struct {
	Addr   []byte `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	PeerId string `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (m *Period) Reset()         { *m = Period{} }
func (m *Period) String() string { return proto.CompactTextString(m) }
func (*Period) ProtoMessage()    {}
func (*Period) Descriptor() ([]byte, []int) {
	return fileDescriptor_dpos_4f5a372154c1a2c0, []int{1}
}
func (m *Period) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Period) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Period.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Period) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Period.Merge(dst, src)
}
func (m *Period) XXX_Size() int {
	return m.Size()
}
func (m *Period) XXX_DiscardUnknown() {
	xxx_messageInfo_Period.DiscardUnknown(m)
}

var xxx_messageInfo_Period proto.InternalMessageInfo

func (m *Period) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *Period) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

type CandidateContext struct {
	Candidates []*Candidate `protobuf:"bytes,2,rep,name=candidates" json:"candidates,omitempty"`
}

func (m *CandidateContext) Reset()         { *m = CandidateContext{} }
func (m *CandidateContext) String() string { return proto.CompactTextString(m) }
func (*CandidateContext) ProtoMessage()    {}
func (*CandidateContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_dpos_4f5a372154c1a2c0, []int{2}
}
func (m *CandidateContext) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CandidateContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CandidateContext.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CandidateContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CandidateContext.Merge(dst, src)
}
func (m *CandidateContext) XXX_Size() int {
	return m.Size()
}
func (m *CandidateContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CandidateContext.DiscardUnknown(m)
}

var xxx_messageInfo_CandidateContext proto.InternalMessageInfo

func (m *CandidateContext) GetCandidates() []*Candidate {
	if m != nil {
		return m.Candidates
	}
	return nil
}

type Candidate struct {
	Addr  []byte `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Votes int64  `protobuf:"varint,2,opt,name=votes,proto3" json:"votes,omitempty"`
	Peer  string `protobuf:"bytes,3,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (m *Candidate) Reset()         { *m = Candidate{} }
func (m *Candidate) String() string { return proto.CompactTextString(m) }
func (*Candidate) ProtoMessage()    {}
func (*Candidate) Descriptor() ([]byte, []int) {
	return fileDescriptor_dpos_4f5a372154c1a2c0, []int{3}
}
func (m *Candidate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Candidate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Candidate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Candidate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candidate.Merge(dst, src)
}
func (m *Candidate) XXX_Size() int {
	return m.Size()
}
func (m *Candidate) XXX_DiscardUnknown() {
	xxx_messageInfo_Candidate.DiscardUnknown(m)
}

var xxx_messageInfo_Candidate proto.InternalMessageInfo

func (m *Candidate) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *Candidate) GetVotes() int64 {
	if m != nil {
		return m.Votes
	}
	return 0
}

func (m *Candidate) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

type EternalBlockMsg struct {
	Hash      []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *EternalBlockMsg) Reset()         { *m = EternalBlockMsg{} }
func (m *EternalBlockMsg) String() string { return proto.CompactTextString(m) }
func (*EternalBlockMsg) ProtoMessage()    {}
func (*EternalBlockMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_dpos_4f5a372154c1a2c0, []int{4}
}
func (m *EternalBlockMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EternalBlockMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EternalBlockMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EternalBlockMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EternalBlockMsg.Merge(dst, src)
}
func (m *EternalBlockMsg) XXX_Size() int {
	return m.Size()
}
func (m *EternalBlockMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_EternalBlockMsg.DiscardUnknown(m)
}

var xxx_messageInfo_EternalBlockMsg proto.InternalMessageInfo

func (m *EternalBlockMsg) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *EternalBlockMsg) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *EternalBlockMsg) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*PeriodContext)(nil), "dpospb.PeriodContext")
	proto.RegisterType((*Period)(nil), "dpospb.Period")
	proto.RegisterType((*CandidateContext)(nil), "dpospb.candidateContext")
	proto.RegisterType((*Candidate)(nil), "dpospb.Candidate")
	proto.RegisterType((*EternalBlockMsg)(nil), "dpospb.EternalBlockMsg")
}
func (m *PeriodContext) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeriodContext) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Period) > 0 {
		for _, msg := range m.Period {
			dAtA[i] = 0xa
			i++
			i = encodeVarintDpos(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.NextPeriod) > 0 {
		for _, msg := range m.NextPeriod {
			dAtA[i] = 0x12
			i++
			i = encodeVarintDpos(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Period) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Period) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDpos(dAtA, i, uint64(len(m.Addr)))
		i += copy(dAtA[i:], m.Addr)
	}
	if len(m.PeerId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDpos(dAtA, i, uint64(len(m.PeerId)))
		i += copy(dAtA[i:], m.PeerId)
	}
	return i, nil
}

func (m *CandidateContext) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CandidateContext) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Candidates) > 0 {
		for _, msg := range m.Candidates {
			dAtA[i] = 0x12
			i++
			i = encodeVarintDpos(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Candidate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Candidate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Addr) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDpos(dAtA, i, uint64(len(m.Addr)))
		i += copy(dAtA[i:], m.Addr)
	}
	if m.Votes != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDpos(dAtA, i, uint64(m.Votes))
	}
	if len(m.Peer) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDpos(dAtA, i, uint64(len(m.Peer)))
		i += copy(dAtA[i:], m.Peer)
	}
	return i, nil
}

func (m *EternalBlockMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EternalBlockMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDpos(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDpos(dAtA, i, uint64(m.Timestamp))
	}
	if len(m.Signature) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDpos(dAtA, i, uint64(len(m.Signature)))
		i += copy(dAtA[i:], m.Signature)
	}
	return i, nil
}

func encodeVarintDpos(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PeriodContext) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Period) > 0 {
		for _, e := range m.Period {
			l = e.Size()
			n += 1 + l + sovDpos(uint64(l))
		}
	}
	if len(m.NextPeriod) > 0 {
		for _, e := range m.NextPeriod {
			l = e.Size()
			n += 1 + l + sovDpos(uint64(l))
		}
	}
	return n
}

func (m *Period) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovDpos(uint64(l))
	}
	l = len(m.PeerId)
	if l > 0 {
		n += 1 + l + sovDpos(uint64(l))
	}
	return n
}

func (m *CandidateContext) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Candidates) > 0 {
		for _, e := range m.Candidates {
			l = e.Size()
			n += 1 + l + sovDpos(uint64(l))
		}
	}
	return n
}

func (m *Candidate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovDpos(uint64(l))
	}
	if m.Votes != 0 {
		n += 1 + sovDpos(uint64(m.Votes))
	}
	l = len(m.Peer)
	if l > 0 {
		n += 1 + l + sovDpos(uint64(l))
	}
	return n
}

func (m *EternalBlockMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovDpos(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovDpos(uint64(m.Timestamp))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovDpos(uint64(l))
	}
	return n
}

func sovDpos(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDpos(x uint64) (n int) {
	return sovDpos(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PeriodContext) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDpos
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
			return fmt.Errorf("proto: PeriodContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeriodContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpos
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
				return ErrInvalidLengthDpos
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Period = append(m.Period, &Period{})
			if err := m.Period[len(m.Period)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpos
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
				return ErrInvalidLengthDpos
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextPeriod = append(m.NextPeriod, &Period{})
			if err := m.NextPeriod[len(m.NextPeriod)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDpos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDpos
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
func (m *Period) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDpos
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
			return fmt.Errorf("proto: Period: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Period: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpos
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
				return ErrInvalidLengthDpos
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = append(m.Addr[:0], dAtA[iNdEx:postIndex]...)
			if m.Addr == nil {
				m.Addr = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpos
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
				return ErrInvalidLengthDpos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDpos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDpos
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
func (m *CandidateContext) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDpos
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
			return fmt.Errorf("proto: candidateContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: candidateContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Candidates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpos
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
				return ErrInvalidLengthDpos
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Candidates = append(m.Candidates, &Candidate{})
			if err := m.Candidates[len(m.Candidates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDpos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDpos
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
func (m *Candidate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDpos
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
			return fmt.Errorf("proto: Candidate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Candidate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpos
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
				return ErrInvalidLengthDpos
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = append(m.Addr[:0], dAtA[iNdEx:postIndex]...)
			if m.Addr == nil {
				m.Addr = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			m.Votes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Votes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpos
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
				return ErrInvalidLengthDpos
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDpos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDpos
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
func (m *EternalBlockMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDpos
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
			return fmt.Errorf("proto: EternalBlockMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EternalBlockMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpos
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
				return ErrInvalidLengthDpos
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpos
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDpos
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
				return ErrInvalidLengthDpos
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDpos(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDpos
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
func skipDpos(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDpos
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
					return 0, ErrIntOverflowDpos
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
					return 0, ErrIntOverflowDpos
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
				return 0, ErrInvalidLengthDpos
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDpos
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
				next, err := skipDpos(dAtA[start:])
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
	ErrInvalidLengthDpos = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDpos   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("dpos.proto", fileDescriptor_dpos_4f5a372154c1a2c0) }

var fileDescriptor_dpos_4f5a372154c1a2c0 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbf, 0x6a, 0xeb, 0x30,
	0x18, 0xc5, 0xad, 0xeb, 0x5b, 0x17, 0x7f, 0x49, 0xff, 0x89, 0x42, 0x3d, 0x14, 0x11, 0x3c, 0x14,
	0x4f, 0x2e, 0x6d, 0xe9, 0x0b, 0x24, 0x64, 0xc8, 0x50, 0x28, 0x7e, 0x81, 0xa0, 0x44, 0x22, 0x36,
	0x4d, 0x2c, 0x21, 0xa9, 0x25, 0x8f, 0xd1, 0xc7, 0xea, 0x98, 0xb1, 0x63, 0xb1, 0x5f, 0xa4, 0x48,
	0xb2, 0x93, 0x0e, 0xd9, 0x8e, 0xce, 0xf9, 0x71, 0xbe, 0x03, 0x02, 0x60, 0x52, 0xe8, 0x5c, 0x2a,
	0x61, 0x04, 0x8e, 0xac, 0x96, 0x8b, 0xb4, 0x84, 0xb3, 0x57, 0xae, 0x2a, 0xc1, 0x26, 0xa2, 0x36,
	0x7c, 0x6b, 0xf0, 0x1d, 0x44, 0xd2, 0x19, 0x09, 0x1a, 0x85, 0xd9, 0xe0, 0xf1, 0x3c, 0xf7, 0x64,
	0xee, 0xb1, 0xa2, 0x4b, 0xf1, 0x3d, 0x0c, 0x6a, 0xbe, 0x35, 0xf3, 0x0e, 0xfe, 0x77, 0x14, 0x06,
	0x8b, 0x78, 0x9d, 0x3e, 0x43, 0xe4, 0x15, 0xc6, 0xf0, 0x9f, 0x32, 0xa6, 0x12, 0x34, 0x42, 0xd9,
	0xb0, 0x70, 0x1a, 0xdf, 0xc0, 0xa9, 0xe4, 0x5c, 0xcd, 0x2b, 0x5b, 0x85, 0xb2, 0xd8, 0xde, 0xe1,
	0x6a, 0xc6, 0xd2, 0x29, 0x5c, 0x2e, 0x69, 0xcd, 0x2a, 0x46, 0x0d, 0xef, 0x37, 0x3e, 0x00, 0xec,
	0x3d, 0xdd, 0x9d, 0xbe, 0xea, 0x4f, 0x4f, 0xfa, 0xa4, 0xf8, 0x03, 0xa5, 0x33, 0x88, 0xf7, 0xc1,
	0xd1, 0x01, 0xd7, 0x70, 0xf2, 0x21, 0x7c, 0x1d, 0xca, 0xc2, 0xc2, 0x3f, 0x2c, 0x69, 0x77, 0x24,
	0xa1, 0xdb, 0xe4, 0x74, 0x4a, 0xe1, 0x62, 0x6a, 0xb8, 0xaa, 0xe9, 0x7a, 0xbc, 0x16, 0xcb, 0xb7,
	0x17, 0xbd, 0xb2, 0x58, 0x49, 0x75, 0xd9, 0x17, 0x5a, 0x8d, 0x6f, 0x21, 0x36, 0xd5, 0x86, 0x6b,
	0x43, 0x37, 0xb2, 0x2b, 0x3d, 0x18, 0x36, 0xd5, 0xd5, 0xaa, 0xa6, 0xe6, 0x5d, 0x71, 0xd7, 0x3e,
	0x2c, 0x0e, 0xc6, 0x38, 0xf9, 0x6a, 0x08, 0xda, 0x35, 0x04, 0xfd, 0x34, 0x04, 0x7d, 0xb6, 0x24,
	0xd8, 0xb5, 0x24, 0xf8, 0x6e, 0x49, 0xb0, 0x88, 0xdc, 0xf7, 0x3d, 0xfd, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x40, 0xa3, 0x1f, 0xbb, 0xcc, 0x01, 0x00, 0x00,
}
