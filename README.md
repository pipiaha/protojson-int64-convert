# protojson-int64-convert

convert protojson.Marshall result json's string typed int64 into number

## Description

Extra int64 conversion for `protojson.Marshal` in `google.golang.org/protobuf/encoding/protojson` package.

* According to [protobuf](https://developers.google.com/protocol-buffers/docs/proto3#json), the JSON representation of a
  64-bit integer is a string. This is because The `protojson` package follows the standard JSON mapping, which states
  that 64-bit ints are encoded as a JSON string.
  issues can be found [here](https://github.com/golang/protobuf/issues/1414)
* `protojson` package has a lot of `internal` methods, so we can't just ovveride `protojson.MarshalOptions` to achieve
  the goal which is to convert the int64 to number in JSON.

* basic logics: 
  * use `protojson.Marshal` as usual.
  * traverse the proto descriptor to find the int64 fields.
  * convert the int64 string to number in json bytes.

## Usage

### 1. Install

```go
go get github.com/pipiaha/protojson-int64-convert
```

### 2. Import

```go
import "github.com/pipiaha/protojson-int64-convert/conversion"
```

### 3. Convert

```go
// some proto message: pbdata
arr := make([]*transfer.RedPointData, 0)
arr = append(arr, &transfer.RedPointData{
Module: 42,
Cid:    0,
Expire: time.Now().UnixMilli(),
})
pbdata := &transfer.BadgeInfoResp{Data: arr}

// Marshalling proto into json
opt := protojson.MarshalOptions{EmitUnpopulated: true}
bytes, _ := opt.Marshal(pbval)
// Expire is a int64, but in json it is a string

// convert string to number
bytes := conversion.Convert(jsonStr, pbdata)
// now Expire is a number
// bytes: {"data":[{"module":42,"cid":0,"expire":1697030400000}]}

```