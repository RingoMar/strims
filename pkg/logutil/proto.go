// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package logutil

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/MemeLabs/strims/pkg/mathutil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Proto(key string, val proto.Message) zapcore.Field {
	return zap.Object(key, protoMarshaller{val.ProtoReflect()})
}

var _ zapcore.ObjectMarshaler = protoMarshaller{}
var _ zapcore.ObjectMarshaler = protoMapMarshaller{}
var _ zapcore.ArrayMarshaler = protoListMarshaller{}

type protoMarshaller struct {
	m protoreflect.Message
}

func (p protoMarshaller) MarshalLogObject(e zapcore.ObjectEncoder) error {
	fields := p.m.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		k := f.JSONName()
		v := p.m.Get(f)

		if f.IsMap() {
			if m := v.Map(); m.IsValid() {
				e.AddObject(k, protoMapMarshaller{f, m})
			}
		} else if f.IsList() {
			if m := v.List(); m.IsValid() {
				e.AddArray(k, protoListMarshaller{f, m})
			}
		} else {
			marshalProtoField(k, f, v, e)
		}
	}
	return nil
}

type protoMapMarshaller struct {
	f protoreflect.FieldDescriptor
	m protoreflect.Map
}

func (p protoMapMarshaller) MarshalLogObject(e zapcore.ObjectEncoder) error {
	p.m.Range(func(ki protoreflect.MapKey, vi protoreflect.Value) bool {
		var k string
		switch p.f.MapKey().Kind() {
		case protoreflect.BoolKind:
			k = strconv.FormatBool(vi.Bool())
		case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
			k = strconv.FormatInt(vi.Int(), 10)
		case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
			k = strconv.FormatUint(vi.Uint(), 10)
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			k = strconv.FormatFloat(vi.Float(), 'f', -1, 64)
		case protoreflect.StringKind:
			k = vi.String()
		}
		marshalProtoField(k, p.f.MapKey(), vi, e)
		return true
	})
	return nil
}

type protoListMarshaller struct {
	f protoreflect.FieldDescriptor
	m protoreflect.List
}

func (p protoListMarshaller) MarshalLogArray(e zapcore.ArrayEncoder) error {
	for i := 0; i < p.m.Len(); i++ {
		v := p.m.Get(i)
		switch p.f.Kind() {
		case protoreflect.BoolKind:
			e.AppendBool(v.Bool())
		case protoreflect.EnumKind:
			e.AppendInt32(int32(v.Enum()))
		case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
			e.AppendInt64(v.Int())
		case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
			e.AppendUint64(v.Uint())
		case protoreflect.FloatKind, protoreflect.DoubleKind:
			e.AppendFloat64(v.Float())
		case protoreflect.StringKind:
			e.AppendString(v.String())
		case protoreflect.BytesKind:
			e.AppendString(marshalProtoBytes(v.Bytes()))
		case protoreflect.MessageKind:
			e.AppendObject(protoMarshaller{v.Message()})
		}
	}
	return nil
}

func marshalProtoField(k string, f protoreflect.FieldDescriptor, v protoreflect.Value, e zapcore.ObjectEncoder) {
	switch f.Kind() {
	case protoreflect.BoolKind:
		if b := v.Bool(); b {
			e.AddBool(k, b)
		}
	case protoreflect.EnumKind:
		if n := v.Enum(); n != 0 {
			e.AddInt32(k, int32(n))
		}
	case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sint32Kind, protoreflect.Sint64Kind, protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		if n := v.Int(); n != 0 {
			e.AddInt64(k, n)
		}
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind, protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		if n := v.Uint(); n != 0 {
			e.AddUint64(k, n)
		}
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		if n := v.Float(); n != 0 {
			e.AddFloat64(k, n)
		}
	case protoreflect.StringKind:
		if s := v.String(); s != "" {
			e.AddString(k, s)
		}
	case protoreflect.BytesKind:
		if b := v.Bytes(); len(b) != 0 {
			e.AddString(k, marshalProtoBytes(b))
		}
	case protoreflect.MessageKind:
		if m := v.Message(); m.IsValid() {
			e.AddObject(k, protoMarshaller{m})
		}
	}
}

func marshalProtoBytes(b []byte) string {
	n := len(b)
	s := base64.RawStdEncoding.EncodeToString(b[:mathutil.Min(64, n)])
	switch {
	case n <= 64:
		return s
	case n < 1<<10:
		return fmt.Sprintf("%s... (%dbytes)", s, n)
	case n < 1<<20:
		return fmt.Sprintf("%s... (%.2fkB)", s, float64(n)/float64(1<<10))
	case n < 1<<30:
		return fmt.Sprintf("%s... (%.2fMB)", s, float64(n)/float64(1<<20))
	default:
		return fmt.Sprintf("%s... (%.2fGB)", s, float64(n)/float64(1<<30))
	}
}
