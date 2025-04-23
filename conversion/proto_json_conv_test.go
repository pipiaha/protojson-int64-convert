package conversion

import (
	"google.golang.org/protobuf/encoding/protojson"
	"protojson-int64-convert/testcase/transfer"
	"testing"
	"time"
)

func TestInt64Conversion(t *testing.T) {
	arr := make([]*transfer.RedPointData, 0)
	arr = append(arr, &transfer.RedPointData{
		Module: 42,
		Cid:    0,
		Expire: time.Now().UnixMilli(),
	})
	pbval := &transfer.BadgeInfoResp{Data: arr}

	opt := protojson.MarshalOptions{EmitUnpopulated: true}
	bytes, _ := opt.Marshal(pbval)

	t.Log(string(bytes))
	bytes = Convert(bytes, pbval) // int64string->number
	t.Log(string(bytes))
	return
}
