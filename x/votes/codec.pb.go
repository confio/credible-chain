// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/votes/codec.proto

package votes

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Vote struct {
	MainVote             int32    `protobuf:"varint,1,opt,name=main_vote,json=mainVote,proto3" json:"main_vote,omitempty"`
	RepVote              string   `protobuf:"bytes,2,opt,name=rep_vote,json=repVote,proto3" json:"rep_vote,omitempty"`
	Charity              string   `protobuf:"bytes,3,opt,name=charity,proto3" json:"charity,omitempty"`
	PostCode             string   `protobuf:"bytes,4,opt,name=postCode,proto3" json:"postCode,omitempty"`
	BirthYear            int32    `protobuf:"varint,5,opt,name=birth_year,json=birthYear,proto3" json:"birth_year,omitempty"`
	Donation             int32    `protobuf:"varint,6,opt,name=donation,proto3" json:"donation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_codec_a43bd61b6f86adea, []int{0}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(dst, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetMainVote() int32 {
	if m != nil {
		return m.MainVote
	}
	return 0
}

func (m *Vote) GetRepVote() string {
	if m != nil {
		return m.RepVote
	}
	return ""
}

func (m *Vote) GetCharity() string {
	if m != nil {
		return m.Charity
	}
	return ""
}

func (m *Vote) GetPostCode() string {
	if m != nil {
		return m.PostCode
	}
	return ""
}

func (m *Vote) GetBirthYear() int32 {
	if m != nil {
		return m.BirthYear
	}
	return 0
}

func (m *Vote) GetDonation() int32 {
	if m != nil {
		return m.Donation
	}
	return 0
}

type VoteRecord struct {
	Vote                 *Vote      `protobuf:"bytes,1,opt,name=vote" json:"vote,omitempty"`
	Identifier           string     `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	SmsCode              string     `protobuf:"bytes,3,opt,name=sms_code,json=smsCode,proto3" json:"sms_code,omitempty"`
	VotedAt              *time.Time `protobuf:"bytes,5,opt,name=voted_at,json=votedAt,stdtime" json:"voted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *VoteRecord) Reset()         { *m = VoteRecord{} }
func (m *VoteRecord) String() string { return proto.CompactTextString(m) }
func (*VoteRecord) ProtoMessage()    {}
func (*VoteRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_codec_a43bd61b6f86adea, []int{1}
}
func (m *VoteRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *VoteRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteRecord.Merge(dst, src)
}
func (m *VoteRecord) XXX_Size() int {
	return m.Size()
}
func (m *VoteRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteRecord.DiscardUnknown(m)
}

var xxx_messageInfo_VoteRecord proto.InternalMessageInfo

func (m *VoteRecord) GetVote() *Vote {
	if m != nil {
		return m.Vote
	}
	return nil
}

func (m *VoteRecord) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *VoteRecord) GetSmsCode() string {
	if m != nil {
		return m.SmsCode
	}
	return ""
}

func (m *VoteRecord) GetVotedAt() *time.Time {
	if m != nil {
		return m.VotedAt
	}
	return nil
}

type Tally struct {
	Option               string   `protobuf:"bytes,1,opt,name=option,proto3" json:"option,omitempty"`
	Total                int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tally) Reset()         { *m = Tally{} }
func (m *Tally) String() string { return proto.CompactTextString(m) }
func (*Tally) ProtoMessage()    {}
func (*Tally) Descriptor() ([]byte, []int) {
	return fileDescriptor_codec_a43bd61b6f86adea, []int{2}
}
func (m *Tally) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tally) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tally.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Tally) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tally.Merge(dst, src)
}
func (m *Tally) XXX_Size() int {
	return m.Size()
}
func (m *Tally) XXX_DiscardUnknown() {
	xxx_messageInfo_Tally.DiscardUnknown(m)
}

var xxx_messageInfo_Tally proto.InternalMessageInfo

func (m *Tally) GetOption() string {
	if m != nil {
		return m.Option
	}
	return ""
}

func (m *Tally) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*Vote)(nil), "votes.Vote")
	proto.RegisterType((*VoteRecord)(nil), "votes.VoteRecord")
	proto.RegisterType((*Tally)(nil), "votes.Tally")
}
func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MainVote != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.MainVote))
	}
	if len(m.RepVote) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.RepVote)))
		i += copy(dAtA[i:], m.RepVote)
	}
	if len(m.Charity) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Charity)))
		i += copy(dAtA[i:], m.Charity)
	}
	if len(m.PostCode) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.PostCode)))
		i += copy(dAtA[i:], m.PostCode)
	}
	if m.BirthYear != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.BirthYear))
	}
	if m.Donation != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Donation))
	}
	return i, nil
}

func (m *VoteRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteRecord) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Vote != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Vote.Size()))
		n1, err := m.Vote.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Identifier) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Identifier)))
		i += copy(dAtA[i:], m.Identifier)
	}
	if len(m.SmsCode) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.SmsCode)))
		i += copy(dAtA[i:], m.SmsCode)
	}
	if m.VotedAt != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.VotedAt)))
		n2, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.VotedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *Tally) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tally) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Option) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Option)))
		i += copy(dAtA[i:], m.Option)
	}
	if m.Total != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Total))
	}
	return i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Vote) Size() (n int) {
	var l int
	_ = l
	if m.MainVote != 0 {
		n += 1 + sovCodec(uint64(m.MainVote))
	}
	l = len(m.RepVote)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Charity)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.PostCode)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.BirthYear != 0 {
		n += 1 + sovCodec(uint64(m.BirthYear))
	}
	if m.Donation != 0 {
		n += 1 + sovCodec(uint64(m.Donation))
	}
	return n
}

func (m *VoteRecord) Size() (n int) {
	var l int
	_ = l
	if m.Vote != nil {
		l = m.Vote.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.SmsCode)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.VotedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.VotedAt)
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *Tally) Size() (n int) {
	var l int
	_ = l
	l = len(m.Option)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Total != 0 {
		n += 1 + sovCodec(uint64(m.Total))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainVote", wireType)
			}
			m.MainVote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MainVote |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepVote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RepVote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Charity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Charity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BirthYear", wireType)
			}
			m.BirthYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BirthYear |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Donation", wireType)
			}
			m.Donation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Donation |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *VoteRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: VoteRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Vote == nil {
				m.Vote = &Vote{}
			}
			if err := m.Vote.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmsCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SmsCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VotedAt == nil {
				m.VotedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.VotedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func (m *Tally) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: Tally: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tally: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Option", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Option = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
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
				next, err := skipCodec(dAtA[start:])
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
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("x/votes/codec.proto", fileDescriptor_codec_a43bd61b6f86adea) }

var fileDescriptor_codec_a43bd61b6f86adea = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0x41, 0x8e, 0x9b, 0x40,
	0x10, 0xcc, 0xc4, 0x60, 0xe3, 0xf6, 0x25, 0x9a, 0x44, 0x11, 0x21, 0x0a, 0xb6, 0x7c, 0xf2, 0x25,
	0x20, 0x39, 0xca, 0x29, 0xa7, 0x38, 0x3f, 0x40, 0x56, 0xa4, 0x9c, 0xd0, 0x00, 0x63, 0x3c, 0x12,
	0xd0, 0x68, 0x18, 0x47, 0xeb, 0x5f, 0xf8, 0x13, 0x7b, 0xde, 0x6f, 0xec, 0x71, 0x7f, 0xb0, 0x2b,
	0xef, 0x47, 0x56, 0xd3, 0x18, 0xef, 0xde, 0xa6, 0xaa, 0x7a, 0xba, 0xaa, 0x0b, 0x3e, 0xde, 0xc4,
	0xff, 0xd1, 0xc8, 0x2e, 0xce, 0xb1, 0x90, 0x79, 0xd4, 0x6a, 0x34, 0xc8, 0x5d, 0xa2, 0x82, 0x79,
	0x89, 0x58, 0x56, 0x32, 0x26, 0x32, 0x3b, 0xec, 0x62, 0xa3, 0x6a, 0xd9, 0x19, 0x51, 0xb7, 0xfd,
	0x5c, 0xf0, 0xbd, 0x54, 0x66, 0x7f, 0xc8, 0xa2, 0x1c, 0xeb, 0xb8, 0xc4, 0x12, 0x5f, 0x27, 0x2d,
	0x22, 0x40, 0xaf, 0x7e, 0x7c, 0x79, 0xc7, 0xc0, 0xf9, 0x8b, 0x46, 0xf2, 0xaf, 0x30, 0xad, 0x85,
	0x6a, 0x52, 0x6b, 0xe3, 0xb3, 0x05, 0x5b, 0xb9, 0x89, 0x67, 0x09, 0x12, 0xbf, 0x80, 0xa7, 0x65,
	0xdb, 0x6b, 0xef, 0x17, 0x6c, 0x35, 0x4d, 0x26, 0x5a, 0xb6, 0x24, 0xf9, 0x30, 0xc9, 0xf7, 0x42,
	0x2b, 0x73, 0xf4, 0x47, 0xbd, 0x72, 0x81, 0x3c, 0x00, 0xaf, 0xc5, 0xce, 0xfc, 0xc1, 0x42, 0xfa,
	0x0e, 0x49, 0x57, 0xcc, 0xbf, 0x01, 0x64, 0x4a, 0x9b, 0x7d, 0x7a, 0x94, 0x42, 0xfb, 0x2e, 0xd9,
	0x4d, 0x89, 0xf9, 0x27, 0x85, 0xb6, 0x5f, 0x0b, 0x6c, 0x84, 0x51, 0xd8, 0xf8, 0xe3, 0x3e, 0xcb,
	0x80, 0x97, 0xb7, 0x0c, 0xc0, 0x3a, 0x27, 0x32, 0x47, 0x5d, 0xf0, 0x39, 0x38, 0xd7, 0xc8, 0xb3,
	0xf5, 0x2c, 0xa2, 0x9a, 0x22, 0x1a, 0x20, 0x81, 0x87, 0x00, 0xaa, 0x90, 0x8d, 0x51, 0x3b, 0x25,
	0xf5, 0x25, 0xfd, 0x1b, 0xc6, 0xde, 0xd6, 0xd5, 0x5d, 0x6a, 0xbb, 0x1e, 0x2e, 0xe8, 0xea, 0x8e,
	0x52, 0xfe, 0x02, 0xcf, 0xae, 0x28, 0x52, 0x61, 0x28, 0xe3, 0x6c, 0x1d, 0x44, 0x7d, 0xff, 0xd1,
	0xd0, 0x6a, 0xb4, 0x1d, 0xfa, 0xdf, 0x38, 0xa7, 0xc7, 0x39, 0x4b, 0x26, 0xf4, 0xe3, 0xb7, 0x59,
	0xfe, 0x04, 0x77, 0x2b, 0xaa, 0xea, 0xc8, 0x3f, 0xc3, 0x18, 0x5b, 0x3a, 0x85, 0xd1, 0xfa, 0x0b,
	0xe2, 0x9f, 0xc0, 0x35, 0x68, 0x44, 0x45, 0x99, 0x46, 0x49, 0x0f, 0x36, 0x1f, 0xee, 0xcf, 0x21,
	0x7b, 0x38, 0x87, 0xec, 0xe9, 0x1c, 0xb2, 0xd3, 0x73, 0xf8, 0x2e, 0x1b, 0x93, 0xd7, 0x8f, 0x97,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0x5a, 0x98, 0xb6, 0x17, 0x02, 0x00, 0x00,
}
