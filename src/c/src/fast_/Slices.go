// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fast_

import (
	flatbuffers "github.com/google/flatbuffers/go"

	fast___Slices_ "fast_/Slices_"
)

type Slices struct {
	_tab flatbuffers.Table
}

func GetRootAsSlices(buf []byte, offset flatbuffers.UOffsetT) *Slices {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Slices{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Slices) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Slices) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Slices) Slice(obj *fast___Slices_.Slice, j int) bool {
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

func (rcv *Slices) SliceLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func SlicesStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SlicesAddSlice(builder *flatbuffers.Builder, slice flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(slice), 0)
}
func SlicesStartSliceVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SlicesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
