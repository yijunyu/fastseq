// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Element_

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Anonymous0 struct {
	_tab flatbuffers.Table
}

func GetRootAsAnonymous0(buf []byte, offset flatbuffers.UOffsetT) *Anonymous0 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Anonymous0{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Anonymous0) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Anonymous0) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Anonymous0) SrcmlKind() SrcmlKind {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Anonymous0) MutateSrcmlKind(n SrcmlKind) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *Anonymous0) SmaliKind() SmaliKind {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Anonymous0) MutateSmaliKind(n SmaliKind) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *Anonymous0) Python3Kind() Python3Kind {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Anonymous0) MutatePython3Kind(n Python3Kind) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Anonymous0) SolidityKind() SolidityKind {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Anonymous0) MutateSolidityKind(n SolidityKind) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func (rcv *Anonymous0) Cobol85Kind() Cobol85Kind {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Anonymous0) MutateCobol85Kind(n Cobol85Kind) bool {
	return rcv._tab.MutateInt32Slot(12, n)
}

func Anonymous0Start(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func Anonymous0AddSrcmlKind(builder *flatbuffers.Builder, srcmlKind int32) {
	builder.PrependInt32Slot(0, srcmlKind, 0)
}
func Anonymous0AddSmaliKind(builder *flatbuffers.Builder, smaliKind int32) {
	builder.PrependInt32Slot(1, smaliKind, 0)
}
func Anonymous0AddPython3Kind(builder *flatbuffers.Builder, python3Kind int32) {
	builder.PrependInt32Slot(2, python3Kind, 0)
}
func Anonymous0AddSolidityKind(builder *flatbuffers.Builder, solidityKind int32) {
	builder.PrependInt32Slot(3, solidityKind, 0)
}
func Anonymous0AddCobol85Kind(builder *flatbuffers.Builder, cobol85Kind int32) {
	builder.PrependInt32Slot(4, cobol85Kind, 0)
}
func Anonymous0End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
