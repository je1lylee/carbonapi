// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package nntbs

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
)

type NumericValueT struct {
	Type  NumericValue
	Value interface{}
}

func NumericValuePack(builder *flatbuffers.Builder, t *NumericValueT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case NumericValueIntValue:
		return IntValuePack(builder, t.Value.(*IntValueT))
	case NumericValueUintValue:
		return UintValuePack(builder, t.Value.(*UintValueT))
	case NumericValueLongValue:
		return LongValuePack(builder, t.Value.(*LongValueT))
	case NumericValueUlongValue:
		return UlongValuePack(builder, t.Value.(*UlongValueT))
	case NumericValueDoubleValue:
		return DoubleValuePack(builder, t.Value.(*DoubleValueT))
	case NumericValueAbsentValue:
		return AbsentValuePack(builder, t.Value.(*AbsentValueT))
	case NumericValueNullValue:
		return NullValuePack(builder, t.Value.(*NullValueT))
	}
	return 0
}

func NumericValueUnPack(t NumericValue, table flatbuffers.Table) *NumericValueT {
	switch t {
	case NumericValueIntValue:
		x := IntValue{_tab: table}
		return &NumericValueT{Type: NumericValueIntValue, Value: x.UnPack()}
	case NumericValueUintValue:
		x := UintValue{_tab: table}
		return &NumericValueT{Type: NumericValueUintValue, Value: x.UnPack()}
	case NumericValueLongValue:
		x := LongValue{_tab: table}
		return &NumericValueT{Type: NumericValueLongValue, Value: x.UnPack()}
	case NumericValueUlongValue:
		x := UlongValue{_tab: table}
		return &NumericValueT{Type: NumericValueUlongValue, Value: x.UnPack()}
	case NumericValueDoubleValue:
		x := DoubleValue{_tab: table}
		return &NumericValueT{Type: NumericValueDoubleValue, Value: x.UnPack()}
	case NumericValueAbsentValue:
		x := AbsentValue{_tab: table}
		return &NumericValueT{Type: NumericValueAbsentValue, Value: x.UnPack()}
	case NumericValueNullValue:
		x := NullValue{_tab: table}
		return &NumericValueT{Type: NumericValueNullValue, Value: x.UnPack()}
	}
	return nil
}

/*type ReplicationValueT struct {
	Type  ReplicationValue
	Value interface{}
}

func ReplicationValuePack(builder *flatbuffers.Builder, t *ReplicationValueT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	switch t.Type {
	case ReplicationValueNNT_V2:
		return NNT_V2Pack(builder, t.Value.(*NNT_V2T))
	case ReplicationValueText:
		return TextPack(builder, t.Value.(*TextT))
	case ReplicationValueHistogram:
		return HistogramPack(builder, t.Value.(*HistogramT))
	case ReplicationValueRawNumeric:
		return RawNumericPack(builder, t.Value.(*RawNumericT))
	case ReplicationValueLogMsg:
		return LogMsgPack(builder, t.Value.(*LogMsgT))
	case ReplicationValueLoadLuaMsg:
		return LoadLuaMsgPack(builder, t.Value.(*LoadLuaMsgT))
	case ReplicationValueDeleteMsg:
		return DeleteMsgPack(builder, t.Value.(*DeleteMsgT))
	case ReplicationValueSurrogatePut:
		return SurrogatePutPack(builder, t.Value.(*SurrogatePutT))
	case ReplicationValueActivityRebuild:
		return ActivityRebuildPack(builder, t.Value.(*ActivityRebuildT))
	case ReplicationValueSetCrdtAlter:
		return SetCrdtAlterPack(builder, t.Value.(*SetCrdtAlterT))
	case ReplicationValueSetCrdtMergeRequest:
		return SetCrdtMergeRequestPack(builder, t.Value.(*SetCrdtMergeRequestT))
	}
	return 0
}

func ReplicationValueUnPack(t ReplicationValue, table flatbuffers.Table) *ReplicationValueT {
	switch t {
	case ReplicationValueNNT_V2:
		x := NNT_V2{_tab: table}
		return &ReplicationValueT{Type: ReplicationValueNNT_V2, Value: x.UnPack()}
	case ReplicationValueText:
		x := Text{_tab: table}
		return &ReplicationValueT{Type: ReplicationValueText, Value: x.UnPack()}
	case ReplicationValueHistogram:
		x := Histogram{_tab: table}
		return &ReplicationValueT{Type: ReplicationValueHistogram, Value: x.UnPack()}
	case ReplicationValueRawNumeric:
		x := RawNumeric{_tab: table}
		return &ReplicationValueT{Type: ReplicationValueRawNumeric, Value: x.UnPack()}
	case ReplicationValueLogMsg:
		x := LogMsg{_tab: table}
		return &ReplicationValueT{Type: ReplicationValueLogMsg, Value: x.UnPack()}
	case ReplicationValueLoadLuaMsg:
		x := LoadLuaMsg{_tab: table}
		return &ReplicationValueT{Type: ReplicationValueLoadLuaMsg, Value: x.UnPack()}
	case ReplicationValueDeleteMsg:
		x := DeleteMsg{_tab: table}
		return &ReplicationValueT{Type: ReplicationValueDeleteMsg, Value: x.UnPack()}
	case ReplicationValueSurrogatePut:
		x := SurrogatePut{_tab: table}
		return &ReplicationValueT{Type: ReplicationValueSurrogatePut, Value: x.UnPack()}
	case ReplicationValueActivityRebuild:
		x := ActivityRebuild{_tab: table}
		return &ReplicationValueT{Type: ReplicationValueActivityRebuild, Value: x.UnPack()}
	case ReplicationValueSetCrdtAlter:
		x := SetCrdtAlter{_tab: table}
		return &ReplicationValueT{Type: ReplicationValueSetCrdtAlter, Value: x.UnPack()}
	case ReplicationValueSetCrdtMergeRequest:
		x := SetCrdtMergeRequest{_tab: table}
		return &ReplicationValueT{Type: ReplicationValueSetCrdtMergeRequest, Value: x.UnPack()}
	}
	return nil
}*/

