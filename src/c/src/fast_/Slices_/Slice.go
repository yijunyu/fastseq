// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Slices_

import (
	flatbuffers "github.com/google/flatbuffers/go"

	fast___Slices___Slice_ "fast_/Slices_/Slice_"
)

type Slice struct {
	_tab flatbuffers.Table
}

func GetRootAsSlice(buf []byte, offset flatbuffers.UOffsetT) *Slice {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Slice{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Slice) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Slice) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Slice) File(obj *fast___Slices___Slice_.SourceFile, j int) bool {
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

func (rcv *Slice) FileLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Slice) Hash() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func SliceStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SliceAddFile(builder *flatbuffers.Builder, file flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(file), 0)
}
func SliceStartFileVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SliceAddHash(builder *flatbuffers.Builder, hash flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(hash), 0)
}
func SliceEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
