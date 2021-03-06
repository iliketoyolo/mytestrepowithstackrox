package protoutils

import (
	"github.com/gogo/protobuf/proto"
)

func protoEqualWrapper(a, b interface{}) bool {
	return proto.Equal(a.(proto.Message), b.(proto.Message))
}
