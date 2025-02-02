// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/v1beta1/osmosis.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgSwapExactAmountIn stores the tx Msg type to swap tokens in the trade ICA
type MsgSwapExactAmountIn struct {
	Sender            string                `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	Routes            []SwapAmountInRoute   `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes"`
	TokenIn           types.Coin            `protobuf:"bytes,3,opt,name=token_in,json=tokenIn,proto3" json:"token_in" yaml:"token_in"`
	TokenOutMinAmount cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=token_out_min_amount,json=tokenOutMinAmount,proto3,customtype=cosmossdk.io/math.Int" json:"token_out_min_amount" yaml:"token_out_min_amount"`
}

func (m *MsgSwapExactAmountIn) Reset()         { *m = MsgSwapExactAmountIn{} }
func (m *MsgSwapExactAmountIn) String() string { return proto.CompactTextString(m) }
func (*MsgSwapExactAmountIn) ProtoMessage()    {}
func (*MsgSwapExactAmountIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_39668b2e9488de8c, []int{0}
}
func (m *MsgSwapExactAmountIn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapExactAmountIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapExactAmountIn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapExactAmountIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapExactAmountIn.Merge(m, src)
}
func (m *MsgSwapExactAmountIn) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapExactAmountIn) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapExactAmountIn.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapExactAmountIn proto.InternalMessageInfo

func (m *MsgSwapExactAmountIn) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgSwapExactAmountIn) GetRoutes() []SwapAmountInRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *MsgSwapExactAmountIn) GetTokenIn() types.Coin {
	if m != nil {
		return m.TokenIn
	}
	return types.Coin{}
}

type SwapAmountInRoute struct {
	PoolId        uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty" yaml:"pool_id"`
	TokenOutDenom string `protobuf:"bytes,2,opt,name=token_out_denom,json=tokenOutDenom,proto3" json:"token_out_denom,omitempty" yaml:"token_out_denom"`
}

func (m *SwapAmountInRoute) Reset()         { *m = SwapAmountInRoute{} }
func (m *SwapAmountInRoute) String() string { return proto.CompactTextString(m) }
func (*SwapAmountInRoute) ProtoMessage()    {}
func (*SwapAmountInRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_39668b2e9488de8c, []int{1}
}
func (m *SwapAmountInRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SwapAmountInRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SwapAmountInRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SwapAmountInRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapAmountInRoute.Merge(m, src)
}
func (m *SwapAmountInRoute) XXX_Size() int {
	return m.Size()
}
func (m *SwapAmountInRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapAmountInRoute.DiscardUnknown(m)
}

var xxx_messageInfo_SwapAmountInRoute proto.InternalMessageInfo

func (m *SwapAmountInRoute) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *SwapAmountInRoute) GetTokenOutDenom() string {
	if m != nil {
		return m.TokenOutDenom
	}
	return ""
}

