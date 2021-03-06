// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Graph_

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ContextGraph struct {
	_tab flatbuffers.Table
}

func GetRootAsContextGraph(buf []byte, offset flatbuffers.UOffsetT) *ContextGraph {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ContextGraph{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ContextGraph) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ContextGraph) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ContextGraph) Edges(obj *ContextEdges) *ContextEdges {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ContextEdges)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ContextGraph) NodeLabel(obj *NodeLabel, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ContextGraph) NodeLabelLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ContextGraph) NodeType(obj *NodeType, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ContextGraph) NodeTypeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ContextGraphStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ContextGraphAddEdges(builder *flatbuffers.Builder, edges flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(edges), 0)
}
func ContextGraphAddNodeLabel(builder *flatbuffers.Builder, nodeLabel flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(nodeLabel), 0)
}
func ContextGraphStartNodeLabelVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ContextGraphAddNodeType(builder *flatbuffers.Builder, nodeType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(nodeType), 0)
}
func ContextGraphStartNodeTypeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ContextGraphEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
