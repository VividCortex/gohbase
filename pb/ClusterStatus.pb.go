// Code generated by protoc-gen-go.
// source: ClusterStatus.proto
// DO NOT EDIT!

package pb

import proto "github.com/VividCortex/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type RegionState_State int32

const (
	RegionState_OFFLINE       RegionState_State = 0
	RegionState_PENDING_OPEN  RegionState_State = 1
	RegionState_OPENING       RegionState_State = 2
	RegionState_OPEN          RegionState_State = 3
	RegionState_PENDING_CLOSE RegionState_State = 4
	RegionState_CLOSING       RegionState_State = 5
	RegionState_CLOSED        RegionState_State = 6
	RegionState_SPLITTING     RegionState_State = 7
	RegionState_SPLIT         RegionState_State = 8
	RegionState_FAILED_OPEN   RegionState_State = 9
	RegionState_FAILED_CLOSE  RegionState_State = 10
	RegionState_MERGING       RegionState_State = 11
	RegionState_MERGED        RegionState_State = 12
	RegionState_SPLITTING_NEW RegionState_State = 13
	// region but hasn't be created yet, or master doesn't
	// know it's already created
	RegionState_MERGING_NEW RegionState_State = 14
)

var RegionState_State_name = map[int32]string{
	0:  "OFFLINE",
	1:  "PENDING_OPEN",
	2:  "OPENING",
	3:  "OPEN",
	4:  "PENDING_CLOSE",
	5:  "CLOSING",
	6:  "CLOSED",
	7:  "SPLITTING",
	8:  "SPLIT",
	9:  "FAILED_OPEN",
	10: "FAILED_CLOSE",
	11: "MERGING",
	12: "MERGED",
	13: "SPLITTING_NEW",
	14: "MERGING_NEW",
}
var RegionState_State_value = map[string]int32{
	"OFFLINE":       0,
	"PENDING_OPEN":  1,
	"OPENING":       2,
	"OPEN":          3,
	"PENDING_CLOSE": 4,
	"CLOSING":       5,
	"CLOSED":        6,
	"SPLITTING":     7,
	"SPLIT":         8,
	"FAILED_OPEN":   9,
	"FAILED_CLOSE":  10,
	"MERGING":       11,
	"MERGED":        12,
	"SPLITTING_NEW": 13,
	"MERGING_NEW":   14,
}

func (x RegionState_State) Enum() *RegionState_State {
	p := new(RegionState_State)
	*p = x
	return p
}
func (x RegionState_State) String() string {
	return proto.EnumName(RegionState_State_name, int32(x))
}
func (x *RegionState_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RegionState_State_value, data, "RegionState_State")
	if err != nil {
		return err
	}
	*x = RegionState_State(value)
	return nil
}

