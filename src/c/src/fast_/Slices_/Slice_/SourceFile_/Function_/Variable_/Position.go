// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Variable_

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Position struct {
	_tab flatbuffers.Table
}

func GetRootAsPosition(buf []byte, offset flatbuffers.UOffsetT) *Position {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Position{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Position) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Position) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Position) Location() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Position) Type() ChangeType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Position) MutateType(n ChangeType) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *Position) DeltaLocation() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func PositionStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func PositionAddLocation(builder *flatbuffers.Builder, location flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(location), 0)
}
func PositionAddType(builder *flatbuffers.Builder, type_ int32) {
	builder.PrependInt32Slot(1, type_, 0)
}
func PositionAddDeltaLocation(builder *flatbuffers.Builder, deltaLocation flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(deltaLocation), 0)
}
func PositionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
