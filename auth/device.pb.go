// Code generated by protoc-gen-go.
// source: device.proto
// DO NOT EDIT!

package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Device struct {
	Os            string `protobuf:"bytes,1,opt,name=os" json:"os,omitempty"`
	OsVersion     string `protobuf:"bytes,2,opt,name=osVersion" json:"osVersion,omitempty"`
	Resolution    string `protobuf:"bytes,3,opt,name=resolution" json:"resolution,omitempty"`
	Ua            string `protobuf:"bytes,4,opt,name=ua" json:"ua,omitempty"`
	ClientVersion string `protobuf:"bytes,5,opt,name=clientVersion" json:"clientVersion,omitempty"`
	Mac           string `protobuf:"bytes,6,opt,name=mac" json:"mac,omitempty"`
	Udid          string `protobuf:"bytes,7,opt,name=udid" json:"udid,omitempty"`
	Idfa          string `protobuf:"bytes,8,opt,name=idfa" json:"idfa,omitempty"`
	Idfv          string `protobuf:"bytes,9,opt,name=idfv" json:"idfv,omitempty"`
	Appid         string `protobuf:"bytes,10,opt,name=appid" json:"appid,omitempty"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}

func init() {
	proto.RegisterType((*Device)(nil), "auth.Device")
}
