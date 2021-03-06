// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Bugs_

import (
	flatbuffers "github.com/google/flatbuffers/go"

	fast___Bugs___Bug_ "fast_/Bugs_/Bug_"
)

type Bug struct {
	_tab flatbuffers.Table
}

func GetRootAsBug(buf []byte, offset flatbuffers.UOffsetT) *Bug {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Bug{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Bug) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Bug) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Bug) Id() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Bug) Opendate() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Bug) Fixdate() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Bug) Buginfo(obj *fast___Bugs___Bug_.Info) *fast___Bugs___Bug_.Info {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(fast___Bugs___Bug_.Info)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Bug) FixedFile(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Bug) FixedFileLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func BugStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func BugAddId(builder *flatbuffers.Builder, id flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(id), 0)
}
func BugAddOpendate(builder *flatbuffers.Builder, opendate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(opendate), 0)
}
func BugAddFixdate(builder *flatbuffers.Builder, fixdate flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(fixdate), 0)
}
func BugAddBuginfo(builder *flatbuffers.Builder, buginfo flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(buginfo), 0)
}
func BugAddFixedFile(builder *flatbuffers.Builder, fixedFile flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(fixedFile), 0)
}
func BugStartFixedFileVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BugEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
