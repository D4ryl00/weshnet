// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errcode.proto

package errcode

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ErrCode int32

const (
	Undefined                                  ErrCode = 0
	TODO                                       ErrCode = 666
	ErrNotImplemented                          ErrCode = 777
	ErrInternal                                ErrCode = 888
	ErrInvalidInput                            ErrCode = 100
	ErrInvalidRange                            ErrCode = 101
	ErrMissingInput                            ErrCode = 102
	ErrSerialization                           ErrCode = 103
	ErrDeserialization                         ErrCode = 104
	ErrStreamRead                              ErrCode = 105
	ErrStreamWrite                             ErrCode = 106
	ErrMissingMapKey                           ErrCode = 107
	ErrDBWrite                                 ErrCode = 108
	ErrDBRead                                  ErrCode = 109
	ErrCryptoRandomGeneration                  ErrCode = 200
	ErrCryptoKeyGeneration                     ErrCode = 201
	ErrCryptoNonceGeneration                   ErrCode = 202
	ErrCryptoSignature                         ErrCode = 203
	ErrCryptoSignatureVerification             ErrCode = 204
	ErrCryptoDecrypt                           ErrCode = 205
	ErrCryptoDecryptPayload                    ErrCode = 206
	ErrCryptoEncrypt                           ErrCode = 207
	ErrCryptoKeyConversion                     ErrCode = 208
	ErrOrbitDBInit                             ErrCode = 1000
	ErrOrbitDBOpen                             ErrCode = 1001
	ErrOrbitDBAppend                           ErrCode = 1002
	ErrOrbitDBDeserialization                  ErrCode = 1003
	ErrOrbitDBStoreCast                        ErrCode = 1004
	ErrHandshakeOwnEphemeralKeyGenSend         ErrCode = 1100
	ErrHandshakePeerEphemeralKeyRecv           ErrCode = 1101
	ErrHandshakeRequesterAuthenticateBoxKeyGen ErrCode = 1102
	ErrHandshakeResponderAcceptBoxKeyGen       ErrCode = 1103
	ErrHandshakeRequesterHello                 ErrCode = 1104
	ErrHandshakeResponderHello                 ErrCode = 1105
	ErrHandshakeRequesterAuthenticate          ErrCode = 1106
	ErrHandshakeResponderAccept                ErrCode = 1107
	ErrHandshakeRequesterAcknowledge           ErrCode = 1108
	ErrContactRequestSameAccount               ErrCode = 1200
	ErrContactRequestContactAlreadyAdded       ErrCode = 1201
	ErrContactRequestContactBlocked            ErrCode = 1202
	ErrContactRequestContactUndefined          ErrCode = 1203
	ErrContactRequestIncomingAlreadyReceived   ErrCode = 1204
	ErrGroupMemberLogEventOpen                 ErrCode = 1300
	ErrGroupMemberLogEventSignature            ErrCode = 1301
	ErrGroupMemberUnknownGroupID               ErrCode = 1302
	ErrGroupSecretOtherDestMember              ErrCode = 1303
	ErrGroupSecretAlreadySentToMember          ErrCode = 1304
	ErrGroupInvalidType                        ErrCode = 1305
	ErrGroupMissing                            ErrCode = 1306
	ErrGroupActivate                           ErrCode = 1307
	ErrGroupDeactivate                         ErrCode = 1308
	ErrGroupInfo                               ErrCode = 1309
	// Event errors
	ErrEventListMetadata                   ErrCode = 1400
	ErrEventListMessage                    ErrCode = 1401
	ErrMessageKeyPersistencePut            ErrCode = 1500
	ErrMessageKeyPersistenceGet            ErrCode = 1501
	ErrBridgeInterrupted                   ErrCode = 1600
	ErrBridgeNotRunning                    ErrCode = 1601
	ErrMessengerInvalidDeepLink            ErrCode = 2000
	ErrDBEntryAlreadyExists                ErrCode = 2100
	ErrDBAddConversation                   ErrCode = 2101
	ErrDBAddContactRequestOutgoingSent     ErrCode = 2102
	ErrDBAddContactRequestOutgoingEnqueud  ErrCode = 2103
	ErrDBAddContactRequestIncomingReceived ErrCode = 2104
	ErrDBAddContactRequestIncomingAccepted ErrCode = 2105
	ErrDBAddGroupMemberDeviceAdded         ErrCode = 2106
	ErrDBMultipleRecords                   ErrCode = 2107
	ErrReplayProcessGroupMetadata          ErrCode = 2200
	ErrReplayProcessGroupMessage           ErrCode = 2201
	ErrCLINoTermcaps                       ErrCode = 3001
	ErrServicesAuth                        ErrCode = 4000
	ErrServicesAuthNotInitialized          ErrCode = 4001
	ErrServicesAuthWrongState              ErrCode = 4002
	ErrServicesAuthInvalidResponse         ErrCode = 4003
	ErrServicesAuthServer                  ErrCode = 4004
	ErrServicesAuthCodeChallenge           ErrCode = 4005
	ErrServicesAuthServiceInvalidToken     ErrCode = 4006
	ErrServicesAuthServiceNotSupported     ErrCode = 4007
	ErrServicesAuthUnknownToken            ErrCode = 4008
	ErrServicesAuthInvalidURL              ErrCode = 4009
	ErrServiceReplication                  ErrCode = 4100
	ErrServiceReplicationServer            ErrCode = 4101
	ErrServiceReplicationMissingEndpoint   ErrCode = 4102
	ErrBertyAccount                        ErrCode = 5000
	ErrBertyAccountNoIDSpecified           ErrCode = 5001
	ErrBertyAccountAlreadyOpened           ErrCode = 5002
	ErrBertyAccountInvalidIDFormat         ErrCode = 5003
	ErrBertyAccountLoggerDecorator         ErrCode = 5004
	ErrBertyAccountGRPCClient              ErrCode = 5005
	ErrBertyAccountOpenAccount             ErrCode = 5006
	ErrBertyAccountDataNotFound            ErrCode = 5007
	ErrBertyAccountMetadataUpdate          ErrCode = 5008
	ErrBertyAccountManagerOpen             ErrCode = 5009
	ErrBertyAccountManagerClose            ErrCode = 5010
	ErrBertyAccountInvalidCLIArgs          ErrCode = 5011
	ErrBertyAccountFSError                 ErrCode = 5012
	ErrBertyAccountAlreadyExists           ErrCode = 5013
	ErrBertyAccountNoBackupSpecified       ErrCode = 5014
	ErrBertyAccountIDGenFailed             ErrCode = 5015
	ErrBertyAccountCreationFailed          ErrCode = 5016
)

