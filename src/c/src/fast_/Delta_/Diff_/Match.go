// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Diff_

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Match struct {
	_tab flatbuffers.Table
}

func GetRootAsMatch(buf []byte, offset flatbuffers.UOffsetT) *Match {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Match{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Match) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Match) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Match) Src() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Match) MutateSrc(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *Match) Dst() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Match) MutateDst(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func MatchStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MatchAddSrc(builder *flatbuffers.Builder, src int32) {
	builder.PrependInt32Slot(0, src, 0)
}
func MatchAddDst(builder *flatbuffers.Builder, dst int32) {
	builder.PrependInt32Slot(1, dst, 0)
}
func MatchEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
