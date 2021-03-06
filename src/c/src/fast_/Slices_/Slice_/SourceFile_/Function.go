// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package SourceFile_

import (
	flatbuffers "github.com/google/flatbuffers/go"

	fast___Slices___Slice___SourceFile___Function_ "fast_/Slices_/Slice_/SourceFile_/Function_"
)

type Function struct {
	_tab flatbuffers.Table
}

func GetRootAsFunction(buf []byte, offset flatbuffers.UOffsetT) *Function {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Function{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Function) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Function) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Function) Variable(obj *fast___Slices___Slice___SourceFile___Function_.Variable, j int) bool {
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

func (rcv *Function) VariableLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Function) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Function) Type() ChangeType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Function) MutateType(n ChangeType) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func FunctionStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func FunctionAddVariable(builder *flatbuffers.Builder, variable flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(variable), 0)
}
func FunctionStartVariableVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func FunctionAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func FunctionAddType(builder *flatbuffers.Builder, type_ int32) {
	builder.PrependInt32Slot(2, type_, 0)
}
func FunctionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