// A TwapRecord stores the most recent price of a pair of denom's
type OsmosisTwapRecord struct {
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// Lexicographically smaller denom of the pair
	Asset0Denom string `protobuf:"bytes,2,opt,name=asset0_denom,json=asset0Denom,proto3" json:"asset0_denom,omitempty"`
	// Lexicographically larger denom of the pair
	Asset1Denom string `protobuf:"bytes,3,opt,name=asset1_denom,json=asset1Denom,proto3" json:"asset1_denom,omitempty"`
	// height this record corresponds to, for debugging purposes
	Height int64 `protobuf:"varint,4,opt,name=height,proto3" json:"record_height" yaml:"record_height"`
	// This field should only exist until we have a global registry in the state
	// machine, mapping prior block heights within {TIME RANGE} to times.
	Time time.Time `protobuf:"bytes,5,opt,name=time,proto3,stdtime" json:"time" yaml:"record_time"`
	// We store the last spot prices in the struct, so that we can interpolate
	// accumulator values for times between when accumulator records are stored.
	P0LastSpotPrice             cosmossdk_io_math.LegacyDec `protobuf:"bytes,6,opt,name=p0_last_spot_price,json=p0LastSpotPrice,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"p0_last_spot_price"`
	P1LastSpotPrice             cosmossdk_io_math.LegacyDec `protobuf:"bytes,7,opt,name=p1_last_spot_price,json=p1LastSpotPrice,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"p1_last_spot_price"`
	P0ArithmeticTwapAccumulator cosmossdk_io_math.LegacyDec `protobuf:"bytes,8,opt,name=p0_arithmetic_twap_accumulator,json=p0ArithmeticTwapAccumulator,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"p0_arithmetic_twap_accumulator"`
	P1ArithmeticTwapAccumulator cosmossdk_io_math.LegacyDec `protobuf:"bytes,9,opt,name=p1_arithmetic_twap_accumulator,json=p1ArithmeticTwapAccumulator,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"p1_arithmetic_twap_accumulator"`
	GeometricTwapAccumulator    cosmossdk_io_math.LegacyDec `protobuf:"bytes,10,opt,name=geometric_twap_accumulator,json=geometricTwapAccumulator,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"geometric_twap_accumulator"`
	// This field contains the time in which the last spot price error occured.
	// It is used to alert the caller if they are getting a potentially erroneous
	// TWAP, due to an unforeseen underlying error.
	LastErrorTime time.Time `protobuf:"bytes,11,opt,name=last_error_time,json=lastErrorTime,proto3,stdtime" json:"last_error_time" yaml:"last_error_time"`
}

func (m *OsmosisTwapRecord) Reset()         { *m = OsmosisTwapRecord{} }
func (m *OsmosisTwapRecord) String() string { return proto.CompactTextString(m) }
func (*OsmosisTwapRecord) ProtoMessage()    {}
func (*OsmosisTwapRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_39668b2e9488de8c, []int{2}
}
func (m *OsmosisTwapRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OsmosisTwapRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OsmosisTwapRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OsmosisTwapRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OsmosisTwapRecord.Merge(m, src)
}
func (m *OsmosisTwapRecord) XXX_Size() int {
	return m.Size()
}
func (m *OsmosisTwapRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_OsmosisTwapRecord.DiscardUnknown(m)
}

var xxx_messageInfo_OsmosisTwapRecord proto.InternalMessageInfo

func (m *OsmosisTwapRecord) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *OsmosisTwapRecord) GetAsset0Denom() string {
	if m != nil {
		return m.Asset0Denom
	}
	return ""
}

func (m *OsmosisTwapRecord) GetAsset1Denom() string {
	if m != nil {
		return m.Asset1Denom
	}
	return ""
}

func (m *OsmosisTwapRecord) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *OsmosisTwapRecord) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *OsmosisTwapRecord) GetLastErrorTime() time.Time {
	if m != nil {
		return m.LastErrorTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*MsgSwapExactAmountIn)(nil), "osmosis.gamm.v1beta1.MsgSwapExactAmountIn")
	proto.RegisterType((*SwapAmountInRoute)(nil), "osmosis.gamm.v1beta1.SwapAmountInRoute")
	proto.RegisterType((*OsmosisTwapRecord)(nil), "osmosis.gamm.v1beta1.OsmosisTwapRecord")
}

func init() {
	proto.RegisterFile("osmosis/gamm/v1beta1/osmosis.proto", fileDescriptor_39668b2e9488de8c)
}

