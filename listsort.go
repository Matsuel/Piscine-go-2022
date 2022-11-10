package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	courant := l
	if courant == nil {
		return nil
	}
	courant.Next = ListSort(courant.Next)
	if courant.Next != nil && courant.Data > courant.Next.Data {
		courant = move(courant)
	}
	return courant
}

func move(l *NodeI) *NodeI {
	p := l
	n := l.Next
	ret := n
	for n != nil && l.Data > n.Data {
		p = n
		n = n.Next
	}
	p.Next = l
	l.Next = n
	return ret
}
