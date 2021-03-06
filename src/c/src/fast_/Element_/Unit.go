// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Element_

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

func (rcv *Unit) Revision() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Unit) Language() LanguageType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Unit) MutateLanguage(n LanguageType) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Unit) Item() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Unit) MutateItem(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func UnitStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func UnitAddFilename(builder *flatbuffers.Builder, filename flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(filename), 0)
}
func UnitAddRevision(builder *flatbuffers.Builder, revision flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(revision), 0)
}
func UnitAddLanguage(builder *flatbuffers.Builder, language int32) {
	builder.PrependInt32Slot(2, language, 0)
}
func UnitAddItem(builder *flatbuffers.Builder, item int32) {
	builder.PrependInt32Slot(3, item, 0)
}
func UnitEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
