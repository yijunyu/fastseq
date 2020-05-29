// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Slice_

import (
	flatbuffers "github.com/google/flatbuffers/go"

	fast___Slices___Slice___SourceFile_ "fast_/Slices_/Slice_/SourceFile_"
)

type SourceFile struct {
	_tab flatbuffers.Table
}

func GetRootAsSourceFile(buf []byte, offset flatbuffers.UOffsetT) *SourceFile {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SourceFile{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SourceFile) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SourceFile) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SourceFile) Function(obj *fast___Slices___Slice___SourceFile_.Function, j int) bool {
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

func (rcv *SourceFile) FunctionLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SourceFile) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *SourceFile) Type() ChangeType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SourceFile) MutateType(n ChangeType) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func SourceFileStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func SourceFileAddFunction(builder *flatbuffers.Builder, function flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(function), 0)
}
func SourceFileStartFunctionVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SourceFileAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func SourceFileAddType(builder *flatbuffers.Builder, type_ int32) {
	builder.PrependInt32Slot(2, type_, 0)
}
func SourceFileEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}