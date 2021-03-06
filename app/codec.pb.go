// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/codec.proto

package credchain

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import votes "github.com/confio/credible-chain/x/votes"
import _ "github.com/gogo/protobuf/gogoproto"
import multisig "github.com/iov-one/weave/x/multisig"
import sigs "github.com/iov-one/weave/x/sigs"

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

// Tx contains the message.
//
// When extending Tx, follow the rules:
// - range 1-50 is reserved for middlewares,
// - range 51-inf is reserved for different message types,
// - keep the same numbers for the same message types in both bcpd and bnsd
//   applications. For example, FeeInfo field is used by both and indexed at
//   first position. Skip unused fields (leave index unused or comment out for
//   clarity).
type Tx struct {
	Signatures []*sigs.StdSignature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
	// ID of a multisig contract.
	Multisig [][]byte `protobuf:"bytes,4,rep,name=multisig" json:"multisig,omitempty"`
	// msg is a sum type over all allowed messages on this chain.
	//
	// Types that are valid to be assigned to Sum:
	//	*Tx_CreateContractMsg
	//	*Tx_UpdateContractMsg
	//	*Tx_RecordVoteMsg
	Sum                  isTx_Sum `protobuf_oneof:"sum"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_codec_4d50eb422ace2421, []int{0}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(dst, src)
}
func (m *Tx) XXX_Size() int {
	return m.Size()
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

type isTx_Sum interface {
	isTx_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Tx_CreateContractMsg struct {
	CreateContractMsg *multisig.CreateContractMsg `protobuf:"bytes,56,opt,name=create_contract_msg,json=createContractMsg,oneof"`
}
type Tx_UpdateContractMsg struct {
	UpdateContractMsg *multisig.UpdateContractMsg `protobuf:"bytes,57,opt,name=update_contract_msg,json=updateContractMsg,oneof"`
}
type Tx_RecordVoteMsg struct {
	RecordVoteMsg *votes.VoteRecord `protobuf:"bytes,100,opt,name=record_vote_msg,json=recordVoteMsg,oneof"`
}

func (*Tx_CreateContractMsg) isTx_Sum() {}
func (*Tx_UpdateContractMsg) isTx_Sum() {}
func (*Tx_RecordVoteMsg) isTx_Sum()     {}

func (m *Tx) GetSum() isTx_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Tx) GetSignatures() []*sigs.StdSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Tx) GetMultisig() [][]byte {
	if m != nil {
		return m.Multisig
	}
	return nil
}

func (m *Tx) GetCreateContractMsg() *multisig.CreateContractMsg {
	if x, ok := m.GetSum().(*Tx_CreateContractMsg); ok {
		return x.CreateContractMsg
	}
	return nil
}

func (m *Tx) GetUpdateContractMsg() *multisig.UpdateContractMsg {
	if x, ok := m.GetSum().(*Tx_UpdateContractMsg); ok {
		return x.UpdateContractMsg
	}
	return nil
}

func (m *Tx) GetRecordVoteMsg() *votes.VoteRecord {
	if x, ok := m.GetSum().(*Tx_RecordVoteMsg); ok {
		return x.RecordVoteMsg
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Tx) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Tx_OneofMarshaler, _Tx_OneofUnmarshaler, _Tx_OneofSizer, []interface{}{
		(*Tx_CreateContractMsg)(nil),
		(*Tx_UpdateContractMsg)(nil),
		(*Tx_RecordVoteMsg)(nil),
	}
}

func _Tx_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Tx)
	// sum
	switch x := m.Sum.(type) {
	case *Tx_CreateContractMsg:
		_ = b.EncodeVarint(56<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateContractMsg); err != nil {
			return err
		}
	case *Tx_UpdateContractMsg:
		_ = b.EncodeVarint(57<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateContractMsg); err != nil {
			return err
		}
	case *Tx_RecordVoteMsg:
		_ = b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RecordVoteMsg); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Tx.Sum has unexpected type %T", x)
	}
	return nil
}

func _Tx_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Tx)
	switch tag {
	case 56: // sum.create_contract_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(multisig.CreateContractMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_CreateContractMsg{msg}
		return true, err
	case 57: // sum.update_contract_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(multisig.UpdateContractMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_UpdateContractMsg{msg}
		return true, err
	case 100: // sum.record_vote_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(votes.VoteRecord)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_RecordVoteMsg{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Tx_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Tx)
	// sum
	switch x := m.Sum.(type) {
	case *Tx_CreateContractMsg:
		s := proto.Size(x.CreateContractMsg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_UpdateContractMsg:
		s := proto.Size(x.UpdateContractMsg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_RecordVoteMsg:
		s := proto.Size(x.RecordVoteMsg)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Tx)(nil), "credchain.Tx")
}
func (m *Tx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tx) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for _, msg := range m.Signatures {
			dAtA[i] = 0x12
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Multisig) > 0 {
		for _, b := range m.Multisig {
			dAtA[i] = 0x22
			i++
			i = encodeVarintCodec(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if m.Sum != nil {
		nn1, err := m.Sum.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *Tx_CreateContractMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.CreateContractMsg != nil {
		dAtA[i] = 0xc2
		i++
		dAtA[i] = 0x3
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.CreateContractMsg.Size()))
		n2, err := m.CreateContractMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *Tx_UpdateContractMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.UpdateContractMsg != nil {
		dAtA[i] = 0xca
		i++
		dAtA[i] = 0x3
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.UpdateContractMsg.Size()))
		n3, err := m.UpdateContractMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *Tx_RecordVoteMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.RecordVoteMsg != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.RecordVoteMsg.Size()))
		n4, err := m.RecordVoteMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
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
func (m *Tx) Size() (n int) {
	var l int
	_ = l
	if len(m.Signatures) > 0 {
		for _, e := range m.Signatures {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	if len(m.Multisig) > 0 {
		for _, b := range m.Multisig {
			l = len(b)
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Tx_CreateContractMsg) Size() (n int) {
	var l int
	_ = l
	if m.CreateContractMsg != nil {
		l = m.CreateContractMsg.Size()
		n += 2 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_UpdateContractMsg) Size() (n int) {
	var l int
	_ = l
	if m.UpdateContractMsg != nil {
		l = m.UpdateContractMsg.Size()
		n += 2 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_RecordVoteMsg) Size() (n int) {
	var l int
	_ = l
	if m.RecordVoteMsg != nil {
		l = m.RecordVoteMsg.Size()
		n += 2 + l + sovCodec(uint64(l))
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
func (m *Tx) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Tx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
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
			m.Signatures = append(m.Signatures, &sigs.StdSignature{})
			if err := m.Signatures[len(m.Signatures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multisig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Multisig = append(m.Multisig, make([]byte, postIndex-iNdEx))
			copy(m.Multisig[len(m.Multisig)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 56:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateContractMsg", wireType)
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
			v := &multisig.CreateContractMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_CreateContractMsg{v}
			iNdEx = postIndex
		case 57:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateContractMsg", wireType)
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
			v := &multisig.UpdateContractMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_UpdateContractMsg{v}
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordVoteMsg", wireType)
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
			v := &votes.VoteRecord{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_RecordVoteMsg{v}
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

func init() { proto.RegisterFile("app/codec.proto", fileDescriptor_codec_4d50eb422ace2421) }

var fileDescriptor_codec_4d50eb422ace2421 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0x87, 0x29, 0xa8, 0xd1, 0x45, 0x83, 0xd4, 0x0b, 0xa9, 0x49, 0x43, 0x3c, 0x91, 0x18, 0x76,
	0x13, 0xf4, 0xa0, 0xf1, 0x06, 0x17, 0x2e, 0x5c, 0x8a, 0x7a, 0x25, 0xed, 0xee, 0xb2, 0x6c, 0x42,
	0x3b, 0xcd, 0xfe, 0x41, 0x1e, 0xc3, 0xa7, 0xf1, 0x19, 0x3c, 0xfa, 0x08, 0x06, 0x5f, 0xc4, 0x74,
	0x11, 0x53, 0x6a, 0xe2, 0x6d, 0x67, 0xf6, 0x9b, 0x6f, 0x26, 0x3f, 0xd4, 0x8a, 0xf3, 0x9c, 0x50,
	0x60, 0x9c, 0xe2, 0x5c, 0x81, 0x01, 0xff, 0x84, 0x2a, 0xce, 0xe8, 0x22, 0x96, 0x59, 0xd0, 0x17,
	0xd2, 0x2c, 0x6c, 0x82, 0x29, 0xa4, 0x44, 0x80, 0x00, 0xe2, 0x88, 0xc4, 0xce, 0x5d, 0xe5, 0x0a,
	0xf7, 0xda, 0x4e, 0x06, 0xa4, 0x84, 0x4b, 0x58, 0xf5, 0x21, 0xe3, 0xe4, 0x85, 0xc7, 0x2b, 0x4e,
	0xd6, 0x24, 0xb5, 0x4b, 0x23, 0xb5, 0x14, 0xe5, 0x55, 0xc1, 0xf5, 0x3f, 0x03, 0x5a, 0x0a, 0xbd,
	0x07, 0xdf, 0x96, 0x60, 0x0a, 0xd9, 0x5c, 0x02, 0x29, 0x2e, 0x95, 0xc9, 0x92, 0xf7, 0xdd, 0xb9,
	0x64, 0x4d, 0x56, 0x60, 0xf8, 0xde, 0xd4, 0xd5, 0x5b, 0x1d, 0xd5, 0x1f, 0xd7, 0xfe, 0x00, 0x21,
	0x2d, 0x45, 0x16, 0x1b, 0xab, 0xb8, 0xee, 0xd4, 0xbb, 0x8d, 0x5e, 0x73, 0xe0, 0xe3, 0x62, 0x07,
	0x9e, 0x1a, 0x36, 0xdd, 0x7d, 0x45, 0x25, 0xca, 0x0f, 0xd0, 0xf1, 0xee, 0xea, 0xce, 0x41, 0xb7,
	0xd1, 0x3b, 0x8d, 0x7e, 0x6b, 0x7f, 0x82, 0x2e, 0xa8, 0xe2, 0xb1, 0xe1, 0x33, 0x0a, 0x99, 0x51,
	0x31, 0x35, 0xb3, 0x54, 0x8b, 0xce, 0x5d, 0xd7, 0xeb, 0x35, 0x07, 0x97, 0x78, 0xc7, 0xe1, 0x91,
	0x83, 0x46, 0x3f, 0xcc, 0x44, 0x8b, 0x71, 0x2d, 0x6a, 0xd3, 0x6a, 0xb3, 0xd0, 0xd9, 0x9c, 0xfd,
	0xd1, 0xdd, 0x57, 0x75, 0x4f, 0x0e, 0xaa, 0xe8, 0x6c, 0xb5, 0xe9, 0x3f, 0xa0, 0x96, 0xe2, 0x14,
	0x14, 0x9b, 0x15, 0x81, 0x38, 0x15, 0x73, 0xaa, 0x36, 0x76, 0x09, 0xe1, 0x67, 0x30, 0x3c, 0x72,
	0xc4, 0xb8, 0x16, 0x9d, 0x6d, 0xd9, 0xa2, 0x37, 0xd1, 0x62, 0x78, 0x88, 0x1a, 0xda, 0xa6, 0xc3,
	0xf3, 0xf7, 0x4d, 0xe8, 0x7d, 0x6c, 0x42, 0xef, 0x73, 0x13, 0x7a, 0xaf, 0x5f, 0x61, 0x2d, 0x39,
	0x72, 0x89, 0xde, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xfa, 0x31, 0x0e, 0x32, 0x02, 0x00,
	0x00,
}