var ErrCode_name = map[int32]string{
	0:    "Undefined",
	666:  "TODO",
	777:  "ErrNotImplemented",
	888:  "ErrInternal",
	100:  "ErrInvalidInput",
	101:  "ErrInvalidRange",
	102:  "ErrMissingInput",
	103:  "ErrSerialization",
	104:  "ErrDeserialization",
	105:  "ErrStreamRead",
	106:  "ErrStreamWrite",
	107:  "ErrMissingMapKey",
	108:  "ErrDBWrite",
	109:  "ErrDBRead",
	200:  "ErrCryptoRandomGeneration",
	201:  "ErrCryptoKeyGeneration",
	202:  "ErrCryptoNonceGeneration",
	203:  "ErrCryptoSignature",
	204:  "ErrCryptoSignatureVerification",
	205:  "ErrCryptoDecrypt",
	206:  "ErrCryptoDecryptPayload",
	207:  "ErrCryptoEncrypt",
	208:  "ErrCryptoKeyConversion",
	1000: "ErrOrbitDBInit",
	1001: "ErrOrbitDBOpen",
	1002: "ErrOrbitDBAppend",
	1003: "ErrOrbitDBDeserialization",
	1004: "ErrOrbitDBStoreCast",
	1100: "ErrHandshakeOwnEphemeralKeyGenSend",
	1101: "ErrHandshakePeerEphemeralKeyRecv",
	1102: "ErrHandshakeRequesterAuthenticateBoxKeyGen",
	1103: "ErrHandshakeResponderAcceptBoxKeyGen",
	1104: "ErrHandshakeRequesterHello",
	1105: "ErrHandshakeResponderHello",
	1106: "ErrHandshakeRequesterAuthenticate",
	1107: "ErrHandshakeResponderAccept",
	1108: "ErrHandshakeRequesterAcknowledge",
	1200: "ErrContactRequestSameAccount",
	1201: "ErrContactRequestContactAlreadyAdded",
	1202: "ErrContactRequestContactBlocked",
	1203: "ErrContactRequestContactUndefined",
	1204: "ErrContactRequestIncomingAlreadyReceived",
	1300: "ErrGroupMemberLogEventOpen",
	1301: "ErrGroupMemberLogEventSignature",
	1302: "ErrGroupMemberUnknownGroupID",
	1303: "ErrGroupSecretOtherDestMember",
	1304: "ErrGroupSecretAlreadySentToMember",
	1305: "ErrGroupInvalidType",
	1306: "ErrGroupMissing",
	1307: "ErrGroupActivate",
	1308: "ErrGroupDeactivate",
	1309: "ErrGroupInfo",
	1400: "ErrEventListMetadata",
	1401: "ErrEventListMessage",
	1500: "ErrMessageKeyPersistencePut",
	1501: "ErrMessageKeyPersistenceGet",
	1600: "ErrBridgeInterrupted",
	1601: "ErrBridgeNotRunning",
	2000: "ErrMessengerInvalidDeepLink",
	2100: "ErrDBEntryAlreadyExists",
	2101: "ErrDBAddConversation",
	2102: "ErrDBAddContactRequestOutgoingSent",
	2103: "ErrDBAddContactRequestOutgoingEnqueud",
	2104: "ErrDBAddContactRequestIncomingReceived",
	2105: "ErrDBAddContactRequestIncomingAccepted",
	2106: "ErrDBAddGroupMemberDeviceAdded",
	2107: "ErrDBMultipleRecords",
	2200: "ErrReplayProcessGroupMetadata",
	2201: "ErrReplayProcessGroupMessage",
	3001: "ErrCLINoTermcaps",
	4000: "ErrServicesAuth",
	4001: "ErrServicesAuthNotInitialized",
	4002: "ErrServicesAuthWrongState",
	4003: "ErrServicesAuthInvalidResponse",
	4004: "ErrServicesAuthServer",
	4005: "ErrServicesAuthCodeChallenge",
	4006: "ErrServicesAuthServiceInvalidToken",
	4007: "ErrServicesAuthServiceNotSupported",
	4008: "ErrServicesAuthUnknownToken",
	4009: "ErrServicesAuthInvalidURL",
	4100: "ErrServiceReplication",
	4101: "ErrServiceReplicationServer",
	4102: "ErrServiceReplicationMissingEndpoint",
	5000: "ErrBertyAccount",
	5001: "ErrBertyAccountNoIDSpecified",
	5002: "ErrBertyAccountAlreadyOpened",
	5003: "ErrBertyAccountInvalidIDFormat",
	5004: "ErrBertyAccountLoggerDecorator",
	5005: "ErrBertyAccountGRPCClient",
	5006: "ErrBertyAccountOpenAccount",
	5007: "ErrBertyAccountDataNotFound",
	5008: "ErrBertyAccountMetadataUpdate",
	5009: "ErrBertyAccountManagerOpen",
	5010: "ErrBertyAccountManagerClose",
	5011: "ErrBertyAccountInvalidCLIArgs",
	5012: "ErrBertyAccountFSError",
	5013: "ErrBertyAccountAlreadyExists",
	5014: "ErrBertyAccountNoBackupSpecified",
	5015: "ErrBertyAccountIDGenFailed",
	5016: "ErrBertyAccountCreationFailed",
}