type BlockType int8

const (
	BlockTypeAll        BlockType = 0
	BlockTypeAverage    BlockType = 1
	BlockTypeCounter    BlockType = 2
	BlockTypeDerivative BlockType = 3
)

var EnumNamesBlockType = map[BlockType]string{
	BlockTypeAll:        "All",
	BlockTypeAverage:    "Average",
	BlockTypeCounter:    "Counter",
	BlockTypeDerivative: "Derivative",
}

var EnumValuesBlockType = map[string]BlockType{
	"All":        BlockTypeAll,
	"Average":    BlockTypeAverage,
	"Counter":    BlockTypeCounter,
	"Derivative": BlockTypeDerivative,
}

func (v BlockType) String() string {
	if s, ok := EnumNamesBlockType[v]; ok {
		return s
	}
	return "BlockType(" + strconv.FormatInt(int64(v), 10) + ")"
}

type NNTDatumT struct {
	Zero             bool
	Count            uint32
	Stddev           float32
	Value            *NumericValueT
	Derivative       float32
	DerivativeStddev float32
	Counter          float32
	CounterStddev    float32
}

func NNTDatumPack(builder *flatbuffers.Builder, t *NNTDatumT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	valueOffset := NumericValuePack(builder, t.Value)

	NNTDatumStart(builder)
	NNTDatumAddZero(builder, t.Zero)
	NNTDatumAddCount(builder, t.Count)
	NNTDatumAddStddev(builder, t.Stddev)
	if t.Value != nil {
		NNTDatumAddValueType(builder, t.Value.Type)
	}
	NNTDatumAddValue(builder, valueOffset)
	NNTDatumAddDerivative(builder, t.Derivative)
	NNTDatumAddDerivativeStddev(builder, t.DerivativeStddev)
	NNTDatumAddCounter(builder, t.Counter)
	NNTDatumAddCounterStddev(builder, t.CounterStddev)
	return NNTDatumEnd(builder)
}

func (rcv *NNTDatum) UnPack() *NNTDatumT {
	if rcv == nil {
		return nil
	}
	t := &NNTDatumT{}
	t.Zero = rcv.Zero()
	t.Count = rcv.Count()
	t.Stddev = rcv.Stddev()
	valueTable := flatbuffers.Table{}
	if rcv.Value(&valueTable) {
		t.Value = NumericValueUnPack(rcv.ValueType(), valueTable)
	}
	t.Derivative = rcv.Derivative()
	t.DerivativeStddev = rcv.DerivativeStddev()
	t.Counter = rcv.Counter()
	t.CounterStddev = rcv.CounterStddev()
	return t
}

type NNTDatum struct {
	_tab flatbuffers.Table
}

