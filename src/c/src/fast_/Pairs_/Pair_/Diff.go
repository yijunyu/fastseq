// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Pair_

import (
	flatbuffers "github.com/google/flatbuffers/go"

	fast_ "fast_"
)

type Diff struct {
	_tab flatbuffers.Table
}

func GetRootAsDiff(buf []byte, offset flatbuffers.UOffsetT) *Diff {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Diff{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Diff) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Diff) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Diff) Project() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Diff) LeftLine() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Diff) MutateLeftLine(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *Diff) LeftColumn() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Diff) MutateLeftColumn(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Diff) RightLine() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Diff) MutateRightLine(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *Diff) RightColumn() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Diff) MutateRightColumn(n int32) bool {
	return rcv._tab.MutateInt32Slot(12, n)
}

func (rcv *Diff) OldCode(obj *fast_.Element) *fast_.Element {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(fast_.Element)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Diff) NewCode(obj *fast_.Element) *fast_.Element {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(fast_.Element)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Diff) Hash() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Diff) Slices(obj *fast_.Slices) *fast_.Slices {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(fast_.Slices)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DiffStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
}
func DiffAddProject(builder *flatbuffers.Builder, project flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(project), 0)
}
func DiffAddLeftLine(builder *flatbuffers.Builder, leftLine int32) {
	builder.PrependInt32Slot(1, leftLine, 0)
}
func DiffAddLeftColumn(builder *flatbuffers.Builder, leftColumn int32) {
	builder.PrependInt32Slot(2, leftColumn, 0)
}
func DiffAddRightLine(builder *flatbuffers.Builder, rightLine int32) {
	builder.PrependInt32Slot(3, rightLine, 0)
}
func DiffAddRightColumn(builder *flatbuffers.Builder, rightColumn int32) {
	builder.PrependInt32Slot(4, rightColumn, 0)
}
func DiffAddOldCode(builder *flatbuffers.Builder, oldCode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(oldCode), 0)
}
func DiffAddNewCode(builder *flatbuffers.Builder, newCode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(newCode), 0)
}
func DiffAddHash(builder *flatbuffers.Builder, hash flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(hash), 0)
}
func DiffAddSlices(builder *flatbuffers.Builder, slices flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(slices), 0)
}
func DiffEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
