// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Function_

import (
	flatbuffers "github.com/google/flatbuffers/go"

	fast___Slices___Slice___SourceFile___Function___Variable_ "fast_/Slices_/Slice_/SourceFile_/Function_/Variable_"
)

type Variable struct {
	_tab flatbuffers.Table
}

func GetRootAsVariable(buf []byte, offset flatbuffers.UOffsetT) *Variable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Variable{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Variable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Variable) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Variable) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Variable) Pos(obj *fast___Slices___Slice___SourceFile___Function___Variable_.Position) *fast___Slices___Slice___SourceFile___Function___Variable_.Position {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(fast___Slices___Slice___SourceFile___Function___Variable_.Position)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Variable) Type() ChangeType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Variable) MutateType(n ChangeType) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *Variable) Defn(obj *fast___Slices___Slice___SourceFile___Function___Variable_.Position, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Variable) DefnLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Variable) Use(obj *fast___Slices___Slice___SourceFile___Function___Variable_.Position, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Variable) UseLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Variable) Dvar(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Variable) DvarLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Variable) Alias(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Variable) AliasLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Variable) Cfunc(obj *fast___Slices___Slice___SourceFile___Function___Variable_.FunctionDecl, j int) bool {
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

func (rcv *Variable) CfuncLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func VariableStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func VariableAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func VariableAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(pos), 0)
}
func VariableAddType(builder *flatbuffers.Builder, type_ int32) {
	builder.PrependInt32Slot(2, type_, 0)
}
func VariableAddDefn(builder *flatbuffers.Builder, defn flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(defn), 0)
}
func VariableStartDefnVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func VariableAddUse(builder *flatbuffers.Builder, use flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(use), 0)
}
func VariableStartUseVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func VariableAddDvar(builder *flatbuffers.Builder, dvar flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(dvar), 0)
}
func VariableStartDvarVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func VariableAddAlias(builder *flatbuffers.Builder, alias flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(alias), 0)
}
func VariableStartAliasVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func VariableAddCfunc(builder *flatbuffers.Builder, cfunc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(cfunc), 0)
}
func VariableStartCfuncVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func VariableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}