func GetRootAsNNTDatum(buf []byte, offset flatbuffers.UOffsetT) *NNTDatum {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NNTDatum{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *NNTDatum) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NNTDatum) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NNTDatum) Zero() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *NNTDatum) MutateZero(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func (rcv *NNTDatum) Count() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NNTDatum) MutateCount(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *NNTDatum) Stddev() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *NNTDatum) MutateStddev(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func (rcv *NNTDatum) ValueType() NumericValue {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return NumericValue(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *NNTDatum) MutateValueType(n NumericValue) bool {
	return rcv._tab.MutateByteSlot(10, byte(n))
}

func (rcv *NNTDatum) Value(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *NNTDatum) Derivative() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *NNTDatum) MutateDerivative(n float32) bool {
	return rcv._tab.MutateFloat32Slot(14, n)
}

func (rcv *NNTDatum) DerivativeStddev() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *NNTDatum) MutateDerivativeStddev(n float32) bool {
	return rcv._tab.MutateFloat32Slot(16, n)
}

func (rcv *NNTDatum) Counter() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *NNTDatum) MutateCounter(n float32) bool {
	return rcv._tab.MutateFloat32Slot(18, n)
}

func (rcv *NNTDatum) CounterStddev() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *NNTDatum) MutateCounterStddev(n float32) bool {
	return rcv._tab.MutateFloat32Slot(20, n)
}

func NNTDatumStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
}
func NNTDatumAddZero(builder *flatbuffers.Builder, zero bool) {
	builder.PrependBoolSlot(0, zero, false)
}
func NNTDatumAddCount(builder *flatbuffers.Builder, count uint32) {
	builder.PrependUint32Slot(1, count, 0)
}
func NNTDatumAddStddev(builder *flatbuffers.Builder, stddev float32) {
	builder.PrependFloat32Slot(2, stddev, 0.0)
}
func NNTDatumAddValueType(builder *flatbuffers.Builder, valueType NumericValue) {
	builder.PrependByteSlot(3, byte(valueType), 0)
}
func NNTDatumAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(value), 0)
}
func NNTDatumAddDerivative(builder *flatbuffers.Builder, derivative float32) {
	builder.PrependFloat32Slot(5, derivative, 0.0)
}
func NNTDatumAddDerivativeStddev(builder *flatbuffers.Builder, derivativeStddev float32) {
	builder.PrependFloat32Slot(6, derivativeStddev, 0.0)
}
func NNTDatumAddCounter(builder *flatbuffers.Builder, counter float32) {
	builder.PrependFloat32Slot(7, counter, 0.0)
}
func NNTDatumAddCounterStddev(builder *flatbuffers.Builder, counterStddev float32) {
	builder.PrependFloat32Slot(8, counterStddev, 0.0)
}
func NNTDatumEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type NNTBlockT struct {
	Type       BlockType
	Crc        uint32
	Version    uint32
	BlockFloor uint64
	CreationUs uint64
	Data       []*NNTDatumT
}

func NNTBlockPack(builder *flatbuffers.Builder, t *NNTBlockT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	dataOffset := flatbuffers.UOffsetT(0)
	if t.Data != nil {
		dataLength := len(t.Data)
		dataOffsets := make([]flatbuffers.UOffsetT, dataLength)
		for j := 0; j < dataLength; j++ {
			dataOffsets[j] = NNTDatumPack(builder, t.Data[j])
		}
		NNTBlockStartDataVector(builder, dataLength)
		for j := dataLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(dataOffsets[j])
		}
		dataOffset = builder.EndVector(dataLength)
	}
	NNTBlockStart(builder)
	NNTBlockAddType(builder, t.Type)
	NNTBlockAddCrc(builder, t.Crc)
	NNTBlockAddVersion(builder, t.Version)
	NNTBlockAddBlockFloor(builder, t.BlockFloor)
	NNTBlockAddCreationUs(builder, t.CreationUs)
	NNTBlockAddData(builder, dataOffset)
	return NNTBlockEnd(builder)
}

func (rcv *NNTBlock) UnPack() *NNTBlockT {
	if rcv == nil {
		return nil
	}
	t := &NNTBlockT{}
	t.Type = rcv.Type()
	t.Crc = rcv.Crc()
	t.Version = rcv.Version()
	t.BlockFloor = rcv.BlockFloor()
	t.CreationUs = rcv.CreationUs()
	dataLength := rcv.DataLength()
	t.Data = make([]*NNTDatumT, dataLength)
	for j := 0; j < dataLength; j++ {
		x := NNTDatum{}
		rcv.Data(&x, j)
		t.Data[j] = x.UnPack()
	}
	return t
}

