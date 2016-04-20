// Code generated by protoc-gen-go.
// source: ClusterId.proto
// DO NOT EDIT!

package pb

import proto "github.com/VividCortex/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// *
// Content of the '/hbase/hbaseid', cluster id, znode.
// Also cluster of the ${HBASE_ROOTDIR}/hbase.id file.
type ClusterId struct {
	// This is the cluster id, a uuid as a String
	ClusterId        *string `protobuf:"bytes,1,req,name=cluster_id" json:"cluster_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClusterId) Reset()         { *m = ClusterId{} }
func (m *ClusterId) String() string { return proto.CompactTextString(m) }
func (*ClusterId) ProtoMessage()    {}

func (m *ClusterId) GetClusterId() string {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return ""
}

func init() {
}
