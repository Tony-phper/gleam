// Code generated by protoc-gen-go.
// source: execution_instruction.proto
// DO NOT EDIT!

package cmd

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type InstructionSet struct {
	Instructions     []*Instruction `protobuf:"bytes,1,rep,name=instructions" json:"instructions,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *InstructionSet) Reset()                    { *m = InstructionSet{} }
func (m *InstructionSet) String() string            { return proto.CompactTextString(m) }
func (*InstructionSet) ProtoMessage()               {}
func (*InstructionSet) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *InstructionSet) GetInstructions() []*Instruction {
	if m != nil {
		return m.Instructions
	}
	return nil
}

type Instruction struct {
	ScatterPartitions        *ScatterPartitions        `protobuf:"bytes,1,opt,name=scatterPartitions" json:"scatterPartitions,omitempty"`
	CollectPartitions        *CollectPartitions        `protobuf:"bytes,2,opt,name=collectPartitions" json:"collectPartitions,omitempty"`
	LocalSort                *LocalSort                `protobuf:"bytes,3,opt,name=localSort" json:"localSort,omitempty"`
	MergeSortedTo            *MergeSortedTo            `protobuf:"bytes,4,opt,name=mergeSortedTo" json:"mergeSortedTo,omitempty"`
	JoinPartitionedSorted    *JoinPartitionedSorted    `protobuf:"bytes,5,opt,name=joinPartitionedSorted" json:"joinPartitionedSorted,omitempty"`
	CoGroupPartitionedSorted *CoGroupPartitionedSorted `protobuf:"bytes,6,opt,name=coGroupPartitionedSorted" json:"coGroupPartitionedSorted,omitempty"`
	Script                   *Script                   `protobuf:"bytes,7,opt,name=script" json:"script,omitempty"`
	XXX_unrecognized         []byte                    `json:"-"`
}

func (m *Instruction) Reset()                    { *m = Instruction{} }
func (m *Instruction) String() string            { return proto.CompactTextString(m) }
func (*Instruction) ProtoMessage()               {}
func (*Instruction) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Instruction) GetScatterPartitions() *ScatterPartitions {
	if m != nil {
		return m.ScatterPartitions
	}
	return nil
}

func (m *Instruction) GetCollectPartitions() *CollectPartitions {
	if m != nil {
		return m.CollectPartitions
	}
	return nil
}

func (m *Instruction) GetLocalSort() *LocalSort {
	if m != nil {
		return m.LocalSort
	}
	return nil
}

func (m *Instruction) GetMergeSortedTo() *MergeSortedTo {
	if m != nil {
		return m.MergeSortedTo
	}
	return nil
}

func (m *Instruction) GetJoinPartitionedSorted() *JoinPartitionedSorted {
	if m != nil {
		return m.JoinPartitionedSorted
	}
	return nil
}

func (m *Instruction) GetCoGroupPartitionedSorted() *CoGroupPartitionedSorted {
	if m != nil {
		return m.CoGroupPartitionedSorted
	}
	return nil
}

func (m *Instruction) GetScript() *Script {
	if m != nil {
		return m.Script
	}
	return nil
}

type ScatterPartitions struct {
	InputShard           *DatasetShard           `protobuf:"bytes,1,req,name=InputShard,json=inputShard" json:"InputShard,omitempty"`
	OutputShardLocations []*DatasetShardLocation `protobuf:"bytes,2,rep,name=OutputShardLocations,json=outputShardLocations" json:"OutputShardLocations,omitempty"`
	ShardCount           *int32                  `protobuf:"varint,3,req,name=shardCount" json:"shardCount,omitempty"`
	XXX_unrecognized     []byte                  `json:"-"`
}

func (m *ScatterPartitions) Reset()                    { *m = ScatterPartitions{} }
func (m *ScatterPartitions) String() string            { return proto.CompactTextString(m) }
func (*ScatterPartitions) ProtoMessage()               {}
func (*ScatterPartitions) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ScatterPartitions) GetInputShard() *DatasetShard {
	if m != nil {
		return m.InputShard
	}
	return nil
}

func (m *ScatterPartitions) GetOutputShardLocations() []*DatasetShardLocation {
	if m != nil {
		return m.OutputShardLocations
	}
	return nil
}

func (m *ScatterPartitions) GetShardCount() int32 {
	if m != nil && m.ShardCount != nil {
		return *m.ShardCount
	}
	return 0
}

type CollectPartitions struct {
	InputShardLocations []*DatasetShardLocation `protobuf:"bytes,1,rep,name=inputShardLocations" json:"inputShardLocations,omitempty"`
	OutputShard         *DatasetShard           `protobuf:"bytes,2,req,name=OutputShard,json=outputShard" json:"OutputShard,omitempty"`
	XXX_unrecognized    []byte                  `json:"-"`
}

func (m *CollectPartitions) Reset()                    { *m = CollectPartitions{} }
func (m *CollectPartitions) String() string            { return proto.CompactTextString(m) }
func (*CollectPartitions) ProtoMessage()               {}
func (*CollectPartitions) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *CollectPartitions) GetInputShardLocations() []*DatasetShardLocation {
	if m != nil {
		return m.InputShardLocations
	}
	return nil
}

func (m *CollectPartitions) GetOutputShard() *DatasetShard {
	if m != nil {
		return m.OutputShard
	}
	return nil
}

type LocalSort struct {
	InputShard       *DatasetShard `protobuf:"bytes,1,req,name=InputShard,json=inputShard" json:"InputShard,omitempty"`
	OutputShard      *DatasetShard `protobuf:"bytes,2,req,name=OutputShard,json=outputShard" json:"OutputShard,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *LocalSort) Reset()                    { *m = LocalSort{} }
func (m *LocalSort) String() string            { return proto.CompactTextString(m) }
func (*LocalSort) ProtoMessage()               {}
func (*LocalSort) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *LocalSort) GetInputShard() *DatasetShard {
	if m != nil {
		return m.InputShard
	}
	return nil
}