type NNTBlock struct {
	_tab flatbuffers.Table
}

func GetRootAsNNTBlock(buf []byte, offset flatbuffers.UOffsetT) *NNTBlock {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NNTBlock{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *NNTBlock) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NNTBlock) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NNTBlock) Type() BlockType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return BlockType(rcv._tab.GetInt8(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *NNTBlock) MutateType(n BlockType) bool {
	return rcv._tab.MutateInt8Slot(4, int8(n))
}

func (rcv *NNTBlock) Crc() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NNTBlock) MutateCrc(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *NNTBlock) Version() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NNTBlock) MutateVersion(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *NNTBlock) BlockFloor() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NNTBlock) MutateBlockFloor(n uint64) bool {
	return rcv._tab.MutateUint64Slot(10, n)
}

func (rcv *NNTBlock) CreationUs() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NNTBlock) MutateCreationUs(n uint64) bool {
	return rcv._tab.MutateUint64Slot(12, n)
}

func (rcv *NNTBlock) Data(obj *NNTDatum, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *NNTBlock) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func NNTBlockStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func NNTBlockAddType(builder *flatbuffers.Builder, type_ BlockType) {
	builder.PrependInt8Slot(0, int8(type_), 0)
}
func NNTBlockAddCrc(builder *flatbuffers.Builder, crc uint32) {
	builder.PrependUint32Slot(1, crc, 0)
}
func NNTBlockAddVersion(builder *flatbuffers.Builder, version uint32) {
	builder.PrependUint32Slot(2, version, 0)
}
func NNTBlockAddBlockFloor(builder *flatbuffers.Builder, blockFloor uint64) {
	builder.PrependUint64Slot(3, blockFloor, 0)
}
func NNTBlockAddCreationUs(builder *flatbuffers.Builder, creationUs uint64) {
	builder.PrependUint64Slot(4, creationUs, 0)
}
func NNTBlockAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(data), 0)
}
func NNTBlockStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func NNTBlockEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type NNTT struct {
	Epoch      uint64
	Apocalypse uint64
	Period     uint32
	Blocks     []*NNTBlockT
}

func NNTPack(builder *flatbuffers.Builder, t *NNTT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	blocksOffset := flatbuffers.UOffsetT(0)
	if t.Blocks != nil {
		blocksLength := len(t.Blocks)
		blocksOffsets := make([]flatbuffers.UOffsetT, blocksLength)
		for j := 0; j < blocksLength; j++ {
			blocksOffsets[j] = NNTBlockPack(builder, t.Blocks[j])
		}
		NNTStartBlocksVector(builder, blocksLength)
		for j := blocksLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(blocksOffsets[j])
		}
		blocksOffset = builder.EndVector(blocksLength)
	}
	NNTStart(builder)
	NNTAddEpoch(builder, t.Epoch)
	NNTAddApocalypse(builder, t.Apocalypse)
	NNTAddPeriod(builder, t.Period)
	NNTAddBlocks(builder, blocksOffset)
	return NNTEnd(builder)
}

func (rcv *NNT) UnPack() *NNTT {
	if rcv == nil {
		return nil
	}
	t := &NNTT{}
	t.Epoch = rcv.Epoch()
	t.Apocalypse = rcv.Apocalypse()
	t.Period = rcv.Period()
	blocksLength := rcv.BlocksLength()
	t.Blocks = make([]*NNTBlockT, blocksLength)
	for j := 0; j < blocksLength; j++ {
		x := NNTBlock{}
		rcv.Blocks(&x, j)
		t.Blocks[j] = x.UnPack()
	}
	return t
}

type NNT struct {
	_tab flatbuffers.Table
}

func GetRootAsNNT(buf []byte, offset flatbuffers.UOffsetT) *NNT {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NNT{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *NNT) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NNT) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NNT) Epoch() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NNT) MutateEpoch(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *NNT) Apocalypse() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NNT) MutateApocalypse(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *NNT) Period() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *NNT) MutatePeriod(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func (rcv *NNT) Blocks(obj *NNTBlock, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *NNT) BlocksLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func NNTStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func NNTAddEpoch(builder *flatbuffers.Builder, epoch uint64) {
	builder.PrependUint64Slot(0, epoch, 0)
}
func NNTAddApocalypse(builder *flatbuffers.Builder, apocalypse uint64) {
	builder.PrependUint64Slot(1, apocalypse, 0)
}
func NNTAddPeriod(builder *flatbuffers.Builder, period uint32) {
	builder.PrependUint32Slot(2, period, 0)
}
func NNTAddBlocks(builder *flatbuffers.Builder, blocks flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(blocks), 0)
}
func NNTStartBlocksVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func NNTEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type NNTMergeOpT struct {
	Metric *MetricInfoT
	Nnt    []*NNTT
}