var fileDescriptor_39668b2e9488de8c = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x8f, 0x9b, 0x90, 0xee, 0x4e, 0x08, 0x51, 0xac, 0xc0, 0x9a, 0x54, 0xb2, 0xbb, 0x46, 0x82,
	0x2e, 0x28, 0x76, 0x5c, 0x24, 0x0e, 0x2b, 0x2e, 0x31, 0xdb, 0x43, 0xa4, 0x94, 0x5d, 0xb9, 0x7b,
	0xe2, 0x62, 0x4d, 0x9c, 0xa9, 0x33, 0x6a, 0xc6, 0x63, 0x79, 0xc6, 0x4d, 0x7b, 0xe7, 0xc4, 0xa9,
	0x1f, 0x85, 0x8f, 0xd1, 0x63, 0x8f, 0x88, 0x83, 0x41, 0xed, 0x01, 0x89, 0x03, 0x87, 0x7c, 0x02,
	0x34, 0x33, 0x4e, 0xda, 0xa4, 0x85, 0x6d, 0x2f, 0xd1, 0xf8, 0xcd, 0xef, 0xcf, 0x7b, 0xf3, 0x7e,
	0x01, 0x36, 0x65, 0x84, 0x32, 0xcc, 0xdc, 0x18, 0x12, 0xe2, 0x9e, 0x7a, 0x63, 0xc4, 0xa1, 0xe7,
	0x96, 0x45, 0x27, 0xcd, 0x28, 0xa7, 0x7a, 0x67, 0xf9, 0x29, 0x30, 0x4e, 0x89, 0xe9, 0x76, 0x62,
	0x1a, 0x53, 0x09, 0x70, 0xc5, 0x49, 0x61, 0xbb, 0x6d, 0x48, 0x70, 0x42, 0x5d, 0xf9, 0x5b, 0x96,
	0xcc, 0x48, 0xf2, 0xdd, 0x31, 0x64, 0x68, 0xe5, 0x10, 0x51, 0x9c, 0x94, 0xf7, 0x56, 0x4c, 0x69,
	0x3c, 0x43, 0xae, 0xfc, 0x1a, 0xe7, 0xc7, 0x2e, 0xc7, 0x04, 0x31, 0x0e, 0x49, 0xaa, 0x00, 0xf6,
	0x3f, 0x5b, 0xa0, 0x73, 0xc8, 0xe2, 0xa3, 0x39, 0x4c, 0x0f, 0xce, 0x60, 0xc4, 0x07, 0x84, 0xe6,
	0x09, 0x1f, 0x26, 0xfa, 0x2b, 0x50, 0x67, 0x28, 0x99, 0xa0, 0xcc, 0xd0, 0x76, 0xb5, 0xbd, 0xe7,
	0x7e, 0x7b, 0x51, 0x58, 0xcd, 0x73, 0x48, 0x66, 0xaf, 0x6d, 0x55, 0xb7, 0x83, 0x12, 0xa0, 0x1f,
	0x80, 0x7a, 0x46, 0x73, 0x8e, 0x98, 0xb1, 0xb5, 0x5b, 0xdd, 0x6b, 0xec, 0x7f, 0xe5, 0x3c, 0x34,
	0x94, 0x23, 0x3c, 0x96, 0xf2, 0x81, 0xc0, 0xfb, 0xb5, 0xcb, 0xc2, 0xaa, 0x04, 0x25, 0x59, 0x3f,
	0x04, 0xcf, 0x38, 0x3d, 0x41, 0x49, 0x88, 0x13, 0xa3, 0xba, 0xab, 0xed, 0x35, 0xf6, 0x3f, 0x77,
	0xd4, 0x78, 0x8e, 0x18, 0x6f, 0xa5, 0xf3, 0x03, 0xc5, 0x89, 0xff, 0x42, 0x50, 0x17, 0x85, 0xd5,
	0x52, 0x2d, 0x2d, 0x89, 0x76, 0xb0, 0x2d, 0x8f, 0xc3, 0x44, 0x27, 0xa0, 0xa3, 0xaa, 0x34, 0xe7,
	0x21, 0xc1, 0x49, 0x08, 0xa5, 0xb7, 0x51, 0x93, 0xe3, 0x7c, 0x2f, 0xf8, 0xbf, 0x17, 0xd6, 0xa7,
	0xca, 0x81, 0x4d, 0x4e, 0x1c, 0x4c, 0x5d, 0x02, 0xf9, 0xd4, 0x19, 0x26, 0x7c, 0x51, 0x58, 0x3b,
	0x77, 0x85, 0xd7, 0x25, 0xec, 0xa0, 0x2d, 0xcb, 0x6f, 0x73, 0x7e, 0x88, 0x13, 0x35, 0xd2, 0xeb,
	0x2f, 0x7f, 0xf9, 0xeb, 0xd7, 0xaf, 0x5f, 0xae, 0x6d, 0x9c, 0xcd, 0x61, 0xda, 0x43, 0xe2, 0x55,
	0x7b, 0x8a, 0xd8, 0xc3, 0x89, 0xfd, 0xb3, 0x06, 0xda, 0xf7, 0x5e, 0x42, 0xff, 0x06, 0x6c, 0xa7,
	0x94, 0xce, 0x42, 0x3c, 0x91, 0xcf, 0x5d, 0xf3, 0xf5, 0x45, 0x61, 0x7d, 0xa2, 0x5a, 0x28, 0x2f,
	0xec, 0xa0, 0x2e, 0x4e, 0xc3, 0x89, 0xee, 0x83, 0xd6, 0x6d, 0x5b, 0x13, 0x94, 0x50, 0x62, 0x6c,
	0xc9, 0xa1, 0xba, 0x8b, 0xc2, 0xfa, 0x6c, 0xb3, 0x6f, 0x09, 0xb0, 0x83, 0xe6, 0xb2, 0xe5, 0x37,
	0xf2, 0xfb, 0xb2, 0x0e, 0xda, 0x6f, 0x55, 0xb3, 0xef, 0xe7, 0x30, 0x0d, 0x50, 0x44, 0xb3, 0x89,
	0xfe, 0x62, 0xa3, 0x8d, 0x95, 0xe5, 0x4b, 0xf0, 0x31, 0x64, 0x0c, 0xf1, 0xfe, 0x5d, 0xbf, 0xa0,
	0xa1, 0x6a, 0x52, 0x71, 0x05, 0xf1, 0x4a, 0x48, 0xf5, 0x0e, 0xc4, 0x53, 0x90, 0x01, 0xa8, 0x4f,
	0x11, 0x8e, 0xa7, 0x6a, 0x09, 0x55, 0xff, 0xd5, 0xdf, 0x85, 0xd5, 0xcc, 0xa4, 0x75, 0xa8, 0x2e,
	0x16, 0x85, 0xd5, 0x51, 0x03, 0xac, 0x95, 0xed, 0xa0, 0x24, 0xea, 0x3f, 0x82, 0x9a, 0x88, 0xb0,
	0xf1, 0x91, 0x0c, 0x48, 0xd7, 0x51, 0xf9, 0x76, 0x96, 0xf9, 0x76, 0xde, 0x2f, 0xf3, 0xed, 0x9b,
	0x65, 0x42, 0xf4, 0x35, 0x3d, 0x41, 0xb6, 0x2f, 0xfe, 0xb0, 0xb4, 0x40, 0xea, 0xe8, 0xef, 0x80,
	0x9e, 0xf6, 0xc3, 0x19, 0x64, 0x3c, 0x64, 0x29, 0xe5, 0x61, 0x9a, 0xe1, 0x08, 0x19, 0x75, 0xf9,
	0x9c, 0x5f, 0x94, 0x19, 0xd9, 0xb9, 0x9f, 0x91, 0x11, 0x8a, 0x61, 0x74, 0xfe, 0x06, 0x45, 0x41,
	0x2b, 0xed, 0x8f, 0x20, 0xe3, 0x47, 0x29, 0xe5, 0xef, 0x04, 0x57, 0x2a, 0x7a, 0xf7, 0x14, 0xb7,
	0x9f, 0xa2, 0xe8, 0xad, 0x2b, 0x4e, 0x81, 0x99, 0xf6, 0x43, 0x98, 0x61, 0x3e, 0x25, 0x88, 0xe3,
	0x28, 0xe4, 0x73, 0x98, 0x86, 0x30, 0x8a, 0x72, 0x92, 0xcf, 0x20, 0xa7, 0x99, 0xf1, 0xec, 0xf1,
	0xea, 0x3b, 0x69, 0x7f, 0xb0, 0x52, 0x12, 0xab, 0x1f, 0xdc, 0xea, 0x48, 0x27, 0xef, 0x7f, 0x9d,
	0x9e, 0x3f, 0xc5, 0xc9, 0xfb, 0x6f, 0x27, 0x08, 0xba, 0x31, 0xa2, 0x04, 0xf1, 0xec, 0x21, 0x17,
	0xf0, 0x78, 0x17, 0x63, 0x25, 0xb3, 0x69, 0x71, 0x0c, 0x5a, 0x72, 0x0b, 0x28, 0xcb, 0x68, 0x26,
	0x17, 0x6f, 0x34, 0x3e, 0x98, 0x1a, 0xbb, 0x4c, 0x4d, 0xf9, 0x37, 0xda, 0x10, 0x50, 0xc9, 0x69,
	0x8a, 0xea, 0x81, 0x28, 0x0a, 0x9e, 0x3f, 0xba, 0xbc, 0x36, 0xb5, 0xab, 0x6b, 0x53, 0xfb, 0xf3,
	0xda, 0xd4, 0x2e, 0x6e, 0xcc, 0xca, 0xd5, 0x8d, 0x59, 0xf9, 0xed, 0xc6, 0xac, 0xfc, 0xb4, 0x1f,
	0x63, 0x3e, 0xcd, 0xc7, 0x4e, 0x44, 0x89, 0x7b, 0xc4, 0x33, 0x3c, 0x41, 0xbd, 0x11, 0x1c, 0x33,
	0x97, 0xc9, 0xb3, 0x7b, 0xea, 0x7d, 0xe7, 0x9e, 0xb9, 0x8c, 0xc3, 0x13, 0x84, 0xc7, 0x91, 0xcb,
	0xcf, 0x53, 0xc4, 0xc6, 0x75, 0xd9, 0xd4, 0xb7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xfc,
	0x9f, 0x22, 0x3d, 0x06, 0x00, 0x00,
}

