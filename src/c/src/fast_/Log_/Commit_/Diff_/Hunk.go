// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Diff_

import (
	flatbuffers "github.com/google/flatbuffers/go"

	fast_ "fast_"
	fast___Log___Commit___Diff___Hunk_ "fast_/Log_/Commit_/Diff_/Hunk_"
)

type Hunk struct {
	_tab flatbuffers.Table
}

func GetRootAsHunk(buf []byte, offset flatbuffers.UOffsetT) *Hunk {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Hunk{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Hunk) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Hunk) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Hunk) FromLineno() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Hunk) MutateFromLineno(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *Hunk) FromColumn() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Hunk) MutateFromColumn(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *Hunk) ToLineno() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Hunk) MutateToLineno(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Hunk) ToColumn() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Hunk) MutateToColumn(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *Hunk) Context() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Hunk) Element(obj *fast_.Element, j int) bool {
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

func (rcv *Hunk) ElementLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Hunk) Graph(obj *fast_.Graph, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Hunk) GraphLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Hunk) Mod(obj *fast___Log___Commit___Diff___Hunk_.ModLine, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Hunk) ModLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Hunk) Slice(obj *fast_.Slices) *fast_.Slices {
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

func HunkStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
}
func HunkAddFromLineno(builder *flatbuffers.Builder, fromLineno int32) {
	builder.PrependInt32Slot(0, fromLineno, 0)
}
func HunkAddFromColumn(builder *flatbuffers.Builder, fromColumn int32) {
	builder.PrependInt32Slot(1, fromColumn, 0)
}
func HunkAddToLineno(builder *flatbuffers.Builder, toLineno int32) {
	builder.PrependInt32Slot(2, toLineno, 0)
}
func HunkAddToColumn(builder *flatbuffers.Builder, toColumn int32) {
	builder.PrependInt32Slot(3, toColumn, 0)
}
func HunkAddContext(builder *flatbuffers.Builder, context flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(context), 0)
}
func HunkAddElement(builder *flatbuffers.Builder, element flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(element), 0)
}
func HunkStartElementVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func HunkAddGraph(builder *flatbuffers.Builder, graph flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(graph), 0)
}
func HunkStartGraphVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func HunkAddMod(builder *flatbuffers.Builder, mod flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(mod), 0)
}
func HunkStartModVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func HunkAddSlice(builder *flatbuffers.Builder, slice flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(slice), 0)
}
func HunkEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