func NNTMergeOpPack(builder *flatbuffers.Builder, t *NNTMergeOpT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	metricOffset := MetricInfoPack(builder, t.Metric)
	nntOffset := flatbuffers.UOffsetT(0)
	if t.Nnt != nil {
		nntLength := len(t.Nnt)
		nntOffsets := make([]flatbuffers.UOffsetT, nntLength)
		for j := 0; j < nntLength; j++ {
			nntOffsets[j] = NNTPack(builder, t.Nnt[j])
		}
		NNTMergeOpStartNntVector(builder, nntLength)
		for j := nntLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(nntOffsets[j])
		}
		nntOffset = builder.EndVector(nntLength)
	}
	NNTMergeOpStart(builder)
	NNTMergeOpAddMetric(builder, metricOffset)
	NNTMergeOpAddNnt(builder, nntOffset)
	return NNTMergeOpEnd(builder)
}

func (rcv *NNTMergeOp) UnPack() *NNTMergeOpT {
	if rcv == nil {
		return nil
	}
	t := &NNTMergeOpT{}
	t.Metric = rcv.Metric(nil).UnPack()
	nntLength := rcv.NntLength()
	t.Nnt = make([]*NNTT, nntLength)
	for j := 0; j < nntLength; j++ {
		x := NNT{}
		rcv.Nnt(&x, j)
		t.Nnt[j] = x.UnPack()
	}
	return t
}

type NNTMergeOp struct {
	_tab flatbuffers.Table
}

func GetRootAsNNTMergeOp(buf []byte, offset flatbuffers.UOffsetT) *NNTMergeOp {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NNTMergeOp{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *NNTMergeOp) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NNTMergeOp) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NNTMergeOp) Metric(obj *MetricInfo) *MetricInfo {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MetricInfo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *NNTMergeOp) Nnt(obj *NNT, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *NNTMergeOp) NntLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func NNTMergeOpStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func NNTMergeOpAddMetric(builder *flatbuffers.Builder, metric flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(metric), 0)
}
func NNTMergeOpAddNnt(builder *flatbuffers.Builder, nnt flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(nnt), 0)
}
func NNTMergeOpStartNntVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func NNTMergeOpEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type NNTMergeT struct {
	Ops []*NNTMergeOpT
}

func NNTMergePack(builder *flatbuffers.Builder, t *NNTMergeT) flatbuffers.UOffsetT {
	if t == nil {
		return 0
	}
	opsOffset := flatbuffers.UOffsetT(0)
	if t.Ops != nil {
		opsLength := len(t.Ops)
		opsOffsets := make([]flatbuffers.UOffsetT, opsLength)
		for j := 0; j < opsLength; j++ {
			opsOffsets[j] = NNTMergeOpPack(builder, t.Ops[j])
		}
		NNTMergeStartOpsVector(builder, opsLength)
		for j := opsLength - 1; j >= 0; j-- {
			builder.PrependUOffsetT(opsOffsets[j])
		}
		opsOffset = builder.EndVector(opsLength)
	}
	NNTMergeStart(builder)
	NNTMergeAddOps(builder, opsOffset)
	return NNTMergeEnd(builder)
}

func (rcv *NNTMerge) UnPack() *NNTMergeT {
	if rcv == nil {
		return nil
	}
	t := &NNTMergeT{}
	opsLength := rcv.OpsLength()
	t.Ops = make([]*NNTMergeOpT, opsLength)
	for j := 0; j < opsLength; j++ {
		x := NNTMergeOp{}
		rcv.Ops(&x, j)
		t.Ops[j] = x.UnPack()
	}
	return t
}

type NNTMerge struct {
	_tab flatbuffers.Table
}

func GetRootAsNNTMerge(buf []byte, offset flatbuffers.UOffsetT) *NNTMerge {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &NNTMerge{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *NNTMerge) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *NNTMerge) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *NNTMerge) Ops(obj *NNTMergeOp, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *NNTMerge) OpsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func NNTMergeStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func NNTMergeAddOps(builder *flatbuffers.Builder, ops flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ops), 0)
}
func NNTMergeStartOpsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func NNTMergeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}