func (m *LocalSort) GetOutputShard() *DatasetShard {
	if m != nil {
		return m.OutputShard
	}
	return nil
}

type MergeSortedTo struct {
	InputShardLocations []*DatasetShardLocation `protobuf:"bytes,1,rep,name=inputShardLocations" json:"inputShardLocations,omitempty"`
	OutputShard         *DatasetShard           `protobuf:"bytes,2,req,name=OutputShard,json=outputShard" json:"OutputShard,omitempty"`
	XXX_unrecognized    []byte                  `json:"-"`
}

func (m *MergeSortedTo) Reset()                    { *m = MergeSortedTo{} }
func (m *MergeSortedTo) String() string            { return proto.CompactTextString(m) }
func (*MergeSortedTo) ProtoMessage()               {}
func (*MergeSortedTo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *MergeSortedTo) GetInputShardLocations() []*DatasetShardLocation {
	if m != nil {
		return m.InputShardLocations
	}
	return nil
}

func (m *MergeSortedTo) GetOutputShard() *DatasetShard {
	if m != nil {
		return m.OutputShard
	}
	return nil
}

type JoinPartitionedSorted struct {
	LeftInputShardLocations  *DatasetShardLocation `protobuf:"bytes,1,req,name=leftInputShardLocations" json:"leftInputShardLocations,omitempty"`
	RightInputShardLocations *DatasetShardLocation `protobuf:"bytes,2,req,name=rightInputShardLocations" json:"rightInputShardLocations,omitempty"`
	OutputShard              *DatasetShard         `protobuf:"bytes,3,req,name=OutputShard,json=outputShard" json:"OutputShard,omitempty"`
	XXX_unrecognized         []byte                `json:"-"`
}

func (m *JoinPartitionedSorted) Reset()                    { *m = JoinPartitionedSorted{} }
func (m *JoinPartitionedSorted) String() string            { return proto.CompactTextString(m) }
func (*JoinPartitionedSorted) ProtoMessage()               {}
func (*JoinPartitionedSorted) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *JoinPartitionedSorted) GetLeftInputShardLocations() *DatasetShardLocation {
	if m != nil {
		return m.LeftInputShardLocations
	}
	return nil
}

func (m *JoinPartitionedSorted) GetRightInputShardLocations() *DatasetShardLocation {
	if m != nil {
		return m.RightInputShardLocations
	}
	return nil
}

func (m *JoinPartitionedSorted) GetOutputShard() *DatasetShard {
	if m != nil {
		return m.OutputShard
	}
	return nil
}

type CoGroupPartitionedSorted struct {
	LeftInputShardLocations  *DatasetShardLocation `protobuf:"bytes,1,req,name=leftInputShardLocations" json:"leftInputShardLocations,omitempty"`
	RightInputShardLocations *DatasetShardLocation `protobuf:"bytes,2,req,name=rightInputShardLocations" json:"rightInputShardLocations,omitempty"`
	OutputShard              *DatasetShard         `protobuf:"bytes,3,req,name=OutputShard,json=outputShard" json:"OutputShard,omitempty"`
	XXX_unrecognized         []byte                `json:"-"`
}

