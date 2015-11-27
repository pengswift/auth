// Code generated by protoc-gen-go.
// source: sdk.proto
// DO NOT EDIT!

package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SDKType int32

const (
	SDKType_SDK_UNKNOWN  SDKType = 0
	SDKType_SDK_OFFICIAL SDKType = 1
	SDKType_SDK_QQ       SDKType = 2
	SDKType_SDK_WEIXIN   SDKType = 3
	SDKType_SDK_WEIBO    SDKType = 4
)

var SDKType_name = map[int32]string{
	0: "SDK_UNKNOWN",
	1: "SDK_OFFICIAL",
	2: "SDK_QQ",
	3: "SDK_WEIXIN",
	4: "SDK_WEIBO",
}
var SDKType_value = map[string]int32{
	"SDK_UNKNOWN":  0,
	"SDK_OFFICIAL": 1,
	"SDK_QQ":       2,
	"SDK_WEIXIN":   3,
	"SDK_WEIBO":    4,
}

func (x SDKType) String() string {
	return proto.EnumName(SDKType_name, int32(x))
}

func init() {
	proto.RegisterEnum("auth.SDKType", SDKType_name, SDKType_value)
}