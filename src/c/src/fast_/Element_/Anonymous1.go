// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Element_

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Anonymous1 struct {
	_tab flatbuffers.Table
}

func GetRootAsAnonymous1(buf []byte, offset flatbuffers.UOffsetT) *Anonymous1 {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Anonymous1{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Anonymous1) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Anonymous1) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Anonymous1) Unit(obj *Unit) *Unit {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Unit)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Anonymous1) Literal(obj *Literal) *Literal {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Literal)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func Anonymous1Start(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Anonymous1AddUnit(builder *flatbuffers.Builder, unit flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(unit), 0)
}
func Anonymous1AddLiteral(builder *flatbuffers.Builder, literal flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(literal), 0)
}
func Anonymous1End(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
