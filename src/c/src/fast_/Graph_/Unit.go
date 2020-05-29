// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Graph_

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Unit struct {
	_tab flatbuffers.Table
}

func GetRootAsUnit(buf []byte, offset flatbuffers.UOffsetT) *Unit {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Unit{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Unit) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Unit) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Unit) Filename() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Unit) SlotTokenIdx() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Unit) MutateSlotTokenIdx(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *Unit) Graph(obj *ContextGraph) *ContextGraph {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ContextGraph)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Unit) SlotDummyNode() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Unit) MutateSlotDummyNode(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *Unit) SymbolCandidate(obj *SymbolCandidate, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Unit) SymbolCandidateLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func UnitStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func UnitAddFilename(builder *flatbuffers.Builder, filename flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(filename), 0)
}
func UnitAddSlotTokenIdx(builder *flatbuffers.Builder, slotTokenIdx int32) {
	builder.PrependInt32Slot(1, slotTokenIdx, 0)
}
func UnitAddGraph(builder *flatbuffers.Builder, graph flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(graph), 0)
}
func UnitAddSlotDummyNode(builder *flatbuffers.Builder, SlotDummyNode int32) {
	builder.PrependInt32Slot(3, SlotDummyNode, 0)
}
func UnitAddSymbolCandidate(builder *flatbuffers.Builder, symbolCandidate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(symbolCandidate), 0)
}
func UnitStartSymbolCandidateVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func UnitEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