var ErrCode_value = map[string]int32{
	"Undefined":                                  0,
	"TODO":                                       666,
	"ErrNotImplemented":                          777,
	"ErrInternal":                                888,
	"ErrInvalidInput":                            100,
	"ErrInvalidRange":                            101,
	"ErrMissingInput":                            102,
	"ErrSerialization":                           103,
	"ErrDeserialization":                         104,
	"ErrStreamRead":                              105,
	"ErrStreamWrite":                             106,
	"ErrMissingMapKey":                           107,
	"ErrDBWrite":                                 108,
	"ErrDBRead":                                  109,
	"ErrCryptoRandomGeneration":                  200,
	"ErrCryptoKeyGeneration":                     201,
	"ErrCryptoNonceGeneration":                   202,
	"ErrCryptoSignature":                         203,
	"ErrCryptoSignatureVerification":             204,
	"ErrCryptoDecrypt":                           205,
	"ErrCryptoDecryptPayload":                    206,
	"ErrCryptoEncrypt":                           207,
	"ErrCryptoKeyConversion":                     208,
	"ErrOrbitDBInit":                             1000,
	"ErrOrbitDBOpen":                             1001,
	"ErrOrbitDBAppend":                           1002,
	"ErrOrbitDBDeserialization":                  1003,
	"ErrOrbitDBStoreCast":                        1004,
	"ErrHandshakeOwnEphemeralKeyGenSend":         1100,
	"ErrHandshakePeerEphemeralKeyRecv":           1101,
	"ErrHandshakeRequesterAuthenticateBoxKeyGen": 1102,
	"ErrHandshakeResponderAcceptBoxKeyGen":       1103,
	"ErrHandshakeRequesterHello":                 1104,
	"ErrHandshakeResponderHello":                 1105,
	"ErrHandshakeRequesterAuthenticate":          1106,
	"ErrHandshakeResponderAccept":                1107,
	"ErrHandshakeRequesterAcknowledge":           1108,
	"ErrContactRequestSameAccount":               1200,
	"ErrContactRequestContactAlreadyAdded":       1201,
	"ErrContactRequestContactBlocked":            1202,
	"ErrContactRequestContactUndefined":          1203,
	"ErrContactRequestIncomingAlreadyReceived":   1204,
	"ErrGroupMemberLogEventOpen":                 1300,
	"ErrGroupMemberLogEventSignature":            1301,
	"ErrGroupMemberUnknownGroupID":               1302,
	"ErrGroupSecretOtherDestMember":              1303,
	"ErrGroupSecretAlreadySentToMember":          1304,
	"ErrGroupInvalidType":                        1305,
	"ErrGroupMissing":                            1306,
	"ErrGroupActivate":                           1307,
	"ErrGroupDeactivate":                         1308,
	"ErrGroupInfo":                               1309,
	"ErrEventListMetadata":                       1400,
	"ErrEventListMessage":                        1401,
	"ErrMessageKeyPersistencePut":                1500,
	"ErrMessageKeyPersistenceGet":                1501,
	"ErrBridgeInterrupted":                       1600,
	"ErrBridgeNotRunning":                        1601,
	"ErrMessengerInvalidDeepLink":                2000,
	"ErrDBEntryAlreadyExists":                    2100,
	"ErrDBAddConversation":                       2101,
	"ErrDBAddContactRequestOutgoingSent":         2102,
	"ErrDBAddContactRequestOutgoingEnqueud":      2103,
	"ErrDBAddContactRequestIncomingReceived":     2104,
	"ErrDBAddContactRequestIncomingAccepted":     2105,
	"ErrDBAddGroupMemberDeviceAdded":             2106,
	"ErrDBMultipleRecords":                       2107,
	"ErrReplayProcessGroupMetadata":              2200,
	"ErrReplayProcessGroupMessage":               2201,
	"ErrCLINoTermcaps":                           3001,
	"ErrServicesAuth":                            4000,
	"ErrServicesAuthNotInitialized":              4001,
	"ErrServicesAuthWrongState":                  4002,
	"ErrServicesAuthInvalidResponse":             4003,
	"ErrServicesAuthServer":                      4004,
	"ErrServicesAuthCodeChallenge":               4005,
	"ErrServicesAuthServiceInvalidToken":         4006,
	"ErrServicesAuthServiceNotSupported":         4007,
	"ErrServicesAuthUnknownToken":                4008,
	"ErrServicesAuthInvalidURL":                  4009,
	"ErrServiceReplication":                      4100,
	"ErrServiceReplicationServer":                4101,
	"ErrServiceReplicationMissingEndpoint":       4102,
	"ErrBertyAccount":                            5000,
	"ErrBertyAccountNoIDSpecified":               5001,
	"ErrBertyAccountAlreadyOpened":               5002,
	"ErrBertyAccountInvalidIDFormat":             5003,
	"ErrBertyAccountLoggerDecorator":             5004,
	"ErrBertyAccountGRPCClient":                  5005,
	"ErrBertyAccountOpenAccount":                 5006,
	"ErrBertyAccountDataNotFound":                5007,
	"ErrBertyAccountMetadataUpdate":              5008,
	"ErrBertyAccountManagerOpen":                 5009,
	"ErrBertyAccountManagerClose":                5010,
	"ErrBertyAccountInvalidCLIArgs":              5011,
	"ErrBertyAccountFSError":                     5012,
	"ErrBertyAccountAlreadyExists":               5013,
	"ErrBertyAccountNoBackupSpecified":           5014,
	"ErrBertyAccountIDGenFailed":                 5015,
	"ErrBertyAccountCreationFailed":              5016,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4240057316120df7, []int{0}
}

