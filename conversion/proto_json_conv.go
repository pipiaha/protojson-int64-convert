package conversion

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

// Convert convert json: int64 string to number.
// fuck protojson btw
func Convert(json []byte, msg proto.Message) []byte {
	m := msg.ProtoReflect()
	json = convertMsg(json, m)
	return json
}

func convertMsg(json []byte, m protoreflect.Message) []byte {
	if m == nil {
		return json
	}
	for i := 0; i < m.Descriptor().Fields().Len(); i++ {
		fd := m.Descriptor().Fields().Get(i)
		v := m.Get(fd)
		json = convertValue(json, fd, v)
	}
	return json
}

func convertValue(json []byte, fd protoreflect.FieldDescriptor, v protoreflect.Value) []byte {
	switch {
	case fd.IsList():
		return convertList(json, fd, v.List())
	case fd.IsMap():
		return convertMap(json, fd, v)
	default:
		return convertSingular(json, fd, v)
	}
}

func convertList(json []byte, fd protoreflect.FieldDescriptor, list protoreflect.List) []byte {
	for i := 0; i < list.Len(); i++ {
		item := list.Get(i)
		if err := convertSingular(json, fd, item); err != nil {
			return err
		}
	}
	return json
}

func convertMap(json []byte, fd protoreflect.FieldDescriptor, v protoreflect.Value) []byte {
	v.Map().Range(func(key protoreflect.MapKey, value protoreflect.Value) bool {
		json = convertValue(json, fd.MapValue(), value)
		return true
	})
	return json
}

func convertSingular(json []byte, fd protoreflect.FieldDescriptor, v protoreflect.Value) []byte {
	switch fd.Kind() {
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind,
		protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind:
		name := fd.JSONName()
		old := fmt.Sprintf("\"%s\":\"%d\"", name, v.Int())
		n := fmt.Sprintf("\"%s\":%d", name, v.Int())
		str := strings.Replace(string(json), old, n, -1)
		json = []byte(str)
	case protoreflect.MessageKind:
		json = convertMsg(json, v.Message())
	}
	return json
}
