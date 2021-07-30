package collections

type SingleLinkedNode struct {
	value interface{}

	next *SingleLinkedNode
}

func NewSingleLinkedNode() *SingleLinkedNode {
	return new(SingleLinkedNode)
}