type RegionState struct {
	RegionInfo       *RegionInfo        `protobuf:"bytes,1,req,name=region_info" json:"region_info,omitempty"`
	State            *RegionState_State `protobuf:"varint,2,req,name=state,enum=pb.RegionState_State" json:"state,omitempty"`
	Stamp            *uint64            `protobuf:"varint,3,opt,name=stamp" json:"stamp,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *RegionState) Reset()         { *m = RegionState{} }
func (m *RegionState) String() string { return proto.CompactTextString(m) }
func (*RegionState) ProtoMessage()    {}

func (m *RegionState) GetRegionInfo() *RegionInfo {
	if m != nil {
		return m.RegionInfo
	}
	return nil
}

func (m *RegionState) GetState() RegionState_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return RegionState_OFFLINE
}

func (m *RegionState) GetStamp() uint64 {
	if m != nil && m.Stamp != nil {
		return *m.Stamp
	}
	return 0
}

type RegionInTransition struct {
	Spec             *RegionSpecifier `protobuf:"bytes,1,req,name=spec" json:"spec,omitempty"`
	RegionState      *RegionState     `protobuf:"bytes,2,req,name=region_state" json:"region_state,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *RegionInTransition) Reset()         { *m = RegionInTransition{} }
func (m *RegionInTransition) String() string { return proto.CompactTextString(m) }
func (*RegionInTransition) ProtoMessage()    {}

func (m *RegionInTransition) GetSpec() *RegionSpecifier {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *RegionInTransition) GetRegionState() *RegionState {
	if m != nil {
		return m.RegionState
	}
	return nil
}

// *
// sequence Id of a store
type StoreSequenceId struct {
	FamilyName       []byte  `protobuf:"bytes,1,req,name=family_name" json:"family_name,omitempty"`
	SequenceId       *uint64 `protobuf:"varint,2,req,name=sequence_id" json:"sequence_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StoreSequenceId) Reset()         { *m = StoreSequenceId{} }
func (m *StoreSequenceId) String() string { return proto.CompactTextString(m) }
func (*StoreSequenceId) ProtoMessage()    {}

func (m *StoreSequenceId) GetFamilyName() []byte {
	if m != nil {
		return m.FamilyName
	}
	return nil
}

func (m *StoreSequenceId) GetSequenceId() uint64 {
	if m != nil && m.SequenceId != nil {
		return *m.SequenceId
	}
	return 0
}

// *
// contains a sequence id of a region which should be the minimum of its store sequence ids and
// list of sequence ids of the region's stores
type RegionStoreSequenceIds struct {
	LastFlushedSequenceId *uint64            `protobuf:"varint,1,req,name=last_flushed_sequence_id" json:"last_flushed_sequence_id,omitempty"`
	StoreSequenceId       []*StoreSequenceId `protobuf:"bytes,2,rep,name=store_sequence_id" json:"store_sequence_id,omitempty"`
	XXX_unrecognized      []byte             `json:"-"`
}

func (m *RegionStoreSequenceIds) Reset()         { *m = RegionStoreSequenceIds{} }
func (m *RegionStoreSequenceIds) String() string { return proto.CompactTextString(m) }
func (*RegionStoreSequenceIds) ProtoMessage()    {}

func (m *RegionStoreSequenceIds) GetLastFlushedSequenceId() uint64 {
	if m != nil && m.LastFlushedSequenceId != nil {
		return *m.LastFlushedSequenceId
	}
	return 0
}

func (m *RegionStoreSequenceIds) GetStoreSequenceId() []*StoreSequenceId {
	if m != nil {
		return m.StoreSequenceId
	}
	return nil
}

type RegionLoad struct {
	// * the region specifier
	RegionSpecifier *RegionSpecifier `protobuf:"bytes,1,req,name=region_specifier" json:"region_specifier,omitempty"`
	// * the number of stores for the region
	Stores *uint32 `protobuf:"varint,2,opt,name=stores" json:"stores,omitempty"`
	// * the number of storefiles for the region
	Storefiles *uint32 `protobuf:"varint,3,opt,name=storefiles" json:"storefiles,omitempty"`
	// * the total size of the store files for the region, uncompressed, in MB
	StoreUncompressedSize_MB *uint32 `protobuf:"varint,4,opt,name=store_uncompressed_size_MB" json:"store_uncompressed_size_MB,omitempty"`
	// * the current total size of the store files for the region, in MB
	StorefileSize_MB *uint32 `protobuf:"varint,5,opt,name=storefile_size_MB" json:"storefile_size_MB,omitempty"`
	// * the current size of the memstore for the region, in MB
	MemstoreSize_MB *uint32 `protobuf:"varint,6,opt,name=memstore_size_MB" json:"memstore_size_MB,omitempty"`
	// *
	// The current total size of root-level store file indexes for the region,
	// in MB. The same as {@link #rootIndexSizeKB} but in MB.
	StorefileIndexSize_MB *uint32 `protobuf:"varint,7,opt,name=storefile_index_size_MB" json:"storefile_index_size_MB,omitempty"`
	// * the current total read requests made to region
	ReadRequestsCount *uint64 `protobuf:"varint,8,opt,name=read_requests_count" json:"read_requests_count,omitempty"`
	// * the current total write requests made to region
	WriteRequestsCount *uint64 `protobuf:"varint,9,opt,name=write_requests_count" json:"write_requests_count,omitempty"`
	// * the total compacting key values in currently running compaction
	TotalCompacting_KVs *uint64 `protobuf:"varint,10,opt,name=total_compacting_KVs" json:"total_compacting_KVs,omitempty"`
	// * the completed count of key values in currently running compaction
	CurrentCompacted_KVs *uint64 `protobuf:"varint,11,opt,name=current_compacted_KVs" json:"current_compacted_KVs,omitempty"`
	// * The current total size of root-level indexes for the region, in KB.
	RootIndexSize_KB *uint32 `protobuf:"varint,12,opt,name=root_index_size_KB" json:"root_index_size_KB,omitempty"`
	// * The total size of all index blocks, not just the root level, in KB.
	TotalStaticIndexSize_KB *uint32 `protobuf:"varint,13,opt,name=total_static_index_size_KB" json:"total_static_index_size_KB,omitempty"`
	// *
	// The total size of all Bloom filter blocks, not just loaded into the
	// block cache, in KB.
	TotalStaticBloomSize_KB *uint32 `protobuf:"varint,14,opt,name=total_static_bloom_size_KB" json:"total_static_bloom_size_KB,omitempty"`
	// * the most recent sequence Id from cache flush
	CompleteSequenceId *uint64 `protobuf:"varint,15,opt,name=complete_sequence_id" json:"complete_sequence_id,omitempty"`
	// * The current data locality for region in the regionserver
	DataLocality          *float32 `protobuf:"fixed32,16,opt,name=data_locality" json:"data_locality,omitempty"`
	LastMajorCompactionTs *uint64  `protobuf:"varint,17,opt,name=last_major_compaction_ts,def=0" json:"last_major_compaction_ts,omitempty"`
	// * the most recent sequence Id of store from cache flush
	StoreCompleteSequenceId []*StoreSequenceId `protobuf:"bytes,18,rep,name=store_complete_sequence_id" json:"store_complete_sequence_id,omitempty"`
	XXX_unrecognized        []byte             `json:"-"`
}

func (m *RegionLoad) Reset()         { *m = RegionLoad{} }
func (m *RegionLoad) String() string { return proto.CompactTextString(m) }
func (*RegionLoad) ProtoMessage()    {}

const Default_RegionLoad_LastMajorCompactionTs uint64 = 0

func (m *RegionLoad) GetRegionSpecifier() *RegionSpecifier {
	if m != nil {
		return m.RegionSpecifier
	}
	return nil
}

func (m *RegionLoad) GetStores() uint32 {
	if m != nil && m.Stores != nil {
		return *m.Stores
	}
	return 0
}

func (m *RegionLoad) GetStorefiles() uint32 {
	if m != nil && m.Storefiles != nil {
		return *m.Storefiles
	}
	return 0
}

func (m *RegionLoad) GetStoreUncompressedSize_MB() uint32 {
	if m != nil && m.StoreUncompressedSize_MB != nil {
		return *m.StoreUncompressedSize_MB
	}
	return 0
}

func (m *RegionLoad) GetStorefileSize_MB() uint32 {
	if m != nil && m.StorefileSize_MB != nil {
		return *m.StorefileSize_MB
	}
	return 0
}

func (m *RegionLoad) GetMemstoreSize_MB() uint32 {
	if m != nil && m.MemstoreSize_MB != nil {
		return *m.MemstoreSize_MB
	}
	return 0
}

func (m *RegionLoad) GetStorefileIndexSize_MB() uint32 {
	if m != nil && m.StorefileIndexSize_MB != nil {
		return *m.StorefileIndexSize_MB
	}
	return 0
}

func (m *RegionLoad) GetReadRequestsCount() uint64 {
	if m != nil && m.ReadRequestsCount != nil {
		return *m.ReadRequestsCount
	}
	return 0
}

func (m *RegionLoad) GetWriteRequestsCount() uint64 {
	if m != nil && m.WriteRequestsCount != nil {
		return *m.WriteRequestsCount
	}
	return 0
}

func (m *RegionLoad) GetTotalCompacting_KVs() uint64 {
	if m != nil && m.TotalCompacting_KVs != nil {
		return *m.TotalCompacting_KVs
	}
	return 0
}

func (m *RegionLoad) GetCurrentCompacted_KVs() uint64 {
	if m != nil && m.CurrentCompacted_KVs != nil {
		return *m.CurrentCompacted_KVs
	}
	return 0
}

func (m *RegionLoad) GetRootIndexSize_KB() uint32 {
	if m != nil && m.RootIndexSize_KB != nil {
		return *m.RootIndexSize_KB
	}
	return 0
}

func (m *RegionLoad) GetTotalStaticIndexSize_KB() uint32 {
	if m != nil && m.TotalStaticIndexSize_KB != nil {
		return *m.TotalStaticIndexSize_KB
	}
	return 0
}

func (m *RegionLoad) GetTotalStaticBloomSize_KB() uint32 {
	if m != nil && m.TotalStaticBloomSize_KB != nil {
		return *m.TotalStaticBloomSize_KB
	}
	return 0
}

func (m *RegionLoad) GetCompleteSequenceId() uint64 {
	if m != nil && m.CompleteSequenceId != nil {
		return *m.CompleteSequenceId
	}
	return 0
}

func (m *RegionLoad) GetDataLocality() float32 {
	if m != nil && m.DataLocality != nil {
		return *m.DataLocality
	}
	return 0
}

func (m *RegionLoad) GetLastMajorCompactionTs() uint64 {
	if m != nil && m.LastMajorCompactionTs != nil {
		return *m.LastMajorCompactionTs
	}
	return Default_RegionLoad_LastMajorCompactionTs
}

func (m *RegionLoad) GetStoreCompleteSequenceId() []*StoreSequenceId {
	if m != nil {
		return m.StoreCompleteSequenceId
	}
	return nil
}

type ReplicationLoadSink struct {
	AgeOfLastAppliedOp        *uint64 `protobuf:"varint,1,req,name=ageOfLastAppliedOp" json:"ageOfLastAppliedOp,omitempty"`
	TimeStampsOfLastAppliedOp *uint64 `protobuf:"varint,2,req,name=timeStampsOfLastAppliedOp" json:"timeStampsOfLastAppliedOp,omitempty"`
	XXX_unrecognized          []byte  `json:"-"`
}

func (m *ReplicationLoadSink) Reset()         { *m = ReplicationLoadSink{} }
func (m *ReplicationLoadSink) String() string { return proto.CompactTextString(m) }
func (*ReplicationLoadSink) ProtoMessage()    {}

func (m *ReplicationLoadSink) GetAgeOfLastAppliedOp() uint64 {
	if m != nil && m.AgeOfLastAppliedOp != nil {
		return *m.AgeOfLastAppliedOp
	}
	return 0
}

func (m *ReplicationLoadSink) GetTimeStampsOfLastAppliedOp() uint64 {
	if m != nil && m.TimeStampsOfLastAppliedOp != nil {
		return *m.TimeStampsOfLastAppliedOp
	}
	return 0
}

type ReplicationLoadSource struct {
	PeerID                   *string `protobuf:"bytes,1,req,name=peerID" json:"peerID,omitempty"`
	AgeOfLastShippedOp       *uint64 `protobuf:"varint,2,req,name=ageOfLastShippedOp" json:"ageOfLastShippedOp,omitempty"`
	SizeOfLogQueue           *uint32 `protobuf:"varint,3,req,name=sizeOfLogQueue" json:"sizeOfLogQueue,omitempty"`
	TimeStampOfLastShippedOp *uint64 `protobuf:"varint,4,req,name=timeStampOfLastShippedOp" json:"timeStampOfLastShippedOp,omitempty"`
	ReplicationLag           *uint64 `protobuf:"varint,5,req,name=replicationLag" json:"replicationLag,omitempty"`
	XXX_unrecognized         []byte  `json:"-"`
}

func (m *ReplicationLoadSource) Reset()         { *m = ReplicationLoadSource{} }
func (m *ReplicationLoadSource) String() string { return proto.CompactTextString(m) }
func (*ReplicationLoadSource) ProtoMessage()    {}

func (m *ReplicationLoadSource) GetPeerID() string {
	if m != nil && m.PeerID != nil {
		return *m.PeerID
	}
	return ""
}

func (m *ReplicationLoadSource) GetAgeOfLastShippedOp() uint64 {
	if m != nil && m.AgeOfLastShippedOp != nil {
		return *m.AgeOfLastShippedOp
	}
	return 0
}

func (m *ReplicationLoadSource) GetSizeOfLogQueue() uint32 {
	if m != nil && m.SizeOfLogQueue != nil {
		return *m.SizeOfLogQueue
	}
	return 0
}

func (m *ReplicationLoadSource) GetTimeStampOfLastShippedOp() uint64 {
	if m != nil && m.TimeStampOfLastShippedOp != nil {
		return *m.TimeStampOfLastShippedOp
	}
	return 0
}

func (m *ReplicationLoadSource) GetReplicationLag() uint64 {
	if m != nil && m.ReplicationLag != nil {
		return *m.ReplicationLag
	}
	return 0
}

type ServerLoad struct {
	// * Number of requests since last report.
	NumberOfRequests *uint64 `protobuf:"varint,1,opt,name=number_of_requests" json:"number_of_requests,omitempty"`
	// * Total Number of requests from the start of the region server.
	TotalNumberOfRequests *uint64 `protobuf:"varint,2,opt,name=total_number_of_requests" json:"total_number_of_requests,omitempty"`
	// * the amount of used heap, in MB.
	UsedHeap_MB *uint32 `protobuf:"varint,3,opt,name=used_heap_MB" json:"used_heap_MB,omitempty"`
	// * the maximum allowable size of the heap, in MB.
	MaxHeap_MB *uint32 `protobuf:"varint,4,opt,name=max_heap_MB" json:"max_heap_MB,omitempty"`
	// * Information on the load of individual regions.
	RegionLoads []*RegionLoad `protobuf:"bytes,5,rep,name=region_loads" json:"region_loads,omitempty"`
	// *
	// Regionserver-level coprocessors, e.g., WALObserver implementations.
	// Region-level coprocessors, on the other hand, are stored inside RegionLoad
	// objects.
	Coprocessors []*Coprocessor `protobuf:"bytes,6,rep,name=coprocessors" json:"coprocessors,omitempty"`
	// *
	// Time when incremental (non-total) counts began being calculated (e.g. number_of_requests)
	// time is measured as the difference, measured in milliseconds, between the current time
	// and midnight, January 1, 1970 UTC.
	ReportStartTime *uint64 `protobuf:"varint,7,opt,name=report_start_time" json:"report_start_time,omitempty"`
	// *
	// Time when report was generated.
	// time is measured as the difference, measured in milliseconds, between the current time
	// and midnight, January 1, 1970 UTC.
	ReportEndTime *uint64 `protobuf:"varint,8,opt,name=report_end_time" json:"report_end_time,omitempty"`
	// *
	// The port number that this region server is hosing an info server on.
	InfoServerPort *uint32 `protobuf:"varint,9,opt,name=info_server_port" json:"info_server_port,omitempty"`
	// *
	// The replicationLoadSource for the replication Source status of this region server.
	ReplLoadSource []*ReplicationLoadSource `protobuf:"bytes,10,rep,name=replLoadSource" json:"replLoadSource,omitempty"`
	// *
	// The replicationLoadSink for the replication Sink status of this region server.
	ReplLoadSink     *ReplicationLoadSink `protobuf:"bytes,11,opt,name=replLoadSink" json:"replLoadSink,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *ServerLoad) Reset()         { *m = ServerLoad{} }
func (m *ServerLoad) String() string { return proto.CompactTextString(m) }
func (*ServerLoad) ProtoMessage()    {}

func (m *ServerLoad) GetNumberOfRequests() uint64 {
	if m != nil && m.NumberOfRequests != nil {
		return *m.NumberOfRequests
	}
	return 0
}

func (m *ServerLoad) GetTotalNumberOfRequests() uint64 {
	if m != nil && m.TotalNumberOfRequests != nil {
		return *m.TotalNumberOfRequests
	}
	return 0
}

func (m *ServerLoad) GetUsedHeap_MB() uint32 {
	if m != nil && m.UsedHeap_MB != nil {
		return *m.UsedHeap_MB
	}
	return 0
}

func (m *ServerLoad) GetMaxHeap_MB() uint32 {
	if m != nil && m.MaxHeap_MB != nil {
		return *m.MaxHeap_MB
	}
	return 0
}

func (m *ServerLoad) GetRegionLoads() []*RegionLoad {
	if m != nil {
		return m.RegionLoads
	}
	return nil
}

func (m *ServerLoad) GetCoprocessors() []*Coprocessor {
	if m != nil {
		return m.Coprocessors
	}
	return nil
}

func (m *ServerLoad) GetReportStartTime() uint64 {
	if m != nil && m.ReportStartTime != nil {
		return *m.ReportStartTime
	}
	return 0
}

func (m *ServerLoad) GetReportEndTime() uint64 {
	if m != nil && m.ReportEndTime != nil {
		return *m.ReportEndTime
	}
	return 0
}

func (m *ServerLoad) GetInfoServerPort() uint32 {
	if m != nil && m.InfoServerPort != nil {
		return *m.InfoServerPort
	}
	return 0
}

func (m *ServerLoad) GetReplLoadSource() []*ReplicationLoadSource {
	if m != nil {
		return m.ReplLoadSource
	}
	return nil
}

func (m *ServerLoad) GetReplLoadSink() *ReplicationLoadSink {
	if m != nil {
		return m.ReplLoadSink
	}
	return nil
}

type LiveServerInfo struct {
	Server           *ServerName `protobuf:"bytes,1,req,name=server" json:"server,omitempty"`
	ServerLoad       *ServerLoad `protobuf:"bytes,2,req,name=server_load" json:"server_load,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *LiveServerInfo) Reset()         { *m = LiveServerInfo{} }
func (m *LiveServerInfo) String() string { return proto.CompactTextString(m) }
func (*LiveServerInfo) ProtoMessage()    {}

func (m *LiveServerInfo) GetServer() *ServerName {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *LiveServerInfo) GetServerLoad() *ServerLoad {
	if m != nil {
		return m.ServerLoad
	}
	return nil
}

type ClusterStatus struct {
	HbaseVersion        *HBaseVersionFileContent `protobuf:"bytes,1,opt,name=hbase_version" json:"hbase_version,omitempty"`
	LiveServers         []*LiveServerInfo        `protobuf:"bytes,2,rep,name=live_servers" json:"live_servers,omitempty"`
	DeadServers         []*ServerName            `protobuf:"bytes,3,rep,name=dead_servers" json:"dead_servers,omitempty"`
	RegionsInTransition []*RegionInTransition    `protobuf:"bytes,4,rep,name=regions_in_transition" json:"regions_in_transition,omitempty"`
	ClusterId           *ClusterId               `protobuf:"bytes,5,opt,name=cluster_id" json:"cluster_id,omitempty"`
	MasterCoprocessors  []*Coprocessor           `protobuf:"bytes,6,rep,name=master_coprocessors" json:"master_coprocessors,omitempty"`
	Master              *ServerName              `protobuf:"bytes,7,opt,name=master" json:"master,omitempty"`
	BackupMasters       []*ServerName            `protobuf:"bytes,8,rep,name=backup_masters" json:"backup_masters,omitempty"`
	BalancerOn          *bool                    `protobuf:"varint,9,opt,name=balancer_on" json:"balancer_on,omitempty"`
	XXX_unrecognized    []byte                   `json:"-"`
}

func (m *ClusterStatus) Reset()         { *m = ClusterStatus{} }
func (m *ClusterStatus) String() string { return proto.CompactTextString(m) }
func (*ClusterStatus) ProtoMessage()    {}

func (m *ClusterStatus) GetHbaseVersion() *HBaseVersionFileContent {
	if m != nil {
		return m.HbaseVersion
	}
	return nil
}

func (m *ClusterStatus) GetLiveServers() []*LiveServerInfo {
	if m != nil {
		return m.LiveServers
	}
	return nil
}

func (m *ClusterStatus) GetDeadServers() []*ServerName {
	if m != nil {
		return m.DeadServers
	}
	return nil
}

func (m *ClusterStatus) GetRegionsInTransition() []*RegionInTransition {
	if m != nil {
		return m.RegionsInTransition
	}
	return nil
}

func (m *ClusterStatus) GetClusterId() *ClusterId {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *ClusterStatus) GetMasterCoprocessors() []*Coprocessor {
	if m != nil {
		return m.MasterCoprocessors
	}
	return nil
}

func (m *ClusterStatus) GetMaster() *ServerName {
	if m != nil {
		return m.Master
	}
	return nil
}

func (m *ClusterStatus) GetBackupMasters() []*ServerName {
	if m != nil {
		return m.BackupMasters
	}
	return nil
}

func (m *ClusterStatus) GetBalancerOn() bool {
	if m != nil && m.BalancerOn != nil {
		return *m.BalancerOn
	}
	return false
}

func init() {
	proto.RegisterEnum("pb.RegionState_State", RegionState_State_name, RegionState_State_value)
}
