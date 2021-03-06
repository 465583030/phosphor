// Code generated by protoc-gen-go.
// source: github.com/mondough/phosphor/proto/trace.proto
// DO NOT EDIT!

/*
Package traceproto is a generated protocol buffer package.

It is generated from these files:
	github.com/mondough/phosphor/proto/trace.proto

It has these top-level messages:
	Annotation
	KeyValue
*/
package traceproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AnnotationType int32

const (
	AnnotationType_UNKNOWN     AnnotationType = 0
	AnnotationType_CLIENT_SEND AnnotationType = 1
	AnnotationType_CLIENT_RECV AnnotationType = 2
	AnnotationType_SERVER_RECV AnnotationType = 3
	AnnotationType_SERVER_SEND AnnotationType = 4
	AnnotationType_TIMEOUT     AnnotationType = 5
	AnnotationType_ANNOTATION  AnnotationType = 6
)

var AnnotationType_name = map[int32]string{
	0: "UNKNOWN",
	1: "CLIENT_SEND",
	2: "CLIENT_RECV",
	3: "SERVER_RECV",
	4: "SERVER_SEND",
	5: "TIMEOUT",
	6: "ANNOTATION",
}
var AnnotationType_value = map[string]int32{
	"UNKNOWN":     0,
	"CLIENT_SEND": 1,
	"CLIENT_RECV": 2,
	"SERVER_RECV": 3,
	"SERVER_SEND": 4,
	"TIMEOUT":     5,
	"ANNOTATION":  6,
}

func (x AnnotationType) String() string {
	return proto.EnumName(AnnotationType_name, int32(x))
}

type Annotation struct {
	// The ID of the trace this annotation is a component of
	TraceId string `protobuf:"bytes,1,opt,name=trace_id" json:"trace_id,omitempty"`
	// The span this trace corresponds to, in the case this
	// is representing a service (REQ/REP) call
	SpanId string `protobuf:"bytes,2,opt,name=span_id" json:"span_id,omitempty"`
	// The parent span this trace corresponds to, allowing us
	// to correlate trace frames and reconstruct the request
	ParentId string `protobuf:"bytes,3,opt,name=parent_id" json:"parent_id,omitempty"`
	// The type of annotation we're capturing
	Type AnnotationType `protobuf:"varint,4,opt,name=type,enum=traceproto.AnnotationType" json:"type,omitempty"`
	// Flag to indicate this is an asynchronous span, which will not have a
	// response - eg. just client send and server recv annotations
	Async bool `protobuf:"varint,5,opt,name=async" json:"async,omitempty"`
	// Time since the epoch in microseconds
	Timestamp int64 `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	// Duration in microseconds
	// This should only be used to measure time on the same node
	// eg. the duration of service / rpc calls
	Duration int64 `protobuf:"varint,7,opt,name=duration" json:"duration,omitempty"`
	// Machine hostname, container name etc
	Hostname string `protobuf:"bytes,8,opt,name=hostname" json:"hostname,omitempty"`
	// Origin of this annotation, likely a service or application for a RPC
	Origin string `protobuf:"bytes,9,opt,name=origin" json:"origin,omitempty"`
	// Destination of this annotations action
	// eg. the service which a request was destined for
	// likely not set for annotations
	Destination string `protobuf:"bytes,10,opt,name=destination" json:"destination,omitempty"`
	// Payload as a string - eg. JSON encoded
	Payload string `protobuf:"bytes,11,opt,name=payload" json:"payload,omitempty"`
	// Repeated series of key value fields for arbitrary data
	KeyValue []*KeyValue `protobuf:"bytes,12,rep,name=key_value" json:"key_value,omitempty"`
}

func (m *Annotation) Reset()         { *m = Annotation{} }
func (m *Annotation) String() string { return proto.CompactTextString(m) }
func (*Annotation) ProtoMessage()    {}

func (m *Annotation) GetKeyValue() []*KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

type KeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("traceproto.AnnotationType", AnnotationType_name, AnnotationType_value)
}
