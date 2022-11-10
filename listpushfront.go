package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {
	val := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = val
		l.Head = l.Head
	} else {
		newNode := val
		newNode.Next, l.Head = l.Head, newNode
	}
}
