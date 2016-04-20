// Code generated by protoc-gen-go.
// source: Cell.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	Cell.proto
	Client.proto
	ClusterId.proto
	ClusterStatus.proto
	Comparator.proto
	ErrorHandling.proto
	Filter.proto
	FS.proto
	HBase.proto
	Master.proto
	MultiRowMutation.proto
	Quota.proto
	RPC.proto
	Tracing.proto
	ZooKeeper.proto

It has these top-level messages:
	Cell
	KeyValue
*/
package pb

import proto "github.com/VividCortex/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// *
// The type of the key in a Cell
type CellType int32

const (
	CellType_MINIMUM       CellType = 0
	CellType_PUT           CellType = 4
	CellType_DELETE        CellType = 8
	CellType_DELETE_COLUMN CellType = 12
	CellType_DELETE_FAMILY CellType = 14
	// MAXIMUM is used when searching; you look from maximum on down.
	CellType_MAXIMUM CellType = 255
)

var CellType_name = map[int32]string{
	0:   "MINIMUM",
	4:   "PUT",
	8:   "DELETE",
	12:  "DELETE_COLUMN",
	14:  "DELETE_FAMILY",
	255: "MAXIMUM",
}
var CellType_value = map[string]int32{
	"MINIMUM":       0,
	"PUT":           4,
	"DELETE":        8,
	"DELETE_COLUMN": 12,
	"DELETE_FAMILY": 14,
	"MAXIMUM":       255,
}

func (x CellType) Enum() *CellType {
	p := new(CellType)
	*p = x
	return p
}
func (x CellType) String() string {
	return proto.EnumName(CellType_name, int32(x))
}
func (x *CellType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CellType_value, data, "CellType")
	if err != nil {
		return err
	}
	*x = CellType(value)
	return nil
}

// *
// Protocol buffer version of Cell.
type Cell struct {
	Row              []byte    `protobuf:"bytes,1,opt,name=row" json:"row,omitempty"`
	Family           []byte    `protobuf:"bytes,2,opt,name=family" json:"family,omitempty"`
	Qualifier        []byte    `protobuf:"bytes,3,opt,name=qualifier" json:"qualifier,omitempty"`
	Timestamp        *uint64   `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	CellType         *CellType `protobuf:"varint,5,opt,name=cell_type,enum=pb.CellType" json:"cell_type,omitempty"`
	Value            []byte    `protobuf:"bytes,6,opt,name=value" json:"value,omitempty"`
	Tags             []byte    `protobuf:"bytes,7,opt,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Cell) Reset()         { *m = Cell{} }
func (m *Cell) String() string { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()    {}

func (m *Cell) GetRow() []byte {
	if m != nil {
		return m.Row
	}
	return nil
}

func (m *Cell) GetFamily() []byte {
	if m != nil {
		return m.Family
	}
	return nil
}

func (m *Cell) GetQualifier() []byte {
	if m != nil {
		return m.Qualifier
	}
	return nil
}

func (m *Cell) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *Cell) GetCellType() CellType {
	if m != nil && m.CellType != nil {
		return *m.CellType
	}
	return CellType_MINIMUM
}

func (m *Cell) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Cell) GetTags() []byte {
	if m != nil {
		return m.Tags
	}
	return nil
}

// *
// Protocol buffer version of KeyValue.
// It doesn't have those transient parameters
type KeyValue struct {
	Row              []byte    `protobuf:"bytes,1,req,name=row" json:"row,omitempty"`
	Family           []byte    `protobuf:"bytes,2,req,name=family" json:"family,omitempty"`
	Qualifier        []byte    `protobuf:"bytes,3,req,name=qualifier" json:"qualifier,omitempty"`
	Timestamp        *uint64   `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	KeyType          *CellType `protobuf:"varint,5,opt,name=key_type,enum=pb.CellType" json:"key_type,omitempty"`
	Value            []byte    `protobuf:"bytes,6,opt,name=value" json:"value,omitempty"`
	Tags             []byte    `protobuf:"bytes,7,opt,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}

func (m *KeyValue) GetRow() []byte {
	if m != nil {
		return m.Row
	}
	return nil
}

func (m *KeyValue) GetFamily() []byte {
	if m != nil {
		return m.Family
	}
	return nil
}

func (m *KeyValue) GetQualifier() []byte {
	if m != nil {
		return m.Qualifier
	}
	return nil
}

func (m *KeyValue) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *KeyValue) GetKeyType() CellType {
	if m != nil && m.KeyType != nil {
		return *m.KeyType
	}
	return CellType_MINIMUM
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KeyValue) GetTags() []byte {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.CellType", CellType_name, CellType_value)
}
