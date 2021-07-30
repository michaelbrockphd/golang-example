package collections

import (
	"errors"
)

type NodeStack struct {
	length int

	topNode *SingleLinkedNode
}

func (ns *NodeStack) Length() int {
	return ns.length
}

func (ns *NodeStack) Top() interface{} {
	if ns.topNode != nil {
		return ns.topNode.value
	}

	return nil
}

func (ns *NodeStack) Push(i interface{}) {
	newTop := NewSingleLinkedNode()

	newTop.value = i

	newTop.next = ns.topNode

	ns.topNode = newTop

	ns.length += 1
}

func (ns *NodeStack) Pop() error {
	if ns.length < 1 {
		return errors.New("cannot pop from an empty stack")
	}

	oldTop := ns.topNode

	ns.topNode = oldTop.next

	oldTop.next = nil

	ns.length -= 1

	return nil
}

func NewNodeStack() Stack {
	rtn := new(NodeStack)

	rtn.length = 0

	rtn.topNode = nil

	return rtn
}