func (m *CoGroupPartitionedSorted) Reset()                    { *m = CoGroupPartitionedSorted{} }
func (m *CoGroupPartitionedSorted) String() string            { return proto.CompactTextString(m) }
func (*CoGroupPartitionedSorted) ProtoMessage()               {}
func (*CoGroupPartitionedSorted) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *CoGroupPartitionedSorted) GetLeftInputShardLocations() *DatasetShardLocation {
	if m != nil {
		return m.LeftInputShardLocations
	}
	return nil
}

func (m *CoGroupPartitionedSorted) GetRightInputShardLocations() *DatasetShardLocation {
	if m != nil {
		return m.RightInputShardLocations
	}
	return nil
}

func (m *CoGroupPartitionedSorted) GetOutputShard() *DatasetShard {
	if m != nil {
		return m.OutputShard
	}
	return nil
}

type Script struct {
	InputShard       *DatasetShard `protobuf:"bytes,1,req,name=InputShard,json=inputShard" json:"InputShard,omitempty"`
	OutputShard      *DatasetShard `protobuf:"bytes,2,req,name=OutputShard,json=outputShard" json:"OutputShard,omitempty"`
	Name             *string       `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	IsPipe           *bool         `protobuf:"varint,4,req,name=isPipe" json:"isPipe,omitempty"`
	Path             *string       `protobuf:"bytes,5,req,name=path" json:"path,omitempty"`
	Args             []string      `protobuf:"bytes,6,rep,name=args" json:"args,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Script) Reset()                    { *m = Script{} }
func (m *Script) String() string            { return proto.CompactTextString(m) }
func (*Script) ProtoMessage()               {}
func (*Script) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *Script) GetInputShard() *DatasetShard {
	if m != nil {
		return m.InputShard
	}
	return nil
}

func (m *Script) GetOutputShard() *DatasetShard {
	if m != nil {
		return m.OutputShard
	}
	return nil
}

func (m *Script) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Script) GetIsPipe() bool {
	if m != nil && m.IsPipe != nil {
		return *m.IsPipe
	}
	return false
}

func (m *Script) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *Script) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type DatasetShard struct {
	FlowName         *string `protobuf:"bytes,1,req,name=FlowName,json=flowName" json:"FlowName,omitempty"`
	DatasetId        *int32  `protobuf:"varint,2,req,name=DatasetId,json=datasetId" json:"DatasetId,omitempty"`
	DatasetShardId   *int32  `protobuf:"varint,3,req,name=DatasetShardId,json=datasetShardId" json:"DatasetShardId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DatasetShard) Reset()                    { *m = DatasetShard{} }
func (m *DatasetShard) String() string            { return proto.CompactTextString(m) }
func (*DatasetShard) ProtoMessage()               {}
func (*DatasetShard) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *DatasetShard) GetFlowName() string {
	if m != nil && m.FlowName != nil {
		return *m.FlowName
	}
	return ""
}

func (m *DatasetShard) GetDatasetId() int32 {
	if m != nil && m.DatasetId != nil {
		return *m.DatasetId
	}
	return 0
}

func (m *DatasetShard) GetDatasetShardId() int32 {
	if m != nil && m.DatasetShardId != nil {
		return *m.DatasetShardId
	}
	return 0
}

type DatasetShardLocation struct {
	Shard            *DatasetShard `protobuf:"bytes,1,req,name=shard" json:"shard,omitempty"`
	Host             *string       `protobuf:"bytes,2,req,name=Host,json=host" json:"Host,omitempty"`
	Port             *int32        `protobuf:"varint,3,req,name=Port,json=port" json:"Port,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *DatasetShardLocation) Reset()                    { *m = DatasetShardLocation{} }
func (m *DatasetShardLocation) String() string            { return proto.CompactTextString(m) }
func (*DatasetShardLocation) ProtoMessage()               {}
func (*DatasetShardLocation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *DatasetShardLocation) GetShard() *DatasetShard {
	if m != nil {
		return m.Shard
	}
	return nil
}

func (m *DatasetShardLocation) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *DatasetShardLocation) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*InstructionSet)(nil), "cmd.InstructionSet")
	proto.RegisterType((*Instruction)(nil), "cmd.Instruction")
	proto.RegisterType((*ScatterPartitions)(nil), "cmd.ScatterPartitions")
	proto.RegisterType((*CollectPartitions)(nil), "cmd.CollectPartitions")
	proto.RegisterType((*LocalSort)(nil), "cmd.LocalSort")
	proto.RegisterType((*MergeSortedTo)(nil), "cmd.MergeSortedTo")
	proto.RegisterType((*JoinPartitionedSorted)(nil), "cmd.JoinPartitionedSorted")
	proto.RegisterType((*CoGroupPartitionedSorted)(nil), "cmd.CoGroupPartitionedSorted")
	proto.RegisterType((*Script)(nil), "cmd.Script")
	proto.RegisterType((*DatasetShard)(nil), "cmd.DatasetShard")
	proto.RegisterType((*DatasetShardLocation)(nil), "cmd.DatasetShardLocation")
}

