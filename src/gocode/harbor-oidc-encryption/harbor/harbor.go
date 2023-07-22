package harbor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"io"
	"reflect"
	"strings"
	"sync"

	// "github.com/golang/protobuf/proto"		// 这个包已被弃用，使用下面的包
	"golang.org/x/crypto/pbkdf2"
	"google.golang.org/protobuf/proto"

	// "google.golang.org/protobuf/proto"	// 这个包被声明但是没有被使用
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// Marshal converts a protobuf message to a URL legal string.
func Marshal(message proto.Message) (string, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(data), nil
}

// Unmarshal decodes a protobuf message.
func Unmarshal(s string, message proto.Message) error {
	data, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	return proto.Unmarshal(data, message)
}

// func main() {
// 	fmt.Println(Marshal(&IDTokenSubject{ConnId: "cas", UserId: "17339872165"}))
// 	//data, err := base64.StdEncoding.DecodeString("CiQwOGE4Njg0Yi1kYjg4LTRiNzMtOTBhOS0zY2QxNjYxZjU0NjYSBWxvY2Fs")
// 	//fmt.Println(string(data), err)
// 	//fmt.Println(ReversibleDecrypt("<enc-v1>YVzL0B165Xg+iP2JbtRxe01P9IbvieI00nZu3A8sdKF9FlpcTbjoNQ==", "Cckbd4bz2882hRJy"))
// 	//fmt.Println(ReversibleDecrypt("<enc-v1>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", "Cckbd4bz2882hRJy"))
// }

// IDTokenSubject represents both the userID and connID which is returned
// as the "sub" claim in the ID Token.
type IDTokenSubject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ConnId string `protobuf:"bytes,2,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
}

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RefreshToken is a message that holds refresh token data used by dex.
type RefreshToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshId string `protobuf:"bytes,1,opt,name=refresh_id,json=refreshId,proto3" json:"refresh_id,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RefreshToken) Reset() {
	*x = RefreshToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_internal_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshToken) ProtoMessage() {}

func (x *RefreshToken) ProtoReflect() protoreflect.Message {
	mi := &file_server_internal_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshToken.ProtoReflect.Descriptor instead.
func (*RefreshToken) Descriptor() ([]byte, []int) {
	return file_server_internal_types_proto_rawDescGZIP(), []int{0}
}

func (x *RefreshToken) GetRefreshId() string {
	if x != nil {
		return x.RefreshId
	}
	return ""
}

func (x *RefreshToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *IDTokenSubject) Reset() {
	*x = IDTokenSubject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_internal_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDTokenSubject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDTokenSubject) ProtoMessage() {}

func (x *IDTokenSubject) ProtoReflect() protoreflect.Message {
	mi := &file_server_internal_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDTokenSubject.ProtoReflect.Descriptor instead.
func (*IDTokenSubject) Descriptor() ([]byte, []int) {
	return file_server_internal_types_proto_rawDescGZIP(), []int{1}
}

func (x *IDTokenSubject) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IDTokenSubject) GetConnId() string {
	if x != nil {
		return x.ConnId
	}
	return ""
}

var File_server_internal_types_proto protoreflect.FileDescriptor

var file_server_internal_types_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x43, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x42, 0x0a, 0x0e,
	0x49, 0x44, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x64,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x65, 0x78, 0x69, 0x64, 0x70, 0x2f, 0x64, 0x65, 0x78, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_server_internal_types_proto_rawDescOnce sync.Once
	file_server_internal_types_proto_rawDescData = file_server_internal_types_proto_rawDesc
)

func file_server_internal_types_proto_rawDescGZIP() []byte {
	file_server_internal_types_proto_rawDescOnce.Do(func() {
		file_server_internal_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_internal_types_proto_rawDescData)
	})
	return file_server_internal_types_proto_rawDescData
}

var file_server_internal_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_internal_types_proto_goTypes = []interface{}{
	(*RefreshToken)(nil),   // 0: internal.RefreshToken
	(*IDTokenSubject)(nil), // 1: internal.IDTokenSubject
}
var file_server_internal_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_internal_types_proto_init() }
func file_server_internal_types_proto_init() {
	if File_server_internal_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_internal_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_internal_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDTokenSubject); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_internal_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_internal_types_proto_goTypes,
		DependencyIndexes: file_server_internal_types_proto_depIdxs,
		MessageInfos:      file_server_internal_types_proto_msgTypes,
	}.Build()
	File_server_internal_types_proto = out.File
	file_server_internal_types_proto_rawDesc = nil
	file_server_internal_types_proto_goTypes = nil
	file_server_internal_types_proto_depIdxs = nil
}

const (
	// EncryptHeaderV1 ...
	EncryptHeaderV1 = "<enc-v1>"
	// SHA1 is the name of sha1 hash alg
	SHA1 = "sha1"
	// SHA256 is the name of sha256 hash alg
	SHA256 = "sha256"
)

// HashAlg used to get correct alg for hash
var HashAlg = map[string]func() hash.Hash{
	SHA1:   sha1.New,
	SHA256: sha256.New,
}

// Encrypt encrypts the content with salt
func Encrypt(content string, salt string, encrptAlg string) string {
	return fmt.Sprintf("%x", pbkdf2.Key([]byte(content), []byte(salt), 4096, 16, HashAlg[encrptAlg]))
}

// ReversibleEncrypt encrypts the str with aes/base64
func ReversibleEncrypt(str, key string) (string, error) {
	keyBytes := []byte(key)
	var block cipher.Block
	var err error

	if block, err = aes.NewCipher(keyBytes); err != nil {
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(str))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(str))
	encrypted := EncryptHeaderV1 + base64.StdEncoding.EncodeToString(cipherText)
	return encrypted, nil
}

// ReversibleDecrypt decrypts the str with aes/base64 or base 64 depending on "header"
func ReversibleDecrypt(str, key string) (string, error) {
	if strings.HasPrefix(str, EncryptHeaderV1) {
		str = str[len(EncryptHeaderV1):]
		return decryptAES(str, key)
	}
	// fallback to base64
	return decodeB64(str)
}

func decodeB64(str string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(str)
	return string(cipherText), err
}

func decryptAES(str, key string) (string, error) {
	keyBytes := []byte(key)
	var block cipher.Block
	var cipherText []byte
	var err error

	if block, err = aes.NewCipher(keyBytes); err != nil {
		return "", err
	}
	if cipherText, err = base64.StdEncoding.DecodeString(str); err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		err = errors.New("cipherText too short")
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}