func (m *MsgSwapExactAmountIn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapExactAmountIn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapExactAmountIn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TokenOutMinAmount.Size()
		i -= size
		if _, err := m.TokenOutMinAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOsmosis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.TokenIn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOsmosis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Routes) > 0 {
		for iNdEx := len(m.Routes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Routes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOsmosis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintOsmosis(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SwapAmountInRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SwapAmountInRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SwapAmountInRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenOutDenom) > 0 {
		i -= len(m.TokenOutDenom)
		copy(dAtA[i:], m.TokenOutDenom)
		i = encodeVarintOsmosis(dAtA, i, uint64(len(m.TokenOutDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintOsmosis(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OsmosisTwapRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OsmosisTwapRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OsmosisTwapRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.LastErrorTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastErrorTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintOsmosis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x5a
	{
		size := m.GeometricTwapAccumulator.Size()
		i -= size
		if _, err := m.GeometricTwapAccumulator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOsmosis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.P1ArithmeticTwapAccumulator.Size()
		i -= size
		if _, err := m.P1ArithmeticTwapAccumulator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOsmosis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.P0ArithmeticTwapAccumulator.Size()
		i -= size
		if _, err := m.P0ArithmeticTwapAccumulator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOsmosis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.P1LastSpotPrice.Size()
		i -= size
		if _, err := m.P1LastSpotPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOsmosis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.P0LastSpotPrice.Size()
		i -= size
		if _, err := m.P0LastSpotPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOsmosis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintOsmosis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x2a
	if m.Height != 0 {
		i = encodeVarintOsmosis(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Asset1Denom) > 0 {
		i -= len(m.Asset1Denom)
		copy(dAtA[i:], m.Asset1Denom)
		i = encodeVarintOsmosis(dAtA, i, uint64(len(m.Asset1Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Asset0Denom) > 0 {
		i -= len(m.Asset0Denom)
		copy(dAtA[i:], m.Asset0Denom)
		i = encodeVarintOsmosis(dAtA, i, uint64(len(m.Asset0Denom)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintOsmosis(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOsmosis(dAtA []byte, offset int, v uint64) int {
	offset -= sovOsmosis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSwapExactAmountIn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovOsmosis(uint64(l))
	}
	if len(m.Routes) > 0 {
		for _, e := range m.Routes {
			l = e.Size()
			n += 1 + l + sovOsmosis(uint64(l))
		}
	}
	l = m.TokenIn.Size()
	n += 1 + l + sovOsmosis(uint64(l))
	l = m.TokenOutMinAmount.Size()
	n += 1 + l + sovOsmosis(uint64(l))
	return n
}

func (m *SwapAmountInRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovOsmosis(uint64(m.PoolId))
	}
	l = len(m.TokenOutDenom)
	if l > 0 {
		n += 1 + l + sovOsmosis(uint64(l))
	}
	return n
}

func (m *OsmosisTwapRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovOsmosis(uint64(m.PoolId))
	}
	l = len(m.Asset0Denom)
	if l > 0 {
		n += 1 + l + sovOsmosis(uint64(l))
	}
	l = len(m.Asset1Denom)
	if l > 0 {
		n += 1 + l + sovOsmosis(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovOsmosis(uint64(m.Height))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovOsmosis(uint64(l))
	l = m.P0LastSpotPrice.Size()
	n += 1 + l + sovOsmosis(uint64(l))
	l = m.P1LastSpotPrice.Size()
	n += 1 + l + sovOsmosis(uint64(l))
	l = m.P0ArithmeticTwapAccumulator.Size()
	n += 1 + l + sovOsmosis(uint64(l))
	l = m.P1ArithmeticTwapAccumulator.Size()
	n += 1 + l + sovOsmosis(uint64(l))
	l = m.GeometricTwapAccumulator.Size()
	n += 1 + l + sovOsmosis(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastErrorTime)
	n += 1 + l + sovOsmosis(uint64(l))
	return n
}

func sovOsmosis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOsmosis(x uint64) (n int) {
	return sovOsmosis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSwapExactAmountIn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosis
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
			return fmt.Errorf("proto: MsgSwapExactAmountIn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapExactAmountIn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Routes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Routes = append(m.Routes, SwapAmountInRoute{})
			if err := m.Routes[len(m.Routes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutMinAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenOutMinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosis
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
func (m *SwapAmountInRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosis
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
			return fmt.Errorf("proto: SwapAmountInRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SwapAmountInRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOutDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosis
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
func (m *OsmosisTwapRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOsmosis
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
			return fmt.Errorf("proto: OsmosisTwapRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OsmosisTwapRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset0Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset0Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset1Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset1Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P0LastSpotPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P0LastSpotPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P1LastSpotPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P1LastSpotPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P0ArithmeticTwapAccumulator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P0ArithmeticTwapAccumulator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P1ArithmeticTwapAccumulator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P1ArithmeticTwapAccumulator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GeometricTwapAccumulator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GeometricTwapAccumulator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastErrorTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOsmosis
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
				return ErrInvalidLengthOsmosis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOsmosis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.LastErrorTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOsmosis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOsmosis
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
func skipOsmosis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOsmosis
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
					return 0, ErrIntOverflowOsmosis
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
					return 0, ErrIntOverflowOsmosis
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
				return 0, ErrInvalidLengthOsmosis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOsmosis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOsmosis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOsmosis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOsmosis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOsmosis = fmt.Errorf("proto: unexpected end of group")
)