func init() { proto.RegisterFile("execution_instruction.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xdc, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0xed, 0x24, 0x5f, 0x7d, 0xdd, 0x46, 0x5f, 0x2e, 0x6d, 0x19, 0xca, 0x8f, 0x22, 0x23,
	0x41, 0x16, 0xa8, 0x12, 0x85, 0x05, 0xfb, 0x56, 0x85, 0x00, 0x85, 0x68, 0x02, 0x0b, 0x56, 0xc8,
	0xb2, 0xa7, 0x89, 0x91, 0xe3, 0xb1, 0x66, 0xc6, 0x82, 0xd7, 0x60, 0xc1, 0xa3, 0xf0, 0x06, 0xbc,
	0x01, 0x2f, 0x83, 0xc4, 0x06, 0xcd, 0xd8, 0x6e, 0x9c, 0xda, 0x69, 0x51, 0x17, 0x08, 0xb1, 0x9b,
	0xb9, 0xf7, 0x9c, 0x73, 0x7f, 0x3d, 0x86, 0x9b, 0xec, 0x13, 0x0b, 0x73, 0x15, 0xf3, 0xf4, 0x7d,
	0x9c, 0x4a, 0x25, 0xf2, 0x50, 0x9f, 0xf7, 0x33, 0xc1, 0x15, 0x47, 0x27, 0x5c, 0x44, 0xfe, 0x31,
	0xf4, 0xc7, 0x4b, 0xcf, 0x94, 0x29, 0x7c, 0x0c, 0x9b, 0x35, 0xac, 0x24, 0xd6, 0xd0, 0x19, 0x79,
	0x07, 0xff, 0xef, 0x87, 0x8b, 0x68, 0xbf, 0x06, 0xa5, 0x2b, 0x28, 0xff, 0xbb, 0x03, 0x5e, 0xcd,
	0x8b, 0x47, 0x30, 0x90, 0x61, 0xa0, 0x14, 0x13, 0x93, 0x40, 0xa8, 0xb8, 0x92, 0xb2, 0x46, 0xde,
	0xc1, 0xae, 0x91, 0x9a, 0x9e, 0xf7, 0xd2, 0x26, 0x41, 0xab, 0x84, 0x3c, 0x49, 0x58, 0xa8, 0x6a,
	0x2a, 0x76, 0x4d, 0xe5, 0xf0, 0xbc, 0x97, 0x36, 0x09, 0xf8, 0x00, 0xdc, 0x84, 0x87, 0x41, 0x32,
	0xe5, 0x42, 0x11, 0xc7, 0xb0, 0xfb, 0x86, 0xfd, 0xb2, 0xb2, 0xd2, 0x25, 0x00, 0x9f, 0xc0, 0xd6,
	0x82, 0x89, 0x19, 0xd3, 0x17, 0x16, 0xbd, 0xe1, 0xa4, 0x63, 0x18, 0x68, 0x18, 0x27, 0x75, 0x0f,
	0x5d, 0x05, 0xe2, 0x04, 0x76, 0x3e, 0xf0, 0x38, 0x3d, 0x8b, 0xcc, 0xa2, 0xc2, 0x45, 0xba, 0x46,
	0x61, 0xcf, 0x28, 0x3c, 0x6f, 0x43, 0xd0, 0x76, 0x22, 0xbe, 0x03, 0x12, 0xf2, 0xa7, 0x82, 0xe7,
	0x59, 0x53, 0xb4, 0x67, 0x44, 0x6f, 0x97, 0x6d, 0x68, 0x07, 0xd1, 0xb5, 0x74, 0xbc, 0x0b, 0x3d,
	0x19, 0x8a, 0x38, 0x53, 0xe4, 0x3f, 0x23, 0xe4, 0x95, 0x53, 0xd1, 0x26, 0x5a, 0xba, 0xfc, 0xaf,
	0x16, 0x0c, 0x1a, 0x83, 0xc2, 0x87, 0x00, 0xe3, 0x34, 0xcb, 0xd5, 0x74, 0x1e, 0x88, 0x88, 0x58,
	0x43, 0x7b, 0xe4, 0x1d, 0x0c, 0x0c, 0xfd, 0x28, 0x50, 0x81, 0x64, 0x85, 0x83, 0x42, 0x7c, 0x06,
	0xc2, 0x13, 0xd8, 0x7e, 0x9d, 0xab, 0xea, 0xaa, 0xfb, 0x5e, 0xcd, 0x52, 0x2f, 0xd7, 0x8d, 0x06,
	0xb9, 0x42, 0xd0, 0x6d, 0xde, 0x42, 0xc3, 0x3b, 0x00, 0x52, 0x5b, 0x0e, 0x79, 0x9e, 0xea, 0x91,
	0xda, 0xa3, 0x2e, 0xad, 0x59, 0xfc, 0x2f, 0x16, 0x0c, 0x1a, 0xab, 0x81, 0x2f, 0xe0, 0xda, 0x32,
	0xa5, 0x65, 0x0e, 0xd6, 0x65, 0x39, 0xb4, 0xb1, 0xf0, 0x11, 0x78, 0xb5, 0x8a, 0x88, 0xbd, 0xae,
	0x0b, 0x5e, 0xad, 0x00, 0x5f, 0x82, 0x7b, 0xb6, 0x73, 0x57, 0x69, 0xe3, 0x95, 0x82, 0x7e, 0xb6,
	0x60, 0x6b, 0x65, 0x6f, 0xff, 0x82, 0x46, 0xfc, 0xb0, 0x60, 0xa7, 0xf5, 0x4b, 0xc0, 0x29, 0x5c,
	0x4f, 0xd8, 0xa9, 0x1a, 0xb7, 0xe6, 0x67, 0x5f, 0x9c, 0xdf, 0x3a, 0x26, 0xbe, 0x05, 0x22, 0xe2,
	0xd9, 0xbc, 0x55, 0xd5, 0xbe, 0x4c, 0x75, 0x2d, 0xf5, 0x7c, 0xe9, 0xce, 0x6f, 0x95, 0xfe, 0xd3,
	0x02, 0xb2, 0xee, 0x7b, 0xfd, 0xf7, 0xab, 0xff, 0x66, 0x41, 0xaf, 0x78, 0x64, 0xfe, 0xd4, 0xfe,
	0x23, 0x42, 0x27, 0x0d, 0x16, 0xcc, 0x24, 0xe8, 0x52, 0x73, 0xc6, 0x5d, 0xe8, 0xc5, 0x72, 0x12,
	0x67, 0x8c, 0x74, 0x86, 0xf6, 0x68, 0x83, 0x96, 0x37, 0x8d, 0xcd, 0x02, 0x35, 0x27, 0xdd, 0x02,
	0xab, 0xcf, 0xda, 0x16, 0x88, 0x99, 0x24, 0xbd, 0xa1, 0xa3, 0x6d, 0xfa, 0xec, 0x67, 0xb0, 0x59,
	0x0f, 0x88, 0x7b, 0xb0, 0x71, 0x9c, 0xf0, 0x8f, 0xaf, 0x74, 0x1c, 0xcb, 0x70, 0x37, 0x4e, 0xcb,
	0x3b, 0xde, 0x02, 0xb7, 0xc4, 0x8e, 0x8b, 0x94, 0xbb, 0xd4, 0x8d, 0x2a, 0x03, 0xde, 0x83, 0x7e,
	0x5d, 0x69, 0x1c, 0x95, 0xcf, 0x59, 0x3f, 0x5a, 0xb1, 0xfa, 0x33, 0xd8, 0x6e, 0x9b, 0x0f, 0xde,
	0x87, 0xae, 0xbc, 0xb8, 0x81, 0x85, 0x5f, 0x97, 0xf1, 0x8c, 0x4b, 0x65, 0x32, 0x70, 0x69, 0x67,
	0xce, 0xa5, 0xd2, 0xb6, 0x49, 0xf1, 0x53, 0xd4, 0x21, 0x3b, 0x19, 0x17, 0xea, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x6a, 0xe6, 0x21, 0xa6, 0x34, 0x08, 0x00, 0x00,
}