type ErrDetails struct {
	Codes                []ErrCode `protobuf:"varint,1,rep,packed,name=codes,proto3,enum=berty.errcode.ErrCode" json:"codes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ErrDetails) Reset()         { *m = ErrDetails{} }
func (m *ErrDetails) String() string { return proto.CompactTextString(m) }
func (*ErrDetails) ProtoMessage()    {}
func (*ErrDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_4240057316120df7, []int{0}
}

func (m *ErrDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrDetails.Unmarshal(m, b)
}

func (m *ErrDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrDetails.Marshal(b, m, deterministic)
}

func (m *ErrDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrDetails.Merge(m, src)
}

func (m *ErrDetails) XXX_Size() int {
	return xxx_messageInfo_ErrDetails.Size(m)
}

func (m *ErrDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ErrDetails proto.InternalMessageInfo

func (m *ErrDetails) GetCodes() []ErrCode {
	if m != nil {
		return m.Codes
	}
	return nil
}

func init() {
	proto.RegisterEnum("berty.errcode.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterType((*ErrDetails)(nil), "berty.errcode.ErrDetails")
}

func init() { proto.RegisterFile("errcode.proto", fileDescriptor_4240057316120df7) }

var fileDescriptor_4240057316120df7 = []byte{
	// 1487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x97, 0x49, 0x73, 0x1c, 0x49,
	0x15, 0x80, 0xa7, 0x03, 0xc6, 0x92, 0xd2, 0x96, 0xe7, 0x29, 0xad, 0x91, 0x35, 0x9e, 0x19, 0xb7,
	0xc6, 0x8c, 0xa7, 0x07, 0x03, 0x52, 0x04, 0xdc, 0xb8, 0xf5, 0x26, 0x4d, 0x87, 0xb5, 0x45, 0xb7,
	0xc4, 0x44, 0x70, 0x4b, 0x55, 0x3e, 0x95, 0x92, 0xae, 0xce, 0xac, 0xc9, 0xca, 0xd2, 0x4c, 0x73,
	0x06, 0x82, 0x7d, 0xf5, 0x0e, 0x44, 0xb0, 0x2f, 0x37, 0x16, 0x03, 0x36, 0x5c, 0xe0, 0xc6, 0xe2,
	0x45, 0x2c, 0x47, 0x38, 0x70, 0x63, 0xfb, 0x01, 0xe6, 0x46, 0x64, 0xd5, 0x6b, 0x75, 0xb5, 0xd4,
	0x08, 0x4e, 0x6a, 0xbd, 0xf7, 0xe5, 0xdb, 0xf2, 0xe5, 0xcb, 0x2c, 0x36, 0x8d, 0xd6, 0x06, 0x46,
	0xe2, 0x62, 0x6c, 0x8d, 0x33, 0x7c, 0x7a, 0x07, 0xad, 0xeb, 0x2f, 0x92, 0xf0, 0xc2, 0x6c, 0x68,
	0x42, 0x93, 0x69, 0x96, 0xfc, 0xaf, 0x1c, 0xba, 0xf4, 0x7e, 0xc6, 0x9a, 0xd6, 0x36, 0xd0, 0x09,
	0x15, 0x25, 0xfc, 0xdd, 0xec, 0x69, 0xcf, 0x26, 0xf3, 0xa5, 0x85, 0xb7, 0xbd, 0x7a, 0xf6, 0xbd,
	0x73, 0x8b, 0x23, 0x26, 0x16, 0x9b, 0xd6, 0xd6, 0x8d, 0xc4, 0x76, 0x0e, 0x5d, 0xb9, 0x37, 0xcf,
	0x26, 0x48, 0xc4, 0xa7, 0xd9, 0xd4, 0xb6, 0x96, 0xb8, 0xab, 0x34, 0x4a, 0x78, 0x8a, 0x4f, 0xb1,
	0xb7, 0x6f, 0x6d, 0x34, 0x36, 0xe0, 0xce, 0xd3, 0x7c, 0x8e, 0xcd, 0x34, 0xad, 0x5d, 0x37, 0xae,
	0xd5, 0x8b, 0x23, 0xec, 0xa1, 0x76, 0x28, 0xe1, 0x13, 0xa7, 0x38, 0xb0, 0xd3, 0x4d, 0x6b, 0x5b,
	0xda, 0xa1, 0xd5, 0x22, 0x82, 0x27, 0xa7, 0xf8, 0x39, 0xf6, 0x4c, 0x26, 0xd9, 0x17, 0x91, 0x92,
	0x2d, 0x1d, 0xa7, 0x0e, 0xe4, 0xa8, 0xb0, 0x2d, 0x74, 0x88, 0x80, 0x24, 0x5c, 0x53, 0x49, 0xa2,
	0x74, 0x98, 0x93, 0xbb, 0x7c, 0x96, 0x41, 0xd3, 0xda, 0x0e, 0x5a, 0x25, 0x22, 0xf5, 0x61, 0xe1,
	0x94, 0xd1, 0x10, 0xf2, 0x39, 0xc6, 0xb3, 0x04, 0x93, 0x11, 0xf9, 0x1e, 0x9f, 0x61, 0xd3, 0x9e,
	0x76, 0x16, 0x45, 0xaf, 0x8d, 0x42, 0x82, 0xe2, 0x9c, 0x9d, 0x3d, 0x14, 0xbd, 0x6e, 0x95, 0x43,
	0xf8, 0x10, 0x19, 0x25, 0x4f, 0x6b, 0x22, 0xbe, 0x8a, 0x7d, 0xe8, 0xf2, 0xb3, 0x79, 0xd5, 0x6a,
	0x39, 0x15, 0xf9, 0xec, 0xb3, 0xff, 0x33, 0x43, 0x3d, 0x7e, 0x91, 0x3d, 0xe7, 0xeb, 0x62, 0xfb,
	0xb1, 0x33, 0x6d, 0xa1, 0xa5, 0xe9, 0xad, 0xa0, 0x46, 0x9b, 0xbb, 0xfe, 0x75, 0x89, 0x3f, 0xcf,
	0xe6, 0x0e, 0xf5, 0x57, 0xb1, 0x5f, 0x50, 0xfe, 0xa6, 0xc4, 0x5f, 0x64, 0xf3, 0x87, 0xca, 0x75,
	0xa3, 0x03, 0x2c, 0xa8, 0x7f, 0x5b, 0xe2, 0xe7, 0xb3, 0x7c, 0x72, 0x75, 0x47, 0x85, 0x5a, 0xb8,
	0xd4, 0x22, 0xfc, 0xae, 0xc4, 0xdf, 0xc1, 0x2e, 0x1e, 0x57, 0x7c, 0x00, 0xad, 0xda, 0x55, 0x41,
	0xbe, 0xfa, 0x41, 0x89, 0x3f, 0x9b, 0xa5, 0x93, 0x43, 0x0d, 0x0c, 0xfc, 0x5f, 0x78, 0x58, 0xe2,
	0x2f, 0xb0, 0xf3, 0x47, 0xc5, 0x9b, 0xa2, 0x1f, 0x19, 0x21, 0xe1, 0xd1, 0xe8, 0xa2, 0xa6, 0xce,
	0x17, 0x3d, 0x3e, 0x96, 0x45, 0xdd, 0xe8, 0x7d, 0xb4, 0x89, 0x77, 0x74, 0x50, 0xe2, 0xe7, 0xb2,
	0x5a, 0x6e, 0xd8, 0x1d, 0xe5, 0x1a, 0xb5, 0x96, 0x56, 0x0e, 0xfe, 0x36, 0x31, 0x2a, 0xdc, 0x88,
	0x51, 0xc3, 0xdf, 0x27, 0xc8, 0x3a, 0x09, 0xab, 0x71, 0x8c, 0x5a, 0xc2, 0x3f, 0x26, 0xa8, 0x86,
	0x24, 0x3e, 0xba, 0x7d, 0xff, 0x9c, 0xe0, 0xf3, 0xec, 0xdc, 0x50, 0xdf, 0x71, 0xc6, 0x62, 0x5d,
	0x24, 0x0e, 0xfe, 0x35, 0xc1, 0x2b, 0xec, 0x52, 0xd3, 0xda, 0xd7, 0x84, 0x96, 0xc9, 0x9e, 0xe8,
	0xe2, 0xc6, 0x9b, 0xba, 0x19, 0xef, 0x61, 0x0f, 0xad, 0x88, 0xf2, 0x62, 0x77, 0xbc, 0x8b, 0x07,
	0x93, 0xfc, 0x32, 0x5b, 0x28, 0x82, 0x9b, 0x88, 0xb6, 0x48, 0xb6, 0x31, 0xd8, 0x87, 0x87, 0x93,
	0x7c, 0x89, 0x5d, 0x29, 0x62, 0x6d, 0x7c, 0x23, 0xc5, 0xc4, 0xa1, 0xad, 0xa6, 0x6e, 0x0f, 0xb5,
	0xf3, 0xd5, 0xc5, 0x9a, 0x79, 0x2b, 0xb7, 0x0d, 0x8f, 0x26, 0xf9, 0x3b, 0xd9, 0xcb, 0xa3, 0x0b,
	0x92, 0xd8, 0x68, 0x89, 0xb6, 0x1a, 0x04, 0x18, 0xbb, 0x21, 0xfa, 0x78, 0x92, 0x97, 0xd9, 0x85,
	0xb1, 0xb6, 0x5f, 0xc3, 0x28, 0x32, 0x70, 0x30, 0x06, 0x20, 0x5b, 0x39, 0xf0, 0xfb, 0x49, 0xfe,
	0x0a, 0x7b, 0xe9, 0x7f, 0x46, 0x07, 0x7f, 0x98, 0xe4, 0x0b, 0xec, 0xf9, 0x13, 0x82, 0x82, 0x3f,
	0x1e, 0x2b, 0xc7, 0xd0, 0x52, 0xd0, 0xd5, 0xe6, 0xcd, 0x08, 0x65, 0x88, 0xf0, 0xa7, 0x49, 0xfe,
	0x12, 0x7b, 0x21, 0x3b, 0xf4, 0xda, 0x89, 0xc0, 0x11, 0xd4, 0x11, 0x3d, 0xac, 0x06, 0x81, 0x49,
	0xb5, 0x83, 0x1f, 0x4c, 0x51, 0x01, 0x46, 0x11, 0xfa, 0xaf, 0x1a, 0x59, 0x14, 0xb2, 0x5f, 0x95,
	0x12, 0x25, 0xfc, 0x70, 0x8a, 0xbf, 0xcc, 0xca, 0xff, 0x0d, 0xad, 0x45, 0x26, 0xe8, 0xa2, 0x84,
	0x1f, 0x4d, 0x51, 0x92, 0x63, 0xa9, 0xe1, 0xd4, 0xf9, 0xf1, 0x14, 0x7f, 0x0f, 0x7b, 0xf5, 0x18,
	0xd7, 0xd2, 0x81, 0xe9, 0x29, 0x1d, 0x92, 0xe7, 0x36, 0x06, 0xa8, 0xf6, 0x51, 0xc2, 0xdd, 0x29,
	0x2a, 0xee, 0x8a, 0x35, 0x69, 0xbc, 0x86, 0xbd, 0x1d, 0xb4, 0xab, 0x26, 0x6c, 0xee, 0xa3, 0x76,
	0x59, 0x6f, 0x5e, 0x63, 0x14, 0xdd, 0x18, 0x60, 0x78, 0xf2, 0xae, 0x33, 0xaa, 0x48, 0x81, 0xda,
	0xd6, 0xbe, 0x62, 0x3a, 0x93, 0xb4, 0x1a, 0x70, 0x83, 0xf1, 0x4b, 0xec, 0xc5, 0x01, 0xd2, 0xc1,
	0xc0, 0xa2, 0xdb, 0x70, 0x7b, 0xe8, 0xa7, 0x92, 0xcb, 0x57, 0xc0, 0x4d, 0x46, 0x49, 0x16, 0x18,
	0x8a, 0xb8, 0x83, 0xda, 0x6d, 0x19, 0xe2, 0x6e, 0x31, 0xea, 0xfc, 0xdc, 0x78, 0x3e, 0x16, 0xb7,
	0xfa, 0x31, 0xc2, 0x6d, 0xc6, 0x67, 0xb3, 0xb1, 0x98, 0x07, 0x92, 0x4f, 0x2c, 0xb8, 0xc3, 0xe8,
	0x80, 0x65, 0xd2, 0x6a, 0xe0, 0xd4, 0xbe, 0x6f, 0x88, 0x2f, 0x33, 0x1a, 0x24, 0x99, 0xb8, 0x81,
	0x62, 0xa0, 0xf8, 0x0a, 0xe3, 0x33, 0xec, 0xcc, 0xd0, 0xfe, 0xae, 0x81, 0xaf, 0x32, 0xfe, 0x1c,
	0x9b, 0x6d, 0x5a, 0x9b, 0x65, 0xbe, 0xaa, 0x7c, 0xcc, 0x4e, 0x48, 0xe1, 0x04, 0x3c, 0x19, 0x44,
	0x53, 0x50, 0x25, 0x89, 0x08, 0x11, 0xfe, 0xcd, 0xa8, 0xe3, 0x48, 0x70, 0x15, 0xfb, 0x9b, 0x7e,
	0x3a, 0x24, 0x0e, 0x75, 0x80, 0x9b, 0xa9, 0x83, 0x3f, 0x9f, 0x3e, 0x89, 0x58, 0x41, 0x07, 0x7f,
	0x39, 0x4d, 0x8e, 0x6b, 0x56, 0xc9, 0x10, 0xb3, 0xab, 0xc2, 0xa6, 0xb1, 0xbf, 0x3f, 0x7e, 0x79,
	0x86, 0x1c, 0xe7, 0xaa, 0x75, 0xe3, 0xda, 0xa9, 0xd6, 0x3e, 0xe1, 0x5f, 0x9d, 0x29, 0x98, 0x45,
	0x1d, 0xe2, 0xe0, 0xee, 0x68, 0x20, 0xc6, 0xab, 0x4a, 0x77, 0xe1, 0xe0, 0x19, 0x9a, 0x77, 0x8d,
	0x5a, 0x53, 0x3b, 0xdb, 0xa7, 0x32, 0x37, 0xdf, 0x52, 0x89, 0x4b, 0xe0, 0x2e, 0x90, 0xd3, 0x46,
	0xad, 0x2a, 0x25, 0x0d, 0xb5, 0x7c, 0xea, 0xfc, 0x04, 0x68, 0xb6, 0x0c, 0x54, 0x85, 0x2e, 0xdb,
	0x48, 0x5d, 0x68, 0x94, 0x0e, 0xfd, 0x66, 0xc1, 0x4f, 0x81, 0x5f, 0x61, 0x97, 0x4f, 0x06, 0x9b,
	0xfa, 0x8d, 0x14, 0x53, 0x09, 0x3f, 0x03, 0xfe, 0x2e, 0xf6, 0xca, 0x78, 0x76, 0xd0, 0xba, 0x87,
	0x3d, 0x7b, 0xef, 0xff, 0x80, 0xf3, 0x03, 0x8d, 0x12, 0xee, 0x03, 0xdd, 0x09, 0x19, 0x5c, 0x68,
	0xcf, 0x06, 0xee, 0xab, 0x00, 0xf3, 0x23, 0xf8, 0xf3, 0x61, 0xba, 0x6b, 0x69, 0xe4, 0x54, 0x1c,
	0x61, 0x1b, 0x03, 0x63, 0x65, 0x02, 0xbf, 0x00, 0x6a, 0xdb, 0x36, 0xc6, 0x91, 0xe8, 0x6f, 0x5a,
	0x13, 0x60, 0x92, 0x90, 0x1d, 0x6a, 0x80, 0x5b, 0x33, 0xd4, 0xfd, 0xe3, 0x98, 0xbc, 0x13, 0x6e,
	0xcf, 0x0c, 0x2e, 0x90, 0xd5, 0xd6, 0xba, 0xd9, 0x42, 0xdb, 0x0b, 0x44, 0x9c, 0xc0, 0xfd, 0xf3,
	0xd4, 0xae, 0x1d, 0xb4, 0x3e, 0x9c, 0xc4, 0x0f, 0x2c, 0xf8, 0x5a, 0x99, 0x7c, 0x16, 0xa5, 0xfe,
	0xed, 0xa0, 0x95, 0xcb, 0xe6, 0x3f, 0x4a, 0xf8, 0x7a, 0x99, 0x2e, 0x87, 0x22, 0xf3, 0xba, 0x35,
	0x3a, 0xec, 0x38, 0xdf, 0xc2, 0xdf, 0x28, 0x53, 0xde, 0x45, 0xfd, 0xe0, 0x01, 0x91, 0x8d, 0xbd,
	0x04, 0xe1, 0x9b, 0x65, 0x7e, 0x81, 0x3d, 0x7b, 0x04, 0xf2, 0xbf, 0xd1, 0xc2, 0xb7, 0xca, 0x94,
	0x54, 0x51, 0xe7, 0x5f, 0x39, 0xf5, 0x3d, 0x11, 0x45, 0xbe, 0xa7, 0xe0, 0xdb, 0x65, 0x6a, 0x85,
	0xa3, 0xcb, 0x55, 0x80, 0x83, 0x43, 0x69, 0xba, 0xa8, 0xe1, 0x3b, 0x27, 0x80, 0xeb, 0xc6, 0x75,
	0xd2, 0x38, 0x36, 0xd6, 0xef, 0xd6, 0x77, 0xcb, 0xd4, 0xb7, 0x45, 0x90, 0x06, 0x49, 0x6e, 0xea,
	0x7b, 0xe3, 0xf2, 0x26, 0x67, 0xdb, 0xed, 0x55, 0xf8, 0xfe, 0x91, 0x94, 0xfc, 0x96, 0x0c, 0xae,
	0xfe, 0x8f, 0x2c, 0x8c, 0x5a, 0x2f, 0xe8, 0x28, 0xe9, 0x8f, 0x2e, 0xd0, 0xd8, 0x3e, 0x4e, 0xd0,
	0x2c, 0x69, 0x6a, 0x19, 0x1b, 0xa5, 0x1d, 0x7c, 0x6c, 0x81, 0xb6, 0xae, 0xe6, 0x5f, 0x87, 0x83,
	0xb9, 0xff, 0xf1, 0x0a, 0x55, 0xad, 0x28, 0x5d, 0x37, 0xad, 0x46, 0x27, 0xc6, 0x40, 0xed, 0x2a,
	0xff, 0xea, 0x1b, 0x87, 0xd0, 0xf1, 0xf3, 0x23, 0x17, 0x25, 0x7c, 0xb2, 0x42, 0x9b, 0x57, 0x44,
	0x06, 0x4f, 0xc2, 0xc6, 0xb2, 0xb1, 0x3d, 0xe1, 0xe0, 0x53, 0xe3, 0xa0, 0x55, 0x13, 0x86, 0xbe,
	0xb7, 0x03, 0x63, 0x85, 0x33, 0x16, 0x3e, 0x5d, 0xa1, 0x72, 0x15, 0xa1, 0x95, 0xf6, 0x66, 0xbd,
	0x1e, 0x29, 0x7f, 0x48, 0x3f, 0x53, 0xa1, 0xf9, 0x5f, 0xd4, 0xfb, 0x28, 0x06, 0x09, 0x7d, 0xb6,
	0x42, 0x35, 0x2b, 0x02, 0x0d, 0xe1, 0xc4, 0xba, 0x71, 0xcb, 0x26, 0xd5, 0x12, 0x3e, 0x57, 0xa1,
	0x6e, 0x2d, 0x12, 0x83, 0xb3, 0xb1, 0x1d, 0x4b, 0xdf, 0x8d, 0x9f, 0x1f, 0xe7, 0x66, 0x4d, 0x68,
	0x11, 0xa2, 0xcd, 0xae, 0x99, 0x2f, 0x8c, 0x73, 0x43, 0x40, 0x3d, 0x32, 0x09, 0xc2, 0x17, 0xc7,
	0xb9, 0xa1, 0x9a, 0xd4, 0x57, 0x5b, 0x55, 0x1b, 0x26, 0xf0, 0xa5, 0x0a, 0xbd, 0xc7, 0x8a, 0xcc,
	0x72, 0xa7, 0x69, 0xad, 0xb1, 0x70, 0xed, 0x84, 0xba, 0xd3, 0xd8, 0xbb, 0x5e, 0xa1, 0xfb, 0x7f,
	0x74, 0xf7, 0x6a, 0x22, 0xe8, 0xa6, 0xf1, 0x70, 0x07, 0x6f, 0x8c, 0xcb, 0xa6, 0xd5, 0x58, 0x41,
	0xbd, 0x2c, 0x54, 0x84, 0x12, 0x6e, 0x8e, 0x8b, 0xb5, 0x6e, 0x31, 0x6b, 0x24, 0x62, 0x6e, 0x55,
	0x6a, 0x97, 0x0f, 0xfe, 0x7a, 0xf1, 0xa9, 0x0f, 0x96, 0xf3, 0xcf, 0x0b, 0x87, 0xc1, 0xde, 0x52,
	0xf6, 0x73, 0x29, 0x34, 0x4b, 0x71, 0x37, 0x5c, 0xa2, 0x0f, 0x8e, 0x9d, 0x53, 0xd9, 0x47, 0xca,
	0xfb, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xcb, 0xd5, 0xef, 0xda, 0x0c, 0x00, 0x00,
}
